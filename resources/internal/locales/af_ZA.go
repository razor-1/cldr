// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_af_ZA() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "af_ZA",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE dd MMMM y", Long: "dd MMMM y", Medium: "dd MMM y", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan.", Feb: "Feb.", Mar: "Mrt.", Apr: "Apr.", May: "Mei", Jun: "Jun.", Jul: "Jul.", Aug: "Aug.", Sep: "Sep.", Oct: "Okt.", Nov: "Nov.", Dec: "Des."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januarie", Feb: "Februarie", Mar: "Maart", Apr: "April", May: "Mei", Jun: "Junie", Jul: "Julie", Aug: "Augustus", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Desember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "So.", Mon: "Ma.", Tue: "Di.", Wed: "Wo.", Thu: "Do.", Fri: "Vr.", Sat: "Sa."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "W", Thu: "D", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "So.", Mon: "Ma.", Tue: "Di.", Wed: "Wo.", Thu: "Do.", Fri: "Vr.", Sat: "Sa."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sondag", Mon: "Maandag", Tue: "Dinsdag", Wed: "Woensdag", Thu: "Donderdag", Fri: "Vrydag", Sat: "Saterdag"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "vm.", PM: "nm."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "v", PM: "n"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "vm.", PM: "nm."}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Verenigde Arabiese Emirate-dirham", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afgaanse afgani", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Albanese lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armeense dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Nederlands-Antilliaanse gulde", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angolese kwanza", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Argentynse peso", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Australiese dollar", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Arubaanse floryn", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Azerbeidjaanse manat", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Bosnies-Herzegowiniese omskakelbare marka", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados-dollar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladesjiese taka", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Bulgaarse lev", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Bahreinse dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundiese frank", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda-dollar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Broeneise dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviaanse boliviano", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Brasilliaanse reaal", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Bahamiaanse dollar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Bhoetanese ngoeltroem", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Botswana-pula", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Belarusiese roebel", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Belo-Russiese roebel (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Beliziese dollar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanadese dollar", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "Kongolese frank", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Switserse frank", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Chileense peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Chinese joean (buiteland)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Chinese joean", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Colombiaanse peso", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Costa Ricaanse colón", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Kubaanse omskakelbare peso", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kubaanse peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Kaap Verdiese escudo", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Tsjeggiese kroon", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Djiboeti-frank", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Deense kroon", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dominikaanse peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Algeriese dinar", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Egiptiese pond", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrese nakfa", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Etiopiese birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Fidjiaanse dollar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Falkland-eilandse pond", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Britse pond", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Georgiese lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ghanese cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ghanese cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltarese pond", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambiese dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Guinese frank", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Guinese syli", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemalaanse quetzal", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Guyanese dollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkongse dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Hondurese lempira", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kroatiese kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haïtiaanse gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Hongaarse florint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesiese roepia", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Israeliese nuwe sikkel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indiese roepee", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Irakse dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Iranse rial", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Yslandse kroon", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Italiaanse lier", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Jamaikaanse dollar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Jordaniese dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Japannese jen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Keniaanse sjieling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kirgisiese som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kambodjaanse riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Comoraanse frank", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Noord-Koreaanse won", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Suid-Koreaanse won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Koeweitse dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Cayman-eilandse dollar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kazakse tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laosiaanse kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libanese pond", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lankaanse roepee", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberiese dollar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesotho loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litause litas", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "Lettiese lats", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "Libiese dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Marokkaanse dirham", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Moldowiese leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Malgassiese ariary", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Macedoniese denar", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Mianmese kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongoolse toegrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Macaose pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mauritaniese ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Mauritaniese ouguiya", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Mauritiaanse roepee", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Malediviese rufia", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malawiese kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Meksikaanse peso", Symbol: "MXN"},
				currency.MYR: cldr.Currency{DisplayName: "Maleisiese ringgit", Symbol: "MYR"},
				currency.MZM: cldr.Currency{DisplayName: "Mosambiekse metical (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mosambiekse metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibiese dollar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigeriese naira", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Nicaraguaanse córdoba", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Noorse kroon", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepalese roepee", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Nieu-Seelandse dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Omaanse rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panamese balboa", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Peruaanse sol", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Papoea-Nieu-Guinese kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filippynse peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistanse roepee", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Poolse zloty", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Paraguaanse guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Katarrese rial", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Roemeense leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Serwiese dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Russiese roebel", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Rwandese frank", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saoedi-Arabiese riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Salomonseilandse dollar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seychellese roepee", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Soedannese pond", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Soedannese pond (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Sweedse kroon", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapoer-dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Sint Helena-pond", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leoniese leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somaliese sjieling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Surinaamse dollar", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Suid-Soedanese pond", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "São Tomé en Príncipe dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "São Tomé en Príncipe-dobra", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Siriese pond", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Swazilandse lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Thaise baht", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Tadjikse somoni", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Turkmeense manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunisiese dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tongaanse pa’anga", Symbol: "TOP"},
				currency.TRL: cldr.Currency{DisplayName: "Turkse lier (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Turkse lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad en Tobago-dollar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Nuwe Taiwanese dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzaniese sjieling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Oekraïnse hriwna", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Ugandese sjieling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "VSA-dollar", Symbol: "USD"},
				currency.UYU: cldr.Currency{DisplayName: "Uruguaanse peso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Oezbekiese som", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Venezolaanse bolivar", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venezolaanse bolívar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Viëtnamese dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatuse vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoaanse tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Sentraal Afrikaanse CFA-frank", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Oos-Karibiese dollar", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Wes-Afrikaanse CFA-frank", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP-frank", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "onbekende geldeenheid", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Jemenitiese rial", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Suid-Afrikaanse rand", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambiese kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Zambiese kwacha", Symbol: "ZMW"},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwiese dollar", Symbol: ""},
			},
		},
	}
}
