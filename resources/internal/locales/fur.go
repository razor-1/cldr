// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_fur() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "fur",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d 'di' MMMM 'dal' y", Long: "d 'di' MMMM 'dal' y", Medium: "dd/MM/y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Zen", Feb: "Fev", Mar: "Mar", Apr: "Avr", May: "Mai", Jun: "Jug", Jul: "Lui", Aug: "Avo", Sep: "Set", Oct: "Otu", Nov: "Nov", Dec: "Dic"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Z", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "L", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Zenâr", Feb: "Fevrâr", Mar: "Març", Apr: "Avrîl", May: "Mai", Jun: "Jugn", Jul: "Lui", Aug: "Avost", Sep: "Setembar", Oct: "Otubar", Nov: "Novembar", Dec: "Dicembar"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dom", Mon: "lun", Tue: "mar", Wed: "mie", Thu: "joi", Fri: "vin", Sat: "sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "L", Tue: "M", Wed: "M", Thu: "J", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "domenie", Mon: "lunis", Tue: "martars", Wed: "miercus", Thu: "joibe", Fri: "vinars", Sat: "sabide"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.", PM: "p."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "a.", PM: "p."}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AMD: cldr.Currency{DisplayName: "Dram armen", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Peso argjentin", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "Selin austriac", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "dolar australian", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BEF: cldr.Currency{DisplayName: "Franc de Belgjiche", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Franc burundês", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Dolar dal Brunei", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "real brasilian", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Rubli bielorùs", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "Rubli bielorùs (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "dolar canadês", Symbol: "CA$"},
				currency.CHF: cldr.Currency{DisplayName: "franc svuizar", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "yuan cinês", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CSD: cldr.Currency{DisplayName: "Vieri dinar serp", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Peso cuban", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "Corone de Republiche Ceche", Symbol: "Kč"},
				currency.DEM: cldr.Currency{DisplayName: "Marc todesc", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "corone danese", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar algerin", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.FRF: cldr.Currency{DisplayName: "Franc francês", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "sterline britaniche", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "dolar di Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRD: cldr.Currency{DisplayName: "Dinar cravuat", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Kuna cravuate", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "rupiah indonesiane", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "rupie indiane", Symbol: "₹"},
				currency.IRR: cldr.Currency{DisplayName: "Rial iranian", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.ITL: cldr.Currency{DisplayName: "Lire taliane", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "yen gjaponês", Symbol: "JP¥"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "won de Coree dal Sud", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "Lats leton", Symbol: "Ls"},
				currency.MGA: cldr.Currency{DisplayName: "", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MXN: cldr.Currency{DisplayName: "peso messican", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "Dolar namibian", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Córdoba oro nicaraguan", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "corone norvegjese", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Dollar neozelandês", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Rupie pachistane", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "zloty polac", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar serp", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "rubli rus", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "riyal de Arabie Saudite", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "corone svedese", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SIT: cldr.Currency{DisplayName: "Talar sloven", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Corone slovache", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "baht tailandês", Symbol: "฿"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRL: cldr.Currency{DisplayName: "Viere Lire turche", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "lire turche", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "gnûf dolar taiwanês", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "dolar american", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "Dolar american (prossime zornade)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "Dolar american (stesse zornade)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Arint", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Aur", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Unitât composite europeane", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Unitât monetarie europeane", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Unitât di acont europeane (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Unitât di acont europeane (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Dirits speciâi di incas", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Franc aur francês", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Franc UIC francês", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Paladi", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platin", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "fonts RINET", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "codiç di verifiche de monede", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Monede no valide o no cognossude", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "rand sudafrican", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "afar",
			language.AB:      "abcazian",
			language.AE:      "avestan",
			language.AF:      "afrikaans",
			language.AM:      "amaric",
			language.AN:      "aragonês",
			language.ANG:     "vieri inglês",
			language.AR:      "arap",
			language.ARC:     "aramaic",
			language.AS:      "assamês",
			language.AST:     "asturian",
			language.AV:      "avar",
			language.AY:      "aymarà",
			language.AZ:      "azerbaijani",
			language.BE:      "bielorùs",
			language.BG:      "bulgar",
			language.BN:      "bengalês",
			language.BO:      "tibetan",
			language.BR:      "breton",
			language.BS:      "bosniac",
			language.CA:      "catalan",
			language.CE:      "cecen",
			language.CH:      "chamorro",
			language.CO:      "cors",
			language.COP:     "coptic",
			language.CR:      "cree",
			language.CS:      "cec",
			language.CU:      "sclâf de glesie",
			language.CY:      "galês",
			language.DA:      "danês",
			language.DE:      "todesc",
			language.DE_AT:   "todesc de Austrie",
			language.DE_CH:   "alt todesc de Svuizare",
			language.DEN:     "sclâf",
			language.EGY:     "vieri egjizian",
			language.EL:      "grêc",
			language.EN:      "inglês",
			language.EN_AU:   "inglês australian",
			language.EN_CA:   "inglês canadês",
			language.EN_GB:   "inglês britanic",
			language.EN_US:   "ingles merecan",
			language.EO:      "esperanto",
			language.ES:      "spagnûl",
			language.ES_419:  "spagnûl de Americhe Latine",
			language.ES_ES:   "spagnûl iberic",
			language.ET:      "eston",
			language.EU:      "basc",
			language.FA:      "persian",
			language.FF:      "fulah",
			language.FI:      "finlandês",
			language.FIL:     "filipin",
			language.FJ:      "fizian",
			language.FO:      "faroês",
			language.FR:      "francês",
			language.FR_CA:   "francês dal Canade",
			language.FR_CH:   "francês de Svuizare",
			language.FRO:     "vieri francês",
			language.FUR:     "furlan",
			language.FY:      "frisian",
			language.GA:      "gaelic irlandês",
			language.GD:      "gaelic scozês",
			language.GL:      "galizian",
			language.GN:      "guaranì",
			language.GOT:     "gotic",
			language.GRC:     "vieri grêc",
			language.GU:      "gujarati",
			language.GV:      "manx",
			language.HE:      "ebraic",
			language.HI:      "hindi",
			language.HR:      "cravuat",
			language.HT:      "haitian",
			language.HU:      "ongjarês",
			language.HY:      "armen",
			language.ID:      "indonesian",
			language.IG:      "igbo",
			language.IK:      "inupiaq",
			language.IO:      "ido",
			language.IS:      "islandês",
			language.IT:      "talian",
			language.IU:      "inuktitut",
			language.JA:      "gjaponês",
			language.KA:      "gjeorgjian",
			language.KK:      "kazac",
			language.KL:      "kalaallisut",
			language.KM:      "khmer",
			language.KN:      "kannada",
			language.KO:      "corean",
			language.KU:      "curd",
			language.KW:      "cornualiês",
			language.LA:      "latin",
			language.LAD:     "ladin",
			language.LB:      "lussemburghês",
			language.LI:      "limburghês",
			language.LN:      "lingala",
			language.LO:      "lao",
			language.LT:      "lituan",
			language.LV:      "leton",
			language.MG:      "malagasy",
			language.MI:      "maori",
			language.MK:      "macedon",
			language.ML:      "malayalam",
			language.MN:      "mongul",
			language.MR:      "marathi",
			language.MS:      "malês",
			language.MT:      "maltês",
			language.MUL:     "lenghis multiplis",
			language.MWL:     "mirandês",
			language.NAP:     "napoletan",
			language.NB:      "norvegjês bokmål",
			language.ND:      "ndebele setentrionâl",
			language.NDS:     "bas todesc",
			language.NE:      "nepalês",
			language.NL:      "olandês",
			language.NL_BE:   "flamant",
			language.NN:      "norvegjês nynorsk",
			language.NO:      "norvegjês",
			language.NON:     "vieri norvegjês",
			language.NSO:     "sotho setentrionâl",
			language.NV:      "navajo",
			language.OC:      "ocitan",
			language.OR:      "oriya",
			language.OS:      "osetic",
			language.OTA:     "turc otoman",
			language.PA:      "punjabi",
			language.PAP:     "papiamento",
			language.PEO:     "vieri persian",
			language.PL:      "polac",
			language.PRO:     "vieri provenzâl",
			language.PS:      "pashto",
			language.PT:      "portughês",
			language.PT_BR:   "portughês brasilian",
			language.PT_PT:   "portughês iberic",
			language.QU:      "quechua",
			language.RM:      "rumanç",
			language.RO:      "romen",
			language.RO_MD:   "moldâf",
			language.RU:      "rus",
			language.SA:      "sanscrit",
			language.SC:      "sardegnûl",
			language.SCN:     "sicilian",
			language.SCO:     "scozês",
			language.SD:      "sindhi",
			language.SE:      "sami setentrionâl",
			language.SG:      "sango",
			language.SGA:     "vieri irlandês",
			language.SI:      "sinalês",
			language.SK:      "slovac",
			language.SL:      "sloven",
			language.SM:      "samoan",
			language.SO:      "somal",
			language.SQ:      "albanês",
			language.SR:      "serp",
			language.SS:      "swati",
			language.ST:      "sotho meridionâl",
			language.SU:      "sundanês",
			language.SUX:     "sumeric",
			language.SV:      "svedês",
			language.SW:      "swahili",
			language.TA:      "tamil",
			language.TE:      "telegu",
			language.TET:     "tetum",
			language.TG:      "tagic",
			language.TH:      "thai",
			language.TK:      "turcmen",
			language.TL:      "tagalog",
			language.TR:      "turc",
			language.TT:      "tartar",
			language.TY:      "tahitian",
			language.UG:      "uigur",
			language.UK:      "ucrain",
			language.UND:     "indeterminade",
			language.UR:      "urdu",
			language.UZ:      "uzbec",
			language.VE:      "venda",
			language.VI:      "vietnamite",
			language.WA:      "valon",
			language.WO:      "wolof",
			language.XH:      "xhosa",
			language.YI:      "yiddish",
			language.YO:      "yoruba",
			language.ZH:      "cinês",
			language.ZH_HANS: "cinês semplificât",
			language.ZH_HANT: "cinês tradizionâl",
			language.ZU:      "zulu",
		},
		Territories: cldr.Territories{
			territory.V_001: "Mont",
			territory.V_002: "Afriche",
			territory.V_003: "Americhe dal Nord",
			territory.V_005: "Americhe meridionâl",
			territory.V_009: "Oceanie",
			territory.V_011: "Afriche ocidentâl",
			territory.V_013: "Americhe centrâl",
			territory.V_014: "Afriche orientâl",
			territory.V_015: "Afriche setentrionâl",
			territory.V_017: "Afriche di mieç",
			territory.V_018: "Afriche meridionâl",
			territory.V_019: "Americhis",
			territory.V_021: "Americhe setentrionâl",
			territory.V_029: "caraibic",
			territory.V_030: "Asie orientâl",
			territory.V_034: "Asie meridionâl",
			territory.V_035: "Asie sud orientâl",
			territory.V_039: "Europe meridionâl",
			territory.V_053: "Australie e Gnove Zelande",
			territory.V_054: "Melanesie",
			territory.V_057: "Regjon de Micronesie",
			territory.V_061: "Polinesie",
			territory.V_142: "Asie",
			territory.V_143: "Asie centrâl",
			territory.V_145: "Asie ocidentâl",
			territory.V_150: "Europe",
			territory.V_151: "Europe orientâl",
			territory.V_154: "Europe setentrionâl",
			territory.V_155: "Europe ocidentâl",
			territory.V_419: "Americhe latine",
			territory.AD:    "Andorra",
			territory.AE:    "Emirâts araps unîts",
			territory.AF:    "Afghanistan",
			territory.AG:    "Antigua e Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albanie",
			territory.AM:    "Armenie",
			territory.AO:    "Angola",
			territory.AQ:    "Antartic",
			territory.AR:    "Argjentine",
			territory.AS:    "Samoa merecanis",
			territory.AT:    "Austrie",
			territory.AU:    "Australie",
			territory.AW:    "Aruba",
			territory.AX:    "Isulis Aland",
			territory.AZ:    "Azerbaigian",
			territory.BA:    "Bosnie e Ercegovine",
			territory.BB:    "Barbados",
			territory.BD:    "Bangladesh",
			territory.BE:    "Belgjiche",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgarie",
			territory.BH:    "Bahrain",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "Sant Barthélemy",
			territory.BM:    "Bermuda",
			territory.BN:    "Brunei",
			territory.BO:    "Bolivie",
			territory.BR:    "Brasîl",
			territory.BS:    "Bahamas",
			territory.BT:    "Bhutan",
			territory.BV:    "Isule Bouvet",
			territory.BW:    "Botswana",
			territory.BY:    "Bielorussie",
			territory.BZ:    "Belize",
			territory.CA:    "Canade",
			territory.CC:    "Isulis Cocos",
			territory.CD:    "Republiche Democratiche dal Congo",
			territory.CF:    "Republiche centri africane",
			territory.CG:    "Congo - Brazzaville",
			territory.CH:    "Svuizare",
			territory.CI:    "Cueste di Avoli",
			territory.CK:    "Isulis Cook",
			territory.CL:    "Cile",
			territory.CM:    "Camerun",
			territory.CN:    "Cine",
			territory.CO:    "Colombie",
			territory.CP:    "Isule Clipperton",
			territory.CR:    "Costa Rica",
			territory.CU:    "Cuba",
			territory.CV:    "Cjâf vert",
			territory.CX:    "Isule Christmas",
			territory.CY:    "Cipri",
			territory.CZ:    "Republiche ceche",
			territory.DE:    "Gjermanie",
			territory.DG:    "Diego Garcia",
			territory.DJ:    "Gibuti",
			territory.DK:    "Danimarcje",
			territory.DM:    "Dominiche",
			territory.DO:    "Republiche dominicane",
			territory.DZ:    "Alzerie",
			territory.EA:    "Ceuta e Melilla",
			territory.EC:    "Ecuador",
			territory.EE:    "Estonie",
			territory.EG:    "Egjit",
			territory.EH:    "Sahara ocidentâl",
			territory.ER:    "Eritree",
			territory.ES:    "Spagne",
			territory.ET:    "Etiopie",
			territory.EU:    "Union europeane",
			territory.FI:    "Finlandie",
			territory.FJ:    "Fizi",
			territory.FK:    "Isulis Falkland",
			territory.FM:    "Micronesie",
			territory.FO:    "Isulis Faroe",
			territory.FR:    "France",
			territory.GA:    "Gabon",
			territory.GB:    "Ream unît",
			territory.GD:    "Grenada",
			territory.GE:    "Gjeorgjie",
			territory.GF:    "Guiana francês",
			territory.GG:    "Guernsey",
			territory.GH:    "Ghana",
			territory.GI:    "Gjibraltar",
			territory.GL:    "Groenlande",
			territory.GM:    "Gambia",
			territory.GN:    "Guinee",
			territory.GP:    "Guadalupe",
			territory.GQ:    "Guinee ecuatoriâl",
			territory.GR:    "Grecie",
			territory.GS:    "Georgia dal Sud e Isulis Sandwich dal Sud",
			territory.GT:    "Guatemala",
			territory.GU:    "Guam",
			territory.GW:    "Guinea-Bissau",
			territory.GY:    "Guyana",
			territory.HK:    "Regjon aministrative speciâl de Cine di Hong Kong",
			territory.HM:    "Isule Heard e Isulis McDonald",
			territory.HN:    "Honduras",
			territory.HR:    "Cravuazie",
			territory.HT:    "Haiti",
			territory.HU:    "Ongjarie",
			territory.IC:    "Isulis Canariis",
			territory.ID:    "Indonesie",
			territory.IE:    "Irlande",
			territory.IL:    "Israêl",
			territory.IM:    "Isule di Man",
			territory.IN:    "India",
			territory.IO:    "Teritori britanic dal Ocean Indian",
			territory.IQ:    "Iraq",
			territory.IR:    "Iran",
			territory.IS:    "Islande",
			territory.IT:    "Italie",
			territory.JE:    "Jersey",
			territory.JM:    "Gjamaiche",
			territory.JO:    "Jordanie",
			territory.JP:    "Gjapon",
			territory.KE:    "Kenya",
			territory.KG:    "Kirghizstan",
			territory.KH:    "Camboze",
			territory.KI:    "Kiribati",
			territory.KM:    "Comoris",
			territory.KN:    "San Kitts e Nevis",
			territory.KP:    "Coree dal nord",
			territory.KR:    "Coree dal sud",
			territory.KW:    "Kuwait",
			territory.KY:    "Isulis Cayman",
			territory.KZ:    "Kazachistan",
			territory.LA:    "Laos",
			territory.LB:    "Liban",
			territory.LC:    "Sante Lusie",
			territory.LI:    "Liechtenstein",
			territory.LK:    "Sri Lanka",
			territory.LR:    "Liberie",
			territory.LS:    "Lesotho",
			territory.LT:    "Lituanie",
			territory.LU:    "Lussemburc",
			territory.LV:    "Letonie",
			territory.LY:    "Libie",
			territory.MA:    "Maroc",
			territory.MC:    "Monaco",
			territory.MD:    "Moldavie",
			territory.ME:    "Montenegro",
			territory.MF:    "Sant Martin",
			territory.MG:    "Madagascar",
			territory.MH:    "Isulis Marshall",
			territory.ML:    "Mali",
			territory.MM:    "Birmanie",
			territory.MN:    "Mongolie",
			territory.MO:    "Regjon aministrative speciâl de Cine di Macao",
			territory.MP:    "Isulis Mariana dal Nord",
			territory.MQ:    "Martiniche",
			territory.MR:    "Mauritanie",
			territory.MS:    "Montserrat",
			territory.MT:    "Malta",
			territory.MU:    "Maurizi",
			territory.MV:    "Maldivis",
			territory.MW:    "Malawi",
			territory.MX:    "Messic",
			territory.MY:    "Malaysia",
			territory.MZ:    "Mozambic",
			territory.NA:    "Namibie",
			territory.NC:    "Gnove Caledonie",
			territory.NE:    "Niger",
			territory.NF:    "Isole Norfolk",
			territory.NG:    "Nigerie",
			territory.NI:    "Nicaragua",
			territory.NL:    "Paîs bas",
			territory.NO:    "Norvegje",
			territory.NP:    "Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Gnove Zelande",
			territory.OM:    "Oman",
			territory.PA:    "Panamà",
			territory.PE:    "Perù",
			territory.PF:    "Polinesie francês",
			territory.PG:    "Papue Gnove Guinee",
			territory.PH:    "Filipinis",
			territory.PK:    "Pakistan",
			territory.PL:    "Polonie",
			territory.PM:    "San Pierre e Miquelon",
			territory.PN:    "Pitcairn",
			territory.PR:    "Porto Rico",
			territory.PS:    "Teritoris palestinês",
			territory.PT:    "Portugal",
			territory.PW:    "Palau",
			territory.PY:    "Paraguay",
			territory.QA:    "Qatar",
			territory.QO:    "Oceanie periferiche",
			territory.RE:    "Reunion",
			territory.RO:    "Romanie",
			territory.RS:    "Serbie",
			territory.RU:    "Russie",
			territory.RW:    "Ruande",
			territory.SA:    "Arabie Saudide",
			territory.SB:    "Isulis Salomon",
			territory.SC:    "Seychelles",
			territory.SD:    "Sudan",
			territory.SE:    "Svezie",
			territory.SG:    "Singapore",
			territory.SH:    "Sante Eline",
			territory.SI:    "Slovenie",
			territory.SJ:    "Svalbard e Jan Mayen",
			territory.SK:    "Slovachie",
			territory.SL:    "Sierra Leone",
			territory.SM:    "San Marin",
			territory.SN:    "Senegal",
			territory.SO:    "Somalie",
			territory.SR:    "Suriname",
			territory.ST:    "Sao Tomè e Principe",
			territory.SV:    "El Salvador",
			territory.SY:    "Sirie",
			territory.SZ:    "Swaziland",
			territory.TA:    "Tristan da Cunha",
			territory.TC:    "Isulis Turks e Caicos",
			territory.TD:    "Çad",
			territory.TF:    "Teritoris meridionâi francês",
			territory.TG:    "Togo",
			territory.TH:    "Tailandie",
			territory.TJ:    "Tazikistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Timor orientâl",
			territory.TM:    "Turkmenistan",
			territory.TN:    "Tunisie",
			territory.TO:    "Tonga",
			territory.TR:    "Turchie",
			territory.TT:    "Trinidad e Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tanzanie",
			territory.UA:    "Ucraine",
			territory.UG:    "Uganda",
			territory.UM:    "Isulis periferichis minôrs dai Stâts Unîts",
			territory.US:    "Stâts Unîts",
			territory.UY:    "Uruguay",
			territory.UZ:    "Uzbechistan",
			territory.VA:    "Vatican",
			territory.VC:    "San Vincent e lis Grenadinis",
			territory.VE:    "Venezuela",
			territory.VG:    "Isulis vergjinis britanichis",
			territory.VI:    "Isulis vergjinis americanis",
			territory.VN:    "Vietnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis e Futuna",
			territory.WS:    "Samoa",
			territory.YE:    "Yemen",
			territory.YT:    "Mayotte",
			territory.ZA:    "Sud Afriche",
			territory.ZM:    "Zambia",
			territory.ZW:    "Zimbabwe",
			territory.ZZ:    "Regjon no cognossude o no valide",
		},
	}
}
