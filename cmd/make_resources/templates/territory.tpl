// Code generated by make_resources.go; DO NOT EDIT.
package territory

const (
{{- range $ter, $_ := .}}
    {{$ter | ToUpperIdent}} = "{{$ter}}"
{{- end}}
)
