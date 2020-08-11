// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_mfe_MU() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "mfe_MU",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "zan", Feb: "fev", Mar: "mar", Apr: "avr", May: "me", Jun: "zin", Jul: "zil", Aug: "out", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "z", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "z", Jul: "z", Aug: "o", Sep: "s", Oct: "o", Nov: "n", Dec: "d"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "zanvie", Feb: "fevriye", Mar: "mars", Apr: "avril", May: "me", Jun: "zin", Jul: "zilye", Aug: "out", Sep: "septam", Oct: "oktob", Nov: "novam", Dec: "desam"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dim", Mon: "lin", Tue: "mar", Wed: "mer", Thu: "ze", Fri: "van", Sat: "sam"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "d", Mon: "l", Tue: "m", Wed: "m", Thu: "z", Fri: "v", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "dimans", Mon: "lindi", Tue: "mardi", Wed: "merkredi", Thu: "zedi", Fri: "vandredi", Sat: "samdi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "¤\u00a00K", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "dirham Emira arab ini", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angole", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "dolar ostralien", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "dinar bahreïn", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "fran burunde", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "pula ya botswane", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "dolar kanadien", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "fran kongole", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "fran swis", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "yuan renminbi sinwa", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "eskudo kapverdien", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "fran djiboutien", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "dinar alzerien", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "liv ezipsien", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "nafka erythreen", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr etiopien", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "liv sterlin", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "sedi ganeen", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi gambien", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "fran gineen", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "roupi", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "yen zapone", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "shiling kenyan", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "fran komorien", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "dolar liberien", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "loti lezoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "dinar libien", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "dirham marokin", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "fran malgas", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "ouguiya moritanien (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ouguiya moritanien", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "roupi morisien", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha malawit", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "metikal mozanbikin", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "dolar namibien", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "naira nizerian", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "fran rwande", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "rial saoudien", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "roupi seselwa", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "dinar soudane", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "liv soudane", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "liv Sainte-Hélène", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "leonn Sierra-Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "shilingi somalien", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "dobra santomeen (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "dobra santomeen", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni swazi", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "dinar tinizien", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "shiling tanzanien", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "shiling ougande", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "dolar amerikin", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "fran CFA (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "fran CFA (BCEAO)", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "rand sid-afrikin", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "kwacha zanbien (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha zanbien", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "dolar zimbawe", Symbol: ""},
			},
		},
	}
}
