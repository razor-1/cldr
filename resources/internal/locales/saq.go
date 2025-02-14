// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_saq() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "saq",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Obo", Feb: "Waa", Mar: "Oku", Apr: "Ong", May: "Ime", Jun: "Ile", Jul: "Sap", Aug: "Isi", Sep: "Saa", Oct: "Tom", Nov: "Tob", Dec: "Tow"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "O", Feb: "W", Mar: "O", Apr: "O", May: "I", Jun: "I", Jul: "S", Aug: "I", Sep: "S", Oct: "T", Nov: "T", Dec: "T"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Lapa le obo", Feb: "Lapa le waare", Mar: "Lapa le okuni", Apr: "Lapa le ong’wan", May: "Lapa le imet", Jun: "Lapa le ile", Jul: "Lapa le sapa", Aug: "Lapa le isiet", Sep: "Lapa le saal", Oct: "Lapa le tomon", Nov: "Lapa le tomon obo", Dec: "Lapa le tomon waare"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Are", Mon: "Kun", Tue: "Ong", Wed: "Ine", Thu: "Ile", Fri: "Sap", Sat: "Kwe"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "K", Tue: "O", Wed: "I", Thu: "I", Fri: "S", Sat: "K"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Mderot ee are", Mon: "Mderot ee kuni", Tue: "Mderot ee ong’wan", Wed: "Mderot ee inet", Thu: "Mderot ee ile", Fri: "Mderot ee sapa", Sat: "Mderot ee kwe"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Tesiran", PM: "Teipa"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Njilingi eel Falme za Kiarabu", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "Njilingi eel Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Njilingi eel Australia", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Njilingi eel Bahareni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Njilingi eel Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Njilingi eel Botswana", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Njilingi eel Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Njilingi eel Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Njilingi eel Uswisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Njilingi eel China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Njilingi eel Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Njilingi eel Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Njilingi eel Aljeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Njilingi eel Misri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Njilingi eel Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Njilingi eel Uhabeshi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Njilingi eel Uingereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Njilingi eel Ghana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Njilingi eel Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Njilingi eel Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Njilingi eel India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Njilingi eel Kijapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Njilingi eel Kenya", Symbol: "Ksh"},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Njilingi eel Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dola eel Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Njilingi eel Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Njilingi eel Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Njilingi eel Moroko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Njilingi eel Bukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Njilingi eel Moritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Njilingi eel Moritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Njilingi eel Morisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Njilingi eel Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Njilingi eel Msumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Njilingi eel Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Njilingi eel Nijeria", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Njilingi eel Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Njilingi eel Saudia", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Njilingi eel Shelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Paunt eel Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Paunt eel Santahelena", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Leoni", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leoni (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Njilingi eel Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Njilingi eel Sao Tome na Principe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Njilingi eel Sao Tome na Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Njilingi eel Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Njilingi eel Tanzania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Njilingi eel Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dola eel Marekani", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Njilingi eel CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Njilingi eel CFA BCEAO", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Njilingi eel Afrika Kusini", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Njilingi eel Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Njilingi eel Zambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dola eel Zimbabwe", Symbol: ""},
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
			language.SAQ: "Kisampur",
			language.SO:  "Kisomali",
			language.SV:  "Kiswidi",
			language.TA:  "Kitamil",
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
			territory.KE: "Kenya",
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
