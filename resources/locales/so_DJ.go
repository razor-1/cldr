// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_so_DJ() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM dd, y", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Abr", May: "May", Jun: "Jun", Jul: "Lul", Aug: "Ogs", Sep: "Seb", Oct: "Okt", Nov: "Nof", Dec: "Dis"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "L", Aug: "O", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Jannaayo", Feb: "Febraayo", Mar: "Maarso", Apr: "Abriil", May: "May", Jun: "Juun", Jul: "Luuliyo", Aug: "Ogost", Sep: "Sebtembar", Oct: "Oktoobar", Nov: "Nofembar", Dec: "Desembar"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Axd", Mon: "Isn", Tue: "Tldo", Wed: "Arbc", Thu: "Khms", Fri: "Jmc", Sat: "Sbti"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "I", Tue: "T", Wed: "A", Thu: "Kh", Fri: "J", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Axd", Mon: "Isn", Tue: "Tldo", Wed: "Arbc", Thu: "Khms", Fri: "Jmc", Sat: "Sbti"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Axad", Mon: "Isniin", Tue: "Talaado", Wed: "Arbaco", Thu: "Khamiis", Fri: "Jimco", Sat: "Sabti"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "GH", PM: "GD"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "h", PM: "d"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "GH", PM: "GD"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_so]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "0 Kun", Currency: "¤#,##0.00", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirhamka Isutaga Imaaraatka Carabta", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afgan Afgani", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Lekta Albaniya", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Daraamka Armeniya", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Galdarka Nadarlaan Antiliyaan", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kawansada Angola", Symbol: "Kz"},
				currency.ARA: cldr.Currency{DisplayName: "Argentine Austral", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Beesada Ley ee Arjentiin (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Beesada Ley ee Arjentiin (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Beesada Ley ee Arjentiin (1883–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Beesada Arjentiin", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Doolarka Astaraaliya", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Foloorinta Aruban", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Manaata Asarbeyjan", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "Diinaarka BBosnia-Hersogofina 1.00 konfatibal maakta Bosnia-Hersogofina 1 konfatibal maaka Bosnia-Hersogofina (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Konfatibal Maakta Bosnia-Hersogofina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Doolarka Barbaadiyaan", Symbol: "DBB"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Bangledesh", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Lefta Bulgariya", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dinaarka Baxreyn", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Faranka Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Doolarka Barmuuda", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Doolarka Buruney", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Bolifiyanada Bolifiya", Symbol: "Bs"},
				currency.BOL: cldr.Currency{DisplayName: "Bolifiyaabka Bolifiyaano(1863–1963)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Realka Barasil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Doolarka Bahamaas", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Nugultaramta Butan", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Buulada Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Rubalka Belarus", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "Doolarka Beelisa", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Doolarka Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faranka Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faranka Iswiska", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Beesada Jili", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Yuwanta Shiinaha (Ofshoor)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuwanta Shiinaha", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Beesada Kolombiya", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Kolonka Kosta Riika", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Beesada Konfatibal ee Kuuba", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Beesada Kuuba", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudo Keyb Farde", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Korunada Jeek", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faran Jabuuti", Symbol: "Fdj"},
				currency.DKK: cldr.Currency{DisplayName: "Koronka Danishka", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Beesada Dominiika", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinaarka Aljeriya", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Kroonka Estooniya", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Bowndka Masar", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfada Eritriya", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birta Itoobbiya", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuuroo", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Markkada Fiinishka ah", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Doolarka Fiji", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Bowndka Faalklaan Aylaanis", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Bowndka Biritishka", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Laariga Joorjiya", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "Sedida Gana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Bowndka Gibraltar", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasida Gambiya", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Faranka Gini", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Kuwestalka Guwatemala", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Doolarka Guyanes", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Doolarka Hoon Koon", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lembirada Honduras", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kunada Korooshiya", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Goordada Hiyati", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Forintiska Hangari", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Rubiah Indonesiya", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "baawnka Ayrishka", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Niyuu Shekelka Israaiil", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rubiga Hindiya", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinaarka Ciraaq", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Riyaalka Iran", Symbol: ""},
				currency.ISJ: cldr.Currency{DisplayName: "krónaha Iceland (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Koronada Eysland", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Doolarka Jamayka", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Dinaarka Urdun", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yenta Jabaan", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilingka Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Somta Kiyrgiystan", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Riyf kambodiya", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faranka Komoros", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Wonka Waqooyiga Kuuriya", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Wonka Koonfur Kuuriya", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinaarka Kuweyt", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Doolarka Kayman Aylaanis", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Tengeda Kasakhstan", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Kib Laoti", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Bowndka Lubnaan", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Rubiga Siri lanka", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Doolarka Liberiya", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "Rubalka Latfiya", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Dinaarka Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirhamka Moroko", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Leeyuuda Moldofa", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Arayrida Madagaskar", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Denaarka Masedoniya", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kayatda Mayanmaar", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrikta Mongoliya", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Bataka Makana", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Oogiya Mawritaniya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Oogiyada Mawritaaniya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rubiga Mowrishiya", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyada Maldifiya", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Kawajada Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Beesada Meksiko", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringitda Malayshiya", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Metikalka Mosambik", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Doolarka Namibiya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nairada Neyjeeriya", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Kordobada Nikargow", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Koronka Norway", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Rubiga Nebal", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Doolarka Niyuu Siyalaan", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Riyaalka Cumaan", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Balbow Banama", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Solsha Beeru", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kinada Babua Niyuu Gini", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Biso Filibin", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Rubiga Bakistan", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Solotida Bolaan", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Guranida Baraguway", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Riyaalka Qatar", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Liyuuda Romaniya", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dinaarka Serbiya", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Rubalka Ruushka", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Faranka Ruwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyaalka Sacuudiga", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Doolarka Solomon Aylaanis", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rubiga Siisalis", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Bowndka Suudaan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Koronka Isweden", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Doolarka Singabuur", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Bowndka St Helen", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leonka Sira Leon", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilingka Soomaaliya", Symbol: "S"},
				currency.SRD: cldr.Currency{DisplayName: "Doolarka Surinamees", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Bowndka Koonfurta Suudaan", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "Dobra Sao Tome & Birinsibi", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Bowndka Suuriya", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeenida iswaasi", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Baatka Taylaan", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Somoonida Tajikistan", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Manaata Turkmenistan", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Dinaarka Tunisiya", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Ba’angada Tonga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Liirada Turkiga", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Doolarka Tirinidad iyo Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Doolarka Taywaanta Cusub", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingka Tansaaniya", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Hirfiniyada Yukreeyn", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingka Yugandha", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Doolarka Mareeykanka", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Beesada Urugway", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Somta Usbekistan", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Bolifar Fenesuala (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Bolifarada Fenesuwela", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Dongta Fitnaam", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Fatu Fanuatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tala Samao", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Faranka CFA ee Bartamaha Afrika", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Doolarka Iist Kaaribyan", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Faranka CFA Galbeedka Afrika", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Faranka CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Lacag aan la aqoon", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Riyaalka Yemen", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Randka Koonfur Afrika", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Kawajada Sambiya", Symbol: "ZK"},
			},
		},
	}
}
