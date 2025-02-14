// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_scn_IT() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "scn_IT",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jin", Feb: "fri", Mar: "mar", Apr: "apr", May: "maj", Jun: "giu", Jul: "gnt", Aug: "agu", Sep: "sit", Oct: "utt", Nov: "nuv", Dec: "dic"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "G", Jul: "G", Aug: "A", Sep: "S", Oct: "U", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "jinnaru", Feb: "frivaru", Mar: "marzu", Apr: "aprili", May: "maju", Jun: "giugnu", Jul: "giugnettu", Aug: "agustu", Sep: "sittèmmiru", Oct: "uttùviru", Nov: "nuvèmmiru", Dec: "dicèmmiru"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dum", Mon: "lun", Tue: "mar", Wed: "mer", Thu: "jov", Fri: "ven", Sat: "sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "d", Mon: "l", Tue: "m", Wed: "m", Thu: "j", Fri: "v", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "du", Mon: "lu", Tue: "ma", Wed: "me", Thu: "jo", Fri: "ve", Sat: "sa"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "dumìnica", Mon: "lunnidìa", Tue: "martidìa", Wed: "mercuridìa", Thu: "jovidìa", Fri: "venniridìa", Sat: "sàbbatu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham di l’Emirati Àrabbi Junciuti", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghani afghanu", Symbol: "؋"},
				currency.ALL: cldr.Currency{DisplayName: "Lek arbanisi", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Dram armenu", Symbol: "֏"},
				currency.ANG: cldr.Currency{DisplayName: "Ciurinu di l’Antiḍḍi Ulannisi", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza angulisi", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Pesu argintinu", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dòllaru australianu", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Ciurinu d’Arubba", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Manat azzeru", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "Marcu cummirtìbbili dâ Bosnia-Herzegòvina", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Dòllaru dî Barbados", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Taka dû Bàngladesh", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Lev bùrgaru", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dìnaru dû Bahrain", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Francu dû Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Dòllaru dî Birmuda", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Dòllaru dû Brunei", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Bulivianu", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Riali brasilianu", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dòllaru dî Bahamas", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum butanisi", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Pula dû Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Rubblu belurrussu", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Dòllaru dû Bilisi", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dòllaru canadisi", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Francu cungulisi", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Francu sbìzziru", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Pesu cilenu", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Yuan cinisi (di fora)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuan cinisi", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Pesu culummianu", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Culón dâ Custa Rica", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Pesu cubbanu cummirtìbbili", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Pesu cubbanu", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Scudu capuvirdisi", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Curuna ceca", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Francu di Gibbuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Curuna danisi", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Pesu duminicanu", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dìnaru argirinu", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Stirlina eggizziana", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nafka eritreu", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr etiupi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euru", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dòllaru dî Figi", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Stirlina di l’Ìsuli Falkland", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Stirlina Britànnica", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Lari giurgianu", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "Cedi ganisi", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "Stirlina di Gibbirterra", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi dû Gammia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Francu dâ Guinìa", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal dâ Guatimala", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Dòllaru dâ Guiana", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Dòllaru di Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Limpira di l’Hunnuras", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna cruata", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Gordu d’Haiti", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Ciurinu unghirisi", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Rupìa innunisiana", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Novu siclu isdraelianu", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupìa inniana", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dìnaru irachenu", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Riali iranianu", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Curuna islannisi", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Dòllaru giamaicanu", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Dìnaru giurdanu", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yen giappunisi", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Scillinu dû Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Som dû Kyrgyzistan", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "Riel cambuggianu", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Francu dî Cumori", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Won dâ Curìa di Tramuntana", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Won dâ Curìa di Sciroccu", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dìnaru dû Kuwait", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Dòllaru di l’Ìsuli Cayman", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge dû Kazzàkistan", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Kip lauisi", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Stirlina libbanisi", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Rupìa dû Sri Lanka", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dòllaru dâ Libberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti dû Lisothu", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dìnaru lìbbicu", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirham marucchinu", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Leu murdavu", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariary margasciu", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Dìnaru macèduni", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kyat dû Myanmar", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik mòngulu", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca di Macau", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya mauritanu", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupìa di Mauritius", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa dî Mardivi", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha dû Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Pesu missicanu", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit malisi", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Mètical dû Muzzammicu", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dòllaru dâ Namibbia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira niggirianu", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Còrdubba dâ Nicaragua", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Curuna nurviggisi", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Rupìa nipalisi", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Dòllaru dâ Nova Zilannia", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Riali di l’Oman", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Barboa di Pànama", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Suli piruvianu", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina dâ Papua Nova Guinìa", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Pesu filippinu", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Rupìa pakistana", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty pulaccu", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Guaranì dû Paraguay", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Riali dû Qatar", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Leu rumenu", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dìnaru serbu", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Rubblu russu", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "Francu dû Ruanna", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riali di l’Arabbia Saudita", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Dòllaru di l’Ìsuli Salumuni", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupìa dî Seychelles", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Stirlina sudanisi", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Curuna svidisi", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Dòllaru di Singapuri", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Stirlina di Sant’Èlina", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Liuni dâ Sierra Liuni", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Liuni dâ Sierra Liuni (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Scillinu sòmalu", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Dòllaru dû Surinami", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Stirlina dû Sudan di sciroccu", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "Dobra di São Tomé & Príncipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Stirlina siriana", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni di Eswatini", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Baht tailannisi", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Somoni dû Tajìkistan", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Manat turkmenu", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Dìnaru tunisinu", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Tongan Paʻanga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Lira turca", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Dòllaru di Trinidad e Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Novu dòllaru taiwanisi", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Scillinu dâ Tanzania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Grivnia ucràina", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Scillinu di l’Uganna", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dòllaru miricanu", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Pesu di l’Uruguay", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Som di l’Uzbèkistan", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Bulivar dû Vinizzuela", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Dong vietnamisi", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu di Vanuatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tala samuanu", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Francu CFA di l’Àfrica cintrali", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dòllaru dî Caraibbi di livanti", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Francu CFA di l’Àfrica di punenti", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "Francu CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Munita Scanusciuta", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Riali dû Yemen", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Rand di l’Àfrica di Sciroccu", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha dâ Zammia", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "afar",
			language.AB:      "abkhasu",
			language.AF:      "afrikaans",
			language.AGQ:     "aghem",
			language.AK:      "akan",
			language.AM:      "amàricu",
			language.AN:      "aragunisi",
			language.APC:     "àrabbu livantinu di tramuntana",
			language.AR:      "àrabbu",
			language.AR_001:  "àrabbu nadaru mudernu",
			language.ARN:     "mapuche",
			language.AS:      "assamisi",
			language.ASA:     "asu",
			language.AST:     "asturianu",
			language.AZ:      "azzeru",
			language.BA:      "bashkir",
			language.BAL:     "baluchi",
			language.BAS:     "basaa",
			language.BE:      "belurrussu",
			language.BEM:     "bemba",
			language.BEW:     "betawi",
			language.BEZ:     "bena",
			language.BG:      "bùrgaru",
			language.BGC:     "haryanvi",
			language.BHO:     "bhojipuri",
			language.BLO:     "anii",
			language.BM:      "bambara",
			language.BN:      "bangladisi",
			language.BR:      "brètuni",
			language.BRX:     "bodu",
			language.BS:      "busnìacu",
			language.BSS:     "akoose",
			language.BYN:     "blin",
			language.CA:      "catalanu",
			language.CAD:     "caddu",
			language.CCH:     "atsam",
			language.CCP:     "chakma",
			language.CE:      "cicenu",
			language.CEB:     "cebbuanu",
			language.CGG:     "chiga",
			language.CHO:     "choctaw",
			language.CHR:     "cherokee",
			language.CIC:     "chickasaw",
			language.CKB:     "curdu Sorani",
			language.CO:      "corsu",
			language.CS:      "cecu",
			language.CSW:     "swampy cree",
			language.CU:      "slavu dâ cresia",
			language.CV:      "ciuvasciu",
			language.CY:      "gallisi",
			language.DA:      "danisi",
			language.DE:      "tidiscu",
			language.DE_AT:   "tidiscu austrìacu",
			language.DE_CH:   "tidiscu autu sbìzziru",
			language.DOI:     "dogri",
			language.DSB:     "sorbu suttanu",
			language.DUA:     "duala",
			language.DV:      "divehi",
			language.DYO:     "jola-fonyi",
			language.DZ:      "dzongkha",
			language.EBU:     "embu",
			language.EE:      "ewe",
			language.EL:      "grecu",
			language.EN:      "ngrisi",
			language.EN_AU:   "ngrisi australianu",
			language.EN_CA:   "ngrisi canadisi",
			language.EN_GB:   "ngrisi britànnicu",
			language.EN_US:   "ngrisi miricanu",
			language.EO:      "spirantu",
			language.ES:      "spagnolu",
			language.ES_419:  "spagnolu dâ mèrica latina",
			language.ES_ES:   "spagnolu eurupeu",
			language.ES_MX:   "spagnolu missicanu",
			language.ET:      "èstuni",
			language.EU:      "bascu",
			language.EWO:     "ewondo",
			language.FA:      "pirsianu",
			language.FF:      "fula",
			language.FI:      "fillannisi",
			language.FIL:     "filippinu",
			language.FO:      "faruisi",
			language.FR:      "francisi",
			language.FR_CA:   "francisi canadisi",
			language.FR_CH:   "francisi sbìzziru",
			language.FRC:     "francisi Cajun",
			language.FUR:     "friulanu",
			language.FY:      "frìsuni uccidintali",
			language.GA:      "irlannisi",
			language.GAA:     "ga",
			language.GD:      "gaèlicu scuzzisi",
			language.GEZ:     "geez",
			language.GL:      "galizzianu",
			language.GN:      "guarani",
			language.GSW:     "tidiscu sbìzziru",
			language.GU:      "gujarati",
			language.GUZ:     "gusii",
			language.GV:      "mannisi",
			language.HA:      "hausa",
			language.HAW:     "hawaiianu",
			language.HE:      "ebbràicu",
			language.HI:      "innianu",
			language.HI_LATN: "innianu ngrisi",
			language.HNJ:     "hmong njua",
			language.HR:      "cruatu",
			language.HSB:     "surbu supranu",
			language.HU:      "unghirisi",
			language.HY:      "armenu",
			language.IA:      "ntirlingua",
			language.ID:      "innunisianu",
			language.IE:      "ntirlingui",
			language.IG:      "igbo",
			language.IO:      "idu",
			language.IS:      "islannisi",
			language.IT:      "talianu",
			language.IU:      "inuktitut",
			language.JA:      "giappunisi",
			language.JBO:     "lojban",
			language.JGO:     "ngomba",
			language.JMC:     "machame",
			language.JV:      "giavanisi",
			language.KA:      "giurgianu",
			language.KAA:     "kara-kalpak",
			language.KAB:     "kabyle",
			language.KAJ:     "jiu",
			language.KAM:     "kamba",
			language.KDE:     "makonde",
			language.KEA:     "capuvirdianu",
			language.KEN:     "kenyang",
			language.KGP:     "kaingang",
			language.KHQ:     "koyra chiini",
			language.KI:      "kikuyu",
			language.KK:      "kazaku",
			language.KKJ:     "kako",
			language.KL:      "kalaallisut",
			language.KLN:     "kalenjin",
			language.KM:      "khmer",
			language.KN:      "kannada",
			language.KO:      "curianu",
			language.KOK:     "konkani",
			language.KPE:     "kpelle",
			language.KS:      "kashmiri",
			language.KSF:     "bafia",
			language.KSH:     "culunisi",
			language.KU:      "curdu",
			language.KW:      "còrnicu",
			language.KXV:     "kuvi",
			language.KY:      "kirghizu",
			language.LA:      "latinu",
			language.LAG:     "langi",
			language.LB:      "lussimmurghisi",
			language.LG:      "ganda",
			language.LIJ:     "lìguri",
			language.LKT:     "lakota",
			language.LLD:     "ild",
			language.LMO:     "lummardu",
			language.LN:      "lingala",
			language.LO:      "lau",
			language.LOU:     "criolu dâ Louisiana",
			language.LT:      "lituanu",
			language.LTG:     "latgallisi",
			language.LU:      "luba-katanga",
			language.LUY:     "luyia",
			language.LV:      "lèttuni",
			language.MAI:     "maithili",
			language.MAS:     "masai",
			language.MDF:     "moksha",
			language.MER:     "meru",
			language.MFE:     "murisianu",
			language.MG:      "margasciu",
			language.MGH:     "makhuwa-meetto",
			language.MGO:     "metaʼ",
			language.MHN:     "muchenu",
			language.MI:      "māori",
			language.MIC:     "mi'kmaw",
			language.MK:      "macèduni",
			language.ML:      "malayalam",
			language.MN:      "mòngulu",
			language.MNI:     "manipuri",
			language.MOH:     "mohawk",
			language.MR:      "marathi",
			language.MS:      "malisi",
			language.MT:      "mautisi",
			language.MUA:     "mundang",
			language.MUL:     "assai lingui",
			language.MUS:     "muscogee",
			language.MY:      "burmisi",
			language.MYV:     "erzya",
			language.MZN:     "mazanderani",
			language.NAQ:     "nama",
			language.NB:      "nurviggisi Bokmål",
			language.NDS:     "tidiscu suttanu",
			language.NDS_NL:  "sàssuni suttanu",
			language.NE:      "nipalisi",
			language.NL:      "ulannisi",
			language.NL_BE:   "ciammingu",
			language.NMG:     "kwasio",
			language.NN:      "nurviggisi Nynorsk",
			language.NNH:     "ngiemboon",
			language.NO:      "nurviggisi",
			language.NQO:     "n’ko",
			language.NV:      "navajo",
			language.OC:      "uccitanu",
			language.OR:      "odia",
			language.PA:      "punjabi",
			language.PCM:     "pidgin niggirianu",
			language.PL:      "pulaccu",
			language.PRG:     "prussianu",
			language.PS:      "pashto",
			language.PT:      "purtughisi",
			language.PT_BR:   "purtughisi brasilianu",
			language.PT_PT:   "purtughisi eurupeu",
			language.QU:      "quechua",
			language.QUC:     "kʼicheʼ",
			language.RAJ:     "rajasthani",
			language.RM:      "rumanciu",
			language.RO:      "rumenu",
			language.RO_MD:   "murdavu",
			language.RU:      "russu",
			language.RW:      "kinyarwanda",
			language.SA:      "sàscritu",
			language.SAH:     "yakut",
			language.SAT:     "santali",
			language.SC:      "sardu",
			language.SCN:     "sicilianu",
			language.SD:      "sindhi",
			language.SES:     "koyraboro senni",
			language.SI:      "sinhala",
			language.SK:      "sluvaccu",
			language.SL:      "sluvenu",
			language.SMJ:     "lule sami",
			language.SMN:     "inari sami",
			language.SO:      "sòmalu",
			language.SQ:      "arbanisi",
			language.SR:      "serbu",
			language.SU:      "sunnanisi",
			language.SV:      "svidisi",
			language.SW:      "swahili",
			language.SYR:     "sirìacu",
			language.SZL:     "silisianu",
			language.TA:      "tamil",
			language.TE:      "telugu",
			language.TG:      "tajik",
			language.TH:      "tailannisi",
			language.TI:      "tigrinya",
			language.TK:      "turkmenu",
			language.TO:      "tunganu",
			language.TR:      "turcu",
			language.TT:      "tàtaru",
			language.TZM:     "tamazight di l’Atlanti Cintrali",
			language.UG:      "uyghur",
			language.UK:      "ucrainu",
			language.UND:     "lingua scanusciuta",
			language.UR:      "urdu",
			language.UZ:      "uzbeku",
			language.VEC:     "vènitu",
			language.VI:      "vietnamisi",
			language.VMW:     "makhuwa",
			language.VO:      "volapük",
			language.WO:      "wolof",
			language.XH:      "xhosa",
			language.XNR:     "kangri",
			language.YO:      "yoruba",
			language.YRL:     "nheengatu",
			language.YUE:     "cinisi cantunisi",
			language.ZA:      "zhuang",
			language.ZH:      "cinisi, mannarinu",
			language.ZH_HANS: "cinisi mannarinu simprificatu",
			language.ZH_HANT: "cinisi mannarinu tradizziunali",
			language.ZU:      "zulu",
			language.ZXX:     "nuḍḍu cuntinutu linguìsticu",
		},
		Territories: cldr.Territories{
			territory.V_001: "Munnu",
			territory.V_002: "Àfrica",
			territory.V_003: "Mèrica di tramuntana cuntinintali",
			territory.V_005: "Mèrica di sciroccu",
			territory.V_009: "Uciània",
			territory.V_011: "Àfrica punintina",
			territory.V_013: "Mèrica cintrali",
			territory.V_014: "Àfrica livantina",
			territory.V_015: "Àfrica di tramuntana",
			territory.V_017: "Àfrica di menzu",
			territory.V_018: "Àfrica di sciroccu",
			territory.V_019: "Mèrichi",
			territory.V_021: "Mèrica di tramuntana",
			territory.V_029: "Caràibbi",
			territory.V_030: "Asia livantina",
			territory.V_034: "Asia di sciroccu",
			territory.V_035: "Asia di sciroccu-livantina",
			territory.V_039: "Europa di sciroccu",
			territory.V_053: "Australasia",
			territory.V_054: "Milanesia",
			territory.V_057: "Riggiuni dâ Micrunesia",
			territory.V_061: "Pulinesia",
			territory.V_142: "Asia",
			territory.V_143: "Asia cintrali",
			territory.V_145: "Asia punintina",
			territory.V_150: "Europa",
			territory.V_151: "Europa livantina",
			territory.V_154: "Europa di tramuntana",
			territory.V_155: "Europa punintina",
			territory.V_202: "Àfrica sutta-sahariana",
			territory.V_419: "Mèrica latina",
			territory.AC:    "Ìsula d’Ascinziuni",
			territory.AD:    "Annorra",
			territory.AE:    "Emirati Àrabbi Junciuti",
			territory.AF:    "Afghànistan",
			territory.AG:    "Antigua e Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Arbanìa",
			territory.AM:    "Armenia",
			territory.AO:    "Angola",
			territory.AQ:    "Antàrtidi",
			territory.AR:    "Argintina",
			territory.AS:    "Samoa Miricani",
			territory.AT:    "Austria",
			territory.AU:    "Australia",
			territory.AW:    "Aruba",
			territory.AX:    "Ìsuli Åland",
			territory.AZ:    "Azerbaijan",
			territory.BA:    "Bosnia e Herzegòvina",
			territory.BB:    "Barbados",
			territory.BD:    "Bàngladesh",
			territory.BE:    "Bergiu",
			territory.BF:    "Burkina Fasu",
			territory.BG:    "Burgarìa",
			territory.BH:    "Bahrain",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "St. Barthélemy",
			territory.BM:    "Birmuda",
			territory.BN:    "Brunei",
			territory.BO:    "Bulivia",
			territory.BQ:    "Caràibbi ulannisi",
			territory.BR:    "Brasili",
			territory.BS:    "Bahamas",
			territory.BT:    "Bhutan",
			territory.BV:    "Ìsula Bouvet",
			territory.BW:    "Botswana",
			territory.BY:    "Belurussia",
			territory.BZ:    "Bilisi",
			territory.CA:    "Cànada",
			territory.CC:    "Ìsuli Cocos (Keeling)",
			territory.CD:    "Congu - Kinshasa",
			territory.CF:    "Ripùbblica Centrafricana",
			territory.CG:    "Congu - Brazzaville",
			territory.CH:    "Sbìzzira",
			territory.CI:    "Custa d’Avoriu",
			territory.CK:    "Ìsuli Cook",
			territory.CL:    "Cili",
			territory.CM:    "Càmerun",
			territory.CN:    "Cina",
			territory.CO:    "Culommia",
			territory.CP:    "Ìsula di Clipperton",
			territory.CR:    "Custa Rica",
			territory.CU:    "Cubba",
			territory.CV:    "Capu Virdi",
			territory.CW:    "Curaçao",
			territory.CX:    "Ìsula di Natali",
			territory.CY:    "Cipru",
			territory.CZ:    "Cechia",
			territory.DE:    "Girmania",
			territory.DG:    "Diego Garcia",
			territory.DJ:    "Gibbuti",
			territory.DK:    "Danimarca",
			territory.DM:    "Dumìnica",
			territory.DO:    "Ripùbblica Duminicana",
			territory.DZ:    "Algirìa",
			territory.EA:    "Ceuta e Miliḍḍa",
			territory.EC:    "Ècuador",
			territory.EE:    "Estonia",
			territory.EG:    "Eggittu",
			territory.EH:    "Sahara punintinu",
			territory.ER:    "Eritrea",
			territory.ES:    "Spagna",
			territory.ET:    "Etiopia",
			territory.EU:    "Uniuni Eurupea",
			territory.EZ:    "Zuna Euru",
			territory.FI:    "Fillannia",
			territory.FJ:    "Figi",
			territory.FK:    "Ìsuli Falkland",
			territory.FM:    "Micrunisia",
			territory.FO:    "Ìsuli Faroe",
			territory.FR:    "Franza",
			territory.GA:    "Gabon",
			territory.GB:    "Regnu Junciutu",
			territory.GD:    "Grenada",
			territory.GE:    "Giorgia",
			territory.GF:    "Guiana Francisi",
			territory.GG:    "Guernsey",
			territory.GH:    "Ghana",
			territory.GI:    "Gibbirterra",
			territory.GL:    "Gruillannia",
			territory.GM:    "Gammia",
			territory.GN:    "Guinìa",
			territory.GP:    "Guadalupa",
			territory.GQ:    "Guinìa Equaturiali",
			territory.GR:    "Grecia",
			territory.GS:    "Giorgia di sciroccu e Ìsuli Sandwich australi",
			territory.GT:    "Guatimala",
			territory.GU:    "Guam",
			territory.GW:    "Guinìa-Bissau",
			territory.GY:    "Guiana",
			territory.HK:    "Hong Kong RAS dâ Cina",
			territory.HM:    "Ìsuli Heard e McDonald",
			territory.HN:    "Hunnuras",
			territory.HR:    "Cruazzia",
			territory.HT:    "Haiti",
			territory.HU:    "Ungarìa",
			territory.IC:    "Ìsuli Canari",
			territory.ID:    "Innunesia",
			territory.IE:    "Irlanna",
			territory.IL:    "Isdraeli",
			territory.IM:    "Ìsula di Man",
			territory.IN:    "Innia",
			territory.IO:    "Tirritoriu Ucianicu di l’Innia Britànnica",
			territory.IQ:    "Iraq",
			territory.IR:    "Iran",
			territory.IS:    "Islanna",
			territory.IT:    "Italia",
			territory.JE:    "Jersey",
			territory.JM:    "Giamàica",
			territory.JO:    "Giurdania",
			territory.JP:    "Giappuni",
			territory.KE:    "Kenya",
			territory.KG:    "Kyrgyzistan",
			territory.KH:    "Camboggia",
			territory.KI:    "Kiribati",
			territory.KM:    "Còmoros",
			territory.KN:    "S. Kitts e Nevis",
			territory.KP:    "Curìa di Tramuntana",
			territory.KR:    "Curìa di Sciroccu",
			territory.KW:    "Kuwait",
			territory.KY:    "Ìsuli Cayman",
			territory.KZ:    "Kazzàkistan",
			territory.LA:    "Laos",
			territory.LB:    "Lìbbanu",
			territory.LC:    "Santa Lucìa",
			territory.LI:    "Liechtenstein",
			territory.LK:    "Sri Lanka",
			territory.LR:    "Libberia",
			territory.LS:    "Lisothu",
			territory.LT:    "Lituania",
			territory.LU:    "Lussimmurgu",
			territory.LV:    "Littonia",
			territory.LY:    "Libbia",
			territory.MA:    "Maroccu",
			territory.MC:    "Mònacu",
			territory.MD:    "Murdova",
			territory.ME:    "Muntinegru",
			territory.MF:    "San Martinu",
			territory.MG:    "Madagascàr",
			territory.MH:    "Ìsuli Marshall",
			territory.MK:    "Macidonia di tramuntana",
			territory.ML:    "Mali",
			territory.MM:    "Myanmar (Burma)",
			territory.MN:    "Mungolia",
			territory.MO:    "Macau RAS dâ Cina",
			territory.MP:    "Ìsuli Marianna di Tramuntana",
			territory.MQ:    "Martinica",
			territory.MR:    "Mauritania",
			territory.MS:    "Munzirratu",
			territory.MT:    "Mauta",
			territory.MU:    "Mauritius",
			territory.MV:    "Mardivi",
			territory.MW:    "Malawi",
			territory.MX:    "Mèssicu",
			territory.MY:    "Malesia",
			territory.MZ:    "Muzzammicu",
			territory.NA:    "Namibbia",
			territory.NC:    "Nova Calidonia",
			territory.NE:    "Niger",
			territory.NF:    "Ìsula Norfolk",
			territory.NG:    "Niggeria",
			territory.NI:    "Nicaragua",
			territory.NL:    "Paisi Vasci",
			territory.NO:    "Nurveggia",
			territory.NP:    "Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Nova Zilannia",
			territory.OM:    "Oman",
			territory.PA:    "Pànama",
			territory.PE:    "Pirù",
			territory.PF:    "Pulinisia Francisi",
			territory.PG:    "Papua Nova Guinìa",
			territory.PH:    "Filippini",
			territory.PK:    "Pàkistan",
			territory.PL:    "Pulonia",
			territory.PM:    "S. Pierre e Miquelon",
			territory.PN:    "Ìsuli Pitcairn",
			territory.PR:    "Portu Ricu",
			territory.PS:    "Tirritori Palistinesi",
			territory.PT:    "Purtugallu",
			territory.PW:    "Palau",
			territory.PY:    "Paraguay",
			territory.QA:    "Qatar",
			territory.QO:    "Uciània di fora",
			territory.RE:    "Réunion",
			territory.RO:    "Rumanìa",
			territory.RS:    "Serbia",
			territory.RU:    "Russia",
			territory.RW:    "Ruanna",
			territory.SA:    "Arabbia Saudita",
			territory.SB:    "Ìsuli Salumuni",
			territory.SC:    "Seychelles",
			territory.SD:    "Sudan",
			territory.SE:    "Sbezzia",
			territory.SG:    "Singapuri",
			territory.SH:    "Sant’Èlina",
			territory.SI:    "Sluvenia",
			territory.SJ:    "Svalbard e Jan Mayen",
			territory.SK:    "Sluvacchia",
			territory.SL:    "Sierra Liuni",
			territory.SM:    "San Marinu",
			territory.SN:    "Sènigal",
			territory.SO:    "Sumalia",
			territory.SR:    "Surinami",
			territory.SS:    "Sudan di sciroccu",
			territory.ST:    "São Tomé e Príncipe",
			territory.SV:    "El Salvador",
			territory.SX:    "Sint Maarten",
			territory.SY:    "Siria",
			territory.SZ:    "Eswatini",
			territory.TA:    "Tristan da Cunha",
			territory.TC:    "Ìsuli Turks e Càicos",
			territory.TD:    "Chad",
			territory.TF:    "Tirritori Francisi di Sciroccu",
			territory.TG:    "Togu",
			territory.TH:    "Tailannia",
			territory.TJ:    "Tajìkistan",
			territory.TK:    "Tukilau",
			territory.TL:    "Timor-Leste",
			territory.TM:    "Turkmènistan",
			territory.TN:    "Tunisìa",
			territory.TO:    "Tonga",
			territory.TR:    "Turchìa",
			territory.TT:    "Trinidad e Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tanzania",
			territory.UA:    "Ucraina",
			territory.UG:    "Uganna",
			territory.UM:    "Ìsuli Miricani di Fora",
			territory.UN:    "Nazziuna Junciuti",
			territory.US:    "Stati Junciuti",
			territory.UY:    "Uruguay",
			territory.UZ:    "Uzbèkistan",
			territory.VA:    "Città dû Vaticanu",
			territory.VC:    "S. Vincent e Grenadine",
			territory.VE:    "Vinizzuela",
			territory.VG:    "Ìsuli Vìrgini Britànnichi",
			territory.VI:    "Ìsuli Vìrgini Miricani",
			territory.VN:    "Vietnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis e Futuna",
			territory.WS:    "Samoa",
			territory.XA:    "Accenti fausi",
			territory.XB:    "Bidirizziunali fausu",
			territory.XK:    "Kòssuvu",
			territory.YE:    "Yemen",
			territory.YT:    "Mayotte",
			territory.ZA:    "Africa di sciroccu",
			territory.ZM:    "Zammia",
			territory.ZW:    "Zimbabwe",
			territory.ZZ:    "Riggiuni scanusciuta",
		},
	}
}
