// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_fo_FO() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'kl'. {0}", Long: "{1} 'kl'. {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "mai", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "mars", Apr: "apríl", May: "mai", Jun: "juni", Jul: "juli", Aug: "august", Sep: "september", Oct: "oktober", Nov: "november", Dec: "desember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sun", Mon: "mán", Tue: "týs", Wed: "mik", Thu: "hós", Fri: "frí", Sat: "ley"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "M", Thu: "H", Fri: "F", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "su", Mon: "má", Tue: "tý", Wed: "mi", Thu: "hó", Fri: "fr", Sat: "le"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sunnudagur", Mon: "mánadagur", Tue: "týsdagur", Wed: "mikudagur", Thu: "hósdagur", Fri: "fríggjadagur", Sat: "leygardagur"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_fo]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Sameindu Emirríkini dirham", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afganistan afghani", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Albania lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armenia dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Niðurlonds Karibia gyllin", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angola kwanza", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Argentina peso", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Avstralskur dollari", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba florin", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Aserbadjan manat", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Bosnia-Hersegovina mark (kann vekslast)", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados dollari", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladesj taka", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Bulgaria lev", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Barein dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundi frankur", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda dollari", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Brunei dollari", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Bolivia boliviano", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Brasilianskur real", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Bahamaoyggjar dollari", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Butan ngultrum", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Botsvana pula", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Hvítarussland ruble", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Hvítarussland ruble (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Belis dollari", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanada dollari", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo frankur", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "sveisiskur frankur", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Kili peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "kinesiskur yuan (úr landi)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "kinesiskur yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolombia peso", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Kosta Rika colón", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Kuba peso (sum kann vekslast)", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kuba peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Grønhøvdaoyggjar escudo", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Kekkia koruna", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Djibuti frankur", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "donsk króna", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Dominika peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Algeria dinar", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Egyptaland pund", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrea nakfa", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Etiopia birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Evra", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Fiji dollari", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Falklandsoyggjar pund", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "bretsk pund", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Georgia lari", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Gana cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltar pund", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambia dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Guinea frankur", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemala quetzal", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Gujana dollari", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hong Kong dollari", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Honduras lempira", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kroatia kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haiti gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Ungarn forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesia rupiah", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Ísrael new shekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "indiskir rupis", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Irak dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "iranskir rials", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "íslendsk króna", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaika dollari", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Jordan dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japanskur yen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "kenjanskur skillingur", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kirgisia som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kambodja riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komoroyggjar frankur", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Norðurkorea won", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Suðurkorea won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuvait dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Caymanoyggjar dollari", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kasakstan tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laos kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Libanon pund", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lanka rupi", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberia dollari", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Libya dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Marokko dirham", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Moldova leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Madagaskar ariary", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Makedónia denar", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Myanmar (Burma) kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongolia tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Makao pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Móritania ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Móritania ouguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Móritius rupi", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Maldivoyggjar rufiyaa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malavi kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Meksiko peso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Malaisia ringgit", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Mosambik metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibia dollari", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigeria naira", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Nikaragua córdoba", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "norsk króna", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepal rupi", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Nýsæland dollari", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Oman rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panama balboa", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Peru sol", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Papua Nýguinea kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filipsoyggjar peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistan rupi", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Pólland zloty", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Paraguai guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Katar rial", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Rumenia leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Serbia dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Russland ruble", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ruanda frankur", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudiarabia riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Salomonoyggjar dollari", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seyskelloyggjar rupi", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Sudan pund", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "svensk króna", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapor dollari", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "St. Helena pund", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leona leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somalia skillingur", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Surinam dollari", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Suðursudan pund", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Sao Tome & Prinsipi dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Sao Tome & Prinsipi dobra", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Sýria pund", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Svasiland lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Tailand baht", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Tadsjikistan somoni", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Turkmenistan manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunesia dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tonga paʻanga", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Turkaland liri", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad & Tobago dollari", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Taivan new dollari", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tansania skillingur", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukraina hryvnia", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Uganda skillingur", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "US dollari", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Uruguai peso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Usbekistan som", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Venesuela bolívar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venesuela bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Vjetnam dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoa tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Miðafrika CFA frankur", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "unse sølv", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "unse guld", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Eystur Karibia dollari", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Vesturafrika CFA frankur", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "unse palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP frankur", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "unse platin", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "ókent gjaldoyra", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Jemen rial", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Suðurafrika rand", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Sambia kwacha", Symbol: "ZMW"},
			},
		},
	}
}