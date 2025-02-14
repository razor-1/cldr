// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ie() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ie",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d.M.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "TMG{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "febr.", Mar: "mar.", Apr: "apr.", May: "may", Jun: "jun.", Jul: "julí", Aug: "aug.", Sep: "sept.", Oct: "oct.", Nov: "nov.", Dec: "dec."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "januar", Feb: "februar", Mar: "marte", Apr: "april", May: "may", Jun: "junio", Jul: "julí", Aug: "august", Sep: "septembre", Oct: "octobre", Nov: "novembre", Dec: "decembre"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sol.", Mon: "lun.", Tue: "mar.", Wed: "mer.", Thu: "jov.", Fri: "ven.", Sat: "sat."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "L", Tue: "M", Wed: "M", Thu: "J", Fri: "V", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "So", Mon: "Lu", Tue: "Ma", Wed: "Me", Thu: "Jo", Fri: "Ve", Sat: "Sa"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "soledí", Mon: "lunedí", Tue: "mardí", Wed: "mercurdí", Thu: "jovedí", Fri: "venerdí", Sat: "saturdí"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a.", PM: "p."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ante midí", PM: "pos midí"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00;¤\u00a0-#,##0.00", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "UAE dirham", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afghan afghani", Symbol: "؋"},
				currency.ALL: cldr.Currency{DisplayName: "albanian lek", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "armenian dram", Symbol: "֏"},
				currency.ANG: cldr.Currency{DisplayName: "guilder del Nederlandesi Antilles", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "angolan kwanza", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "argentinian peso", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "australian dollar", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "aruban florin", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "azerbaidjanesi manat", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "convertibil mark de Bosnia-Herzogovina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "barbadan dollar", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "bangladeshan taka", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "bulgarian lev", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "bahrainesi dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "burundian franc", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "bermudan dollar", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "brunesi dollar", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "bolivian boliviano", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "brasilian real", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "bahaman dollar", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "bhutanesi ngultrum", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "botswanan pula", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "bieloruss ruble", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "belizan dollar", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "canadan dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "congolesi franc", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "sviss franc", Symbol: "F.Sv."},
				currency.CLP: cldr.Currency{DisplayName: "chilan peso", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "chinesi yuan (extraterritorial)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "chinesi yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "columbian peso", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "costa-rican colon", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "cuban convertibil peso", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "cuban peso", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "cap-verdesi escudo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "tchec koruna", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "djibutian franc", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "danesi krone", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "dominican peso", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "algerian dinar", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "egiptian pund", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "eritrean nakfa", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "etiopian birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "fidjian dollar", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "pund de Falkland", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "pund sterling", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "georgian lari", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "ghanan cedi", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "pund de Gibraltar", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "gambian dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "guinean franc", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "guatemalan quetzal", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "guyanesi dollar", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "dollar de Hong-Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "honduresi lempira", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "croatian kuna", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "haitian gourde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "hungarian forint", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "indonesian rupia", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "israelesi nov shekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "indian rupia", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "irakesi dinar", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "iranesi real", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "islandesi krona", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "jamaican dollar", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "jordanian dinar", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "japanesi yen", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "kenian shilling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "kirgistanesi som", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "cambodjan riel", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "comoran franc", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "nord-korean won", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "sud-korean won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "kuwaitesi dinar", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "caymanesi dollar", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "kazakhstanesi tenge", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "laotic kip", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "libanesi pund", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "sri-lankan rupia", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "liberian dollar", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "lesothan loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "libian dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "maroccan dirham", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "moldovan lei", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "malgachic ariary", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "macedonian denar", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "myanmaran kyat", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "mongolian tugric", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "macan pataca", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "mauritanian uguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "maurician rupia", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "maldivan rufia", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "malawian kwacha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "mexican peso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "malaysian ringgit", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "mozambican metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "namibian dollar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "nigerian naira", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "nicaraguan cordoba", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "norvegian krone", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "nepalesi rupia", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "nov-zelandesi dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "omanesi rial", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "panamesi balboa", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "peruan sol", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "kina de Papua Nov-Guinea", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "filipinesi peso", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "pakistani rupia", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "polonesi zloty", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "paraguayan guarani", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "qataresi riyal", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "rumanian lei", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "serbian dinar", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "russian ruble", Symbol: "Rub."},
				currency.RWF: cldr.Currency{DisplayName: "rwandan franc", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "sauditic riyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "dollar del Insules Solomon", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "seychellan rupia", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sudanesi pund", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "sved krona", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "singapuran dollar", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "pund de Sant-Helena", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "leone de Sierra-Leone", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "leone de Sierra-Leone (1964..2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "somalian shilling", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "surinamesi dollar", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "sud-sudanesi pund", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "dobra de São Tomé e Príncipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "sirian pund", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "swazian lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "thai baht", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "tadjikistanesi somoni", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "turcmen manat", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "tunisian dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "tongan pa’anga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "turc lira", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "dollar de Trinidad e Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "nov taiwanesi dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "tanzanian shilling", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "ukrainan hrivnia", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "ugandesi shilling", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "american dollar", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "uruguayan peso", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "uzbec som", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "venezuelan bolivar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "vietnamesi dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "vanuatuan vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "samoan tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "CFA franc", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "argent", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "aur", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "ost-caribean dollar", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "west-african CFA franc", Symbol: "F\u202fCFA"},
				currency.XPD: cldr.Currency{DisplayName: "palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP franc", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platine", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "ínconosset valuta", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "yemenesi real", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "sud-african rand", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "zambian kwacha", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AB:      "abkhazian",
			language.ADY:     "adyghean",
			language.AF:      "afrikaans",
			language.ALE:     "aleut",
			language.ALT:     "sud-altaic",
			language.AM:      "amharesi",
			language.AN:      "aragonesi",
			language.AR:      "arabic",
			language.ARN:     "mapuche",
			language.ARS:     "arabic najdi",
			language.AS:      "assamesi",
			language.AST:     "asturian",
			language.AV:      "avar",
			language.AY:      "aymaran",
			language.AZ:      "azerbaidjanesi",
			language.BA:      "bashkir",
			language.BAN:     "balinesi",
			language.BE:      "bieloruss",
			language.BG:      "bulgarian",
			language.BI:      "Bislama",
			language.BM:      "bambaran",
			language.BN:      "bangla",
			language.BO:      "tibetan",
			language.BR:      "bretonesi",
			language.BRX:     "bodo",
			language.BS:      "bosnian",
			language.CA:      "catalan",
			language.CE:      "tchetchen",
			language.CEB:     "cebuano",
			language.CH:      "chamorro",
			language.CHR:     "cherokee",
			language.CHY:     "cheyenne",
			language.CKB:     "kurd, sorani",
			language.CO:      "corsican",
			language.CRJ:     "sudost-cree",
			language.CRK:     "planuran cree",
			language.CRL:     "nordost-cree",
			language.CRM:     "cree de Moose",
			language.CRR:     "algonquinesi de Carolina",
			language.CS:      "tchec",
			language.CSW:     "paludan cree",
			language.CV:      "chuvash",
			language.CY:      "vallesi",
			language.DA:      "danesi",
			language.DAK:     "dakota",
			language.DE:      "german",
			language.DE_AT:   "austrian german",
			language.DE_CH:   "sviss alt-german",
			language.DSB:     "platt sorbic",
			language.DZ:      "dzongkha",
			language.EL:      "grec",
			language.EN:      "anglesi",
			language.EN_AU:   "australian anglesi",
			language.EN_CA:   "canadian anglesi",
			language.EN_GB:   "anglesi de UR",
			language.EN_US:   "anglesi de USA",
			language.EO:      "Esperanto",
			language.ES:      "hispan",
			language.ES_419:  "hispan del latin America",
			language.ES_ES:   "europan hispan",
			language.ES_MX:   "mexican hispan",
			language.ET:      "estonian",
			language.EU:      "basc",
			language.FA:      "persian",
			language.FA_AF:   "dari",
			language.FI:      "finn",
			language.FIL:     "filipinesi",
			language.FJ:      "fidjian",
			language.FO:      "feroesi",
			language.FR:      "francesi",
			language.FR_CA:   "canadian francesi",
			language.FR_CH:   "sviss francesi",
			language.FRC:     "cadjun-francesi",
			language.FRR:     "nord-frisian",
			language.FY:      "west-frisian",
			language.GA:      "irlandesi",
			language.GD:      "scotian gaelic",
			language.GL:      "galician",
			language.GSW:     "swiss-aleman",
			language.HA:      "hausa",
			language.HAI:     "haidan",
			language.HAW:     "hawaian",
			language.HAX:     "sud-haidan",
			language.HE:      "hebreic",
			language.HI:      "hindi",
			language.HI_LATN: "hinglish",
			language.HR:      "croatian",
			language.HSB:     "montan sorbic",
			language.HT:      "haitian creol",
			language.HU:      "hungarian",
			language.HY:      "armenian",
			language.IA:      "Interlingua",
			language.ID:      "indonesian",
			language.IE:      "Interlingue",
			language.II:      "yi de Sichuan",
			language.INH:     "ingush",
			language.IO:      "Ido",
			language.IS:      "islandesi",
			language.IT:      "italian",
			language.JA:      "japanesi",
			language.JBO:     "Lojban",
			language.JV:      "javan",
			language.KA:      "georgian",
			language.KK:      "kazakh",
			language.KO:      "korean",
			language.KRL:     "karelian",
			language.KS:      "kashmiran",
			language.KSH:     "kölnesi",
			language.KU:      "kurdesi",
			language.KW:      "cornwallesi",
			language.LA:      "latin",
			language.LAD:     "ladino",
			language.LB:      "luxemburgic",
			language.LIJ:     "ligurian",
			language.LKT:     "lakota",
			language.LOU:     "creol de Louisiana",
			language.LRC:     "nord-lurian",
			language.LT:      "lituan",
			language.LUO:     "luo",
			language.LV:      "lettonian",
			language.MAD:     "maduran",
			language.MAK:     "macassaresi",
			language.MDF:     "mokshan",
			language.MG:      "malgachic",
			language.MH:      "marshallesi",
			language.MI:      "maoric",
			language.MK:      "macedonian",
			language.ML:      "malayalam",
			language.MN:      "mongolian",
			language.MS:      "malayesi",
			language.MT:      "maltesi",
			language.MUL:     "multiplic lingues",
			language.MWL:     "mirandesi",
			language.MY:      "birman",
			language.NA:      "nauru",
			language.NAP:     "neapolitan",
			language.NB:      "norvegian, bokmål",
			language.ND:      "nord-ndebele",
			language.NDS:     "bass-german",
			language.NE:      "nepalesi",
			language.NL:      "hollandesi",
			language.NL_BE:   "flandrian",
			language.NN:      "neo-norvegian",
			language.NO:      "norvegian",
			language.NR:      "sud-ndebele",
			language.NSO:     "nord-sotho",
			language.NV:      "navajo",
			language.OC:      "occitan",
			language.OJC:     "central odjibwe",
			language.OJS:     "odji-cree",
			language.OS:      "ossetian",
			language.PA:      "pandjabic",
			language.PAP:     "Papiamento",
			language.PAU:     "palauan",
			language.PCM:     "pidgin-nigerian",
			language.PL:      "polonesi",
			language.PRG:     "prussian",
			language.PS:      "pashto",
			language.PT:      "portugalesi",
			language.PT_BR:   "brasilian portugalesi",
			language.PT_PT:   "europan portugalesi",
			language.QU:      "quechua",
			language.RAP:     "rapanuic",
			language.RM:      "reto-romanch",
			language.RO:      "rumanian",
			language.RU:      "russ",
			language.SA:      "sanscrit",
			language.SAH:     "yakutan",
			language.SAT:     "santalesi",
			language.SC:      "sardinian",
			language.SCN:     "sicilian",
			language.SCO:     "Scots",
			language.SE:      "nord-samian",
			language.SI:      "sinhalesi",
			language.SK:      "slovac",
			language.SL:      "slovenian",
			language.SM:      "samoan",
			language.SN:      "shonan",
			language.SO:      "somalian",
			language.SQ:      "albanesi",
			language.SR:      "serbian",
			language.SRN:     "sranan",
			language.SS:      "swati",
			language.ST:      "sud-sotho",
			language.STR:     "salishan del Straits",
			language.SU:      "sundan",
			language.SV:      "sved",
			language.SW:      "swahili",
			language.SWB:     "maoresi comoran",
			language.SYR:     "sirian",
			language.TG:      "tadjic",
			language.TH:      "thai",
			language.TK:      "turcmen",
			language.TLH:     "Klingon",
			language.TLI:     "tlingit",
			language.TOK:     "Toki Pona",
			language.TR:      "turc",
			language.TT:      "tataric",
			language.TY:      "tahitian",
			language.TYV:     "tuvan",
			language.TZM:     "tamazight del central Atlas",
			language.UDM:     "udmurtan",
			language.UG:      "uyghur",
			language.UK:      "ukrainan",
			language.UND:     "ínconosset lingue",
			language.UR:      "urdu",
			language.UZ:      "uzbec",
			language.VE:      "venda",
			language.VI:      "vietnamesi",
			language.WA:      "wallonesi",
			language.XH:      "xhosa",
			language.YI:      "yiddish",
			language.YO:      "yoruba",
			language.YUE:     "chinesi, cantonesi",
			language.ZGH:     "standard maroccan tamazight",
			language.ZH:      "chinesi, mandarin",
			language.ZH_HANS: "chinesi simplificat",
			language.ZH_HANT: "chinesi traditional",
			language.ZU:      "zulu",
			language.ZXX:     "sin linguistic contenete",
		},
		Territories: cldr.Territories{
			territory.V_001: "munde",
			territory.V_002: "Africa",
			territory.V_003: "septentrional America",
			territory.V_005: "Sud-America",
			territory.V_009: "Oceania",
			territory.V_011: "West-Africa",
			territory.V_013: "central America",
			territory.V_014: "Ost-Africa",
			territory.V_015: "Nord-Africa",
			territory.V_017: "central Africa",
			territory.V_018: "meridional Africa",
			territory.V_019: "Americas",
			territory.V_021: "Nord-America",
			territory.V_029: "Caribes",
			territory.V_030: "Ost-Asia",
			territory.V_034: "Sud-Asia",
			territory.V_035: "Sudost-Asia",
			territory.V_039: "Sud-Europa",
			territory.V_053: "Australasia",
			territory.V_054: "Melanesia",
			territory.V_057: "region de Micronesia",
			territory.V_061: "Polinesia",
			territory.V_142: "Asia",
			territory.V_143: "Central Asia",
			territory.V_145: "West-Asia",
			territory.V_150: "Europa",
			territory.V_151: "Ost-Europa",
			territory.V_154: "Nord-Europa",
			territory.V_155: "West-Europa",
			territory.V_202: "Sub-Saharan Africa",
			territory.V_419: "latin America",
			territory.AC:    "Insul de Ascension",
			territory.AD:    "Andorra",
			territory.AE:    "Unit Arab Emiratus",
			territory.AF:    "Afghanistan",
			territory.AG:    "Antigua e Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albania",
			territory.AM:    "Armenia",
			territory.AO:    "Angola",
			territory.AQ:    "Antarctica",
			territory.AR:    "Argentinia",
			territory.AS:    "American Samoa",
			territory.AT:    "Austria",
			territory.AU:    "Australia",
			territory.AW:    "Aruba",
			territory.AX:    "Insules Åland",
			territory.AZ:    "Azerbaidjan",
			territory.BA:    "Bosnia e Herzegovina",
			territory.BB:    "Barbados",
			territory.BD:    "Bangladesh",
			territory.BE:    "Belgia",
			territory.BF:    "Burkina-Faso",
			territory.BG:    "Bulgaria",
			territory.BH:    "Bahrain",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "Sant-Bartolomeo",
			territory.BM:    "Bermuda",
			territory.BN:    "Brunei",
			territory.BO:    "Bolivia",
			territory.BQ:    "Caribean Nederland",
			territory.BR:    "Brasilia",
			territory.BS:    "Bahamas",
			territory.BT:    "Bhutan",
			territory.BV:    "Insul Bouvet",
			territory.BW:    "Botswana",
			territory.BY:    "Bielorussia",
			territory.BZ:    "Belize",
			territory.CA:    "Canada",
			territory.CC:    "Insules Cocos (Keeling)",
			territory.CD:    "Congo (Kinshasa)",
			territory.CF:    "Central African Republica",
			territory.CG:    "Congo (Brazzaville)",
			territory.CH:    "Svissia",
			territory.CI:    "Coste de Ivor",
			territory.CK:    "Insules Cook",
			territory.CL:    "Chile",
			territory.CM:    "Camerun",
			territory.CN:    "China",
			territory.CO:    "Columbia",
			territory.CP:    "Insul Clipperton",
			territory.CR:    "Costa-Rica",
			territory.CU:    "Cuba",
			territory.CV:    "Cap-Verdi",
			territory.CW:    "Curaçao",
			territory.CX:    "Insul Christmas",
			territory.CY:    "Cypria",
			territory.CZ:    "Tchekia",
			territory.DE:    "Germania",
			territory.DG:    "Diego-Garcia",
			territory.DJ:    "Djibuti",
			territory.DK:    "Dania",
			territory.DM:    "Dominica",
			territory.DO:    "Dominican Republica",
			territory.DZ:    "Algeria",
			territory.EA:    "Ceuta e Melilla",
			territory.EC:    "Ecuador",
			territory.EE:    "Estonia",
			territory.EG:    "Egiptia",
			territory.EH:    "West-Sahara",
			territory.ER:    "Eritrea",
			territory.ES:    "Hispania",
			territory.ET:    "Etiopia",
			territory.EU:    "Union Europan",
			territory.EZ:    "Zone de euro",
			territory.FI:    "Finland",
			territory.FJ:    "Fidji",
			territory.FK:    "Insules Falkland",
			territory.FM:    "Micronesia",
			territory.FO:    "Insulas Feroe",
			territory.FR:    "Francia",
			territory.GA:    "Gabon",
			territory.GB:    "Unit Reyia",
			territory.GD:    "Grenada",
			territory.GE:    "Georgia",
			territory.GF:    "Francesi Guiana",
			territory.GG:    "Guernsey",
			territory.GH:    "Ghana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Greenland",
			territory.GM:    "Gambia",
			territory.GN:    "Guinea",
			territory.GP:    "Guadelup",
			territory.GQ:    "Equatoral Guinea",
			territory.GR:    "Grecia",
			territory.GS:    "Insules Sud-Georgia e Sud-Sandwich",
			territory.GT:    "Guatemala",
			territory.GU:    "Guam",
			territory.GW:    "Guinea-Bissau",
			territory.GY:    "Guyana",
			territory.HK:    "Hong-Kong (SAR de China)",
			territory.HM:    "Insules Heard e McDonald Islands",
			territory.HN:    "Hondura",
			territory.HR:    "Croatia",
			territory.HT:    "Haiti",
			territory.HU:    "Hungaria",
			territory.IC:    "Insules Canarias",
			territory.ID:    "Indonesia",
			territory.IE:    "Irland",
			territory.IL:    "Israel",
			territory.IM:    "Insul de Man",
			territory.IN:    "India",
			territory.IO:    "Chagos (BTIO)",
			territory.IQ:    "Irak",
			territory.IR:    "Iran",
			territory.IS:    "Island",
			territory.IT:    "Italia",
			territory.JE:    "Jersey",
			territory.JM:    "Jamaica",
			territory.JO:    "Jordania",
			territory.JP:    "Japan",
			territory.KE:    "Kenia",
			territory.KG:    "Kirgizstan",
			territory.KH:    "Cambodja",
			territory.KI:    "Kiribati",
			territory.KM:    "Comoros",
			territory.KN:    "St. Kitts e Nevis",
			territory.KP:    "Nord-Korea",
			territory.KR:    "Sud-Korea",
			territory.KW:    "Kuwait",
			territory.KY:    "Insules Cayman",
			territory.KZ:    "Kazakhstan",
			territory.LA:    "Laos",
			territory.LB:    "Liban",
			territory.LC:    "St.-Lucia",
			territory.LI:    "Liechtenstein",
			territory.LK:    "Sri-Lanka",
			territory.LR:    "Liberia",
			territory.LS:    "Lesotho",
			territory.LT:    "Lituania",
			territory.LU:    "Luxemburg",
			territory.LV:    "Lettonia",
			territory.LY:    "Libia",
			territory.MA:    "Marocco",
			territory.MC:    "Mónaco",
			territory.MD:    "Moldova",
			territory.ME:    "Montenegro",
			territory.MF:    "St.-Martin",
			territory.MG:    "Madagascar",
			territory.MH:    "Insules Marshall",
			territory.MK:    "Nord-Macedonia",
			territory.ML:    "Mali",
			territory.MM:    "Myanmar (Birmania)",
			territory.MN:    "Mongolia",
			territory.MO:    "Macao (SAR de China)",
			territory.MP:    "Insules Nord Mariana",
			territory.MQ:    "Martinica",
			territory.MR:    "Mauretania",
			territory.MS:    "Montserrat",
			territory.MT:    "Malta",
			territory.MU:    "Mauricio",
			territory.MV:    "Maldivas",
			territory.MW:    "Malawi",
			territory.MX:    "Mexico",
			territory.MY:    "Malaysia",
			territory.MZ:    "Mozambic",
			territory.NA:    "Namibia",
			territory.NC:    "Nov-Caledonia",
			territory.NE:    "Niger",
			territory.NF:    "Insul Norfolk",
			territory.NG:    "Nigeria",
			territory.NI:    "Nicaragua",
			territory.NL:    "Nederland",
			territory.NO:    "Norvegia",
			territory.NP:    "Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Nov-Zeland",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Perú",
			territory.PF:    "Francesi Polinesia",
			territory.PG:    "Papua Nov-Guinea",
			territory.PH:    "Filipines",
			territory.PK:    "Pakistan",
			territory.PL:    "Polonia",
			territory.PM:    "St.-Pierre e Miquelon",
			territory.PN:    "Insules Pitcairn",
			territory.PR:    "Porto-Rico",
			territory.PS:    "Territorias de Palestina",
			territory.PT:    "Portugal",
			territory.PW:    "Palau",
			territory.PY:    "Paraguay",
			territory.QA:    "Qatar",
			territory.QO:    "periferic Oceania",
			territory.RE:    "Reunion",
			territory.RO:    "Rumania",
			territory.RS:    "Serbia",
			territory.RU:    "Russia",
			territory.RW:    "Rwanda",
			territory.SA:    "Sauditic Arabia",
			territory.SB:    "Insules Solomon",
			territory.SC:    "Seychelles",
			territory.SD:    "Sudan",
			territory.SE:    "Svedia",
			territory.SG:    "Singapur",
			territory.SH:    "Sant-Helena",
			territory.SI:    "Slovenia",
			territory.SJ:    "Svalbard e Jan Mayen",
			territory.SK:    "Slovakia",
			territory.SL:    "Sierra-Leone",
			territory.SM:    "San-Marino",
			territory.SN:    "Senegal",
			territory.SO:    "Somalia",
			territory.SR:    "Surinam",
			territory.SS:    "Sud-Sudan",
			territory.ST:    "São-Tomé e Príncipe",
			territory.SV:    "El-Salvador",
			territory.SX:    "Sint-Maarten",
			territory.SY:    "Siria",
			territory.SZ:    "Eswatini",
			territory.TA:    "Tristan-da-Cunha",
			territory.TC:    "Turks e Caicos",
			territory.TD:    "Tchad",
			territory.TF:    "Territorias meridional de Francia",
			territory.TG:    "Togo",
			territory.TH:    "Thailand",
			territory.TJ:    "Tadjikistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Ost-Timor",
			territory.TM:    "Turkmenistan",
			territory.TN:    "Tunisia",
			territory.TO:    "Tonga",
			territory.TR:    "Turcia",
			territory.TT:    "Trinidad e Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tanzania",
			territory.UA:    "Ukraina",
			territory.UG:    "Uganda",
			territory.UM:    "Insules periferic de USA",
			territory.UN:    "Unit Nationes",
			territory.US:    "Unit States",
			territory.UY:    "Uruguay",
			territory.UZ:    "Uzbekistan",
			territory.VA:    "Cité de Vatican",
			territory.VC:    "St. Vincent e Grenadines",
			territory.VE:    "Venezuela",
			territory.VG:    "Insules Vírginas (UR)",
			territory.VI:    "Insules Vírginas (USA)",
			territory.VN:    "Vietnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis e Futuna",
			territory.WS:    "Samoa",
			territory.XA:    "pseudo-diacritica",
			territory.XB:    "pseudo-bidi",
			territory.XK:    "Kosovo",
			territory.YE:    "Yemen",
			territory.YT:    "Mayotte",
			territory.ZA:    "Sud-Africa",
			territory.ZM:    "Zambia",
			territory.ZW:    "Zimbabwe",
			territory.ZZ:    "ínconosset region",
		},
	}
}
