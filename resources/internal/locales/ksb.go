// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ksb() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ksb",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januali", Feb: "Febluali", Mar: "Machi", Apr: "Aplili", May: "Mei", Jun: "Juni", Jul: "Julai", Aug: "Agosti", Sep: "Septemba", Oct: "Oktoba", Nov: "Novemba", Dec: "Desemba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jpi", Mon: "Jtt", Tue: "Jmn", Wed: "Jtn", Thu: "Alh", Fri: "Iju", Sat: "Jmo"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "2", Mon: "3", Tue: "4", Wed: "5", Thu: "A", Fri: "I", Sat: "1"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Jumaapii", Mon: "Jumaatatu", Tue: "Jumaane", Wed: "Jumaatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumaamosi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "makeo", PM: "nyiaghuo"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "dilham ya Falme za Kialabu", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "kwanza ya Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "dola ya Austlalia", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "dinali ya Bahaleni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "falanga ya Bulundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "pula ya Botswana", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "dola ya Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "falanga ya Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "falanga ya Uswisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "yaun lenminbi ya China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "eskudo ya Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "falanga ya Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "dinali ya Aljelia", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "pauni ya Misli", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa ya Elitlea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "bil ya Uhabeshi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "yulo", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "pauni ya Uingeeza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "sedi ya Ghana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi ya Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "falanga ya Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "lupia ya India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "salafu ya Kijapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "shilingi ya Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "falanga ya Komolo", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "dola ya Libelia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "loti ya Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "dinali ya Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "dilham ya Moloko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "falanga ya Bukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "ugwiya ya Molitania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ugwiya ya Molitania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "lupia ya Molisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha ya Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "metikali ya Msumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "dola ya Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "naila ya Naijelia", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "falanga ya Lwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "liyal ya Saudia", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "lupia ya Shelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "dinali ya Sudani", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "pauni ya Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "pauni ya Santahelena", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "leoni", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "leoni (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "shilingi ya Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "dobla ya Sao Tome na Plincipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "dobla ya Sao Tome na Plincipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "dinali ya Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "shilingi ya Tanzania", Symbol: "TSh"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "shilingi ya Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "dola ya Malekani", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "falanga CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "falanga CFA BCEAO", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "landi ya Aflika Kusini", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "kwacha ya Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha ya Zambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "dola ya Zimbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Kiakan",
			language.AM:  "Kiamhali",
			language.AR:  "Kialabu",
			language.BE:  "Kibelaausi",
			language.BG:  "Kibulgalia",
			language.BN:  "Kibangla",
			language.CS:  "Kichecki",
			language.DE:  "Kijeumani",
			language.EL:  "Kigiiki",
			language.EN:  "Kiingeeza",
			language.ES:  "Kihispania",
			language.FA:  "Kiajemi",
			language.FR:  "Kifalansa",
			language.HA:  "Kihausa",
			language.HI:  "Kihindi",
			language.HU:  "Kihungai",
			language.ID:  "Kiindonesia",
			language.IG:  "Kiigbo",
			language.IT:  "Kiitaliano",
			language.JA:  "Kijapani",
			language.JV:  "Kijava",
			language.KM:  "Kikambodia",
			language.KO:  "Kikolea",
			language.KSB: "Kishambaa",
			language.MS:  "Kimalesia",
			language.MY:  "Kibulma",
			language.NE:  "Kinepali",
			language.NL:  "Kiholanzi",
			language.PA:  "Kipunjabi",
			language.PL:  "Kipolandi",
			language.PT:  "Kileno",
			language.RO:  "Kiomania",
			language.RU:  "Kilusi",
			language.RW:  "Kinyalwanda",
			language.SO:  "Kisomali",
			language.SV:  "Kiswidi",
			language.TA:  "Kitamil",
			language.TH:  "Kitailandi",
			language.TR:  "Kituuki",
			language.UK:  "Kiuklania",
			language.UR:  "Kiuldu",
			language.VI:  "Kivietinamu",
			language.YO:  "Kiyoluba",
			language.ZH:  "Kichina",
			language.ZU:  "Kizulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andola",
			territory.AE: "Falme za Kialabu",
			territory.AF: "Afuganistani",
			territory.AG: "Antigua na Balbuda",
			territory.AI: "Anguilla",
			territory.AL: "Albania",
			territory.AM: "Almenia",
			territory.AO: "Angola",
			territory.AR: "Ajentina",
			territory.AS: "Samoa ya Malekani",
			territory.AT: "Austlia",
			territory.AU: "Austlalia",
			territory.AW: "Aluba",
			territory.AZ: "Azabajani",
			territory.BA: "Bosnia na Hezegovina",
			territory.BB: "Babadosi",
			territory.BD: "Bangladeshi",
			territory.BF: "Bukinafaso",
			territory.BG: "Bulgalia",
			territory.BH: "Bahaleni",
			territory.BI: "Bulundi",
			territory.BJ: "Benini",
			territory.BM: "Belmuda",
			territory.BN: "Blunei",
			territory.BO: "Bolivia",
			territory.BR: "Blazili",
			territory.BS: "Bahama",
			territory.BT: "Butani",
			territory.BW: "Botswana",
			territory.BY: "Belalusi",
			territory.BZ: "Belize",
			territory.CA: "Kanada",
			territory.CD: "Jamhuli ya Kidemoklasia ya Kongo",
			territory.CF: "Jamhuli ya Afrika ya Gati",
			territory.CG: "Kongo",
			territory.CH: "Uswisi",
			territory.CI: "Kodivaa",
			territory.CK: "Visiwa vya Cook",
			territory.CL: "Chile",
			territory.CM: "Kameluni",
			territory.CN: "China",
			territory.CO: "Kolombia",
			territory.CR: "Kostalika",
			territory.CU: "Kuba",
			territory.CV: "Kepuvede",
			territory.CY: "Kuplosi",
			territory.CZ: "Jamhuli ya Cheki",
			territory.DE: "Ujeumani",
			territory.DJ: "Jibuti",
			territory.DK: "Denmaki",
			territory.DM: "Dominika",
			territory.DO: "Jamhuli ya Dominika",
			territory.DZ: "Aljelia",
			territory.EC: "Ekwado",
			territory.EE: "Estonia",
			territory.EG: "Misli",
			territory.ER: "Elitlea",
			territory.ES: "Hispania",
			territory.ET: "Uhabeshi",
			territory.FI: "Ufini",
			territory.FJ: "Fiji",
			territory.FK: "Visiwa vya Falkland",
			territory.FM: "Miklonesia",
			territory.FR: "Ufalansa",
			territory.GA: "Gaboni",
			territory.GB: "Uingeeza",
			territory.GD: "Glenada",
			territory.GE: "Jojia",
			territory.GF: "Gwiyana ya Ufalansa",
			territory.GH: "Ghana",
			territory.GI: "Jiblalta",
			territory.GL: "Glinlandi",
			territory.GM: "Gambia",
			territory.GN: "Gine",
			territory.GP: "Gwadelupe",
			territory.GQ: "Ginekweta",
			territory.GR: "Ugiiki",
			territory.GT: "Gwatemala",
			territory.GU: "Gwam",
			territory.GW: "Ginebisau",
			territory.GY: "Guyana",
			territory.HN: "Honduasi",
			territory.HR: "Kolasia",
			territory.HT: "Haiti",
			territory.HU: "Hungalia",
			territory.ID: "Indonesia",
			territory.IE: "Ayalandi",
			territory.IL: "Islaeli",
			territory.IN: "India",
			territory.IO: "Eneo ja Uingeeza mwe Bahali Hindi",
			territory.IQ: "Ilaki",
			territory.IR: "Uajemi",
			territory.IS: "Aislandi",
			territory.IT: "Italia",
			territory.JM: "Jamaika",
			territory.JO: "Yoldani",
			territory.JP: "Japani",
			territory.KE: "Kenya",
			territory.KG: "Kiigizistani",
			territory.KH: "Kambodia",
			territory.KI: "Kiibati",
			territory.KM: "Komolo",
			territory.KN: "Santakitzi na Nevis",
			territory.KP: "Kolea Kaskazini",
			territory.KR: "Kolea Kusini",
			territory.KW: "Kuwaiti",
			territory.KY: "Visiwa vya Kayman",
			territory.KZ: "Kazakistani",
			territory.LA: "Laosi",
			territory.LB: "Lebanoni",
			territory.LC: "Santalusia",
			territory.LI: "Lishenteni",
			territory.LK: "Sililanka",
			territory.LR: "Libelia",
			territory.LS: "Lesoto",
			territory.LT: "Litwania",
			territory.LU: "Lasembagi",
			territory.LV: "Lativia",
			territory.LY: "Libya",
			territory.MA: "Moloko",
			territory.MC: "Monako",
			territory.MD: "Moldova",
			territory.MG: "Bukini",
			territory.MH: "Visiwa vya Mashal",
			territory.ML: "Mali",
			territory.MM: "Myama",
			territory.MN: "Mongolia",
			territory.MP: "Visiwa vya Maliana vya Kaskazini",
			territory.MQ: "Maltiniki",
			territory.MR: "Maulitania",
			territory.MS: "Montselati",
			territory.MT: "Malta",
			territory.MU: "Molisi",
			territory.MV: "Modivu",
			territory.MW: "Malawi",
			territory.MX: "Meksiko",
			territory.MY: "Malesia",
			territory.MZ: "Msumbiji",
			territory.NA: "Namibia",
			territory.NC: "Nyukaledonia",
			territory.NE: "Naija",
			territory.NF: "Kisiwa cha Nolfok",
			territory.NG: "Naijelia",
			territory.NI: "Nikalagwa",
			territory.NL: "Uholanzi",
			territory.NO: "Nolwei",
			territory.NP: "Nepali",
			territory.NR: "Naulu",
			territory.NU: "Niue",
			territory.NZ: "Nyuzilandi",
			territory.OM: "Omani",
			territory.PA: "Panama",
			territory.PE: "Pelu",
			territory.PF: "Polinesia ya Ufalansa",
			territory.PG: "Papua",
			territory.PH: "Filipino",
			territory.PK: "Pakistani",
			territory.PL: "Polandi",
			territory.PM: "Santapieli na Mikeloni",
			territory.PN: "Pitkailni",
			territory.PR: "Pwetoliko",
			territory.PS: "Ukingo wa Maghalibi na Ukanda wa Gaza wa Palestina",
			territory.PT: "Uleno",
			territory.PW: "Palau",
			territory.PY: "Palagwai",
			territory.QA: "Katali",
			territory.RE: "Liyunioni",
			territory.RO: "Lomania",
			territory.RU: "Ulusi",
			territory.RW: "Lwanda",
			territory.SA: "Saudi",
			territory.SB: "Visiwa vya Solomon",
			territory.SC: "Shelisheli",
			territory.SD: "Sudani",
			territory.SE: "Uswidi",
			territory.SG: "Singapoo",
			territory.SH: "Santahelena",
			territory.SI: "Slovenia",
			territory.SK: "Slovakia",
			territory.SL: "Siela Leoni",
			territory.SM: "Samalino",
			territory.SN: "Senegali",
			territory.SO: "Somalia",
			territory.SR: "Sulinamu",
			territory.ST: "Sao Tome na Plincipe",
			territory.SV: "Elsavado",
			territory.SY: "Silia",
			territory.SZ: "Uswazi",
			territory.TC: "Visiwa vya Tulki na Kaiko",
			territory.TD: "Chadi",
			territory.TG: "Togo",
			territory.TH: "Tailandi",
			territory.TJ: "Tajikistani",
			territory.TK: "Tokelau",
			territory.TL: "Timoli ya Mashaliki",
			territory.TM: "Tulukimenistani",
			territory.TN: "Tunisia",
			territory.TO: "Tonga",
			territory.TR: "Utuluki",
			territory.TT: "Tlinidad na Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwani",
			territory.TZ: "Tanzania",
			territory.UA: "Uklaini",
			territory.UG: "Uganda",
			territory.US: "Malekani",
			territory.UY: "Ulugwai",
			territory.UZ: "Uzibekistani",
			territory.VA: "Vatikani",
			territory.VC: "Santavisenti na Glenadini",
			territory.VE: "Venezuela",
			territory.VG: "Visiwa vya Vilgin vya Uingeeza",
			territory.VI: "Visiwa vya Vilgin vya Malekani",
			territory.VN: "Vietinamu",
			territory.VU: "Vanuatu",
			territory.WF: "Walis na Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yemeni",
			territory.YT: "Mayotte",
			territory.ZA: "Aflika Kusini",
			territory.ZM: "Zambia",
			territory.ZW: "Zimbabwe",
		},
	}
}
