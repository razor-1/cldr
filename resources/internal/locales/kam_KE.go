// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_kam_KE() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "kam_KE",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Mbe", Feb: "Kel", Mar: "Ktũ", Apr: "Kan", May: "Ktn", Jun: "Tha", Jul: "Moo", Aug: "Nya", Sep: "Knd", Oct: "Ĩku", Nov: "Ĩkm", Dec: "Ĩkl"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "M", Feb: "K", Mar: "K", Apr: "K", May: "K", Jun: "T", Jul: "M", Aug: "N", Sep: "K", Oct: "Ĩ", Nov: "Ĩ", Dec: "Ĩ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Mwai wa mbee", Feb: "Mwai wa kelĩ", Mar: "Mwai wa katatũ", Apr: "Mwai wa kana", May: "Mwai wa katano", Jun: "Mwai wa thanthatũ", Jul: "Mwai wa muonza", Aug: "Mwai wa nyaanya", Sep: "Mwai wa kenda", Oct: "Mwai wa ĩkumi", Nov: "Mwai wa ĩkumi na ĩmwe", Dec: "Mwai wa ĩkumi na ilĩ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Wky", Mon: "Wkw", Tue: "Wkl", Wed: "Wtũ", Thu: "Wkn", Fri: "Wtn", Sat: "Wth"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Y", Mon: "W", Tue: "E", Wed: "A", Thu: "A", Fri: "A", Sat: "A"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Wa kyumwa", Mon: "Wa kwambĩlĩlya", Tue: "Wa kelĩ", Wed: "Wa katatũ", Thu: "Wa kana", Fri: "Wa katano", Sat: "Wa thanthatũ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Ĩyakwakya", PM: "Ĩyawĩoo"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Ĩyakwakya", PM: "Ĩyawĩoo"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham ya Falme za Kiarabu", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza ya Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Ndola ya Australia", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinari ya Bahareni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Faranga ya Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pula ya Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Ndola ya Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faranga ya Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faranga ya Uswisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Renminbi ya China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudo ya Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faranga ya Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinari ya Aljeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Vaundi ya Misili", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa ya Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Bir ya Uhabeshi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pauni ya Uingereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi ya Ghana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Ndalasi ya Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Faranga ya Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia ya India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Sarafu ya Kijapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Silingi ya Kenya", Symbol: "Ksh"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faranga ya Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dola ya Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ya Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinari ya Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirham ya Moroko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariary ya Bukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugwiya ya Moritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugwiya ya Moritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupia ya Morisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha ya Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikali ya Msumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Ndola ya Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira ya Nijeria", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Faranga ya Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal ya Saudia", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia ya Shelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Vaũndi ya Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Vaũndi ya Santahelena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leoni", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Silingi ya Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Ndinari ya Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Silingi ya Tanzania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Silingi ya Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Ndola ya Marekani", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Faranga CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faranga CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Randi ya Afrika Kusini", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha ya Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha ya Zambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Ndola ya Zimbabwe", Symbol: ""},
			},
		},
	}
}
