// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_sbp_TZ() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "sbp_TZ",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Mup", Feb: "Mwi", Mar: "Msh", Apr: "Mun", May: "Mag", Jun: "Muj", Jul: "Msp", Aug: "Mpg", Sep: "Mye", Oct: "Mok", Nov: "Mus", Dec: "Muh"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Mupalangulwa", Feb: "Mwitope", Mar: "Mushende", Apr: "Munyi", May: "Mushende Magali", Jun: "Mujimbi", Jul: "Mushipepo", Aug: "Mupuguto", Sep: "Munyense", Oct: "Mokhu", Nov: "Musongandembwe", Dec: "Muhaano"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Mul", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Alh", Fri: "Iju", Sat: "Jmo"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "M", Mon: "J", Tue: "J", Wed: "J", Thu: "A", Fri: "I", Sat: "J"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Mulungu", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alahamisi", Fri: "Ijumaa", Sat: "Jumamosi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Lwamilawu", PM: "Pashamihe"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Lwamilawu", PM: "Pashamihe"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00¤", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Ihela ya Shitwa sha Shiyalabu", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Ihela ya Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Ihela ya Awusitilaliya", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Ihela ya Bahaleni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Ihela ya Bulundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Ihela ya Botiswana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Ihela ya Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Ihela ya Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Ihela ya Uswisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Ihela ya Shina", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Ihela ya Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Ihela ya Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Ihela ya Alijeliya", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Ihela ya Misili", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Ihela ya Elitileya", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ihela ya Uhabeshi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Ihela ya Ulaya", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Ihela ya Uwingelesa", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Ihela ya Ghana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Ihela ya Gambiya", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Ihela ya Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Ihela ya Indiya", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Ihela ya Japani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Ihela ya Kenya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Ihela ya Komolo", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Ihela ya Libeliya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Ihela ya Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Ihela ya Libiya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Ihela ya Moloko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ihela ya Bukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ihela ya Molitaniya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ihela ya Molitaniya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Ihela ya Molisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Ihela ya Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Ihela ya Musumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Ihela ya Namibiya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Ihela ya Nijeliya", Symbol: "₦"},
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
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ihela ya Lwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Ihela ya Sawudiya", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Ihela ya Shelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Ihela ya Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Ihela ya Santahelena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Ihela ya Siela Liyoni", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Ihela ya Somaliya", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Ihela ya Sao Tome ni Pilinsipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Ihela ya Sao Tome ni Pilinsipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Ihela ya Uswasi", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Ihela ya Tunisiya", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Ihela ya Tansaniya", Symbol: "TSh"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Ihela ya Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Ihela ya Malekani", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Ihela ya CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Ihela ya CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Ihela ya Afilika Kusini", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Ihela ya Sambiya (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Ihela ya Sambiya", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Ihela ya Simbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Ishiyakani",
			language.AM:  "Ishiyamuhali",
			language.AR:  "Ishiyalabu",
			language.BE:  "Ishibelalusi",
			language.BG:  "Ishibulugalia",
			language.BN:  "Ishibangila",
			language.CS:  "Ishisheki",
			language.DE:  "Ishijelumani",
			language.EL:  "Ishigiliki",
			language.EN:  "Ishingelesa",
			language.ES:  "Ishihisipaniya",
			language.FA:  "Ishiajemi",
			language.FR:  "Ishifalansa",
			language.HA:  "Ishihawusa",
			language.HI:  "Ishihindi",
			language.HU:  "Ishihungali",
			language.ID:  "Ishihindonesia",
			language.IG:  "Ishihigibo",
			language.IT:  "Ishihitaliyano",
			language.JA:  "Ishijapani",
			language.JV:  "Ishijava",
			language.KM:  "Ishikambodia",
			language.KO:  "Ishikoleya",
			language.MS:  "Ishimalesiya",
			language.MY:  "Ishibuluma",
			language.NE:  "Ishinepali",
			language.NL:  "Ishiholansi",
			language.PA:  "Ishipunjabi",
			language.PL:  "Ishipolandi",
			language.PT:  "Ishileno",
			language.RO:  "Ishilomaniya",
			language.RU:  "Ishilusi",
			language.RW:  "Ishinyalwanda",
			language.SBP: "Ishisangu",
			language.SO:  "Ishisomali",
			language.SV:  "Ishiswidi",
			language.TA:  "Ishitamili",
			language.TH:  "Ishitayilandi",
			language.TR:  "Ishituluki",
			language.UK:  "Ishiyukilaniya",
			language.UR:  "Ishiwuludi",
			language.VI:  "Ishivietinamu",
			language.YO:  "Ishiyoluba",
			language.ZH:  "Ishishina",
			language.ZU:  "Ishisulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andola",
			territory.AE: "Wutwa wa shiyalabu",
			territory.AF: "Afuganisitani",
			territory.AG: "Anitiguya ni Balubuda",
			territory.AI: "Anguilla",
			territory.AL: "Alubaniya",
			territory.AM: "Alimeniya",
			territory.AO: "Angola",
			territory.AR: "Ajentina",
			territory.AS: "Samoya ya Malekani",
			territory.AT: "Awusitiliya",
			territory.AU: "Awusitilaliya",
			territory.AW: "Aluba",
			territory.AZ: "Asabajani",
			territory.BA: "Bosiniya ni Hesegovina",
			territory.BB: "Babadosi",
			territory.BD: "Bangiladeshi",
			territory.BE: "Ubeligiji",
			territory.BF: "Bukinafaso",
			territory.BG: "Buligaliya",
			territory.BH: "Bahaleni",
			territory.BI: "Bulundi",
			territory.BJ: "Benini",
			territory.BM: "Belimuda",
			territory.BN: "Buluneyi",
			territory.BO: "Boliviya",
			territory.BR: "Bulasili",
			territory.BS: "Bahama",
			territory.BT: "Butani",
			territory.BW: "Botiswana",
			territory.BY: "Belalusi",
			territory.BZ: "Belise",
			territory.CA: "Kanada",
			territory.CD: "Jamuhuli ya Kidemokilasiya ya Kongo",
			territory.CF: "Jamuhuli ya Afilika ya Pakhati",
			territory.CG: "Kongo",
			territory.CH: "Uswisi",
			territory.CI: "Kodivaya",
			territory.CK: "Figunguli fya Kooki",
			territory.CL: "Shile",
			territory.CM: "Kameruni",
			territory.CN: "Shina",
			territory.CO: "Kolombiya",
			territory.CR: "Kositalika",
			territory.CU: "Kuba",
			territory.CV: "Kepuvede",
			territory.CY: "Kupilosi",
			territory.CZ: "Jamuhuli ya Sheki",
			territory.DE: "Wujelumani",
			territory.DJ: "Jibuti",
			territory.DK: "Denimaki",
			territory.DM: "Dominika",
			territory.DO: "Jamuhuli ya Dominika",
			territory.DZ: "Alijeliya",
			territory.EC: "Ekwado",
			territory.EE: "Esitoniya",
			territory.EG: "Misili",
			territory.ER: "Elitileya",
			territory.ES: "Hisipaniya",
			territory.ET: "Uhabeshi",
			territory.FI: "Wufini",
			territory.FJ: "Fiji",
			territory.FK: "Figunguli fya Fokolendi",
			territory.FM: "Mikilonesiya",
			territory.FR: "Wufalansa",
			territory.GA: "Gaboni",
			territory.GB: "Uwingelesa",
			territory.GD: "Gilenada",
			territory.GE: "Jojiya",
			territory.GF: "Gwiyana ya Wufalansa",
			territory.GH: "Khana",
			territory.GI: "Jibulalita",
			territory.GL: "Gilinilandi",
			territory.GM: "Gambiya",
			territory.GN: "Gine",
			territory.GP: "Gwadelupe",
			territory.GQ: "Ginekweta",
			territory.GR: "Wugiliki",
			territory.GT: "Gwatemala",
			territory.GU: "Gwamu",
			territory.GW: "Ginebisawu",
			territory.GY: "Guyana",
			territory.HN: "Hondulasi",
			territory.HR: "Kolasiya",
			territory.HT: "Hayiti",
			territory.HU: "Hungaliya",
			territory.ID: "Indonesiya",
			territory.IE: "Ayalandi",
			territory.IL: "Isilaeli",
			territory.IN: "Indiya",
			territory.IO: "Uluvala lwa Uwingelesa ku Bahali ya Hindi",
			territory.IQ: "Ilaki",
			territory.IR: "Uwajemi",
			territory.IS: "Ayisilendi",
			territory.IT: "Italiya",
			territory.JM: "Jamaika",
			territory.JO: "Yolodani",
			territory.JP: "Japani",
			territory.KE: "Kenya",
			territory.KG: "Kiligisisitani",
			territory.KH: "Kambodiya",
			territory.KI: "Kilibati",
			territory.KM: "Komolo",
			territory.KN: "Santakitisi ni Nevisi",
			territory.KP: "Koleya ya luvala lwa Kunyamande",
			territory.KR: "Koleya ya Kusini",
			territory.KW: "Kuwaiti",
			territory.KY: "Figunguli ifya Kayimayi",
			territory.KZ: "Kasakisitani",
			territory.LA: "Layosi",
			territory.LB: "Lebanoni",
			territory.LC: "Santalusiya",
			territory.LI: "Lisheniteni",
			territory.LK: "Sililanka",
			territory.LR: "Libeliya",
			territory.LS: "Lesoto",
			territory.LT: "Litwaniya",
			territory.LU: "Lasembagi",
			territory.LV: "Lativiya",
			territory.LY: "Libiya",
			territory.MA: "Moloko",
			territory.MC: "Monako",
			territory.MD: "Molidova",
			territory.MG: "Bukini",
			territory.MH: "Figunguli ifya Malishali",
			territory.ML: "Mali",
			territory.MM: "Muyama",
			territory.MN: "Mongoliya",
			territory.MP: "Figunguli fya Maliyana ifya luvala lwa Kunyamande",
			territory.MQ: "Malitiniki",
			territory.MR: "Molitaniya",
			territory.MS: "Monitiselati",
			territory.MT: "Malita",
			territory.MU: "Molisi",
			territory.MV: "Modivu",
			territory.MW: "Malawi",
			territory.MX: "Mekisiko",
			territory.MY: "Malesiya",
			territory.MZ: "Musumbiji",
			territory.NA: "Namibiya",
			territory.NC: "Nyukaledoniya",
			territory.NE: "Nijeli",
			territory.NF: "Shigunguli sha Nolifoki",
			territory.NG: "Nijeliya",
			territory.NI: "Nikalagwa",
			territory.NL: "Wuholansi",
			territory.NO: "Nolwe",
			territory.NP: "Nepali",
			territory.NR: "Nawulu",
			territory.NU: "Niwue",
			territory.NZ: "Nyusilendi",
			territory.OM: "Omani",
			territory.PA: "Panama",
			territory.PE: "Pelu",
			territory.PF: "Polinesiya ya Wufalansa",
			territory.PG: "Papuwa",
			territory.PH: "Filipino",
			territory.PK: "Pakisitani",
			territory.PL: "Polandi",
			territory.PM: "Santapieli ni Mikeloni",
			territory.PN: "Pitikailini",
			territory.PR: "Pwetoliko",
			territory.PS: "Munjema gwa Kusikha nu Luvala lwa Gasa lwa Palesit",
			territory.PT: "Wuleno",
			territory.PW: "Palawu",
			territory.PY: "Palagwayi",
			territory.QA: "Katali",
			territory.RE: "Liyunioni",
			territory.RO: "Lomaniya",
			territory.RU: "Wulusi",
			territory.RW: "Lwanda",
			territory.SA: "Sawudi",
			territory.SB: "Figunguli fya Solomoni",
			territory.SC: "Shelisheli",
			territory.SD: "Sudani",
			territory.SE: "Uswidi",
			territory.SG: "Singapoo",
			territory.SH: "Santahelena",
			territory.SI: "Siloveniya",
			territory.SK: "Silovakiya",
			territory.SL: "Siela Liyoni",
			territory.SM: "Samalino",
			territory.SN: "Senegali",
			territory.SO: "Somaliya",
			territory.SR: "Sulinamu",
			territory.ST: "Sayo Tome ni Pilinikipe",
			territory.SV: "Elisavado",
			territory.SY: "Siliya",
			territory.SZ: "Uswasi",
			territory.TC: "Figunguli fya Tuliki ni Kaiko",
			territory.TD: "Shadi",
			territory.TG: "Togo",
			territory.TH: "Tailandi",
			territory.TJ: "Tajikisitani",
			territory.TK: "Tokelawu",
			territory.TL: "Timoli ya kunena",
			territory.TM: "Tulukimenisitani",
			territory.TN: "Tunisiya",
			territory.TO: "Tonga",
			territory.TR: "Utuluki",
			territory.TT: "Tilinidadi ni Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwani",
			territory.TZ: "Tansaniya",
			territory.UA: "Yukileini",
			territory.UG: "Uganda",
			territory.US: "Malekani",
			territory.UY: "Ulugwayi",
			territory.UZ: "Usibekisitani",
			territory.VA: "Vatikani",
			territory.VC: "Santavisenti na Gilenadini",
			territory.VE: "Venesuela",
			territory.VG: "Figunguli ifya Viliginiya ifya Uwingelesa",
			territory.VI: "Figunguli fya Viliginiya ifya Malekani",
			territory.VN: "Vietinamu",
			territory.VU: "Vanuatu",
			territory.WF: "Walisi ni Futuna",
			territory.WS: "Samoya",
			territory.YE: "Yemeni",
			territory.YT: "Mayote",
			territory.ZA: "Afilika Kusini",
			territory.ZM: "Sambiya",
			territory.ZW: "Simbabwe",
		},
	}
}
