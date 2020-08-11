// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_da() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "da",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE 'den' d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd.MM.y"}, Time: cldr.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'kl'. {0}", Long: "{1} 'kl'. {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "feb.", Mar: "mar.", Apr: "apr.", May: "maj", Jun: "jun.", Jul: "jul.", Aug: "aug.", Sep: "sep.", Oct: "okt.", Nov: "nov.", Dec: "dec."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "marts", Apr: "april", May: "maj", Jun: "juni", Jul: "juli", Aug: "august", Sep: "september", Oct: "oktober", Nov: "november", Dec: "december"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "søn", Mon: "man", Tue: "tir", Wed: "ons", Thu: "tor", Fri: "fre", Sat: "lør"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "O", Thu: "T", Fri: "F", Sat: "L"}, Short: cldr.CalendarDayFormatNameValue{Sun: "sø", Mon: "ma", Tue: "ti", Wed: "on", Thu: "to", Fri: "fr", Sat: "lø"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "søndag", Mon: "mandag", Tue: "tirsdag", Wed: "onsdag", Thu: "torsdag", Fri: "fredag", Sat: "lørdag"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorransk peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "dirham fra de Forenede Arabiske Emirater", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afghansk afghani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afghansk afghani", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "albansk lek (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "albansk lek", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "armensk dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Nederlandske Antiller-gylden", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "angolansk kwanza", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Angolansk kwanza (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Angolansk nye kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Angolansk kwanza (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Argentinsk austral", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "argentinsk peso ley (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "argentinsk peso (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentinsk peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "argentinsk peso", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Østrigsk schilling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "australsk dollar", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "arubansk florin", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Aserbajdsjansk manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "aserbajdsjansk manat", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Bosnien-Hercegovinsk dinar", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "bosnien-hercegovinsk konvertibel mark", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "ny bosnien-hercegovinsk dinar (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "barbadisk dollar", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "bangladeshisk taka", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belgisk franc (konvertibel)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Belgisk franc", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Belgisk franc (financial)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Bulgarsk hard lev", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "bulgarsk socialistisk lev", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "bulgarsk lev", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "bulgarsk lev (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "bahrainsk dinar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "burundisk franc", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "bermudansk dollar", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "bruneisk dollar", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "boliviansk boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "boliviansk boliviano (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Boliviansk peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Boliviansk mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Brasiliansk cruzeiro novo (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Brasiliansk cruzado (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Brasiliansk cruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "brasiliansk real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Brasiliansk cruzado novo (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Brasiliansk cruzeiro (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "brasiliansk cruzeiro (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "bahamansk dollar", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "bhutansk ngultrum", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Burmesisk kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "botswansk pula", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Hviderussisk rubel (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "hviderussisk rubel", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "hviderussisk rubel (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "belizisk dollar", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "canadisk dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "congolesisk franc", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR euro", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "schweizerfranc", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR franc", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "chilensk escudo", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "chilensk peso", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "kinesisk yuan (offshore)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "kinesisk yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "colombiansk peso", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "costaricansk colón", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Serbisk dinar (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Tjekkoslovakisk hard koruna", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "cubansk konvertibel peso", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "cubansk peso", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "kapverdisk escudo", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Cypriotisk pund", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "tjekkisk koruna", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Østtysk mark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Tysk mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "djiboutisk franc", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "dansk krone", Symbol: "kr."},
				currency.DOP: cldr.Currency{DisplayName: "dominikansk peso", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "algerisk dinar", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadoriansk sucre", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Estisk kroon", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "egyptisk pund", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "eritreisk nakfa", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Spansk peseta (A–konto)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Spansk peseta (konvertibel konto)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Spansk peseta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "etiopisk birr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Finsk mark", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "fijiansk dollar", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "pund fra Falklandsøerne", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Fransk franc", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "britisk pund", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Georgisk kupon larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "georgisk lari", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Ghanesisk cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "ghanesisk cedi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "gibraltarisk pund", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "gambisk dalasi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "guineansk franc", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Guineansk syli", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Ækvatorialguineask ekwele", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Græsk drakme", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "guatemalansk quetzal", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portugisisk guinea escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guineansk peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "guyansk dollar", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "hongkongsk dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "honduransk lempira", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Kroatisk dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "kroatisk kuna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "haitisk gourde", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ungarsk forint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "indonesisk rupiah", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Irsk pund", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Israelsk pund", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "israelsk shekel (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "ny israelsk shekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "indisk rupee", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "irakisk dinar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "iransk rial", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "islandsk krone (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "islandsk krone", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Italiensk lire", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "jamaicansk dollar", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "jordansk dinar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "japansk yen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "kenyansk shilling", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "kirgisisk som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "cambodjansk riel", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "comorisk franc", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "nordkoreansk won", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "sydkoreansk hwan (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "sydkoreansk won (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "sydkoreansk won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "kuwaitisk dinar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "caymansk dollar", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "kasakhisk tenge", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "laotisk kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libanesisk pund", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "srilankansk rupee", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "liberisk dollar", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesothisk loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litauisk litas", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Litauisk talonas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Luxembourgsk konvertibel franc", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Luxembourgsk franc", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Luxembourgsk finansiel franc", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Lettisk lat", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Lettisk rubel", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "libysk dinar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "marokkansk dirham", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Marokkansk franc", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "monegaskisk franc", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "moldovisk cupon", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldovisk leu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "madagaskisk ariary", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Madagaskisk franc", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "makedonsk denar", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "makedonsk denar (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Malisk franc", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "myanmarsk kyat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "mongolsk tugrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "macaosk pataca", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "mauritansk ouguiya (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "mauritansk ouguiya", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Maltesisk lira", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Maltesisk pund", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "mauritisk rupee", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "maldivisk rupi (1947–1981)", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "maldivisk rufiyaa", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "malawisk kwacha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "mexicansk peso", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Mexicansk silver peso (1861–1992)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "malaysisk ringgit", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mozambiquisk escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Mozambiquisk metical (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "mozambiquisk metical", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "namibisk dollar", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "nigeriansk naira", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Nicaraguansk cordoba (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "nicaraguansk cordoba", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Hollandsk guilder", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "norsk krone", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "nepalesisk rupee", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "newzealandsk dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "omansk rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "panamansk balboa", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "peruviansk inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "peruansk sol", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "peruviansk sol (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "papuansk kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "filippinsk peso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "pakistansk rupee", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "polsk zloty", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Polsk zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portugisisk escudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "paraguaysk guarani", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "qatarsk rial", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "rhodesisk dollar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Rumænsk leu (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "rumænsk leu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "serbisk dinar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "russisk rubel", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Russisk rubel (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "rwandisk franc", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "saudiarabisk riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "salomonsk dollar", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "seychellisk rupee", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Sudansk dinar (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sudansk pund", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Sudansk pund (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "svensk krone", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "singaporeansk dollar", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "pund fra Saint Helena", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Slovensk tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Slovakisk koruna", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "sierraleonsk leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "somalisk shilling", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "surinamsk dollar", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Surinamsk guilder", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "sydsudansk pund", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "dobra fra Sao Tome og Principe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "dobra fra Sao Tome og Principe", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Sovjetisk rubel", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Salvadoransk colon", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "syrisk pund", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "swazilandsk lilangeni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "thailandsk baht", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tadsjikisk rubel", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "tadsjikisk somoni", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Turkmensk manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "turkmensk manat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "tunesisk dinar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "tongansk paʻanga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Escudo fra Timor", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Tyrkisk lire (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "tyrkisk lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "dollar fra Trinidad og Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "ny taiwansk dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "tanzanisk shilling", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ukrainsk grynia", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrainsk karbovanetz", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Ugandisk shilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "ugandisk shilling", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "amerikansk dollar", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "Amerikansk dollar (næste dag)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "Amerikansk dollar (samme dag)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Uruguayansk peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "uruguayansk peso", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "usbekisk sum", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Venezuelansk bolivar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "venezuelansk bolivar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "venezuelansk bolivar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "vietnamesisk dong", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "vietnamesisk dong (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "vanuaisk vatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "samoansk tala", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "CFA-franc (BEAC)", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Sølv", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Guld", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "EURCO", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "EMU", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "østkaribisk dollar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "SDR", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "ECU", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Fransk guldfranc", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Fransk UIC-franc", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "CFA-franc BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP-franc", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platin", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET-fond", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "testvalutakode", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "ukendt valuta", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Yemenitisk dinar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "yemenitisk rial", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Jugoslavisk hard dinar (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Jugoslavisk noviy dinar (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Jugoslavisk konvertibel dinar (1990–1992)", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "jugoslavisk reformeret dinar (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Sydafrikansk rand (financial)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "sydafrikansk rand", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambisk kwacha (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "zambisk kwacha", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Ny zairisk zaire (1993–1998)", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zairisk zaire (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwisk dollar (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabwisk dollar (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabwisk dollar (2008)", Symbol: ""},
			},
		},
	}
}
