// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_gaa() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "gaa",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss\u202fa zzzz", Long: "h:mm:ss\u202fa z", Medium: "h:mm:ss\u202fa", Short: "h:mm\u202fa"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'be' 'ni' 'atswa' {0}", Long: "{1} 'be' 'ni' 'atswa' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Aha", Feb: "Ofl", Mar: "Ots", Apr: "Abe", May: "Agb", Jun: "Otu", Jul: "Maa", Aug: "Man", Sep: "Gbo", Oct: "Ant", Nov: "Ale", Dec: "Afu"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "A", Feb: "O", Mar: "O", Apr: "A", May: "A", Jun: "O", Jul: "M", Aug: "M", Sep: "G", Oct: "A", Nov: "A", Dec: "A"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Aharabata", Feb: "Oflɔ", Mar: "Otsokrikri", Apr: "Abɛibe", May: "Agbiɛnaa", Jun: "Otukwajan", Jul: "Maawɛ", Aug: "Manyawale", Sep: "Gbo", Oct: "Antɔŋ", Nov: "Alemle", Dec: "Afuabe"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Hɔg", Mon: "Ju", Tue: "Juf", Wed: "Shɔ", Thu: "Soo", Fri: "Soh", Sat: "Hɔɔ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "H", Mon: "J", Tue: "J", Wed: "S", Thu: "S", Fri: "S", Sat: "H"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Hɔgbaa", Mon: "Ju", Tue: "Jufɔ", Wed: "Shɔ", Thu: "Soo", Fri: "Sohaa", Sat: "Hɔɔ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "LB", PM: "SN"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "LEEBI", PM: "SHWANE"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.ALL: cldr.Currency{DisplayName: "Albania Leki", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.ANG: cldr.Currency{DisplayName: "Netherlands Antillea Guilda", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Angola Kwanza", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Argentina Peso", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba Florin", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "Bosnia-Herzegovina Marki Ni Hiɔ Tsakemɔ", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados Dɔla", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Bulgaria Levi", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Burundi Franki", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda Dɔla", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Bolivia Boliviano", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Brazil Real", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Bahamas Dɔla", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Botswana Pula", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Belarus Rubol", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Belize Dɔla", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Kanada Dɔla", Symbol: "KA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo Franki", Symbol: "KDF"},
				currency.CHF: cldr.Currency{DisplayName: "Switzerland Frank", Symbol: "SZF"},
				currency.CLP: cldr.Currency{DisplayName: "Tsili Peso", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolombia Peso", Symbol: "KOP"},
				currency.CRC: cldr.Currency{DisplayName: "Kosta Rika Kolón", Symbol: "KRK"},
				currency.CUC: cldr.Currency{DisplayName: "Kuba Peso Ni Hiɔ Tsakemɔ", Symbol: "KUK"},
				currency.CUP: cldr.Currency{DisplayName: "Kuba Peso", Symbol: "KUP"},
				currency.CZK: cldr.Currency{DisplayName: "Tsek Koruna", Symbol: "TSK"},
				currency.DJF: cldr.Currency{DisplayName: "Djibouti Franki", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Denmark Krone", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Dominika Peso", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Algeria Dinar", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Ejipt Pound", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrea Nakfa", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ethiopia Birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Falkland Ŋshɔkpɔi Pound", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Britain Pound", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sidi", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ghana Sidi", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltar Pound", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Gambia Dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Guinea Franki", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemala Kuetzal", Symbol: "GTK"},
				currency.GYD: cldr.Currency{DisplayName: "Guyana Dɔla", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Hondura Lempira", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kroatia Kuna", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Haiti Gourde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Hungary Forinti", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "Aisland Króna", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaika Dɔla", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenya Sheleŋ", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Komoros Franki", Symbol: "KF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "Kayman Ŋshɔkpɔi Dɔla", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Liberia Dɔla", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Libia Dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Moroko Dirham", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Moldova Leu", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Madagaska Ariari", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Makedonia Denari", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRU: cldr.Currency{DisplayName: "Mauritania Ouguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mauritius Rupi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Malawi Kwatsa", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Meziko Peso", Symbol: "MZ$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Mozambik Metikal", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Namibia Dɔla", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Anago Naira", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Nikaragua Kórdoba", Symbol: "K$"},
				currency.NOK: cldr.Currency{DisplayName: "Norway Krone", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PAB: cldr.Currency{DisplayName: "Panama Balboa", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Peru Sol", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Poland Zloti", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Paraguay Guarani", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "Romania Leu", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Serbia Dinari", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Russia Rubol", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "Rwanda Franki", Symbol: "RF"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Seyshɛl Rupi", Symbol: "SSR"},
				currency.SDG: cldr.Currency{DisplayName: "Sudan Pound", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Sweden Krona", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "St. Helena Pound", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Sierra Leone Leone", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leone Leone (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Somali Sheleŋ", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Surinam Dɔla", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Sudan Anai Pound", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Swazi Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Tunisia Dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad Kɛ Tobago Dɔla", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzania Sheleŋ", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Ukrainia Hryvnia", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Uganda Sjeleŋ", Symbol: "UGS"},
				currency.USD: cldr.Currency{DisplayName: "US Dɔla", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Uruguay Peso", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Venezuela Bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Karibbean Bokã Dɔla", Symbol: "KB$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Afrika Anai Sefa Franki", Symbol: "SFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Shika Ko Ni Gbɛ́i Bɛ Mli", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "South Afrika Randi", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Zambia Kwatsa", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AB:      "Abkhazia",
			language.ACE:     "Achinese",
			language.ADA:     "Dangme",
			language.ADY:     "Adyghe",
			language.AF:      "Afrikaans",
			language.AGQ:     "Aghem",
			language.AIN:     "Ainu",
			language.AK:      "Akan",
			language.ALE:     "Aleut",
			language.AM:      "Amharic",
			language.AN:      "Aragonese",
			language.ANP:     "Angika",
			language.AR:      "Arabik",
			language.AR_001:  "Ŋmɛnɛŋmɛnɛ Beiaŋ Arabik",
			language.ARP:     "Arapaho",
			language.AS:      "Assamese",
			language.ASA:     "Asu",
			language.AST:     "Asturian",
			language.ATJ:     "Atikamekw",
			language.AV:      "Avaric",
			language.AWA:     "Awadhi",
			language.AY:      "Aymara",
			language.AZ:      "Azeri",
			language.BA:      "Bashkir",
			language.BAN:     "Balinese",
			language.BAS:     "Basaa",
			language.BE:      "Belarusian",
			language.BEM:     "Bemba",
			language.BEZ:     "Bena",
			language.BG:      "Bulgarian",
			language.BHO:     "Bhojpuri",
			language.BI:      "Bislama",
			language.BIN:     "Bini",
			language.BM:      "Bambara",
			language.BN:      "Bangla",
			language.BR:      "Breton",
			language.BRX:     "Bodo",
			language.BS:      "Bosnian",
			language.BUG:     "Buginese",
			language.BYN:     "Blin",
			language.CA:      "Katalan",
			language.CAY:     "Kayuga",
			language.CCP:     "Tsakma",
			language.CEB:     "Sebuano",
			language.CRR:     "Karolina Algonkian",
			language.DE:      "German",
			language.DE_AT:   "Austria German",
			language.DE_CH:   "Switzerland German Krɔŋŋ",
			language.EN:      "Blɔfo",
			language.EN_AU:   "Australia Blɔfo",
			language.EN_CA:   "Kanada Blɔfo",
			language.EN_GB:   "UK Blɔfo",
			language.EN_US:   "US Blɔfo",
			language.ES:      "Spanish",
			language.ES_419:  "Romanse Amerika Spanish",
			language.ES_ES:   "Yuropa Spanish",
			language.ES_MX:   "Meziko Spanish",
			language.EU:      "Baske",
			language.FR:      "Frɛntsi",
			language.FR_CA:   "Kanada Frɛntsi",
			language.FR_CH:   "Switzerland Frɛntsi",
			language.FRC:     "Kajun Frɛntsi",
			language.GAA:     "Gã",
			language.HI:      "Hindi",
			language.HY:      "Armenian",
			language.ID:      "Indonesian",
			language.IT:      "Italian",
			language.JA:      "Japanese",
			language.KO:      "Korean",
			language.KSF:     "Bafia",
			language.MY:      "Burmese",
			language.NL:      "Daatsi",
			language.NL_BE:   "Flemish",
			language.PL:      "Polish",
			language.PT:      "Portuguese",
			language.PT_BR:   "Brazil Portuguese",
			language.PT_PT:   "Yuropa Portuguese",
			language.RU:      "Russian",
			language.RUP:     "Aromanian",
			language.SQ:      "Albanian",
			language.TH:      "Thai",
			language.TR:      "Turkish",
			language.UND:     "Wiemɔ ko ni gbɛ́i bɛ mli",
			language.YUE:     "Tsainesi, Kantonese",
			language.ZH:      "Tsainese, Mandarin",
			language.ZH_HANS: "Mandarin Tsainese Ni Waaa",
			language.ZH_HANT: "Blema Mandarin Tsainese",
		},
		Territories: cldr.Territories{
			territory.V_001: "Jeŋ Fɛɛ",
			territory.V_002: "Afrika",
			territory.V_003: "Kooyigbɛ Amerika",
			territory.V_005: "Wuoyigbɛ Amerika",
			territory.V_009: "Ŋshɔkpɔi",
			territory.V_011: "Afrika Anaigbɛ",
			territory.V_013: "Teŋgbɛ Amerika",
			territory.V_014: "Afrika Bokagbɛ",
			territory.V_015: "Afrika Kooyigbɛ",
			territory.V_017: "Afrika Teŋgbɛ",
			territory.V_018: "Afrika Wuoyigbɛ",
			territory.V_019: "Amerika Niiaŋ",
			territory.V_021: "Kooyigbɛ Shɔŋŋ Amerika",
			territory.V_029: "Karibean",
			territory.V_030: "Asia Bokagbɛ",
			territory.V_034: "Asia Wuoyigbɛ",
			territory.V_035: "Asia Wuoyi-Bokagbɛ",
			territory.V_039: "Yuropa Wuoyigbɛ",
			territory.V_053: "Australasia",
			territory.V_054: "Melanesia",
			territory.V_057: "Ŋshɔkpɔi Bibii",
			territory.V_061: "Ŋshɔkpɔi Bibii Pii",
			territory.V_142: "Asia",
			territory.V_143: "Asia Teŋgbɛ",
			territory.V_145: "Asia Anaigbɛ",
			territory.V_150: "Yuropa",
			territory.V_151: "Yuropa Bokagbɛ",
			territory.V_154: "Yuropa Kooyigbɛ",
			territory.V_155: "Yuropa Anaigbɛ",
			territory.V_202: "Afrika Fã Ni Yɔɔ Sahara Lɛ Shishi",
			territory.V_419: "Romanse Amerika",
			territory.AG:    "Antigua Kɛ Barbuda",
			territory.AI:    "Anguilla",
			territory.AO:    "Angola",
			territory.AR:    "Argentina",
			territory.AW:    "Aruba",
			territory.BB:    "Barbados",
			territory.BF:    "Burkina Faso",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "St. Barthélemy",
			territory.BM:    "Bermuda",
			territory.BO:    "Bolivia",
			territory.BQ:    "Netherlands Ni Yɔɔ Karibean",
			territory.BR:    "Brazil",
			territory.BS:    "Bahamas",
			territory.BV:    "Bouvet Ŋshɔkpɔ",
			territory.BW:    "Botswana",
			territory.BZ:    "Belize",
			territory.CA:    "Kanada",
			territory.CD:    "Kongo - Kinshasa",
			territory.CF:    "Teŋgbɛ Afrika Jeŋmaŋ",
			territory.CG:    "Kongo - Brazzaville",
			territory.CI:    "Ko Divua",
			territory.CL:    "Tsili",
			territory.CM:    "Kameroon",
			territory.CN:    "Tsaina",
			territory.CO:    "Kolombia",
			territory.CR:    "Kosta Rika",
			territory.CU:    "Kuba",
			territory.CV:    "Kape Verde",
			territory.CW:    "Kurasao",
			territory.DJ:    "Djibouti",
			territory.DM:    "Dominika",
			territory.DO:    "Dominika Republik",
			territory.DZ:    "Algeria",
			territory.EA:    "Keuta Kɛ Melilla",
			territory.EC:    "Ekuador",
			territory.EG:    "Ejipt",
			territory.EH:    "Sahara Wuoyigbɛ",
			territory.ER:    "Eritrea",
			territory.ET:    "Etiopia",
			territory.EU:    "Yuropa Maji Ekomefeemɔ",
			territory.EZ:    "Yuropaniiaŋ",
			territory.FK:    "Falkland Ŋshɔkpɔi",
			territory.GA:    "Gabon",
			territory.GD:    "Grenada",
			territory.GF:    "Frentsibii Guiana",
			territory.GH:    "Ghana",
			territory.GL:    "Greenland",
			territory.GM:    "Gambia",
			territory.GN:    "Guinea",
			territory.GP:    "Guadeloupe",
			territory.GQ:    "Ekuatorial Guinea",
			territory.GS:    "Georgia Wuoyi Kɛ Sandwitsi Ŋshɔkpɔi Ni Yɔɔ Wuoyi",
			territory.GT:    "Guatemala",
			territory.GW:    "Guinea-Bissau",
			territory.GY:    "Guyana",
			territory.HN:    "Honduras",
			territory.HT:    "Haiti",
			territory.IC:    "Kanary Ŋshɔkpɔi",
			territory.IN:    "India",
			territory.IO:    "Britain Shikpɔji Ni Yɔɔ Indian Ŋshɔ Lɛ Mli",
			territory.JM:    "Jamaika",
			territory.JP:    "Japan",
			territory.KE:    "Kenya",
			territory.KM:    "Komoros",
			territory.KN:    "St. Kitts Kɛ Nevis",
			territory.KY:    "Kayman Ŋshɔkpɔi",
			territory.LC:    "St. Lusia",
			territory.LR:    "Liberia",
			territory.LS:    "Lesotho",
			territory.LY:    "Libia",
			territory.MA:    "Moroko",
			territory.MF:    "St. Martin",
			territory.MG:    "Madagaskar",
			territory.ML:    "Mali",
			territory.MO:    "Makao SAR Tsaina",
			territory.MQ:    "Martinik",
			territory.MR:    "Mauritania",
			territory.MS:    "Montserrat",
			territory.MU:    "Mauritius",
			territory.MW:    "Malawi",
			territory.MX:    "Meziko",
			territory.MZ:    "Mozambik",
			territory.NA:    "Namibia",
			territory.NE:    "Niger",
			territory.NG:    "Anago",
			territory.NI:    "Nikaragua",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PM:    "St. Pierre Kɛ Mikelon",
			territory.PR:    "Puerto Riko",
			territory.PY:    "Paraguay",
			territory.QO:    "Ŋshɔkpɔi Ni Yɔɔ Shɔŋŋ",
			territory.RE:    "Réunion",
			territory.RW:    "Rwanda",
			territory.SC:    "Seyshelles",
			territory.SD:    "Sudan",
			territory.SH:    "St. Helena",
			territory.SL:    "Sierra Leone",
			territory.SN:    "Senegal",
			territory.SO:    "Somalia",
			territory.SR:    "Suriname",
			territory.SS:    "Sudan Wuoyi",
			territory.ST:    "São Tomé Kɛ Prínsipe",
			territory.SV:    "El Salvador",
			territory.SX:    "Sint Maarten",
			territory.SZ:    "Eswatini",
			territory.TC:    "Turks Kɛ Kaikos Ŋshɔkpɔi",
			territory.TD:    "Tsad",
			territory.TF:    "Frentsibii Ashikpɔji Ni Yɔɔ Wuoyi",
			territory.TG:    "Togo",
			territory.TN:    "Tunisia",
			territory.TT:    "Trinidad Kɛ Tobago",
			territory.TZ:    "Tanzania",
			territory.UG:    "Uganda",
			territory.UN:    "Jeŋmaji Ekomefeemɔ",
			territory.US:    "United States",
			territory.UY:    "Uruguay",
			territory.VC:    "St. Vinsent Kɛ Grenadines",
			territory.VE:    "Venezuela",
			territory.VG:    "Britain Ŋshɔkpɔi Ni Atarako Amɛhe",
			territory.VI:    "US Ŋshɔkpɔi Ni Atarako Amɛhe",
			territory.XA:    "Eyaa Ŋwɛi Kɛ Shikpɔŋ Fɛɛ",
			territory.XB:    "Eyaa Biɛ Kɛ Biɛ Fɛɛ",
			territory.YT:    "Mayotte",
			territory.ZA:    "South Afrika",
			territory.ZM:    "Zambia",
			territory.ZW:    "Zimbabwe",
			territory.ZZ:    "He Ko Ni Gbɛ́i Bɛ Mli",
		},
	}
}
