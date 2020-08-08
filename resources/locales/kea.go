// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_kea() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d 'di' MMMM 'di' y", Long: "d 'di' MMMM 'di' y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Set", Oct: "Otu", Nov: "Nuv", Dec: "Diz"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Janeru", Feb: "Febreru", Mar: "Marsu", Apr: "Abril", May: "Maiu", Jun: "Junhu", Jul: "Julhu", Aug: "Agostu", Sep: "Setenbru", Oct: "Otubru", Nov: "Nuvenbru", Dec: "Dizenbru"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dum", Mon: "sig", Tue: "ter", Wed: "kua", Thu: "kin", Fri: "ses", Sat: "sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "S", Tue: "T", Wed: "K", Thu: "K", Fri: "S", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "du", Mon: "si", Tue: "te", Wed: "ku", Thu: "ki", Fri: "se", Sat: "sa"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "dumingu", Mon: "sigunda-fera", Tue: "tersa-fera", Wed: "kuarta-fera", Thu: "kinta-fera", Fri: "sesta-fera", Sat: "sábadu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "am", PM: "pm"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_kea]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Diren di Emiradus Arabi Unidu", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kuanza", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dola australianu", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Dinar di Barain", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Franku borundes", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Rial brazileru", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Pula di Botsuana", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dola kanadianu", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Franku kongoles", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Franku suisu", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Iuan xines", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Skudu Kabuverdianu", Symbol: "\u200b"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Franku di Djibuti", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Kuroa dinamarkeza", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar arjelinu", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Libra ejípsiu", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nafka di Eritreia", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Bir etiópiku", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Libra britániku", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi di Gana (1979–2007)", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Sili", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Dola di Ong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupia indoneziu", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia indianu", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Ieni japones", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Xelin kenianu", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Franku di Komoris", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Won sul-koreanu", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Dola liberianu", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti di Lezotu", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinar líbiu", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Diren marokinu", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Ariari di Madagaskar", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ougia (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ougia", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupia di Maurisias", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kuaxa di Malaui", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Pezu mexikanu", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "MYR"},
				currency.MZM: cldr.Currency{DisplayName: "Metikal", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Dola namibianu", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Kuroa norueges", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty polaku", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rublu rusu", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Franku ruandes", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Rial saudita", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia di Seixelis", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Libra sudanes", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Libra sudanes antigu (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Kuroa sueku", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Libra di Santa Ilena", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Leone di Sera Leoa", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Xelin somalianu", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra di Sãu Tume i Prínsipi (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Dobra di San Tume i Prínsipi", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilanjeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Baht tailandes", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinar tunizianu", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Lira turku", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Dola Novu di Taiwan", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Xelin di Tanzania", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Xelin ugandensi", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Dola merkanu", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "VEF"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Franku CFA (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Franku CFA (BCEAO)", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Mueda diskonxedu", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Rand sulafrikanu", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Kuaxa zambianu (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kuaxa zambianu", Symbol: "ZMW"},
				currency.ZWD: cldr.Currency{DisplayName: "Dola di Zimbabue (1980–2008)", Symbol: ""},
			},
		},
	}
}