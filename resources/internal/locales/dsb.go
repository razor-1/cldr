// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_dsb() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "dsb",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.yy"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "měr", Apr: "apr", May: "maj", Jun: "jun", Jul: "jul", Aug: "awg", Sep: "sep", Oct: "okt", Nov: "now", Dec: "dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "měrc", Apr: "apryl", May: "maj", Jun: "junij", Jul: "julij", Aug: "awgust", Sep: "september", Oct: "oktober", Nov: "nowember", Dec: "december"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "nje", Mon: "pón", Tue: "wał", Wed: "srj", Thu: "stw", Fri: "pět", Sat: "sob"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "w", Wed: "s", Thu: "s", Fri: "p", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "nj", Mon: "pó", Tue: "wa", Wed: "sr", Thu: "st", Fri: "pě", Sat: "so"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "njeźela", Mon: "pónjeźele", Tue: "wałtora", Wed: "srjoda", Thu: "stwórtk", Fri: "pětk", Sat: "sobota"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "dopołdnja", PM: "wótpołdnja"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "wótp."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "dopołdnja", PM: "wótpołdnja"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "andorraska peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "ZAE dirham", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afghaniski afgani", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "albański lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "armeński dram", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "nižozemsko-antilski gulden", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "angolska kwanza", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "angolska kwanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "angolska nowa kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "angolska kwanza reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "argentinski austral", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "argentinski peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "argentinski peso", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "rakuski šiling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "awstralski dolar", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "aruba-florin", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "azerbajdžaniski manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "azerbajdžaniski manat", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "bosniski dinar", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "bosniska konwertibelna marka", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "barbadoski dolar", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "bangladešska taka", Symbol: "৳"},
				currency.BEC: cldr.Currency{DisplayName: "belgiski frank (konwertibelny)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "belgiski frank", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "belgiski financny frank", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "bulgarski lew (1962–1999)", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "bulgarski lew", Symbol: "BGN"},
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
				currency.BTN: cldr.Currency{DisplayName: "bhutański ngultrum", Symbol: ""},
				currency.BUK: cldr.Currency{DisplayName: "burmaski kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "botswaniska pula", Symbol: "P"},
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
				currency.CRC: cldr.Currency{DisplayName: "kosta-rikański colón", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "kubański konwertibelny peso", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "kubański peso", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "kapverdski escudo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "česka krona", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "dźibutiski frank", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "dańska krona", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "dominikański peso", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "algeriski dinar", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "egyptojski punt", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "eritrejska nakfa", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "etiopiski birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "fidźiski dolar", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "falklandski punt", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "britiski punt", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "georgiski lari", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "ghanaski cedi", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "gibraltiski punt", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "gambiski dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "guineski frank", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "guatemalski quetzal", Symbol: "Q"},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissau peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "guyański dolar", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "hongkongski dolar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "honduraska lempira", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "chorwatska kuna", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "haitiska gourda", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "madźarski forint", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "indoneska rupija", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "israelski nowy šekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "indiska rupija", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "irakski dinar", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "irański rial", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "islandska krona", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "jamaiski dolar", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "jordaniski dinar", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "japański yen", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "keniaski šiling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "kirgiski som", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "kambodžaski riel", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "komorski frank", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "pódpołnocnokorejski won", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "pódpołdnjowokorejski won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "kuwaitski dinar", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "kajmaniski dolar", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "kazachski tenge", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "laoski kip", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "libanoński punt", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "sri-lankaska rupija", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "liberiski dolar", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "litawski litas", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "letiski lat", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "libyski dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "marokkoski dirham", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldawiski leu", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "madagaskarski ariary", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "makedoński denar", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "myanmarski kyat", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "mongolski tugrik", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "macaoska pataca", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "mauretański ouguiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "mauretański ouguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "mauriciska rupija", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "malediwiska rupija", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "malawiski kwacha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "mexiski peso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "malajziski ringgit", Symbol: "RM"},
				currency.MZE: cldr.Currency{DisplayName: "Mozabicke escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "mosambikski metical (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "mosambikski metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "namibiski dolar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "nigeriska naira", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "nikaraguaska cordoba", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "norwegska krona", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "nepalska rupija", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "nowoseelandski dolar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "omański rial", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "panamaski balboa", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "peruski sol", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "papua-neuguinejska kina", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "filipinski peso", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "pakistańska rupija", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "pólski złoty", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "paraguayski guarani", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "katarski rial", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "rumuński leu", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "serbiski dinar", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "ruski rubl", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ruandiski frank", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "saudi-arabiski rial", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "salomoński dolar", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "seychelska rupija", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sudański punt", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "šwedska krona", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "singapurski dolar", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "St. Helena punt", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "sierra-leoneski leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "somaliski šiling", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "surinamski dolar", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "pódpołdnjowosudański punt", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "são-tomeska dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "são-tomeska dobra", Symbol: "Db"},
				currency.SVC: cldr.Currency{DisplayName: "el-salvadorski colón", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "syriski punt", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "swasiski lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "thaiski baht", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "tadźikiski somoni", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "turkmeniski manat", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "tuneziski dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "tongaski paʻanga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "turkojska lira", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "trinidad-tobagoski dolar", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "nowy taiwański dolar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "tansaniski šiling", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "ukrainska griwna", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "ugandaski šiling", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "ameriski dolar", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "uruguayski peso", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "usbekiski sum", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "venezuelski bolívar (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "venezuelski bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "vietnamski dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "vanuatski vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "samoaska tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "CFA-frank (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "pódzajtšnokaribiski dolar", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "CFA-frank (BCEAO)", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP-frank", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "njeznate pjenjeze", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "jemeński rial", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "pódpołdnjowoafriski rand", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "sambiska kwacha", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "afaršćina",
			language.AB:      "abchazšćina",
			language.AF:      "afrikans",
			language.AGQ:     "aghem",
			language.AK:      "akanšćina",
			language.AM:      "amharšćina",
			language.AN:      "aragonšćina",
			language.ANG:     "anglosaksojšćina",
			language.AR:      "arabšćina",
			language.AR_001:  "moderna wusokoarabšćina",
			language.ARN:     "arawkašćina",
			language.AS:      "asamšćina",
			language.ASA:     "pare",
			language.AST:     "asturšćina",
			language.AV:      "awaršćina",
			language.AY:      "aymaršćina",
			language.AZ:      "azerbajdžanšćina",
			language.BA:      "baškiršćina",
			language.BE:      "běłorušćina",
			language.BEM:     "bemba",
			language.BEZ:     "bena",
			language.BG:      "bulgaršćina",
			language.BI:      "bislamšćina",
			language.BM:      "bambara",
			language.BN:      "bengalšćina",
			language.BO:      "tibetšćina",
			language.BR:      "bretonšćina",
			language.BRX:     "bodo",
			language.BS:      "bosnišćina",
			language.BUG:     "bugišćina",
			language.CA:      "katanlanšćina",
			language.CGG:     "chiga",
			language.CH:      "čamoršćina",
			language.CHO:     "choctawšćina",
			language.CHR:     "cherokee",
			language.CKB:     "sorani",
			language.CO:      "korsišćina",
			language.CR:      "kri",
			language.CS:      "češćina",
			language.CY:      "walizišćina",
			language.DA:      "danšćina",
			language.DAV:     "taita",
			language.DE:      "nimšćina",
			language.DE_AT:   "awstriska nimšćina",
			language.DE_CH:   "šwicarska wusokonimšćina",
			language.DJE:     "zarma",
			language.DSB:     "dolnoserbšćina",
			language.DUA:     "duala",
			language.DV:      "divehi",
			language.DYO:     "jola-fonyi",
			language.DZ:      "dzongkha",
			language.EBU:     "embu",
			language.EE:      "ewe",
			language.EL:      "grichišćina",
			language.EN:      "engelšćina",
			language.EN_AU:   "awstralska engelšćina",
			language.EN_CA:   "kanadiska engelšćina",
			language.EN_GB:   "UK-engelšćina",
			language.EN_US:   "US-engelšćina",
			language.EO:      "esperanto",
			language.ES:      "špańšćina",
			language.ES_419:  "łatyńskoamerikańska špańšćina",
			language.ES_ES:   "europejska špańšćina",
			language.ES_MX:   "mexikańska špańšćina",
			language.ET:      "estišćina",
			language.EU:      "baskišćina",
			language.FA:      "persišćina",
			language.FI:      "finšćina",
			language.FIL:     "filipinšćina",
			language.FJ:      "fidžišćina",
			language.FO:      "ferejšćina",
			language.FR:      "francojšćina",
			language.FR_CA:   "kanadiska francojšćina",
			language.FR_CH:   "šwicarska francojšćina",
			language.FY:      "frizišćina",
			language.GA:      "iršćina",
			language.GAG:     "gagauzšćina",
			language.GD:      "šotišćina",
			language.GL:      "galicišćina",
			language.GN:      "guarani",
			language.GOT:     "gotišćina",
			language.GSW:     "šwicarska nimšćina",
			language.GU:      "gudžaratšćina",
			language.GUZ:     "gusii",
			language.GV:      "manšćina",
			language.HA:      "hausa",
			language.HAW:     "hawaiišćina",
			language.HE:      "hebrejšćina",
			language.HI:      "hindišćina",
			language.HR:      "chorwatšćina",
			language.HSB:     "górnoserbšćina",
			language.HT:      "haitišćina",
			language.HU:      "hungoršćina",
			language.HY:      "armeńšćina",
			language.IA:      "interlingua",
			language.ID:      "indonešćina",
			language.IG:      "igbo",
			language.II:      "sichuan yi",
			language.IK:      "inupiak",
			language.IO:      "ido",
			language.IS:      "islandšćina",
			language.IT:      "italšćina",
			language.IU:      "inuitšćina",
			language.JA:      "japańšćina",
			language.JGO:     "ngomba",
			language.JMC:     "machame",
			language.JV:      "javašćina",
			language.KA:      "georgišćina",
			language.KAB:     "kabylšćina",
			language.KAM:     "kamba",
			language.KDE:     "makonde",
			language.KEA:     "kapverdšćina",
			language.KHQ:     "koyra chiini",
			language.KI:      "kikuyu",
			language.KK:      "kazachšćina",
			language.KL:      "grönlandšćina",
			language.KLN:     "kalenjin",
			language.KM:      "kambodžanšćina",
			language.KN:      "kannadšćina",
			language.KO:      "korejańšćina",
			language.KOI:     "komi-permyak",
			language.KOK:     "konkani",
			language.KRI:     "krio",
			language.KS:      "kašmiršćina",
			language.KSB:     "šambala",
			language.KSF:     "bafia",
			language.KU:      "kurdišćina",
			language.KW:      "kornišćina",
			language.KY:      "kirgišćina",
			language.LA:      "łatyńšćina",
			language.LAG:     "langi",
			language.LB:      "luxemburgšćina",
			language.LG:      "gandšćina",
			language.LI:      "limburšćina",
			language.LKT:     "lakotšćina",
			language.LN:      "lingala",
			language.LO:      "laošćina",
			language.LT:      "litawšćina",
			language.LU:      "luba-katanga",
			language.LUO:     "luo",
			language.LUY:     "luhya",
			language.LV:      "letišćina",
			language.MAS:     "masaišćina",
			language.MER:     "meru",
			language.MFE:     "mauriciska kreolšćina",
			language.MG:      "malgašćina",
			language.MGH:     "makhuwa-meetto",
			language.MGO:     "meta’",
			language.MI:      "maorišćina",
			language.MK:      "makedońšćina",
			language.ML:      "malajamšćina",
			language.MN:      "mongolšćina",
			language.MOH:     "mohawkšćina",
			language.MR:      "maratišćina",
			language.MS:      "malajšćina",
			language.MT:      "maltašćina",
			language.MUA:     "mundang",
			language.MUS:     "krik",
			language.MY:      "burmašćina",
			language.NA:      "naurušćina",
			language.NAQ:     "nama",
			language.NB:      "norwegske bokmål",
			language.ND:      "pódpołnocne ndebele",
			language.NDS:     "dolnonimšćina",
			language.NE:      "nepalšćina",
			language.NL:      "nižozemšćina",
			language.NL_BE:   "flamšćina",
			language.NMG:     "kwasio",
			language.NN:      "norwegske nynorsk",
			language.NO:      "norwegšćina",
			language.NQO:     "n’ko",
			language.NUS:     "nuer",
			language.NV:      "navaho",
			language.NYN:     "nyankole",
			language.OC:      "okcitanšćina",
			language.OM:      "oromo",
			language.OR:      "orojišćina",
			language.PA:      "pandžabšćina",
			language.PL:      "pólšćina",
			language.PRG:     "prusčina",
			language.PS:      "paštunšćina",
			language.PT:      "portugalšćina",
			language.PT_BR:   "brazilska portugalšćina",
			language.PT_PT:   "europejska portugalšćina",
			language.QU:      "kečua",
			language.QUC:     "kʼicheʼ",
			language.RM:      "retoromańšćina",
			language.RN:      "kirundišćina",
			language.RO:      "rumunšćina",
			language.RO_MD:   "moldawišćina",
			language.ROF:     "rombo",
			language.RU:      "rušćina",
			language.RW:      "kinjarwanda",
			language.RWK:     "rwa",
			language.SA:      "sanskrit",
			language.SAQ:     "samburu",
			language.SBP:     "sangu",
			language.SC:      "sardinšćina",
			language.SCN:     "sicilianišćina",
			language.SD:      "sindšćina",
			language.SE:      "lapšćina",
			language.SEH:     "sena",
			language.SES:     "koyra senni",
			language.SG:      "sango",
			language.SH:      "serbochorwatšćina",
			language.SHI:     "tašelhit",
			language.SI:      "singalšćina",
			language.SK:      "słowakšćina",
			language.SL:      "słowjeńšćina",
			language.SM:      "samošćina",
			language.SMA:     "pódpołdnjowa samišćina",
			language.SMJ:     "lule-samišćina",
			language.SMN:     "inari-samišćina",
			language.SMS:     "skolt-samišćina",
			language.SN:      "šonšćina",
			language.SO:      "somališćina",
			language.SQ:      "albanšćina",
			language.SR:      "serbišćina",
			language.SS:      "siswati",
			language.ST:      "pódpołdnjowa sotšćina (Sesotho)",
			language.STQ:     "saterfrizišćina",
			language.SU:      "sundanšćina",
			language.SV:      "šwedšćina",
			language.SW:      "swahilišćina",
			language.SW_CD:   "kongojska swahilišćina",
			language.TA:      "tamilšćina",
			language.TE:      "telugšćina",
			language.TEO:     "teso",
			language.TG:      "tadžikišćina",
			language.TH:      "thailandšćina",
			language.TI:      "tigrinja",
			language.TK:      "turkmeńšćina",
			language.TL:      "tagalog",
			language.TN:      "tswana",
			language.TO:      "tonganšćina",
			language.TR:      "turkojšćina",
			language.TS:      "tsonga",
			language.TT:      "tataršćina",
			language.TW:      "twi",
			language.TWQ:     "tasawaq",
			language.TY:      "tahitišćina",
			language.TZM:     "centralnoatlaski tamazight",
			language.UG:      "ujguršćina",
			language.UK:      "ukrainšćina",
			language.UND:     "njeznata rěc",
			language.UR:      "urdušćina",
			language.UZ:      "usbekšćina",
			language.VAI:     "vai",
			language.VI:      "vietnamšćina",
			language.VO:      "volapük",
			language.VUN:     "vunjo",
			language.WA:      "walonšćina",
			language.WO:      "wolof",
			language.XH:      "xhosa",
			language.XOG:     "soga",
			language.YI:      "jidišćina",
			language.YO:      "jorubšćina",
			language.ZA:      "zhuang",
			language.ZGH:     "standardny marokkański tamazight",
			language.ZH:      "chinšćina",
			language.ZH_HANS: "chinšćina (zjadnorjona)",
			language.ZH_HANT: "chinšćina (tradicionalna)",
			language.ZU:      "zulu",
			language.ZXX:     "žedno rěcne wopśimjeśe",
		},
		Territories: cldr.Territories{
			territory.V_001: "swět",
			territory.V_002: "Afrika",
			territory.V_003: "Pódpołnocna Amerika",
			territory.V_005: "Pódpołdnjowa Amerika",
			territory.V_009: "Oceaniska",
			territory.V_011: "Pódwjacorna Afrika",
			territory.V_013: "Srjejźna Amerika",
			territory.V_014: "pódzajtšna Afrika",
			territory.V_015: "pódpołnocna Afrika",
			territory.V_017: "srjejźna Afrika",
			territory.V_018: "pódpołdnjowa Afrika",
			territory.V_019: "Amerika",
			territory.V_021: "pódpołnocny ameriski kontinent",
			territory.V_029: "Karibiska",
			territory.V_030: "pódzajtšna Azija",
			territory.V_034: "pódpołdnjowa Azija",
			territory.V_035: "krotkozajtšna Azija",
			territory.V_039: "pódpołdnjowa Europa",
			territory.V_053: "Awstralazija",
			territory.V_054: "Melaneziska",
			territory.V_057: "Mikroneziska (kupowy region)",
			territory.V_061: "Polyneziska",
			territory.V_142: "Azija",
			territory.V_143: "centralna Azija",
			territory.V_145: "pódwjacorna Azija",
			territory.V_150: "Europa",
			territory.V_151: "pódzajtšna Europa",
			territory.V_154: "pódpołnocna Europa",
			territory.V_155: "pódwjacorna Europa",
			territory.V_419: "Łatyńska Amerika",
			territory.AC:    "Ascension",
			territory.AD:    "Andorra",
			territory.AE:    "Zjadnośone arabiske emiraty",
			territory.AF:    "Afghanistan",
			territory.AG:    "Antigua a Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albańska",
			territory.AM:    "Armeńska",
			territory.AO:    "Angola",
			territory.AQ:    "Antarktis",
			territory.AR:    "Argentinska",
			territory.AS:    "Ameriska Samoa",
			territory.AT:    "Awstriska",
			territory.AU:    "Awstralska",
			territory.AW:    "Aruba",
			territory.AX:    "Åland",
			territory.AZ:    "Azerbajdžan",
			territory.BA:    "Bosniska a Hercegowina",
			territory.BB:    "Barbados",
			territory.BD:    "Bangladeš",
			territory.BE:    "Belgiska",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgarska",
			territory.BH:    "Bahrain",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "St. Barthélemy",
			territory.BM:    "Bermudy",
			territory.BN:    "Brunei",
			territory.BO:    "Boliwiska",
			territory.BQ:    "Karibiska Nižozemska",
			territory.BR:    "Brazilska",
			territory.BS:    "Bahamy",
			territory.BT:    "Bhutan",
			territory.BV:    "Bouvetowa kupa",
			territory.BW:    "Botswana",
			territory.BY:    "Běłoruska",
			territory.BZ:    "Belize",
			territory.CA:    "Kanada",
			territory.CC:    "Kokosowe kupy",
			territory.CD:    "Kongo (Demokratiska republika)",
			territory.CF:    "Centralnoafriska republika",
			territory.CG:    "Kongo (Republika)",
			territory.CH:    "Šwicarska",
			territory.CI:    "Słonowokósćowy pśibrjog",
			territory.CK:    "Cookowe kupy",
			territory.CL:    "Chilska",
			territory.CM:    "Kamerun",
			territory.CN:    "China",
			territory.CO:    "Kolumbiska",
			territory.CP:    "Clippertonowa kupa",
			territory.CR:    "Kosta Rika",
			territory.CU:    "Kuba",
			territory.CV:    "Kap Verde",
			territory.CW:    "Curaçao",
			territory.CX:    "Gódowne kupy",
			territory.CY:    "Cypriska",
			territory.CZ:    "Česka republika",
			territory.DE:    "Nimska",
			territory.DG:    "Diego Garcia",
			territory.DJ:    "Džibuti",
			territory.DK:    "Dańska",
			territory.DM:    "Dominika",
			territory.DO:    "Dominikańska republika",
			territory.DZ:    "Algeriska",
			territory.EA:    "Ceuta a Melilla",
			territory.EC:    "Ekwador",
			territory.EE:    "Estniska",
			territory.EG:    "Egyptojska",
			territory.EH:    "Pódwjacorna Sahara",
			territory.ER:    "Eritreja",
			territory.ES:    "Špańska",
			territory.ET:    "Etiopiska",
			territory.EU:    "Europska unija",
			territory.FI:    "Finska",
			territory.FJ:    "Fidži",
			territory.FK:    "Falklandske kupy (Malwiny)",
			territory.FM:    "Mikroneziska",
			territory.FO:    "Färöje",
			territory.FR:    "Francojska",
			territory.GA:    "Gabun",
			territory.GB:    "UK",
			territory.GD:    "Grenada",
			territory.GE:    "Georgiska",
			territory.GF:    "Francojska Guyana",
			territory.GG:    "Guernsey",
			territory.GH:    "Ghana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Grönlandska",
			territory.GM:    "Gambija",
			territory.GN:    "Gineja",
			territory.GP:    "Guadeloupe",
			territory.GQ:    "Ekwatorialna Gineja",
			territory.GR:    "Grichiska",
			territory.GS:    "Pódpołdnjowa Georgiska a Pódpołdnjowe Sandwichowe kupy",
			territory.GT:    "Guatemala",
			territory.GU:    "Guam",
			territory.GW:    "Gineja-Bissau",
			territory.GY:    "Guyana",
			territory.HK:    "Hongkong",
			territory.HM:    "Heardowa kupa a McDonaldowe kupy",
			territory.HN:    "Honduras",
			territory.HR:    "Chorwatska",
			territory.HT:    "Haiti",
			territory.HU:    "Hungorska",
			territory.IC:    "Kanariske kupy",
			territory.ID:    "Indoneziska",
			territory.IE:    "Irska",
			territory.IL:    "Israel",
			territory.IM:    "Man",
			territory.IN:    "Indiska",
			territory.IO:    "Britiski indiskooceaniski teritorium",
			territory.IQ:    "Irak",
			territory.IR:    "Iran",
			territory.IS:    "Islandska",
			territory.IT:    "Italska",
			territory.JE:    "Jersey",
			territory.JM:    "Jamaika",
			territory.JO:    "Jordaniska",
			territory.JP:    "Japańska",
			territory.KE:    "Kenia",
			territory.KG:    "Kirgizistan",
			territory.KH:    "Kambodža",
			territory.KI:    "Kiribati",
			territory.KM:    "Komory",
			territory.KN:    "St. Kitts a Nevis",
			territory.KP:    "Pódpołnocna Koreja",
			territory.KR:    "Pódpołdnjowa Koreja",
			territory.KW:    "Kuwait",
			territory.KY:    "Kajmaniske kupy",
			territory.KZ:    "Kazachstan",
			territory.LA:    "Laos",
			territory.LB:    "Libanon",
			territory.LC:    "St. Lucia",
			territory.LI:    "Liechtenstein",
			territory.LK:    "Sri Lanka",
			territory.LR:    "Liberija",
			territory.LS:    "Lesotho",
			territory.LT:    "Litawska",
			territory.LU:    "Luxemburgska",
			territory.LV:    "Letiska",
			territory.LY:    "Libyska",
			territory.MA:    "Marokko",
			territory.MC:    "Monaco",
			territory.MD:    "Moldawska",
			territory.ME:    "Carna Góra",
			territory.MF:    "St. Martin",
			territory.MG:    "Madagaskar",
			territory.MH:    "Marshallowe kupy",
			territory.ML:    "Mali",
			territory.MM:    "Myanmar",
			territory.MN:    "Mongolska",
			territory.MO:    "Macao",
			territory.MP:    "Pódpołnocne Mariany",
			territory.MQ:    "Martinique",
			territory.MR:    "Mawretańska",
			territory.MS:    "Montserrat",
			territory.MT:    "Malta",
			territory.MU:    "Mauritius",
			territory.MV:    "Malediwy",
			territory.MW:    "Malawi",
			territory.MX:    "Mexiko",
			territory.MY:    "Malajzija",
			territory.MZ:    "Mosambik",
			territory.NA:    "Namibija",
			territory.NC:    "Nowa Kaledoniska",
			territory.NE:    "Niger",
			territory.NF:    "Norfolkowa kupa",
			territory.NG:    "Nigerija",
			territory.NI:    "Nikaragua",
			territory.NL:    "Nižozemska",
			territory.NO:    "Norwegska",
			territory.NP:    "Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Nowoseelandska",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Francojska Polyneziska",
			territory.PG:    "Papua-Neuguinea",
			territory.PH:    "Filipiny",
			territory.PK:    "Pakistan",
			territory.PL:    "Pólska",
			territory.PM:    "St. Pierre a Miquelon",
			territory.PN:    "Pitcairnowe kupy",
			territory.PR:    "Puerto Rico",
			territory.PS:    "Palestina",
			territory.PT:    "Portugalska",
			territory.PW:    "Palau",
			territory.PY:    "Paraguay",
			territory.QA:    "Katar",
			territory.QO:    "wenkowna Oceaniska",
			territory.RE:    "Réunion",
			territory.RO:    "Rumuńska",
			territory.RS:    "Serbiska",
			territory.RU:    "Ruska",
			territory.RW:    "Ruanda",
			territory.SA:    "Saudi-Arabiska",
			territory.SB:    "Salomony",
			territory.SC:    "Seychelle",
			territory.SD:    "Sudan",
			territory.SE:    "Šwedska",
			territory.SG:    "Singapur",
			territory.SH:    "St. Helena",
			territory.SI:    "Słowjeńska",
			territory.SJ:    "Svalbard a Jan Mayen",
			territory.SK:    "Słowakska",
			territory.SL:    "Sierra Leone",
			territory.SM:    "San Marino",
			territory.SN:    "Senegal",
			territory.SO:    "Somalija",
			territory.SR:    "Surinamska",
			territory.SS:    "Pódpołdnjowy Sudan",
			territory.ST:    "São Tomé a Príncipe",
			territory.SV:    "El Salvador",
			territory.SX:    "Sint Maarten",
			territory.SY:    "Syriska",
			territory.SZ:    "Swasiska",
			territory.TA:    "Tristan da Cunha",
			territory.TC:    "Turks a Caicos kupy",
			territory.TD:    "Čad",
			territory.TF:    "Francojski pódpołdnjowy a antarktiski teritorium",
			territory.TG:    "Togo",
			territory.TH:    "Thailandska",
			territory.TJ:    "Tadźikistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Pódzajtšny Timor",
			territory.TM:    "Turkmeniska",
			territory.TN:    "Tuneziska",
			territory.TO:    "Tonga",
			territory.TR:    "Turkojska",
			territory.TT:    "Trinidad a Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tansanija",
			territory.UA:    "Ukraina",
			territory.UG:    "Uganda",
			territory.UM:    "Ameriska Oceaniska",
			territory.US:    "USA",
			territory.UY:    "Uruguay",
			territory.UZ:    "Uzbekistan",
			territory.VA:    "Vatikańske město",
			territory.VC:    "St. Vincent a Grenadiny",
			territory.VE:    "Venezuela",
			territory.VG:    "Britiske kněžniske kupy",
			territory.VI:    "Ameriske kněžniske kupy",
			territory.VN:    "Vietnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis a Futuna",
			territory.WS:    "Samoa",
			territory.XK:    "Kosowo",
			territory.YE:    "Jemen",
			territory.YT:    "Mayotte",
			territory.ZA:    "Pódpołdnjowa Afrika (Republika)",
			territory.ZM:    "Sambija",
			territory.ZW:    "Simbabwe",
			territory.ZZ:    "njeznaty region",
		},
	}
}
