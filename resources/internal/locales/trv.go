// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_trv() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "trv",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss\u202fa zzzz", Long: "h:mm:ss\u202fa z", Medium: "h:mm:ss\u202fa", Short: "h:mm\u202fa"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "JQG{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Kii", Feb: "Dhi", Mar: "Tri", Apr: "Spi", May: "Rii", Jun: "Mti", Jul: "Emi", Aug: "Mai", Sep: "Mni", Oct: "Mxi", Nov: "Mxk", Dec: "Mxd"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "K", Feb: "D", Mar: "T", Apr: "S", May: "R", Jun: "M", Jul: "E", Aug: "P", Sep: "A", Oct: "M", Nov: "K", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Kingal idas", Feb: "Dha idas", Mar: "Tru idas", Apr: "Spat idas", May: "Rima idas", Jun: "Mataru idas", Jul: "Empitu idas", Aug: "Maspat idas", Sep: "Mngari idas", Oct: "Maxal idas", Nov: "Maxal kingal idas", Dec: "Maxal dha idas"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Emp", Mon: "Kin", Tue: "Dha", Wed: "Tru", Thu: "Spa", Fri: "Rim", Sat: "Mat"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "E", Mon: "K", Tue: "D", Wed: "T", Thu: "S", Fri: "R", Sat: "M"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Jiyax sngayan", Mon: "tgKingal jiyax iyax sngayan", Tue: "tgDha jiyax iyax sngayan", Wed: "tgTru jiyax iyax sngayan", Thu: "tgSpac jiyax iyax sngayan", Fri: "tgRima jiyax iyax sngayan", Sat: "tgMataru jiyax iyax sngayan"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Brax kndaax", PM: "Baubau kndaax"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "pila Autaria", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "pila Pajey", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "pila Ipaw", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "pila Irow", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "pila Inglis", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "pila Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "pila Intia", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "pila Nihong", Symbol: "JP¥"},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.MGA: cldr.Currency{DisplayName: "", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "pila Macao", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "pila Nowey", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "pila Ruski", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "pila Taiwan", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "pila America", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "ini klayi pila ni", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.BN:      "patas Monchiara",
			language.DE:      "patas Towjih",
			language.EN:      "patas Ingrisi",
			language.ES:      "patas Espanniu",
			language.FR:      "patas Bosey",
			language.HI:      "patas Heyti",
			language.IT:      "patas Itariya",
			language.JA:      "patas Nihong",
			language.PT:      "patas Bowdu",
			language.PT_BR:   "patas Pajey",
			language.RU:      "patas Ruski",
			language.SR:      "patas Srpian",
			language.TRV:     "patas Taroko",
			language.UND:     "Ini klayna patas ni",
			language.UR:      "patas Yurtu",
			language.ZH:      "patas Ipaw",
			language.ZH_HANS: "Qantan Ipaw patas",
			language.ZH_HANT: "Baday Ipaw patas",
		},
		Territories: cldr.Territories{
			territory.AQ: "alang Nanci",
			territory.BA: "alang Posniya",
			territory.BR: "alang Pajey",
			territory.BV: "alang Puwei",
			territory.CH: "alang Switjrrant",
			territory.CN: "alang Ipaw",
			territory.DE: "alang Towjih",
			territory.FR: "alang Posey",
			territory.GB: "alang Inglis",
			territory.GS: "alang Nanjiouya ni Nansanminji",
			territory.HM: "alang Htee ni Mayktan",
			territory.HR: "alang Krowtia",
			territory.IN: "alang Intu",
			territory.IO: "alang Inglis niq Intu",
			territory.IT: "alang Itariya",
			territory.JP: "alang Nihong",
			territory.ME: "alang Mondineygrw",
			territory.RS: "alang Srbia",
			territory.RU: "alang Ruski",
			territory.SM: "alang Snmarinow",
			territory.TF: "alang Posey niq Nan",
			territory.US: "alang Amarika",
			territory.ZZ: "ini klayi na alang ni",
		},
	}
}
