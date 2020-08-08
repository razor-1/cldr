// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_fr_BL() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}, GMT: "UTC{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "janv.", Feb: "févr.", Mar: "mars", Apr: "avr.", May: "mai", Jun: "juin", Jul: "juil.", Aug: "août", Sep: "sept.", Oct: "oct.", Nov: "nov.", Dec: "déc."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "janvier", Feb: "février", Mar: "mars", Apr: "avril", May: "mai", Jun: "juin", Jul: "juillet", Aug: "août", Sep: "septembre", Oct: "octobre", Nov: "novembre", Dec: "décembre"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dim.", Mon: "lun.", Tue: "mar.", Wed: "mer.", Thu: "jeu.", Fri: "ven.", Sat: "sam."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "J", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "di", Mon: "lu", Tue: "ma", Wed: "me", Thu: "je", Fri: "ve", Sat: "sa"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "dimanche", Mon: "lundi", Tue: "mardi", Wed: "mercredi", Thu: "jeudi", Fri: "vendredi", Sat: "samedi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_fr]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u200f−", Percent: "٪", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "peseta andorrane", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "dirham des Émirats arabes unis", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "afghani (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "afghani afghan", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "lek albanais (1947–1961)", Symbol: "ALK"},
				currency.ALL: cldr.Currency{DisplayName: "lek albanais", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "dram arménien", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "florin antillais", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angolais", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "kwanza angolais (1977–1990)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "nouveau kwanza angolais (1990–2000)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "kwanza angolais réajusté (1995–1999)", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "austral argentin", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "peso lourd argentin (1970–1983)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "peso argentin (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "peso argentin (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "peso argentin", Symbol: "$AR"},
				currency.ATS: cldr.Currency{DisplayName: "schilling autrichien", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "dollar australien", Symbol: "$AU"},
				currency.AWG: cldr.Currency{DisplayName: "florin arubais", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "manat azéri (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "manat azéri", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "dinar bosniaque", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "mark convertible bosniaque", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "nouveau dinar bosniaque (1994–1997)", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "dollar barbadien", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "taka bangladeshi", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "franc belge (convertible)", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "franc belge", Symbol: "FB"},
				currency.BEL: cldr.Currency{DisplayName: "franc belge (financier)", Symbol: "BEL"},
				currency.BGL: cldr.Currency{DisplayName: "lev bulgare (1962–1999)", Symbol: "BGL"},
				currency.BGM: cldr.Currency{DisplayName: "lev socialiste bulgare", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "lev bulgare", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "lev bulgare (1879–1952)", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "dinar bahreïni", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "franc burundais", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "dollar bermudien", Symbol: "$BM"},
				currency.BND: cldr.Currency{DisplayName: "dollar brunéien", Symbol: "$BN"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano bolivien", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "boliviano bolivien (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "peso bolivien", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "mvdol bolivien", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "nouveau cruzeiro brésilien (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "cruzado brésilien (1986–1989)", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "cruzeiro brésilien (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "réal brésilien", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "nouveau cruzado", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "cruzeiro", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "cruzeiro brésilien (1942–1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "dollar bahaméen", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ngultrum bouthanais", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "kyat birman", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "pula botswanais", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "nouveau rouble biélorusse (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "rouble biélorusse", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "rouble biélorusse (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "dollar bélizéen", Symbol: "$BZ"},
				currency.CAD: cldr.Currency{DisplayName: "dollar canadien", Symbol: "$CA"},
				currency.CDF: cldr.Currency{DisplayName: "franc congolais", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "euro WIR", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "franc suisse", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "franc WIR", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "escudo chilien", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "unité d’investissement chilienne", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "peso chilien", Symbol: "$CL"},
				currency.CNH: cldr.Currency{DisplayName: "yuan chinois (zone extracôtière)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "dollar de la Banque populaire chinoise", Symbol: "CNX"},
				currency.CNY: cldr.Currency{DisplayName: "yuan renminbi chinois", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "peso colombien", Symbol: "$CO"},
				currency.COU: cldr.Currency{DisplayName: "unité de valeur réelle colombienne", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "colón costaricain", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "dinar serbo-monténégrin", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "couronne forte tchécoslovaque", Symbol: "CSK"},
				currency.CUC: cldr.Currency{DisplayName: "peso cubain convertible", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "peso cubain", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "escudo capverdien", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "livre chypriote", Symbol: "£CY"},
				currency.CZK: cldr.Currency{DisplayName: "couronne tchèque", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "mark est-allemand", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "mark allemand", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "franc djiboutien", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "couronne danoise", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "peso dominicain", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "dinar algérien", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "sucre équatorien", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "unité de valeur constante équatoriale (UVC)", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "couronne estonienne", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "livre égyptienne", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "nafka érythréen", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "peseta espagnole (compte A)", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "peseta espagnole (compte convertible)", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "peseta espagnole", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "birr éthiopien", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "mark finlandais", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "dollar fidjien", Symbol: "$FJ"},
				currency.FKP: cldr.Currency{DisplayName: "livre des îles Malouines", Symbol: "£FK"},
				currency.FRF: cldr.Currency{DisplayName: "franc français", Symbol: "F"},
				currency.GBP: cldr.Currency{DisplayName: "livre sterling", Symbol: "£GB"},
				currency.GEK: cldr.Currency{DisplayName: "coupon de lari géorgien", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "lari géorgien", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "cédi", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "cédi ghanéen", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "livre de Gibraltar", Symbol: "£GI"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi gambien", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "franc guinéen", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "syli guinéen", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "ekwélé équatoguinéen", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "drachme grecque", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "quetzal guatémaltèque", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "escudo de Guinée portugaise", Symbol: "GWE"},
				currency.GWP: cldr.Currency{DisplayName: "peso bissau-guinéen", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "dollar du Guyana", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "dollar de Hong Kong", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "lempira hondurien", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "dinar croate", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "kuna croate", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "gourde haïtienne", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "forint hongrois", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "roupie indonésienne", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "livre irlandaise", Symbol: "£IE"},
				currency.ILP: cldr.Currency{DisplayName: "livre israélienne", Symbol: "£IL"},
				currency.ILR: cldr.Currency{DisplayName: "shekel israélien (1980–1985)", Symbol: "ILR"},
				currency.ILS: cldr.Currency{DisplayName: "nouveau shekel israélien", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "roupie indienne", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "dinar irakien", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "riyal iranien", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "couronne islandaise (1918–1981)", Symbol: "ISJ"},
				currency.ISK: cldr.Currency{DisplayName: "couronne islandaise", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "lire italienne", Symbol: "₤IT"},
				currency.JMD: cldr.Currency{DisplayName: "dollar jamaïcain", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "dinar jordanien", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "yen japonais", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "shilling kényan", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "som kirghize", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "riel cambodgien", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "franc comorien", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "won nord-coréen", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "hwan sud-coréen (1953–1962)", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "won sud-coréen (1945–1953)", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "won sud-coréen", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "dinar koweïtien", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "dollar des îles Caïmans", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "tenge kazakh", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "kip loatien", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "livre libanaise", Symbol: "£LB"},
				currency.LKR: cldr.Currency{DisplayName: "roupie srilankaise", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "dollar libérien", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "loti lesothan", Symbol: "lLS"},
				currency.LTL: cldr.Currency{DisplayName: "litas lituanien", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "talonas lituanien", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "franc convertible luxembourgeois", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "franc luxembourgeois", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "franc financier luxembourgeois", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "lats letton", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "rouble letton", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "dinar libyen", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "dirham marocain", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "franc marocain", Symbol: "fMA"},
				currency.MCF: cldr.Currency{DisplayName: "franc monégasque", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "cupon moldave", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "leu moldave", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ariary malgache", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "franc malgache", Symbol: "Fmg"},
				currency.MKD: cldr.Currency{DisplayName: "denar macédonien", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "denar macédonien (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "franc malien", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "kyat myanmarais", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "tugrik mongol", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "pataca macanaise", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "ouguiya mauritanien (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "ouguiya mauritanien", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "lire maltaise", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "livre maltaise", Symbol: "£MT"},
				currency.MUR: cldr.Currency{DisplayName: "roupie mauricienne", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "roupie maldivienne (1947–1981)", Symbol: "MVP"},
				currency.MVR: cldr.Currency{DisplayName: "rufiyaa maldivien", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha malawite", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "peso mexicain", Symbol: "$MX"},
				currency.MXP: cldr.Currency{DisplayName: "peso d’argent mexicain (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "unité de conversion mexicaine (UDI)", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "ringgit malais", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "escudo mozambicain", Symbol: "MZE"},
				currency.MZM: cldr.Currency{DisplayName: "métical", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "metical mozambicain", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "dollar namibien", Symbol: "$NA"},
				currency.NGN: cldr.Currency{DisplayName: "naira nigérian", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "cordoba", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "córdoba oro nicaraguayen", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "florin néerlandais", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "couronne norvégienne", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "roupie népalaise", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "dollar néo-zélandais", Symbol: "$NZ"},
				currency.OMR: cldr.Currency{DisplayName: "riyal omanais", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "balboa panaméen", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "inti péruvien", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "sol péruvien", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "sol péruvien (1863–1985)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "kina papouan-néo-guinéen", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "peso philippin", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "roupie pakistanaise", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "zloty polonais", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "zloty (1950–1995)", Symbol: "PLZ"},
				currency.PTE: cldr.Currency{DisplayName: "escudo portugais", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "guaraní paraguayen", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "riyal qatari", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "dollar rhodésien", Symbol: "$RH"},
				currency.ROL: cldr.Currency{DisplayName: "ancien leu roumain", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "leu roumain", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "dinar serbe", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "rouble russe", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "rouble russe (1991–1998)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "franc rwandais", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "riyal saoudien", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "dollar des îles Salomon", Symbol: "$SB"},
				currency.SCR: cldr.Currency{DisplayName: "roupie des Seychelles", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "dinar soudanais", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "livre soudanaise", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "livre soudanaise (1956–2007)", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "couronne suédoise", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "dollar de Singapour", Symbol: "$SG"},
				currency.SHP: cldr.Currency{DisplayName: "livre de Sainte-Hélène", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "tolar slovène", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "couronne slovaque", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "leone sierra-léonais", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "shilling somalien", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "dollar surinamais", Symbol: "$SR"},
				currency.SRG: cldr.Currency{DisplayName: "florin surinamais", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "livre sud-soudanaise", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "dobra santoméen (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "dobra santoméen", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "rouble soviétique", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "colón salvadorien", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "livre syrienne", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni swazi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "baht thaïlandais", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "rouble tadjik", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "somoni tadjik", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "manat turkmène", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "nouveau manat turkmène", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "dinar tunisien", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "pa’anga tongan", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "escudo timorais", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "livre turque (1844–2005)", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "livre turque", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "dollar trinidadien", Symbol: "$TT"},
				currency.TWD: cldr.Currency{DisplayName: "nouveau dollar taïwanais", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "shilling tanzanien", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "hryvnia ukrainienne", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "karbovanetz", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "shilling ougandais (1966–1987)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "shilling ougandais", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "dollar des États-Unis", Symbol: "$US"},
				currency.USN: cldr.Currency{DisplayName: "dollar des Etats-Unis (jour suivant)", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "dollar des Etats-Unis (jour même)", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "peso uruguayen (unités indexées)", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "peso uruguayen (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "peso uruguayen", Symbol: "$UY"},
				currency.UYW: cldr.Currency{DisplayName: "unité de salaire nominal uruguayenne", Symbol: "UYW"},
				currency.UZS: cldr.Currency{DisplayName: "sum ouzbek", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "bolivar vénézuélien (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "bolivar vénézuélien (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "bolivar vénézuélien", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "dông vietnamien", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "dông vietnamien (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "vatu vanuatuan", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "tala samoan", Symbol: "$WS"},
				currency.XAF: cldr.Currency{DisplayName: "franc CFA (BEAC)", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "argent", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "or", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "unité européenne composée", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "unité monétaire européenne", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "unité de compte européenne (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "unité de compte européenne (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "dollar des Caraïbes orientales", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "droit de tirage spécial", Symbol: "DTS"},
				currency.XEU: cldr.Currency{DisplayName: "unité de compte européenne (ECU)", Symbol: "XEU"},
				currency.XFO: cldr.Currency{DisplayName: "franc or", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "franc UIC", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "franc CFA (BCEAO)", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "palladium", Symbol: "XPD"},
				currency.XPF: cldr.Currency{DisplayName: "franc CFP", Symbol: "FCFP"},
				currency.XPT: cldr.Currency{DisplayName: "platine", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "type de fonds RINET", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "sucre", Symbol: "XSU"},
				currency.XTS: cldr.Currency{DisplayName: "(devise de test)", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "unité de compte ADB", Symbol: "XUA"},
				currency.XXX: cldr.Currency{DisplayName: "devise inconnue ou non valide", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "dinar du Yémen", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "riyal yéménite", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "nouveau dinar yougoslave", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "dinar yougoslave Noviy", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "dinar yougoslave convertible", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "dinar réformé yougoslave (1992–1993)", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "rand sud-africain (financier)", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "rand sud-africain", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "kwacha zambien (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha zambien", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "nouveau zaïre zaïrien", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "zaïre zaïrois", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "dollar zimbabwéen", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "dollar zimbabwéen (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "dollar zimbabwéen (2008)", Symbol: "ZWR"},
			},
		},
	}
}
