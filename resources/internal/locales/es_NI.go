// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_es_NI() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "es_NI",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ene.", Feb: "feb.", Mar: "mar.", Apr: "abr.", May: "may.", Jun: "jun.", Jul: "jul.", Aug: "ago.", Sep: "sept.", Oct: "oct.", Nov: "nov.", Dec: "dic."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "E", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "enero", Feb: "febrero", Mar: "marzo", Apr: "abril", May: "mayo", Jun: "junio", Jul: "julio", Aug: "agosto", Sep: "septiembre", Oct: "octubre", Nov: "noviembre", Dec: "diciembre"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dom.", Mon: "lun.", Tue: "mar.", Wed: "mié.", Thu: "jue.", Fri: "vie.", Sat: "sáb."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "X", Thu: "J", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "DO", Mon: "LU", Tue: "MA", Wed: "MI", Thu: "JU", Fri: "VI", Sat: "SA"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "domingo", Mon: "lunes", Tue: "martes", Wed: "miércoles", Thu: "jueves", Fri: "viernes", Sat: "sábado"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.\u00a0m.", PM: "p.\u00a0m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a.\u00a0m.", PM: "p.\u00a0m."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "a.\u00a0m.", PM: "p.\u00a0m."}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "peseta andorrana", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "dírham de los Emiratos Árabes Unidos", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "afgani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afgani", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "florín de las Antillas Neerlandesas", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "kuanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "kwanza angoleño (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "nuevo kwanza angoleño (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "kwanza reajustado angoleño (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "austral argentino", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "peso argentino (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "peso argentino", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "chelín austriaco", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "dólar australiano", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "florín arubeño", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "manat azerí (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "manat azerbaiyano", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "dinar bosnio", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "marco convertible de Bosnia y Herzegovina", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "dólar barbadense", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "franco belga (convertible)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "franco belga", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "franco belga (financiero)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "lev fuerte búlgaro", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "lev búlgaro", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "dinar bahreiní", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "franco burundés", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "dólar de Bermudas", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "dólar bruneano", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "peso boliviano", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "MVDOL boliviano", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "nuevo cruceiro brasileño (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "cruzado brasileño", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "cruceiro brasileño (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "real brasileño", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "nuevo cruzado brasileño", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "cruceiro brasileño", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "dólar bahameño", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "gultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "kyat birmano", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "nuevo rublo bielorruso (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "rublo bielorruso", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "rublo bielorruso (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "dólar beliceño", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "dólar canadiense", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "franco congoleño", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "euro WIR", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "franco suizo", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "franco WIR", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "unidad de fomento chilena", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "peso chileno", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "yuan chino (extracontinental)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "yuan", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "peso colombiano", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "unidad de valor real colombiana", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "colón costarricense", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "antiguo dinar serbio", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "corona fuerte checoslovaca", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "peso cubano convertible", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "peso cubano", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "escudo de Cabo Verde", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "libra chipriota", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "corona checa", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "ostmark de Alemania del Este", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "marco alemán", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "franco yibutiano", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "corona danesa", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "peso dominicano", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "dinar argelino", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "sucre ecuatoriano", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "unidad de valor constante (UVC) ecuatoriana", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "corona estonia", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "libra egipcia", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "peseta española (cuenta A)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "peseta española (cuenta convertible)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "peseta española", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "bir", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "marco finlandés", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "dólar fiyiano", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "libra malvinense", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "franco francés", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "libra esterlina", Symbol: "GBP"},
				currency.GEK: cldr.Currency{DisplayName: "kupon larit georgiano", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "cedi ghanés (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "libra gibraltareña", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "franco guineano", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "syli guineano", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "ekuele de Guinea Ecuatorial", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "dracma griego", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "quetzal guatemalteco", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "escudo de Guinea Portuguesa", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "peso de Guinea-Bissáu", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "dólar guyanés", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "dólar hongkonés", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "lempira hondureño", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "dinar croata", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "gourde haitiano", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "forinto húngaro", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "rupia indonesia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "libra irlandesa", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "libra israelí", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "nuevo séquel israelí", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "rupia india", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "dinar iraquí", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "rial iraní", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "corona islandesa", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "lira italiana", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "dólar jamaicano", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "dinar jordano", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "yen", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "chelín keniano", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "franco comorense", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "won norcoreano", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "won surcoreano", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "dinar kuwaití", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "dólar de las Islas Caimán", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "tenge kazako", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libra libanesa", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "rupia esrilanquesa", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "dólar liberiano", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "loti lesothense", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "litas lituano", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "talonas lituano", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "franco convertible luxemburgués", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "franco luxemburgués", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "franco financiero luxemburgués", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "lats letón", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "rublo letón", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "dinar libio", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "dírham marroquí", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "franco marroquí", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "leu moldavo", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ariari", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "franco malgache", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "dinar macedonio", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "franco malí", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "kiat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "pataca de Macao", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "uguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "uguiya", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "lira maltesa", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "libra maltesa", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "rupia mauriciana", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "rufiya", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "kuacha malauí", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "peso mexicano", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "peso de plata mexicano (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "unidad de inversión (UDI) mexicana", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "ringit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "escudo mozambiqueño", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "antiguo metical mozambiqueño", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "dólar namibio", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "córdoba nicaragüense (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "córdoba nicaragüense", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "florín neerlandés", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "corona noruega", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "rupia nepalí", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "dólar neozelandés", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "rial omaní", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "balboa panameño", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "inti peruano", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "sol peruano", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "sol peruano (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "peso filipino", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "rupia pakistaní", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "esloti", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "zloty polaco (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "escudo portugués", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "guaraní paraguayo", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "rial catarí", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "dólar rodesiano", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "antiguo leu rumano", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "leu rumano", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "dinar serbio", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "rublo ruso", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "rublo ruso (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "franco ruandés", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "rial saudí", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "dólar salomonense", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "rupia seychellense", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "dinar sudanés", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "libra sudanesa", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "libra sudanesa antigua", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "corona sueca", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "dólar singapurense", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "libra de Santa Elena", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "tólar esloveno", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "corona eslovaca", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "leona", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "chelín somalí", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "dólar surinamés", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "florín surinamés", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "libra sursudanesa", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "dobra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "rublo soviético", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "colón salvadoreño", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "libra siria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "bat", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "rublo tayiko", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "somoni tayiko", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "manat turcomano (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "manat turcomano", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "dinar tunecino", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "paanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "escudo timorense", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "lira turca (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "lira turca", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "dólar de Trinidad y Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "nuevo dólar taiwanés", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "chelín tanzano", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "grivna", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "karbovanet ucraniano", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "chelín ugandés (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "chelín ugandés", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "dólar estadounidense", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "dólar estadounidense (día siguiente)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "dólar estadounidense (mismo día)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "peso uruguayo en unidades indexadas", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "peso uruguayo (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "peso uruguayo", Symbol: "UYU"},
				currency.UYW: cldr.Currency{DisplayName: "unidad previsional uruguayo", Symbol: "UYW"},
				currency.UZS: cldr.Currency{DisplayName: "sum", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "bolívar venezolano (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "bolívar venezolano (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "bolívar venezolano", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "franco CFA de África Central", Symbol: "XAF"},
				currency.XAG: cldr.Currency{DisplayName: "plata", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "oro", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "unidad compuesta europea", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "unidad monetaria europea", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "unidad de cuenta europea (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "unidad de cuenta europea (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "dólar del Caribe Oriental", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "derechos especiales de giro", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "unidad de moneda europea", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "franco oro francés", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "franco UIC francés", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "franco CFA de África Occidental", Symbol: "XOF"},
				currency.XPD: cldr.Currency{DisplayName: "paladio", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "franco CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platino", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "fondos RINET", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "código reservado para pruebas", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "moneda desconocida", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "dinar yemení", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "rial yemení", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "dinar fuerte yugoslavo", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "super dinar yugoslavo", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "dinar convertible yugoslavo", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "rand sudafricano (financiero)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "kwacha zambiano (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kuacha zambiano", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "nuevo zaire zaireño", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "zaire zaireño", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "dólar de Zimbabue", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "dólar zimbabuense", Symbol: ""},
			},
		},
	}
}
