// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_mg_MG() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "mg_MG",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "Mey", Jun: "Jon", Jul: "Jol", Aug: "Aog", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Janoary", Feb: "Febroary", Mar: "Martsa", Apr: "Aprily", May: "Mey", Jun: "Jona", Jul: "Jolay", Aug: "Aogositra", Sep: "Septambra", Oct: "Oktobra", Nov: "Novambra", Dec: "Desambra"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Alah", Mon: "Alats", Tue: "Tal", Wed: "Alar", Thu: "Alak", Fri: "Zom", Sat: "Asab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "A", Mon: "A", Tue: "T", Wed: "A", Thu: "A", Fri: "Z", Sat: "A"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Alah", Mon: "Alats", Tue: "Tal", Wed: "Alar", Thu: "Alak", Fri: "Zom", Sat: "Asab"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Alahady", Mon: "Alatsinainy", Tue: "Talata", Wed: "Alarobia", Thu: "Alakamisy", Fri: "Zoma", Sat: "Asabotsy"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", CurrencyAccounting: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angoley", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dolara aostralianina", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "dinar bahreïni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Farantsa Borondi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pola botsoaney", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dôlara Kanadianina", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Farantsa kôngôley", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Farantsa soisa", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yoan sinoa Renminbi", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Farantsa Djibotianina", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinara alzerianina", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "vola venty ejipsiana", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfà Eritreanina", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Bir etiopianina", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Eoro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "livre sterling", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "cédi", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi gambianina", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Farantsa Gineanina", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Ropia Indianina", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yen Japoney", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilling kenianina", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Farantsa Komorianina", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dôlara Liberianina", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinara Libyanina", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirham marokianina", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariary", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya moritanianina (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya moritanianina", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Ropia maorisianina", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha malawite", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikaly", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dolara namibianina", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira nigerianina", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Farantsa Roande", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Rial saodianina", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Ropia Seysheloà", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dinara Sodaney", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "livre soudanaise (1956–2007)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "livre de Sainte-Hélène", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilling somalianina", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinar tonizianina", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilling tanzanianina", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilling ogandianina", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dolara amerikanina", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Farantsa CFA (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Farantsa CFA (BCEAO)", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Rand afrikanina tatsimo", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha zambianina (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha zambianina", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dôlara Zimbaboeanina", Symbol: ""},
			},
		},
		Languages: cldr.Languages{
			language.AK: "Akan",
			language.AM: "Amharika",
			language.AR: "Arabo",
			language.BE: "Bielorosy",
			language.BG: "Biolgara",
			language.BN: "Bengali",
			language.CS: "Tseky",
			language.DE: "Alemanina",
			language.EL: "Grika",
			language.EN: "Anglisy",
			language.ES: "Espaniola",
			language.FA: "Persa",
			language.FR: "Frantsay",
			language.HA: "haoussa",
			language.HI: "hindi",
			language.HU: "hongroà",
			language.ID: "Indonezianina",
			language.IG: "igbo",
			language.IT: "Italianina",
			language.JA: "Japoney",
			language.JV: "Javaney",
			language.KM: "khmer",
			language.KO: "Koreanina",
			language.MG: "Malagasy",
			language.MS: "Malay",
			language.MY: "Birmana",
			language.NE: "Nepale",
			language.NL: "Holandey",
			language.PA: "Penjabi",
			language.PL: "Poloney",
			language.PT: "Portiogey",
			language.RO: "Romanianina",
			language.RU: "Rosianina",
			language.RW: "Roande",
			language.SO: "Somalianina",
			language.SV: "Soisa",
			language.TA: "Tamoila",
			language.TH: "Taioaney",
			language.TR: "Tiorka",
			language.UK: "Okrainianina",
			language.UR: "Ordò",
			language.VI: "Vietnamianina",
			language.YO: "Yôrobà",
			language.ZH: "Sinoa, Mandarin",
			language.ZU: "Zolò",
		},
		Territories: cldr.Territories{
			territory.AD: "Andorra",
			territory.AE: "Emirà Arabo mitambatra",
			territory.AF: "Afghanistan",
			territory.AG: "Antiga sy Barboda",
			territory.AI: "Anguilla",
			territory.AL: "Albania",
			territory.AM: "Armenia",
			territory.AO: "Angola",
			territory.AR: "Arzantina",
			territory.AS: "Samoa amerikanina",
			territory.AT: "Aotrisy",
			territory.AU: "Aostralia",
			territory.AW: "Arobà",
			territory.AZ: "Azerbaidjan",
			territory.BA: "Bosnia sy Herzegovina",
			territory.BB: "Barbady",
			territory.BD: "Bangladesy",
			territory.BE: "Belzika",
			territory.BF: "Borkina Faso",
			territory.BG: "Biolgaria",
			territory.BH: "Bahrain",
			territory.BI: "Borondi",
			territory.BJ: "Benin",
			territory.BM: "Bermioda",
			territory.BN: "Brunei",
			territory.BO: "Bolivia",
			territory.BR: "Brezila",
			territory.BS: "Bahamas",
			territory.BT: "Bhotana",
			territory.BW: "Botsoana",
			territory.BY: "Belarosy",
			territory.BZ: "Belize",
			territory.CA: "Kanada",
			territory.CD: "Repoblikan’i Kongo",
			territory.CF: "Repoblika Ivon’Afrika",
			territory.CG: "Kôngô",
			territory.CH: "Soisa",
			territory.CI: "Côte d’Ivoire",
			territory.CK: "Nosy Kook",
			territory.CL: "Shili",
			territory.CM: "Kamerona",
			territory.CN: "Sina",
			territory.CO: "Kôlômbia",
			territory.CR: "Kosta Rikà",
			territory.CU: "Kiobà",
			territory.CV: "Nosy Cap-Vert",
			territory.CY: "Sypra",
			territory.CZ: "Repoblikan’i Tseky",
			territory.DE: "Alemaina",
			territory.DJ: "Djiboti",
			territory.DK: "Danmarka",
			territory.DM: "Dominika",
			territory.DO: "Repoblika Dominikanina",
			territory.DZ: "Alzeria",
			territory.EC: "Ekoatera",
			territory.EE: "Estonia",
			territory.EG: "Ejypta",
			territory.ER: "Eritrea",
			territory.ES: "Espaina",
			territory.ET: "Ethiopia",
			territory.FI: "Finlandy",
			territory.FJ: "Fidji",
			territory.FK: "Nosy Falkand",
			territory.FM: "Mikrônezia",
			territory.FR: "Frantsa",
			territory.GA: "Gabon",
			territory.GB: "Angletera",
			territory.GD: "Grenady",
			territory.GE: "Zeorzia",
			territory.GF: "Guyana frantsay",
			territory.GH: "Ghana",
			territory.GI: "Zibraltara",
			territory.GL: "Groenland",
			territory.GM: "Gambia",
			territory.GN: "Ginea",
			territory.GP: "Goadelopy",
			territory.GQ: "Guinea Ekoatera",
			territory.GR: "Gresy",
			territory.GT: "Goatemalà",
			territory.GU: "Guam",
			territory.GW: "Giné-Bisao",
			territory.GY: "Guyana",
			territory.HN: "Hondiorasy",
			territory.HR: "Kroasia",
			territory.HT: "Haiti",
			territory.HU: "Hongria",
			territory.ID: "Indonezia",
			territory.IE: "Irlandy",
			territory.IL: "Israely",
			territory.IN: "Indy",
			territory.IO: "Faridranomasina indiana britanika",
			territory.IQ: "Irak",
			territory.IR: "Iran",
			territory.IS: "Islandy",
			territory.IT: "Italia",
			territory.JM: "Jamaïka",
			territory.JO: "Jordania",
			territory.JP: "Japana",
			territory.KE: "Kenya",
			territory.KG: "Kiordistan",
			territory.KH: "Kambôdja",
			territory.KI: "Kiribati",
			territory.KM: "Kômaoro",
			territory.KN: "Saint-Christophe-et-Niévès",
			territory.KP: "Korea Avaratra",
			territory.KR: "Korea Atsimo",
			territory.KW: "Kôeity",
			territory.KY: "Nosy Kayman",
			territory.KZ: "Kazakhstan",
			territory.LA: "Laôs",
			territory.LB: "Libana",
			territory.LC: "Sainte-Lucie",
			territory.LI: "Listenstein",
			territory.LK: "Sri Lanka",
			territory.LR: "Liberia",
			territory.LS: "Lesotho",
			territory.LT: "Litoania",
			territory.LU: "Lioksamboro",
			territory.LV: "Letonia",
			territory.LY: "Libya",
			territory.MA: "Marôka",
			territory.MC: "Mônakô",
			territory.MD: "Môldavia",
			territory.MG: "Madagasikara",
			territory.MH: "Nosy Marshall",
			territory.ML: "Mali",
			territory.MM: "Myanmar",
			territory.MN: "Môngôlia",
			territory.MP: "Nosy Mariana Atsinanana",
			territory.MQ: "Martinika",
			territory.MR: "Maoritania",
			territory.MS: "Montserrat",
			territory.MT: "Malta",
			territory.MU: "Maorisy",
			territory.MV: "Maldiva",
			territory.MW: "Malaoì",
			territory.MX: "Meksika",
			territory.MY: "Malaizia",
			territory.MZ: "Mozambika",
			territory.NA: "Namibia",
			territory.NC: "Nouvelle-Calédonie",
			territory.NE: "Niger",
			territory.NF: "Nosy Norfolk",
			territory.NG: "Nizeria",
			territory.NI: "Nikaragoà",
			territory.NL: "Holanda",
			territory.NO: "Nôrvezy",
			territory.NP: "Nepala",
			territory.NR: "Naorò",
			territory.NU: "Nioé",
			territory.NZ: "Nouvelle-Zélande",
			territory.OM: "Oman",
			territory.PA: "Panama",
			territory.PE: "Peroa",
			territory.PF: "Polynezia frantsay",
			territory.PG: "Papouasie-Nouvelle-Guinée",
			territory.PH: "Filipina",
			territory.PK: "Pakistan",
			territory.PL: "Pôlôna",
			territory.PM: "Saint-Pierre-et-Miquelon",
			territory.PN: "Pitkairn",
			territory.PR: "Pôrtô Rikô",
			territory.PS: "Palestina",
			territory.PT: "Pôrtiogala",
			territory.PW: "Palao",
			territory.PY: "Paragoay",
			territory.QA: "Katar",
			territory.RE: "Larenion",
			territory.RO: "Romania",
			territory.RU: "Rosia",
			territory.RW: "Roanda",
			territory.SA: "Arabia saodita",
			territory.SB: "Nosy Salomona",
			territory.SC: "Seyshela",
			territory.SD: "Sodan",
			territory.SE: "Soedy",
			territory.SG: "Singaporo",
			territory.SH: "Sainte-Hélène",
			territory.SI: "Slovenia",
			territory.SK: "Slovakia",
			territory.SL: "Sierra Leone",
			territory.SM: "Saint-Marin",
			territory.SN: "Senegal",
			territory.SO: "Somalia",
			territory.SR: "Sorinam",
			territory.ST: "São Tomé-et-Príncipe",
			territory.SV: "El Salvador",
			territory.SY: "Syria",
			territory.SZ: "Soazilandy",
			territory.TC: "Nosy Turks sy Caïques",
			territory.TD: "Tsady",
			territory.TG: "Togo",
			territory.TH: "Thailandy",
			territory.TJ: "Tajikistan",
			territory.TK: "Tokelao",
			territory.TL: "Timor Atsinanana",
			territory.TM: "Torkmenistan",
			territory.TN: "Tonizia",
			territory.TO: "Tongà",
			territory.TR: "Torkia",
			territory.TT: "Trinidad sy Tobagô",
			territory.TV: "Tovalò",
			territory.TW: "Taioana",
			territory.TZ: "Tanzania",
			territory.UA: "Okraina",
			territory.UG: "Oganda",
			territory.US: "Etazonia",
			territory.UY: "Orogoay",
			territory.UZ: "Ozbekistan",
			territory.VA: "Firenen’i Vatikana",
			territory.VC: "Saint-Vincent-et-les Grenadines",
			territory.VE: "Venezoelà",
			territory.VG: "Nosy britanika virijiny",
			territory.VI: "Nosy Virijiny Etazonia",
			territory.VN: "Vietnam",
			territory.VU: "Vanoatò",
			territory.WF: "Wallis sy Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yemen",
			territory.YT: "Mayôty",
			territory.ZA: "Afrika Atsimo",
			territory.ZM: "Zambia",
			territory.ZW: "Zimbaboe",
		},
	}
}
