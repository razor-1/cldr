// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_la() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "la",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, 'die' d MMMM y G", Long: "'die' d MMMM y G", Medium: "'die' d MMM y G", Short: "d M y G"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'de' {0}", Long: "{1} 'de' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Ian", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "Mai", Jun: "Iun", Jul: "Iul", Aug: "Aug", Sep: "Sep", Oct: "Oct", Nov: "Nov", Dec: "Dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Ianuarius", Feb: "Februarius", Mar: "Martius", Apr: "Aprilis", May: "Maius", Jun: "Iunius", Jul: "Iulius", Aug: "Augustus", Sep: "September", Oct: "October", Nov: "November", Dec: "December"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Dom", Mon: "Lun", Tue: "Mar", Wed: "Mer", Thu: "Iov", Fri: "Ven", Sat: "Sab"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Dominica", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: "Sabbatum"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
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
				currency.CHF: cldr.Currency{DisplayName: "Francus Helveticus", Symbol: ""},
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
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Marca Finniae", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Libra sterlingorum", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GRD: cldr.Currency{DisplayName: "Drachma", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Dollarium Hongkongense", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.ITL: cldr.Currency{DisplayName: "Lira Italiana", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yen", Symbol: "JP¥"},
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
				currency.MXN: cldr.Currency{DisplayName: "Pensum Mexicanum", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Corona Norvegiae", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PEN: cldr.Currency{DisplayName: "Nuevo Sol", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "Rubelus Russicus", Symbol: ""},
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
				currency.USD: cldr.Currency{DisplayName: "Dollarium Civitatum Foederatarum", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Argentum", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Aurum", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "F\u202fCFA"},
				currency.XPD: cldr.Currency{DisplayName: "Palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platinum", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA: "Afarica",
			language.AB: "Abasca",
			language.AE: "Avestana",
			language.AF: "Africana",
			language.AM: "Amharica",
			language.AN: "Aragonensis",
			language.AR: "Arabica",
			language.AS: "Assamica",
			language.AY: "Aymara",
			language.AZ: "Atropatenica",
			language.BE: "Ruthenica Alba",
			language.BG: "Bulgarica",
			language.BH: "Bihari",
			language.BN: "Bengalica",
			language.BO: "Tibetana",
			language.BR: "Britonica",
			language.BS: "Bosnica",
			language.CA: "Catalana",
			language.CH: "Chamoruana",
			language.CO: "Corsa",
			language.CS: "Bohemica",
			language.CU: "Slavica Ecclesiastica",
			language.CV: "Chuvassica",
			language.CY: "Cambrica",
			language.DA: "Danica",
			language.DE: "Germanica",
			language.DV: "Dhivehi",
			language.DZ: "Dzongkha",
			language.EL: "Neograeca",
			language.EN: "Anglica",
			language.EO: "Esperantica",
			language.ES: "Hispanica",
			language.ET: "Estonica",
			language.EU: "Vasconica",
			language.FA: "Persica",
			language.FI: "Finnica",
			language.FO: "Faroensis",
			language.FR: "Francogallica",
			language.GA: "Hibernica",
			language.GD: "Scotica",
			language.GL: "Gallaica",
			language.GN: "Guaranica",
			language.GU: "Gujaratensis",
			language.GV: "Monensis",
			language.HI: "Hindica",
			language.HR: "Croatica",
			language.HT: "Creolus Haitianus",
			language.HU: "Hungarica",
			language.HY: "Armenica",
			language.IA: "Interlingua",
			language.IE: "Interlingue",
			language.IG: "Igbonica",
			language.IN: "Indonesia",
			language.IS: "Islandica",
			language.IT: "Italiana",
			language.IW: "Hebraica",
			language.JA: "Iaponica",
			language.JI: "Iudaeogermanica",
			language.JV: "Iavensis",
			language.KA: "Georgiana",
			language.KK: "Cazachica",
			language.KL: "Groenlandica",
			language.KM: "Chmerica",
			language.KN: "Cannadica",
			language.KO: "Coreana",
			language.KS: "Casmirica",
			language.KU: "Curdica",
			language.KW: "Cornubica",
			language.KY: "Chirgisica",
			language.LA: "Latina",
			language.LB: "Luxemburgica",
			language.LI: "Limburgica",
			language.LT: "Lithuanica",
			language.LV: "Lettonica",
			language.MG: "Malagasiana",
			language.MI: "Maoriana",
			language.MK: "Macedonica",
			language.ML: "Malabarica",
			language.MN: "Mongolica",
			language.MR: "Marathica",
			language.MS: "Malayana",
			language.MT: "Melitensis",
			language.MY: "Birmanica",
			language.NE: "Nepalensis",
			language.NL: "Batava",
			language.NO: "Norvegica",
			language.OC: "Occitana",
			language.OJ: "Ojibwayensis",
			language.OR: "Orissensis",
			language.OS: "Ossetica",
			language.PA: "Panjabica",
			language.PI: "Palica",
			language.PL: "Polonica",
			language.PS: "Afganica",
			language.PT: "Lusitana",
			language.QU: "Quechuae",
			language.RM: "Rhaetica",
			language.RO: "Dacoromanica",
			language.RU: "Russica",
			language.SA: "Sanscrita",
			language.SC: "Sarda",
			language.SD: "Sindhuica",
			language.SE: "Samica septentrionalis",
			language.SI: "Singhalensis",
			language.SK: "Slovaca",
			language.SL: "Slovena",
			language.SM: "Samoana",
			language.SO: "Somalica",
			language.SQ: "Albanica",
			language.SR: "Serbica",
			language.SV: "Suecica",
			language.SW: "Suahili",
			language.TA: "Tamulica",
			language.TE: "Telingana",
			language.TG: "Tadzikica",
			language.TH: "Thai",
			language.TI: "Tigrinya",
			language.TK: "Turcomannica",
			language.TL: "Tagalog",
			language.TO: "Tongana",
			language.TR: "Turcica",
			language.TT: "Tatarica",
			language.UK: "Ucrainica",
			language.UR: "Urdu",
			language.UZ: "Uzbecica",
			language.VI: "Vietnamica",
			language.ZH: "Sinica",
			language.ZU: "Zuluana",
		},
		Territories: cldr.Territories{
			territory.V_001: "Mundus",
			territory.V_002: "Africa",
			territory.V_003: "America Septentrionalis",
			territory.V_005: "America Australis",
			territory.V_009: "Oceania",
			territory.V_011: "Africa Occidentalis",
			territory.V_013: "America Centralis",
			territory.V_014: "Africa Orientalis",
			territory.V_015: "Africa Septentrionalis",
			territory.V_018: "Africa Australis (regio)",
			territory.V_019: "America",
			territory.V_029: "Caribaeum",
			territory.V_030: "Asia Orientalis",
			territory.V_035: "Asia Meridiorientalis",
			territory.V_039: "Europa Centralis",
			territory.V_053: "Australasia",
			territory.V_054: "Melanesia",
			territory.V_057: "Micronesia (regio)",
			territory.V_061: "Polynesia",
			territory.V_142: "Asia",
			territory.V_143: "Media Asia",
			territory.V_150: "Europa",
			territory.V_151: "Europa Orientalis",
			territory.V_154: "Europa Septentrionalis",
			territory.V_155: "Europa Occidentalis",
			territory.AD:    "Andorra",
			territory.AE:    "Phylarchiarum Arabicarum Confoederatio",
			territory.AF:    "Afgania",
			territory.AG:    "Antiqua et Barbuda",
			territory.AI:    "Anguilla",
			territory.AL:    "Albania",
			territory.AM:    "Armenia",
			territory.AO:    "Angolia",
			territory.AR:    "Argentina",
			territory.AS:    "Samoa Americana",
			territory.AT:    "Austria",
			territory.AU:    "Australia",
			territory.AW:    "Aruba",
			territory.AX:    "Alandia",
			territory.AZ:    "Atropatene",
			territory.BA:    "Bosnia et Herzegovina",
			territory.BB:    "Barbata",
			territory.BD:    "Bangladesha",
			territory.BE:    "Belgica",
			territory.BF:    "Burkina Faso",
			territory.BG:    "Bulgaria",
			territory.BH:    "Baharina",
			territory.BI:    "Burundia",
			territory.BJ:    "Beninum",
			territory.BL:    "Insula Sancti Bartholomaei",
			territory.BM:    "Bermuda",
			territory.BN:    "Bruneium",
			territory.BO:    "Bolivia",
			territory.BR:    "Brasilia",
			territory.BS:    "Insulae Bahamenses",
			territory.BT:    "Butania",
			territory.BU:    "Birmania",
			territory.BV:    "Insula Bouvet",
			territory.BW:    "Botswana",
			territory.BY:    "Ruthenia Alba",
			territory.BZ:    "Beliza",
			territory.CA:    "Canada",
			territory.CC:    "Insulae Cocos seu Keeling",
			territory.CD:    "Res publica Democratica Congensis",
			territory.CF:    "Res publica Africae Mediae",
			territory.CG:    "Res publica Congoliae",
			territory.CH:    "Helvetia",
			territory.CI:    "Litus Eburneum",
			territory.CK:    "Insulae Cook",
			territory.CL:    "Chilia",
			territory.CM:    "Cameronia",
			territory.CN:    "Res publica popularis Sinarum",
			territory.CO:    "Columbia",
			territory.CR:    "Costarica",
			territory.CU:    "Cuba",
			territory.CV:    "Res publica Capitis Viridis",
			territory.CW:    "Insula Curacensis",
			territory.CX:    "Insula Christi Natalis",
			territory.CY:    "Cyprus",
			territory.CZ:    "Cechia",
			territory.DD:    "Res publica Democratica Germanica",
			territory.DE:    "Germania",
			territory.DJ:    "Gibutum",
			territory.DK:    "Dania",
			territory.DM:    "Dominica",
			territory.DO:    "Res publica Dominicana",
			territory.DZ:    "Algerium",
			territory.EC:    "Aequatoria",
			territory.EE:    "Estonia",
			territory.EG:    "Aegyptus",
			territory.EH:    "Sahara Occidentalis",
			territory.ER:    "Erythraea",
			territory.ES:    "Hispania",
			territory.ET:    "Aethiopia",
			territory.FI:    "Finnia",
			territory.FJ:    "Viti",
			territory.FK:    "Insulae Malvinae",
			territory.FM:    "Micronesia",
			territory.FO:    "Faeroae insulae",
			territory.FR:    "Francia",
			territory.GA:    "Gabonia",
			territory.GB:    "Britanniarum Regnum",
			territory.GD:    "Granata",
			territory.GE:    "Georgia",
			territory.GF:    "Guiana Francica",
			territory.GG:    "Lisia",
			territory.GH:    "Gana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Groenlandia",
			territory.GM:    "Gambia",
			territory.GN:    "Guinea",
			territory.GP:    "Guadalupa",
			territory.GQ:    "Guinea Aequatorensis",
			territory.GR:    "Graecia",
			territory.GT:    "Guatimalia",
			territory.GU:    "Guama",
			territory.GW:    "Guinea Bissaviensis",
			territory.GY:    "Guiana",
			territory.HK:    "Hongcongum",
			territory.HN:    "Honduria",
			territory.HR:    "Croatia",
			territory.HT:    "Haitia",
			territory.HU:    "Hungaria",
			territory.ID:    "Indonesia",
			territory.IE:    "Hibernia",
			territory.IL:    "Israël",
			territory.IM:    "Monapia",
			territory.IN:    "India",
			territory.IQ:    "Iracum",
			territory.IR:    "Irania",
			territory.IS:    "Islandia",
			territory.IT:    "Italia",
			territory.JE:    "Caesarea Insula",
			territory.JM:    "Iamaica",
			territory.JO:    "Iordania",
			territory.JP:    "Iaponia",
			territory.KE:    "Kenia",
			territory.KG:    "Chirgisia",
			territory.KH:    "Cambosia",
			territory.KI:    "Kiribati",
			territory.KM:    "Insulae Comorianae",
			territory.KN:    "Sanctus Christophorus et Nevis",
			territory.KP:    "Res publica Popularis Democratica Coreana",
			territory.KR:    "Res publica Coreana",
			territory.KW:    "Cuvaitum",
			territory.KY:    "Insulae Caimanenses",
			territory.KZ:    "Kazachstania",
			territory.LA:    "Laotia",
			territory.LB:    "Libanus",
			territory.LC:    "Sancta Lucia",
			territory.LI:    "Lichtenstenum",
			territory.LK:    "Taprobane",
			territory.LR:    "Liberia",
			territory.LS:    "Lesothum",
			territory.LT:    "Lituania",
			territory.LU:    "Luxemburgum",
			territory.LV:    "Lettonia",
			territory.LY:    "Libya",
			territory.MA:    "Marocum",
			territory.MC:    "Monoecus",
			territory.MD:    "Res publica Moldavica",
			territory.ME:    "Mons Niger",
			territory.MG:    "Madagascaria",
			territory.MH:    "Insulae Marsalienses",
			territory.MK:    "Res publica Macedonica",
			territory.ML:    "Malium",
			territory.MN:    "Mongolia",
			territory.MO:    "Macaum",
			territory.MP:    "Insulae Marianae Septentrionales",
			territory.MQ:    "Martinica",
			territory.MR:    "Mauritania",
			territory.MS:    "Montserrat",
			territory.MT:    "Melita",
			territory.MU:    "Mauritia",
			territory.MV:    "Insulae Maldivae",
			territory.MW:    "Malavium",
			territory.MX:    "Mexicum",
			territory.MY:    "Malaesia",
			territory.MZ:    "Mozambicum",
			territory.NA:    "Namibia",
			territory.NC:    "Nova Caledonia",
			territory.NE:    "Res publica Nigritana",
			territory.NF:    "Insula Norfolcia",
			territory.NG:    "Nigeria",
			territory.NI:    "Nicaragua",
			territory.NL:    "Nederlandia",
			territory.NO:    "Norvegia",
			territory.NP:    "Nepalia",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Nova Zelandia",
			territory.OM:    "Omania",
			territory.PA:    "Panama",
			territory.PE:    "Peruvia",
			territory.PF:    "Polynesia Francica",
			territory.PG:    "Papua Nova Guinea",
			territory.PH:    "Philippinae",
			territory.PK:    "Pakistania",
			territory.PL:    "Polonia",
			territory.PM:    "Insulae Sancti Petri et Miquelonensis",
			territory.PN:    "Insulae Pitcairn",
			territory.PR:    "Portus Dives",
			territory.PS:    "Territoria Palaestinensia",
			territory.PT:    "Portugallia",
			territory.PW:    "Belavia",
			territory.PY:    "Paraquaria",
			territory.QA:    "Quataria",
			territory.RE:    "Reunio",
			territory.RO:    "Romania",
			territory.RS:    "Serbia",
			territory.RU:    "Russia",
			territory.RW:    "Ruanda",
			territory.SA:    "Arabia Saudiana",
			territory.SB:    "Insulae Salomonis",
			territory.SC:    "Insulae Seisellenses",
			territory.SD:    "Sudania",
			territory.SE:    "Suecia",
			territory.SG:    "Singapura",
			territory.SH:    "Sancta Helena, Ascensio et Tristan da Cunha",
			territory.SI:    "Slovenia",
			territory.SK:    "Slovacia",
			territory.SL:    "Mons Leoninus",
			territory.SM:    "Res publica Sancti Marini",
			territory.SN:    "Senegalia",
			territory.SO:    "Somalia",
			territory.SR:    "Surinamia",
			territory.SS:    "Sudania Australis",
			territory.ST:    "Insulae Sancti Thomae et Principis",
			territory.SV:    "Salvatoria",
			territory.SY:    "Syria",
			territory.SZ:    "Swazia",
			territory.TC:    "Insulae Turcenses et Caicenses",
			territory.TD:    "Tzadia",
			territory.TG:    "Togum",
			territory.TH:    "Thailandia",
			territory.TJ:    "Tadzikistania",
			territory.TK:    "Tokelau",
			territory.TL:    "Timoria Orientalis",
			territory.TM:    "Turcomannia",
			territory.TN:    "Tunesia",
			territory.TO:    "Tonga",
			territory.TR:    "Turcia",
			territory.TT:    "Trinitas et Tabacum",
			territory.TV:    "Tuvalu",
			territory.TW:    "Res publica Sinarum",
			territory.TZ:    "Tanzania",
			territory.UA:    "Ucraina",
			territory.UG:    "Uganda",
			territory.US:    "Civitates Foederatae Americae",
			territory.UY:    "Uraquaria",
			territory.UZ:    "Uzbecia",
			territory.VA:    "Civitas Vaticana",
			territory.VC:    "Sanctus Vincentius et Granatinae",
			territory.VE:    "Venetiola",
			territory.VG:    "Virginis Insulae Britannicae",
			territory.VI:    "Virginis Insulae Americanae",
			territory.VN:    "Vietnamia",
			territory.VU:    "Vanuatu",
			territory.WF:    "Vallis et Futuna",
			territory.WS:    "Samoa",
			territory.XK:    "Kosovia",
			territory.YE:    "Iemenia",
			territory.YT:    "Maiotta",
			territory.YU:    "Iugoslavia",
			territory.ZA:    "Africa Australis",
			territory.ZM:    "Zambia",
			territory.ZW:    "Zimbabua",
		},
	}
}
