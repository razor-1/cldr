// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_mi() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "mi",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-y"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss\u202fa zzzz", Long: "h:mm:ss\u202fa z", Medium: "h:mm:ss\u202fa", Short: "h:mm\u202fa"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Hān", Feb: "Pēp", Mar: "Māe", Apr: "Āpe", May: "Mei", Jun: "Hun", Jul: "Hūr", Aug: "Āku", Sep: "Hep", Oct: "Oke", Nov: "Noe", Dec: "Tīh"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "H", Feb: "P", Mar: "M", Apr: "Ā", May: "M", Jun: "H", Jul: "H", Aug: "Ā", Sep: "H", Oct: "O", Nov: "N", Dec: "T"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Hānuere", Feb: "Pēpuere", Mar: "Māehe", Apr: "Āperira", May: "Mei", Jun: "Hune", Jul: "Hūrae", Aug: "Ākuhata", Sep: "Hepetema", Oct: "Oketopa", Nov: "Noema", Dec: "Tīhema"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Rāt", Mon: "Man", Tue: "Tūr", Wed: "Wen", Thu: "Tāi", Fri: "Par", Sat: "Rāh"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Rt", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "P", Sat: "Rh"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Rāt", Mon: "Man", Tue: "Tū", Wed: "Wen", Thu: "Tāi", Fri: "Par", Sat: "Rāh"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Rātapu", Mon: "Mane", Tue: "Tūrei", Wed: "Wenerei", Thu: "Tāite", Fri: "Paraire", Sat: "Rāhoroi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham UAE", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghani Awhekenetāna", Symbol: "؋"},
				currency.ALL: cldr.Currency{DisplayName: "Lek Arapeinia", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Dram Āmenia", Symbol: "֏"},
				currency.ANG: cldr.Currency{DisplayName: "Guilder Anatiri Hōrana", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza Anakora", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Peso Āketina", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Tāra Ahitereiria", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Florin Arūpa", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Manat Atepaihānia", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "Mark Pōngia-Herekōwini takahuri", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Tāra Papatohe", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Pākaratēhi", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Leva Purukāria", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dinar Pāreina", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Franc Puruniti", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Tāra Pāmura", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Tāra Poronai", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano Poriwia", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Real Parahi", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Tāra Pahama", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum Pūtana", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Pula Poriwana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Ruble Pērara", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Tāra Pērihi", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Tāra Kānata", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Franc Kōngo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Franc Huiterangi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Peso Hiri", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan Haina (ki waho)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Haina", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso Koromōpia", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Colon Koto Rika", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Peso Kiupa takahuri", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Peso Kiupa", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Kūrae Matomato", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Koruna Tieke", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Franc Tepūti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Kroner Tenemāka", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Tominika", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar Aratiria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Pāuna Īhipa", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa Eriterea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr Etiopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Tāra Whītī", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Pāuna Whākana", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pāuna Piritene", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lari Hōria", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "Cedi Kāna", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "Pāuna Kāmaka", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi Kamopia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Franc Kini", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal Kuatamāra", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Tāra Kaiana", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Tāra Hongipua", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira Honotura", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna Koroātia", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde Haiti", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Forint Hanekari", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Rupiah Initonīhia", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Shekel Hou Iharaira", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupī Iniana", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar Irāka", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Rial Irāna", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Kronur Tiorangi", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Tāra Hemeika", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar Hōrano", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yen Hapanihi", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Hereni Kenia", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Som Kikitānga", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "Riel Kamapōtia", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Franc Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Won Kōrea ki te Raki", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Won Kōrea ki te Tonga", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar Kūweiti", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Tāra Kāmana", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge Katatānga", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "kip Rāoho", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Pāuna Repanona", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Rupee Hiri Ranaka", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Tāra Raipiria", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti Teroto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinar Ripia", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirham Moroko", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Leu Morotawa", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariary Matakāhika", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Denar Makerōnia", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kyat Pēma", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "tugrik Mongōria", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca Makau", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya Mauritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupee Marihi", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa Māratiri", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha Marāwi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Peso Mēhiko", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit Mareia", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Metical Mohapiki", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Tāra Namipia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira Ngāitīria", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Cordoba Nikarāhua", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Kroner Nōwei", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Rupee Nepōra", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Tāra o Aotearoa", Symbol: "$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial Ōmana", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Balboa Panama", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sole Peru", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina Papua Nūkini", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Peso Piripīni", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Rupee Pakitāne", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty Pōrana", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani Parakai", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Riyal Katā", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Leu Romeinia", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar Hirupia", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Rūpera Ruhiana", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "Franc Rāwana", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal Hauri", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Tāra Moutere Horomona", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupee Heikere", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Pāuna Hūtāne", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Kronor Huitene", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Tāra Hingapoa", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Pāuna Hato Herena", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Leone Araone", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leone Araon (1964—2022)e", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Hereni Hūmārie", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Tāra Huriname", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Pāuna Hūtāne Tonga", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "Dobra Hao Tome me Pirinihipi", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Pāuna Hiria", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni Warerangi", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Baht Tairanga", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni Takiritānga", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Manat Tukumanatānga", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Dinar Tūnihia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Pa’anga Tonga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Lira Tākei", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Tāra Tirinaki Tōpako", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Tāra Taiwana Hou", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Hereni Tānahia", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia Ukareinga", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "hereni Ukānga", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Tāra US", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Peso Urukoi", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Som Uhipeketāne", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Bolivar Penehūera", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Dong Whitināmu", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu Whenuatū", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tala Hāmoa", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Franc CFA Āwherika Waenga", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Tāra Karapīana Rāwhiti", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Franc CFA Āwherika ki te Uru", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "Franc CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Moni Tē Mōhiotia", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Rial Īmene", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Rand Āwherika ki te Tonga", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha Tāmipia", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AB:      "Apakāhiana",
			language.ACE:     "Akanīhi",
			language.ADA:     "Atāngami",
			language.ADY:     "Āteke",
			language.AF:      "Awherikāna",
			language.AGQ:     "Ākeme",
			language.AIN:     "Ainu",
			language.AK:      "Ākana",
			language.ALE:     "Ariuta",
			language.ALT:     "Ātai ki te Tonga",
			language.AM:      "Amahereka",
			language.AN:      "Arakonihi",
			language.ANN:     "Ōporo",
			language.ANP:     "Anahika",
			language.AR:      "Ārapi",
			language.AR_001:  "Ārapi Moroki",
			language.ARN:     "Mapūte",
			language.ARP:     "Arapaho",
			language.ARS:     "Arapika Nahāri",
			language.AS:      "Āhamēhi",
			language.ASA:     "Ahu",
			language.AST:     "Ahitūriana",
			language.ATJ:     "Atikameke",
			language.AV:      "Āwhāriki",
			language.AWA:     "Awāti",
			language.AY:      "Aimāra",
			language.AZ:      "Ahēri",
			language.BA:      "Pākīra",
			language.BAN:     "Pārinīhi",
			language.BAS:     "Pahā",
			language.BE:      "Perarūhiana",
			language.BEM:     "Pema",
			language.BEZ:     "Pena",
			language.BG:      "Purukāriana",
			language.BGC:     "Herianawhi",
			language.BHO:     "Pōhipuri",
			language.BI:      "Pihirāma",
			language.BIN:     "Pini",
			language.BLA:     "Hihika",
			language.BM:      "Pāpara",
			language.BN:      "Pākara",
			language.BO:      "Tipete",
			language.BR:      "Peretana",
			language.BRX:     "Pōto",
			language.BS:      "Pōngiana",
			language.BUG:     "Pukenīhi",
			language.BYN:     "Pirina",
			language.CA:      "Katarana",
			language.CAY:     "Keiūka",
			language.CCP:     "Tiakamā",
			language.CE:      "Tietiene",
			language.CEB:     "Hepuano",
			language.CGG:     "Tieka",
			language.CH:      "Tiamoro",
			language.CHK:     "Tiukīhi",
			language.CHM:     "Mari",
			language.CHO:     "Tiokatō",
			language.CHP:     "Tiepewaiana",
			language.CHR:     "Tierokī",
			language.CHY:     "Haiene",
			language.CKB:     "Kūrihi, Hōrani",
			language.CLC:     "Tiekautini",
			language.CO:      "Kōhikana",
			language.CRG:     "Mītiwhi",
			language.CRJ:     "Kirī Tonga-mā-Rāwhiti",
			language.CRK:     "Pareina Kirī",
			language.CRL:     "Kirī Raki-mā-Rāwhiti",
			language.CRM:     "Mūhi Kirī",
			language.CRR:     "Arakōkiana Kararaina",
			language.CS:      "Tieke",
			language.CSW:     "Wāpi Kirī",
			language.CV:      "Tiuwhāhi",
			language.CY:      "Werehi",
			language.DA:      "Teina",
			language.DAK:     "Takōta",
			language.DAR:     "Tākawa",
			language.DAV:     "Taita",
			language.DE:      "Tiamana",
			language.DE_AT:   "Tiamana Ateriana",
			language.DE_CH:   "Tiamana Ōkawa Huiterangi",
			language.DGR:     "Tōkiripi",
			language.DJE:     "Tāma",
			language.DOI:     "Tōkiri",
			language.DSB:     "Hōpiana Hakahaka",
			language.DUA:     "Tuāra",
			language.DV:      "Tīwhehi",
			language.DYO:     "Hora-Whōni",
			language.DZ:      "Tonoka",
			language.DZG:     "Tahāka",
			language.EBU:     "Emepū",
			language.EE:      "Ewe",
			language.EFI:     "Ewhiki",
			language.EKA:     "Ekatika",
			language.EL:      "Kariki",
			language.EN:      "Ingarihi",
			language.EN_AU:   "Ingarihi Ahitereiriana",
			language.EN_CA:   "Ingarihi Kānata",
			language.EN_GB:   "Ingarihi UK",
			language.EN_US:   "Ingarihi US",
			language.EO:      "Eheperāto",
			language.ES:      "Pāniora",
			language.ES_419:  "Pāniora Amerikana ki te Tonga",
			language.ES_ES:   "Pāniora Ūropi",
			language.ES_MX:   "Pāniora Mehikana",
			language.ET:      "Etōniana",
			language.EU:      "Pākihi",
			language.EWO:     "Ewāto",
			language.FA:      "Pāhiana",
			language.FA_AF:   "Tāri",
			language.FF:      "Whūra",
			language.FI:      "Whinirānia",
			language.FIL:     "Piripīno",
			language.FJ:      "Whītīana",
			language.FO:      "Wharoīhi",
			language.FON:     "Whāna",
			language.FR:      "Wīwī",
			language.FR_CA:   "Wīwī Kānata",
			language.FR_CH:   "Wīwī Huiterangi",
			language.FRC:     "Wīwī Keihana",
			language.FRR:     "Whirīhiana ki te Raki",
			language.FUR:     "Whiriūriana",
			language.FY:      "Whirīhiana ki te Uru",
			language.GA:      "Airihi",
			language.GAA:     "Kā",
			language.GD:      "Keiriki Kotimana",
			language.GEZ:     "Kīhi",
			language.GIL:     "Kiripatīhi",
			language.GL:      "Karīhia",
			language.GN:      "Kuaranī",
			language.GOR:     "Korōtaro",
			language.GSW:     "Tiamana Huiterangi",
			language.GU:      "Kutarāti",
			language.GUZ:     "Kūhī",
			language.GV:      "Manaki",
			language.GWI:     "Kuitīna",
			language.HA:      "Hauha",
			language.HAI:     "Heira",
			language.HAW:     "Wāhu",
			language.HAX:     "Haira ki te Tonga",
			language.HE:      "Hīperu",
			language.HI:      "Hīni",
			language.HI_LATN: "Hīngarihi",
			language.HIL:     "Hirikaina",
			language.HMN:     "Mōnga",
			language.HR:      "Koroātiana",
			language.HSB:     "Hōpiana Maunga",
			language.HT:      "Kereō Haiti",
			language.HU:      "Hanekari",
			language.HUP:     "Hupa",
			language.HUR:     "Hākomerema",
			language.HY:      "Āmeniana",
			language.HZ:      "Herero",
			language.IA:      "Inarīngua",
			language.IBA:     "Īpana",
			language.IBB:     "Ipīpio",
			language.ID:      "Initonīhiana",
			language.IG:      "Ikapo",
			language.II:      "Hīhuana Eī",
			language.IKT:     "Inukitetūta Kānata ki te Uru",
			language.ILO:     "Iroko",
			language.INH:     "Inguihi",
			language.IO:      "Īto",
			language.IS:      "Tiorangi",
			language.IT:      "Itāriana",
			language.IU:      "Inukitetūta",
			language.JA:      "Hapanihi",
			language.JBO:     "Rōpāna",
			language.JGO:     "Nakōma",
			language.JMC:     "Mākame",
			language.JV:      "Hāwhanihi",
			language.KA:      "Hōriana",
			language.KAB:     "Kapāiro",
			language.KAC:     "Katīana",
			language.KAJ:     "Heiho",
			language.KAM:     "Kāmapa",
			language.KBD:     "Kapāriana",
			language.KCG:     "Tiapa",
			language.KDE:     "Makonote",
			language.KEA:     "Kapuwētianu",
			language.KFO:     "Koro",
			language.KGP:     "Keinganga",
			language.KHA:     "Kahi",
			language.KHQ:     "Kōira Tīni",
			language.KI:      "Kikūiu",
			language.KJ:      "Kuoniāma",
			language.KK:      "Kahāka",
			language.KKJ:     "Kako",
			language.KL:      "Kararīhutu",
			language.KLN:     "Karenini",
			language.KM:      "Kimēra",
			language.KMB:     "Kimipunu",
			language.KN:      "Kanara",
			language.KO:      "Kōreana",
			language.KOK:     "Kōkani",
			language.KPE:     "Kepēre",
			language.KR:      "Kanuri",
			language.KRC:     "Karatai-Pāka",
			language.KRL:     "Kareriana",
			language.KRU:     "Kuruka",
			language.KS:      "Kahimiri",
			language.KSB:     "Hapāra",
			language.KSF:     "Pāwhia",
			language.KSH:     "Korōniana",
			language.KU:      "Kūrihi",
			language.KUM:     "Kumiki",
			language.KV:      "Komi",
			language.KW:      "Kōnihi",
			language.KWK:     "Kuakawara",
			language.KY:      "Kiakihi",
			language.LA:      "Rātini",
			language.LAD:     "Ratino",
			language.LAG:     "Rangi",
			language.LB:      "Rakapuō",
			language.LEZ:     "Rēhiana",
			language.LG:      "Kānata",
			language.LI:      "Ripūkuihi",
			language.LIL:     "Riruete",
			language.LKT:     "Rakōta",
			language.LN:      "Ringāra",
			language.LO:      "Rao",
			language.LOU:     "Kreōro Ruihiana",
			language.LOZ:     "Rohi",
			language.LRC:     "Ruri ki te Raki",
			language.LSM:     "Hāmia",
			language.LT:      "Rituānia",
			language.LU:      "Rupa Katanga",
			language.LUA:     "Rupa Rurua",
			language.LUN:     "Runa",
			language.LUO:     "Ruo",
			language.LUS:     "Mīho",
			language.LUY:     "Rūia",
			language.LV:      "Rāwhia",
			language.MAD:     "Maturīhi",
			language.MAG:     "Makāhi",
			language.MAI:     "Maitiri",
			language.MAK:     "Makahā",
			language.MAS:     "Māhai",
			language.MDF:     "Mōkaha",
			language.MEN:     "Menēte",
			language.MER:     "Meru",
			language.MFE:     "Morihiene",
			language.MG:      "Marakāhi",
			language.MGH:     "Makuwa-Mēto",
			language.MGO:     "Meta",
			language.MH:      "Mararīhi",
			language.MI:      "Māori",
			language.MIC:     "Mīkamā",
			language.MIN:     "Minākapao",
			language.MK:      "Makerōnia",
			language.ML:      "Mareiārama",
			language.MN:      "Mongōria",
			language.MNI:     "Manipuri",
			language.MOE:     "Inu-aimuna",
			language.MOH:     "Mauhōka",
			language.MOS:     "Mohi",
			language.MR:      "Marati",
			language.MS:      "Marei",
			language.MT:      "Mārata",
			language.MUA:     "Mūnatanga",
			language.MUL:     "Ngā reo maha",
			language.MUS:     "Mukōki",
			language.MWL:     "Miranatīhi",
			language.MY:      "Pēmīhi",
			language.MYV:     "Erehīa",
			language.MZN:     "Mahaterani",
			language.NA:      "Nauru",
			language.NAP:     "Neaporitana",
			language.NAQ:     "Nama",
			language.NB:      "Pakamō Nōwei",
			language.ND:      "Enetepēra ki te Raki",
			language.NDS:     "Tiamana Hakahaka",
			language.NE:      "Nepari",
			language.NEW:     "Newari",
			language.NG:      "Natōka",
			language.NIA:     "Niāhi",
			language.NIU:     "Niueana",
			language.NL:      "Tati",
			language.NL_BE:   "Tati Whēmirihi",
			language.NMG:     "Kuahio",
			language.NN:      "Nīnōka Nōwei",
			language.NNH:     "Nekiepūna",
			language.NO:      "Nōwei",
			language.NOG:     "Nōkai",
			language.NQO:     "Unukō",
			language.NR:      "Enetepēra ki te Tonga",
			language.NSO:     "Hoto ki te Raki",
			language.NUS:     "Nua",
			language.NV:      "Nawahō",
			language.NY:      "Niānia",
			language.NYN:     "Niānakore",
			language.OC:      "Ōkitana",
			language.OJB:     "Ōtīpia Raki-mā-Uru",
			language.OJC:     "Ohīpawe Waenga",
			language.OJS:     "Ōti-Kirī",
			language.OJW:     "Ōhīpiwa ki te Uru",
			language.OKA:     "Ōkanakana",
			language.OM:      "Ōromo",
			language.OR:      "Ōtia",
			language.OS:      "Ōtītiki",
			language.PA:      "Punutapi",
			language.PAG:     "Pāngahina",
			language.PAM:     "Pamapaka",
			language.PAP:     "Papiamēto",
			language.PAU:     "Pārau",
			language.PCM:     "Ngāitiriana Kōrapurapu",
			language.PIS:     "Pītini",
			language.PL:      "Pōrihi",
			language.PQM:     "Marahiti-Pehamakoare",
			language.PS:      "Pāhitō",
			language.PT:      "Pōtukīhi",
			language.PT_BR:   "Pōtukīhi Parahi",
			language.PT_PT:   "Pōtukīhi Uropi",
			language.QU:      "Kētua",
			language.RAJ:     "Ratiahitani",
			language.RAP:     "Rapanui",
			language.RAR:     "Rarotonga",
			language.RHG:     "Rohingia",
			language.RM:      "Romānihi",
			language.RN:      "Rūniti",
			language.RO:      "Romeinia",
			language.ROF:     "Romopo",
			language.RU:      "Ruhiana",
			language.RUP:     "Aromeiniana",
			language.RW:      "Kiniawāna",
			language.RWK:     "Rawa",
			language.SA:      "Hanahiti",
			language.SAD:     "Hātawe",
			language.SAH:     "Hakūta",
			language.SAQ:     "Hāmapuru",
			language.SAT:     "Hatāri",
			language.SBA:     "Nekāpei",
			language.SBP:     "Hāngu",
			language.SC:      "Hārinia",
			language.SCN:     "Hihiriana",
			language.SCO:     "Kotimana",
			language.SD:      "Hiniti",
			language.SE:      "Hami ki te Raki",
			language.SEH:     "Hena",
			language.SES:     "Kōiraporo Heni",
			language.SG:      "Hāngo",
			language.SHI:     "Tāhehita",
			language.SHN:     "Hāna",
			language.SI:      "Hinihāra",
			language.SK:      "Horowākia",
			language.SL:      "Horowinia",
			language.SLH:     "Ratūti ki te Tonga",
			language.SM:      "Hāmoa",
			language.SMN:     "Inari Hami",
			language.SMS:     "Hakoto Hāmi",
			language.SN:      "Hōna",
			language.SNK:     "Honīke",
			language.SO:      "Hamāri",
			language.SQ:      "Arapeiniana",
			language.SR:      "Hirupia",
			language.SRN:     "Harāna Tongo",
			language.SS:      "Wāti",
			language.ST:      "Hōto ki te Tonga",
			language.STR:     "Hārihi Kuititanga",
			language.SU:      "Hunanīhi",
			language.SUK:     "Hukuma",
			language.SV:      "Huitene",
			language.SW:      "Wāhīri",
			language.SWB:     "Komōriana",
			language.SYR:     "Hīriaka",
			language.TA:      "Tamira",
			language.TCE:     "Tatōne ki te Tonga",
			language.TE:      "Teruku",
			language.TEM:     "Tīmene",
			language.TEO:     "Teho",
			language.TET:     "Tetumu",
			language.TG:      "Tāhiki",
			language.TGX:     "Tākihi",
			language.TH:      "Tai",
			language.THT:     "Tātana",
			language.TI:      "Tekirinia",
			language.TIG:     "Tīkara",
			language.TK:      "Tākamana",
			language.TLH:     "Kirīngona",
			language.TLI:     "Tirīkiti",
			language.TN:      "Hawāna",
			language.TO:      "Tonga",
			language.TOK:     "Toki Pona",
			language.TPI:     "Toko Pīhini",
			language.TR:      "Tākei",
			language.TRV:     "Taroko",
			language.TS:      "Honga",
			language.TT:      "Tatā",
			language.TTM:     "Tūtone ki te Raki",
			language.TUM:     "Tumūka",
			language.TVL:     "Tuwaru",
			language.TWQ:     "Tahawaka",
			language.TY:      "Tahiti",
			language.TYV:     "Tuwīniana",
			language.TZM:     "Tamahīta Te Puku o Atarihi",
			language.UDM:     "Ūmutu",
			language.UG:      "Wīkura",
			language.UK:      "Ukareinga",
			language.UMB:     "Ūpunu",
			language.UND:     "Reo Tē Mōhiotia",
			language.UR:      "Ūrutu",
			language.UZ:      "Ūpeke",
			language.VAI:     "Wai",
			language.VE:      "Wēnera",
			language.VI:      "Whitināmu",
			language.VUN:     "Whunio",
			language.WA:      "Warūna",
			language.WAE:     "Wāhere",
			language.WAL:     "Wareita",
			language.WAR:     "Warei",
			language.WO:      "Warawhe",
			language.WUU:     "Hainamana Wū",
			language.XAL:     "Karamiki",
			language.XH:      "Tōha",
			language.XOG:     "Hoka",
			language.YAV:     "Angapene",
			language.YBB:     "Emapa",
			language.YI:      "Irihi",
			language.YO:      "Ōrūpa",
			language.YRL:     "Nīkātū",
			language.YUE:     "Hainamana, Katonīhi",
			language.ZGH:     "Moroko Tamatai",
			language.ZH:      "Hainamana Manarini",
			language.ZH_HANS: "Hainamana Māmā",
			language.ZH_HANT: "Hainamana Tukuiho",
			language.ZU:      "Tūru",
			language.ZUN:     "Tuni",
			language.ZXX:     "Wetereo kiko kore",
			language.ZZA:     "Tātā",
		},
		Territories: cldr.Territories{
			territory.V_001: "te ao",
			territory.V_002: "Āwherika",
			territory.V_003: "Amerika ki te Raki",
			territory.V_005: "Amerika ki te Tonga",
			territory.V_009: "Ngā Moutere-a-Kiwa",
			territory.V_011: "Āwherika ki te Uru",
			territory.V_013: "Te Puku o Amerika",
			territory.V_014: "Āwherika ki te Rāwhiti",
			territory.V_015: "Āwherika ki te Raki",
			territory.V_017: "Te Pokapū o Āwherika",
			territory.V_018: "Āwherika Whakatetonga",
			territory.V_019: "Ngā Amerika",
			territory.V_021: "Te Raki o Amerika",
			territory.V_029: "Karapīana",
			territory.V_030: "Āhia ki te Rāwhiti",
			territory.V_034: "Āhia ki te Tonga",
			territory.V_035: "Āhia ki te Tonga-mā-rāwhiti",
			territory.V_039: "Ūropi ki te Tonga",
			territory.V_053: "Atareiria",
			territory.V_054: "Meranīhia",
			territory.V_057: "Te Rohe o Mekanēhia",
			territory.V_061: "Te Moana-nui-a-Kiwa",
			territory.V_142: "Āhia",
			territory.V_143: "Te Puku o Āhia",
			territory.V_145: "Āhia ki te Uru",
			territory.V_150: "Ūropi",
			territory.V_151: "Ūropi ki te Rāwhiti",
			territory.V_154: "Ūropi ki te Raki",
			territory.V_155: "Ūropi ki te Uru",
			territory.V_202: "Āwherika ki te Tonga o Te Hahāra",
			territory.V_419: "Amerika Rātini",
			territory.AC:    "Te Moutere Aupikinga",
			territory.AD:    "Anatōra",
			territory.AE:    "Kotahitanga o ngā Whenua o Ārapi",
			territory.AF:    "Awhekenetāna",
			territory.AG:    "Motu Nehe me Pāputa",
			territory.AI:    "Anguira",
			territory.AL:    "Arapeinia",
			territory.AM:    "Āmenia",
			territory.AO:    "Anakora",
			territory.AQ:    "Te Kōpakatanga ki te Tonga",
			territory.AR:    "Āketina",
			territory.AS:    "Hāmoa-Amerika",
			territory.AT:    "Ataria",
			territory.AU:    "Ahitereiria",
			territory.AW:    "Arūpa",
			territory.AX:    "Motu Ōrana",
			territory.AZ:    "Atepaihānia",
			territory.BA:    "Pōngia-Herekōwini",
			territory.BB:    "Papatohe",
			territory.BD:    "Pākaratēhi",
			territory.BE:    "Peretiama",
			territory.BF:    "Pākina Wharo",
			territory.BG:    "Purukāria",
			territory.BH:    "Pāreina",
			territory.BI:    "Puruniti",
			territory.BJ:    "Penīna",
			territory.BL:    "Hato Pāteremi",
			territory.BM:    "Pāmura",
			territory.BN:    "Poronai",
			territory.BO:    "Poriwia",
			territory.BQ:    "Karapīana Hōrana",
			territory.BR:    "Parīhi",
			territory.BS:    "Pahama",
			territory.BT:    "Pūtana",
			territory.BV:    "Motu Pūwei",
			territory.BW:    "Poriwana",
			territory.BY:    "Pērara",
			territory.BZ:    "Perīhi",
			territory.CA:    "Kānata",
			territory.CC:    "Ngā Moutere Kokoko (Kirini)",
			territory.CD:    "Kōngo - Kinihāha",
			territory.CF:    "Te Whenua Tūhake o Āwherika Waenga",
			territory.CG:    "Kōngo - Pārawhe",
			territory.CH:    "Huiterangi",
			territory.CI:    "Te Tai Rei",
			territory.CK:    "Kuki Airani",
			territory.CL:    "Hiri",
			territory.CM:    "Kamarūna",
			territory.CN:    "Haina",
			territory.CO:    "Koromōpia",
			territory.CP:    "Te Moutere Kiripetone",
			territory.CR:    "Koto Rīka",
			territory.CU:    "Kiupa",
			territory.CV:    "Te Kūrae Matomato",
			territory.CW:    "Kurahao",
			territory.CX:    "Te Moutere Kirihimete",
			territory.CY:    "Haipara",
			territory.CZ:    "Tiekia",
			territory.DE:    "Tiamana",
			territory.DG:    "Tieko Kāhia",
			territory.DJ:    "Tipūti",
			territory.DK:    "Tenemāka",
			territory.DM:    "Tominika",
			territory.DO:    "Te Whenua Tūhake o Tominika",
			territory.DZ:    "Aratiria",
			territory.EA:    "Hūta me Merera",
			territory.EC:    "Ekuatoa",
			territory.EE:    "Etōnia",
			territory.EG:    "Īhipa",
			territory.EH:    "Hahāra ki te Tonga",
			territory.ER:    "Eritēria",
			territory.ES:    "Peina",
			territory.ET:    "Etiopia",
			territory.EU:    "Te Uniana o Ūropi",
			territory.EZ:    "Te Rohe o Ūropi",
			territory.FI:    "Whinarana",
			territory.FJ:    "Whītī",
			territory.FK:    "Motu Whākarangi",
			territory.FM:    "Mekanēhia",
			territory.FO:    "Motu Wharau",
			territory.FR:    "Wīwī",
			territory.GA:    "Kāpona",
			territory.GB:    "Te Hononga o Piritene",
			territory.GD:    "Kerenāta",
			territory.GE:    "Hōria",
			territory.GF:    "Kiāna Wīwī",
			territory.GG:    "Kōnihi",
			territory.GH:    "Kāna",
			territory.GI:    "Kāmaka",
			territory.GL:    "Whenuakāriki",
			territory.GM:    "Kamopia",
			territory.GN:    "Kini",
			territory.GP:    "Kuatarupa",
			territory.GQ:    "Kini Ekuatoria",
			territory.GR:    "Kirihi",
			territory.GS:    "Hōria ki te Tonga me ngā Motu Hanawiti ki te Tonga",
			territory.GT:    "Kuatamāra",
			territory.GU:    "Kuama",
			territory.GW:    "Kini-Pihao",
			territory.GY:    "Kaiana",
			territory.HK:    "Hongipua Haina",
			territory.HM:    "Ngā Moutere Heriti me Makitānara",
			territory.HN:    "Honotura",
			territory.HR:    "Koroātia",
			territory.HT:    "Haiti",
			territory.HU:    "Hanekari",
			territory.IC:    "Motu Kanēre",
			territory.ID:    "Initonīhia",
			territory.IE:    "Airani",
			territory.IL:    "Iharaira",
			territory.IM:    "Te Moutere Mana",
			territory.IN:    "Inia",
			territory.IO:    "Te Rohe o te Moana Īniana Piritihi",
			territory.IQ:    "Irāka",
			territory.IR:    "Irāna",
			territory.IS:    "Tiorangi",
			territory.IT:    "Itāria",
			territory.JE:    "Tōrehe",
			territory.JM:    "Hemeika",
			territory.JO:    "Hōrano",
			territory.JP:    "Hapani",
			territory.KE:    "Kenia",
			territory.KG:    "Kikitānga",
			territory.KH:    "Kamapōtia",
			territory.KI:    "Kiripati",
			territory.KM:    "Komoro",
			territory.KN:    "Hato Kiti me Newhi",
			territory.KP:    "Kōrea ki te Raki",
			territory.KR:    "Kōrea ki te Tonga",
			territory.KW:    "Kūweiti",
			territory.KY:    "Ngā Motu Keimana",
			territory.KZ:    "Katatānga",
			territory.LA:    "Rāoho",
			territory.LB:    "Repanona",
			territory.LC:    "Hato Ruhia",
			territory.LI:    "Rīkenetaina",
			territory.LK:    "Hiri Rānaka",
			territory.LR:    "Raipiria",
			territory.LS:    "Teroto",
			territory.LT:    "Rituānia",
			territory.LU:    "Rakapuō",
			territory.LV:    "Rāwhia",
			territory.LY:    "Ripia",
			territory.MA:    "Moroko",
			territory.MC:    "Monāko",
			territory.MD:    "Morotawa",
			territory.ME:    "Maungakororiko",
			territory.MF:    "Hato Mātene",
			territory.MG:    "Matakāhika",
			territory.MH:    "Ngā Motu Māhara",
			territory.MK:    "Makerōnia ki te Raki",
			territory.ML:    "Māri",
			territory.MM:    "Pēma",
			territory.MN:    "Mongōria",
			territory.MO:    "Makau Haina",
			territory.MP:    "Ngā Motu Mariana ki te Raki",
			territory.MQ:    "Mātiniki",
			territory.MR:    "Mauritānia",
			territory.MS:    "Monoterā",
			territory.MT:    "Mārata",
			territory.MU:    "Marihi",
			territory.MV:    "Māratiri",
			territory.MW:    "Marāwi",
			territory.MX:    "Mēhiko",
			territory.MY:    "Mareia",
			territory.MZ:    "Mohapiki",
			territory.NA:    "Namipia",
			territory.NC:    "Whenua Kanaki",
			territory.NE:    "Ngāika",
			territory.NF:    "Te Moutere Nōpoke",
			territory.NG:    "Ngāitiria",
			territory.NI:    "Nikarāhua",
			territory.NL:    "Hōrana",
			territory.NO:    "Nōwei",
			territory.NP:    "Nepōra",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Aotearoa",
			territory.OM:    "Ōmana",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Poronēhia Wīwī",
			territory.PG:    "Papua Nūkini",
			territory.PH:    "Piripīni",
			territory.PK:    "Pakitāne",
			territory.PL:    "Pōrana",
			territory.PM:    "Hato Piere & Mikerona",
			territory.PN:    "Pitikeina",
			territory.PR:    "Peta Riko",
			territory.PS:    "Ngā Rohe o Parihitini",
			territory.PT:    "Potukara",
			territory.PW:    "Pārau",
			territory.PY:    "Parakai",
			territory.QA:    "Katā",
			territory.QO:    "Ngā Moutere-a-Kiwa o Waho atu",
			territory.RE:    "Reūnio",
			territory.RO:    "Romeinia",
			territory.RS:    "Hirupia",
			territory.RU:    "Rūhia",
			territory.RW:    "Rāwana",
			territory.SA:    "Hauri Arāpia",
			territory.SB:    "Ngā Motu Horomona",
			territory.SC:    "Heikere",
			territory.SD:    "Hūtāne",
			territory.SE:    "Huitene",
			territory.SG:    "Hingapoa",
			territory.SH:    "Hato Hērena",
			territory.SI:    "Horowinia",
			territory.SJ:    "Heopara me Iana Maiana",
			territory.SK:    "Horowākia",
			territory.SL:    "Te Araone",
			territory.SM:    "Hana Marino",
			territory.SN:    "Henekara",
			territory.SO:    "Hūmārie",
			territory.SR:    "Huriname",
			territory.SS:    "Hūtāne ki te Tonga",
			territory.ST:    "Hato Tomei me Pirinipei",
			territory.SV:    "Whakaora",
			territory.SX:    "Hiti Mātene",
			territory.SY:    "Hiria",
			territory.SZ:    "Ehiwatini",
			territory.TA:    "Tiritana da Kunia",
			territory.TC:    "Koru-Kākoa",
			territory.TD:    "Kāta",
			territory.TF:    "Ngā Rohe o Wīwī ki te Tonga",
			territory.TG:    "Toko",
			territory.TH:    "Tairanga",
			territory.TJ:    "Takiritānga",
			territory.TK:    "Tokerau",
			territory.TL:    "Tīmoa ki te Rāwhiti",
			territory.TM:    "Tukumanatānga",
			territory.TN:    "Tūnihia",
			territory.TO:    "Tonga",
			territory.TR:    "Tākei",
			territory.TT:    "Tirinaki Tōpako",
			territory.TV:    "Tūwaru",
			territory.TW:    "Taiwana",
			territory.TZ:    "Tānahia",
			territory.UA:    "Ukareinga",
			territory.UG:    "Ukānga",
			territory.UM:    "Ngā Moutere Amerika o Waho",
			territory.UN:    "Te Rūnanga Whakakotahi i ngā Iwi o te Ao",
			territory.US:    "Hononga o Amerika",
			territory.UY:    "Urukoi",
			territory.UZ:    "Uhipeketāne",
			territory.VA:    "Te Poho-o-Pita",
			territory.VC:    "Hato Wēneti me Keretīni",
			territory.VE:    "Penehūera",
			territory.VG:    "Ngā Moutere Puhi Piritene",
			territory.VI:    "Ngā Moutere Puhi Amerika",
			territory.VN:    "Whitināmu",
			territory.VU:    "Whenuatū",
			territory.WF:    "Warihi me Whutuna",
			territory.WS:    "Hāmoa",
			territory.XA:    "Mita kikoika",
			territory.XB:    "Piri Kikoika",
			territory.XK:    "Kōhoro",
			territory.YE:    "Īmene",
			territory.YT:    "Māiota",
			territory.ZA:    "Āwherika ki te Tonga",
			territory.ZM:    "Tāmipia",
			territory.ZW:    "Timuwawe",
			territory.ZZ:    "Rohe Tē Mōhiotia",
		},
	}
}
