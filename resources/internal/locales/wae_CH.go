// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_wae_CH() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "wae_CH",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jen", Feb: "Hor", Mar: "Mär", Apr: "Abr", May: "Mei", Jun: "Brá", Jul: "Hei", Aug: "Öig", Sep: "Her", Oct: "Wím", Nov: "Win", Dec: "Chr"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "H", Mar: "M", Apr: "A", May: "M", Jun: "B", Jul: "H", Aug: "Ö", Sep: "H", Oct: "W", Nov: "W", Dec: "C"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Jenner", Feb: "Hornig", Mar: "Märze", Apr: "Abrille", May: "Meije", Jun: "Bráčet", Jul: "Heiwet", Aug: "Öigšte", Sep: "Herbštmánet", Oct: "Wímánet", Nov: "Wintermánet", Dec: "Chrištmánet"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sun", Mon: "Män", Tue: "Ziš", Wed: "Mit", Thu: "Fró", Fri: "Fri", Sat: "Sam"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "Z", Wed: "M", Thu: "F", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sunntag", Mon: "Mäntag", Tue: "Zištag", Wed: "Mittwuč", Thu: "Fróntag", Fri: "Fritag", Sat: "Samštag"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "’", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "¤\u00a00K", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Brasilianiši Real", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Chinesiši Yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pfund", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indiši Rupie", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yen", Symbol: "¥"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
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
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "Rubel", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "Dollar", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Unbekannti Wãrig", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AB:      "Abčasiš",
			language.AF:      "Afrikáns",
			language.AM:      "Amhariš",
			language.AR:      "Arabiš",
			language.AS:      "Assamesiš",
			language.AY:      "Aymara",
			language.AZ:      "Serbaidšaniš",
			language.BE:      "Wísrussiš",
			language.BG:      "Bulgariš",
			language.BN:      "Bengališ",
			language.BO:      "Tibetiš",
			language.BS:      "Bosniš",
			language.CA:      "Katalaniš",
			language.CS:      "Tšečiš",
			language.CY:      "Walisiš",
			language.DA:      "Däniš",
			language.DE:      "Titš",
			language.DE_AT:   "Öštričišes Titš",
			language.DE_CH:   "Schwizer Hočtitš",
			language.DV:      "Malediwiš",
			language.DZ:      "Butaniš",
			language.EFI:     "Efik",
			language.EL:      "Gričiš",
			language.EN:      "Engliš",
			language.EN_AU:   "Auštrališes Engliš",
			language.EN_CA:   "Kanadišes Engliš",
			language.EN_GB:   "Britišes Engliš",
			language.EN_US:   "Amerikanišes Engliš",
			language.ES:      "Schpaniš",
			language.ES_419:  "Latiamerikanišes Schpaniš",
			language.ES_ES:   "Iberišes Schpaniš",
			language.ET:      "Estniš",
			language.EU:      "Baskiš",
			language.FA:      "Persiš",
			language.FI:      "Finiš",
			language.FIL:     "Filipiniš",
			language.FJ:      "Fidšianiš",
			language.FR:      "Wälš",
			language.FR_CA:   "Kanadišes Wälš",
			language.FR_CH:   "Schwizer Wälš",
			language.GA:      "Iriš",
			language.GL:      "Galiziš",
			language.GN:      "Guarani",
			language.GU:      "Gujarati",
			language.HA:      "Hausa",
			language.HAW:     "Hawaíaniš",
			language.HE:      "Hebräiš",
			language.HI:      "Hindi",
			language.HR:      "Kroatiš",
			language.HT:      "Haitianiš",
			language.HU:      "Ungariš",
			language.HY:      "Armeniš",
			language.ID:      "Indonesiš",
			language.IG:      "Igbo",
			language.IS:      "Iisländiš",
			language.IT:      "Italieniš",
			language.JA:      "Japaniš",
			language.KA:      "Georgiš",
			language.KK:      "Kazačiš",
			language.KM:      "Kambodšaniš",
			language.KN:      "Kannada",
			language.KO:      "Koreaniš",
			language.KS:      "Kašmiriš",
			language.KU:      "Kurdiš",
			language.KY:      "Kirgisiš",
			language.LA:      "Latiniš",
			language.LB:      "Luxemburgiš",
			language.LN:      "Lingala",
			language.LO:      "Laotiš",
			language.LT:      "Litauiš",
			language.LV:      "Lettiš",
			language.MG:      "Malagási",
			language.MI:      "Maori",
			language.MK:      "Mazedoniš",
			language.ML:      "Malayalam",
			language.MN:      "Mongoliš",
			language.MR:      "Marathi",
			language.MS:      "Malaíš",
			language.MT:      "Maltesiš",
			language.MY:      "Burmesiš",
			language.NB:      "Norwegiš Bokmål",
			language.ND:      "Nordndebele",
			language.NE:      "Nepalesiš",
			language.NL:      "Holändiš",
			language.NL_BE:   "Flämiš",
			language.NN:      "Norwegiš Nynorsk",
			language.NSO:     "Nordsotho",
			language.NY:      "Nyanja",
			language.OR:      "Oriya",
			language.OS:      "Osétiš",
			language.PA:      "Pandšabiš",
			language.PL:      "Polniš",
			language.PS:      "Paštu",
			language.PT:      "Portugisiš",
			language.PT_BR:   "Brasilianišes Portugisiš",
			language.PT_PT:   "Iberišes Portugisiš",
			language.QU:      "Quečua",
			language.RM:      "Rätromaniš",
			language.RN:      "Rundi",
			language.RO:      "Rumäniš",
			language.RU:      "Rusiš",
			language.RW:      "Ruandiš",
			language.SA:      "Sanskrit",
			language.SAH:     "Jakutiš",
			language.SD:      "Sindhi",
			language.SE:      "Nordsamiš",
			language.SG:      "Sango",
			language.SI:      "Singalesiš",
			language.SK:      "Slowakiš",
			language.SL:      "Sloweniš",
			language.SM:      "Samoaniš",
			language.SN:      "Shona",
			language.SO:      "Somališ",
			language.SQ:      "Albaniš",
			language.SR:      "Serbiš",
			language.SS:      "Swazi",
			language.ST:      "Südsotho",
			language.SU:      "Sundanesiš",
			language.SV:      "Schwediš",
			language.SW:      "Suaheliš",
			language.TA:      "Tamiliš",
			language.TE:      "Telugu",
			language.TET:     "Tetum",
			language.TG:      "Tadšikiš",
			language.TH:      "Thailändiš",
			language.TI:      "Tigrinja",
			language.TK:      "Turkmeniš",
			language.TN:      "Tswana",
			language.TO:      "Tonga",
			language.TPI:     "Niwmelanesiš",
			language.TR:      "Türkiš",
			language.TS:      "Tsonga",
			language.TY:      "Taitiš",
			language.UG:      "Uiguriš",
			language.UK:      "Ukrainiš",
			language.UND:     "Unbekannti Schprač",
			language.UR:      "Urdu",
			language.UZ:      "Usbekiš",
			language.VE:      "Venda",
			language.VI:      "Vietnamesiš",
			language.WAE:     "Walser",
			language.WO:      "Wolof",
			language.XH:      "Xhosa",
			language.YO:      "Yoruba",
			language.ZH:      "Chinesiš",
			language.ZH_HANS: "Vereifačts Chinesiš",
			language.ZH_HANT: "Traditionells Chinesiš",
			language.ZU:      "Zulu",
		},
		Territories: cldr.Territories{
			territory.V_001: "Wäld",
			territory.V_002: "Afrika",
			territory.V_003: "Nordamerika",
			territory.V_005: "Südamerika",
			territory.V_009: "Ozeanie",
			territory.V_011: "Weštafrika",
			territory.V_013: "Zentralamerika",
			territory.V_014: "Oštafrika",
			territory.V_015: "Nordafrika",
			territory.V_017: "Mittelafrika",
			territory.V_018: "Südličs Afrika",
			territory.V_019: "Amerikaniš Kontinänt",
			territory.V_021: "Nördličs Amerika",
			territory.V_029: "Karibik",
			territory.V_030: "Oštasie",
			territory.V_034: "Südasie",
			territory.V_035: "Südoštasie",
			territory.V_039: "Südeuropa",
			territory.V_053: "Auštralie und Niwséland",
			territory.V_054: "Melanesie",
			territory.V_057: "Mikronesišes Inselgebiet",
			territory.V_061: "Polinesie",
			territory.V_142: "Asie",
			territory.V_143: "Zentralasie",
			territory.V_145: "Weštasie",
			territory.V_150: "Europa",
			territory.V_151: "Ošteuropa",
			territory.V_154: "Nordeuropa",
			territory.V_155: "Wešteuropa",
			territory.V_419: "Latíamerika",
			territory.AC:    "Himmelfártsinsla",
			territory.AD:    "Andorra",
			territory.AE:    "Vereinigti Arabiše Emirat",
			territory.AF:    "Afganištan",
			territory.AG:    "Antigua und Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albanie",
			territory.AM:    "Armenie",
			territory.AO:    "Angola",
			territory.AQ:    "Antarktis",
			territory.AR:    "Argentinie",
			territory.AS:    "Amerikaniš Samoa",
			territory.AT:    "Öštrič",
			territory.AU:    "Australie",
			territory.AW:    "Aruba",
			territory.AX:    "Alandinslä",
			territory.AZ:    "Aserbaidšan",
			territory.BA:    "Bosnie und Herzegovina",
			territory.BB:    "Barbados",
			territory.BD:    "Bangladeš",
			territory.BE:    "Belgie",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgarie",
			territory.BH:    "Bačrain",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "St. Bartholomäus-Insla",
			territory.BM:    "Bermuda",
			territory.BN:    "Brunei",
			territory.BO:    "Boliwie",
			territory.BR:    "Brasilie",
			territory.BS:    "Bahamas",
			territory.BT:    "Bhutan",
			territory.BV:    "Bouvetinsla",
			territory.BW:    "Botswana",
			territory.BY:    "Wísrussland",
			territory.BZ:    "Belize",
			territory.CA:    "Kanada",
			territory.CC:    "Kokosinslä",
			territory.CD:    "Kongo-Kinshasa",
			territory.CF:    "Zentralafrikaniši Rebublik",
			territory.CG:    "Kongo Brazzaville",
			territory.CH:    "Schwiz",
			territory.CI:    "Elfebeiküšta",
			territory.CK:    "Cookinslä",
			territory.CL:    "Tšile",
			territory.CM:    "Kamerun",
			territory.CN:    "China",
			territory.CO:    "Kolumbie",
			territory.CP:    "Clipperton Insla",
			territory.CR:    "Costa Rica",
			territory.CU:    "Kuba",
			territory.CV:    "Kap Verde",
			territory.CX:    "Wienäčtsinslä",
			territory.CY:    "Zypre",
			territory.CZ:    "Tšečie",
			territory.DE:    "Titšland",
			territory.DG:    "Diego Garcia",
			territory.DJ:    "Dšibuti",
			territory.DK:    "Dänemark",
			territory.DM:    "Doninica",
			territory.DO:    "Dominikaniši Rebublik",
			territory.DZ:    "Algerie",
			territory.EA:    "Ceuta und Melilla",
			territory.EC:    "Ecuador",
			territory.EE:    "Eštland",
			territory.EG:    "Egypte",
			territory.EH:    "Weštsahara",
			territory.ER:    "Eritrea",
			territory.ES:    "Schpanie",
			territory.ET:    "Ethiopie",
			territory.EU:    "Europäiši Unio",
			territory.FI:    "Finnland",
			territory.FJ:    "Fidši",
			territory.FK:    "Falklandinslä",
			territory.FM:    "Mikronesie",
			territory.FO:    "Färöe",
			territory.FR:    "Frankrič",
			territory.GA:    "Gabon",
			territory.GB:    "England",
			territory.GD:    "Grenada",
			territory.GE:    "Georgie",
			territory.GF:    "Französiš Guiana",
			territory.GG:    "Guernsey",
			territory.GH:    "Gana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Grönland",
			territory.GM:    "Gambia",
			territory.GN:    "Ginea",
			territory.GP:    "Guadeloupe",
			territory.GQ:    "Equatorialginea",
			territory.GR:    "Gričeland",
			territory.GS:    "Südgeorgie und d’südliče Senwičinslä",
			territory.GT:    "Guatemala",
			territory.GU:    "Guam",
			territory.GW:    "Ginea Bissau",
			territory.GY:    "Guyana",
			territory.HK:    "Sonderverwaltigszona Hongkong",
			territory.HM:    "Heard- und McDonald-Inslä",
			territory.HN:    "Honduras",
			territory.HR:    "Kroatie",
			territory.HT:    "Haiti",
			territory.HU:    "Ungare",
			territory.IC:    "Kanariše Inslä",
			territory.ID:    "Indonesie",
			territory.IE:    "Irland",
			territory.IL:    "Israel",
			territory.IM:    "Isle of Man",
			territory.IN:    "Indie",
			territory.IO:    "Britišes Territorium em indiše Ozean",
			territory.IQ:    "Irak",
			territory.IR:    "Iran",
			territory.IS:    "Island",
			territory.IT:    "Italie",
			territory.JE:    "Jersey",
			territory.JM:    "Jamaika",
			territory.JO:    "Jordanie",
			territory.JP:    "Japan",
			territory.KE:    "Kenya",
			territory.KG:    "Kirgištan",
			territory.KH:    "Kambodša",
			territory.KI:    "Kiribati",
			territory.KM:    "Komore",
			territory.KN:    "St. Kitts und Nevis",
			territory.KP:    "Nordkorea",
			territory.KR:    "Südkorea",
			territory.KW:    "Kuweit",
			territory.KY:    "Kaimaninslä",
			territory.KZ:    "Kasačstan",
			territory.LA:    "Laos",
			territory.LB:    "Libanon",
			territory.LC:    "St. Lucia",
			territory.LI:    "Liečteštei",
			territory.LK:    "Sri Lanka",
			territory.LR:    "Liberia",
			territory.LS:    "Lesotho",
			territory.LT:    "Litaue",
			territory.LU:    "Luxeburg",
			territory.LV:    "Lettland",
			territory.LY:    "Lübie",
			territory.MA:    "Maroko",
			territory.MC:    "Monago",
			territory.MD:    "Moldau",
			territory.ME:    "Montenegro",
			territory.MF:    "St. Martin",
			territory.MG:    "Madagaskar",
			territory.MH:    "Maršalinslä",
			territory.ML:    "Mali",
			territory.MM:    "Burma",
			territory.MN:    "Mongolei",
			territory.MO:    "Sonderverwaltigszona Makau",
			territory.MP:    "Nördliči Mariane",
			territory.MQ:    "Martinique",
			territory.MR:    "Mauretanie",
			territory.MS:    "Monserrat",
			territory.MT:    "Malta",
			territory.MU:    "Mauritius",
			territory.MV:    "Malediwe",
			territory.MW:    "Malawi",
			territory.MX:    "Mexiko",
			territory.MY:    "Malaysia",
			territory.MZ:    "Mosambik",
			territory.NA:    "Namibia",
			territory.NC:    "Niwkaledonie",
			territory.NE:    "Niger",
			territory.NF:    "Norfolkinsla",
			territory.NG:    "Nigeria",
			territory.NI:    "Nicaragua",
			territory.NL:    "Holand",
			territory.NO:    "Norwäge",
			territory.NP:    "Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Niwséland",
			territory.OM:    "Oman",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Französiš Polinesie",
			territory.PG:    "Papua Niwginea",
			territory.PH:    "Philippine",
			territory.PK:    "Pakištan",
			territory.PL:    "Pole",
			territory.PM:    "St. Pierre und Miquelon",
			territory.PN:    "Pitcairn",
			territory.PR:    "Puerto Rico",
			territory.PS:    "Paleština",
			territory.PT:    "Portugal",
			territory.PW:    "Palau",
			territory.PY:    "Paraguai",
			territory.QA:    "Katar",
			territory.QO:    "Üssers Ozeanie",
			territory.RE:    "Réunion",
			territory.RO:    "Rumänie",
			territory.RS:    "Serbie",
			territory.RU:    "Russland",
			territory.RW:    "Ruanda",
			territory.SA:    "Saudi Arabie",
			territory.SB:    "Salomone",
			territory.SC:    "Sečelle",
			territory.SD:    "Sudan",
			territory.SE:    "Schwede",
			territory.SG:    "Singapur",
			territory.SH:    "St. Helena",
			territory.SI:    "Slowenie",
			territory.SJ:    "Svalbard und Jan Mayen",
			territory.SK:    "Slowakei",
			territory.SL:    "Sierra Leone",
			territory.SM:    "San Marino",
			territory.SN:    "Senegal",
			territory.SO:    "Somalia",
			territory.SR:    "Suriname",
			territory.ST:    "São Tomé and Príncipe",
			territory.SV:    "El Salvador",
			territory.SY:    "Sürie",
			territory.SZ:    "Swasiland",
			territory.TA:    "Tristan da Cunha",
			territory.TC:    "Turks- und Caicosinslä",
			territory.TD:    "Tšad",
			territory.TF:    "Französiši Süd- und Antarktisgebiet",
			territory.TG:    "Togo",
			territory.TH:    "Thailand",
			territory.TJ:    "Tadšikistan",
			territory.TK:    "Tokelau",
			territory.TL:    "Ošttimor",
			territory.TM:    "Turkmeništan",
			territory.TN:    "Tunesie",
			territory.TO:    "Tonga",
			territory.TR:    "Türkei",
			territory.TT:    "Trinidad und Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwan",
			territory.TZ:    "Tansania",
			territory.UA:    "Ukraine",
			territory.UG:    "Uganda",
			territory.UM:    "Amerikaniš Ozeanie",
			territory.US:    "Amerika",
			territory.UY:    "Urugauy",
			territory.UZ:    "Usbekištan",
			territory.VA:    "Vatikan",
			territory.VC:    "St. Vincent und d’Grenadine",
			territory.VE:    "Venezuela",
			territory.VG:    "Britiši Jungfröiwinslä",
			territory.VI:    "Amerikaniši Jungfröiwinslä",
			territory.VN:    "Vietnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Wallis und Futuna",
			territory.WS:    "Samoa",
			territory.YE:    "Jéme",
			territory.YT:    "Moyette",
			territory.ZA:    "Südafrika",
			territory.ZM:    "Sambia",
			territory.ZW:    "Simbabwe",
			territory.ZZ:    "Unbekannti Regio",
		},
	}
}
