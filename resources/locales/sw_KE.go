// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_sw_KE() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y G", Long: "d MMMM y G", Medium: "d MMM y G", Short: "dd/MM/y GGGGG"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'saa' {0}", Long: "{1} 'saa' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Machi", Apr: "Aprili", May: "Mei", Jun: "Juni", Jul: "Julai", Aug: "Agosti", Sep: "Septemba", Oct: "Oktoba", Nov: "Novemba", Dec: "Desemba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jumapili", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumamosi"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Jumapili", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumamosi"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Jumapili", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumamosi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_sw]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "elfu 0;elfu -0", Currency: "¤\u00a0M0;¤-M0", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Diramu ya Falme za Kiarabu", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afghani ya Afghanistani", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Lek ya Albania", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Dram ya Armenia", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Gilda ya Antili ya Uholanzi", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza ya Angola", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Peso ya Ajentina", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Dola ya Australia", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "florin ya Aruba", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Manati ya Azabajani", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Maki ya Bosnia na Hezegovina Inayoweza Kubadilishwa", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Dola ya Babadosi", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Taka ya Bangladeshi", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Lev ya Bulgaria", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Dinari ya Bahareni", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Faranga ya Burundi", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Dola ya Bamuda", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Dola ya Brunei", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano ya Bolivia", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Reale ya Brazili", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dola ya Bahama", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrumi ya Bhutani", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Pula ya Botswana", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Ruble ya Belarusi", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Ruble ya Belarusi (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Dola ya Belize", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Dola ya Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faranga ya Kongo", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Faranga ya Uswisi", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "peso ya Chile", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan ya China (huru)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan ya China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso ya Kolombia", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Colon ya Kostarika", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Peso ya Kuba Inayoweza Kubadilishwa", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Peso ya Kuba", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudo ya Kepuvede", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Koruna ya Cheki", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Faranga ya Jibuti", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Kroni ya Denmaki", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso ya Dominika", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Dinari ya Aljeria", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Pauni ya Misri", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa ya Eritrea", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr ya Uhabeshi", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dola ya Fiji", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "pauni ya Visiwa vya Falkland", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Pauni ya Uingereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lari ya Jiojia", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi ya Ghana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Sidi ya Ghana", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Pauni ya Jibrata", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ya Gambia", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Faranga ya Guinea", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Faranga ya Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal ya Guatemala", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "dola ya Guyana", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Dola ya Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira ya Hondurasi", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna ya Kroeshia", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Godi ya Haiti", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Forinti ya Hungaria", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupia ya Indonesia", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Shekeli Mpya ya Israel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia ya India", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinari ya Iraki", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Riali ya Irani", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Krona ya Aisilandi", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Dola ya Jamaika", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Dinari ya Yordani", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Yeni ya Japani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilingi ya Kenya", Symbol: "Ksh"},
				currency.KGS: cldr.Currency{DisplayName: "Som ya Kyrgystan", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Rieli ya Kambodia", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Faranga ya Komoro", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Won ya Korea Kaskazini", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Won ya Korea Kusini", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinari ya Kuwait", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "dola ya Visiwa vya Cayman", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge ya Kazakistani", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Kip ya Laosi", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Pauni ya Lebanoni", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Rupia ya Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Dola ya Liberia", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ya Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litas ya Lithuania", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "Lats ya Lativia", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "Dinari ya Libya", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Diramu ya Moroko", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Leu ya Moldova", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Ariari ya Madagaska", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Dinari ya Masedonia", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Kiati ya Myama", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik ya Mongolia", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Pataka ya Macau", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya ya Mauritania (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya ya Mauritania", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Rupia ya Mauritius", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa ya Maldivi", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha ya Malawi", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Peso ya Meksiko", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringgiti ya Malesia", Symbol: "MYR"},
				currency.MZM: cldr.Currency{DisplayName: "Metikali ya Msumbiji (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Metikali ya Msumbiji", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Dola ya Namibia", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naira ya Naijeria", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Cordoba ya Nikaragwa", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Kroni ya Norwe", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Rupia ya Nepali", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Dola ya Nyuzilandi", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Riali ya Omani", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "balboa ya Panama", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "sol ya Peru", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Kina ya Papua New Guinea", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Peso ya Ufilipino", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Rupia ya Pakistani", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Zloti ya Polandi", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani ya Paragwai", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Riali ya Katari", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Leu ya Romania", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Dinari ya Serbia", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Ruble ya Urusi", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Faranga ya Rwanda", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyali ya Saudia", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Dola ya Visiwa vya Solomoni", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia ya Ushelisheli", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Pauni ya Sudani", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Pauni ya Sudani (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Krona ya Uswidi", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Dola ya Singapoo", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Pauni ya St. Helena", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Leoni ya Siera Leoni", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Shilingi ya Somalia", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "dola ya Suriname", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Pauni ya Sudani Kusini", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Pauni ya Syria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni ya Uswazi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Bahti ya Tailandi", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni ya Tajikistani", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Manati ya Turkmenistani", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinari ya Tunisia", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga ya Tonga", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Lira ya Uturuki", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Dola ya Trinidadi na Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Dola ya Taiwani", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingi ya Tanzania", Symbol: "TSh"},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia ya Ukraini", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingi ya Uganda", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Dola ya Marekani", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Peso ya Urugwai", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Som ya Uzbekistani", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Bolivar ya Venezuela (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Boliva ya Venezuela", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Dong ya Vietnamu", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu ya Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Tala ya Samoa", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Faranga ya CFA ya Afrika ya Kati", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dola ya Karibi Mashariki", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faranga ya CFA ya Afrika Magharibi", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Faranga ya CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Sarafu Isiyojulikana", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Riali ya Yemeni", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Randi ya Afrika Kusini", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha ya Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha ya Zambia", Symbol: "ZMW"},
				currency.ZWD: cldr.Currency{DisplayName: "Dola ya Zimbabwe", Symbol: ""},
			},
		},
	}
}
