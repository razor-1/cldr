// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_ro() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ian.", Feb: "feb.", Mar: "mar.", Apr: "apr.", May: "mai", Jun: "iun.", Jul: "iul.", Aug: "aug.", Sep: "sept.", Oct: "oct.", Nov: "nov.", Dec: "dec."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "I", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "I", Jul: "I", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ianuarie", Feb: "februarie", Mar: "martie", Apr: "aprilie", May: "mai", Jun: "iunie", Jul: "iulie", Aug: "august", Sep: "septembrie", Oct: "octombrie", Nov: "noiembrie", Dec: "decembrie"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dum.", Mon: "lun.", Tue: "mar.", Wed: "mie.", Thu: "joi", Fri: "vin.", Sat: "sâm."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "J", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "du.", Mon: "lu.", Tue: "ma.", Wed: "mi.", Thu: "joi", Fri: "vi.", Sat: "sâ."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "duminică", Mon: "luni", Tue: "marți", Wed: "miercuri", Thu: "joi", Fri: "vineri", Sat: "sâmbătă"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ro]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "pesetă andorrană", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "dirham din Emiratele Arabe Unite", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "afgani afgan", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "leka albaneză", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "dram armenesc", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "gulden din Antilele Olandeze", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angoleză", Symbol: "AOA"},
				currency.ARP: cldr.Currency{DisplayName: "peso argentinian (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "peso argentinian", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "șiling austriac", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "dolar australian", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "florin aruban", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "manat azer (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "manat azer", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "dinar Bosnia-Herțegovina (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "marcă convertibilă din Bosnia și Herțegovina", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "dolar din Barbados", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "taka din Bangladesh", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "franc belgian (convertibil)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "franc belgian", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "franc belgian (financiar)", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "leva bulgărească", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "dinar din Bahrain", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "franc burundez", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "dolar din Bermuda", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "dolar din Brunei", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano bolivian", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "peso bolivian", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "mvdol bolivian", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "cruzeiro brazilian (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "real brazilian", Symbol: "BRL"},
				currency.BRR: cldr.Currency{DisplayName: "cruzeiro brazilian (1993–1994)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "dolar din Bahamas", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ngultrum din Bhutan", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "kyat birman", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "pula Botswana", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "rublă belarusă", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "rublă belarusă (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "dolar din Belize", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "dolar canadian", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "franc congolez", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "franc elvețian", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "peso chilian", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "yuan chinezesc (offshore)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "yuan chinezesc", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "peso columbian", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "colon costarican", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "dinar Serbia și Muntenegru (2002–2006)", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "peso cubanez convertibil", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "peso cubanez", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "escudo din Capul Verde", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "liră cipriotă", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "coroană cehă", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "marcă est-germană", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "marcă germană", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "franc djiboutian", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "coroană daneză", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "peso dominican", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "dinar algerian", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "sucre Ecuador", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "coroană estoniană", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "liră egipteană", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa eritreeană", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "peseta spaniolă (cont A)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "peseta spaniolă (cont convertibil)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "pesetă spaniolă", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr etiopian", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "EUR"},
				currency.FIM: cldr.Currency{DisplayName: "marcă finlandeză", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "dolar fijian", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "liră din Insulele Falkland", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "franc francez", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "liră sterlină", Symbol: "GBP"},
				currency.GEL: cldr.Currency{DisplayName: "lari georgian", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "cedi Ghana (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "cedi ghanez", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "liră din Gibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi din Gambia", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "franc guineean", Symbol: "GNF"},
				currency.GRD: cldr.Currency{DisplayName: "drahmă grecească", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "quetzal guatemalez", Symbol: "GTQ"},
				currency.GWP: cldr.Currency{DisplayName: "peso Guineea-Bissau", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "dolar guyanez", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "dolar din Hong Kong", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "lempira honduriană", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "dinar croat", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "kuna croată", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "gourde din Haiti", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "forint maghiar", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "rupie indoneziană", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "liră irlandeză", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "liră israeliană", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "șechel israelian nou", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "rupie indiană", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "dinar irakian", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "rial iranian", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "coroană islandeză", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "liră italiană", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "dolar jamaican", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "dinar iordanian", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "yen japonez", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "șiling kenyan", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "som kârgâz", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "riel cambodgian", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "franc comorian", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "won nord-coreean", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "won sud-coreean", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "dinar kuweitian", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "dolar din Insulele Cayman", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "tenge kazahă", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "kip laoțian", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "liră libaneză", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "rupie srilankeză", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "dolar liberian", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "loti lesothian", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "litu lituanian", Symbol: "LTL"},
				currency.LUC: cldr.Currency{DisplayName: "franc convertibil luxemburghez", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "franc luxemburghez", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "franc financiar luxemburghez", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "lats letonian", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "rublă Letonia", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "dinar libian", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "dirham marocan", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "franc marocan", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "leu moldovenesc", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ariary malgaș", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "franc Madagascar", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "dinar macedonean", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "franc Mali", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "kyat din Myanmar", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "tugrik mongol", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "pataca din Macao", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "ouguiya mauritană (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "ouguiya mauritană", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "liră malteză", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "rupie mauritiană", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "rufiyaa maldiviană", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha malawiană", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "peso mexican", Symbol: "MXN"},
				currency.MXP: cldr.Currency{DisplayName: "peso mexican de argint (1861–1992)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "ringgit malaiezian", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "escudo Mozambic", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "metical Mozambic vechi", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "metical mozambican", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "dolar namibian", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "naira nigeriană", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "cordoba nicaraguană (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "cordoba nicaraguană", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "gulden olandez", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "coroană norvegiană", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "rupie nepaleză", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "dolar neozeelandez", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "rial omanez", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "balboa panameză", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "inti peruvian", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "sol peruvian", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "sol peruvian (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "kina din Papua-Noua Guinee", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "peso filipinez", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "rupie pakistaneză", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "zlot polonez", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "zlot polonez (1950–1995)", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "guarani paraguayan", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "rial qatarian", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "dolar rhodesian", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "leu românesc (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "leu românesc", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "dinar sârbesc", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "rublă rusească", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "franc rwandez", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "rial saudit", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "dolar din Insulele Solomon", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "rupie din Seychelles", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "dinar sudanez", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "liră sudaneză", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "liră sudaneză (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "coroană suedeză", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "dolar singaporez", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "liră Insula Sf. Elena", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "tolar sloven", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "coroană slovacă", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "leone din Sierra Leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "șiling somalez", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "dolar surinamez", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "gulden Surinam", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "liră din Sudanul de Sud", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "dobra Sao Tome și Principe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "dobra Sao Tome și Principe", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "rublă sovietică", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "colon El Salvador", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "liră siriană", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni din Swaziland", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "baht thailandez", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "rublă Tadjikistan", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "somoni tadjic", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "manat turkmen (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "manat turkmen", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "dinar tunisian", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "pa’anga tongană", Symbol: "TOP"},
				currency.TRL: cldr.Currency{DisplayName: "liră turcească (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "liră turcească", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "dolar din Trinidad-Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "dolar nou din Taiwan", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "șiling tanzanian", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "hryvna ucraineană", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "carboavă ucraineană", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "șiling ugandez (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "șiling ugandez", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "dolar american", Symbol: "USD"},
				currency.USN: cldr.Currency{DisplayName: "dolar american (ziua următoare)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "dolar american (aceeași zi)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "peso Uruguay (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "peso uruguayan", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "sum Uzbekistan", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "bolivar Venezuela (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "bolivar venezuelean (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "bolivar venezuelean", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "dong vietnamez", Symbol: "VND"},
				currency.VUV: cldr.Currency{DisplayName: "vatu din Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "tala samoană", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "franc CFA BEAC", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "argint", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "aur", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "unitate compusă europeană", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "unitate monetară europeană", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "unitate de cont europeană (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "unitate de cont europeană (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "dolar din Caraibele de Est", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "drepturi speciale de tragere", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "unitate de monedă europeană", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "franc francez de aur", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "franc UIC francez", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "franc CFA BCEAO", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "paladiu", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "franc CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platină", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "cod monetar de test", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "monedă necunoscută", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "dinar Yemen", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "rial yemenit", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "dinar iugoslav greu", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "dinar iugoslav nou", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "dinar iugoslav convertibil", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "rand sud-african (financiar)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "rand sud-african", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "kwacha zambian (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha zambian", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "zair nou", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "dolar Zimbabwe (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "dolar Zimbabwe (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "dolar Zimbabwe (2008)", Symbol: ""},
			},
		},
	}
}
