// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_vai_Vaii_LR() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "vai_Vaii_LR",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ꖨꖕꔞ", Feb: "ꕒꕡ", Mar: "ꕾꖺ", Apr: "ꖢꖕ", May: "ꖑꕱ", Jun: "ꖱꘋ", Jul: "ꖱꕞ", Aug: "ꗛꔕ", Sep: "ꕢꕌ", Oct: "ꕭꖃ", Nov: "ꔞꘋ", Dec: "ꖨꖕꗏ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ꖨꖕ ꕪꕴ ꔞꔀꕮꕊ", Feb: "ꕒꕡꖝꖕ", Mar: "ꕾꖺ", Apr: "ꖢꖕ", May: "ꖑꕱ", Jun: "ꖱꘋ", Jul: "ꖱꕞꔤ", Aug: "ꗛꔕ", Sep: "ꕢꕌ", Oct: "ꕭꖃ", Nov: "ꔞꘋꕔꕿ ꕸꖃꗏ", Dec: "ꖨꖕ ꕪꕴ ꗏꖺꕮꕊ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ꕞꕌꔵ", Mon: "ꗳꗡꘉ", Tue: "ꕚꕞꕚ", Wed: "ꕉꕞꕒ", Thu: "ꕉꔤꕆꕢ", Fri: "ꕉꔤꕀꕮ", Sat: "ꔻꔬꔳ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "ꖳꕯꔤꗳ ꕉꕟꔬ ꗡꕆꔓꔻ ꔵꕌꕆ", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "ꕉꖐꕞ ꖴꕎꘋꕤ", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "ꖺꔻꖤꔃꔷꕩ ꕜꕞꕌ", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "ꕑꗸꘋ", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "ꖜꖩꔺ ꖢꕟꘋꕃ", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "ꕷꖬꕎꕯ ꖛꕞ", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "ꕪꕯꕜ ꕜꕞꕌ", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "ꖏꖐꕱ ꖢꕟꘋꕃ", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "ꖬꔃꕤ ꖨꕮꕊ ꖢꕟꘋꕃ", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "ꕦꕇꔧ ꖳꕎꘋ ꔓꕆꘋꔬ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "ꗡꔻꖴꖁ ꕪꕷꗲꗡꔵꕩꖆ", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "ꕀꖜꔳ ꖢꕟꘋꕃ", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "ꕉꔷꕀꔸꕩ ꔵꕯ", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "ꕆꔻꕞ ꗁꖻꘋ", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "ꔀꔸꔳꕟ ꗁꖻꘋ", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ꔤꕿꖎꔪꕩ ꔫꔤ", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "ꖳꖄ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "ꔛꔟꔻ ꗁꖻꘋ ꔻꗳꔷꘋ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "ꕭꕌꕯ ꔻꔵ", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "ꕭꔭꕩ ꕜꕞꔻ", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "ꕅꔤꕇ ꖢꕟꘋꕃ", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ꔤꔺꕩ ꖩꔪ", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "ꕧꕐꕇꔧ ꘂꘋ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ꔞꕰ ꔻꔝꘋ", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "ꖏꖒꖄ ꖢꕟꘋꕃ", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "ꕞꔤꔫꕩ ꕜꕞꕌ", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "ꔷꖇꕿ ꖃꔳ", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "ꔷꔫꕩ ꔵꕯ", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "ꗞꕟꖏ ꔵꕌꕆ", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "ꕮꕞꕭꕌꔻ ꕉꔸꕩꔸ", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "ꗞꔸꕚꕇꕰ ꖳꕅꕩ (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ꗞꔸꕚꕇꕰ ꖳꕅꕩ", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "ꗞꔓꗔ ꖩꔪ", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "ꕮꕞꕌꔨ ꖴꕎꕦ", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "ꗞꕤꔭꕃ ꕆꔳꕪ", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "ꕯꕆꔫꕩ ꕜꕞꕌ", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "ꕯꔤꕀꔸꕩ ꕯꔤꕟ", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "ꕟꖙꕡ ꖢꕟꘋꕃ", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "ꕢꖙꔵ ꔸꕩꔷ", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "ꔖꗼꔷ ꖩꔪ", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "ꖬꗵꘋ ꗁꖻꘋ", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "ꔻꘋ ꗥꔷꕯ ꗁꖻꘋ", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "ꔷꗚꘋ", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "ꔷꗚꘋ (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "ꖇꕮꔷ ꔻꔝꘋ", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "ꕢꕴ ꕿꔈ ꗪ ꕉ ꕗꕴ ꖁꖜꕟ (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "ꕢꕴ ꕿꔈ ꗪ ꕉ ꕗꕴ ꖁꖜꕟ", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "ꔷꕞꔟꕇ", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "ꖤꕇꔻꕩ ꔵꕯ", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "ꕚꘋꕤꕇꕰ ꔻꔝꘋ", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "ꖳꕭꕡ ꔻꔝꘋ", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "ꕶꕱ ꕜꕞ", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "ꕉꔱꔸꕪ ꗛꔤ ꔒꘋꗣ ꗏ ꕟꘋꔵ", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "ꕤꔭꕩ ꖴꕎꕦ (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "ꕤꔭꕩ ꖴꕎꕦ", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "ꔽꕓꖜꔃ ꕜꕞ", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "ꕉꕪꘋ",
			language.AM:  "ꕉꕆꕌꔸ",
			language.AR:  "ꕞꕌꖝ",
			language.BE:  "ꔆꕞꖩꔻ",
			language.BG:  "ꗂꔠꗸꘋ",
			language.BN:  "ꗩꕭꔷ",
			language.CS:  "ꗿꗡ",
			language.DE:  "ꕧꕮꔧ",
			language.EL:  "ꗥꗷꘋ",
			language.EN:  "ꕶꕱ",
			language.ES:  "ꕐꘊꔧ",
			language.FA:  "ꗨꗡꔻꘂꘋ",
			language.FR:  "ꗱꘋꔻ",
			language.HA:  "ꕌꖙꕢ",
			language.HI:  "ꔦꔺ",
			language.HU:  "ꖽꔟꗸꘋ",
			language.ID:  "ꔤꖆꕇꔻꘂꘋ",
			language.IG:  "ꔤꕼ",
			language.IT:  "ꔤꕚꔷꘂꘋ",
			language.JA:  "ꕧꕐꕇꔧ",
			language.JV:  "ꕧꕙꕇꔧ",
			language.KM:  "ꕃꘈꗢ",
			language.KO:  "ꖏꔸꘂꘋ",
			language.MS:  "ꕮꔒꔀ",
			language.MY:  "ꗩꕆꔻ",
			language.NE:  "ꕇꕐꔷ",
			language.NL:  "ꗍꔿ",
			language.PA:  "ꖛꕨꔬ",
			language.PL:  "ꗁꔒꔻ",
			language.PT:  "ꕶꕿꕃꔤ",
			language.RO:  "ꖄꕆꕇꘂꘋ",
			language.RU:  "ꗐꖺꔻꘂꘋ",
			language.RW:  "ꕟꖙꕡ",
			language.SO:  "ꖇꕮꔷ",
			language.SV:  "ꖬꔨꗵꘋ",
			language.TA:  "ꕚꕆꔷ",
			language.TH:  "ꕚꔤ",
			language.TR:  "ꗋꕃ",
			language.UK:  "ꖳꖴꔓꕇꘂꘋ",
			language.UR:  "ꖺꖦ",
			language.VAI: "ꕙꔤ",
			language.VI:  "ꔲꕩꕯꕆꔧ",
			language.YO:  "ꖎꖄꕑ",
			language.ZH:  "ꕦꕇꔧ",
			language.ZU:  "ꖮꖨ",
		},
		Territories: cldr.Territories{
			territory.AC: "ꗻꗡ ꕒꕡꕌ ꗏ ꔳꘋꗣ",
			territory.AD: "ꕉꖆꕟ",
			territory.AE: "ꖳꕯꔤꗳ ꕉꕟꔬ ꗡꕆꔓꔻ",
			territory.AF: "ꕉꔱꕭꔕꔻꕚꘋ",
			territory.AG: "ꕉꘋꔳꖶꕎ ꗪ ꕑꖜꕜ",
			territory.AI: "ꕉꕄꕞ",
			territory.AL: "ꕉꔷꕑꕇꕩ",
			territory.AM: "ꕉꕆꕯ",
			territory.AO: "ꕉꖐꕞ",
			territory.AQ: "ꕉꘋꕚꔳꕪ",
			territory.AR: "ꕉꘀꘋꔳꕯ",
			territory.AS: "ꕶꕱ ꕢꕹꕎ",
			territory.AT: "ꖺꔻꖤꕎ",
			territory.AU: "ꖺꖬꖤꔃꔷꕩ",
			territory.AW: "ꕉꖩꕑ",
			territory.AX: "ꕉꕞꔺ",
			territory.AZ: "ꕉꕤꕑꔤꕧꘋ",
			territory.BA: "ꕷꔻꕇꕰ ꗪ ꗥꕤꖑꔲꕯ",
			territory.BB: "ꕑꔆꖁꔻ",
			territory.BD: "ꕑꕅꕞꗵꔼ",
			territory.BE: "ꗩꕀꗚꘋ",
			territory.BF: "ꕷꕃꕯ ꕘꖇ",
			territory.BG: "ꗂꔠꔸꕩ",
			territory.BH: "ꕑꗸꘋ",
			territory.BI: "ꖜꖩꔺ",
			territory.BJ: "ꗩꕇꘋ",
			territory.BL: "ꕪꘋꕓ ꗞꗢ ꕒꕚꕞꕆ",
			territory.BM: "ꗩꖷꕜ",
			territory.BN: "ꖜꖩꘉꔧ",
			territory.BO: "ꕷꔷꔲꕩ",
			territory.BQ: "ꕪꔓꔬꘂꘋ ꖨꕮ ꗨꗳꗣ",
			territory.BR: "ꖜꕟꔘꔀ",
			territory.BS: "ꕑꕌꕮꔻ",
			territory.BT: "ꖜꕚꘋ",
			territory.BV: "ꖜꔍꔳ ꔳꘋꗣ",
			territory.BW: "ꕷꖬꕎꕯ",
			territory.BY: "ꗩꕞꖩꔻ",
			territory.BZ: "ꔆꔷꔘ",
			territory.CA: "ꕪꕯꕜ",
			territory.CC: "ꖏꖏꔻ (ꔞꔀꔷꘋ) ꔳꘋꗣ",
			territory.CD: "ꖏꖐ ꗵꗞꖴꕟꔎ ꕸꖃꔀ",
			territory.CF: "ꕉꔱꔸꕪ ꗳ ꗳ ꕸꖃꔀ",
			territory.CG: "ꖏꖐ",
			territory.CH: "ꖬꔃꕤ ꖨꕮꕊ",
			territory.CI: "ꖏꔳ ꕾꕎ",
			territory.CK: "ꖏꕃ ꔳꘋꗣ",
			territory.CL: "ꔚꔷ",
			territory.CM: "ꕪꔈꖩꘋ",
			territory.CN: "ꕦꔤꕯ",
			territory.CO: "ꗛꗏꔭꕩ",
			territory.CP: "ꕃꔒꕐꗋꘋ ꔳꘋꗣ",
			territory.CR: "ꖏꔻꕚ ꔸꕪ",
			territory.CU: "ꕃꖳꕑ",
			territory.CV: "ꔞꔪ ꗲꔵ ꔳꘋꗣ",
			territory.CW: "ꖴꕟꖇꕱ",
			territory.CX: "ꔞꔻꕮꔻ ꔳꘋꗣ",
			territory.CY: "ꕢꗡꖛꗐꔻ",
			territory.CZ: "ꗿꕃ ꕸꖃꔀ",
			territory.DE: "ꕧꕮꔧ",
			territory.DG: "ꔵꔀꖑ ꔳꘋꗣ",
			territory.DJ: "ꕀꖜꔳ",
			territory.DK: "ꕜꕇꕮꕃ",
			territory.DM: "ꖁꕆꕇꕪ",
			territory.DO: "ꖁꕆꕇꕪꘋ ꕸꕱꔀ",
			territory.DZ: "ꕉꔷꔠꔸꕩ",
			territory.EA: "ꗻꕚ ꗪ ꔡꔷꕞ",
			territory.EC: "ꗡꖴꔃꗍ",
			territory.EE: "ꗡꔻꕿꕇꕰ",
			territory.EG: "ꕆꔖꕞ",
			territory.EH: "ꕢꕌꕟ ꔎꔒ ꕀꔤ",
			territory.ER: "ꔀꔸꔳꕟ",
			territory.ES: "ꕐꘊꔧ",
			territory.ET: "ꔤꔳꖎꔪꕩ",
			territory.FI: "ꔱꘋ ꖨꕮꕊ",
			territory.FJ: "ꔱꔤꕀ",
			territory.FK: "ꕘꔷꕃ ꖨꕮ ꔳꘋꗣ",
			territory.FM: "ꕆꖏꕇꔻꕩ",
			territory.FO: "ꕘꖄ ꔳꘋꗣ",
			territory.FR: "ꖢꕟꘋꔻ",
			territory.GA: "ꕭꕷꘋ",
			territory.GB: "ꖕꕯꔤꗳ",
			territory.GD: "ꖶꕟꕯꕜ",
			territory.GE: "ꗘꖺꕀꕩ",
			territory.GF: "ꗱꘋꔻ ꖶꕎꕯ",
			territory.GG: "ꖶꗦꘋꔻ",
			territory.GH: "ꕭꕌꕯ",
			territory.GI: "ꕀꖜꕟꕚ",
			territory.GL: "ꕧꕓ ꖴꕎ ꖨꕮꕊ",
			territory.GM: "ꕭꔭꕩ",
			territory.GN: "ꕅꔤꕇ",
			territory.GP: "ꖶꕎꔐꖨꔅ",
			territory.GQ: "ꖦꕰꕊ ꗳ ꕅꔤꕇ",
			territory.GR: "ꗥꗷꘋ",
			territory.GS: "ꗘꖺꕀꕩ ꗛꔤ ꔒꘋꗣ ꗏ ꗪ ꗇꖢ ꔳꘋꗣ ꗛꔤ ꔒꘋꗣ ꗏ",
			territory.GT: "ꖶꕎꔎꕮꕞ",
			territory.GU: "ꖶꕎꕆ",
			territory.GW: "ꕅꔤꕇ ꔫꕢꕴ",
			territory.GY: "ꖶꕩꕯ",
			territory.HM: "ꗥꗡꔵ ꗪ ꕮꖁꕯ",
			territory.HN: "ꖽꖫꕟ",
			territory.HR: "ꖏꔓꔻꕩ",
			territory.HT: "ꕌꔤꔳ",
			territory.HU: "ꖽꘋꕭꔓ",
			territory.IC: "ꗛꖺꔻꕩ ꔳꘋꗣ",
			territory.ID: "ꔤꖆꕇꔻꕩ",
			territory.IE: "ꕉꔓ ꖨꕮꕊ",
			territory.IL: "ꕑꕇꔻꕞꔤꕞ",
			territory.IM: "ꕮꘋ ꔳꘋꗣ",
			territory.IN: "ꔤꔺꕩ",
			territory.IO: "ꔛꔟꔻ ꔤꔺꕩ ꗛꔤꘂ ꕗꕴꔀ ꕮ",
			territory.IQ: "ꔤꕟꕃ",
			territory.IR: "ꔤꕟꘋ",
			territory.IS: "ꕉꔤꔻ ꖨꕮꕊ",
			territory.IT: "ꔤꕚꔷ",
			territory.JE: "ꘀꗡꔘ",
			territory.JM: "ꕧꕮꔧꕪ",
			territory.JO: "ꗘꖺꗵꘋ",
			territory.JP: "ꔛꗨꗢ",
			territory.KE: "ꔞꕰ",
			territory.KG: "ꕃꕅꔻꕚꘋ",
			territory.KH: "ꕪꕹꔵꕩ",
			territory.KI: "ꕃꔸꕑꔳ",
			territory.KM: "ꖏꕹꖄꔻ",
			territory.KN: "ꔻꘋ ꕃꔳꔻ ꗪ ꔕꔲꔻ",
			territory.KP: "ꖏꔸꕩ ꗛꔤ ꕪꘋꗒ",
			territory.KR: "ꖏꔸꕩ ꗛꔤ ꔒꘋꗣ ꗏ",
			territory.KW: "ꖴꔃꔳ",
			territory.KY: "ꔞꔀꕮꘋ ꔳꘋꗣ",
			territory.KZ: "ꕪꕤꔻꕚꘋ",
			territory.LA: "ꕞꕴꔻ",
			territory.LB: "ꔒꕑꗟꘋ",
			territory.LC: "ꔻꘋ ꖨꔻꕩ",
			territory.LI: "ꔷꗿꘋꔻꗳꘋ",
			territory.LK: "ꖬꔸ ꕞꘋꕪ",
			territory.LR: "ꕞꔤꔫꕩ",
			territory.LS: "ꔷꖇꕿ",
			territory.LT: "ꔷꖤꔃꕇꕰ",
			territory.LU: "ꗏꔻꘋꗂꖺ",
			territory.LV: "ꕞꔳꔲꕩ",
			territory.LY: "ꔒꔫꕩ",
			territory.MA: "ꗞꕟꖏ",
			territory.MC: "ꗞꕯꖏ",
			territory.MD: "ꖒꔷꖁꕙ",
			territory.ME: "ꗞꔳꕇꖶꖄ",
			territory.MF: "ꕪꘋꕓ ꗞꗢ ꕮꕊꔳꘋ",
			territory.MG: "ꕮꕜꕭꔻꕪ",
			territory.MH: "ꕮꕊꕣ ꔳꘋꗣ",
			territory.ML: "ꕮꔷ",
			territory.MM: "ꕆꕩꘋꕮ",
			territory.MN: "ꗞꖐꔷꕩ",
			territory.MO: "ꕮꗛꖺ",
			territory.MP: "ꗛꔤ ꕪꘋꗒ ꕮꔸꕩꕯ ꔳꘋꗣ",
			territory.MQ: "ꕮꔳꕇꕃ",
			territory.MR: "ꗞꔓꔎꕇꕰ",
			territory.MS: "ꗞꘋꔖꕟꔳ",
			territory.MT: "ꕮꕊꕚ",
			territory.MU: "ꗞꔓꗔ",
			territory.MV: "ꕮꔷꕜꔍ",
			territory.MW: "ꕮꕞꕌꔨ",
			territory.MX: "ꘈꔻꖏ",
			territory.MY: "ꕮꔒꔻꕩ",
			territory.MZ: "ꕹꕤꔭꕃ",
			territory.NA: "ꕯꕆꔫꕩ",
			territory.NC: "ꕪꔷꖁꕇꕰ ꕯꕮꕊ",
			territory.NE: "ꕯꔤꕧ",
			territory.NF: "ꗟꖺꗉ ꔳꘋꗣ",
			territory.NG: "ꕯꔤꕀꔸꕩ",
			territory.NI: "ꕇꕪꕟꖶꕎ",
			territory.NL: "ꘉꕜ ꖨꕮꕊ",
			territory.NO: "ꗟꖺꔃ",
			territory.NP: "ꕇꕐꔷ",
			territory.NR: "ꖆꖩ",
			territory.NU: "ꖸꔃꔤ",
			territory.NZ: "ꔽꔤ ꖨꕮ ꕯꕮꕊ",
			territory.OM: "ꕱꕮꘋ",
			territory.PA: "ꕐꕯꕮ",
			territory.PE: "ꗨꗡꖩ",
			territory.PF: "ꗱꘋꔻ ꕶꔷꕇꔻꕩ",
			territory.PG: "ꕐꖛꕎ ꕅꔤꕇ ꕯꕮꕊ",
			territory.PH: "ꔱꔒꔪꘋ",
			territory.PK: "ꕐꕃꔻꕚꘋ",
			territory.PL: "ꕶꗷꘋ",
			territory.PM: "ꔻꘋ ꔪꘂ ꗪ ꕆꔞꗏꘋ",
			territory.PN: "ꔪꔳꕪꕆ",
			territory.PR: "ꔪꖳꕿ ꔸꖏ",
			territory.PS: "ꕐꔒꔻꔳꕯ ꔎꔒ ꕀꔤ ꗛꔤ ꕞ ꗱ ꗪ ꕭꕌꕤ",
			territory.PT: "ꕶꕿꕃꔤ ꕸꖃꔀ",
			territory.PW: "ꕐꖃ",
			territory.PY: "ꕐꕟꗝꔀ",
			territory.QA: "ꕪꕚꕌ",
			territory.RE: "ꔓꗠꖻ",
			territory.RO: "ꖄꕆꕇꕰ",
			territory.RS: "ꗻꗡꔬꕩ",
			territory.RU: "ꗐꖺꔻꕩ",
			territory.RW: "ꕟꖙꕡ",
			territory.SA: "ꕞꕌꖝ ꕸꖃꔀ",
			territory.SB: "ꖬꕞꔤꕮꕊꕯ ꔳꘋꗣ",
			territory.SC: "ꔖꗼꔷ",
			territory.SD: "ꖬꗵꘋ",
			territory.SE: "ꖬꔨꗵꘋ",
			territory.SG: "ꔻꕬꕶꕱ",
			territory.SH: "ꔻꘋ ꗥꔷꕯ",
			territory.SI: "ꔻꖃꔍꕇꕰ",
			territory.SJ: "ꔻꕙꕒꔵ ꗪ ꕧꘋ ꕮꘂꘋ",
			territory.SK: "ꔻꖃꕙꕃꕩ",
			territory.SL: "ꔋꕩ ꕒꕌꖺ ꕸꖃꔀ",
			territory.SM: "ꕮꔸꖆ ꕢꘋ",
			territory.SN: "ꔻꕇꕭꕌ",
			territory.SO: "ꖇꕮꔷꕩ",
			territory.SR: "ꖬꔸꕯꔈ",
			territory.SS: "ꖬꕜꘋ ꗛꔤ ꔒꘋꗣ ꗏ",
			territory.ST: "ꕢꕴ ꕿꔈ ꗪ ꕉ ꕮꔧ ꕗꕴꔀ",
			territory.SV: "ꗡꗷ ꕢꔍꗍꖺ",
			territory.SX: "ꔻꘋꔳ ꕮꕊꗳꘋ",
			territory.SY: "ꔻꕩꘋ",
			territory.SZ: "ꖬꕎꔽ ꖨꕮꕊ",
			territory.TA: "ꔳꔻꕚꘋ ꕜ ꖴꕯ",
			territory.TC: "ꗋꖺꕃꔻ ꗪ ꕪꔤꖏꔻ ꔳꘋꗣ",
			territory.TD: "ꕦꔵ",
			territory.TF: "ꔱꗷꘋꔻ ꗛꔤ ꔒꘋꗣ ꗏ ꕸꖃꔀ ꖸ",
			territory.TG: "ꕿꖑ",
			territory.TH: "ꕚꔤ ꖨꕮꕊ",
			territory.TJ: "ꕚꕀꕃꔻꕚꘋ",
			territory.TK: "ꕿꔞꖃ",
			territory.TL: "ꔎꔒ ꗃ ꔳꗞꖻ",
			territory.TM: "ꗋꖺꕃꕮꕇꔻꕚꘋ",
			territory.TN: "ꖤꕇꔻꕩ",
			territory.TO: "ꗋꕬ",
			territory.TR: "ꗋꖺꕃ",
			territory.TT: "ꖤꔸꔕꕜ ꗪ ꕿꔆꖑ",
			territory.TV: "ꕚꖣꖨ",
			territory.TW: "ꕚꔤꕎꘋ",
			territory.TZ: "ꕚꘋꕤꕇꕰ",
			territory.UA: "ꖳꖴꔓꘋ",
			territory.UG: "ꖳꕭꕡ",
			territory.UM: "ꕶꕱ ꕪꘋ ꗅꘋ ꔳꘋꗣ ꖸ",
			territory.US: "ꕶꕱ",
			territory.UY: "ꖳꔓꗝꔀ",
			territory.UZ: "ꖳꗩꕃꔻꕚꘋ",
			territory.VA: "ꕙꔳꕪꘋ ꕢꕨꕌ",
			territory.VC: "ꔻꘋ ꔲꘋꔻꘋ ꗪ ꖶꔓꕯꔵꘋ ꖸ",
			territory.VE: "ꕙꔳꕪꘋ ꕸꖃꔀ",
			territory.VG: "ꔛꔟꔻ ꗩꗡ ꗏ ꖷꖬ ꔳꘋꗣ",
			territory.VI: "ꕶꕱ ꗩꗡ ꗏ ꖷꖬ ꔳꘋꗣ",
			territory.VN: "ꗲꕇꖮꔃꕞ",
			territory.VU: "ꕙꖸꕎꖤ",
			territory.WF: "ꕎꔷꔻ ꗪ ꖢꖤꕯ",
			territory.WS: "ꕢꕹꖙꕉ",
			territory.XK: "ꖏꖇꕾ",
			territory.YE: "ꔝꘈꘋ",
			territory.YT: "ꕮꗚꔎ",
			territory.ZA: "ꕉꔱꔸꕪ ꗛꔤ ꔒꘋꗣ ꗏ ꕸꖃꔀ",
			territory.ZM: "ꕤꔭꕩ",
			territory.ZW: "ꔽꕓꖜꔃ",
		},
	}
}
