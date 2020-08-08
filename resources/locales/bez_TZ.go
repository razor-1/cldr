// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_bez_TZ() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Hut", Feb: "Vil", Mar: "Dat", Apr: "Tai", May: "Han", Jun: "Sit", Jul: "Sab", Aug: "Nan", Sep: "Tis", Oct: "Kum", Nov: "Kmj", Dec: "Kmb"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "H", Feb: "V", Mar: "D", Apr: "T", May: "H", Jun: "S", Jul: "S", Aug: "N", Sep: "T", Oct: "K", Nov: "K", Dec: "K"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "pa mwedzi gwa hutala", Feb: "pa mwedzi gwa wuvili", Mar: "pa mwedzi gwa wudatu", Apr: "pa mwedzi gwa wutai", May: "pa mwedzi gwa wuhanu", Jun: "pa mwedzi gwa sita", Jul: "pa mwedzi gwa saba", Aug: "pa mwedzi gwa nane", Sep: "pa mwedzi gwa tisa", Oct: "pa mwedzi gwa kumi", Nov: "pa mwedzi gwa kumi na moja", Dec: "pa mwedzi gwa kumi na mbili"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Mul", Mon: "Vil", Tue: "Hiv", Wed: "Hid", Thu: "Hit", Fri: "Hih", Sat: "Lem"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "M", Mon: "J", Tue: "H", Wed: "H", Thu: "H", Fri: "W", Sat: "J"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "pa mulungu", Mon: "pa shahuviluha", Tue: "pa hivili", Wed: "pa hidatu", Thu: "pa hitayi", Fri: "pa hihanu", Sat: "pa shahulembela"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "pamilau", PM: "pamunyi"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "pamilau", PM: "pamunyi"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_bez]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "#,##0.00¤", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Lupila lwa Hufalme dza Huhihalabu", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Lupila lwa Huangola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Lupila lwa Huaustlalia", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Lupila lwa Hubahareni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Lupila lwa Huburundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Lupila lwa Hubotswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Lupila lwa Hukanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Lupila lwa Hukongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Lupila lwa Huuswisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Lupila lwa Huchina", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Lupila lwa Hukepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Lupila lwa Hujibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Lupila lwa Hualjelia", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Lupila lwa Humisri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Lupila lwa Hueritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Lupila lwa Huuhabeshi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Lupila lwa Yulo", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Lupila lwa Huuingereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Lupila lwa Hughana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Lupila lwa Hugambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Lupila lwa Hujine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Lupila lwa Huindia", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Lupila lwa Hijapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilingi ya Hukenya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Lupila lwa Hukomoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Lupila lwa Hulibelia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lupila lwa Hulesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Lupila lwa Hulibya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Lupila lwa Humoloko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Lupila lwa Hubukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Lupila lwa Humolitania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Lupila lwa Humolitania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Lupila lwa Humolisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Lupila lwa Humalawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Lupila lwa Humsumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Lupila lwa Hunamibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Lupila lwa Hunijelia", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Lupila lwa Hurwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Lupila lwa Husaudi", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Lupila lwa Hushelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Lupila lwa Husudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Lupila lwa Husantahelena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Lupila lwa Lioni", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Lupila lwa Husomalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Lupila lwa Husaotome na Huprinisipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Lupila lwa Husaotome na Huprinisipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lupila lwa Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Lupila lwa Hutunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingi ya Hutanzania", Symbol: "TSh"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingi ya Huuganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Lupila lwa Humalekani", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Lupila lwa CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Lupila lwa CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Lupila lwa Huafriaka ya Hukusini", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Lupila lwa Huzambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Lupila lwa Huzambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Lupila lwa Huzimbabwe", Symbol: ""},
			},
		},
	}
}
