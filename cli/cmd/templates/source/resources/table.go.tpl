package resources

import (
	"github.com/{{.Org}}/cq-source-{{.Name}}/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:     "{{.Name}}_sample_table",
		Resolver: fetchSampleTable,
		Columns: []schema.Column{
      {
        Name: "column",
        Type: schema.String,
      },
    },
  }
}

func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
  return fmt.Errorf("not implemented")
}