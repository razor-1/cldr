// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_bm_ML() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "bm_ML",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "zan", Feb: "feb", Mar: "mar", Apr: "awi", May: "mɛ", Jun: "zuw", Jul: "zul", Aug: "uti", Sep: "sɛt", Oct: "ɔku", Nov: "now", Dec: "des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Z", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "Z", Jul: "Z", Aug: "U", Sep: "S", Oct: "Ɔ", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "zanwuye", Feb: "feburuye", Mar: "marisi", Apr: "awirili", May: "mɛ", Jun: "zuwɛn", Jul: "zuluye", Aug: "uti", Sep: "sɛtanburu", Oct: "ɔkutɔburu", Nov: "nowanburu", Dec: "desanburu"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "kar", Mon: "ntɛ", Tue: "tar", Wed: "ara", Thu: "ala", Fri: "jum", Sat: "sib"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "K", Mon: "N", Tue: "T", Wed: "A", Thu: "A", Fri: "J", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "kari", Mon: "ntɛnɛ", Tue: "tarata", Wed: "araba", Thu: "alamisa", Fri: "juma", Sat: "sibiri"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "arabu mara kafoli Diram", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "angola Kwanza", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "ositirali Dolar", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "bareyini Dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "burundi Fraŋ", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "bɔtisiwana Pula", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "kanada Dolar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "kongole Fraŋ", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "suwisi Fraŋ", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "siniwa Yuwan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "capivɛrdi Esekudo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "jibuti Fraŋ", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "alizeri Dinar", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "eziputi Livri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "eritere Nafika", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "etiopi Bir", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "ero", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "angilɛ Livri", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "gana Sedi", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "gambi Dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "gine Fraŋ", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Ɛndu Rupi", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "zapɔne Yɛn", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "keniya Siling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "komɔri Fraŋ", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "liberiya Dolar", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "lesoto Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "libi Dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "marɔku Diram", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "madagasikari Fraŋ", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "mɔritani Uguwiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "mɔritani Uguwiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "morisi Rupi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "malawi Kwaca", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "mozanbiki Metikali", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "namibi Dolar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "nizeriya Nɛra", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "ruwanda Fraŋ", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "sawudiya Riyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "sesɛli Rupi", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "sudani Dinar", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "sudani Livri", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Ɛlɛni-Senu Livri", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "siyeralewɔni Lewɔni", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "siyeralewɔni Lewɔni (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "somali Siling", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "sawotome Dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "sawotome Dobra", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "swazilandi Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "tunizi Dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "tanzani Siling", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "uganda Siling", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "ameriki Dolar", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "sefa Fraŋ (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "sefa Fraŋ (BCEAO)", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "sudafriki Randi", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "zambi Kwaca (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "zambi Kwaca", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "zimbabuwe Dolar", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK: "akankan",
			language.AM: "amarikikan",
			language.AR: "larabukan",
			language.BE: "biyelorisikan",
			language.BG: "buligarikan",
			language.BM: "bamanakan",
			language.BN: "bɛngalikan",
			language.CS: "cɛkikan",
			language.DE: "alimaɲikan",
			language.EL: "gɛrɛsikan",
			language.EN: "angilɛkan",
			language.ES: "esipaɲolkan",
			language.FA: "perisanikan",
			language.FR: "tubabukan",
			language.HA: "awusakan",
			language.HI: "inidikan",
			language.HU: "oŋirikan",
			language.ID: "Ɛndonezikan",
			language.IG: "igibokan",
			language.IT: "italikan",
			language.JA: "zapɔnekan",
			language.JV: "javanekan",
			language.KM: "kambojikan",
			language.KO: "korekan",
			language.MS: "malɛzikan",
			language.MY: "birimanikan",
			language.NE: "nepalekan",
			language.NL: "olandekan",
			language.PA: "pɛnijabikan",
			language.PL: "polonekan",
			language.PT: "pɔritigalikan",
			language.RO: "rumanikan",
			language.RU: "irisikan",
			language.RW: "ruwandakan",
			language.SO: "somalikan",
			language.SV: "suwɛdikan",
			language.TA: "tamulikan",
			language.TH: "tayikan",
			language.TR: "turikikan",
			language.UK: "ukɛrɛnikan",
			language.UR: "urudukan",
			language.VI: "wiyɛtinamukan",
			language.YO: "yorubakan",
			language.ZH: "siniwakan",
			language.ZU: "zulukan",
		},
		Territories: cldr.Territories{
			territory.AD: "Andɔr",
			territory.AE: "Arabu mara kafoli",
			territory.AF: "Afiganistaŋ",
			territory.AG: "Antiga-ni-Barbuda",
			territory.AI: "Angiya",
			territory.AL: "Alibani",
			territory.AM: "Arimeni",
			territory.AO: "Angola",
			territory.AR: "Arizantin",
			territory.AS: "Samowa amerikani",
			territory.AT: "Otirisi",
			territory.AU: "Ositirali",
			territory.AW: "Aruba",
			territory.AZ: "Azɛrbayjaŋ",
			territory.BA: "Bozni-Ɛrizigovini",
			territory.BB: "Barbadi",
			territory.BD: "Bɛngiladɛsi",
			territory.BE: "Bɛliziki",
			territory.BF: "Burukina Faso",
			territory.BG: "Buligari",
			territory.BH: "Bareyini",
			territory.BI: "Burundi",
			territory.BJ: "Benɛn",
			territory.BM: "Bermudi",
			territory.BN: "Burinɛyi",
			territory.BO: "Bolivi",
			territory.BR: "Berezili",
			territory.BS: "Bahamasi",
			territory.BT: "Butaŋ",
			territory.BW: "Bɔtisiwana",
			territory.BY: "Belarusi",
			territory.BZ: "Belizi",
			territory.CA: "Kanada",
			territory.CD: "Kongo ka republiki demɔkratiki",
			territory.CF: "Santarafiriki",
			territory.CG: "Kongo",
			territory.CH: "Suwisi",
			territory.CI: "Kodiwari",
			territory.CK: "Kuki Gun",
			territory.CL: "Sili",
			territory.CM: "Kameruni",
			territory.CN: "Siniwajamana",
			territory.CO: "Kolombi",
			territory.CR: "Kɔsitarika",
			territory.CU: "Kuba",
			territory.CV: "Capivɛrdi",
			territory.CY: "Cipri",
			territory.CZ: "Ceki republiki",
			territory.DE: "Alimaɲi",
			territory.DJ: "Jibuti",
			territory.DK: "Danemarki",
			territory.DM: "Dɔminiki",
			territory.DO: "Dɔmimiki republiki",
			territory.DZ: "Alizeri",
			territory.EC: "Ekwatɔr",
			territory.EE: "Esetoni",
			territory.EG: "Eziputi",
			territory.ER: "Eritere",
			territory.ES: "Esipaɲi",
			territory.ET: "Etiopi",
			territory.FI: "Finilandi",
			territory.FJ: "Fiji",
			territory.FK: "Maluwini Gun",
			territory.FM: "Mikironesi",
			territory.FR: "Faransi",
			territory.GA: "Gabɔŋ",
			territory.GB: "Angilɛtɛri",
			territory.GD: "Granadi",
			territory.GE: "Zeyɔrzi",
			territory.GF: "Faransi ka gwiyani",
			territory.GH: "Gana",
			territory.GI: "Zibralitari",
			territory.GL: "Gɔrɔhenelandi",
			territory.GM: "Ganbi",
			territory.GN: "Gine",
			territory.GP: "Gwadelup",
			territory.GQ: "Gine ekwatɔri",
			territory.GR: "Gɛrɛsi",
			territory.GT: "Gwatemala",
			territory.GU: "Gwam",
			territory.GW: "Gine Bisawo",
			territory.GY: "Gwiyana",
			territory.HN: "Hɔndirasi",
			territory.HR: "Kroasi",
			territory.HT: "Ayiti",
			territory.HU: "Hɔngri",
			territory.ID: "Ɛndonezi",
			territory.IE: "Irilandi",
			territory.IL: "Isirayeli",
			territory.IN: "Ɛndujamana",
			territory.IO: "Angilɛ ka ɛndu dugukolo",
			territory.IQ: "Iraki",
			territory.IR: "Iraŋ",
			territory.IS: "Isilandi",
			territory.IT: "Itali",
			territory.JM: "Zamayiki",
			territory.JO: "Zɔrdani",
			territory.JP: "Zapɔn",
			territory.KE: "Keniya",
			territory.KG: "Kirigizisitaŋ",
			territory.KH: "Kamboji",
			territory.KI: "Kiribati",
			territory.KM: "Komɔri",
			territory.KN: "Kristɔfo-Senu-ni-Ɲevɛs",
			territory.KP: "Kɛɲɛka Kore",
			territory.KR: "Worodugu Kore",
			territory.KW: "Kowɛti",
			territory.KY: "Bama Gun",
			territory.KZ: "Kazakistaŋ",
			territory.LA: "Layosi",
			territory.LB: "Libaŋ",
			territory.LC: "Lusi-Senu",
			territory.LI: "Lisɛnsitayini",
			territory.LK: "Sirilanka",
			territory.LR: "Liberiya",
			territory.LS: "Lesoto",
			territory.LT: "Lituyani",
			territory.LU: "Likisanburu",
			territory.LV: "Letoni",
			territory.LY: "Libi",
			territory.MA: "Marɔku",
			territory.MC: "Monako",
			territory.MD: "Molidavi",
			territory.MG: "Madagasikari",
			territory.MH: "Marisali Gun",
			territory.ML: "Mali",
			territory.MM: "Myanimari",
			territory.MN: "Moŋoli",
			territory.MP: "Kɛɲɛka Mariyani Gun",
			territory.MQ: "Maritiniki",
			territory.MR: "Mɔritani",
			territory.MS: "Moŋsera",
			territory.MT: "Malti",
			territory.MU: "Morisi",
			territory.MV: "Maldivi",
			territory.MW: "Malawi",
			territory.MX: "Meksiki",
			territory.MY: "Malɛzi",
			territory.MZ: "Mozanbiki",
			territory.NA: "Namibi",
			territory.NC: "Kaledoni Koura",
			territory.NE: "Nizɛri",
			territory.NF: "Nɔrofoliki Gun",
			territory.NG: "Nizeriya",
			territory.NI: "Nikaragwa",
			territory.NL: "Peyiba",
			territory.NO: "Nɔriwɛzi",
			territory.NP: "Nepali",
			territory.NR: "Nawuru",
			territory.NU: "Nyuwe",
			territory.NZ: "Zelandi Koura",
			territory.OM: "Omaŋ",
			territory.PA: "Panama",
			territory.PE: "Peru",
			territory.PF: "Faransi ka polinezi",
			territory.PG: "Papuwasi-Gine-Koura",
			territory.PH: "Filipini",
			territory.PK: "Pakisitaŋ",
			territory.PL: "Poloɲi",
			territory.PM: "Piyɛri-Senu-ni-Mikelɔŋ",
			territory.PN: "Pitikarini",
			territory.PR: "Pɔrotoriko",
			territory.PS: "Palesitini",
			territory.PT: "Pɔritigali",
			territory.PW: "Palawu",
			territory.PY: "Paraguwayi",
			territory.QA: "Katari",
			territory.RE: "Reyuɲɔŋ",
			territory.RO: "Rumani",
			territory.RU: "Irisi",
			territory.RW: "Ruwanda",
			territory.SA: "Arabiya Sawudiya",
			territory.SB: "Salomo Gun",
			territory.SC: "Sesɛli",
			territory.SD: "Sudaŋ",
			territory.SE: "Suwɛdi",
			territory.SG: "Sɛngapuri",
			territory.SH: "Ɛlɛni Senu",
			territory.SI: "Sloveni",
			territory.SK: "Slowaki",
			territory.SL: "Siyera Lewɔni",
			territory.SM: "Marini-Senu",
			territory.SN: "Senegali",
			territory.SO: "Somali",
			territory.SR: "Surinami",
			territory.ST: "Sawo Tome-ni-Prinicipe",
			territory.SV: "Salivadɔr",
			territory.SY: "Siri",
			territory.SZ: "Swazilandi",
			territory.TC: "Turiki Gun ni Kayiki",
			territory.TD: "Cadi",
			territory.TG: "Togo",
			territory.TH: "Tayilandi",
			territory.TJ: "Tajikisitani",
			territory.TK: "Tokelo",
			territory.TL: "Kɔrɔn Timɔr",
			territory.TM: "Turikimenisitani",
			territory.TN: "Tunizi",
			territory.TO: "Tonga",
			territory.TR: "Turiki",
			territory.TT: "Trinite-ni-Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Tayiwani",
			territory.TZ: "Tanzani",
			territory.UA: "Ukɛrɛni",
			territory.UG: "Uganda",
			territory.US: "Ameriki",
			territory.UY: "Urugwayi",
			territory.UZ: "Uzebekisitani",
			territory.VA: "Vatikaŋ",
			territory.VC: "Vinisɛn-Senu-ni-Grenadini",
			territory.VE: "Venezuwela",
			territory.VG: "Angilɛ ka Sungurunnin Gun",
			territory.VI: "Ameriki ka Sungurunnin Gun",
			territory.VN: "Wiyɛtinamu",
			territory.VU: "Vanuwatu",
			territory.WF: "Walisi-ni-Futuna",
			territory.WS: "Samowa",
			territory.YE: "Yemɛni",
			territory.YT: "Mayoti",
			territory.ZA: "Worodugu Afriki",
			territory.ZM: "Zanbi",
			territory.ZW: "Zimbabuwe",
		},
	}
}
