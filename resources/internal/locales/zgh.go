// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_zgh() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "zgh",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ⵉⵏⵏ", Feb: "ⴱⵕⴰ", Mar: "ⵎⴰⵕ", Apr: "ⵉⴱⵔ", May: "ⵎⴰⵢ", Jun: "ⵢⵓⵏ", Jul: "ⵢⵓⵍ", Aug: "ⵖⵓⵛ", Sep: "ⵛⵓⵜ", Oct: "ⴽⵜⵓ", Nov: "ⵏⵓⵡ", Dec: "ⴷⵓⵊ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ⵉ", Feb: "ⴱ", Mar: "ⵎ", Apr: "ⵉ", May: "ⵎ", Jun: "ⵢ", Jul: "ⵢ", Aug: "ⵖ", Sep: "ⵛ", Oct: "ⴽ", Nov: "ⵏ", Dec: "ⴷ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ⵉⵏⵏⴰⵢⵔ", Feb: "ⴱⵕⴰⵢⵕ", Mar: "ⵎⴰⵕⵚ", Apr: "ⵉⴱⵔⵉⵔ", May: "ⵎⴰⵢⵢⵓ", Jun: "ⵢⵓⵏⵢⵓ", Jul: "ⵢⵓⵍⵢⵓⵣ", Aug: "ⵖⵓⵛⵜ", Sep: "ⵛⵓⵜⴰⵏⴱⵉⵔ", Oct: "ⴽⵜⵓⴱⵔ", Nov: "ⵏⵓⵡⴰⵏⴱⵉⵔ", Dec: "ⴷⵓⵊⴰⵏⴱⵉⵔ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ⴰⵙⴰ", Mon: "ⴰⵢⵏ", Tue: "ⴰⵙⵉ", Wed: "ⴰⴽⵕ", Thu: "ⴰⴽⵡ", Fri: "ⴰⵙⵉⵎ", Sat: "ⴰⵙⵉⴹ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ⴰⵙⴰ", Mon: "ⴰⵢⵏ", Tue: "ⴰⵙⵉ", Wed: "ⴰⴽⵕ", Thu: "ⴰⴽⵡ", Fri: "ⴰⵙⵉⵎ", Sat: "ⴰⵙⵉⴹ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ⴰⵙⴰⵎⴰⵙ", Mon: "ⴰⵢⵏⴰⵙ", Tue: "ⴰⵙⵉⵏⴰⵙ", Wed: "ⴰⴽⵕⴰⵙ", Thu: "ⴰⴽⵡⴰⵙ", Fri: "ⴰⵙⵉⵎⵡⴰⵙ", Sat: "ⴰⵙⵉⴹⵢⴰⵙ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ⵜⵉⴼⴰⵡⵜ", PM: "ⵜⴰⴷⴳⴳⵯⴰⵜ"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ⵜⵉⴼⴰⵡⵜ", PM: "ⵜⴰⴷⴳⴳⵯⴰⵜ"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ⵜⵉⴼⴰⵡⵜ", PM: "ⵜⴰⴷⴳⴳⵯⴰⵜ"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00¤", CurrencyAccounting: "#,##0.00¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "ⴰⴷⵔⵉⵎ ⵏ ⵍⵉⵎⴰⵔⴰⵜ", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "ⴽⵡⴰⵏⵣⴰ ⵏ ⴰⵏⴳⵓⵍⴰ", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵓⵙⵜⵔⴰⵍⵢⴰ", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "ⴰⴷⵉⵏⴰⵔ ⵏ ⴱⵃⵔⴰⵢⵏ", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "ⴼⵔⴰⵏⴽ ⵏ ⴱⵓⵔⵓⵏⴷⵉ", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "ⴰⴱⵓⵍⴰ ⵏ ⴱⵓⵜⵙⵡⴰⵏⴰ", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⴽⴰⵏⴰⴷⴰ", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "ⴼⵔⴰⵏⴽ ⵏ ⴽⵓⵏⴳⵓ", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "ⴰⴼⵔⴰⵏⴽ ⵏ ⵙⵡⵉⵙⵔⴰ", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "ⴰⵢⴰⵏ ⵏ ⵛⵛⵉⵏⵡⴰ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "ⵉⵙⴽⵓⴷⵓ ⵏ ⴽⴰⴱⴱⵉⵔⴷⵉ", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "ⴼⵔⴰⵏⴽ ⵏ ⴷⵊⵉⴱⵓⵜⵉ", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "ⴰⴷⵉⵏⴰⵔ ⵏ ⴷⵣⴰⵢⵔ", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "ⴰⵊⵏⵉⵀ ⵏ ⵎⵉⵚⵕ", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "ⵏⴰⴼⴽⴰ ⵏ ⵉⵔⵉⵜⵉⵔⵢⴰ", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ⴱⵉⵔ ⵏ ⵉⵜⵢⵓⴱⵢⴰ", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "ⵓⵔⵓ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "ⴰⵊⵏⵉⵀ ⴰⵙⵜⵔⵍⵉⵏⵉ ⵏ ⵏⵏⴳⵍⵉⵣ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "ⵙⵉⴷⵉ ⵏ ⵖⴰⵏⴰ", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "ⴷⴰⵍⴰⵙⵉ ⵏ ⴳⴰⵎⴱⵢⴰ", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ⴰⵔⵓⴱⵉ ⵏ ⵍⵀⵉⵏⴷ", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "ⴰⵢⴰⵏ ⵏ ⵍⵢⴰⴱⴰⵏ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ⴰⵛⵉⵍⵉⵏ ⵏ ⴽⵉⵏⵢⴰ", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "ⴼⵔⴰⵏⴽ ⵏ ⵇⵓⵎⵓⵕ", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵍⵉⴱⵉⵔⵢⴰ", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "ⵍⵓⵜⵉ ⵏ ⵍⵉⵚⵓⵟⵓ", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "ⴰⴷⵉⵏⴰⵔ ⵏ ⵍⵉⴱⵢⴰ", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "ⴰⴷⵔⵉⵎ ⵏ ⵍⵎⵖⵔⵉⴱ", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "ⴼⵔⴰⵏⴽ ⵏ ⵎⴰⴷⴰⵖⴰⵛⵇⴰⵔ", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "ⵓⵇⵉⵢⵢⴰ ⵏ ⵎⵓⵕⵉⵟⴰⵏⵢⴰ (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ⵓⵇⵉⵢⵢⴰ ⵏ ⵎⵓⵕⵉⵟⴰⵏⵢⴰ", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "ⴰⵔⵓⴱⵉ ⵏ ⵎⵓⵔⵉⵙ", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "ⴽⵡⴰⵛⴰ ⵏ ⵎⴰⵍⴰⵡⵉ", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "ⴰⵎⵉⵜⵉⴽⵍ ⵏ ⵎⵓⵣⵏⴱⵉⵇ", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵏⴰⵎⵉⴱⵢⴰ", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "ⵏⴰⵢⵔⴰ ⵏ ⵏⵉⵊⵉⵔⵢⴰ", Symbol: "₦"},
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
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ⴰⴼⵔⴰⵏⴽ ⵏ ⵔⵡⴰⵏⴷⴰ", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "ⴰⵔⵢⴰⵍ ⵏ ⵙⵙⴰⵄⵓⴷⵉⵢⴰ", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "ⴰⵔⵓⴱⵉ ⵏ ⵙⵙⵉⵛⵉⵍ", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "ⴰⴷⵉⵏⴰⵔ ⵏ ⵙⵙⵓⴷⴰⵏ", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "ⴰⵊⵏⵉⵀ ⵏ ⵙⵙⵓⴷⴰⵏ", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "ⴰⵊⵏⵉⵀ ⵏ ⵙⴰⵏⵜⵉⵍⵉⵏ", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "ⵍⵉⵢⵓⵏ", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "ⴰⵛⵉⵍⵉⵏ ⵏ ⵚⵚⵓⵎⴰⵍ", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "ⴰⴷⵓⴱⵔⴰ ⵏ ⵙⴰⵏⵟⵓⵎⵉ (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "ⴰⴷⵓⴱⵔⴰ ⵏ ⵙⴰⵏⵟⵓⵎⵉ", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "ⵍⵉⵍⴰⵏⵊⵉⵏⵉ", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "ⴰⴷⵉⵏⴰⵔ ⵏ ⵜⵓⵏⵙ", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "ⴰⵛⵉⵍⵉⵏ ⵏ ⵟⴰⵏⵥⴰⵏⵢⴰ", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "ⴰⵛⵉⵍⵉⵏ ⵏ ⵓⵖⴰⵏⴷⴰ", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵉⵡⵓⵏⴰⴽ ⵉⵎⵓⵏⵏ", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "ⴼⵔⴰⵏⴽ ⵚⵉⴼⴰ", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "ⴼⵔⴰⵏⴽ ⵚⵉⴼⴰ ⴱⵉⵙⴰⵡ", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "ⴰⵔⴰⵏⴷ ⵏ ⴰⴼⵔⵉⵇⵢⴰ ⵏ ⵉⴼⴼⵓⵙ", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "ⴰⴽⵡⴰⵛⴰ ⵏ ⵣⴰⵎⴱⵢⴰ (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "ⴰⴽⵡⴰⵛⴰ ⵏ ⵣⴰⵎⴱⵢⴰ", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵣⵉⵎⴱⴰⴱⵡⵉ (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵣⵉⵎⴱⴰⴱⵡⵉ (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵣⵉⵎⴱⴰⴱⵡⵉ (2008)", Symbol: ""},
			},
		},
		Languages: cldr.Languages{
			language.AK:  "ⵜⴰⴽⴰⵏⵜ",
			language.AM:  "ⵜⴰⵎⵀⴰⵔⵉⵜ",
			language.AR:  "ⵜⴰⵄⵔⴰⴱⵜ",
			language.BE:  "ⵜⴰⴱⵉⵍⴰⵔⵓⵙⵜ",
			language.BG:  "ⵜⴰⴱⵍⵖⴰⵔⵉⵜ",
			language.BN:  "ⵜⴰⴱⵏⵖⴰⵍⵉⵜ",
			language.CS:  "ⵜⴰⵜⵛⵉⴽⵉⵜ",
			language.DE:  "ⵜⴰⵍⵉⵎⴰⵏⵜ",
			language.EL:  "ⵜⴰⴳⵔⵉⴳⵉⵜ",
			language.EN:  "ⵜⴰⵏⴳⵍⵉⵣⵜ",
			language.ES:  "ⵜⴰⵙⴱⵏⵢⵓⵍⵉⵜ",
			language.FA:  "ⵜⴰⴼⵓⵔⵙⵉⵜ",
			language.FR:  "ⵜⴰⴼⵔⴰⵏⵙⵉⵙⵜ",
			language.HA:  "ⵜⴰⵀⴰⵡⵙⴰⵜ",
			language.HI:  "ⵜⴰⵀⵉⵏⴷⵉⵜ",
			language.HU:  "ⵜⴰⵀⵏⵖⴰⵔⵉⵜ",
			language.ID:  "ⵜⴰⵏⴷⵓⵏⵉⵙⵉⵜ",
			language.IG:  "ⵜⵉⴳⴱⵓⵜ",
			language.IT:  "ⵜⴰⵟⴰⵍⵢⴰⵏⵜ",
			language.JA:  "ⵜⴰⵊⴰⴱⴱⵓⵏⵉⵜ",
			language.JV:  "ⵜⴰⵊⴰⴱⴰⵏⵉⵜ",
			language.KM:  "ⵜⴰⵅⵎⵉⵔⵜ",
			language.KO:  "ⵜⴰⴽⵓⵔⵉⵜ",
			language.MS:  "ⵜⴰⵎⴰⵍⴰⵡⵉⵜ",
			language.MY:  "ⵜⴰⴱⵉⵔⵎⴰⵏⵉⵜ",
			language.NE:  "ⵜⴰⵏⵉⴱⴰⵍⵉⵜ",
			language.NL:  "ⵜⴰⵀⵓⵍⴰⵏⴷⵉⵜ",
			language.PA:  "ⵜⴰⴱⵏⵊⴰⴱⵉⵜ",
			language.PL:  "ⵜⴰⴱⵓⵍⵓⵏⵉⵜ",
			language.PT:  "ⵜⴰⴱⵕⵟⵇⵉⵣⵜ",
			language.RO:  "ⵜⴰⵔⵓⵎⴰⵏⵉⵜ",
			language.RU:  "ⵜⴰⵔⵓⵙⵉⵜ",
			language.RW:  "ⵜⴰⵔⵓⵡⴰⵏⴷⵉⵜ",
			language.SO:  "ⵜⴰⵙⵓⵎⴰⵍⵉⵜ",
			language.SV:  "ⵜⴰⵙⵡⵉⴷⵉⵜ",
			language.TA:  "ⵜⴰⵜⴰⵎⵉⵍⵜ",
			language.TH:  "ⵜⴰⵜⴰⵢⵍⴰⵏⴷⵉⵜ",
			language.TR:  "ⵜⴰⵜⵓⵔⴽⵉⵜ",
			language.UK:  "ⵜⵓⴽⵔⴰⵏⵉⵜ",
			language.UR:  "ⵜⵓⵔⴷⵓⵜ",
			language.VI:  "ⵜⴰⴱⵉⵜⵏⴰⵎⵉⵜ",
			language.YO:  "ⵜⴰⵢⵔⵓⴱⴰⵜ",
			language.ZGH: "ⵜⴰⵎⴰⵣⵉⵖⵜ",
			language.ZH:  "ⵜⴰⵛⵉⵏⵡⵉⵜ",
			language.ZU:  "ⵜⴰⵣⵓⵍⵓⵜ",
		},
		Territories: cldr.Territories{
			territory.AD: "ⴰⵏⴷⵓⵔⴰ",
			territory.AE: "ⵍⵉⵎⴰⵔⴰⵜ",
			territory.AF: "ⴰⴼⵖⴰⵏⵉⵙⵜⴰⵏ",
			territory.AG: "ⴰⵏⵜⵉⴳⴰ ⴷ ⴱⵔⴱⵓⴷⴰ",
			territory.AI: "ⴰⵏⴳⵉⵍⴰ",
			territory.AL: "ⴰⵍⴱⴰⵏⵢⴰ",
			territory.AM: "ⴰⵔⵎⵉⵏⵢⴰ",
			territory.AO: "ⴰⵏⴳⵓⵍⴰ",
			territory.AR: "ⴰⵔⵊⴰⵏⵜⵉⵏ",
			territory.AS: "ⵙⴰⵎⵡⴰ ⵜⴰⵎⵉⵔⵉⴽⴰⵏⵉⵜ",
			territory.AT: "ⵏⵏⵎⵙⴰ",
			territory.AU: "ⵓⵙⵜⵔⴰⵍⵢⴰ",
			territory.AW: "ⴰⵔⵓⴱⴰ",
			territory.AZ: "ⴰⴷⵔⴰⴱⵉⵊⴰⵏ",
			territory.BA: "ⴱⵓⵙⵏⴰ ⴷ ⵀⵉⵔⵙⵉⴽ",
			territory.BB: "ⴱⴰⵔⴱⴰⴷ",
			territory.BD: "ⴱⴰⵏⴳⵍⴰⴷⵉⵛ",
			territory.BE: "ⴱⵍⵊⵉⴽⴰ",
			territory.BF: "ⴱⵓⵔⴽⵉⵏⴰ ⴼⴰⵙⵓ",
			territory.BG: "ⴱⵍⵖⴰⵔⵢⴰ",
			territory.BH: "ⴱⵃⵔⴰⵢⵏ",
			territory.BI: "ⴱⵓⵔⵓⵏⴷⵉ",
			territory.BJ: "ⴱⵉⵏⵉⵏ",
			territory.BM: "ⴱⵔⵎⵓⴷⴰ",
			territory.BN: "ⴱⵔⵓⵏⵉ",
			territory.BO: "ⴱⵓⵍⵉⴱⵢⴰ",
			territory.BR: "ⴱⵔⴰⵣⵉⵍ",
			territory.BS: "ⴱⴰⵀⴰⵎⴰⵙ",
			territory.BT: "ⴱⵀⵓⵜⴰⵏ",
			territory.BW: "ⴱⵓⵜⵙⵡⴰⵏⴰ",
			territory.BY: "ⴱⵉⵍⴰⵔⵓⵙⵢⴰ",
			territory.BZ: "ⴱⵉⵍⵉⵣ",
			territory.CA: "ⴽⴰⵏⴰⴷⴰ",
			territory.CD: "ⵜⴰⴳⴷⵓⴷⴰⵏⵜ ⵜⴰⴷⵉⵎⵓⵇⵔⴰⵜⵉⵜ ⵏ ⴽⵓⵏⴳⵓ",
			territory.CF: "ⵜⴰⴳⴷⵓⴷⴰⵏⵜ ⵜⴰⵏⴰⵎⵎⴰⵙⵜ ⵏ ⵉⴼⵔⵉⵇⵢⴰ",
			territory.CG: "ⴽⵓⵏⴳⵓ",
			territory.CH: "ⵙⵡⵉⵙⵔⴰ",
			territory.CI: "ⴽⵓⵜ ⴷⵉⴼⵡⴰⵔ",
			territory.CK: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵏ ⴽⵓⴽ",
			territory.CL: "ⵛⵛⵉⵍⵉ",
			territory.CM: "ⴽⴰⵎⵉⵔⵓⵏ",
			territory.CN: "ⵛⵛⵉⵏⵡⴰ",
			territory.CO: "ⴽⵓⵍⵓⵎⴱⵢⴰ",
			territory.CR: "ⴽⵓⵙⵜⴰ ⵔⵉⴽⴰ",
			territory.CU: "ⴽⵓⴱⴰ",
			territory.CV: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵏ ⴽⴰⴱⴱⵉⵔⴷⵉ",
			territory.CY: "ⵇⵓⴱⵔⵓⵙ",
			territory.CZ: "ⵜⴰⴳⴷⵓⴷⴰⵏⵜ ⵜⴰⵜⵛⵉⴽⵉⵜ",
			territory.DE: "ⴰⵍⵎⴰⵏⵢⴰ",
			territory.DJ: "ⴷⵊⵉⴱⵓⵜⵉ",
			territory.DK: "ⴷⴰⵏⵎⴰⵔⴽ",
			territory.DM: "ⴷⵓⵎⵉⵏⵉⴽ",
			territory.DO: "ⵜⴰⴳⴷⵓⴷⴰⵏⵜ ⵜⴰⴷⵓⵎⵉⵏⵉⴽⵜ",
			territory.DZ: "ⴷⵣⴰⵢⵔ",
			territory.EC: "ⵉⴽⵡⴰⴷⵓⵔ",
			territory.EE: "ⵉⵙⵜⵓⵏⵢⴰ",
			territory.EG: "ⵎⵉⵚⵕ",
			territory.ER: "ⵉⵔⵉⵜⵉⵔⵢⴰ",
			territory.ES: "ⵙⴱⴰⵏⵢⴰ",
			territory.ET: "ⵉⵜⵢⵓⴱⵢⴰ",
			territory.FI: "ⴼⵉⵍⵍⴰⵏⴷⴰ",
			territory.FJ: "ⴼⵉⴷⵊⵉ",
			territory.FK: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵏ ⵎⴰⵍⴰⵡⵉ",
			territory.FM: "ⵎⵉⴽⵔⵓⵏⵉⵣⵢⴰ",
			territory.FR: "ⴼⵔⴰⵏⵙⴰ",
			territory.GA: "ⴳⴰⴱⵓⵏ",
			territory.GB: "ⵜⴰⴳⵍⴷⵉⵜ ⵉⵎⵓⵏⵏ",
			territory.GD: "ⵖⵔⵏⴰⵟⴰ",
			territory.GE: "ⵊⵓⵔⵊⵢⴰ",
			territory.GF: "ⴳⵡⵉⵢⴰⵏ ⵜⴰⴼⵔⴰⵏⵙⵉⵙⵜ",
			territory.GH: "ⵖⴰⵏⴰ",
			territory.GI: "ⴰⴷⵔⴰⵔ ⵏ ⵟⴰⵕⵉⵇ",
			territory.GL: "ⴳⵔⵉⵍⴰⵏⴷ",
			territory.GM: "ⴳⴰⵎⴱⵢⴰ",
			territory.GN: "ⵖⵉⵏⵢⴰ",
			territory.GP: "ⴳⵡⴰⴷⴰⵍⵓⴱ",
			territory.GQ: "ⵖⵉⵏⵢⴰ ⵏ ⵉⴽⵡⴰⴷⵓⵔ",
			territory.GR: "ⵍⵢⵓⵏⴰⵏ",
			territory.GT: "ⴳⵡⴰⵜⵉⵎⴰⵍⴰ",
			territory.GU: "ⴳⵡⴰⵎ",
			territory.GW: "ⵖⵉⵏⵢⴰ ⴱⵉⵙⴰⵡ",
			territory.GY: "ⴳⵡⵉⵢⴰⵏⴰ",
			territory.HN: "ⵀⵓⵏⴷⵓⵔⴰⵙ",
			territory.HR: "ⴽⵔⵡⴰⵜⵢⴰ",
			territory.HT: "ⵀⴰⵢⵜⵉ",
			territory.HU: "ⵀⵏⵖⴰⵔⵢⴰ",
			territory.ID: "ⴰⵏⴷⵓⵏⵉⵙⵢⴰ",
			territory.IE: "ⵉⵔⵍⴰⵏⴷⴰ",
			territory.IL: "ⵉⵙⵔⴰⵢⵉⵍ",
			territory.IN: "ⵍⵀⵉⵏⴷ",
			territory.IO: "ⵜⴰⵎⵏⴰⴹⵜ ⵜⴰⵏⴳⵍⵉⵣⵉⵜ ⵏ ⵓⴳⴰⵔⵓ ⴰⵀⵉⵏⴷⵉ",
			territory.IQ: "ⵍⵄⵉⵔⴰⵇ",
			territory.IR: "ⵉⵔⴰⵏ",
			territory.IS: "ⵉⵙⵍⴰⵏⴷ",
			territory.IT: "ⵉⵟⴰⵍⵢⴰ",
			territory.JM: "ⵊⴰⵎⴰⵢⴽⴰ",
			territory.JO: "ⵍⵓⵔⴷⵓⵏ",
			territory.JP: "ⵍⵢⴰⴱⴰⵏ",
			territory.KE: "ⴽⵉⵏⵢⴰ",
			territory.KG: "ⴽⵉⵔⵖⵉⵣⵉⵙⵜⴰⵏ",
			territory.KH: "ⴽⴰⵎⴱⵓⴷⵢⴰ",
			territory.KI: "ⴽⵉⵔⵉⴱⴰⵜⵉ",
			territory.KM: "ⵇⵓⵎⵓⵔ",
			territory.KN: "ⵙⴰⵏⴽⵔⵉⵙ ⴷ ⵏⵉⴼⵉⵙ",
			territory.KP: "ⴽⵓⵔⵢⴰ ⵏ ⵉⵥⵥⵍⵎⴹ",
			territory.KR: "ⴽⵓⵔⵢⴰ ⵏ ⵉⴼⴼⵓⵙ",
			territory.KW: "ⵍⴽⵡⵉⵜ",
			territory.KY: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵏ ⴽⴰⵢⵎⴰⵏ",
			territory.KZ: "ⴽⴰⵣⴰⵅⵙⵜⴰⵏ",
			territory.LA: "ⵍⴰⵡⵙ",
			territory.LB: "ⵍⵓⴱⵏⴰⵏ",
			territory.LC: "ⵙⴰⵏⵜⵍⵓⵙⵉ",
			territory.LI: "ⵍⵉⴽⵉⵏⵛⵜⴰⵢⵏ",
			territory.LK: "ⵙⵔⵉⵍⴰⵏⴽⴰ",
			territory.LR: "ⵍⵉⴱⵉⵔⵢⴰ",
			territory.LS: "ⵍⵉⵚⵓⵟⵓ",
			territory.LT: "ⵍⵉⵜⵡⴰⵏⵢⴰ",
			territory.LU: "ⵍⵓⴽⵙⴰⵏⴱⵓⵔⴳ",
			territory.LV: "ⵍⴰⵜⴼⵢⴰ",
			territory.LY: "ⵍⵉⴱⵢⴰ",
			territory.MA: "ⵍⵎⵖⵔⵉⴱ",
			territory.MC: "ⵎⵓⵏⴰⴽⵓ",
			territory.MD: "ⵎⵓⵍⴷⵓⴼⵢⴰ",
			territory.ME: "ⵎⵓⵏⵜⵉⵏⵉⴳⵔⵓ",
			territory.MG: "ⵎⴰⴷⴰⵖⴰⵛⵇⴰⵔ",
			territory.MH: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵏ ⵎⴰⵔⵛⴰⵍ",
			territory.ML: "ⵎⴰⵍⵉ",
			territory.MM: "ⵎⵢⴰⵏⵎⴰⵔ",
			territory.MN: "ⵎⵏⵖⵓⵍⵢⴰ",
			territory.MP: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵏ ⵎⴰⵔⵢⴰⵏ ⵏ ⵉⵥⵥⵍⵎⴹ",
			territory.MQ: "ⵎⴰⵔⵜⵉⵏⵉⴽ",
			territory.MR: "ⵎⵓⵕⵉⵟⴰⵏⵢⴰ",
			territory.MS: "ⵎⵓⵏⵙⵉⵔⴰⵜ",
			territory.MT: "ⵎⴰⵍⵟⴰ",
			territory.MU: "ⵎⵓⵔⵉⵙ",
			territory.MV: "ⵎⴰⵍⴷⵉⴼ",
			territory.MW: "ⵎⴰⵍⴰⵡⵉ",
			territory.MX: "ⵎⵉⴽⵙⵉⴽ",
			territory.MY: "ⵎⴰⵍⵉⵣⵢⴰ",
			territory.MZ: "ⵎⵓⵣⵏⴱⵉⵇ",
			territory.NA: "ⵏⴰⵎⵉⴱⵢⴰ",
			territory.NC: "ⴽⴰⵍⵉⴷⵓⵏⵢⴰ ⵜⴰⵎⴰⵢⵏⵓⵜ",
			territory.NE: "ⵏⵏⵉⵊⵉⵔ",
			territory.NF: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵏ ⵏⵓⵔⴼⵓⵍⴽ",
			territory.NG: "ⵏⵉⵊⵉⵔⵢⴰ",
			territory.NI: "ⵏⵉⴽⴰⵔⴰⴳⵡⴰ",
			territory.NL: "ⵀⵓⵍⴰⵏⴷⴰ",
			territory.NO: "ⵏⵏⵔⵡⵉⵊ",
			territory.NP: "ⵏⵉⴱⴰⵍ",
			territory.NR: "ⵏⴰⵡⵔⵓ",
			territory.NU: "ⵏⵉⵡⵉ",
			territory.NZ: "ⵏⵢⵓⵣⵉⵍⴰⵏⴷⴰ",
			territory.OM: "ⵄⵓⵎⴰⵏ",
			territory.PA: "ⴱⴰⵏⴰⵎⴰ",
			territory.PE: "ⴱⵉⵔⵓ",
			territory.PF: "ⴱⵓⵍⵉⵏⵉⵣⵢⴰ ⵜⴰⴼⵔⴰⵏⵙⵉⵙⵜ",
			territory.PG: "ⴱⴰⴱⵡⴰ ⵖⵉⵏⵢⴰ ⵜⴰⵎⴰⵢⵏⵓⵜ",
			territory.PH: "ⴼⵉⵍⵉⴱⴱⵉⵏ",
			territory.PK: "ⴱⴰⴽⵉⵙⵜⴰⵏ",
			territory.PL: "ⴱⵓⵍⵓⵏⵢⴰ",
			territory.PM: "ⵙⴰⵏⴱⵢⵉⵔ ⴷ ⵎⵉⴽⵍⵓⵏ",
			territory.PN: "ⴱⵉⵜⴽⴰⵢⵔⵏ",
			territory.PR: "ⴱⵓⵔⵜⵓ ⵔⵉⴽⵓ",
			territory.PS: "ⴰⴳⵎⵎⴰⴹ ⵏ ⵜⴰⴳⵓⵜ ⴷ ⵖⵣⵣⴰ",
			territory.PT: "ⴱⵕⵟⵇⵉⵣ",
			territory.PW: "ⴱⴰⵍⴰⵡ",
			territory.PY: "ⴱⴰⵔⴰⴳⵡⴰⵢ",
			territory.QA: "ⵇⴰⵜⴰⵔ",
			territory.RE: "ⵔⵉⵢⵓⵏⵢⵓⵏ",
			territory.RO: "ⵔⵓⵎⴰⵏⵢⴰ",
			territory.RS: "ⵙⵉⵔⴱⵢⴰ",
			territory.RU: "ⵔⵓⵙⵢⴰ",
			territory.RW: "ⵔⵡⴰⵏⴷⴰ",
			territory.SA: "ⵙⵙⴰⵄⵓⴷⵉⵢⴰ",
			territory.SB: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵏ ⵙⴰⵍⵓⵎⴰⵏ",
			territory.SC: "ⵙⵙⵉⵛⵉⵍ",
			territory.SD: "ⵙⵙⵓⴷⴰⵏ",
			territory.SE: "ⵙⵙⵡⵉⴷ",
			territory.SG: "ⵙⵏⵖⴰⴼⵓⵔⴰ",
			territory.SH: "ⵙⴰⵏⵜⵉⵍⵉⵏ",
			territory.SI: "ⵙⵍⵓⴼⵉⵏⵢⴰ",
			territory.SK: "ⵙⵍⵓⴼⴰⴽⵢⴰ",
			territory.SL: "ⵙⵙⵉⵔⴰⵍⵢⵓⵏ",
			territory.SM: "ⵙⴰⵏⵎⴰⵔⵉⵏⵓ",
			territory.SN: "ⵙⵙⵉⵏⵉⴳⴰⵍ",
			territory.SO: "ⵚⵚⵓⵎⴰⵍ",
			territory.SR: "ⵙⵓⵔⵉⵏⴰⵎ",
			territory.SS: "ⵙⵙⵓⴷⴰⵏ ⵏ ⵉⴼⴼⵓⵙ",
			territory.ST: "ⵙⴰⵡⵟⵓⵎⵉ ⴷ ⴱⵔⴰⵏⵙⵉⴱ",
			territory.SV: "ⵙⴰⵍⴼⴰⴷⵓⵔ",
			territory.SY: "ⵙⵓⵔⵢⴰ",
			territory.SZ: "ⵙⵡⴰⵣⵉⵍⴰⵏⴷⴰ",
			territory.TC: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵏ ⵜⵓⵔⴽⵢⴰ ⴷ ⴽⴰⵢⴽ",
			territory.TD: "ⵜⵛⴰⴷ",
			territory.TG: "ⵟⵓⴳⵓ",
			territory.TH: "ⵟⴰⵢⵍⴰⵏⴷ",
			territory.TJ: "ⵜⴰⴷⵊⴰⴽⵉⵙⵜⴰⵏ",
			territory.TK: "ⵟⵓⴽⵍⴰⵡ",
			territory.TL: "ⵜⵉⵎⵓⵔ ⵏ ⵍⵇⴱⵍⵜ",
			territory.TM: "ⵜⵓⵔⴽⵎⴰⵏⵙⵜⴰⵏ",
			territory.TN: "ⵜⵓⵏⵙ",
			territory.TO: "ⵟⵓⵏⴳⴰ",
			territory.TR: "ⵜⵓⵔⴽⵢⴰ",
			territory.TT: "ⵜⵔⵉⵏⵉⴷⴰⴷ ⴷ ⵟⵓⴱⴰⴳⵓ",
			territory.TV: "ⵜⵓⴼⴰⵍⵓ",
			territory.TW: "ⵟⴰⵢⵡⴰⵏ",
			territory.TZ: "ⵟⴰⵏⵥⴰⵏⵢⴰ",
			territory.UA: "ⵓⴽⵔⴰⵏⵢⴰ",
			territory.UG: "ⵓⵖⴰⵏⴷⴰ",
			territory.US: "ⵉⵡⵓⵏⴰⴽ ⵎⵓⵏⵏⵉⵏ ⵏ ⵎⵉⵔⵉⴽⴰⵏ",
			territory.UY: "ⵓⵔⵓⴳⵡⴰⵢ",
			territory.UZ: "ⵓⵣⴱⴰⴽⵉⵙⵜⴰⵏ",
			territory.VA: "ⴰⵡⴰⵏⴽ ⵏ ⴼⴰⵜⵉⴽⴰⵏ",
			territory.VC: "ⵙⴰⵏⴼⴰⵏⵙⴰⵏ ⴷ ⴳⵔⵉⵏⴰⴷⵉⵏ",
			territory.VE: "ⴼⵉⵏⵣⵡⵉⵍⴰ",
			territory.VG: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵜⵉⵎⴳⴰⴷ ⵏ ⵏⵏⴳⵍⵉⵣ",
			territory.VI: "ⵜⵉⴳⵣⵉⵔⵉⵏ ⵜⵉⵎⴳⴰⴷ ⵏ ⵉⵡⵓⵏⴰⴽ ⵎⵓⵏⵏⵉⵏ",
			territory.VN: "ⴼⵉⵜⵏⴰⵎ",
			territory.VU: "ⴼⴰⵏⵡⴰⵟⵓ",
			territory.WF: "ⵡⴰⵍⵉⵙ ⴷ ⴼⵓⵜⵓⵏⴰ",
			territory.WS: "ⵙⴰⵎⵡⴰ",
			territory.YE: "ⵢⴰⵎⴰⵏ",
			territory.YT: "ⵎⴰⵢⵓⵟ",
			territory.ZA: "ⴰⴼⵔⵉⵇⵢⴰ ⵏ ⵉⴼⴼⵓⵙ",
			territory.ZM: "ⵣⴰⵎⴱⵢⴰ",
			territory.ZW: "ⵣⵉⵎⴱⴰⴱⵡⵉ",
		},
	}
}
