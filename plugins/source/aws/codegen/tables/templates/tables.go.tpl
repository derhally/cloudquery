// Code generated by codegen; DO NOT EDIT.

package plugin

import (
	"github.com/cloudquery/plugin-sdk/schema"
{{- range $resource := . }}
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/{{ $resource.Service }}"
{{- end }}
)

func tables() []*schema.Table {
	return []*schema.Table{
    {{- range $resource := . }}
        {{ $resource.Service }}.{{ $resource.SubService | ToCamel }}(),
    {{- end }}
	}
}
