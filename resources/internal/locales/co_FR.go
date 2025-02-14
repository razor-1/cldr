// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_co_FR() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "co_FR",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM 'di' 'u' y", Long: "d MMMM 'di' 'u' y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}, GMT: "UTC{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ghj.", Feb: "fer.", Mar: "mar.", Apr: "apr.", May: "mag.", Jun: "ghju.", Jul: "lug.", Aug: "aos.", Sep: "sit.", Oct: "ott.", Nov: "nuv.", Dec: "dic."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "G", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "G", Jul: "L", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ghjennaghju", Feb: "ferraghju", Mar: "marzu", Apr: "aprile", May: "maghju", Jun: "ghjugnu", Jul: "lugliu", Aug: "aostu", Sep: "sittembre", Oct: "ottobre", Nov: "nuvembre", Dec: "dicembre"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dum.", Mon: "lun.", Tue: "mar.", Wed: "mer.", Thu: "ghj.", Fri: "ven.", Sat: "sab."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "G", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "du", Mon: "lu", Tue: "ma", Wed: "me", Thu: "gh", Fri: "ve", Sat: "sa"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "dumenica", Mon: "luni", Tue: "marti", Wed: "mercuri", Thu: "ghjovi", Fri: "venneri", Sat: "sabbatu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤;(#,##0.00)\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "yuan chinese", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "EUR"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "libra sterlina", Symbol: "£GB"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "rupia indiana", Symbol: "INR"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "yen giappunese", Symbol: "JP¥"},
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
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "rublu russiu", Symbol: "₽"},
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
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "", Symbol: "$US"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "muneta scunnisciuta", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0} : {1}"},
		Languages: cldr.Languages{
			language.AR:      "arabu",
			language.AR_001:  "arabu mudernu",
			language.CO:      "corsu",
			language.CS:      "ceccu",
			language.DA:      "danese",
			language.DE:      "tedescu",
			language.DE_AT:   "tedescu austriacu",
			language.DE_CH:   "tedescu sguizzeru",
			language.EL:      "grecu",
			language.EN:      "inglese",
			language.EN_AU:   "inglese australianu",
			language.EN_CA:   "inglese canadianu",
			language.EN_US:   "inglese (S.U.)",
			language.ES:      "spagnolu",
			language.FI:      "finlandese",
			language.FR:      "francese",
			language.FR_CA:   "francese canadianu",
			language.FR_CH:   "francese sguizzeru",
			language.HU:      "ungarese",
			language.ID:      "indunesianu",
			language.IT:      "talianu",
			language.JA:      "giappunese",
			language.KO:      "cureanu",
			language.LV:      "lettone",
			language.MT:      "maltese",
			language.NL:      "neerlandese",
			language.NL_BE:   "fiammingu",
			language.PL:      "pulunese",
			language.PT:      "purtughese",
			language.PT_BR:   "purtughese brasilianu",
			language.PT_PT:   "purtughese europeanu",
			language.RU:      "russiu",
			language.SK:      "sluvaccu",
			language.SL:      "sluvenu",
			language.SV:      "svedese",
			language.TH:      "tailandese",
			language.TR:      "turcu",
			language.UND:     "lingua scunnisciuta",
			language.ZH:      "chinese mandarinu",
			language.ZH_HANS: "mandarinu simplificatu",
			language.ZH_HANT: "mandarinu tradiziunale",
		},
		Territories: cldr.Territories{
			territory.V_001: "Mondu",
			territory.V_002: "Africa",
			territory.V_009: "Oceania",
			territory.V_019: "Americhe",
			territory.V_142: "Asia",
			territory.V_150: "Europa",
			territory.V_419: "America latina",
			territory.AQ:    "Antarticu",
			territory.AT:    "Austria",
			territory.AU:    "Australia",
			territory.BE:    "Belgica",
			territory.CA:    "Canada",
			territory.CH:    "Svizzera",
			territory.CN:    "China",
			territory.CR:    "Costa Rica",
			territory.CU:    "Cuba",
			territory.CZ:    "Republica cecca",
			territory.DE:    "Alemagna",
			territory.DK:    "Danimarca",
			territory.DO:    "Republica Duminicana",
			territory.ES:    "Spagna",
			territory.EU:    "Unione europea",
			territory.FI:    "Finlandia",
			territory.FR:    "Francia",
			territory.GB:    "Reame Unitu",
			territory.GR:    "Grecia",
			territory.GT:    "Guatemala",
			territory.HU:    "Ungheria",
			territory.IE:    "Irlanda",
			territory.IL:    "Israele",
			territory.IN:    "India",
			territory.IR:    "Iran",
			territory.IS:    "Islanda",
			territory.IT:    "Italia",
			territory.JP:    "Giappone",
			territory.LB:    "Libanu",
			territory.LC:    "Santa Lucia",
			territory.MF:    "San Martinu",
			territory.MN:    "Mungulia",
			territory.MQ:    "Martinica",
			territory.MX:    "Messicu",
			territory.MY:    "Malesia",
			territory.NI:    "Nicaragua",
			territory.NL:    "Nederlanda",
			territory.NO:    "Nurvegia",
			territory.NZ:    "Nova Zelanda",
			territory.PA:    "Panama",
			territory.PE:    "Perù",
			territory.PH:    "Filippine",
			territory.PS:    "Palestina",
			territory.PT:    "Portugallu",
			territory.RS:    "Serbia",
			territory.RU:    "Russia",
			territory.SK:    "Sluvacchia",
			territory.SY:    "Siria",
			territory.TR:    "Turchia",
			territory.UN:    "Nazioni Unite",
			territory.US:    "Stati Uniti",
			territory.ZZ:    "regione scunnisciuta",
		},
	}
}
