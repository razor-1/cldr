// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_naq() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "naq",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss\u202fa zzzz", Long: "h:mm:ss\u202fa z", Medium: "h:mm:ss\u202fa", Short: "h:mm\u202fa"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "May", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Oct", Nov: "Nov", Dec: "Dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ǃKhanni", Feb: "ǃKhanǀgôab", Mar: "ǀKhuuǁkhâb", Apr: "ǃHôaǂkhaib", May: "ǃKhaitsâb", Jun: "Gamaǀaeb", Jul: "ǂKhoesaob", Aug: "Aoǁkhuumûǁkhâb", Sep: "Taraǀkhuumûǁkhâb", Oct: "ǂNûǁnâiseb", Nov: "ǀHooǂgaeb", Dec: "Hôasoreǁkhâb"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Son", Mon: "Ma", Tue: "De", Wed: "Wu", Thu: "Do", Fri: "Fr", Sat: "Sat"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "E", Wed: "W", Thu: "D", Fri: "F", Sat: "A"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sontaxtsees", Mon: "Mantaxtsees", Tue: "Denstaxtsees", Wed: "Wunstaxtsees", Thu: "Dondertaxtsees", Fri: "Fraitaxtsees", Sat: "Satertaxtsees"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ǁgoagas", PM: "ǃuias"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "United Arab Emirates Dirham", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "Angolan Kwanzab", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Australian Dollari", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Bahrain Dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Burundi Franc", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Botswanan Pulab", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Canadian Dollari", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Congolese Franc", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Swiss Franci", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Chinese Yuan Renminbi", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Caboverdiano", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Djibouti Franc", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Algerian Dinar", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Egytian Ponds", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Eritreian Nakfa", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ethiopian Birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Eurob", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "British Ponds", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Ghana Cedi", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Gambia Dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Guinea Franc", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indian Rupee", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Japanese Yenni", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenyan Shilling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Comorian Franc", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Liberian Dollar", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lesotho Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Libyan Dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Moroccan Dirham", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Malagasy Franc", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Mauritania Ouguiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Mauritania Ouguiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mauritius Rupeeb", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Malawian Kwachab", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Mozambique Metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Namibia Dollari", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nigerian Naira", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Rwanda Franci", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi Riyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Seychelles Rupee", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Sudanese Dinar", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Sudanese Ponds", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "St Helena Ponds", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Leone", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leone (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Somali Shillings", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Sao Tome and Principe Dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Sao Tome and Principe Dobra", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Tunisian Dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzanian Shillings", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Ugandan Shillings", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "US Dollari", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "CFA Franc BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "CFA Franc BCEAO", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "South African Randi", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambian Kwachab (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Zambian Kwachab", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwe Dollari", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Akangowab",
			language.AM:  "Amharicgowab",
			language.AR:  "Arabiǁî gowab",
			language.BE:  "Belarusanǁî gowab",
			language.BG:  "Bulgariaǁî gowab",
			language.BN:  "Bengaliǁî gowab",
			language.CS:  "Czechǁî gowab",
			language.DE:  "Duits",
			language.EL:  "Xriks",
			language.EN:  "Engels",
			language.ES:  "Spaans",
			language.FA:  "Persiaǁî gowab",
			language.FR:  "Frans",
			language.HA:  "Hausagowab",
			language.HI:  "Hindigowab",
			language.HU:  "Hungariaǁî gowab",
			language.ID:  "Indonesiaǁî gowab",
			language.IG:  "Igbogowab",
			language.IT:  "Italians",
			language.JA:  "Japanees",
			language.JV:  "Javanese",
			language.KM:  "Khmerǁî gowab, Central",
			language.KO:  "Koreaǁî gowab",
			language.MS:  "Malayǁî gowab",
			language.MY:  "Burmesǁî gowab",
			language.NAQ: "Khoekhoegowab",
			language.NE:  "Nepalǁî gowab",
			language.NL:  "Hollands",
			language.PA:  "Punjabigowab",
			language.PL:  "Poleǁî gowab",
			language.PT:  "Portugees",
			language.RO:  "Romaniaǁî gowab",
			language.RU:  "Russiaǁî gowab",
			language.RW:  "Rwandaǁî gowab",
			language.SO:  "Somaliǁî gowab",
			language.SV:  "Swedeǁî gowab",
			language.TA:  "Tamilǁî gowab",
			language.TH:  "Thaiǁî gowab",
			language.TR:  "Turkeǁî gowab",
			language.UK:  "Ukrainiaǁî gowab",
			language.UR:  "Urduǁî gowab",
			language.VI:  "Vietnamǁî gowab",
			language.YO:  "Yorubab",
			language.ZH:  "Chineesǁî gowab, Mandarinni",
			language.ZU:  "Zulub",
		},
		Territories: cldr.Territories{
			territory.AD: "Andorrab",
			territory.AE: "United Arab Emirates",
			territory.AF: "Afghanistanni",
			territory.AG: "Antiguab tsî Barbudab",
			territory.AI: "Anguillab",
			territory.AL: "Albaniab",
			territory.AM: "Armeniab",
			territory.AO: "Angolab",
			territory.AR: "Argentinab",
			territory.AS: "Americab Samoab",
			territory.AT: "Austriab",
			territory.AU: "Australieb",
			territory.AW: "Arubab",
			territory.AZ: "Azerbaijanni",
			territory.BA: "Bosniab tsî Herzegovinab",
			territory.BB: "Barbados",
			territory.BD: "Banglades",
			territory.BE: "Belgiummi",
			territory.BF: "Burkina Fasob",
			territory.BG: "Bulgariab",
			territory.BH: "Bahrain",
			territory.BI: "Burundib",
			territory.BJ: "Benins",
			territory.BM: "Bermudas",
			territory.BN: "Brunei",
			territory.BO: "Boliviab",
			territory.BR: "Braziliab",
			territory.BS: "Bahamas",
			territory.BT: "Bhutans",
			territory.BW: "Botswanab",
			territory.BY: "Belarus",
			territory.BZ: "Belize",
			territory.CA: "Kanadab",
			territory.CD: "Democratic Republic of the Congo",
			territory.CF: "Central African Republiki",
			territory.CG: "Congob",
			territory.CH: "Switzerlandi",
			territory.CI: "Ivoorkusi",
			territory.CK: "Cook Islands",
			territory.CL: "Chilib",
			territory.CM: "Cameroonni",
			territory.CN: "Chinab",
			territory.CO: "Colombiab",
			territory.CR: "Costa Rica",
			territory.CU: "Cubab",
			territory.CV: "Cape Verde Islands",
			territory.CY: "Cyprus",
			territory.CZ: "Czech Republiki",
			territory.DE: "Duitslandi",
			territory.DJ: "Djibouti",
			territory.DK: "Denmarki",
			territory.DM: "Dominicab",
			territory.DO: "Dominican Republic",
			territory.DZ: "Algeriab",
			territory.EC: "Ecuadori",
			territory.EE: "Estoniab",
			territory.EG: "Egipteb",
			territory.ER: "Eritreab",
			territory.ES: "Spanieb",
			territory.ET: "Ethiopiab",
			territory.FI: "Finlandi",
			territory.FJ: "Fijib",
			territory.FK: "Falkland Islands",
			territory.FM: "Micronesia",
			territory.FR: "Frankreiki",
			territory.GA: "Gaboni",
			territory.GB: "United Kingdom",
			territory.GD: "Grenada",
			territory.GE: "Georgiab",
			territory.GF: "French Guiana",
			territory.GH: "Ghanab",
			territory.GI: "Gibraltar",
			territory.GL: "Greenland",
			territory.GM: "Gambiab",
			territory.GN: "Guineab",
			territory.GP: "Guadeloupe",
			territory.GQ: "Equatorial Guineab",
			territory.GR: "Xrikelandi",
			territory.GT: "Guatemala",
			territory.GU: "Guam",
			territory.GW: "Guinea-Bissau",
			territory.GY: "Guyana",
			territory.HN: "Honduras",
			territory.HR: "Croatiab",
			territory.HT: "Haiti",
			territory.HU: "Hongareieb",
			territory.ID: "Indonesiab",
			territory.IE: "Irlandi",
			territory.IL: "Israeli",
			territory.IN: "Indiab",
			territory.IO: "British Indian Ocean Territory",
			territory.IQ: "Iraqi",
			territory.IR: "Iranni",
			territory.IS: "Iceland",
			territory.IT: "Italiab",
			territory.JM: "Jamaicab",
			territory.JO: "Jordanni",
			territory.JP: "Japanni",
			territory.KE: "Kenyab",
			territory.KG: "Kyrgyzstanni",
			territory.KH: "Cambodiab",
			territory.KI: "Kiribati",
			territory.KM: "Comoros",
			territory.KN: "Saint Kitts and Nevis",
			territory.KP: "Koreab, Noord",
			territory.KR: "Koreab, Suid",
			territory.KW: "Kuwaiti",
			territory.KY: "Cayman Islands",
			territory.KZ: "Kazakhstanni",
			territory.LA: "Laos",
			territory.LB: "Lebanonni",
			territory.LC: "Saint Lucia",
			territory.LI: "Liechtensteinni",
			territory.LK: "Sri Lankab",
			territory.LR: "Liberiab",
			territory.LS: "Lesothob",
			territory.LT: "Lithuaniab",
			territory.LU: "Luxembourgi",
			territory.LV: "Latvia",
			territory.LY: "Libyab",
			territory.MA: "Morocco",
			territory.MC: "Monaco",
			territory.MD: "Moldova",
			territory.MG: "Madagascari",
			territory.MH: "Marshall Islands",
			territory.ML: "Malib",
			territory.MM: "Myanmar",
			territory.MN: "Mongolia",
			territory.MP: "Northern Mariana Islands",
			territory.MQ: "Martinique",
			territory.MR: "Mauritania",
			territory.MS: "Montserrat",
			territory.MT: "Malta",
			territory.MU: "Mauritius",
			territory.MV: "Maldives",
			territory.MW: "Malawib",
			territory.MX: "Mexicob",
			territory.MY: "Malaysiab",
			territory.MZ: "Mozambiki",
			territory.NA: "Namibiab",
			territory.NC: "New Caledonia",
			territory.NE: "Nigeri",
			territory.NF: "Norfolk Island",
			territory.NG: "Nigerieb",
			territory.NI: "Nicaraguab",
			territory.NL: "Netherlands",
			territory.NO: "Noorweeb",
			territory.NP: "Nepali",
			territory.NR: "Nauru",
			territory.NU: "Niue",
			territory.NZ: "New Zealandi",
			territory.OM: "Oman",
			territory.PA: "Panama",
			territory.PE: "Perub",
			territory.PF: "French Polynesia",
			territory.PG: "Papua New Guineab",
			territory.PH: "Philippinni",
			territory.PK: "Pakistanni",
			territory.PL: "Polandi",
			territory.PM: "Saint Pierre and Miquelon",
			territory.PN: "Pitcairn",
			territory.PR: "Puerto Rico",
			territory.PS: "Palestinian West Bank and Gaza",
			territory.PT: "Portugali",
			territory.PW: "Palau",
			territory.PY: "Paraguaib",
			territory.QA: "Qatar",
			territory.RE: "Réunion",
			territory.RO: "Romania",
			territory.RU: "Rasiab",
			territory.RW: "Rwandab",
			territory.SA: "Saudi Arabiab",
			territory.SB: "Solomon Islands",
			territory.SC: "Seychelles",
			territory.SD: "Sudanni",
			territory.SE: "Swedeb",
			territory.SG: "Singapore",
			territory.SH: "Saint Helena",
			territory.SI: "Slovenia",
			territory.SK: "Slovakia",
			territory.SL: "Sierra Leone",
			territory.SM: "San Marino",
			territory.SN: "Senegali",
			territory.SO: "Somaliab",
			territory.SR: "Suriname",
			territory.ST: "São Tomé and Príncipe",
			territory.SV: "El Salvadori",
			territory.SY: "Syriab",
			territory.SZ: "Swazilandi",
			territory.TC: "Turks and Caicos Islands",
			territory.TD: "Chadi",
			territory.TG: "Togob",
			territory.TH: "Thailandi",
			territory.TJ: "Tajikistan",
			territory.TK: "Tokelau",
			territory.TL: "East Timor",
			territory.TM: "Turkmenistan",
			territory.TN: "Tunisiab",
			territory.TO: "Tonga",
			territory.TR: "Turkeieb",
			territory.TT: "Trinidad and Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwan",
			territory.TZ: "Tanzaniab",
			territory.UA: "Ukraine",
			territory.UG: "Ugandab",
			territory.US: "Amerikab",
			territory.UY: "Uruguaib",
			territory.UZ: "Uzbekistan",
			territory.VA: "Vatican State",
			territory.VC: "Saint Vincent and the Grenadines",
			territory.VE: "Venezuelab",
			territory.VG: "British Virgin Islands",
			territory.VI: "U.S. Virgin Islands",
			territory.VN: "Vietnammi",
			territory.VU: "Vanuatu",
			territory.WF: "Wallis and Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yemen",
			territory.YT: "Mayotte",
			territory.ZA: "Suid Afrikab",
			territory.ZM: "Zambiab",
			territory.ZW: "Zimbabweb",
		},
	}
}
