package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/helpers"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/recipes"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
	pluginschema "github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

const useFullStruct = "." // special case to use full struct instead of a member

//go:embed templates/*.go.tpl
var awsTemplatesFS embed.FS

func main() {
	resources := recipes.AllResources

	for _, r := range resources {
		generateResource(r, false)
		generateResource(r, true)
	}

	for i, r := range resources {
		if r.Parent != nil {
			r.Parent.Table.Relations = append(r.Parent.Table.Relations, r.Table)
			resources[i] = nil
		}
	}
	relationalRes := make([]*recipes.Resource, 0, len(resources))
	for _, r := range resources {
		if r != nil {
			relationalRes = append(relationalRes, r)
		}
	}

	generateProvider(relationalRes)
}

func inferFromRecipe(r *recipes.Resource) {
	var (
		res, items, pag, pget *helpers.InferResult
	)

	if r.ItemsStruct != nil {
		items = helpers.InferFromStructOutput(r.ItemsStruct)
		r.GetMethod = items.Method
		if helpers.BareType(reflect.TypeOf(r.AWSStruct)) == helpers.BareType(reflect.TypeOf(r.ItemsStruct)) {
			// Same type, so we need to save the output struct completely
			r.ResponseItemsName = useFullStruct
		} else {
			needSingular := r.PaginatorStruct != nil
			if !needSingular && len(items.ItemsFieldCandidates(false)) == 0 {
				// we have wrapper? look for a single item
				singleItemMaybeWrapper := items.ItemsFieldCandidates(true)
				if len(singleItemMaybeWrapper) != 1 {
					log.Fatal("Could not determine possible ItemsName wrapper for ", items.Method, ":", len(singleItemMaybeWrapper), " candidates")
				}
				r.ResponseItemsWrapper = singleItemMaybeWrapper[0].Name
				again := helpers.InferFromType(helpers.BareType(singleItemMaybeWrapper[0].Type), "")
				if again.PaginatorTokenField != nil {
					r.WrappedNextTokenName = again.PaginatorTokenField.Name
				}

			} else {
				r.ResponseItemsName = items.ItemsField(needSingular).Name
			}
		}

		if items.PaginatorTokenField != nil {
			r.NextTokenName = items.PaginatorTokenField.Name
		}

		res = items
	}

	var pagSingleItem reflect.Type // This is the single item from the paginator, which is used later to match fields.

	if r.PaginatorStruct != nil {
		pag = helpers.InferFromStructOutput(r.PaginatorStruct)
		var f reflect.StructField
		// we need a slice field, but what if it's wrapped?
		if len(pag.ItemsFieldCandidates(false)) == 0 {
			// we have wrapper? look for a single item
			singleItemMaybeWrapper := pag.ItemsFieldCandidates(true)
			if len(singleItemMaybeWrapper) != 1 {
				log.Fatal("Could not determine possible PaginatorStruct wrapper for ", pag.Method, ":", len(singleItemMaybeWrapper), " candidates")
			}
			r.PaginatorListWrapper = singleItemMaybeWrapper[0].Name

			if singleItemMaybeWrapper[0].Type.Kind() == reflect.Ptr {
				r.PaginatorListWrapperType = "&" + path.Base(singleItemMaybeWrapper[0].Type.Elem().PkgPath()) + "." + singleItemMaybeWrapper[0].Type.Elem().Name()
			} else {
				r.PaginatorListWrapperType = path.Base(singleItemMaybeWrapper[0].Type.PkgPath()) + "." + singleItemMaybeWrapper[0].Type.Name()
			}
			again := helpers.InferFromType(helpers.BareType(singleItemMaybeWrapper[0].Type), "")
			if again.PaginatorTokenField != nil {
				r.WrappedNextTokenName = again.PaginatorTokenField.Name
			}
			f = again.ItemsField(false)
		} else {
			f = pag.ItemsField(false)
		}

		pagSingleItem = f.Type.Elem() // get type of slice

		r.PaginatorListName = f.Name
		r.PaginatorListType = f.Type.Elem().Name() // single type from a slice
		if f.Type.Elem().Kind() == reflect.Struct {
			r.PaginatorListType = "types." + r.PaginatorListType
		}

		r.ListMethod = pag.Method

		if res == nil {
			res = pag
		}
	}

	if r.PaginatorGetStruct != nil {
		if r.ItemsStruct == nil {
			log.Fatal("PaginatorGetStruct requires ItemsStruct on resource ", r.AWSService)
		}

		pget = helpers.InferFromStructInput(r.PaginatorGetStruct)
		if pget.Method != items.Method {
			log.Fatal("PaginatorGetStruct method ", pget.Method, " does not match ItemsStruct method ", items.Method)
		}

		if pag != nil {
			// figure out which fields match to what

			r.AutoCalculated.GetAndListOrder = nil
			r.AutoCalculated.MatchedGetAndListFields = make(map[string]string)

			fields := make(map[string]reflect.Type)
			//log.Println("PROCESSING", pagSingleItem.Name(), pagSingleItem.Kind().String())
			if k := pagSingleItem.Kind(); k == reflect.String {
				// special case for string
				fields[""] = pagSingleItem
			} else {
				for i := 0; i < pagSingleItem.NumField(); i++ {
					f := pagSingleItem.Field(i)
					if f.Name == "noSmithyDocumentSerde" || f.Type.String() == "document.NoSerde" {
						continue
					}
					fields[f.Name] = f.Type
				}
			}

			if len(fields) == 1 && fields[""] != nil {
				// special case for string (not struct)
				found := false
				for _, f := range pget.FieldOrder {
					ff := pget.Fields[f].Type
					if helpers.BareType(ff).Kind() == helpers.BareType(fields[""]).Kind() {
						found = true
						r.AutoCalculated.GetAndListOrder = append(r.AutoCalculated.GetAndListOrder, f)
						r.AutoCalculated.MatchedGetAndListFields[f] = "&item"
						break
					}
				}
				if !found {
					log.Println("PaginatorGetStruct field of type", fields[""].Kind().String(), "not matched in PaginatorStruct in", pagSingleItem.Name())
				}
			} else {
				for _, f := range pget.FieldOrder {
					found := false
					nameMatchFn := func(a, b string) bool { return strings.ToLower(a) == strings.ToLower(b) }

					for attempts := 0; attempts < 2; attempts++ {
						for n, t := range fields {
							if nameMatchFn(n, f) && helpers.BareType(t) == helpers.BareType(pget.Fields[f].Type) {
								found = true
								r.AutoCalculated.GetAndListOrder = append(r.AutoCalculated.GetAndListOrder, f)
								r.AutoCalculated.MatchedGetAndListFields[f] = "item." + n
								break
							}
						}
						if found {
							break
						}
						if !found {
							if attempts == 0 {
								// Either suffix or single field
								nameMatchFn = func(a, b string) bool {
									return (len(pget.FieldOrder) == 1 && a == "Name") || strings.HasSuffix(a, b)
								}

								log.Println("PaginatorGetStruct field", f, "not matched in PaginatorStruct in", pagSingleItem.Name(), "doing heuristic match")
							} else {
								log.Println("PaginatorGetStruct field", f, "not matched in PaginatorStruct in", pagSingleItem.Name(), "even after heuristic match")
							}
						}
					}
				}
			}

			if len(r.AutoCalculated.GetAndListOrder) > 0 {
				log.Println("GetAndListOrder for", pagSingleItem.Name()+":", r.AutoCalculated.GetAndListOrder)
			}
		}

	}

	if items != nil && pag != nil {
		r.AWSSubService = pag.SubService
	} else {
		r.AWSSubService = res.SubService
	}

	if items != nil && pag != nil && strings.TrimSuffix(pag.SubService, "s") != strings.TrimSuffix(items.SubService, "s") { // Certificate vs Certificates
		log.Println("Mismatching subservices between ItemsStruct and PaginatorStruct for resource ", r.AWSService, ": ", items.SubService, " vs ", pag.SubService)
	}

}

func generateResource(r *recipes.Resource, mock bool) {
	var err error
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)

	if r.ItemsStruct != nil {
		inferFromRecipe(r)
	}

	tableNameFromSubService, fetcherNameFromSubService := helpers.TableAndFetcherNames(r)

	r.Table, err = sdkgen.NewTableFromStruct(
		fmt.Sprintf("aws_%s_%s", strings.ToLower(r.AWSService), tableNameFromSubService),
		r.AWSStruct,
		sdkgen.WithSkipFields(append(r.SkipFields, "noSmithyDocumentSerde")),
	)
	if err != nil {
		log.Fatal(err)
	}

	if r.TrimPrefix != "" {
		for i := range r.Table.Columns {
			r.Table.Columns[i].Name = strings.TrimPrefix(r.Table.Columns[i].Name, r.TrimPrefix)
		}
	}

	r.Table.Columns = append(r.DefaultColumns, r.Table.Columns...)
	if r.ColumnOverrides != nil {
		for i, c := range r.Table.Columns {
			override, ok := r.ColumnOverrides[c.Name]
			if !ok {
				continue
			}
			r.Table.Columns[i].Name = helpers.Coalesce(override.Name, r.Table.Columns[i].Name)
			r.Table.Columns[i].Resolver = helpers.Coalesce(override.Resolver, r.Table.Columns[i].Resolver)
			r.Table.Columns[i].Description = helpers.Coalesce(override.Description, r.Table.Columns[i].Description)

			delete(r.ColumnOverrides, c.Name)
		}
		coSlice := make([]string, 0, len(r.ColumnOverrides))
		for k := range r.ColumnOverrides {
			coSlice = append(coSlice, k)
		}
		sort.Strings(coSlice)
		// remaining, unmatched columns are added to the end of the table. Difference from DefaultColumns? none for now
		for _, k := range coSlice {
			c := r.ColumnOverrides[k]
			if c.Type == pluginschema.TypeInvalid {
				if !mock {
					fmt.Println("Not adding unmatched column with unspecified type", k, c)
				}
				continue
			}
			c.Name = helpers.Coalesce(c.Name, k)
			r.Table.Columns = append(r.Table.Columns, c)
		}
	}

	hasReferenceToResolvers := false

	for i := range r.Table.Columns {
		if len(r.Table.Options.PrimaryKeys) == 0 && r.Table.Columns[i].Name == "arn" {
			//	r.Table.Columns[i].Options.PrimaryKey = true
			r.Table.Options.PrimaryKeys = []string{"arn"}
		}
		if r.Table.Columns[i].Name == "tags" {
			r.HasTags = true

			if r.Table.Columns[i].Resolver == recipes.ResolverAuto {
				r.Table.Columns[i].Resolver = "resolve" + r.AWSService + r.AWSSubService + "Tags"
			}
		}
		if strings.HasPrefix(r.Table.Columns[i].Resolver, "resolvers.") {
			hasReferenceToResolvers = true
		}
	}
	r.Table.Multiplex = `client.ServiceAccountRegionMultiplexer("` + helpers.Coalesce(r.MultiplexerServiceOverride, strings.ToLower(r.AWSService)) + `")`

	r.Table.Resolver = "fetch" + r.AWSService + fetcherNameFromSubService
	r.TableFuncName = strings.TrimPrefix(r.Table.Resolver, "fetch")

	if mock {
		r.MockFuncName = "build" + r.TableFuncName
		r.TestFuncName = "Test" + r.TableFuncName
	}

	t := reflect.TypeOf(r.AWSStruct).Elem()
	r.AWSStructName = path.Base(t.PkgPath()) + "." + t.Name() // types.Something or sometimes service.Something
	r.ItemName = helpers.Coalesce(r.ItemName, t.Name())

	if !mock {
		r.Imports = quoteImports(r.Imports)
	} else {
		r.MockImports = quoteImports(r.MockImports)
	}

	r.TypesImport = ""
	sp := t.PkgPath()
	if strings.HasSuffix(sp, "/types") {
		r.TypesImport = sp
		if r.AddTypesImport {
			if !mock {
				r.Imports = append(r.Imports, strconv.Quote(sp))
			} else {
				r.MockImports = append(r.MockImports, strconv.Quote(sp))
			}
		}
		r.Imports = append(r.Imports, strconv.Quote(strings.TrimSuffix(sp, "/types")))         // auto-import main pkg (not "types")
		r.MockImports = append(r.MockImports, strconv.Quote(strings.TrimSuffix(sp, "/types"))) // auto-import main pkg (not "types")
	} else if strings.HasSuffix(sp, "/aws-sdk-go-v2/service/"+strings.ToLower(r.AWSService)) { // main struct lives in main pkg, auto-import that as well
		r.Imports = append(r.Imports, strconv.Quote(sp))
		r.MockImports = append(r.MockImports, strconv.Quote(sp))
	}

	if hasReferenceToResolvers && !mock {
		res := "resolvers " + strconv.Quote(`github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/`+strings.ToLower(r.AWSService))
		if !slices.Contains(r.Imports, res) {
			r.Imports = append(r.Imports, res)
		}
	}

	r.TemplateFilename = r.Template + helpers.StringSwitch(mock, "_mock_test", "") + ".go.tpl"
	tpl, err := template.New(r.TemplateFilename).Funcs(template.FuncMap{
		"ToCamel":  strcase.ToCamel,
		"ToLower":  strings.ToLower,
		"ToSnake":  strcase.ToSnake,
		"Coalesce": func(a1, a2 string) string { return helpers.Coalesce(a2, a1) }, // go templates argument order is backwards
	}).ParseFS(awsTemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse aws templates: %w", err))
	}
	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse codegen template: %w", err))
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	filePath := path.Join(dir, "../codegen", strings.ToLower(r.AWSService))
	if err := os.MkdirAll(filePath, 0755); err != nil {
		log.Fatal(fmt.Errorf("failed to create directory: %w", err))
	}

	fileSuffix := helpers.StringSwitch(mock, "_mock_test.go", ".go")
	filePath = path.Join(filePath, strings.ToLower(r.AWSService)+"_"+tableNameFromSubService+fileSuffix)
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func generateProvider(rr []*recipes.Resource) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)
	tpl, err := template.New("provider.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
		"ToSnake": strcase.ToSnake,
	}).ParseFS(awsTemplatesFS, "templates/provider.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse provider.go.tpl: %w", err))
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, rr); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}

	filePath := path.Join(dir, "../resources/provider/provider_codegen.go")
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func quoteImports(imports []string) []string {
	for i := range imports {
		if !strings.HasSuffix(imports[i], `"`) {
			imports[i] = strconv.Quote(imports[i])
		}
	}
	return imports
}