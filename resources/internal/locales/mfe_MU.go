// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_mfe_MU() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "mfe_MU",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "zan", Feb: "fev", Mar: "mar", Apr: "avr", May: "me", Jun: "zin", Jul: "zil", Aug: "out", Sep: "sep", Oct: "okt", Nov: "nov", Dec: "des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "z", Feb: "f", Mar: "m", Apr: "a", May: "m", Jun: "z", Jul: "z", Aug: "o", Sep: "s", Oct: "o", Nov: "n", Dec: "d"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "zanvie", Feb: "fevriye", Mar: "mars", Apr: "avril", May: "me", Jun: "zin", Jul: "zilye", Aug: "out", Sep: "septam", Oct: "oktob", Nov: "novam", Dec: "desam"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dim", Mon: "lin", Tue: "mar", Wed: "mer", Thu: "ze", Fri: "van", Sat: "sam"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "d", Mon: "l", Tue: "m", Wed: "m", Thu: "z", Fri: "v", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "dimans", Mon: "lindi", Tue: "mardi", Wed: "merkredi", Thu: "zedi", Fri: "vandredi", Sat: "samdi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "dirham Emira arab ini", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angole", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "dolar ostralien", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "dinar bahreïn", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "fran burunde", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "pula ya botswane", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "dolar kanadien", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "fran kongole", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "fran swis", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "yuan renminbi sinwa", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "eskudo kapverdien", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "fran djiboutien", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "dinar alzerien", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "liv ezipsien", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "nafka erythreen", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr etiopien", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "liv sterlin", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "sedi ganeen", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi gambien", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "fran gineen", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "roupi", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "yen zapone", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "shiling kenyan", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "fran komorien", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "dolar liberien", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "loti lezoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "dinar libien", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "dirham marokin", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "fran malgas", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "ouguiya moritanien (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ouguiya moritanien", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "roupi morisien", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha malawit", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "metikal mozanbikin", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "dolar namibien", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "naira nizerian", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "fran rwande", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "rial saoudien", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "roupi seselwa", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "dinar soudane", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "liv soudane", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "liv Sainte-Hélène", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "leonn Sierra-Leone", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "leonn Sierra-Leone (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "shilingi somalien", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "dobra santomeen (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "dobra santomeen", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni swazi", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "dinar tinizien", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "shiling tanzanien", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "shiling ougande", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "dolar amerikin", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "fran CFA (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "fran CFA (BCEAO)", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "rand sid-afrikin", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "kwacha zanbien (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha zanbien", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "dolar zimbawe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "akan",
			language.AM:  "amarik",
			language.AR:  "arab",
			language.BE:  "bieloris",
			language.BG:  "bilgar",
			language.BN:  "bengali",
			language.CS:  "tchek",
			language.DE:  "alman",
			language.EL:  "grek",
			language.EN:  "angle",
			language.ES:  "espagnol",
			language.FA:  "persan",
			language.FR:  "franse",
			language.HA:  "haoussa",
			language.HI:  "hindi",
			language.HU:  "hongrwa",
			language.ID:  "indonezien",
			language.IG:  "igbo",
			language.IT:  "italien",
			language.JA:  "zapone",
			language.JV:  "zavane",
			language.KM:  "khmer, santral",
			language.KO:  "koreen",
			language.MFE: "kreol morisien",
			language.MS:  "male",
			language.MY:  "birman",
			language.NE:  "nepale",
			language.NL:  "olande",
			language.PA:  "penjabi",
			language.PL:  "polone",
			language.PT:  "portige",
			language.RO:  "roumin",
			language.RU:  "ris",
			language.RW:  "rwanda",
			language.SO:  "somali",
			language.SV:  "swedwa",
			language.TA:  "tamoul",
			language.TH:  "thaï",
			language.TR:  "tirk",
			language.UK:  "ikrenien",
			language.UR:  "ourdou",
			language.VI:  "vietnamien",
			language.YO:  "yoruba",
			language.ZH:  "sinwa, mandarin",
			language.ZU:  "zoulou",
		},
		Territories: cldr.Territories{
			territory.AD: "Andor",
			territory.AE: "Emira arab ini",
			territory.AF: "Afganistan",
			territory.AG: "Antigua-ek-Barbuda",
			territory.AI: "Anguilla",
			territory.AL: "Albani",
			territory.AM: "Armeni",
			territory.AO: "Angola",
			territory.AR: "Larzantinn",
			territory.AS: "Samoa amerikin",
			territory.AT: "Lostris",
			territory.AU: "Lostrali",
			territory.AW: "Aruba",
			territory.AZ: "Azerbaïdjan",
			territory.BA: "Bosni-Herzegovinn",
			territory.BB: "Barbad",
			territory.BD: "Banglades",
			territory.BE: "Belzik",
			territory.BF: "Burkina Faso",
			territory.BG: "Bilgari",
			territory.BH: "Bahreïn",
			territory.BI: "Burundi",
			territory.BJ: "Benin",
			territory.BM: "Bermid",
			territory.BN: "Brunei",
			territory.BO: "Bolivi",
			territory.BR: "Brezil",
			territory.BS: "Bahamas",
			territory.BT: "Boutan",
			territory.BW: "Botswana",
			territory.BY: "Belaris",
			territory.BZ: "Beliz",
			territory.CA: "Kanada",
			territory.CD: "Repiblik demokratik Kongo",
			territory.CF: "Repiblik Lafrik Santral",
			territory.CG: "Kongo",
			territory.CH: "Laswis",
			territory.CI: "Côte d’Ivoire",
			territory.CK: "Zil Cook",
			territory.CL: "Shili",
			territory.CM: "Kamerounn",
			territory.CN: "Lasinn",
			territory.CO: "Kolonbi",
			territory.CR: "Costa Rica",
			territory.CU: "Cuba",
			territory.CV: "Kap-Ver",
			territory.CY: "Cyprus",
			territory.CZ: "Repiblik Chek",
			territory.DE: "Almagn",
			territory.DJ: "Djibouti",
			territory.DK: "Dannmark",
			territory.DM: "Dominik",
			territory.DO: "Repiblik dominikin",
			territory.DZ: "Alzeri",
			territory.EC: "Ekwater",
			territory.EE: "Estoni",
			territory.EG: "Lezipt",
			territory.ER: "Erythre",
			territory.ES: "Lespagn",
			territory.ET: "Letiopi",
			territory.FI: "Finland",
			territory.FJ: "Fidji",
			territory.FK: "Zil malwinn",
			territory.FM: "Mikronezi",
			territory.FR: "Lafrans",
			territory.GA: "Gabon",
			territory.GB: "United Kingdom",
			territory.GD: "Grenad",
			territory.GE: "Zeorzi",
			territory.GF: "Gwiyann franse",
			territory.GH: "Ghana",
			territory.GI: "Zibraltar",
			territory.GL: "Greenland",
			territory.GM: "Gambi",
			territory.GN: "Gine",
			territory.GP: "Guadloup",
			territory.GQ: "Gine ekwatoryal",
			territory.GR: "Gres",
			territory.GT: "Guatemala",
			territory.GU: "Guam",
			territory.GW: "Gine-Bisau",
			territory.GY: "Guyana",
			territory.HN: "Honduras",
			territory.HR: "Kroasi",
			territory.HT: "Ayti",
			territory.HU: "Ongri",
			territory.ID: "Indonezi",
			territory.IE: "Irland",
			territory.IL: "Izrael",
			territory.IN: "Lenn",
			territory.IO: "Teritwar Britanik Losean Indien",
			territory.IQ: "Irak",
			territory.IR: "Iran",
			territory.IS: "Island",
			territory.IT: "Itali",
			territory.JM: "Zamaik",
			territory.JO: "Zordani",
			territory.JP: "Zapon",
			territory.KE: "Kenya",
			territory.KG: "Kirghizistan",
			territory.KH: "Kambodj",
			territory.KI: "Kiribati",
			territory.KM: "Komor",
			territory.KN: "Saint-Christophe-ek-Niévès",
			territory.KP: "Lakore-dinor",
			territory.KR: "Lakore-disid",
			territory.KW: "Koweit",
			territory.KY: "Zil Kayman",
			territory.KZ: "Kazakstan",
			territory.LA: "Laos",
			territory.LB: "Liban",
			territory.LC: "Sainte-Lucie",
			territory.LI: "Liechtenstein",
			territory.LK: "Sri Lanka",
			territory.LR: "Liberia",
			territory.LS: "Lezoto",
			territory.LT: "Lituani",
			territory.LU: "Luxembourg",
			territory.LV: "Letoni",
			territory.LY: "Libi",
			territory.MA: "Marok",
			territory.MC: "Monako",
			territory.MD: "Moldavi",
			territory.MG: "Madagaskar",
			territory.MH: "Zil Marshall",
			territory.ML: "Mali",
			territory.MM: "Myanmar",
			territory.MN: "Mongoli",
			territory.MP: "Zil Maryann dinor",
			territory.MQ: "Martinik",
			territory.MR: "Moritani",
			territory.MS: "Montsera",
			territory.MT: "Malt",
			territory.MU: "Moris",
			territory.MV: "Maldiv",
			territory.MW: "Malawi",
			territory.MX: "Mexik",
			territory.MY: "Malezi",
			territory.MZ: "Mozambik",
			territory.NA: "Namibi",
			territory.NC: "Nouvel-Kaledoni",
			territory.NE: "Nizer",
			territory.NF: "Lil Norfolk",
			territory.NG: "Nizeria",
			territory.NI: "Nicaragua",
			territory.NL: "Oland",
			territory.NO: "Norvez",
			territory.NP: "Nepal",
			territory.NR: "Nauru",
			territory.NU: "Niowe",
			territory.NZ: "Nouvel Zeland",
			territory.OM: "Oman",
			territory.PA: "Panama",
			territory.PE: "Perou",
			territory.PF: "Polinezi franse",
			territory.PG: "Papouazi-Nouvel-Gine",
			territory.PH: "Filipinn",
			territory.PK: "Pakistan",
			territory.PL: "Pologn",
			territory.PM: "Saint-Pierre-ek-Miquelon",
			territory.PN: "Pitcairn",
			territory.PR: "Porto Rico",
			territory.PS: "Teritwar Palestinn",
			territory.PT: "Portigal",
			territory.PW: "Palau",
			territory.PY: "Paraguay",
			territory.QA: "Katar",
			territory.RE: "Larenion",
			territory.RO: "Roumani",
			territory.RU: "Larisi",
			territory.RW: "Rwanda",
			territory.SA: "Larabi Saoudit",
			territory.SB: "Zil Salomon",
			territory.SC: "Sesel",
			territory.SD: "Soudan",
			territory.SE: "Laswed",
			territory.SG: "Singapour",
			territory.SH: "Sainte-Hélène",
			territory.SI: "Sloveni",
			territory.SK: "Slovaki",
			territory.SL: "Sierra Leone",
			territory.SM: "Saint-Marin",
			territory.SN: "Senegal",
			territory.SO: "Somali",
			territory.SR: "Surinam",
			territory.ST: "São Tome-ek-Prínsip",
			territory.SV: "Salvador",
			territory.SY: "Lasiri",
			territory.SZ: "Swaziland",
			territory.TC: "Zil Tirk ek Caïcos",
			territory.TD: "Tchad",
			territory.TG: "Togo",
			territory.TH: "Thayland",
			territory.TJ: "Tadjikistan",
			territory.TK: "Tokelau",
			territory.TL: "Timor oriantal",
			territory.TM: "Turkmenistan",
			territory.TN: "Tinizi",
			territory.TO: "Tonga",
			territory.TR: "Tirki",
			territory.TT: "Trinite-ek-Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwan",
			territory.TZ: "Tanzani",
			territory.UA: "Ikrenn",
			territory.UG: "Ouganda",
			territory.US: "Lamerik",
			territory.UY: "Uruguay",
			territory.UZ: "Ouzbekistan",
			territory.VA: "Lata Vatikan",
			territory.VC: "Saint-Vincent-ek-Grenadines",
			territory.VE: "Venezuela",
			territory.VG: "Zil vierz britanik",
			territory.VI: "Zil Vierz Lamerik",
			territory.VN: "Vietnam",
			territory.VU: "Vanuatu",
			territory.WF: "Wallis-ek-Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yemenn",
			territory.YT: "Mayot",
			territory.ZA: "Sid-Afrik",
			territory.ZM: "Zambi",
			territory.ZW: "Zimbabwe",
		},
	}
}
