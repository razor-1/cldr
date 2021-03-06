// Code generated by make_resources.go; DO NOT EDIT.
package resources

import (
    "testing"

    "golang.org/x/text/feature/plural"
    "golang.org/x/text/language"
)

{{range $i, $pg := .PluralGroups}}
    func Test{{$.PluralType}}_{{$i}}(t *testing.T) {
    var tests []pluralFormTest
    {{range $pg.PluralRules}}
        {{if .IntegerExamples}}tests = appendIntegerTests(tests, plural.{{.CountTitle}}, {{printf "%#v" .IntegerExamples}}){{end}}
        {{if .DecimalExamples}}tests = appendDecimalTests(tests, plural.{{.CountTitle}}, {{printf "%#v" .DecimalExamples}}){{end}}
    {{end}}
    tags := []language.Tag{
        {{ range $pg.SplitLocales -}}
            tag_{{.}},
        {{- end}}
        }
    for _, tag := range tags {
        runTests(t, tag, tests)
    }
    }
{{end}}