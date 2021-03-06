// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
    "{{.CLDRPackage}}"
    "{{.CLDRPackage}}/resources/currency"
    {{if .Territories}}"{{.CLDRPackage}}/resources/territory"{{end}}
    {{if .Languages}}"{{.CLDRPackage}}/resources/language"{{end}}
)

func Get_{{.LocaleCode}}() *cldr.Locale {
    return &cldr.Locale {
        Locale: "{{.LocaleCode}}",
        Calendar: {{printf "%#v" .Calendar}},
        Number: cldr.Number{
            Symbols: {{printf "%#v" .Symbols}},
            Formats: {{printf "%#v" .NumberFormats}},
            Currencies: cldr.Currencies{
                {{range $curCode, $currency := .Currencies -}}
                currency.{{$curCode}}: {{printf "%#v" $currency}},
                {{end}}
            },
        },
        Display: {{printf "%#v" .LocaleDisplayPattern}},
        Languages: cldr.Languages{
            {{range $lang, $locLang := .Languages -}}
            language.{{$lang | ToUpperIdent}}: {{printf "%#v" $locLang}},
            {{end}}
        },
        Territories: cldr.Territories{
            {{range $ter, $locTer := .Territories -}}
            territory.{{$ter | ToUpperIdent}}: {{printf "%#v" $locTer}},
            {{end}}
        },
    }
}
