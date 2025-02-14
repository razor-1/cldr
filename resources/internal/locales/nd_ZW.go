// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_nd_ZW() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "nd_ZW",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Zib", Feb: "Nhlo", Mar: "Mbi", Apr: "Mab", May: "Nkw", Jun: "Nhla", Jul: "Ntu", Aug: "Ncw", Sep: "Mpan", Oct: "Mfu", Nov: "Lwe", Dec: "Mpal"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Z", Feb: "N", Mar: "M", Apr: "M", May: "N", Jun: "N", Jul: "N", Aug: "N", Sep: "M", Oct: "M", Nov: "L", Dec: "M"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Zibandlela", Feb: "Nhlolanja", Mar: "Mbimbitho", Apr: "Mabasa", May: "Nkwenkwezi", Jun: "Nhlangula", Jul: "Ntulikazi", Aug: "Ncwabakazi", Sep: "Mpandula", Oct: "Mfumfu", Nov: "Lwezi", Dec: "Mpalakazi"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Son", Mon: "Mvu", Tue: "Sib", Wed: "Sit", Thu: "Sin", Fri: "Sih", Sat: "Mgq"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "S", Wed: "S", Thu: "S", Fri: "S", Sat: "M"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sonto", Mon: "Mvulo", Tue: "Sibili", Wed: "Sithathu", Thu: "Sine", Fri: "Sihlanu", Sat: "Mgqibelo"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dola laseArab", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza yase Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dola yase Australia", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinari yase Bhahareni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Fulenki yase Bhurundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Phula yase Botswana", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dola yase Khanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Fulenki yase Khongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Fulenki yase Swisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Renminbi yase China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Caboverdiano", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Fulenki yase Jibhuthi", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinali yase Aljeriya", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Phawundi laseGibhide", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa yase Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Dola laseEthiopia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Phawundi yase Ngilandi", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi yase Ghana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi yase Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Fulenki yase Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupi yase Indiya", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yeni yase Japhani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilingi yase Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Fulenki yase Khomoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dola yase Libheriya", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lothi yase Lesotho", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinari yase Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirham yase Morokho", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Fulenki yase Malagasi", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugwiya yase Moritaniya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugwiya yase Moritaniya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupi yase Morishasi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha yase Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikali yase Mozambiki", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dola yase Namibiya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nayira yase Nijeriya", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Fulenki yase Ruwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal yase Saudi", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupi yase Seyisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dinari yase Sudani", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Phawundi yase Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Phawundindi laseSt Helena", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Leyoni", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leyoni (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilingi yase Somaliya", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra yase Sao Tome lo Principe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra yase Sao Tome lo Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinari yase Tunisiya", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingi yase Tanzaniya", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingi yase Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dola yase Amelika", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Fulenki CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Fulenki CFA BCEAO", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Randi yase Afrika ye Zanzi", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha yase Zambiya (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha yase Zambiya", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dola yase Zimbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK: "isi-Akhani",
			language.AM: "isi-Amaharikhi",
			language.AR: "isi-Alabhu",
			language.BE: "isi-Bhelarashiyani",
			language.BG: "isi-Bulgaria",
			language.BN: "isi-Bhengali",
			language.CS: "isi-Czech",
			language.DE: "isi-Jalimani",
			language.EL: "isi-Giliki",
			language.EN: "isi-Ngisi",
			language.ES: "isi-Sipeyini",
			language.FA: "isi-Pheshiyani",
			language.FR: "isi-Fulentshi",
			language.HA: "isi-Hausa",
			language.HI: "isi-Hindi",
			language.HU: "isi-Hangari",
			language.ID: "isi-Indonesia",
			language.IG: "isi-Igbo",
			language.IT: "isi-Italiano",
			language.JA: "isi-Japhani",
			language.JV: "isi-Java",
			language.KM: "isi-Khambodiya",
			language.KO: "isi-Koriya",
			language.MS: "isi-Malayi",
			language.MY: "isi-Burma",
			language.ND: "isiNdebele",
			language.NE: "isi-Nepali",
			language.NL: "isi-Dutch",
			language.PA: "isi-Phunjabi",
			language.PL: "isi-Pholoshi",
			language.PT: "isi-Potukezi",
			language.RO: "isi-Romani",
			language.RU: "isi-Rashiya",
			language.RW: "isi-Ruwanda",
			language.SO: "isi-Somali",
			language.SV: "isi-Swidishi",
			language.TA: "isi-Thamil",
			language.TH: "isi-Thayi",
			language.TR: "isi-Thekishi",
			language.UK: "isi-Ukrain",
			language.UR: "isi-Udu",
			language.VI: "isi-Vietnamese",
			language.YO: "isi-Yorubha",
			language.ZH: "isi-China",
			language.ZU: "isi-Zulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andora",
			territory.AE: "United Arab Emirates",
			territory.AF: "Afghanistan",
			territory.AG: "Antigua le Barbuda",
			territory.AI: "Anguilla",
			territory.AL: "Albania",
			territory.AM: "Armenia",
			territory.AO: "Angola",
			territory.AR: "Ajentina",
			territory.AS: "Samoa ye Amelika",
			territory.AT: "Austria",
			territory.AU: "Australia",
			territory.AW: "Arubha",
			territory.AZ: "Azerbaijan",
			territory.BA: "Bhosnia le Herzegovina",
			territory.BB: "Bhabhadosi",
			territory.BD: "Bhangiladeshi",
			territory.BE: "Bhelgium",
			territory.BF: "Bhukina Faso",
			territory.BG: "Bhulgariya",
			territory.BH: "Bhahareni",
			territory.BI: "Bhurundi",
			territory.BJ: "Bhenini",
			territory.BM: "Bhemuda",
			territory.BN: "Brunei",
			territory.BO: "Bholiviya",
			territory.BR: "Brazili",
			territory.BS: "Bhahamas",
			territory.BT: "Bhutani",
			territory.BW: "Botswana",
			territory.BY: "Bhelarusi",
			territory.BZ: "Bhelize",
			territory.CA: "Khanada",
			territory.CD: "Democratic Republic of the Congo",
			territory.CF: "Central African Republic",
			territory.CG: "Khongo",
			territory.CH: "Switzerland",
			territory.CI: "Ivory Coast",
			territory.CK: "Cook Islands",
			territory.CL: "Chile",
			territory.CM: "Khameruni",
			territory.CN: "China",
			territory.CO: "Kholombiya",
			territory.CR: "Khosta Rikha",
			territory.CU: "Cuba",
			territory.CV: "Cape Verde Islands",
			territory.CY: "Cyprus",
			territory.CZ: "Czech Republic",
			territory.DE: "Germany",
			territory.DJ: "Djibouti",
			territory.DK: "Denmakhi",
			territory.DM: "Dominikha",
			territory.DO: "Dominican Republic",
			territory.DZ: "Aljeriya",
			territory.EC: "Ecuador",
			territory.EE: "Estonia",
			territory.EG: "Egypt",
			territory.ER: "Eritrea",
			territory.ES: "Spain",
			territory.ET: "Ethiopia",
			territory.FI: "Finland",
			territory.FJ: "Fiji",
			territory.FK: "Falkland Islands",
			territory.FM: "Micronesia",
			territory.FR: "Furansi",
			territory.GA: "Gabhoni",
			territory.GB: "United Kingdom",
			territory.GD: "Grenada",
			territory.GE: "Georgia",
			territory.GF: "Gwiyana ye Furansi",
			territory.GH: "Ghana",
			territory.GI: "Gibraltar",
			territory.GL: "Greenland",
			territory.GM: "Gambiya",
			territory.GN: "Guinea",
			territory.GP: "Guadeloupe",
			territory.GQ: "Equatorial Guinea",
			territory.GR: "Greece",
			territory.GT: "Guatemala",
			territory.GU: "Guam",
			territory.GW: "Guinea-Bissau",
			territory.GY: "Guyana",
			territory.HN: "Honduras",
			territory.HR: "Croatia",
			territory.HT: "Hayiti",
			territory.HU: "Hungary",
			territory.ID: "Indonesiya",
			territory.IE: "Ireland",
			territory.IL: "Isuraeli",
			territory.IN: "Indiya",
			territory.IO: "British Indian Ocean Territory",
			territory.IQ: "Iraki",
			territory.IR: "Iran",
			territory.IS: "Iceland",
			territory.IT: "Itali",
			territory.JM: "Jamaica",
			territory.JO: "Jodani",
			territory.JP: "Japan",
			territory.KE: "Khenya",
			territory.KG: "Kyrgyzstan",
			territory.KH: "Cambodia",
			territory.KI: "Khiribati",
			territory.KM: "Khomoro",
			territory.KN: "Saint Kitts and Nevis",
			territory.KP: "North Korea",
			territory.KR: "South Korea",
			territory.KW: "Khuweiti",
			territory.KY: "Cayman Islands",
			territory.KZ: "Kazakhstan",
			territory.LA: "Laos",
			territory.LB: "Lebhanoni",
			territory.LC: "Saint Lucia",
			territory.LI: "Liechtenstein",
			territory.LK: "Sri Lanka",
			territory.LR: "Libheriya",
			territory.LS: "Lesotho",
			territory.LT: "Lithuania",
			territory.LU: "Luxembourg",
			territory.LV: "Latvia",
			territory.LY: "Libhiya",
			territory.MA: "Morokho",
			territory.MC: "Monakho",
			territory.MD: "Moldova",
			territory.MG: "Madagaska",
			territory.MH: "Marshall Islands",
			territory.ML: "Mali",
			territory.MM: "Myanmar",
			territory.MN: "Mongolia",
			territory.MP: "Northern Mariana Islands",
			territory.MQ: "Martinique",
			territory.MR: "Mauritania",
			territory.MS: "Montserrat",
			territory.MT: "Malta",
			territory.MU: "Mauritius",
			territory.MV: "Maldives",
			territory.MW: "Malawi",
			territory.MX: "Meksikho",
			territory.MY: "Malezhiya",
			territory.MZ: "Mozambique",
			territory.NA: "Namibhiya",
			territory.NC: "New Caledonia",
			territory.NE: "Niger",
			territory.NF: "Norfolk Island",
			territory.NG: "Nigeriya",
			territory.NI: "Nicaragua",
			territory.NL: "Netherlands",
			territory.NO: "Noweyi",
			territory.NP: "Nephali",
			territory.NR: "Nauru",
			territory.NU: "Niue",
			territory.NZ: "New Zealand",
			territory.OM: "Omani",
			territory.PA: "Panama",
			territory.PE: "Pheru",
			territory.PF: "Pholinesiya ye Fulansi",
			territory.PG: "Papua New Guinea",
			territory.PH: "Philippines",
			territory.PK: "Phakistani",
			territory.PL: "Pholandi",
			territory.PM: "Saint Pierre and Miquelon",
			territory.PN: "Pitcairn",
			territory.PR: "Puerto Rico",
			territory.PS: "Palestinian West Bank and Gaza",
			territory.PT: "Portugal",
			territory.PW: "Palau",
			territory.PY: "Paraguay",
			territory.QA: "Kathari",
			territory.RE: "Réunion",
			territory.RO: "Romania",
			territory.RU: "Rashiya",
			territory.RW: "Ruwanda",
			territory.SA: "Saudi Arabia",
			territory.SB: "Solomon Islands",
			territory.SC: "Seychelles",
			territory.SD: "Sudani",
			territory.SE: "Sweden",
			territory.SG: "Singapore",
			territory.SH: "Saint Helena",
			territory.SI: "Slovenia",
			territory.SK: "Slovakia",
			territory.SL: "Sierra Leone",
			territory.SM: "San Marino",
			territory.SN: "Senegali",
			territory.SO: "Somaliya",
			territory.SR: "Suriname",
			territory.ST: "São Tomé and Príncipe",
			territory.SV: "El Salvador",
			territory.SY: "Syria",
			territory.SZ: "Swaziland",
			territory.TC: "Turks and Caicos Islands",
			territory.TD: "Chadi",
			territory.TG: "Thogo",
			territory.TH: "Thayilandi",
			territory.TJ: "Tajikistan",
			territory.TK: "Thokelawu",
			territory.TL: "East Timor",
			territory.TM: "Turkmenistan",
			territory.TN: "Tunisiya",
			territory.TO: "Thonga",
			territory.TR: "Thekhi",
			territory.TT: "Trinidad le Tobago",
			territory.TV: "Thuvalu",
			territory.TW: "Thayiwani",
			territory.TZ: "Tanzaniya",
			territory.UA: "Yukreini",
			territory.UG: "Uganda",
			territory.US: "Amelika",
			territory.UY: "Yurugwai",
			territory.UZ: "Uzbekistan",
			territory.VA: "Vatican State",
			territory.VC: "Saint Vincent and the Grenadines",
			territory.VE: "Venezuela",
			territory.VG: "British Virgin Islands",
			territory.VI: "U.S. Virgin Islands",
			territory.VN: "Vietnam",
			territory.VU: "Vhanuatu",
			territory.WF: "Wallis and Futuna",
			territory.WS: "Samowa",
			territory.YE: "Yemeni",
			territory.YT: "Mayotte",
			territory.ZA: "Mzansi ye Afrika",
			territory.ZM: "Zambiya",
			territory.ZW: "Zimbabwe",
		},
	}
}
