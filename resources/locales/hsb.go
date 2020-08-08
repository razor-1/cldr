// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_hsb() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.yy"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm 'hodź'."}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "měr", Apr: "apr", May: "mej", Jun: "jun", Jul: "jul", Aug: "awg", Sep: "sep", Oct: "okt", Nov: "now", Dec: "dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "měrc", Apr: "apryl", May: "meja", Jun: "junij", Jul: "julij", Aug: "awgust", Sep: "september", Oct: "oktober", Nov: "nowember", Dec: "december"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "nje", Mon: "pón", Tue: "wut", Wed: "srj", Thu: "štw", Fri: "pja", Sat: "sob"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "w", Wed: "s", Thu: "š", Fri: "p", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "nj", Mon: "pó", Tue: "wu", Wed: "sr", Thu: "št", Fri: "pj", Sat: "so"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "njedźela", Mon: "póndźela", Tue: "wutora", Wed: "srjeda", Thu: "štwórtk", Fri: "pjatk", Sat: "sobota"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "dopołdnja", PM: "popołdnju"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "pop."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "dopołdnja", PM: "popołdnju"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_hsb]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "andorraska peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "ZAE dirham", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afghaniski afghani", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "albanski lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "armenski dram", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "nižozemsko-antilski gulden", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "angolska kwanza", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "angolska kwanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "angolska nowa kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "angolska kwanza reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "argentinski austral", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "argentinski peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "argentinski peso", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "awstriski šiling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "awstralski dolar", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "aruba-florin", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "azerbajdźanski manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "azerbajdźanski manat", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "bosniski dinar", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "bosniska konwertibelna hriwna", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "barbadoski dolar", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "bangladešska taka", Symbol: "৳"},
				currency.BEC: cldr.Currency{DisplayName: "belgiski frank (konwertibelny)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "belgiski frank", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "belgiski finančny frank", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "bołharski lew (1962–1999)", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "bołharski lew", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "bahrainski dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "burundiski frank", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "bermudaski dolar", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "bruneiski dolar", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "boliwiski boliviano", Symbol: "Bs"},
				currency.BOP: cldr.Currency{DisplayName: "boliwiski peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "boliwiski mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "brazilski nowy cruzeiro (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "brazilski cruzado (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "brazilski cruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "brazilski real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "brazilski nowy cruzado (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "brazilski cruzeiro (1993–1994)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "bahamaski dolar", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "bhutanski ngultrum", Symbol: ""},
				currency.BUK: cldr.Currency{DisplayName: "burmaski kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "botswanska pula", Symbol: "P"},
				currency.BYB: cldr.Currency{DisplayName: "běłoruski rubl (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "běłoruski rubl", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "běłoruski rubl (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "belizeski dolar", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "kanadiski dolar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "kongoski frank", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "šwicarski frank", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "chilski peso", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "chinski yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "kolumbiski peso", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "kosta-rikaski colón", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "kubaski konwertibelny peso", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "kubaski peso", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "kapverdski escudo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "čěska króna", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "dźibutiski frank", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "danska króna", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "dominikanski peso", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "algeriski dinar", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "egyptowski punt", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "eritrejska nakfa", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "etiopiski birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "fidźiski dolar", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "falklandski punt", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "britiski punt", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "georgiski lari", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "ghanaski cedi", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "gibraltarski punt", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "gambiski dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "ginejski frank", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "guatemalski quetzal", Symbol: "Q"},
				currency.GWP: cldr.Currency{DisplayName: "ginejsko-bissauski peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "guyanski dolar", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "hongkongski dolar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "honduraska lempira", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "chorwatska kuna", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "haitiska gourda", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "madźarski forint", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "indoneska rupija", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "israelski nowy šekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "indiska rupija", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "irakski dinar", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "iranski rial", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "islandska króna", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "jamaiski dolar", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "jordaniski dinar", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "japanski yen", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "keniaski šiling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "kirgiski som", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "kambodźaski riel", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "komorski frank", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "sewjernokorejski won", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "južnokorejski won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "kuwaitski dinar", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "kajmanski dolar", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "kazachski tenge", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "laoski kip", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "libanonski punt", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "sri-lankaska rupija", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "liberiski dolar", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "litawski litas", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "letiski lat", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "libyski dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "marokkoski dirham", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldawski leu", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "madagaskarski ariary", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "makedonski denar", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "myanmarski kyat", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "mongolski tugrik", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "macaoska pataka", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "mawretanska ouguiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "mawretanska ouguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "mauritiuska rupija", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "malediwiska rupija", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "malawiski kwacha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "mexiski peso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "malajziski ringgit", Symbol: "RM"},
				currency.MZE: cldr.Currency{DisplayName: "mosambikski escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "mosambikski metical (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "mosambikski metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "namibiski dolar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "nigeriski naira", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "nikaraguaski cordoba", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "norwegska króna", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "nepalska rupija", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "nowoseelandski dolar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "omanski rial", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "panamaski balboa", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "peruski sol", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "papua-nowoginejski kina", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "filipinski peso", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "pakistanska rupija", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "pólski złoty", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "paraguayski guarani", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "katarski rial", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "rumunski leu", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "serbiski dinar", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "ruski rubl", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ruandiski frank", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "saudi-arabski rial", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "salomonski dolar", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "seychellska rupija", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sudanski punt", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "šwedska króna", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "singapurski dolar", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "St. Helenski punt", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "sierra-leoneski leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "somaliski šiling", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "surinamski dolar", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "južnosudanski punt", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "são tomeski dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "são tomeski dobra", Symbol: "Db"},
				currency.SVC: cldr.Currency{DisplayName: "el salvadorski colón", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "syriski punt", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "swasiski lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "thaiski baht", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "tadźikski somoni", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "turkmenski manat", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "tuneziski dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "tongaski paʻanga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "turkowska lira", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "trinidad-tobagoski dolar", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "nowy taiwanski dolar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "tansaniski šiling", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "ukrainska hriwna", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "ugandaski šiling", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "ameriski dolar", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "uruguayski peso", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "uzbekski sum", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "venezuelski bolívar (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "venezuelski bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "vietnamski dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "vanuatuski vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "samoaski tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "CFA-frank (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "wuchodnokaribiski dolar", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "CFA-frank (BCEAO)", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP-frank", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "njeznata měna", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "jemenski rial", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "južnoafriski rand", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "sambiski kwacha", Symbol: "ZK"},
			},
		},
	}
}