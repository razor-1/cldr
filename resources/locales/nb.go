// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_nb() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} 'kl'. {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "feb", Mar: "mar", Apr: "apr", May: "mai", Jun: "jun", Jul: "jul", Aug: "aug", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "mars", Apr: "april", May: "mai", Jun: "juni", Jul: "juli", Aug: "august", Sep: "september", Oct: "oktober", Nov: "november", Dec: "desember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "søn.", Mon: "man.", Tue: "tir.", Wed: "ons.", Thu: "tor.", Fri: "fre.", Sat: "lør."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "O", Thu: "T", Fri: "F", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "sø.", Mon: "ma.", Tue: "ti.", Wed: "on.", Thu: "to.", Fri: "fr.", Sat: "lø."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "søndag", Mon: "mandag", Tue: "tirsdag", Wed: "onsdag", Thu: "torsdag", Fri: "fredag", Sat: "lørdag"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_nb]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "\u00a0", Negative: "\u061c−", Percent: "٪\u061c", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "andorranske pesetas", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "emiratarabiske dirham", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "afgansk afghani (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "afghanske afghani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "albanske lek (1946–1965)", Symbol: "ALK"},
				currency.ALL: cldr.Currency{DisplayName: "albanske lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "armenske dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "nederlandske antillegylden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "angolanske kwanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "angolanske kwanza (1977–1990)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "angolanske nye kwanza (1990–2000)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "angolanske omjusterte kwanza (1995–1999)", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "argentinske australer", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "argentinske peso ley", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "argentinsk pesos (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "argentinske pesos (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "argentinske pesos", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "østerrikske shilling", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "australske dollar", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "arubiske floriner", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "aserbajdsjanske manat (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "aserbajdsjanske manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "bosnisk-hercegovinske dinarer (1992–1994)", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "bosnisk-hercegovinske konvertible mark", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "nye bosnisk-hercegovinske dinarer (1994–1997)", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "barbadiske dollar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "bangladeshiske taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "belgiske franc (konvertible)", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "belgiske franc", Symbol: "BEF"},
				currency.BEL: cldr.Currency{DisplayName: "belgiske franc (finansielle)", Symbol: "BEL"},
				currency.BGL: cldr.Currency{DisplayName: "bulgarske lev (hard)", Symbol: "BGL"},
				currency.BGM: cldr.Currency{DisplayName: "bulgarske lev (sosialist)", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "bulgarske lev", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "bulgarske lev (1879–1952)", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "bahrainske dinarer", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "burundiske franc", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "bermudiske dollar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "bruneiske dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "bolivianske boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "bolivianske boliviano (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "bolivianske pesos", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "bolivianske mvdol", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "brasilianske cruzeiro novo (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "brasilianske cruzados (1986–1989)", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "brasilianske cruzeiro (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "brasilianske real", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "brasilianske cruzado novo (1989–1990)", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "brasilianske cruzeiro (1993–1994)", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "brasilianske cruzeiro (1942–1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "bahamanske dollar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "bhutanske ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "burmesiske kyat", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "botswanske pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "hviterussiske nye rubler (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "nye hviterussiske rubler", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "hviterussiske rubler (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "beliziske dollar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "kanadiske dollar", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "kongolesiske franc", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR euro", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "sveitsiske franc", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR franc", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "chilenske escudo", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "chilenske unidades de fomento", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "chilenske pesos", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "kinesiske yuan (offshore)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "Kinas folkebank dollar", Symbol: "CNX"},
				currency.CNY: cldr.Currency{DisplayName: "kinesiske yuan", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "colombianske pesos", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "colombianske unidad de valor real", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "costaricanske colón", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "serbiske dinarer (2002–2006)", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "tsjekkoslovakiske koruna (hard)", Symbol: "CSK"},
				currency.CUC: cldr.Currency{DisplayName: "kubanske konvertible pesos", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "kubanske pesos", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "kappverdiske escudos", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "kypriotiske pund", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "tsjekkiske koruna", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "østtyske mark", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "tyske mark", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "djiboutiske franc", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "danske kroner", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "dominikanske pesos", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "algeriske dinarer", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "ecuadorianske sucre", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "ecuadorianske unidad de valor constante (UVC)", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "estiske kroon", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "egyptiske pund", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "eritreiske nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "spanske peseta (A–konto)", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "spanske peseta (konvertibel konto)", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "spanske peseta", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "etiopiske birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "finske mark", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "fijianske dollar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "falklandspund", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "franske franc", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "britiske pund", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "georgiske kupon larit", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "georgiske lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "ghanesisk cedi (1979–2007)", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "ghanesiske cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "gibraltarske pund", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "gambiske dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "guineanske franc", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "guineanske syli", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "ekvatorialguineanske ekwele guineana", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "greske drakmer", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "guatemalanske quetzal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "portugisiske guinea escudo", Symbol: "GWE"},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissau-pesos", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "guyanske dollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkong-dollar", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "honduranske lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "kroatiske dinarer", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "kroatiske kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "haitiske gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ungarske forinter", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "indonesiske rupier", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "irske pund", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "israelske pund", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "israelske shekler (1980–1985)", Symbol: "ILR"},
				currency.ILS: cldr.Currency{DisplayName: "nye israelske shekler", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "indiske rupier", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "irakske dinarer", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "iranske rialer", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "islandske kroner (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "islandske kroner", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "italienske lire", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "jamaikanske dollar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "jordanske dinarer", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japanske yen", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "kenyanske shilling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "kirgisiske som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "kambodsjanske riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "komoriske franc", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "nordkoreanske won", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "sørkoreanske hwan (1953–1962)", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "sørkoreanske won (1945–1953)", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "sørkoreanske won", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "kuwaitiske dinarer", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "caymanske dollar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "kasakhstanske tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "laotiske kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libanesiske pund", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "srilankiske rupier", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "liberiske dollar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "lesothiske loti", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "litauiske litas", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "litauiske talonas", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "luxemburgske konvertible franc", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "luxemburgske franc", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "luxemburgske finansielle franc", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "latviske lats", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "latviske rubler", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "libyske dinarer", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "marokkanske dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "marokkanske franc", Symbol: "MAF"},
				currency.MCF: cldr.Currency{DisplayName: "monegaskiske franc", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "moldovske cupon", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "moldovske leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "madagassiske ariary", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "madagassiske franc", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "makedonske denarer", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "makedonske denarer (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "maliske franc", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "myanmarske kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "mongolske tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "makaoiske pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "mauritanske ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "mauritanske ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "maltesiske lira", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "maltesiske pund", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "mauritiske rupier", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "maldiviske rupier", Symbol: "MVP"},
				currency.MVR: cldr.Currency{DisplayName: "maldiviske rufiyaa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "malawiske kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "meksikanske pesos", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "meksikanske sølvpesos (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "meksikanske unidad de inversion (UDI)", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "malaysiske ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "mosambikiske escudo", Symbol: "MZE"},
				currency.MZM: cldr.Currency{DisplayName: "gamle mosambikiske metical", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "mosambikiske metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "namibiske dollar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "nigerianske naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "nicaraguanske cordoba (1988–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "nicaraguanske córdoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "nederlandske gylden", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "norske kroner", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "nepalske rupier", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "newzealandske dollar", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "omanske rialer", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "panamanske balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "peruanske inti", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "peruanske sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "peruanske sol (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "papuanske kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "filippinske pesos", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "pakistanske rupier", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "polske zloty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "polske zloty (1950–1995)", Symbol: "PLZ"},
				currency.PTE: cldr.Currency{DisplayName: "portugisiske escudo", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "paraguayanske guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "qatarske rialer", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "rhodesiske dollar", Symbol: "RHD"},
				currency.ROL: cldr.Currency{DisplayName: "rumenske leu (1952–2006)", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "rumenske leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "serbiske dinarer", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "russiske rubler", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "russiske rubler (1991–1998)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "rwandiske franc", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "saudiarabiske riyaler", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "salomonske dollar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "seychelliske rupier", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "sudanesiske dinarer (1992–2007)", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "sudanske pund", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "sudanesiske pund", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "svenske kroner", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "singaporske dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "sankthelenske pund", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "slovenske tolar", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "slovakiske koruna", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "sierraleonske leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "somaliske shilling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "surinamske dollar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "surinamske gylden", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "sørsudanske pund", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "saotomesiske dobra (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "saotomesiske dobra", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "sovjetiske rubler", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "salvadoranske colon", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "syriske pund", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "swazilandske lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "thailandske baht", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "tadsjikiske rubler", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "tadsjikiske somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "turkmenske manat (1993–2009)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "turkmenske manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "tunisiske dinarer", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "tonganske paʻanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "timoresiske escudo", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "tyrkiske lire (1922–2005)", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "tyrkiske lire", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "trinidadiske dollar", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "nye taiwanske dollar", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "tanzanianske shilling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ukrainske hryvnia", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "ukrainske karbovanetz", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "ugandiske shilling (1966–1987)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "ugandiske shilling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "amerikanske dollar", Symbol: "USD"},
				currency.USN: cldr.Currency{DisplayName: "amerikanske dollar (neste dag)", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "amerikanske dollar (samme dag)", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "uruguyanske pesos (indekserte enheter)", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "uruguayanske pesos (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "uruguayanske pesos", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "usbekiske som", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "venezuelanske bolivar (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "venezuelanske bolivar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "venezuelanske bolivar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "vietnamesiske dong", Symbol: "VND"},
				currency.VNN: cldr.Currency{DisplayName: "vietnamesiske dong (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "vanuatiske vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "samoanske tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "sentralafrikanske CFA-franc", Symbol: "XAF"},
				currency.XAG: cldr.Currency{DisplayName: "sølv", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "gull", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "europeisk sammensatt enhet", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "europeisk monetær enhet", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "europeisk kontoenhet (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "europeisk kontoenhet (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "østkaribiske dollar", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "spesielle trekkrettigheter", Symbol: "XDR"},
				currency.XEU: cldr.Currency{DisplayName: "europeisk valutaenhet", Symbol: "XEU"},
				currency.XFO: cldr.Currency{DisplayName: "franske gullfranc", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "franske UIC-franc", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "vestafrikanske CFA-franc", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "palladium", Symbol: "XPD"},
				currency.XPF: cldr.Currency{DisplayName: "CFP-franc", Symbol: "XPF"},
				currency.XPT: cldr.Currency{DisplayName: "platina", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "RINET-fond", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "sucre", Symbol: "XSU"},
				currency.XTS: cldr.Currency{DisplayName: "testvalutakode", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "ADB-kontoenhet", Symbol: "XUA"},
				currency.XXX: cldr.Currency{DisplayName: "ukjent valuta", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "jemenittiske dinarer", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "jemenittiske rialer", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "jugoslaviske dinarer (hard)", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "jugoslaviske noviy-dinarer", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "jugoslaviske konvertible dinarer", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "jugoslaviske reformerte dinarer (1992–1993)", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "sørafrikanske rand (finansielle)", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "sørafrikanske rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "zambiske kwacha (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "zambiske kwacha", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "zairiske nye zaire", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "zairiske zaire", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "zimbabwiske dollar (1980–2008)", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "zimbabwisk dollar (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "zimbabwisk dollar (2008)", Symbol: ""},
			},
		},
	}
}
