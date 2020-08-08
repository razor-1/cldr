// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_jv() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Agt", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Maret", Apr: "April", May: "Mei", Jun: "Juni", Jul: "Juli", Aug: "Agustus", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Desember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ahad", Mon: "Sen", Tue: "Sel", Wed: "Rab", Thu: "Kam", Fri: "Jum", Sat: "Sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "S", Tue: "S", Wed: "R", Thu: "K", Fri: "J", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Ahad", Mon: "Sen", Tue: "Sel", Wed: "Rab", Thu: "Kam", Fri: "Jum", Sat: "Sab"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Ahad", Mon: "Senin", Tue: "Selasa", Wed: "Rabu", Thu: "Kamis", Fri: "Jumat", Sat: "Sabtu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Isuk", PM: "Wengi"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "Isuk", PM: "Wengi"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Isuk", PM: "Wengi"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_jv]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "0 èwu", Currency: "¤\u00a0#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham Uni Emirat Arab", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghani Afganistan", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Lek Albania", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Dram Armenia", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Guilder Antilla Walanda", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Peso Argentina", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dolar Australia", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Florin Aruban", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Manat Azerbaijan", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Mark Konvertibel Bosnia-Herzegovina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Dolar Barbadian", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Bangladesh", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Lev Bulgaria", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bahrain Dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Franc Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Dolar Bermuda", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Dolar Brunai", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano Bolivia", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Real Brasil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dolar Bahamian", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum Bhutan", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Pula Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Ruble Belarusia", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "Dolar Belise", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dolar Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Franc Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Franc Swiss", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Peso Chili", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan Tyongkok (Jaban Rangkah)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Tyongkok", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso Kolumbia", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Colon Kosta Rika", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Peso Konvertibel Kuba", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Peso Kuba", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Tanjung Verde", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Koruna Czech", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Franc Djibouti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Krone Denmark", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Dominika", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar Algeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Pound Mesir", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr Ethiopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dolar Fiji", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Pound Kepuloan Falkland", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pound Inggris", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lari Georgia", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "Cedi Ghana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Pound Gibraltar", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Franc Guinea", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal Guatemala", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Dolar Guyana", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Dolar Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira Honduras", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna Kroasia", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde Haiti", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Forint Hungaria", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Rupiah Indonesia", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Shekel Anyar Israel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupee India", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar Irak", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Rial Iran", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Krona Islandia", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Dolar Jamaika", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar Yordania", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yen Jepang", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilling Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Som Kirgistan", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Riel Kamboja", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Franc Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Won Korea Lor", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Won Korea Kidul", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar Kuwait", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Dolar Kepuloan Caiman", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge Kasakhstan", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Kip Laos", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Pound Libanon", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Rupee Sri Lanka", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dolar Liberia", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinar Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirham Maroko", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Leu Moldova", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariary Malagasi", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Denar Masedonia", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kyat Myanmar", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik Mongol", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca Macau", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya Mauritania (1973 - 2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya Mauritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupee Mauritius", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa Maladewa", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Peso Meksiko", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit Malaysia", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Metical Mosambik", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dolar Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira Nigeria", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Cordoba Nikaragua", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Krone Norwegia", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Rupee Nepal", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Dolar Selandia Anyar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial Oman", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Balboa Panama", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sol Peru", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina Papua Nugini", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Piso Filipina", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Rupee Pakistan", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty Polandia", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani Paraguay", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Rial Qatar", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Leu Rumania", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar Serbia", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Rubel Rusia", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Franc Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal Saudi", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Dolar Kepuloan Solomon", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupee Seichelles", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Pound Sudan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Krona Swedia", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Dolar Singapura", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Pound Santa Helena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leone Sierra Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilling Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Dolar Suriname", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Pound Sudan Kidul", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "Dobra Sao Tome lan Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Pound Siria", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni Swasi", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Baht Thai", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni Tajikistan", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Manat Turmenistan", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Dinar Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga Tonga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Lira Turki", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Dolar Trinidad lan Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Dolar Anyar Taiwan", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilling Tansania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia Ukrania", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilling Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dolar Amerika Serikat", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Peso Uruguay", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Som Usbekistan", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Bolivar Venezuela (2008 - 2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Bolivar Venezuela", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Dong Vietnam", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu Vanuatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tala Samoa", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "CFA Franc Afrika Tengah", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dolar Karibia Wetan", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "CFA Franc Afrika Kulon", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Franc CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Dhuwit Ora Dikenali", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Rial Yaman", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Rand Afrika Kidul", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha Sambia", Symbol: "ZK"},
			},
		},
	}
}