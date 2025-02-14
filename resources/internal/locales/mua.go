// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_mua() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "mua",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "FLO", Feb: "CLA", Mar: "CKI", Apr: "FMF", May: "MAD", Jun: "MBI", Jul: "MLI", Aug: "MAM", Sep: "FDE", Oct: "FMU", Nov: "FGW", Dec: "FYU"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "O", Feb: "A", Mar: "I", Apr: "F", May: "D", Jun: "B", Jul: "L", Aug: "M", Sep: "E", Oct: "U", Nov: "W", Dec: "Y"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Fĩi Loo", Feb: "Cokcwaklaŋne", Mar: "Cokcwaklii", Apr: "Fĩi Marfoo", May: "Madǝǝuutǝbijaŋ", Jun: "Mamǝŋgwãafahbii", Jul: "Mamǝŋgwãalii", Aug: "Madǝmbii", Sep: "Fĩi Dǝɓlii", Oct: "Fĩi Mundaŋ", Nov: "Fĩi Gwahlle", Dec: "Fĩi Yuru"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Cya", Mon: "Cla", Tue: "Czi", Wed: "Cko", Thu: "Cka", Fri: "Cga", Sat: "Cze"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Y", Mon: "L", Tue: "Z", Wed: "O", Thu: "A", Fri: "G", Sat: "E"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Com’yakke", Mon: "Comlaaɗii", Tue: "Comzyiiɗii", Wed: "Comkolle", Thu: "Comkaldǝɓlii", Fri: "Comgaisuu", Sat: "Comzyeɓsuu"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "comme", PM: "lilli"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Solai Arabiya", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "solai Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "solai Australya", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "solai Barenya", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "solai Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "solai Botswana", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "solai Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "solai Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Solai Swiss", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "solai Syiŋ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "solai Kapverdiya", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "solai Djibouti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "solai Algerya", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "solai Egypt", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "solai Eritre", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "solai Etiopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "solai Euro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "solai Britaniya", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "solai Gana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "solai Gambiya", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "solai Guine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "solai India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "solai Japoŋ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "solai Kenia", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "solai Komorya", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "solai Liberiya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "solai Lesotho", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "solai Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Solai Marok", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Solai Malagasya", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Solai Mauritaniya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Solai Mauritaniya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Solai Mauricǝ", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Solai Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Solai Mozambika", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Solai Namibiya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Solai Nigeriya", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Solai Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Solai Saudiya", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Solai Saichel", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Solai Sudaŋ ma dii ne dinar", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Solai Sudaŋ ma dii ne livre", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Solai Helena", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "solai Sierra leonǝ", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "solai Sierra leonǝ (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Solai Somaliya", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Solai Sao Tome (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Solai Sao Tome", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "solai Swaziland", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Solai Tunisiya", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Solai Tanzaniya", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Solai Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Solai Amerika", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "solai BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "solai BCEAO", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Solai Africa nekǝsǝŋ", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Solai Zambiya (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Solai Zambiya", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Solai Zimbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "akaŋ",
			language.AM:  "amharik",
			language.AR:  "arabiya",
			language.BE:  "belarussiya",
			language.BG:  "bulgaria",
			language.BN:  "bengalia",
			language.CS:  "syekya",
			language.DE:  "germaŋ",
			language.EL:  "grek",
			language.EN:  "zah Anglofoŋ",
			language.ES:  "Espaniya",
			language.FA:  "Persia",
			language.FR:  "zah sǝr Franssǝ",
			language.HA:  "haussa",
			language.HI:  "hindi",
			language.HU:  "hungariya",
			language.ID:  "indonesiya",
			language.IG:  "igbo",
			language.IT:  "italiya",
			language.JA:  "zah sǝr Japoŋ",
			language.JV:  "javaniya",
			language.KM:  "kmer",
			language.KO:  "korea",
			language.MS:  "malasiya",
			language.MUA: "MUNDAŊ",
			language.MY:  "birmania",
			language.NE:  "Nepaliya",
			language.NL:  "zah sǝr ma kasǝŋ",
			language.PA:  "Pǝnjabi",
			language.PL:  "Poloniya",
			language.PT:  "Zah sǝr Portugal",
			language.RO:  "Romaniya",
			language.RU:  "Russiya",
			language.RW:  "Zah sǝr Rwanda",
			language.SO:  "Somaliya",
			language.SV:  "Swedia",
			language.TA:  "Tamul",
			language.TR:  "Turk",
			language.UK:  "Ukrainia",
			language.UR:  "Urdu",
			language.VI:  "Vietnamiya",
			language.YO:  "Yoruba",
			language.ZH:  "zah Syiŋ",
			language.ZU:  "Zulu",
		},
		Territories: cldr.Territories{
			territory.AD: "andorra",
			territory.AE: "Sǝr Arabiya ma taini",
			territory.AF: "afghanistaŋ",
			territory.AG: "antiguan ne Barbuda",
			territory.AI: "anguiya",
			territory.AL: "albaniya",
			territory.AM: "armeniya",
			territory.AO: "angola",
			territory.AR: "argentiniya",
			territory.AS: "samoa Amerika",
			territory.AT: "austriya",
			territory.AU: "australiya",
			territory.AW: "aruba",
			territory.AZ: "azerbaijaŋ",
			territory.BA: "bosniya ne Herzegovina",
			territory.BB: "barbadiya",
			territory.BD: "bangladeshiya",
			territory.BE: "belgika",
			territory.BF: "burkina Faso",
			territory.BG: "bulgariya",
			territory.BH: "bahraiŋ",
			territory.BI: "burundi",
			territory.BJ: "beniŋ",
			territory.BM: "bermudiya",
			territory.BN: "bruniya",
			territory.BO: "boliviya",
			territory.BR: "brazilya",
			territory.BS: "bahamas",
			territory.BT: "butaŋ",
			territory.BW: "botswana",
			territory.BY: "belarussiya",
			territory.BZ: "beliziya",
			territory.CA: "kanada",
			territory.CD: "Sǝr Kongo ma dii ne zair",
			territory.CF: "centrafrika",
			territory.CG: "kongo",
			territory.CH: "Sǝr Swiss",
			territory.CI: "ser Ivoiriya",
			territory.CK: "kook ma laŋne",
			territory.CL: "syili",
			territory.CM: "kameruŋ",
			territory.CN: "syiŋ",
			territory.CO: "kolombiya",
			territory.CR: "kosta Rika",
			territory.CU: "Kuba",
			territory.CV: "kap ma laŋne",
			territory.CY: "Syipriya",
			territory.CZ: "Sǝr Syek",
			territory.DE: "Germaniya",
			territory.DJ: "Djibouti",
			territory.DK: "Daŋmark",
			territory.DM: "Dominik",
			territory.DO: "Sǝr Dominik ma lii",
			territory.DZ: "algeriya",
			territory.EC: "Ekwatǝr",
			territory.EE: "Estoniya",
			territory.EG: "Sǝr Egypt",
			territory.ER: "Sǝr Eritre",
			territory.ES: "Espaŋiya",
			territory.ET: "Etiopia",
			territory.FI: "Sǝr Finland",
			territory.FJ: "Sǝr Fiji",
			territory.FK: "Sǝr malouniya ma laŋne",
			territory.FM: "Micronesiya",
			territory.FR: "Franssǝ",
			territory.GA: "Gaboŋ",
			territory.GB: "Sǝr Anglofoŋ",
			territory.GD: "Grenadǝ",
			territory.GE: "Georgiya",
			territory.GF: "Sǝr Guyana ma Franssǝ",
			territory.GH: "Gana",
			territory.GI: "Sǝr Gibraltar",
			territory.GL: "Sǝr Groenland",
			territory.GM: "Gambiya",
			territory.GN: "Guine",
			territory.GP: "Sǝr Gwadeloupǝ",
			territory.GQ: "Sǝr Guine",
			territory.GR: "Sǝr Grek",
			territory.GT: "Gwatemala",
			territory.GU: "Gwam",
			territory.GW: "Guine ma Bissao",
			territory.GY: "Guyana",
			territory.HN: "Sǝr Honduras",
			territory.HR: "kroatiya",
			territory.HT: "Sǝr Haiti",
			territory.HU: "Hungriya",
			territory.ID: "Indonesiya",
			territory.IE: "Sǝr Ireland",
			territory.IL: "Sǝr Israel",
			territory.IN: "Sǝr Indǝ",
			territory.IO: "anglofoŋ ma Indiya",
			territory.IQ: "Irak",
			territory.IR: "Iraŋ",
			territory.IS: "Sǝr Island",
			territory.IT: "Italiya",
			territory.JM: "Jamaika",
			territory.JO: "Jordaniya",
			territory.JP: "Japaŋ",
			territory.KE: "Sǝr Kenya",
			territory.KG: "Kirgizstaŋ",
			territory.KH: "kambodiya",
			territory.KI: "Sǝr Kiribati",
			territory.KM: "komora",
			territory.KN: "Sǝr Kristof ne Nievǝ",
			territory.KP: "Sǝr Kore fah sǝŋ",
			territory.KR: "Sǝr Kore nekǝsǝŋ",
			territory.KW: "Sǝr Kowait",
			territory.KY: "kayman ma laŋne",
			territory.KZ: "Kazakstaŋ",
			territory.LA: "Sǝr Laos",
			territory.LB: "Libaŋ",
			territory.LC: "Sǝr Lucia",
			territory.LI: "Lichtǝnsteiŋ",
			territory.LK: "Sǝr Lanka",
			territory.LR: "Liberiya",
			territory.LS: "Sǝr Lesotho",
			territory.LT: "Lituaniya",
			territory.LU: "Sǝr Luxemburg",
			territory.LV: "Letoniya",
			territory.LY: "Libiya",
			territory.MA: "Marok",
			territory.MC: "Monako",
			territory.MD: "Moldoviya",
			territory.MG: "Madagaskar",
			territory.MH: "Sǝr Marshall ma laŋne",
			territory.ML: "Sǝr Mali",
			territory.MM: "Sǝr Myanmar",
			territory.MN: "Mongoliya",
			territory.MP: "Sǝr Maria ma laŋne",
			territory.MQ: "Martinika",
			territory.MR: "Mauritaniya",
			territory.MS: "Sǝr Montserrat",
			territory.MT: "Sǝr Malta",
			territory.MU: "Sǝr Mauricǝ",
			territory.MV: "Maldivǝ",
			territory.MW: "Sǝr Malawi",
			territory.MX: "Mexiko",
			territory.MY: "Malaysiya",
			territory.MZ: "Mozambika",
			territory.NA: "Namibiya",
			territory.NC: "Kaledoniya mafuu",
			territory.NE: "Sǝr Niger",
			territory.NF: "Norfolk ma laŋne",
			territory.NG: "Nigeriya",
			territory.NI: "Nikaragwa",
			territory.NL: "Sǝr ma kasǝŋ",
			territory.NO: "Norvegǝ",
			territory.NP: "Sǝr Nepal",
			territory.NR: "Sǝr Nauru",
			territory.NU: "Niwe",
			territory.NZ: "Zeland mafuu",
			territory.OM: "Omaŋ",
			territory.PA: "Sǝr Panama",
			territory.PE: "Peru",
			territory.PF: "Sǝr Polynesiya ma Franssǝ",
			territory.PG: "Papuasiya Guine mafuu",
			territory.PH: "Filipiŋ",
			territory.PK: "Pakistaŋ",
			territory.PL: "Pologŋ",
			territory.PM: "Sǝr Pǝtar ne Mikǝlon",
			territory.PN: "Pitkairn",
			territory.PR: "Porto Riko",
			territory.PS: "Sǝr Palestiniya",
			territory.PT: "Sǝr Portugal",
			territory.PW: "Sǝr Palau",
			territory.PY: "Paragwai",
			territory.QA: "Katar",
			territory.RE: "Sǝr Reunion",
			territory.RO: "Romaniya",
			territory.RU: "Russiya",
			territory.RW: "Rwanda",
			territory.SA: "Sǝr Arabiya",
			territory.SB: "Sǝr Salomon ma laŋne",
			territory.SC: "Saichel",
			territory.SD: "Sudaŋ",
			territory.SE: "Sǝr Sued",
			territory.SG: "Singapur",
			territory.SH: "Sǝr Helena",
			territory.SI: "Sloveniya",
			territory.SK: "Slovakiya",
			territory.SL: "Sierra Leonǝ",
			territory.SM: "Sǝr Marino",
			territory.SN: "Senegal",
			territory.SO: "Somaliya",
			territory.SR: "Sǝr Surinam",
			territory.ST: "Sao Tome ne Principe",
			territory.SV: "Sǝr Salvador",
			territory.SY: "Syria",
			territory.SZ: "Sǝr Swaziland",
			territory.TC: "Turkiya ne kaicos ma laŋne",
			territory.TD: "syad",
			territory.TG: "Sǝr Togo",
			territory.TH: "Tailand",
			territory.TJ: "Tajikistaŋ",
			territory.TK: "Sǝr Tokelau",
			territory.TL: "Timoriya",
			territory.TM: "Turkmenistaŋ",
			territory.TN: "Tunisiya",
			territory.TO: "Sǝr Tonga",
			territory.TR: "Turkiya",
			territory.TT: "Trinite ne Tobago",
			territory.TV: "Sǝr Tuvalu",
			territory.TW: "Taiwaŋ",
			territory.TZ: "Tanzaniya",
			territory.UA: "Ukraiŋ",
			territory.UG: "Uganda",
			territory.US: "Amerika",
			territory.UY: "Urugwai",
			territory.UZ: "Uzbekistaŋ",
			territory.VA: "Vaticaŋ",
			territory.VC: "Sǝr Vinceŋ ne Grenadiŋ",
			territory.VE: "Sǝr Venezuela",
			territory.VG: "ser Anglofon ma laŋne",
			territory.VI: "Sǝr amerika ma laŋne",
			territory.VN: "Sǝr Vietnam",
			territory.VU: "Sǝr Vanuatu",
			territory.WF: "Wallis ne Futuna",
			territory.WS: "Sǝr Samoa",
			territory.YE: "Yemeŋ",
			territory.YT: "Mayot",
			territory.ZA: "Afrika nekǝsǝŋ",
			territory.ZM: "Zambiya",
			territory.ZW: "Zimbabwe",
		},
	}
}
