// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_asa() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "asa",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Januari", Feb: "Februari", Mar: "Machi", Apr: "Aprili", May: "Mei", Jun: "Juni", Jul: "Julai", Aug: "Agosti", Sep: "Septemba", Oct: "Oktoba", Nov: "Novemba", Dec: "Desemba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Jpi", Mon: "Jtt", Tue: "Jnn", Wed: "Jtn", Thu: "Alh", Fri: "Ijm", Sat: "Jmo"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "J", Mon: "J", Tue: "J", Wed: "J", Thu: "A", Fri: "I", Sat: "J"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Jumapili", Mon: "Jumatatu", Tue: "Jumanne", Wed: "Jumatano", Thu: "Alhamisi", Fri: "Ijumaa", Sat: "Jumamosi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "icheheavo", PM: "ichamthi"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "icheheavo", PM: "ichamthi"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "dirham ya Falme dha Kiarabu", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "kwandha ya Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "dola ya Authtralia", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "dinari ya Bahareni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "faranga ya Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "pula ya Botthwana", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "dola ya Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "faranga ya Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "faranga ya Uthwithi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "yuan renminbi ya China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "ethkudo ya Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "faranga ya Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "dinari ya Aljeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "pauni ya Mithri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa ya Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "bir ya Uhabeshi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "pauni ya Uingeredha", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "thedi ya Ghana", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "dalathi ya Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "faranga ya Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "rupia ya India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "tharafu ya Kijapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "shilingi ya Kenya", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "faranga ya Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "dola ya Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "loti ya Lethoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "dinari ya Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "dirham ya Moroko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "faranga ya Bukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "ugwiya ya Moritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ugwiya ya Moritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "rupia ya Morithi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha ya Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "metikali ya Mthumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "dola ya Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "naira ya Nijeria", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "faranga ya Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "riyal ya Thaudia", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "rupia ya Shelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "dinari ya Thudani", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "pauni ya Thudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "pauni ya Thantahelena", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "leoni", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "shilingi ya Thomalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "dobra ya Thao Tome na Principe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "dobra ya Thao Tome na Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "dinari ya Tunithia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "shilingi ya Tandhania", Symbol: "TSh"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "shilingi ya Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "dola ya Marekani", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "faranga CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "faranga CFA BCEAO", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "randi ya Afrika Kuthini", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "kwacha ya Dhambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha ya Dhambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "dola ya Dhimbabwe", Symbol: ""},
			},
		},
		Languages: cldr.Languages{
			language.AK:  "Kiakan",
			language.AM:  "Kiamhari",
			language.AR:  "Kiarabu",
			language.ASA: "Kipare",
			language.BE:  "Kibelarusi",
			language.BG:  "Kibulgaria",
			language.BN:  "Kibangla",
			language.CS:  "Kicheki",
			language.DE:  "Kijerumani",
			language.EL:  "Kigiriki",
			language.EN:  "Kiingeredha",
			language.ES:  "Kihithpania",
			language.FA:  "Kiajemi",
			language.FR:  "Kifarantha",
			language.HA:  "Kihautha",
			language.HI:  "Kihindi",
			language.HU:  "Kihungari",
			language.ID:  "Kiindonethia",
			language.IG:  "Kiigbo",
			language.IT:  "Kiitaliaano",
			language.JA:  "Kijapani",
			language.JV:  "Kijava",
			language.KM:  "Kikambodia",
			language.KO:  "Kikorea",
			language.MS:  "Kimalesia",
			language.MY:  "Kiburma",
			language.NE:  "Kinepali",
			language.NL:  "Kiholandhi",
			language.PA:  "Kipunjabi",
			language.PL:  "Kipolandi",
			language.PT:  "Kireno",
			language.RO:  "Kiromania",
			language.RU:  "Kiruthi",
			language.RW:  "Kinyarandwa",
			language.SO:  "Kithomali",
			language.SV:  "Kithwidi",
			language.TA:  "Kitamil",
			language.TH:  "Kitailandi",
			language.TR:  "Kituruki",
			language.UK:  "Kiukrania",
			language.UR:  "Kiurdu",
			language.VI:  "Kivietinamu",
			language.YO:  "Kiyoruba",
			language.ZH:  "Kichina",
			language.ZU:  "Kidhulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andora",
			territory.AE: "Falme dha Kiarabu",
			territory.AF: "Afuganistani",
			territory.AG: "Antigua na Barbuda",
			territory.AI: "Anguilla",
			territory.AL: "Albania",
			territory.AM: "Armenia",
			territory.AO: "Angola",
			territory.AR: "Ajentina",
			territory.AS: "Thamoa ya Marekani",
			territory.AT: "Authtria",
			territory.AU: "Authtralia",
			territory.AW: "Aruba",
			territory.AZ: "Adhabajani",
			territory.BA: "Bothnia na Hedhegovina",
			territory.BB: "Babadothi",
			territory.BD: "Bangladeshi",
			territory.BE: "Ubelgiji",
			territory.BF: "Bukinafatho",
			territory.BG: "Bulgaria",
			territory.BH: "Bahareni",
			territory.BI: "Burundi",
			territory.BJ: "Benini",
			territory.BM: "Bermuda",
			territory.BN: "Brunei",
			territory.BR: "Brazili",
			territory.BS: "Bahama",
			territory.BT: "Butani",
			territory.BW: "Botthwana",
			territory.BY: "Belaruthi",
			territory.BZ: "Belidhe",
			territory.CA: "Kanada",
			territory.CD: "Jamhuri ya Kidemokrathia ya Kongo",
			territory.CF: "Jamhuri ya Afrika ya Kati",
			territory.CG: "Kongo",
			territory.CH: "Uthwithi",
			territory.CI: "Kodivaa",
			territory.CK: "Vithiwa vya Cook",
			territory.CL: "Chile",
			territory.CM: "Kameruni",
			territory.CN: "China",
			territory.CO: "Kolombia",
			territory.CR: "Kothtarika",
			territory.CU: "Kuba",
			territory.CV: "Kepuvede",
			territory.CY: "Kuprothi",
			territory.CZ: "Jamhuri ya Cheki",
			territory.DE: "Ujerumani",
			territory.DJ: "Jibuti",
			territory.DK: "Denmaki",
			territory.DM: "Dominika",
			territory.DO: "Jamhuri ya Dominika",
			territory.DZ: "Aljeria",
			territory.EC: "Ekwado",
			territory.EE: "Ethtonia",
			territory.EG: "Mithri",
			territory.ER: "Eritrea",
			territory.ES: "Hithpania",
			territory.ET: "Uhabeshi",
			territory.FI: "Ufini",
			territory.FJ: "Fiji",
			territory.FK: "Vithiwa vya Falkland",
			territory.FM: "Mikronethia",
			territory.FR: "Ufarantha",
			territory.GA: "Gaboni",
			territory.GB: "Uingeredha",
			territory.GD: "Grenada",
			territory.GE: "Jojia",
			territory.GF: "Gwiyana ya Ufarantha",
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
			territory.HN: "Hondurathi",
			territory.HR: "Korathia",
			territory.HT: "Haiti",
			territory.HU: "Hungaria",
			territory.ID: "Indonethia",
			territory.IE: "Ayalandi",
			territory.IL: "Ithraeli",
			territory.IN: "India",
			territory.IO: "Ieneo la Uingeredha katika Bahari Hindi",
			territory.IQ: "Iraki",
			territory.IR: "Uajemi",
			territory.IS: "Aithlandi",
			territory.IT: "Italia",
			territory.JM: "Jamaika",
			territory.JO: "Yordani",
			territory.JP: "Japani",
			territory.KE: "Kenya",
			territory.KG: "Kirigizithtani",
			territory.KH: "Kambodia",
			territory.KI: "Kiribati",
			territory.KM: "Komoro",
			territory.KN: "Thantakitdhi na Nevith",
			territory.KP: "Korea Kathkazini",
			territory.KR: "Korea Kuthini",
			territory.KW: "Kuwaiti",
			territory.KY: "Vithiwa vya Kayman",
			territory.KZ: "Kazakithtani",
			territory.LA: "Laothi",
			territory.LB: "Lebanoni",
			territory.LC: "Thantaluthia",
			territory.LI: "Lishenteni",
			territory.LK: "Thirilanka",
			territory.LR: "Liberia",
			territory.LS: "Lethoto",
			territory.LT: "Litwania",
			territory.LU: "Lathembagi",
			territory.LV: "Lativia",
			territory.LY: "Libya",
			territory.MA: "Moroko",
			territory.MC: "Monako",
			territory.MD: "Moldova",
			territory.MG: "Bukini",
			territory.MH: "Vithiwa vya Marshal",
			territory.ML: "Mali",
			territory.MM: "Myama",
			territory.MN: "Mongolia",
			territory.MP: "Vithiwa vya Mariana vya Kathkazini",
			territory.MQ: "Martiniki",
			territory.MR: "Moritania",
			territory.MS: "Monttherrati",
			territory.MT: "Malta",
			territory.MU: "Morithi",
			territory.MV: "Modivu",
			territory.MW: "Malawi",
			territory.MX: "Mekthiko",
			territory.MY: "Malethia",
			territory.MZ: "Mthumbiji",
			territory.NA: "Namibia",
			territory.NC: "Nyukaledonia",
			territory.NE: "Nijeri",
			territory.NF: "Kithiwa cha Norfok",
			territory.NG: "Nijeria",
			territory.NI: "Nikaragwa",
			territory.NL: "Uholandhi",
			territory.NO: "Norwe",
			territory.NP: "Nepali",
			territory.NR: "Nauru",
			territory.NU: "Niue",
			territory.NZ: "Nyudhilandi",
			territory.OM: "Omani",
			territory.PA: "Panama",
			territory.PE: "Peru",
			territory.PF: "Polinesia ya Ufarantha",
			territory.PG: "Papua",
			territory.PH: "Filipino",
			territory.PK: "Pakithtani",
			territory.PL: "Polandi",
			territory.PM: "Thantapieri na Mikeloni",
			territory.PN: "Pitkairni",
			territory.PR: "Pwetoriko",
			territory.PS: "Palestina",
			territory.PT: "Ureno",
			territory.PW: "Palau",
			territory.PY: "Paragwai",
			territory.QA: "Katari",
			territory.RE: "Riyunioni",
			territory.RO: "Romania",
			territory.RU: "Uruthi",
			territory.RW: "Rwanda",
			territory.SA: "Thaudi",
			territory.SB: "Vithiwa vya Tholomon",
			territory.SC: "Shelisheli",
			territory.SD: "Thudani",
			territory.SE: "Uthwidi",
			territory.SG: "Thingapoo",
			territory.SH: "Thantahelena",
			territory.SI: "Thlovenia",
			territory.SK: "Tholvakia",
			territory.SL: "Thiera Leoni",
			territory.SM: "Thamarino",
			territory.SN: "Thenegali",
			territory.SO: "Thomalia",
			territory.SR: "Thurinamu",
			territory.ST: "Thao Tome na Principe",
			territory.SV: "Elsavado",
			territory.SY: "Thiria",
			territory.SZ: "Uthwadhi",
			territory.TC: "Vithiwa vya Turki na Kaiko",
			territory.TD: "Chadi",
			territory.TG: "Togo",
			territory.TH: "Tailandi",
			territory.TJ: "Tajikithtani",
			territory.TK: "Tokelau",
			territory.TL: "Timori ya Mashariki",
			territory.TM: "Turukimenithtani",
			territory.TN: "Tunithia",
			territory.TO: "Tonga",
			territory.TR: "Uturuki",
			territory.TT: "Trinidad na Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwani",
			territory.TZ: "Tadhania",
			territory.UG: "Uganda",
			territory.US: "Marekani",
			territory.UY: "Urugwai",
			territory.UZ: "Udhibekithtani",
			territory.VA: "Vatikani",
			territory.VC: "Thantavithenti na Grenadini",
			territory.VE: "Venezuela",
			territory.VG: "Vithiwa vya Virgin vya Uingeredha",
			territory.VI: "Vithiwa vya Virgin vya Marekani",
			territory.VN: "Vietinamu",
			territory.VU: "Vanuatu",
			territory.WF: "Walith na Futuna",
			territory.WS: "Thamoa",
			territory.YE: "Yemeni",
			territory.YT: "Mayotte",
			territory.ZA: "Afrika Kuthini",
			territory.ZM: "Dhambia",
			territory.ZW: "Dhimbabwe",
		},
	}
}
