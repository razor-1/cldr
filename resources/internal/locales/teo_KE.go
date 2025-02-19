// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_teo_KE() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "teo_KE",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Rar", Feb: "Muk", Mar: "Kwa", Apr: "Dun", May: "Mar", Jun: "Mod", Jul: "Jol", Aug: "Ped", Sep: "Sok", Oct: "Tib", Nov: "Lab", Dec: "Poo"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "R", Feb: "M", Mar: "K", Apr: "D", May: "M", Jun: "M", Jul: "J", Aug: "P", Sep: "S", Oct: "T", Nov: "L", Dec: "P"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Orara", Feb: "Omuk", Mar: "Okwamg’", Apr: "Odung’el", May: "Omaruk", Jun: "Omodok’king’ol", Jul: "Ojola", Aug: "Opedel", Sep: "Osokosokoma", Oct: "Otibar", Nov: "Olabor", Dec: "Opoo"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jum", Mon: "Bar", Tue: "Aar", Wed: "Uni", Thu: "Ung", Fri: "Kan", Sat: "Sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "J", Mon: "B", Tue: "A", Wed: "U", Thu: "U", Fri: "K", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Nakaejuma", Mon: "Nakaebarasa", Tue: "Nakaare", Wed: "Nakauni", Thu: "Nakaung’on", Fri: "Nakakany", Sat: "Nakasabiti"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Taparachu", PM: "Ebongi"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Ango’otol lok’ Falme za Kiarabu", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "Ango’otol lok’ Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Ango’otol lok’ Australia", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Ango’otol lok’ Bahareni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Ango’otol lok’ Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Ango’otol lok’ Botswana", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Edola lok’Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Ango’otol lok’ Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Ango’otol lok’ Uswisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Ango’otol lok’ China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Ango’otol lok’ Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Ango’otol lok’ Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Ango’otol lok’ Aljeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Epaunt lok’ Misri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Ango’otol lok’ Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ango’otol lok’ Uhabeshi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Epaunt lok’ Uingereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi ya Ghana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ya Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Ango’otol lok’ Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Ango’otol lok’ India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Ango’otol lok’ Kijapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Ango’otol lok’ Kenya", Symbol: "Ksh"},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Ango’otol lok’ Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "edola lok’ Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Ango’otol lok’ Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Ango’otol lok’ Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Ango’otol lok’ Moroko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ango’otol lok’ Bukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ango’otol lok’ Moritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ango’otol lok’ Moritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Ango’otol lok’ Morisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Ango’otol lok’ Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Ango’otol lok’ Msumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Ango’otol lok’ Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Ango’otol lok’ Nijeria", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Ango’otol lok’ Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Ango’otol lok’ Saudia", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Ango’otol lok’ Shelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Epaunt Lok’ Sudan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Ango’otol lok’ Santahelena", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Leoni", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leoni (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Ango’otol lok’ Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Ango’otol lok’ Sao Tome na Principe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Ango’otol lok’ Sao Tome na Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Ango’otol lok’ Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Ango’otol lok’ Tanzania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Ango’otol lok’ Uganda", Symbol: "USh"},
				currency.USD: cldr.Currency{DisplayName: "edola lok’ Amareka", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Ango’otol lok’ CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Ango’otol lok’ CFA BCEAO", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Ango’otol lok’ Afrika Kusini", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Ango’otol lok’ Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Ango’otol lok’ Zambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Edola lok’Zimbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Kiakan",
			language.AM:  "Kiamhari",
			language.AR:  "Kiarabu",
			language.BE:  "Kibelarusi",
			language.BG:  "Kibulgaria",
			language.BN:  "Kibangla",
			language.CS:  "Kichecki",
			language.DE:  "Kijerumani",
			language.EL:  "Kigiriki",
			language.EN:  "Kingereza",
			language.ES:  "Kihispania",
			language.FA:  "Kiajemi",
			language.FR:  "Kifaransa",
			language.HA:  "Kihausa",
			language.HI:  "Kihindi",
			language.HU:  "Kihungari",
			language.ID:  "Kiindonesia",
			language.IG:  "Kiigbo",
			language.IT:  "Kiitaliano",
			language.JA:  "Kijapani",
			language.JV:  "Kijava",
			language.KM:  "Kikambodia",
			language.KO:  "Kikorea",
			language.MS:  "Kimalesia",
			language.MY:  "Kiburma",
			language.NE:  "Kinepali",
			language.NL:  "Kiholanzi",
			language.PA:  "Kipunjabi",
			language.PL:  "Kipolandi",
			language.PT:  "Kireno",
			language.RO:  "Kiromania",
			language.RU:  "Kirusi",
			language.RW:  "Kinyarwanda",
			language.SO:  "Kisomali",
			language.SV:  "Kiswidi",
			language.TA:  "Kitamil",
			language.TEO: "Kiteso",
			language.TH:  "Kitailandi",
			language.TR:  "Kituruki",
			language.UK:  "Kiukrania",
			language.UR:  "Kiurdu",
			language.VI:  "Kivietinamu",
			language.YO:  "Kiyoruba",
			language.ZH:  "Kichina",
			language.ZU:  "Kizulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andora",
			territory.AE: "Falme za Kiarabu",
			territory.AF: "Afuganistani",
			territory.AG: "Antigua na Barbuda",
			territory.AI: "Anguilla",
			territory.AL: "Albania",
			territory.AM: "Armenia",
			territory.AO: "Angola",
			territory.AR: "Ajentina",
			territory.AS: "Samoa ya Marekani",
			territory.AT: "Austria",
			territory.AU: "Australia",
			territory.AW: "Aruba",
			territory.AZ: "Azabajani",
			territory.BA: "Bosnia na Hezegovina",
			territory.BB: "Babadosi",
			territory.BD: "Bangladeshi",
			territory.BE: "Ubelgiji",
			territory.BF: "Bukinafaso",
			territory.BG: "Bulgaria",
			territory.BH: "Bahareni",
			territory.BI: "Burundi",
			territory.BJ: "Benini",
			territory.BM: "Bermuda",
			territory.BN: "Brunei",
			territory.BO: "Bolivia",
			territory.BR: "Brazili",
			territory.BS: "Bahama",
			territory.BT: "Butani",
			territory.BW: "Botswana",
			territory.BY: "Belarusi",
			territory.BZ: "Belize",
			territory.CA: "Kanada",
			territory.CD: "Jamhuri ya Kidemokrasia ya Kongo",
			territory.CF: "Jamhuri ya Afrika ya Kati",
			territory.CG: "Kongo",
			territory.CH: "Uswisi",
			territory.CI: "Kodivaa",
			territory.CK: "Visiwa vya Cook",
			territory.CL: "Chile",
			territory.CM: "Kameruni",
			territory.CN: "China",
			territory.CO: "Kolombia",
			territory.CR: "Kostarika",
			territory.CU: "Kuba",
			territory.CV: "Kepuvede",
			territory.CY: "Kuprosi",
			territory.CZ: "Jamhuri ya Cheki",
			territory.DE: "Ujerumani",
			territory.DJ: "Jibuti",
			territory.DK: "Denmaki",
			territory.DM: "Dominika",
			territory.DO: "Jamhuri ya Dominika",
			territory.DZ: "Aljeria",
			territory.EC: "Ekwado",
			territory.EE: "Estonia",
			territory.EG: "Misri",
			territory.ER: "Eritrea",
			territory.ES: "Hispania",
			territory.ET: "Uhabeshi",
			territory.FI: "Ufini",
			territory.FJ: "Fiji",
			territory.FK: "Visiwa vya Falkland",
			territory.FM: "Mikronesia",
			territory.FR: "Ufaransa",
			territory.GA: "Gaboni",
			territory.GB: "Uingereza",
			territory.GD: "Grenada",
			territory.GE: "Jojia",
			territory.GF: "Gwiyana ya Ufaransa",
			territory.GH: "Ghana",
			territory.GI: "Jibralta",
			territory.GL: "Grinlandi",
			territory.GM: "Gambia",
			territory.GN: "Gine",
			territory.GP: "Gwadelupe",
			territory.GQ: "Ginekweta",
			territory.GR: "Ugiriki",
			territory.GT: "Gwatemala",
			territory.GU: "Gwam",
			territory.GW: "Ginebisau",
			territory.GY: "Guyana",
			territory.HN: "Hondurasi",
			territory.HR: "Korasia",
			territory.HT: "Haiti",
			territory.HU: "Hungaria",
			territory.ID: "Indonesia",
			territory.IE: "Ayalandi",
			territory.IL: "Israeli",
			territory.IN: "India",
			territory.IO: "Eneo la Uingereza katika Bahari Hindi",
			territory.IQ: "Iraki",
			territory.IR: "Uajemi",
			territory.IS: "Aislandi",
			territory.IT: "Italia",
			territory.JM: "Jamaika",
			territory.JO: "Yordani",
			territory.JP: "Japani",
			territory.KE: "Kenia",
			territory.KG: "Kirigizistani",
			territory.KH: "Kambodia",
			territory.KI: "Kiribati",
			territory.KM: "Komoro",
			territory.KN: "Santakitzi na Nevis",
			territory.KP: "Korea Kaskazini",
			territory.KR: "Korea Kusini",
			territory.KW: "Kuwaiti",
			territory.KY: "Visiwa vya Kayman",
			territory.KZ: "Kazakistani",
			territory.LA: "Laosi",
			territory.LB: "Lebanoni",
			territory.LC: "Santalusia",
			territory.LI: "Lishenteni",
			territory.LK: "Sirilanka",
			territory.LR: "Liberia",
			territory.LS: "Lesoto",
			territory.LT: "Litwania",
			territory.LU: "Lasembagi",
			territory.LV: "Lativia",
			territory.LY: "Libya",
			territory.MA: "Moroko",
			territory.MC: "Monako",
			territory.MD: "Moldova",
			territory.MG: "Bukini",
			territory.MH: "Visiwa vya Marshal",
			territory.ML: "Mali",
			territory.MM: "Myama",
			territory.MN: "Mongolia",
			territory.MP: "Visiwa vya Mariana vya Kaskazini",
			territory.MQ: "Martiniki",
			territory.MR: "Moritania",
			territory.MS: "Montserrati",
			territory.MT: "Malta",
			territory.MU: "Morisi",
			territory.MV: "Modivu",
			territory.MW: "Malawi",
			territory.MX: "Meksiko",
			territory.MY: "Malesia",
			territory.MZ: "Msumbiji",
			territory.NA: "Namibia",
			territory.NC: "Nyukaledonia",
			territory.NE: "Nijeri",
			territory.NF: "Kisiwa cha Norfok",
			territory.NG: "Nijeria",
			territory.NI: "Nikaragwa",
			territory.NL: "Uholanzi",
			territory.NO: "Norwe",
			territory.NP: "Nepali",
			territory.NR: "Nauru",
			territory.NU: "Niue",
			territory.NZ: "Nyuzilandi",
			territory.OM: "Omani",
			territory.PA: "Panama",
			territory.PE: "Peru",
			territory.PF: "Polinesia ya Ufaransa",
			territory.PG: "Papua",
			territory.PH: "Filipino",
			territory.PK: "Pakistani",
			territory.PL: "Polandi",
			territory.PM: "Santapieri na Mikeloni",
			territory.PN: "Pitkairni",
			territory.PR: "Pwetoriko",
			territory.PS: "Ukingo wa Magharibi na Ukanda wa Gaza wa Palestina",
			territory.PT: "Ureno",
			territory.PW: "Palau",
			territory.PY: "Paragwai",
			territory.QA: "Katari",
			territory.RE: "Riyunioni",
			territory.RO: "Romania",
			territory.RU: "Urusi",
			territory.RW: "Rwanda",
			territory.SA: "Saudi",
			territory.SB: "Visiwa vya Solomon",
			territory.SC: "Shelisheli",
			territory.SD: "Sudani",
			territory.SE: "Uswidi",
			territory.SG: "Singapoo",
			territory.SH: "Santahelena",
			territory.SI: "Slovenia",
			territory.SK: "Slovakia",
			territory.SL: "Siera Leoni",
			territory.SM: "Samarino",
			territory.SN: "Senegali",
			territory.SO: "Somalia",
			territory.SR: "Surinamu",
			territory.ST: "Sao Tome na Principe",
			territory.SV: "Elsavado",
			territory.SY: "Siria",
			territory.SZ: "Uswazi",
			territory.TC: "Visiwa vya Turki na Kaiko",
			territory.TD: "Chadi",
			territory.TG: "Togo",
			territory.TH: "Tailandi",
			territory.TJ: "Tajikistani",
			territory.TK: "Tokelau",
			territory.TL: "Timori ya Mashariki",
			territory.TM: "Turukimenistani",
			territory.TN: "Tunisia",
			territory.TO: "Tonga",
			territory.TR: "Uturuki",
			territory.TT: "Trinidad na Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwani",
			territory.TZ: "Tanzania",
			territory.UA: "Ukraini",
			territory.UG: "Uganda",
			territory.US: "Marekani",
			territory.UY: "Urugwai",
			territory.UZ: "Uzibekistani",
			territory.VA: "Vatikani",
			territory.VC: "Santavisenti na Grenadini",
			territory.VE: "Venezuela",
			territory.VG: "Visiwa vya Virgin vya Uingereza",
			territory.VI: "Visiwa vya Virgin vya Marekani",
			territory.VN: "Vietinamu",
			territory.VU: "Vanuatu",
			territory.WF: "Walis na Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yemeni",
			territory.YT: "Mayotte",
			territory.ZA: "Afrika Kusini",
			territory.ZM: "Zambia",
			territory.ZW: "Zimbabwe",
		},
	}
}
