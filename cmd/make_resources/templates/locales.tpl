// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
    "{{.CLDRPackage}}"
    "github.com/razor-1/localizer/resources/currency"
)

func getLocale_{{.LocaleCode}}() *cldr.Locale {
    return &cldr.Locale {
        Calendar: {{printf "%#v" .Calendar}},
        Plural: cldr.Plural{Cardinal: LocalePlural[tag_{{.PluralLocaleCode}}]()},
        Number: cldr.Number{
            Symbols: {{printf "%#v" .Symbols}},
            Formats: {{printf "%#v" .NumberFormats}},
            Currencies: cldr.Currencies{
                {{range $curCode, $currency := .Currencies -}}
                currency.{{$curCode}}: {{printf "%#v" $currency}},
                {{end}}
            },
        },
    }
}
