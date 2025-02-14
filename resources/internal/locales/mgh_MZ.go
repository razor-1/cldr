// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_mgh_MZ() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "mgh_MZ",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Kwa", Feb: "Una", Mar: "Rar", Apr: "Che", May: "Tha", Jun: "Moc", Jul: "Sab", Aug: "Nan", Sep: "Tis", Oct: "Kum", Nov: "Moj", Dec: "Yel"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "K", Feb: "U", Mar: "R", Apr: "C", May: "T", Jun: "M", Jul: "S", Aug: "N", Sep: "T", Oct: "K", Nov: "M", Dec: "Y"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Mweri wo kwanza", Feb: "Mweri wo unayeli", Mar: "Mweri wo uneraru", Apr: "Mweri wo unecheshe", May: "Mweri wo unethanu", Jun: "Mweri wo thanu na mocha", Jul: "Mweri wo saba", Aug: "Mweri wo nane", Sep: "Mweri wo tisa", Oct: "Mweri wo kumi", Nov: "Mweri wo kumi na moja", Dec: "Mweri wo kumi na yel’li"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sab", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Ara", Fri: "Iju", Sat: "Jmo"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "J", Tue: "J", Wed: "J", Thu: "A", Fri: "I", Sat: "J"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sabato", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Arahamisi", Fri: "Ijumaa", Sat: "Jumamosi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "wichishu", PM: "mchochil’l"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
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
				currency.INR: cldr.Currency{DisplayName: "", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "", Symbol: "JP¥"},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
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
				currency.MZN: cldr.Currency{DisplayName: "", Symbol: "MTn"},
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
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
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
				currency.USD: cldr.Currency{DisplayName: "", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Ikan",
			language.AM:  "Imhari",
			language.AR:  "Iarabu",
			language.BE:  "Ibelausi",
			language.BG:  "Ibulgaria",
			language.BN:  "Ibangla",
			language.CS:  "Icheki",
			language.DE:  "Ijerimani",
			language.EL:  "Igiriki",
			language.EN:  "Ingilishi",
			language.ES:  "Ihispaniola",
			language.FA:  "Iajemi",
			language.FR:  "Ifaransa",
			language.HA:  "Ihausa",
			language.HI:  "Ihindi",
			language.HU:  "Ihungari",
			language.IG:  "Igbo",
			language.IT:  "Italiano",
			language.JA:  "Ijapani",
			language.JV:  "Ijava",
			language.KM:  "Ikambodia",
			language.KO:  "Ikorea",
			language.MGH: "Makua",
			language.MS:  "Imalesia",
			language.MY:  "Iburma",
			language.NE:  "Inepali",
			language.NL:  "Iholanzi",
			language.PA:  "Ipunjabi",
			language.PL:  "Ipolandi",
			language.PT:  "Nreno",
			language.RO:  "Iromania",
			language.RU:  "Irisi",
			language.RW:  "Inyaranda",
			language.SO:  "Isomali",
			language.SV:  "Iswidi",
			language.TA:  "Itamil",
			language.TH:  "Itailandi",
			language.TR:  "Ituruki",
			language.UK:  "Iukran",
			language.UR:  "Ihurdu",
			language.VI:  "Ivyetinamu",
			language.YO:  "Iyoruba",
			language.ZH:  "Ichina",
			language.ZU:  "Izulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Uandora",
			territory.AF: "Ufugustani",
			territory.AL: "Ualbania",
			territory.AS: "Usamoa ya Marekani",
			territory.AZ: "Uazabajani",
			territory.BI: "Urundi",
			territory.BJ: "Ubelin",
			territory.CA: "Ukanada",
			territory.CG: "Ukongo",
			territory.CH: "Uswisi",
			territory.CI: "Ukodiva",
			territory.CL: "Uchile",
			territory.CN: "Uchina",
			territory.CO: "Ukolombia",
			territory.CU: "Ukuba",
			territory.CY: "Ukuprosi",
			territory.CZ: "Ucheki",
			territory.DJ: "Ujibuti",
			territory.DK: "Udenimaka",
			territory.DM: "Udominika",
			territory.DZ: "Alujeria",
			territory.EG: "Umisiri",
			territory.ER: "Uriterea",
			territory.ES: "Uhispania",
			territory.ET: "Uhabeshi",
			territory.FI: "Ufini",
			territory.FJ: "Ufiji",
			territory.FR: "Ufaransa",
			territory.GA: "Ugaboni",
			territory.GD: "Ugrenada",
			territory.GE: "Ujojia",
			territory.GF: "Ufaransa yo Gwaya",
			territory.GH: "Ugana",
			territory.GI: "Ujibralta",
			territory.GL: "Ugrinlandi",
			territory.GM: "Ugambia",
			territory.GN: "Ugine",
			territory.GP: "Ugwadelupe",
			territory.GT: "Ugwatemala",
			territory.GU: "Ugwam",
			territory.GW: "Uginebisau",
			territory.GY: "Uguyana",
			territory.HN: "Uhondurasi",
			territory.HR: "Ukorasia",
			territory.HT: "Uhaiti",
			territory.HU: "Uhungaria",
			territory.ID: "Undonesia",
			territory.IE: "Uayalandi",
			territory.IL: "Uisraeli",
			territory.IN: "Uhindini",
			territory.IQ: "Wiraki",
			territory.IT: "Italia",
			territory.JM: "Ujamaika",
			territory.JO: "Uyordani",
			territory.JP: "Ujapani",
			territory.KE: "Ukenya",
			territory.KH: "Ukambodia",
			territory.KM: "Ukomoro",
			territory.KN: "Usantakitzi na Nevis",
			territory.KP: "Ukorea Kaskazini",
			territory.KR: "Ukorea Kusini",
			territory.KZ: "Ukazakistani",
			territory.LB: "Ulebanoni",
			territory.LC: "Usantalusia",
			territory.LI: "Ushenteni",
			territory.LK: "Usirilanka",
			territory.LR: "Uliberia",
			territory.LS: "Ulesoto",
			territory.LT: "Utwania",
			territory.LU: "Usembaji",
			territory.LV: "Ulativia",
			territory.LY: "Ulibya",
			territory.ME: "Umantegro",
			territory.MG: "Ubukini",
			territory.MW: "Umalawi",
			territory.MZ: "Umozambiki",
			territory.NE: "Unijeri",
			territory.NG: "Unijeria",
			territory.NO: "Unorwe",
			territory.OM: "Uomani",
			territory.PA: "Upanama",
			territory.PE: "Uperuu",
			territory.PF: "Ufaransa yo Potina",
			territory.PG: "Upapua",
			territory.PH: "Ufilipino",
			territory.PK: "Upakistani",
			territory.PL: "Upolandi",
			territory.PM: "Usantapieri na Mikeloni",
			territory.PN: "Upitkairni",
			territory.PR: "Upwetoriko",
			territory.PY: "Paragwai",
			territory.QA: "Ukatari",
			territory.RE: "Uriyunioni",
			territory.RO: "Uromania",
			territory.RW: "Urwanda",
			territory.SA: "Usaudi",
			territory.SC: "Ushelisheli",
			territory.SD: "Usudani",
			territory.SE: "Uswidi",
			territory.SG: "Usingapoo",
			territory.SH: "Usantahelena",
			territory.SI: "Uslovenia",
			territory.SK: "Uslovakia",
			territory.SM: "Usamarino",
			territory.SN: "Usenegali",
			territory.SO: "Usomalia",
			territory.SR: "Usurinamu",
			territory.ST: "Usao Tome na Principe",
			territory.SV: "Usalavado",
			territory.SY: "Usiria",
			territory.SZ: "Uswazi",
			territory.TD: "Uchadi",
			territory.TG: "Utogo",
			territory.TH: "Utailandi",
			territory.TJ: "Ujikistani",
			territory.TK: "Utokelau",
			territory.TL: "Utimo Mashariki",
			territory.TM: "Uturukimenistani",
			territory.TN: "Utunisia",
			territory.TO: "Utonga",
			territory.TR: "Utuki",
			territory.TT: "Utrinidad na Tobago",
			territory.TV: "Utuvalu",
			territory.TZ: "Utanzania",
			territory.US: "Umarekani",
			territory.VA: "Uvatikani",
			territory.VC: "Usantavisenti na Grenadini",
			territory.VE: "Uvenezuela",
			territory.VN: "Uvietinamu",
			territory.VU: "Uvanuatu",
			territory.WF: "Uwalis na Futuna",
			territory.WS: "Usamoa",
			territory.YE: "Uyemeni",
			territory.ZA: "Afrika du Sulu",
			territory.ZM: "Uzambia",
			territory.ZW: "Uzimbabwe",
		},
	}
}
