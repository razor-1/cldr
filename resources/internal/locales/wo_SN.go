// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_wo_SN() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "wo_SN",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "dd-MM-y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} - {0}", Long: "{1} - {0}", Medium: "{1} - {0}", Short: "{1} - {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Sam", Feb: "Few", Mar: "Mar", Apr: "Awr", May: "Mee", Jun: "Suw", Jul: "Sul", Aug: "Ut", Sep: "Sàt", Oct: "Okt", Nov: "Now", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Samwiyee", Feb: "Fewriyee", Mar: "Mars", Apr: "Awril", May: "Mee", Jun: "Suwe", Jul: "Sulet", Aug: "Ut", Sep: "Sàttumbar", Oct: "Oktoobar", Nov: "Nowàmbar", Dec: "Desàmbar"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dib", Mon: "Alt", Tue: "Tal", Wed: "Àla", Thu: "Alx", Fri: "Àjj", Sat: "Ase"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Dib", Mon: "Alt", Tue: "Tal", Wed: "Àla", Thu: "Alx", Fri: "Àjj", Sat: "Ase"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Dibéer", Mon: "Altine", Tue: "Talaata", Wed: "Àlarba", Thu: "Alxamis", Fri: "Àjjuma", Sat: "Aseer"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Sub", PM: "Ngo"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤0K", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "United Arab Emirates Dirham", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghan Afghani", Symbol: "؋"},
				currency.ALL: cldr.Currency{DisplayName: "Albanian Lek", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Armenian Dram", Symbol: "֏"},
				currency.ANG: cldr.Currency{DisplayName: "Netherlands Antillean Guilder", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Angolan Kwanza", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Argentine Peso", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Australian Dollar", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruban Florin", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Azerbaijani Manat", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "Bosnia-Herzegovina Convertible Mark", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbadian Dollar", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladeshi Taka", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Bulgarian Lev", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bahraini Dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Burundian Franc", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Vote BMD", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Brunei Dollar", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Bolivian Boliviano", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Real bu Bresil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Bahamian Dollar", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Bhutanese Ngultrum", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Botswanan Pula", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Belarusian Ruble", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Belize Dollar", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Vote CAD", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Congolese Franc", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Swiss Franc", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Chilean Peso", Symbol: "Vote $"},
				currency.CNH: cldr.Currency{DisplayName: "Chinese Yuan (offshore)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Yuan bu Siin", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Colombian Peso", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Costa Rican Colón", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Cuban Convertible Peso", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Cuban Peso", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Cape Verdean Escudo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Czech Koruna", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Djiboutian Franc", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Danish Krone", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Dominican Peso", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Algerian Dinar", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Egyptian Pound", Symbol: "EGPP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrean Nakfa", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ethiopian Birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Fijian Dollar", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "FKPS", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pound bu Grànd Brëtaañ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Georgian Lari", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "Ghanaian Cedi", Symbol: "GHS."},
				currency.GIP: cldr.Currency{DisplayName: "Vote GIP", Symbol: "GIIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambian Dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Guinean Franc", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "GT", Symbol: "GT Q"},
				currency.GYD: cldr.Currency{DisplayName: "Guyanaese Dollar", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Hong Kong Dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Honduran Lempira", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Croatian Kuna", Symbol: "HRKS"},
				currency.HTG: cldr.Currency{DisplayName: "Haitian Gourde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Hungarian Forint", Symbol: "Vote Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesian Rupiah", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Israeli New Shekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupee bu End", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Iraqi Dinar", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Iranian Rial", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Icelandic Króna", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaican Dollar", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Jordanian Dinar", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Yen bu Sapoŋ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenyan Shilling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Kyrgystani Som", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "Cambodian Riel", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Comorian Franc", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "North Korean Won", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "South Korean Won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuwaiti Dinar", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Cayman Islands Dollar", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Kazakhstani Tenge", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Laotian Kip", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Lebanese Pound", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lankan Rupee", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Liberian Dollar", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lesotho Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Libyan Dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Moroccan dirhams", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Moldovan Leu", Symbol: "Vote MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Malagasy Ariary", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Macedonian Denar", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Myanmar Kyat", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Mongolian Tugrik", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Macanese Pataca", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Mauritanian Ouguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mauritian Rupee", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Maldivian Rufiyaa", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Malawian Kwacha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Mexican Peso", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Malaysian Ringgit", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Mozambican Metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Namibian Dollar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nigerian Naira", Symbol: "NGN."},
				currency.NIO: cldr.Currency{DisplayName: "Nicaraguan Córdoba", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Norwegian Krone", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Nepalese Rupee", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "New Zealand Dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Omani Rial", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Panamanian Balboa", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Peruvian Sols", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Papua New Guinean Kina", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Philippine Peso", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistani Rupee", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Polish Zloty", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Paraguayan Guaranis", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Qatari Riyal", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Romanian Leu", Symbol: "Vote lei"},
				currency.RSD: cldr.Currency{DisplayName: "Serbian Dinar", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Ruble bi Rsis", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "Rwandan Franc", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi Riyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Solomon Islands Dollar", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Seychellois Rupee", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Sudanese Pound", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Swedish Krona", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Singapore Dollar", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "St. Helena Pound", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Sierra Leonean Leone", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leonean Leone (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Somali Shilling", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Surinamese Dollar", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "South Sudanese Pound", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "São Tomé & Príncipe Dobra", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Syrian Pound", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Swazi Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Thai Baht", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Tajikistani Somoni", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Turkmenistani Manat", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Tunisian Dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Tongan Paʻanga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Turkish Lira", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad & Tobago Dollar", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "New Taiwan Dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzanian Shilling", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "UAHS", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Ugandan Shilling", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dolaaru US", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Uruguayan Peso", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Uzbekistani Som", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Venezuelan Bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Vietnamese Dong", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu Vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Samoan Tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Central African CFA Franc", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "East Caribbean Dollar", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Franc CFA bu Afrik Sowwu-jant", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP Franc", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Xaalis buñ Xamul", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Yemeni Rial", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "South African Rand", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Zambian Kwacha", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AF:      "Afrikaans",
			language.AM:      "Amharik",
			language.AR:      "Arabic",
			language.AR_001:  "Araab",
			language.AS:      "Asame",
			language.AZ:      "Aserbayjane",
			language.BA:      "Baskir",
			language.BAN:     "Bali",
			language.BE:      "Belaris",
			language.BEM:     "Bemba",
			language.BG:      "Bilgaar",
			language.BN:      "Baŋla",
			language.BO:      "Tibetan",
			language.BR:      "Breton",
			language.BS:      "Bosñak",
			language.CA:      "Katalan",
			language.CEB:     "Sibiyanoo",
			language.CHM:     "Mari",
			language.CHR:     "Ceroki",
			language.CKB:     "Kurdi gu Diggu",
			language.CO:      "Kors",
			language.CS:      "Cek",
			language.CY:      "Wels",
			language.DA:      "Danuwa",
			language.DE:      "Almaa",
			language.DE_AT:   "Almaa bu Ótiriis",
			language.DE_CH:   "Almaa bu Kawe bu Swis",
			language.DSB:     "Sorab-Suuf",
			language.DV:      "Diweyi",
			language.DZ:      "Dsongkaa",
			language.EL:      "Gereg",
			language.EN:      "Àngale",
			language.EN_AU:   "Àngale bu Óstraali",
			language.EN_CA:   "Àngale bu Kanadaa",
			language.EN_GB:   "Àngale (RI)",
			language.EN_US:   "Àngale (ES)",
			language.EO:      "Esperantoo",
			language.ES:      "Español",
			language.ES_419:  "Español bu Amerik Latin",
			language.ES_ES:   "Español bu Tugël",
			language.ES_MX:   "Español bu Meksik",
			language.ET:      "Estoñiye",
			language.EU:      "Bask",
			language.FA:      "Pers",
			language.FF:      "Pël",
			language.FI:      "Feylànde",
			language.FIL:     "Filipiye",
			language.FO:      "Feroos",
			language.FR:      "Farañse",
			language.FR_CA:   "Frañse bu Kanadaa",
			language.FR_CH:   "Frañse bu Swis",
			language.GA:      "Irlànde",
			language.GD:      "Galuwaa bu Ekos",
			language.GL:      "Galisiye",
			language.GN:      "Garani",
			language.GU:      "Gujarati",
			language.HA:      "Hawsa",
			language.HAW:     "Hawaye",
			language.HE:      "Ebrë",
			language.HI:      "Endo",
			language.HI_LATN: "Hindī bu Àngale",
			language.HIL:     "Hiligaynon",
			language.HR:      "Krowat",
			language.HSB:     "Sorab-Kaw",
			language.HT:      "Kereyolu Ayti",
			language.HU:      "Ongruwaa",
			language.HY:      "Armaniye",
			language.HZ:      "Herero",
			language.IBB:     "Ibibiyo",
			language.ID:      "Endonesiye",
			language.IG:      "Igbo",
			language.IS:      "Islànde",
			language.IT:      "Italiye",
			language.IU:      "Inuktitit",
			language.JA:      "Sapone",
			language.KA:      "Sorsiye",
			language.KK:      "Kasax",
			language.KM:      "Xmer",
			language.KN:      "Kannadaa",
			language.KO:      "Koreye",
			language.KOK:     "Konkani",
			language.KR:      "Kanuri",
			language.KRU:     "Kuruks",
			language.KS:      "Kashmiri",
			language.KU:      "Kurdi",
			language.KY:      "Kirgiis",
			language.LA:      "Latin",
			language.LB:      "Liksàmbursuwaa",
			language.LO:      "Laaw",
			language.LT:      "Lituyaniye",
			language.LV:      "Letoniye",
			language.MEN:     "Mende",
			language.MG:      "Malagasi",
			language.MI:      "Mawri",
			language.MK:      "Maseduwaane",
			language.ML:      "Malayalam",
			language.MN:      "Mongoliye",
			language.MNI:     "Manipuri",
			language.MOH:     "Mowak",
			language.MR:      "Marati",
			language.MS:      "Malay",
			language.MT:      "Malt",
			language.MY:      "Birmes",
			language.NE:      "Nepale",
			language.NIU:     "Niweyan",
			language.NL:      "Neyerlànde",
			language.NL_BE:   "Belsig",
			language.NO:      "Nerwesiye",
			language.NY:      "Sewa",
			language.OC:      "Ositan",
			language.OM:      "Oromo",
			language.OR:      "Oja",
			language.PA:      "Punjabi",
			language.PAP:     "Papiyamento",
			language.PL:      "Polone",
			language.PS:      "Pasto",
			language.PT:      "Purtugees",
			language.PT_BR:   "Purtugees bu Bresil",
			language.PT_PT:   "Portugees bu Tugël",
			language.QU:      "Kesuwa",
			language.QUC:     "Kishe",
			language.RM:      "Romaas",
			language.RO:      "Rumaniyee",
			language.RU:      "Rus",
			language.RW:      "Kinyarwànda",
			language.SA:      "Sanskrit",
			language.SAH:     "Saxa",
			language.SAT:     "Santali",
			language.SD:      "Sindi",
			language.SE:      "Penku Sami",
			language.SI:      "Sinala",
			language.SK:      "Eslowaki (Eslowak)",
			language.SL:      "Esloweniye",
			language.SMA:     "Sami gu Saalum",
			language.SMJ:     "Lule Sami",
			language.SMN:     "Inari Sami",
			language.SMS:     "Eskolt Sami",
			language.SO:      "Somali (làkk)",
			language.SQ:      "Albane",
			language.SR:      "Serb",
			language.SV:      "Suweduwaa",
			language.SYR:     "Siryak",
			language.TA:      "Tamil",
			language.TE:      "Telugu",
			language.TG:      "Tajis",
			language.TH:      "Tay",
			language.TI:      "Tigriña",
			language.TK:      "Tirkmen",
			language.TO:      "Tongan",
			language.TR:      "Tirk",
			language.TT:      "Tatar",
			language.TZM:     "Tamasis gu Digg Atlaas",
			language.UG:      "Uygur",
			language.UK:      "Ikreniye",
			language.UND:     "Làkk wuñ xamul",
			language.UR:      "Urdu",
			language.UZ:      "Usbek",
			language.VE:      "Wenda",
			language.VI:      "Wiyetnaamiye",
			language.WO:      "Wolof",
			language.YI:      "Yidis",
			language.YO:      "Yoruba",
			language.ZH:      "Sinuwaa",
			language.ZH_HANS: "Sinuwaa buñ woyofal",
			language.ZH_HANT: "Sinuwaa bu cosaan",
		},
		Territories: cldr.Territories{
			territory.V_001: "àddina",
			territory.V_002: "Africa",
			territory.V_003: "North America",
			territory.V_005: "Amerique du Sud",
			territory.V_009: "Oseani",
			territory.V_011: "Sowwu Afrique",
			territory.V_013: "Amerique Centrale",
			territory.V_014: "Penku Afrique",
			territory.V_015: "Afrique du Nord",
			territory.V_017: "Moyen Afrique",
			territory.V_018: "Afrique du Sud",
			territory.V_019: "Amerika",
			territory.V_021: "amerique du nord",
			territory.V_029: "Caraïbe",
			territory.V_030: "Asie penku",
			territory.V_034: "Asie du Sud",
			territory.V_035: "Asie Sud-est",
			territory.V_039: "Sud Europe",
			territory.V_053: "Ostralasi",
			territory.V_054: "Melanesi",
			territory.V_057: "Mikronesi",
			territory.V_061: "Polineesi",
			territory.V_142: "Asia",
			territory.V_143: "Asie centrale",
			territory.V_145: "Asie sowwu jàng",
			territory.V_150: "Europe",
			territory.V_151: "Europe bu penku",
			territory.V_154: "Europe du nord",
			territory.V_155: "Europe sowwu jàng",
			territory.V_202: "Afrique sub-saharienne",
			territory.V_419: "Amerique Latine",
			territory.AC:    "Ile Ascension",
			territory.AD:    "Andoor",
			territory.AE:    "Emira Arab Ini",
			territory.AF:    "Afganistaŋ",
			territory.AG:    "Antiguwa ak Barbuda",
			territory.AI:    "Angiiy",
			territory.AL:    "Albani",
			territory.AM:    "Armeni",
			territory.AO:    "Àngolaa",
			territory.AQ:    "Antarktik",
			territory.AR:    "Arsàntin",
			territory.AS:    "Samowa bu Amerig",
			territory.AT:    "Ótiriis",
			territory.AU:    "Ostarali",
			territory.AW:    "Aruba",
			territory.AX:    "Duni Aalànd",
			territory.AZ:    "Aserbayjaŋ",
			territory.BA:    "Bosni Ersegowin",
			territory.BB:    "Barbad",
			territory.BD:    "Bengalades",
			territory.BE:    "Belsig",
			territory.BF:    "Burkina Faaso",
			territory.BG:    "Bilgari",
			territory.BH:    "Bahreyin",
			territory.BI:    "Burundi",
			territory.BJ:    "Benee",
			territory.BL:    "Saŋ Bartalemi",
			territory.BM:    "Bermid",
			territory.BN:    "Burney",
			territory.BO:    "Boliwi",
			territory.BQ:    "Pays-Bas bu Caraïbe",
			territory.BR:    "Beresil",
			territory.BS:    "Bahamas",
			territory.BT:    "Butaŋ",
			territory.BV:    "Dunu Buwet",
			territory.BW:    "Botswana",
			territory.BY:    "Belaris",
			territory.BZ:    "Belis",
			territory.CA:    "Kanadaa",
			territory.CC:    "Duni Koko (Kilin)",
			territory.CD:    "Kongo (R K D)",
			territory.CF:    "Repiblik Sàntar Afrik",
			territory.CG:    "Réewum Kongo",
			territory.CH:    "Siwis",
			territory.CI:    "Kodiwaar",
			territory.CK:    "Duni Kuuk",
			territory.CL:    "Sili",
			territory.CM:    "Kamerun",
			territory.CN:    "Siin",
			territory.CO:    "Kolombi",
			territory.CP:    "Ile Clipperton",
			territory.CR:    "Kosta Rika",
			territory.CU:    "Kuba",
			territory.CV:    "Kabo Werde",
			territory.CW:    "Kursawo",
			territory.CX:    "Dunu Kirismas",
			territory.CY:    "Siipar",
			territory.CZ:    "Réewum Cek",
			territory.DE:    "Almaañ",
			territory.DG:    "Garsiya",
			territory.DJ:    "Jibuti",
			territory.DK:    "Danmàrk",
			territory.DM:    "Dominik",
			territory.DO:    "Repiblik Dominiken",
			territory.DZ:    "Alseri",
			territory.EA:    "Ceuta & Melilla",
			territory.EC:    "Ekwaatër",
			territory.EE:    "Estoni",
			territory.EG:    "Esipt",
			territory.EH:    "Sahara bu sowwu",
			territory.ER:    "Eritere",
			territory.ES:    "Españ",
			territory.ET:    "Ecopi",
			territory.EU:    "EZ",
			territory.EZ:    "Eurozone",
			territory.FI:    "Finlànd",
			territory.FJ:    "Fijji",
			territory.FK:    "Duni Falkland",
			territory.FM:    "Mikoronesi",
			territory.FO:    "Duni Faro",
			territory.FR:    "Faraans",
			territory.GA:    "Gaboŋ",
			territory.GB:    "Ruwaayom Ini",
			territory.GD:    "Garanad",
			territory.GE:    "Seworsi",
			territory.GF:    "Guyaan Farañse",
			territory.GG:    "Gernase",
			territory.GH:    "Gana",
			territory.GI:    "Sibraltaar",
			territory.GL:    "Girinlànd",
			territory.GM:    "Gàmbi",
			territory.GN:    "Gine",
			territory.GP:    "Guwaadelup",
			territory.GQ:    "Gine Ekuwatoriyal",
			territory.GR:    "Gerees",
			territory.GS:    "Seworsi di Sid ak Duni Sàndwiis di Sid",
			territory.GT:    "Guwatemala",
			territory.GU:    "Guwam",
			territory.GW:    "Gine-Bisaawóo",
			territory.GY:    "Giyaan",
			territory.HK:    "Ooŋ Koŋ",
			territory.HM:    "Duni Hërd ak Duni MakDonald",
			territory.HN:    "Onduraas",
			territory.HR:    "Korowasi",
			territory.HT:    "Ayti",
			territory.HU:    "Ongari",
			territory.IC:    "Ile Canary",
			territory.ID:    "Indonesi",
			territory.IE:    "Irlànd",
			territory.IL:    "Israyel",
			territory.IM:    "Dunu Maan",
			territory.IN:    "End",
			territory.IO:    "Terituwaaru Brëtaañ ci Oseyaa Enjeŋ",
			territory.IQ:    "Irag",
			territory.IR:    "Iraŋ",
			territory.IS:    "Islànd",
			territory.IT:    "Itali",
			territory.JE:    "Serse",
			territory.JM:    "Samayig",
			territory.JO:    "Sordani",
			territory.JP:    "Sàppoŋ",
			territory.KE:    "Keeña",
			territory.KG:    "Kirgistaŋ",
			territory.KH:    "Kàmboj",
			territory.KI:    "Kiribati",
			territory.KM:    "Komoor",
			territory.KN:    "Saŋ Kits ak Newis",
			territory.KP:    "Kore Noor",
			territory.KR:    "Corée du Sud",
			territory.KW:    "Kowet",
			territory.KY:    "Duni Kaymaŋ",
			territory.KZ:    "Kasaxstaŋ",
			territory.LA:    "Lawos",
			territory.LB:    "Libaa",
			territory.LC:    "Saŋ Lusi",
			territory.LI:    "Liktensteyin",
			territory.LK:    "Siri Lànka",
			territory.LR:    "Liberiya",
			territory.LS:    "Lesoto",
			territory.LT:    "Litiyani",
			territory.LU:    "Liksàmbur",
			territory.LV:    "Letoni",
			territory.LY:    "Libi",
			territory.MA:    "Marog",
			territory.MC:    "Monako",
			territory.MD:    "Moldawi",
			territory.ME:    "Montenegoro",
			territory.MF:    "Saŋ Marteŋ",
			territory.MG:    "Madagaskaar",
			territory.MH:    "Duni Marsaal",
			territory.MK:    "Maseduwaan bëj Gànnaar",
			territory.ML:    "Mali",
			territory.MM:    "Miyanmaar",
			territory.MN:    "Mongoli",
			territory.MO:    "Makaawo",
			territory.MP:    "Duni Mariyaan Noor",
			territory.MQ:    "Martinik",
			territory.MR:    "Mooritani",
			territory.MS:    "Mooseraa",
			territory.MT:    "Malt",
			territory.MU:    "Moriis",
			territory.MV:    "Maldiiw",
			territory.MW:    "Malawi",
			territory.MX:    "Meksiko",
			territory.MY:    "Malesi",
			territory.MZ:    "Mosàmbig",
			territory.NA:    "Namibi",
			territory.NC:    "Nuwel Kaledoni",
			territory.NE:    "Niiseer",
			territory.NF:    "Dunu Norfolk",
			territory.NG:    "Niseriya",
			territory.NI:    "Nikaraguwa",
			territory.NL:    "Peyi Baa",
			territory.NO:    "Norwees",
			territory.NP:    "Nepaal",
			territory.NR:    "Nawru",
			territory.NU:    "Niw",
			territory.NZ:    "Nuwel Selànd",
			territory.OM:    "Omaan",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Polinesi Farañse",
			territory.PG:    "Papuwasi Gine Gu Bees",
			territory.PH:    "Filipin",
			territory.PK:    "Pakistaŋ",
			territory.PL:    "Poloñ",
			territory.PM:    "Saŋ Peer ak Mikeloŋ",
			territory.PN:    "Duni Pitkayirn",
			territory.PR:    "Porto Riko",
			territory.PS:    "réew yu Palestine",
			territory.PT:    "Portigaal",
			territory.PW:    "Palaw",
			territory.PY:    "Paraguwe",
			territory.QA:    "Kataar",
			territory.QO:    "Oceanie",
			territory.RE:    "Reeñoo",
			territory.RO:    "Rumani",
			territory.RS:    "Serbi",
			territory.RU:    "Risi",
			territory.RW:    "Ruwànda",
			territory.SA:    "Arabi Sawudi",
			territory.SB:    "Duni Salmoon",
			territory.SC:    "Seysel",
			territory.SD:    "Sudaŋ",
			territory.SE:    "Suwed",
			territory.SG:    "Singapuur",
			territory.SH:    "Saŋ Eleen",
			territory.SI:    "Esloweni",
			territory.SJ:    "Swalbaar ak Jan Mayen",
			territory.SK:    "Eslowaki",
			territory.SL:    "Siyera Lewon",
			territory.SM:    "San Marino",
			territory.SN:    "Senegaal",
			territory.SO:    "Somali",
			territory.SR:    "Sirinam",
			territory.SS:    "Sudaŋ di Sid",
			territory.ST:    "Sawo Tome ak Pirinsipe",
			territory.SV:    "El Salwadoor",
			territory.SX:    "Sin Marten",
			territory.SY:    "Siri",
			territory.SZ:    "Suwasilànd",
			territory.TA:    "Tristan da Cunha",
			territory.TC:    "Duni Tirk ak Kaykos",
			territory.TD:    "Càdd",
			territory.TF:    "Teer Ostraal gu Fraas",
			territory.TG:    "Togo",
			territory.TH:    "Taylànd",
			territory.TJ:    "Tajikistaŋ",
			territory.TK:    "Tokoloo",
			territory.TL:    "Timor Leste",
			territory.TM:    "Tirkmenistaŋ",
			territory.TN:    "Tinisi",
			territory.TO:    "Tonga",
			territory.TR:    "Tirki",
			territory.TT:    "Tirinite ak Tobago",
			territory.TV:    "Tuwalo",
			territory.TW:    "Taywan",
			territory.TZ:    "Taŋsani",
			territory.UA:    "Ikeren",
			territory.UG:    "Ugànda",
			territory.UM:    "Duni Amerig Utar meer",
			territory.UN:    "United Nations",
			territory.US:    "Etaa Sini",
			territory.UY:    "Uruge",
			territory.UZ:    "Usbekistaŋ",
			territory.VA:    "Site bu Watikaa",
			territory.VC:    "Saŋ Weesaa ak Garanadin",
			territory.VE:    "Wenesiyela",
			territory.VG:    "Duni Wirsin yu Brëtaañ",
			territory.VI:    "Duni Wirsin yu Etaa-sini",
			territory.VN:    "Wiyetnam",
			territory.VU:    "Wanuatu",
			territory.WF:    "Walis ak Futuna",
			territory.WS:    "Samowa",
			territory.XA:    "Pseudo-aksan",
			territory.XB:    "Pseudo-bidi",
			territory.XK:    "Kosowo",
			territory.YE:    "Yaman",
			territory.YT:    "Mayot",
			territory.ZA:    "Afrik di Sid",
			territory.ZM:    "Sàmbi",
			territory.ZW:    "Simbabwe",
			territory.ZZ:    "Gox buñ xamul",
		},
	}
}
