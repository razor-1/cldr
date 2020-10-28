// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ln_CG() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ln_CG",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "yan", Feb: "fbl", Mar: "msi", Apr: "apl", May: "mai", Jun: "yun", Jul: "yul", Aug: "agt", Sep: "stb", Oct: "ɔtb", Nov: "nvb", Dec: "dsb"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "y", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "y", Jul: "y", Aug: "a", Sep: "s", Oct: "ɔ", Nov: "n", Dec: "d"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "sánzá ya yambo", Feb: "sánzá ya míbalé", Mar: "sánzá ya mísáto", Apr: "sánzá ya mínei", May: "sánzá ya mítáno", Jun: "sánzá ya motóbá", Jul: "sánzá ya nsambo", Aug: "sánzá ya mwambe", Sep: "sánzá ya libwa", Oct: "sánzá ya zómi", Nov: "sánzá ya zómi na mɔ̌kɔ́", Dec: "sánzá ya zómi na míbalé"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "eye", Mon: "ybo", Tue: "mbl", Wed: "mst", Thu: "min", Fri: "mtn", Sat: "mps"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "e", Mon: "y", Tue: "m", Wed: "m", Thu: "m", Fri: "m", Sat: "p"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "eyenga", Mon: "mokɔlɔ mwa yambo", Tue: "mokɔlɔ mwa míbalé", Wed: "mokɔlɔ mwa mísáto", Thu: "mokɔlɔ ya mínéi", Fri: "mokɔlɔ ya mítáno", Sat: "mpɔ́sɔ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ntɔ́ngɔ́", PM: "mpókwa"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ntɔ́ngɔ́", PM: "mpókwa"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirihamɛ ya Lémila alabo", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Lek", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza ya Angóla", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Peso y’Argentina", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dolarɛ ya Ositali", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Guldeni y’ Aruba", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Mark ya kobóngwama", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Dolále ya Barbados", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Lev ya Bulgaria", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dinarɛ ya Bahrɛnɛ", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Falánga ya Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Real ya Brazil", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Dolále ya Bahamas", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pula ya Botswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Rubelé ya Bielorusí", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "Rubelé ya Bielorusí (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Dolále ya Belíze", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Dolarɛ ya Kanadá", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Falánga ya Kongó", Symbol: "FC"},
				currency.CHF: cldr.Currency{DisplayName: "Falánga ya Swisɛ", Symbol: "Fr."},
				currency.CLP: cldr.Currency{DisplayName: "Peso ya Shili", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuanɛ Renminbi ya Sinɛ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso ya Kolombi", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Colon ya Kosta Rika", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Peso ya Kuba", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Esikudo ya Kapevɛrɛ", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Motolé Sheki", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Falánga ya Dzibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Motolé ya Danemark", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Dominikani", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinarɛ ya Alizeri", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Paunɛ ya Ezípitɛ", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa ya Elitlɛ", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birɛ ya Etsiópi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Ɛlɔ́", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dolále ya Fiji", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Paunɛ ya Angɛlɛtɛ́lɛ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi ya Gana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Cedi", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Bojito ya Gibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ya Gambi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Falánga ya Gine", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Falánga ya Ginɛ", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna ya Kroasia", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gurde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Folinte", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupi ya Índɛ", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "Motolé ya Islandi", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Dolále ya Jamaïke", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yeni ya Zapɔ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilingɛ ya Kenya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Falánga ya Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "Dolále ya Bisanga bya Kayman", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dolarɛ ya Liberya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ya Lesóto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litas ya Litwani", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "Lats ya Letoni", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinarɛ ya Libí", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirihame ya Marokɛ", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Falánga ya Madagasikarɛ", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Denalé", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugwiya ya Moritani (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugwiya ya Moritani", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupi ya Morisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwasha ya Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Peso ya Mexiko", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikali ya Mozambiki", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Métikal", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dolarɛ ya Namibi", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira ya Nizerya", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Motolé ya Norvej", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Dolále ya Zeland ya Sika", Symbol: "NZ$"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sol Sika", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Sloty", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "Leu Sika", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Dinalé ya Serbia", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rubelé ya Rusí", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Falánga ya Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyalɛ ya Alabi Sawuditɛ", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Dolále ya Bisanga Solomoni", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupi ya Sɛshɛlɛ", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dinarɛ ya Sudá", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Paunɛ ya Sudá", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Motolé ya Swédi", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Paunɛ ya Sántu elena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leonɛ", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilingɛ ya Somali", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Bojito ya Sudaní ya Súdi", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra ya Sao Tomé mpé Presipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra ya Sao Tomé mpé Presipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinarɛ ya Tinizi", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Pa’Anga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Dolále ya Trinidad mpé Tobago", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingɛ ya Tanzani", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Griwná", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingɛ ya Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dolarɛ ya Ameriki", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Falánga CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dolále ya Kalibí Monyɛlɛ", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Falánga CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Falánga CFP", Symbol: "F CFP"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Randɛ ya Afríka Súdi", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwasha ya Zambi (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwasha ya Zambi", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dolarɛ ya Zimbabwɛ", Symbol: ""},
			},
		},
		Languages: cldr.Languages{
			language.AF:     "afrikansi",
			language.AK:     "akan",
			language.AM:     "liamariki",
			language.AR:     "lialabo",
			language.BE:     "libyelorisí",
			language.BG:     "libiligali",
			language.BN:     "libengali",
			language.CS:     "litshekɛ",
			language.DE:     "lialemá",
			language.DE_AT:  "lialémani ya Otrish",
			language.DE_CH:  "lialémani ya Swisi",
			language.EL:     "ligeleki",
			language.EN:     "lingɛlɛ́sa",
			language.EN_CA:  "lingɛlɛ́sa ya Kanadá",
			language.EN_GB:  "lingɛlɛ́sa ya Ingɛlɛ́tɛlɛ",
			language.ES:     "lisipanye",
			language.ES_419: "lispanyoli ya Ameríka Latína",
			language.ES_ES:  "lispanyoli ya Erópa",
			language.FA:     "lipelésanɛ",
			language.FR:     "lifalansɛ́",
			language.FR_CA:  "lifalansɛ́ ya Kanadá",
			language.FR_CH:  "lifalansɛ́ ya Swisi",
			language.GSW:    "lialemaniki",
			language.HA:     "hausa",
			language.HE:     "liébeleo",
			language.HI:     "lihindi",
			language.HU:     "liongili",
			language.ID:     "lindonezi",
			language.IG:     "igbo",
			language.IT:     "litaliano",
			language.JA:     "lizapɔ",
			language.JV:     "lizava",
			language.KG:     "kikɔ́ngɔ",
			language.KM:     "likambodza",
			language.KO:     "likoreya",
			language.LA:     "latina",
			language.LN:     "lingála",
			language.LU:     "kiluba",
			language.LUA:    "ciluba",
			language.MS:     "limalezi",
			language.MY:     "libilimá",
			language.NE:     "linepalɛ",
			language.NL:     "lifalamá",
			language.PA:     "lipendzabi",
			language.PL:     "lipolonɛ",
			language.PT:     "lipulutugɛ́si",
			language.PT_BR:  "lipulutugɛ́si ya Brazil",
			language.PT_PT:  "lipulutugɛ́si ya Erópa",
			language.RM:     "liromansh",
			language.RO:     "liromani",
			language.RU:     "lirisí",
			language.RW:     "kinyarwanda",
			language.SO:     "lisomali",
			language.SV:     "lisuwedɛ",
			language.SW:     "kiswahíli",
			language.TA:     "litamuli",
			language.TH:     "litaye",
			language.TR:     "litiliki",
			language.UK:     "likrɛni",
			language.UR:     "liurdu",
			language.VI:     "liviyetinámi",
			language.YO:     "yoruba",
			language.ZH:     "lisinwa",
			language.ZU:     "zulu",
		},
		Territories: cldr.Territories{
			territory.V_001: "Mabelé",
			territory.V_002: "Afríka",
			territory.V_003: "Ameríka ya Nola",
			territory.V_005: "Ameríka ya Sidi",
			territory.V_011: "Afríka ya Wɛ́sita",
			territory.V_013: "Ameríka ya káti",
			territory.V_014: "Afríka ya Ɛ́sita",
			territory.V_015: "Afríka ya Nola",
			territory.V_017: "Afríka ya Katikáti",
			territory.V_018: "Afríka ya Sidi",
			territory.V_019: "Ameríka",
			territory.V_030: "Azía ya Ɛ́sita",
			territory.V_034: "Azía ya Sidi",
			territory.V_035: "Azía ya Sidi-Ɛ́sita",
			territory.V_039: "Erópa ya Sidi",
			territory.V_142: "Azía",
			territory.V_143: "Azía ya Katikáti",
			territory.V_145: "Azía ya Wɛ́sita",
			territory.V_150: "Erópa",
			territory.V_151: "Erópa ya Ɛ́sita",
			territory.V_154: "Erópa ya Nola",
			territory.V_155: "Erópa ya Wɛ́sita",
			territory.V_419: "Ameríka Latína",
			territory.AC:    "Esanga ya Mbuta o Likoló",
			territory.AD:    "Andorɛ",
			territory.AE:    "Lɛmila alabo",
			territory.AF:    "Afiganisitá",
			territory.AG:    "Antiga mpé Barbuda",
			territory.AI:    "Angiyɛ",
			territory.AL:    "Alibani",
			territory.AM:    "Amɛni",
			territory.AO:    "Angóla",
			territory.AQ:    "Antarctique",
			territory.AR:    "Arizantinɛ",
			territory.AS:    "Samoa ya Ameriki",
			territory.AT:    "Otilisi",
			territory.AU:    "Ositáli",
			territory.AW:    "Aruba",
			territory.AX:    "Bisanga Ɛland",
			territory.AZ:    "Azɛlɛbaizá",
			territory.BA:    "Bosini mpé Hezegovine",
			territory.BB:    "Barɛbadɛ",
			territory.BD:    "Bengalidɛsi",
			territory.BE:    "Beleziki",
			territory.BF:    "Bukina Faso",
			territory.BG:    "Biligari",
			territory.BH:    "Bahrɛnɛ",
			territory.BI:    "Burundi",
			territory.BJ:    "Benɛ",
			territory.BL:    "Sántu Barthélemy",
			territory.BM:    "Bermuda",
			territory.BN:    "Brineyi",
			territory.BO:    "Bolivi",
			territory.BR:    "Brezílɛ",
			territory.BS:    "Bahamasɛ",
			territory.BT:    "Butáni",
			territory.BV:    "Esanga Buvé",
			territory.BW:    "Botswana",
			territory.BY:    "Byelorisi",
			territory.BZ:    "Belizɛ",
			territory.CA:    "Kanada",
			territory.CC:    "Bisanga Kokos",
			territory.CD:    "Kongó-Kinsásá",
			territory.CF:    "Repibiki ya Afríka ya Káti",
			territory.CG:    "Kongó-Brazzaville",
			territory.CH:    "Swisɛ",
			territory.CI:    "Kotídivualɛ",
			territory.CK:    "Bisanga bya Kookɛ",
			territory.CL:    "Síli",
			territory.CM:    "Kamɛrune",
			territory.CN:    "Sinɛ",
			territory.CO:    "Kolombi",
			territory.CP:    "Esanga Clipperton",
			territory.CR:    "Kositarika",
			territory.CU:    "Kiba",
			territory.CV:    "Bisanga bya Kapevɛrɛ",
			territory.CX:    "Esanga ya Mbótama",
			territory.CY:    "Sípɛlɛ",
			territory.CZ:    "Repibiki Tsekɛ",
			territory.DE:    "Alemani",
			territory.DJ:    "Dzibuti",
			territory.DK:    "Danɛmarike",
			territory.DM:    "Domínike",
			territory.DO:    "Repibiki ya Domínikɛ",
			territory.DZ:    "Alizɛri",
			territory.EA:    "Zewta mpé Melílla",
			territory.EC:    "Ekwatɛ́lɛ",
			territory.EE:    "Esitoni",
			territory.EG:    "Ezípite",
			territory.EH:    "Sahara ya Limbɛ",
			territory.ER:    "Elitelɛ",
			territory.ES:    "Esipanye",
			territory.ET:    "Etsíopi",
			territory.EU:    "Lisangá ya Erópa",
			territory.FI:    "Filandɛ",
			territory.FJ:    "Fidzi",
			territory.FK:    "Bisanga bya Falklandí (Bisanga bya Maluni)",
			territory.FM:    "Mikronezi",
			territory.FO:    "Bisanga bya Faloé",
			territory.FR:    "Falánsɛ",
			territory.GA:    "Gabɔ",
			territory.GB:    "Angɛlɛtɛ́lɛ",
			territory.GD:    "Gelenadɛ",
			territory.GE:    "Zorzi",
			territory.GF:    "Giyanɛ ya Falánsɛ",
			territory.GG:    "Guernesey",
			territory.GH:    "Gana",
			territory.GI:    "Zibatalɛ",
			territory.GL:    "Gowelande",
			territory.GM:    "Gambi",
			territory.GN:    "Ginɛ",
			territory.GP:    "Gwadɛlupɛ",
			territory.GQ:    "Ginɛ́kwatɛ́lɛ",
			territory.GR:    "Geleki",
			territory.GS:    "Îles de Géorgie du Sud et Sandwich du Sud",
			territory.GT:    "Gwatémala",
			territory.GU:    "Gwamɛ",
			territory.GW:    "Ginɛbisau",
			territory.GY:    "Giyane",
			territory.HK:    "Hong Kong",
			territory.HM:    "Ile Heard et Iles McDonald",
			territory.HN:    "Ondurasɛ",
			territory.HR:    "Krowasi",
			territory.HT:    "Ayiti",
			territory.HU:    "Ongili",
			territory.IC:    "Bisanga bya Kanári",
			territory.ID:    "Indonezi",
			territory.IE:    "Irelandɛ",
			territory.IL:    "Isirayelɛ",
			territory.IM:    "Esanga ya Man",
			territory.IN:    "Índɛ",
			territory.IO:    "Mabelé ya Angɛlɛtɛ́lɛ na mbú ya Indiya",
			territory.IQ:    "Iraki",
			territory.IR:    "Irâ",
			territory.IS:    "Isilandɛ",
			territory.IT:    "Itali",
			territory.JE:    "Jelezy",
			territory.JM:    "Zamaiki",
			territory.JO:    "Zɔdani",
			territory.JP:    "Zapɔ",
			territory.KE:    "Kenya",
			territory.KG:    "Kigizisitá",
			territory.KH:    "Kambodza",
			territory.KI:    "Kiribati",
			territory.KM:    "Komorɛ",
			territory.KN:    "Sántu krístofe mpé Nevɛ̀s",
			territory.KP:    "Korɛ ya nɔ́rdi",
			territory.KR:    "Korɛ ya súdi",
			territory.KW:    "Koweti",
			territory.KY:    "Bisanga bya Kayíma",
			territory.KZ:    "Kazakisitá",
			territory.LA:    "Lawosi",
			territory.LB:    "Libá",
			territory.LC:    "Sántu lisi",
			territory.LI:    "Lishɛteni",
			territory.LK:    "Sirilanka",
			territory.LR:    "Libériya",
			territory.LS:    "Lesoto",
			territory.LT:    "Litwani",
			territory.LU:    "Likisambulu",
			territory.LV:    "Letoni",
			territory.LY:    "Libí",
			territory.MA:    "Marokɛ",
			territory.MC:    "Monako",
			territory.MD:    "Molidavi",
			territory.ME:    "Monténégro",
			territory.MF:    "Sántu Martin",
			territory.MG:    "Madagasikari",
			territory.MH:    "Bisanga bya Marishalɛ",
			territory.ML:    "Malí",
			territory.MM:    "Birmanie",
			territory.MN:    "Mongolí",
			territory.MO:    "Makau",
			territory.MP:    "Bisanga bya Marianɛ ya nɔ́rdi",
			territory.MQ:    "Martiniki",
			territory.MR:    "Moritani",
			territory.MS:    "Mɔsera",
			territory.MT:    "Malitɛ",
			territory.MU:    "Morisɛ",
			territory.MV:    "Madívɛ",
			territory.MW:    "Malawi",
			territory.MX:    "Meksike",
			territory.MY:    "Malezi",
			territory.MZ:    "Mozambíki",
			territory.NA:    "Namibi",
			territory.NC:    "Kaledoni ya sika",
			territory.NE:    "Nizɛrɛ",
			territory.NF:    "Esanga Norfokɛ",
			territory.NG:    "Nizerya",
			territory.NI:    "Nikaragwa",
			territory.NL:    "Olandɛ",
			territory.NO:    "Norivezɛ",
			territory.NP:    "Nepálɛ",
			territory.NR:    "Nauru",
			territory.NU:    "Nyué",
			territory.NZ:    "Zelandɛ ya sika",
			territory.OM:    "Ománɛ",
			territory.PA:    "Panama",
			territory.PE:    "Péru",
			territory.PF:    "Polinezi ya Falánsɛ",
			territory.PG:    "Papwazi Ginɛ ya sika",
			territory.PH:    "Filipinɛ",
			territory.PK:    "Pakisitá",
			territory.PL:    "Poloni",
			territory.PM:    "Sántu pététo mpé Mikelɔ",
			territory.PN:    "Pikairni",
			territory.PR:    "Pɔtoriko",
			territory.PS:    "Palɛsine",
			territory.PT:    "Putúlugɛsi",
			territory.PW:    "Palau",
			territory.PY:    "Palagwei",
			territory.QA:    "Katari",
			territory.RE:    "Lenyo",
			territory.RO:    "Romani",
			territory.RS:    "Serbie",
			territory.RU:    "Risí",
			territory.RW:    "Rwanda",
			territory.SA:    "Alabi Sawuditɛ",
			territory.SB:    "Bisanga Solomɔ",
			territory.SC:    "Sɛshɛlɛ",
			territory.SD:    "Sudá",
			territory.SE:    "Swédɛ",
			territory.SG:    "Singapurɛ",
			territory.SH:    "Sántu eleni",
			territory.SI:    "Siloveni",
			territory.SJ:    "Svalbard mpé Jan Mayen",
			territory.SK:    "Silovaki",
			territory.SL:    "Siera Leonɛ",
			territory.SM:    "Sántu Marinɛ",
			territory.SN:    "Senegalɛ",
			territory.SO:    "Somali",
			territory.SR:    "Surinamɛ",
			territory.SS:    "Sudani ya Sidi",
			territory.ST:    "Sao Tomé mpé Presipɛ",
			territory.SV:    "Savadɔrɛ",
			territory.SY:    "Sirí",
			territory.SZ:    "Swazilandi",
			territory.TC:    "Bisanga bya Turki mpé Kaiko",
			territory.TD:    "Tsádi",
			territory.TF:    "Terres australes et antarctiques françaises",
			territory.TG:    "Togo",
			territory.TH:    "Tailandɛ",
			territory.TJ:    "Tazikisitá",
			territory.TK:    "Tokelau",
			territory.TL:    "Timor ya Monyɛlɛ",
			territory.TM:    "Tikɛménisitá",
			territory.TN:    "Tinizi",
			territory.TO:    "Tonga",
			territory.TR:    "Tiliki",
			territory.TT:    "Tinidadɛ mpé Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Taiwanin",
			territory.TZ:    "Tanzani",
			territory.UA:    "Ikrɛni",
			territory.UG:    "Uganda",
			territory.UM:    "Bisanga Mokɛ́na Mosíká bya Lisangá lya Ameríka",
			territory.US:    "Ameriki",
			territory.UY:    "Irigwei",
			territory.UZ:    "Uzibɛkisitá",
			territory.VA:    "Vatiká",
			territory.VC:    "Sántu vesá mpé Gelenadinɛ",
			territory.VE:    "Venézuela",
			territory.VG:    "Bisanga bya Vierzi ya Angɛlɛtɛ́lɛ",
			territory.VI:    "Bisanga bya Vierzi ya Ameriki",
			territory.VN:    "Viyetinamɛ",
			territory.VU:    "Vanuatu",
			territory.WF:    "Walisɛ mpé Futuna",
			territory.WS:    "Samoa",
			territory.YE:    "Yemɛnɛ",
			territory.YT:    "Mayotɛ",
			territory.ZA:    "Afríka ya Súdi",
			territory.ZM:    "Zambi",
			territory.ZW:    "Zimbabwe",
		},
	}
}
