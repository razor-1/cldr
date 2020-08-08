// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_gl() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{0} 'do' {1}", Long: "{0} 'do' {1}", Medium: "{0}, {1}", Short: "{0}, {1}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Xan.", Feb: "Feb.", Mar: "Mar.", Apr: "Abr.", May: "Maio", Jun: "Xuño", Jul: "Xul.", Aug: "Ago.", Sep: "Set.", Oct: "Out.", Nov: "Nov.", Dec: "Dec."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "X", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "X", Jul: "X", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Xaneiro", Feb: "Febreiro", Mar: "Marzo", Apr: "Abril", May: "Maio", Jun: "Xuño", Jul: "Xullo", Aug: "Agosto", Sep: "Setembro", Oct: "Outubro", Nov: "Novembro", Dec: "Decembro"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dom.", Mon: "Luns", Tue: "Mar.", Wed: "Mér.", Thu: "Xov.", Fri: "Ven.", Sat: "Sáb."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "X", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Do", Mon: "Lu", Tue: "Ma", Wed: "Mé", Thu: "Xo", Fri: "Ve", Sat: "Sá"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Domingo", Mon: "Luns", Tue: "Martes", Wed: "Mércores", Thu: "Xoves", Fri: "Venres", Sat: "Sábado"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_gl]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "peseta andorrana", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "dirham dos Emiratos Árabes Unidos", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "afgani afgán", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "lek albanés", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "dram armenio", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "florín das Antillas Neerlandesas", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angolano", Symbol: "AOA"},
				currency.ARP: cldr.Currency{DisplayName: "Peso arxentino (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "peso arxentino", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "dólar australiano", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "florín de Aruba", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "manat acerbaixano", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "marco convertible de Bosnia e Hercegovina", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "dólar de Barbados", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "taka de Bangladesh", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Franco belga (convertible)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Franco belga", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Franco belga (financeiro)", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "lev búlgaro", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "dinar de Bahrain", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "franco burundiano", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "dólar bermudano", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "dólar de Brunei", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "Peso boliviano", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "MVDOL boliviano", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Cruzeiro novo brasileiro (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Cruzado brasileiro", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Cruzeiro brasileiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "real brasileiro", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Cruzado novo brasileiro", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Cruzeiro brasileiro", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "dólar bahamés", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ngultrum butanés", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "pula botswaniano", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "rublo belaruso", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Rublo bielorruso (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "dólar belizense", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "dólar canadense", Symbol: "$"},
				currency.CDF: cldr.Currency{DisplayName: "franco congolés", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "franco suízo", Symbol: "CHF"},
				currency.CLF: cldr.Currency{DisplayName: "Unidades de fomento chilenas", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "peso chileno", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "iuán chinés (extracontinental)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "iuán chinés", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "peso colombiano", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "colón costarriqueño", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "peso cubano convertible", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "peso cubano", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "escudo caboverdiano", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "coroa checa", Symbol: "CZK"},
				currency.DEM: cldr.Currency{DisplayName: "Marco alemán", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "franco djibutiano", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "coroa dinamarquesa", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "peso dominicano", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "dinar alxeriano", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Sucre ecuatoriano", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Unidade de valor constante ecuatoriana", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "libra exipcia", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa eritreo", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Peseta española (conta A)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Peseta española (conta convertible)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Peseta española", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr etíope", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "dólar fixiano", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "libra das Illas Malvinas", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Franco francés", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "libra esterlina", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "lari xeorxiano", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "cedi ghanés", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "libra xibraltareña", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi gambiano", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "franco guineano", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Syli guineano", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Ekwele guineana", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Dracma grego", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "quetzal guatemalteco", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "dólar güianés", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "dólar de Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "lempira hondureño", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "kuna croata", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "gourde haitiana", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "florín húngaro", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "rupia indonesia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Libra irlandesa", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "novo shequel israelí", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "rupia india", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "dinar iraquí", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "rial iraniano", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "coroa islandesa", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Lira italiana", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "dólar xamaicano", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "dinar xordano", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "ien xaponés", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "xilin kenyano", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "som kirguiz", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "riel camboxano", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "franco comoriano", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "won norcoreano", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "won surcoreano", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "dinar kuwaití", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "dólar das Illas Caimán", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "tenge kazako", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "kip laosiano", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libra libanesa", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "rupia srilankesa", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "dólar liberiano", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Loti de Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litas lituana", Symbol: "LTL"},
				currency.LUC: cldr.Currency{DisplayName: "Franco convertible luxemburgués", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Franco luxemburgués", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Franco financeiro luxemburgués", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Lats letón", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "dinar libio", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "dirham marroquí", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Franco marroquí", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Leu moldavo", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ariary malgaxe", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "dinar macedonio", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "kyat birmano", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "tugrik mongol", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "pataca macaense", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya mauritano (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "ouguiya mauritano", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "rupia mauriciana", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "rupia maldivana", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha de Malawi", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "peso mexicano", Symbol: "$MX"},
				currency.MXP: cldr.Currency{DisplayName: "Peso de prata mexicano (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Unidade de inversión mexicana", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "ringgit malaio", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "metical mozambicano", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "dólar namibio", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "naira nixeriano", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Córdoba nicaragüense", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "córdoba nicaraguano", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Florín holandés", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "coroa norueguesa", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "rupia nepalesa", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "dólar neozelandés", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "rial omaní", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "balboa panameño", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Inti peruano", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "sol peruano", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Sol peruano (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "kina de Papúa-Nova Guinea", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "peso filipino", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "rupia paquistaní", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "zloty polaco", Symbol: "PLN"},
				currency.PTE: cldr.Currency{DisplayName: "Escudo portugués", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "guaraní paraguaio", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "rial qatarí", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "leu romanés", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "dinar serbio", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "rublo ruso", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Rublo ruso (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "franco ruandés", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "rial saudita", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "dólar das Illas Salomón", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "rupia de Seychelles", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "libra sudanesa", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "coroa sueca", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "dólar de Singapur", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "libra de Santa Helena", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "leone de Serra Leoa", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "xilin somalí", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "dólar surinamés", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "libra sursudanesa", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Dobra de São Tomé e Príncipe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "dobra de San Tomé e Príncipe", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Rublo soviético", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Colón salvadoreño", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "libra siria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni de Eswatini", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "baht tailandés", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "somoni taxiquistano", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "manat turkmeno", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "dinar tunisiano", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "paʻanga tongano", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "lira turca", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "dólar trinitense", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "novo dólar taiwanés", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "xilin tanzano", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "hrivna ucraína", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "xilin ugandés", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "dólar estadounidense", Symbol: "$"},
				currency.UYI: cldr.Currency{DisplayName: "Peso en unidades indexadas uruguaio", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Peso uruguaio (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "peso uruguaio", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "som uzbeko", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Bolívar venezolano (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Bolívar venezolano (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "bolívar venezolano", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "dong vietnamita", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "vatu vanuatiano", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "tala samoano", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "franco CFA (BEAC)", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Prata", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Ouro", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "dólar do Caribe Oriental", Symbol: "XCD"},
				currency.XOF: cldr.Currency{DisplayName: "franco CFA (BCEAO)", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Paladio", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "franco CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platino", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "moeda descoñecida", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "rial iemení", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "rand surafricano", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha zambiano (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha zambiano", Symbol: "ZMW"},
			},
		},
	}
}
