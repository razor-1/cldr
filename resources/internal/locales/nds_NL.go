// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_nds_NL() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "nds_NL",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, 'de' d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "d.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "'Klock' H.mm:ss (zzzz)", Long: "'Klock' H.mm:ss (z)", Medium: "'Klock' H.mm:ss", Short: "'Kl'. H.mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "UTC{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan.", Feb: "Feb.", Mar: "März", Apr: "Apr.", May: "Mai", Jun: "Juni", Jul: "Juli", Aug: "Aug.", Sep: "Sep.", Oct: "Okt.", Nov: "Nov.", Dec: "Dez."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januaar", Feb: "Februaar", Mar: "März", Apr: "April", May: "Mai", Jun: "Juni", Jul: "Juli", Aug: "August", Sep: "September", Oct: "Oktover", Nov: "November", Dec: "Dezember"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sü.", Mon: "Ma.", Tue: "Di.", Wed: "Mi.", Thu: "Du.", Fri: "Fr.", Sat: "Sa."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "M", Thu: "D", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sünndag", Mon: "Maandag", Tue: "Dingsdag", Wed: "Middeweken", Thu: "Dunnersdag", Fri: "Freedag", Sat: "Sünnavend"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "vm", PM: "nm"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "vm", PM: "nm"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Austraalsch Dollar", Symbol: "AU$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Brasiliaansch Real", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Kanaadsch Dollar", Symbol: "CA$"},
				currency.CHF: cldr.Currency{DisplayName: "Swiezer Franken", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Chineesch Yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "Däänsch Kroon", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Britsch Pund Sterling", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Hongkong-Dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Indoneesch Rupje", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indsch Rupje", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Japaansch Yen", Symbol: "JP¥"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Söödkoreansch Won", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.MGA: cldr.Currency{DisplayName: "", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MXN: cldr.Currency{DisplayName: "Mexikaansch Peso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Norweegsch Kroon", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Poolsch Zloty", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "Russ’sch Ruvel", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudsch Rial", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "Sweedsch Kroon", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "Thailannsch Baht", Symbol: "฿"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Törksch Lira", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Nieg Taiwan-Dollar", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "US-Dollar", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Nich begäng Geldsoort", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Söödafrikaansch Rand", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "Afar",
			language.AB:      "Abchaasch",
			language.ACE:     "Aceh",
			language.ACH:     "Acholi",
			language.ADA:     "Adangme",
			language.ADY:     "Adygeisch",
			language.AE:      "Avestsch",
			language.AF:      "Afrikaans",
			language.AFH:     "Afrihili",
			language.AGQ:     "Aghem",
			language.AIN:     "Ainu",
			language.AK:      "Akan",
			language.AKK:     "Akkadsch",
			language.ALE:     "Aleutsch",
			language.ALT:     "Süüd-Altaisch",
			language.AM:      "Amhaarsch",
			language.AN:      "Aragoneesch",
			language.ANG:     "Ooldengelsch",
			language.ANP:     "Angika",
			language.AR:      "Araabsch",
			language.AR_001:  "Standardaraabsch",
			language.ARC:     "Aramääsch",
			language.ARN:     "Araukaansch",
			language.ARP:     "Arapaho",
			language.ARW:     "Arawak-Spraken",
			language.AS:      "Assameesch",
			language.AST:     "Asturiaansch",
			language.AV:      "Awaarsch",
			language.AWA:     "Awadhi",
			language.AY:      "Aymara",
			language.AZ:      "Aserbaidschaansch",
			language.BA:      "Baschkiersch",
			language.BAL:     "Belutschisch",
			language.BAN:     "Balineesch",
			language.BAS:     "Basaa",
			language.BE:      "Wittruss’sch",
			language.BEJ:     "Bedscha",
			language.BEM:     "Bemba",
			language.BG:      "Bulgaarsch",
			language.BHO:     "Bhodschpuri",
			language.BI:      "Bislama",
			language.BIK:     "Bikol",
			language.BIN:     "Bini",
			language.BLA:     "Siksika",
			language.BM:      "Bambara",
			language.BN:      "Bengaalsch",
			language.BO:      "Tibeetsch",
			language.BR:      "Bretoonsch",
			language.BRA:     "Braj-Bhakha",
			language.BRX:     "Bodo",
			language.BS:      "Bosnisch",
			language.BUA:     "Burjaatsch",
			language.BUG:     "Bugineesch",
			language.BYN:     "Blin",
			language.CA:      "Katalaansch",
			language.CAD:     "Caddo",
			language.CAR:     "Kariebsche Spraken",
			language.CCH:     "Atsam",
			language.CE:      "Tschetscheensch",
			language.CEB:     "Cebuano",
			language.CH:      "Chamorro",
			language.CHB:     "Chibcha-Spraken",
			language.CHG:     "Tschagataisch",
			language.CHK:     "Trukeesch",
			language.CHM:     "Mari",
			language.CHN:     "Chinook",
			language.CHO:     "Choctaw",
			language.CHP:     "Chipewyan",
			language.CHR:     "Cherokeesch",
			language.CHY:     "Cheyenne",
			language.CKB:     "Zentraalkurdsch",
			language.CO:      "Koorssch",
			language.COP:     "Koptsch",
			language.CR:      "Cree",
			language.CRH:     "Krimtataarsch",
			language.CS:      "Tschech’sch",
			language.CSB:     "Kaschuubsch",
			language.CU:      "Karkenslaavsch",
			language.CV:      "Tschuwasch’sch",
			language.CY:      "Waliesch",
			language.DA:      "Däänsch",
			language.DAK:     "Dakota",
			language.DAR:     "Dargiensch",
			language.DE:      "Hoochdüütsch",
			language.DE_AT:   "Öösterrieksch Hoochdüütsch",
			language.DE_CH:   "Swiezer Hoochdüütsch",
			language.DEL:     "Delaware",
			language.DEN:     "Slave",
			language.DGR:     "Dogrib",
			language.DIN:     "Dinka",
			language.DOI:     "Dogri",
			language.DSB:     "Neddersorbsch",
			language.DUA:     "Duala",
			language.DUM:     "Middelnedderlandsch",
			language.DV:      "Maledievsch",
			language.DYU:     "Dyula",
			language.DZ:      "Bhutaansch",
			language.EE:      "Ewe",
			language.EFI:     "Efik",
			language.EGY:     "Ägyptsch",
			language.EKA:     "Ekajuk",
			language.EL:      "Greeksch",
			language.ELX:     "Elaamsch",
			language.EN:      "Engelsch",
			language.EN_AU:   "Austraalsch Engelsch",
			language.EN_CA:   "Kanaadsch Engelsch",
			language.EN_GB:   "Engelsch (GB)",
			language.EN_US:   "Engelsch (US)",
			language.ENM:     "Middelengelsch",
			language.EO:      "Esperanto",
			language.ES:      "Spaansch",
			language.ES_419:  "Latienamerikaansch Spaansch",
			language.ES_ES:   "Ibeersch Spaansch",
			language.ES_MX:   "Mexikaansch Spaansch",
			language.ET:      "Eestlannsch",
			language.EU:      "Basksch",
			language.EWO:     "Ewondo",
			language.FA:      "Pers’sch",
			language.FAN:     "Pangwe",
			language.FAT:     "Fanti",
			language.FF:      "Ful",
			language.FI:      "Finnsch",
			language.FIL:     "Philippiensch",
			language.FJ:      "Fidschiaansch",
			language.FO:      "Färöösch",
			language.FON:     "Fon",
			language.FR:      "Franzöösch",
			language.FR_CA:   "Kanaadsch Franzöösch",
			language.FR_CH:   "Swiezer Franzöösch",
			language.FRM:     "Middelfranzöösch",
			language.FRO:     "Ooldfranzöösch",
			language.FRR:     "Noordfreesch",
			language.FRS:     "Saterfreesch",
			language.FUR:     "Friuulsch",
			language.FY:      "Westfreesch",
			language.GA:      "Iersch",
			language.GAA:     "Ga",
			language.GAY:     "Gayo",
			language.GBA:     "Gbaya",
			language.GD:      "Schottsch Gäälsch",
			language.GEZ:     "Geez",
			language.GIL:     "Gilberteesch",
			language.GL:      "Galizsch",
			language.GMH:     "Middelhoochdüütsch",
			language.GN:      "Guarani",
			language.GOH:     "Ooldhoochdüütsch",
			language.GON:     "Gondi",
			language.GOR:     "Gorontalo",
			language.GOT:     "Gootsch",
			language.GRB:     "Grebo",
			language.GRC:     "Ooldgreeksch",
			language.GSW:     "Swiezerdüütsch",
			language.GU:      "Gudscharati",
			language.GV:      "Manx",
			language.GWI:     "Kutchin",
			language.HA:      "Haussa",
			language.HAI:     "Haida",
			language.HAW:     "Hawaiiaansch",
			language.HE:      "Hebrääsch",
			language.HI:      "Hindi",
			language.HIL:     "Hiligaynon",
			language.HIT:     "Hethitsch",
			language.HMN:     "Miao-Spraken",
			language.HO:      "Hiri-Motu",
			language.HR:      "Kroaatsch",
			language.HSB:     "Böversorbsch",
			language.HT:      "Haitiaansch",
			language.HU:      "Ungaarsch",
			language.HUP:     "Hupa",
			language.HY:      "Armeensch",
			language.HZ:      "Herero",
			language.IA:      "Interlingua",
			language.IBA:     "Iban",
			language.ID:      "Indoneesch",
			language.IE:      "Interlingue",
			language.IG:      "Igbo",
			language.II:      "Sichuan Yi",
			language.IK:      "Inupiak",
			language.ILO:     "Ilokano",
			language.INH:     "Ingusch’sch",
			language.IO:      "Ido",
			language.IS:      "Ieslannsch",
			language.IT:      "Italieensch",
			language.IU:      "Inuktitut",
			language.JA:      "Japaansch",
			language.JBO:     "Lojban",
			language.JPR:     "Jöödsch-Pers’sch",
			language.JRB:     "Jöödsch-Araabsch",
			language.JV:      "Javaansch",
			language.KA:      "Georgsch",
			language.KAA:     "Karakalpaksch",
			language.KAB:     "Kabyylsch",
			language.KAC:     "Kachin",
			language.KAJ:     "Jju",
			language.KAM:     "Kamba",
			language.KAW:     "Kawi",
			language.KBD:     "Kabardiensch",
			language.KCG:     "Tyap",
			language.KFO:     "Koro",
			language.KG:      "Kongo",
			language.KHA:     "Khasi",
			language.KHO:     "Saaksch",
			language.KI:      "Kikuyu",
			language.KJ:      "Kwanyama",
			language.KK:      "Kasach’sch",
			language.KL:      "Gröönlannsch",
			language.KM:      "Kambodschaansch",
			language.KMB:     "Kimbundu",
			language.KN:      "Kannada",
			language.KO:      "Koreaansch",
			language.KOI:     "Komipermjaksch",
			language.KOK:     "Konkani",
			language.KOS:     "Kosraeaansch",
			language.KPE:     "Kpelle",
			language.KR:      "Kanuursch",
			language.KRC:     "Karatschaisch-Balkaarsch",
			language.KRL:     "Kareelsch",
			language.KRU:     "Oraon",
			language.KS:      "Kaschmiersch",
			language.KU:      "Kurdsch",
			language.KUM:     "Kumücksch",
			language.KUT:     "Kutenai",
			language.KV:      "Komi",
			language.KW:      "Koornsch",
			language.KY:      "Kirgiesch",
			language.LA:      "Latiensch",
			language.LAD:     "Ladiensch",
			language.LAH:     "Lahnda",
			language.LAM:     "Lamba",
			language.LB:      "Luxemborgsch",
			language.LEZ:     "Lesgisch",
			language.LG:      "Luganda",
			language.LI:      "Limborgsch",
			language.LN:      "Lingala",
			language.LO:      "Laootsch",
			language.LOL:     "Mongo",
			language.LOZ:     "Rotse",
			language.LT:      "Litausch",
			language.LU:      "Luba",
			language.LUA:     "Luba-Lulua",
			language.LUI:     "Luiseno",
			language.LUN:     "Lunda",
			language.LUO:     "Luo",
			language.LUS:     "Lushai",
			language.LV:      "Lettsch",
			language.MAD:     "Madureesch",
			language.MAG:     "Magahi",
			language.MAI:     "Maithili",
			language.MAK:     "Makassarsch",
			language.MAN:     "Manding",
			language.MAS:     "Massai",
			language.MDF:     "Mokscha",
			language.MDR:     "Mandareesch",
			language.MEN:     "Mende",
			language.MG:      "Madagassisch",
			language.MGA:     "Middeliersch",
			language.MH:      "Marschalleesch",
			language.MI:      "Maori",
			language.MIC:     "Micmac",
			language.MIN:     "Minangkabau",
			language.MK:      "Mazedoonsch",
			language.ML:      "Malayalam",
			language.MN:      "Mongoolsch",
			language.MNC:     "Mandschuursch",
			language.MNI:     "Manipuri",
			language.MOH:     "Mohawk",
			language.MOS:     "Mossi",
			language.MR:      "Marathi",
			language.MS:      "Malaisch",
			language.MT:      "Malteesch",
			language.MUL:     "Mehrsprakig",
			language.MUS:     "Muskogee-Spraken",
			language.MWL:     "Mirandeesch",
			language.MWR:     "Marwari",
			language.MY:      "Birmaansch",
			language.MYV:     "Erzya",
			language.NA:      "Nauruusch",
			language.NAP:     "Neapolitaansch",
			language.NB:      "Norweegsch Bokmål",
			language.ND:      "Noord-Ndebele",
			language.NDS:     "Neddersass’sch",
			language.NE:      "Nepaleesch",
			language.NEW:     "Newari",
			language.NG:      "Ndonga",
			language.NIA:     "Nias",
			language.NIU:     "Niue",
			language.NL:      "Nedderlandsch",
			language.NL_BE:   "Fläämsch",
			language.NN:      "Norweegsch Nynorsk",
			language.NO:      "Norweegsch",
			language.NOG:     "Nogai",
			language.NON:     "Ooldnoorsch",
			language.NQO:     "N’Ko",
			language.NR:      "Süüd-Ndebele",
			language.NSO:     "Noord-Sotho",
			language.NV:      "Navajo",
			language.NWC:     "Oold-Newari",
			language.NY:      "Nyanja",
			language.NYM:     "Nyamwezi",
			language.NYN:     "Nyankole",
			language.NYO:     "Nyoro",
			language.NZI:     "Nzima",
			language.OC:      "Okzitaansch",
			language.OJ:      "Ojibwa",
			language.OM:      "Oromo",
			language.OR:      "Orija",
			language.OS:      "Ossetsch",
			language.OSA:     "Osage",
			language.OTA:     "Osmaansch",
			language.PA:      "Pandschaabsch",
			language.PAG:     "Pangasinan",
			language.PAL:     "Middelpers’sch",
			language.PAM:     "Pampanggan",
			language.PAP:     "Papiamento",
			language.PAU:     "Palausch",
			language.PEO:     "Ooldpers’sch",
			language.PHN:     "Phönieksch",
			language.PI:      "Pali",
			language.PL:      "Poolsch",
			language.PON:     "Ponapeaansch",
			language.PRO:     "Ooldprovenzaalsch",
			language.PS:      "Paschtu",
			language.PT:      "Portugeesch",
			language.PT_BR:   "Brasiliaansch Portugeesch",
			language.PT_PT:   "Ibeersch Portugeesch",
			language.QU:      "Quechua",
			language.RAJ:     "Rajasthani",
			language.RAP:     "Oosterinsel-Spraak",
			language.RAR:     "Rarotongaansch",
			language.RM:      "Rätoromaansch",
			language.RN:      "Rundi",
			language.RO:      "Rumäänsch",
			language.RO_MD:   "Moldaawsch",
			language.ROM:     "Romani",
			language.ROOT:    "Wortel",
			language.RU:      "Russ’sch",
			language.RUP:     "Aromuunsch",
			language.RW:      "Ruandsch",
			language.SA:      "Sanskrit",
			language.SAD:     "Sandawe",
			language.SAH:     "Jakuutsch",
			language.SAM:     "Samaritaansch",
			language.SAS:     "Sasak",
			language.SAT:     "Santali",
			language.SC:      "Sardsch",
			language.SCN:     "Siziliaansch",
			language.SCO:     "Schottsch",
			language.SD:      "Sindhi",
			language.SE:      "Noord-Saamsch",
			language.SEL:     "Selkupsch",
			language.SG:      "Sango",
			language.SGA:     "Oold-Iersch",
			language.SHN:     "Schan",
			language.SI:      "Singhaleesch",
			language.SID:     "Sidamo",
			language.SK:      "Slowaaksch",
			language.SL:      "Sloweensch",
			language.SM:      "Samoaansch",
			language.SMA:     "Süüd-Lappsch",
			language.SMJ:     "Lule-Lappsch",
			language.SMN:     "Inari-Lappsch",
			language.SMS:     "Skolt-Lappsch",
			language.SN:      "Schona",
			language.SNK:     "Soninke",
			language.SO:      "Somaalsch",
			language.SOG:     "Sogdisch",
			language.SQ:      "Albaansch",
			language.SR:      "Serbsch",
			language.SRN:     "Surinaamsch",
			language.SRR:     "Serer",
			language.SS:      "Swazi",
			language.ST:      "Süüd-Sotho",
			language.SU:      "Sundaneesch",
			language.SUK:     "Sukuma",
			language.SUS:     "Susu",
			language.SUX:     "Sumersch",
			language.SV:      "Sweedsch",
			language.SW:      "Suaheli",
			language.SYC:     "Oold-Syyrsch",
			language.SYR:     "Syyrsch",
			language.TA:      "Tamilsch",
			language.TE:      "Telugu",
			language.TEM:     "Temne",
			language.TER:     "Tereno",
			language.TET:     "Tetum",
			language.TG:      "Tadschiksch",
			language.TH:      "Thailannsch",
			language.TI:      "Tigrinja",
			language.TIG:     "Tigre",
			language.TIV:     "Tiv",
			language.TK:      "Turkmeensch",
			language.TKL:     "Tokelausch",
			language.TL:      "Tagalog",
			language.TLH:     "Klingoonsch",
			language.TLI:     "Tlingit",
			language.TMH:     "Tamaschek",
			language.TN:      "Tswana",
			language.TO:      "Tongaasch",
			language.TOG:     "Tonga (Nyasa)",
			language.TPI:     "Tok Pisin",
			language.TR:      "Törksch",
			language.TS:      "Tsonga",
			language.TSI:     "Tsimshian",
			language.TT:      "Tataarsch",
			language.TUM:     "Tumbuka",
			language.TVL:     "Elliceaansch",
			language.TW:      "Twi",
			language.TY:      "Tahitsch",
			language.TYV:     "Tuwinsch",
			language.UDM:     "Udmurtsch",
			language.UG:      "Uiguursch",
			language.UGA:     "Ugaritsch",
			language.UK:      "Ukrainsch",
			language.UMB:     "Mbundu",
			language.UND:     "Nich begäng Spraak",
			language.UR:      "Urdu",
			language.UZ:      "Usbeeksch",
			language.VAI:     "Vai",
			language.VE:      "Venda",
			language.VI:      "Vietnameesch",
			language.VO:      "Volapük",
			language.VOT:     "Wootsch",
			language.WA:      "Walloonsch",
			language.WAL:     "Walamo",
			language.WAR:     "Waray",
			language.WAS:     "Washo",
			language.WO:      "Wolof",
			language.XAL:     "Kalmücksch",
			language.XH:      "Xhosa",
			language.YAO:     "Yao",
			language.YAP:     "Yapeesch",
			language.YI:      "Jiddisch",
			language.YO:      "Yoruba",
			language.ZA:      "Zhuang",
			language.ZAP:     "Zapoteeksch",
			language.ZBL:     "Bliss-Symbolen",
			language.ZEN:     "Zenaga",
			language.ZH:      "Chineesch",
			language.ZH_HANS: "Vereenfacht Chineesch",
			language.ZH_HANT: "Traditschonell Chineesch",
			language.ZU:      "Zulu",
			language.ZUN:     "Zuni",
			language.ZXX:     "Keen Spraakinhold",
			language.ZZA:     "Zaza",
		},
		Territories: cldr.Territories{
			territory.V_001: "Welt",
			territory.V_002: "Afrika",
			territory.V_003: "Noordamerika",
			territory.V_005: "Süüdamerika",
			territory.V_009: "Ozeanien",
			territory.V_011: "Westafrika",
			territory.V_013: "Middelamerika",
			territory.V_014: "Oostafrika",
			territory.V_015: "Noordafrika",
			territory.V_017: "Zentralafrika",
			territory.V_018: "Süüdlich Afrika",
			territory.V_019: "Amerika",
			territory.V_029: "Karibik",
			territory.V_030: "Oostasien",
			territory.V_034: "Süüdasien",
			territory.V_035: "Süüdoostasien",
			territory.V_039: "Süüdeuropa",
			territory.V_053: "Australien un Neeseeland",
			territory.V_054: "Melanesien",
			territory.V_061: "Polynesien",
			territory.V_142: "Asien",
			territory.V_143: "Zentralasien",
			territory.V_145: "Westasien",
			territory.V_150: "Europa",
			territory.V_151: "Oosteuropa",
			territory.V_154: "Noordeuropa",
			territory.V_155: "Westeuropa",
			territory.V_419: "Latienamerika",
			territory.AD:    "Andorra",
			territory.AE:    "Vereenigte Araabsche Emiraten",
			territory.AF:    "Afghanistan",
			territory.AG:    "Antigua un Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albanien",
			territory.AM:    "Armenien",
			territory.AO:    "Angola",
			territory.AQ:    "Antarktis",
			territory.AR:    "Argentinien",
			territory.AS:    "Amerikaansch-Samoa",
			territory.AT:    "Öösterriek",
			territory.AU:    "Australien",
			territory.AW:    "Aruba",
			territory.AX:    "Ålandeilannen",
			territory.AZ:    "Aserbaidschan",
			territory.BA:    "Bosnien un Herzegowina",
			territory.BB:    "Barbados",
			territory.BD:    "Bangladesch",
			territory.BE:    "Belgien",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgarien",
			territory.BH:    "Bahrain",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "Saint Barthélemy",
			territory.BM:    "Bermuda",
			territory.BN:    "Brunei Darussalam",
			territory.BO:    "Bolivien",
			territory.BR:    "Brasilien",
			territory.BS:    "Bahamas",
			territory.BT:    "Bhutan",
			territory.BV:    "Bouvet-Eiland",
			territory.BW:    "Botswana",
			territory.BY:    "Wittrussland",
			territory.BZ:    "Belize",
			territory.CA:    "Kanada",
			territory.CC:    "Kokos-Eilannen",
			territory.CD:    "Demokraatsche Republik Kongo",
			territory.CF:    "Zentralafrikaansche Republik",
			territory.CG:    "Republik Kongo",
			territory.CH:    "Swiez",
			territory.CI:    "Elfenbeenküst",
			territory.CK:    "Cook-Eilannen",
			territory.CL:    "Chile",
			territory.CM:    "Kamerun",
			territory.CN:    "China",
			territory.CO:    "Kolumbien",
			territory.CR:    "Costa Rica",
			territory.CU:    "Kuba",
			territory.CV:    "Kap Verde",
			territory.CX:    "Wiehnachtseiland",
			territory.CY:    "Zypern",
			territory.CZ:    "Tschechien",
			territory.DE:    "Düütschland",
			territory.DJ:    "Dschibuti",
			territory.DK:    "Däänmark",
			territory.DM:    "Dominica",
			territory.DO:    "Dominikaansche Republik",
			territory.DZ:    "Algerien",
			territory.EC:    "Ecuador",
			territory.EE:    "Eestland",
			territory.EG:    "Ägypten",
			territory.EH:    "Westsahara",
			territory.ER:    "Eritrea",
			territory.ES:    "Spanien",
			territory.ET:    "Äthiopien",
			territory.EU:    "Europääsche Union",
			territory.FI:    "Finnland",
			territory.FJ:    "Fidschi",
			territory.FK:    "Falkland-Eilannen",
			territory.FM:    "Mikronesien",
			territory.FO:    "Färöer",
			territory.FR:    "Frankriek",
			territory.GA:    "Gabun",
			territory.GB:    "Grootbritannien",
			territory.GD:    "Grenada",
			territory.GE:    "Georgien",
			territory.GF:    "Franzöösch-Guayana",
			territory.GG:    "Guernsey",
			territory.GH:    "Ghana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Gröönland",
			territory.GM:    "Gambia",
			territory.GN:    "Guinea",
			territory.GP:    "Guadeloupe",
			territory.GQ:    "Äquatorialguinea",
			territory.GR:    "Grekenland",
			territory.GS:    "Süüdgeorgien un de Südlichen Sandwich-Eilannen",
			territory.GT:    "Guatemala",
			territory.GU:    "Guam",
			territory.GW:    "Guinea-Bissau",
			territory.GY:    "Guyana",
			territory.HK:    "Hongkong",
			territory.HM:    "Heard- un McDonald-Eilannen",
			territory.HN:    "Honduras",
			territory.HR:    "Kroatien",
			territory.HT:    "Haiti",
			territory.HU:    "Ungarn",
			territory.ID:    "Indonesien",
			territory.IE:    "Irland",
			territory.IL:    "Israel",
			territory.IM:    "Insel Man",
			territory.IN:    "Indien",
			territory.IO:    "Britisch Rebeed in’n Indischen Ozean",
			territory.IQ:    "Irak",
			territory.IR:    "Iran",
			territory.IS:    "Iesland",
			territory.IT:    "Italien",
			territory.JE:    "Jersey",
			territory.JM:    "Jamaika",
			territory.JO:    "Jordanien",
			territory.JP:    "Japan",
			territory.KE:    "Kenia",
			territory.KG:    "Kirgisistan",
			territory.KH:    "Kambodscha",
			territory.KI:    "Kiribati",
			territory.KM:    "Komoren",
			territory.KN:    "St. Kitts un Nevis",
			territory.KP:    "Noordkorea",
			territory.KR:    "Söödkorea",
			territory.KW:    "Kuwait",
			territory.KY:    "Kaiman-Eilannen",
			territory.KZ:    "Kasachstan",
			territory.LA:    "Laos",
			territory.LB:    "Libanon",
			territory.LC:    "St. Lucia",
			territory.LI:    "Liechtensteen",
			territory.LK:    "Sri Lanka",
			territory.LR:    "Liberia",
			territory.LS:    "Lesotho",
			territory.LT:    "Litauen",
			territory.LU:    "Luxemborg",
			territory.LV:    "Lettland",
			territory.LY:    "Libyen",
			territory.MA:    "Marokko",
			territory.MC:    "Monaco",
			territory.MD:    "Moldawien",
			territory.ME:    "Montenegro",
			territory.MF:    "Saint Martin",
			territory.MG:    "Madagaskar",
			territory.MH:    "Marshall-Eilannen",
			territory.MK:    "Makedonien",
			territory.ML:    "Mali",
			territory.MM:    "Birma",
			territory.MN:    "Mongolei",
			territory.MO:    "Macao",
			territory.MP:    "Nöördliche Marianen",
			territory.MQ:    "Martinique",
			territory.MR:    "Mauretanien",
			territory.MS:    "Montserrat",
			territory.MT:    "Malta",
			territory.MU:    "Mauritius",
			territory.MV:    "Malediven",
			territory.MW:    "Malawi",
			territory.MX:    "Mexiko",
			territory.MY:    "Malaysia",
			territory.MZ:    "Mosambik",
			territory.NA:    "Namibia",
			territory.NC:    "Neekaledonien",
			territory.NE:    "Niger",
			territory.NF:    "Norfolk",
			territory.NG:    "Nigeria",
			territory.NI:    "Nikaragua",
			territory.NL:    "Nedderlannen",
			territory.NO:    "Norwegen",
			territory.NP:    "Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Neeseeland",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Franzöösch-Polynesien",
			territory.PG:    "Papua-Neeguinea",
			territory.PH:    "Philippinen",
			territory.PK:    "Pakistan",
			territory.PL:    "Polen",
			territory.PM:    "St. Pierre un Miquelon",
			territory.PN:    "Pitcairn",
			territory.PR:    "Puerto Rico",
			territory.PS:    "Palästinensische Rebeden",
			territory.PT:    "Portugal",
			territory.PW:    "Palau",
			territory.PY:    "Paraguay",
			territory.QA:    "Katar",
			territory.QO:    "Büter Ozeanien",
			territory.RE:    "Reunion",
			territory.RO:    "Rumänien",
			territory.RS:    "Serbien",
			territory.RU:    "Russland",
			territory.RW:    "Ruanda",
			territory.SA:    "Saudi-Arabien",
			territory.SB:    "Salomonen",
			territory.SC:    "Seychellen",
			territory.SD:    "Sudan",
			territory.SE:    "Sweden",
			territory.SG:    "Singapur",
			territory.SH:    "St. Helena",
			territory.SI:    "Slowenien",
			territory.SJ:    "Svalbard un Jan Mayen",
			territory.SK:    "Slowakei",
			territory.SL:    "Sierra Leone",
			territory.SM:    "San Marino",
			territory.SN:    "Senegal",
			territory.SO:    "Somalia",
			territory.SR:    "Surinam",
			territory.ST:    "São Tomé un Príncipe",
			territory.SV:    "El Salvador",
			territory.SY:    "Syrien",
			territory.SZ:    "Swasiland",
			territory.TC:    "Turks- un Caicosinseln",
			territory.TD:    "Tschad",
			territory.TF:    "Franzöösche Süüd- un Antarktisrebeden",
			territory.TG:    "Togo",
			territory.TH:    "Thailand",
			territory.TJ:    "Tadschikistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Oosttimor",
			territory.TM:    "Turkmenistan",
			territory.TN:    "Tunesien",
			territory.TO:    "Tonga",
			territory.TR:    "Törkei",
			territory.TT:    "Trinidad un Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tansania",
			territory.UA:    "Ukraine",
			territory.UG:    "Uganda",
			territory.UM:    "Amerikaansch-Ozeanien",
			territory.US:    "USA",
			territory.UY:    "Uruguay",
			territory.UZ:    "Usbekistan",
			territory.VA:    "Vatikan",
			territory.VC:    "St. Vincent un de Grenadinen",
			territory.VE:    "Venezuela",
			territory.VG:    "Brietsche Jumfern-Eilannen",
			territory.VI:    "Amerikaansche Jumfern-Eilannen",
			territory.VN:    "Vietnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis un Futuna",
			territory.WS:    "Samoa",
			territory.YE:    "Jemen",
			territory.YT:    "Mayotte",
			territory.ZA:    "Söödafrika",
			territory.ZM:    "Sambia",
			territory.ZW:    "Simbabwe",
			territory.ZZ:    "Nich begäng Regioon",
		},
	}
}
