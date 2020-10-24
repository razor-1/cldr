// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_sl() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "sl",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, dd. MMMM y", Long: "dd. MMMM y", Medium: "d. MMM y", Short: "d. MM. yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "feb.", Mar: "mar.", Apr: "apr.", May: "maj", Jun: "jun.", Jul: "jul.", Aug: "avg.", Sep: "sep.", Oct: "okt.", Nov: "nov.", Dec: "dec."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "j", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "j", Jul: "j", Aug: "a", Sep: "s", Oct: "o", Nov: "n", Dec: "d"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "marec", Apr: "april", May: "maj", Jun: "junij", Jul: "julij", Aug: "avgust", Sep: "september", Oct: "oktober", Nov: "november", Dec: "december"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ned.", Mon: "pon.", Tue: "tor.", Wed: "sre.", Thu: "čet.", Fri: "pet.", Sat: "sob."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "p", Tue: "t", Wed: "s", Thu: "č", Fri: "p", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ned.", Mon: "pon.", Tue: "tor.", Wed: "sre.", Thu: "čet.", Fri: "pet.", Sat: "sob."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "nedelja", Mon: "ponedeljek", Tue: "torek", Wed: "sreda", Thu: "četrtek", Fri: "petek", Sat: "sobota"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "dop.", PM: "pop."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "d", PM: "p"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "dopoldne", PM: "popoldne"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤;(#,##0.00\u00a0¤)", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "andorska peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "dirham Združenih arabskih emiratov", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "stari afganistanski afgani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afgani", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "albanski lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "armenski dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "nizozemsko-antilski gulden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "angolska kvanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "stara angolska kvanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "angolska nova kvanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "konvertibilna angolska kvanza (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "argentinski avstral", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "argentinski peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "argentinski peso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "avstrijski šiling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "avstralski dolar", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "arubski florin", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "stari azerbajdžanski manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "azerbajdžanski manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "bosansko-hercegovski dinar", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "bosansko-hercegovska konvertibilna marka", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "barbadoški dolar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "bangladeška taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "belgijski konvertibilni frank", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "belgijski frank", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "belgijski finančni frank", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "stari bolgarski lev", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "bolgarski lev", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "bahranski dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "burundski frank", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "bermudski dolar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "brunejski dolar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "bolivijski boliviano", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "bolivijski peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "bolivijski mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "brazilski novi kruzeiro (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "brazilski kruzado", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "stari brazilski kruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "brazilski real", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "novi brazilski kruzado", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "brazilski kruzeiro", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "bahamski dolar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "butanski ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "burmanski kjat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "bocvanska pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "beloruski novi rubelj (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "beloruski rubelj", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "beloruski rubelj (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "belizejski dolar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "kanadski dolar", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "kongoški frank", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "evro WIR", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "švicarski frank", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "frank WIR", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "čilski unidades de fomento", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "čilski peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "kitajski juan (offshore)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "kitajski juan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "kolumbijski peso", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "kolumbijska enota realne vrednosti", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "kostariški kolon", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "stari srbski dinar", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "češkoslovaška krona", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "kubanski konvertibilni peso", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "kubanski peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "zelenortski eskudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "ciprski funt", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "češka krona", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "vzhodnonemška marka", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "nemška marka", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "džibutski frank", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "danska krona", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "dominikanski peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "alžirski dinar", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "ekvadorski sukre", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "ekvadorska enota realne vrednosti (UVC)", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "estonska krona", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "egiptovski funt", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "eritrejska nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "španska pezeta (račun A)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "španska pezeta (račun B)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "španska pezeta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "etiopski bir", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "evro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "finska marka", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "fidžijski dolar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "falklandski funt", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "francoski frank", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "britanski funt", Symbol: "GBP"},
				currency.GEK: cldr.Currency{DisplayName: "gruzijski bon lari", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "gruzijski lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "stari ganski cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "ganski cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "gibraltarski funt", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "gambijski dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "gvinejski frank", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "gvinejski sili", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "ekwele Ekvatorialne Gvineje", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "grška drahma", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "gvatemalski kecal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "eskudo Portugalske Gvineje", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "peso Gvineje Bissau", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "gvajanski dolar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "hongkonški dolar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "honduraška lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "hrvaški dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "hrvaška kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "haitski gurd", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "madžarski forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "indonezijska rupija", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "irski funt", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "izraelski funt", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "izraelski šekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "indijska rupija", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "iraški dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "iranski rial", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "islandska krona", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "italijanska lira", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "jamajški dolar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "jordanski dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japonski jen", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "kenijski šiling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "kirgiški som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "kamboški riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "komorski frank", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "severnokorejski von", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "južnokorejski von", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "kuvajtski dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "kajmanski dolar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "kazahstanski tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "laoški kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libanonski funt", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "šrilanška rupija", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "liberijski dolar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "lesoški loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "litovski litas", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "litvanski litas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "luksemburški konvertibilni frank", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "luksemburški frank", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "luksemburški finančni frank", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "latvijski lats", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "latvijski rubelj", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "libijski dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "maroški dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "maroški frank", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldavijski leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "malgaški ariarij", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "malgaški frank", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "makedonski denar", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "malijski frank", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "mjanmarski kjat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "mongolski tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "makavska pataka", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "mavretanska uguija (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "mavretanska uguija", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "malteška lira", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "malteški funt", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "mavricijska rupija", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "maldivska rufija", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "malavijska kvača", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "mehiški peso", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "mehiški srebrni peso (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "mehiška inverzna enota (UDI)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "malezijski ringit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "mozambiški eskudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "stari mozambiški metikal", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "mozambiški metikal", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "namibijski dolar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "nigerijska naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "nikaraška kordova", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "nikaraška zlata kordova", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "nizozemski gulden", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "norveška krona", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "nepalska rupija", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "novozelandski dolar", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "omanski rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "panamska balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "perujski inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "perujski sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "perujski sol (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "kina Papue Nove Gvineje", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "filipinski peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "pakistanska rupija", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "poljski novi zlot", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "stari poljski zlot (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "portugalski eskudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "paragvajski gvarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "katarski rial", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "rodezijski dolar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "stari romunski leu", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "romunski leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "srbski dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "ruski rubelj", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "ruski rubelj (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ruandski frank", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "saudski rial", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "solomonski dolar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "sejšelska rupija", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "stari sudanski dinar", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sudanski funt", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "stari sudanski funt", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "švedska krona", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "singapurski dolar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "funt Sv. Helene", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "slovenski tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "slovaška krona", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "sieraleonski leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "somalski šiling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "surinamski dolar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "surinamski gulden", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "južnosudanski funt", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "saotomejska dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "saotomejska dobra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "sovjetski rubelj", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "salvadorski kolon", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "sirijski funt", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "svazijski lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "tajski baht", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "tadžikistanski rubelj", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "tadžikistanski somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "turkmenski manat", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "turkmenistanski novi manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "tunizijski dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "tongovska paanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "timorski eskudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "stara turška lira", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "nova turška lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "dolar Trinidada in Tobaga", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "novi tajvanski dolar", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "tanzanijski šiling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ukrajinska grivna", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "ukrajinski karbovanci", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "stari ugandski šiling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "ugandski šiling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "ameriški dolar", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "ameriški dolar, naslednji dan", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "ameriški dolar, isti dan", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "stari urugvajski peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "urugvajski peso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "uzbeški sum", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "venezuelski bolivar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "venezuelski bolivar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "venezuelski bolivar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "vientnamski dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "vanuatujski vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "samoanska tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "CFA frank BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "srebro", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "zlato", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "evropska sestavljena enota", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "evropska monetarna enota", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "evropska obračunska enota (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "evropska obračunska enota (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "vzhodnokaribski dolar", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "posebne pravice črpanja", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "evropska denarna enota", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "zlati frank", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "frank UIC", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "zahodnoafriški frank CFA", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "paladij", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP frank", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platina", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "koda za potrebe testiranja", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "neznana valuta", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "jemenski dinar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "jemenski rial", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "stari jugoslovanski dinar", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "novi jugoslovanski dinar", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "jugoslovanski konvertibilni dinar", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "južnoafriški finančni rand", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "južnoafriški rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "zambijska kvača (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "zambijska kvača", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "zairski novi zaire", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "zairski zaire", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "zimbabvejski dolar", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "zimbabvejski dolar (2009)", Symbol: ""},
			},
		},
	}
}