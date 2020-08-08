// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
    "golang.org/x/text/feature/plural"
    "golang.org/x/text/language"

    "{{.CLDRPackage}}"
)

type PluralForms []plural.Form
type pluralFunc func(ops *cldr.PluralOperands) plural.Form

{{range $i, $pg := .PluralGroups}}
func cardinal_{{$i}}() cldr.PluralData {
    // for locales: {{printf "%v" $pg.SplitLocales}}
    forms := PluralForms{
        {{- range $pg.PluralRules}}
            plural.{{.CountTitle}},
        {{- end}}
    }
    plFunc := func (ops *cldr.PluralOperands) plural.Form {
        {{- range $pg.PluralRules}}{{if .GoCondition}}
        // {{.Condition}}
        if {{.GoCondition}} {
            return plural.{{.CountTitle}}
        }{{end}}
        {{- end}}
        return plural.Other
    }

    return cldr.PluralData{Forms: forms, Func: plFunc}
}
{{end}}

var LocalePlural = map[language.Tag]func() cldr.PluralData{
{{- range $locale, $funcName := .LocaleMap}}
    tag_{{$locale}}: {{$funcName}},
{{- end}}
}

func intInRange(i, from, to int64) bool {
    return from <= i && i <= to
}

func intEqualsAny(i int64, any ...int64) bool {
    for _, a := range any {
        if i == a {
            return true
        }
    }
    return false
}
