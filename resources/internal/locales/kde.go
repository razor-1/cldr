// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_kde() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "kde",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mac", Apr: "Apr", May: "Mei", Jun: "Jun", Jul: "Jul", Aug: "Ago", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Des"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Mwedi Ntandi", Feb: "Mwedi wa Pili", Mar: "Mwedi wa Tatu", Apr: "Mwedi wa Nchechi", May: "Mwedi wa Nnyano", Jun: "Mwedi wa Nnyano na Umo", Jul: "Mwedi wa Nnyano na Mivili", Aug: "Mwedi wa Nnyano na Mitatu", Sep: "Mwedi wa Nnyano na Nchechi", Oct: "Mwedi wa Nnyano na Nnyano", Nov: "Mwedi wa Nnyano na Nnyano na U", Dec: "Mwedi wa Nnyano na Nnyano na M"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ll2", Mon: "Ll3", Tue: "Ll4", Wed: "Ll5", Thu: "Ll6", Fri: "Ll7", Sat: "Ll1"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "2", Mon: "3", Tue: "4", Wed: "5", Thu: "6", Fri: "7", Sat: "1"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Liduva lyapili", Mon: "Liduva lyatatu", Tue: "Liduva lyanchechi", Wed: "Liduva lyannyano", Thu: "Liduva lyannyano na linji", Fri: "Liduva lyannyano na mavili", Sat: "Liduva litandi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "Muhi", PM: "Chilo"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham ya Falme za Chiarabu", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza ya Angola", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dola ya Australia", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinari ya Bahareni", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Faranga ya Burundi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pula ya Botswana", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dola ya Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faranga ya Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faranga ya Uswisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Renminbi ya China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudo ya Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faranga ya Jibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinari ya Aljeria", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Pauni ya Misri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa ya Eritrea", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Bir ya Uhabeshi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Pauni ya Uingereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi ya Ghana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ya Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Faranga ya Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia ya India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Sarafu ya Chijapani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilingi ya Kenya", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faranga ya Komoro", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dola ya Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ya Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinari ya Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirham ya Moroko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Faranga ya Bukini", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugwiya ya Moritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugwiya ya Moritania", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupia ya Morisi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha ya Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikali ya Msumbiji", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dola ya Namibia", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira ya Nijeria", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Faranga ya Rwanda", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal ya Saudia", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia ya Shelisheli", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dinari ya Sudani", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Pauni ya Sudani", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Pauni ya Santahelena", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Leoni", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leoni (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Shilingi ya Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Principe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinari ya Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilingi ya Tanzania", Symbol: "TSh"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Shilingi ya Uganda", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dola ya Marekani", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Faranga CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Faranga CFA BCEAO", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Randi ya Afrika Kusini", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha ya Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha ya Zambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dola ya Zimbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Chakan",
			language.AM:  "Chamhali",
			language.AR:  "Chalabu",
			language.BE:  "Chibelalusi",
			language.BG:  "Chibulgalia",
			language.BN:  "Chibangla",
			language.CS:  "Chichechi",
			language.DE:  "Chidyelumani",
			language.EL:  "Chigilichi",
			language.EN:  "Chiingeleza",
			language.ES:  "Chihispania",
			language.FA:  "Chiajemi",
			language.FR:  "Chifalansa",
			language.HA:  "Chihausa",
			language.HI:  "Chihindi",
			language.HU:  "Chihungali",
			language.ID:  "Chiiongonesia",
			language.IG:  "Chiigbo",
			language.IT:  "Chiitaliano",
			language.JA:  "Chidyapani",
			language.JV:  "Chidyava",
			language.KDE: "Chimakonde",
			language.KM:  "Chikambodia",
			language.KO:  "Chikolea",
			language.MS:  "Chimalesia",
			language.MY:  "Chibulma",
			language.NE:  "Chinepali",
			language.NL:  "Chiholanzi",
			language.PA:  "Chipunjabi",
			language.PL:  "Chipolandi",
			language.PT:  "Chileno",
			language.RO:  "Chilomania",
			language.RU:  "Chilusi",
			language.RW:  "Chinyalwanda",
			language.SO:  "Chisomali",
			language.SV:  "Chiswidi",
			language.TA:  "Chitamil",
			language.TH:  "Chitailandi",
			language.TR:  "Chituluchi",
			language.UK:  "Chiuklania",
			language.UR:  "Chiuldu",
			language.VI:  "Chivietinamu",
			language.YO:  "Chiyoluba",
			language.ZH:  "Chichina",
			language.ZU:  "Chizulu",
		},
		Territories: cldr.Territories{
			territory.AD: "Andola",
			territory.AE: "Dimiliki dya Vakulungwa va Chalabu",
			territory.AF: "Afuganistani",
			territory.AG: "Antigua na Balbuda",
			territory.AI: "Angwila",
			territory.AL: "Albania",
			territory.AM: "Almenia",
			territory.AO: "Angola",
			territory.AR: "Adyentina",
			territory.AS: "Samoa ya Malekani",
			territory.AT: "Austlia",
			territory.AU: "Austlalia",
			territory.AW: "Aluba",
			territory.AZ: "Azabadyani",
			territory.BA: "Bosnia na Hezegovina",
			territory.BB: "Babadosi",
			territory.BD: "Bangladeshi",
			territory.BE: "Ubelgidi",
			territory.BF: "Buchinafaso",
			territory.BG: "Bulgalia",
			territory.BH: "Bahaleni",
			territory.BI: "Bulundi",
			territory.BJ: "Benini",
			territory.BM: "Belmuda",
			territory.BN: "Blunei",
			territory.BO: "Bolivia",
			territory.BR: "Blazili",
			territory.BS: "Bahama",
			territory.BT: "Butani",
			territory.BW: "Botswana",
			territory.BY: "Belalusi",
			territory.BZ: "Belize",
			territory.CA: "Kanada",
			territory.CD: "Jamuhuli ya Chidemoklasia ya kuKongo",
			territory.CF: "Jamuhuli ya Afilika ya Paching’ati",
			territory.CG: "Kongo",
			territory.CH: "Uswisi",
			territory.CI: "Kodivaa",
			territory.CK: "Chisiwa cha Cook",
			territory.CL: "Chile",
			territory.CM: "Kameluni",
			territory.CN: "China",
			territory.CO: "Kolombia",
			territory.CR: "Kostalika",
			territory.CU: "Kuba",
			territory.CV: "Kepuvede",
			territory.CY: "Kuplosi",
			territory.CZ: "Jamuhuli ya Chechi",
			territory.DE: "Udyerumani",
			territory.DJ: "Dyibuti",
			territory.DK: "Denmaki",
			territory.DM: "Dominika",
			territory.DO: "Jamuhuli ya Dominika",
			territory.DZ: "Aljelia",
			territory.EC: "Ekwado",
			territory.EE: "Estonia",
			territory.EG: "Misli",
			territory.ER: "Elitilea",
			territory.ES: "Hispania",
			territory.ET: "Uhabeshi",
			territory.FI: "Ufini",
			territory.FJ: "Fiji",
			territory.FK: "Chisiwa cha Falkland",
			territory.FM: "Mikilonesia",
			territory.FR: "Ufalansa",
			territory.GA: "Gaboni",
			territory.GB: "Nngalesa",
			territory.GD: "Glenada",
			territory.GE: "Dyodya",
			territory.GF: "Gwiyana ya Ufalansa",
			territory.GH: "Ghana",
			territory.GI: "Diblalta",
			territory.GL: "Glinlandi",
			territory.GM: "Gambia",
			territory.GN: "Gine",
			territory.GP: "Gwadelupe",
			territory.GQ: "Ginekweta",
			territory.GR: "Ugilichi",
			territory.GT: "Gwatemala",
			territory.GU: "Gwam",
			territory.GW: "Ginebisau",
			territory.GY: "Guyana",
			territory.HN: "Hondulasi",
			territory.HR: "Kolasia",
			territory.HT: "Haiti",
			territory.HU: "Hungalia",
			territory.ID: "Indonesia",
			territory.IE: "Ayalandi",
			territory.IL: "Islaeli",
			territory.IN: "India",
			territory.IO: "Lieneo lyaki Nngalesa Nbahali ya Hindi",
			territory.IQ: "Ilaki",
			territory.IR: "Uadyemi",
			territory.IS: "Aislandi",
			territory.IT: "Italia",
			territory.JM: "Dyamaika",
			territory.JO: "Yordani",
			territory.JP: "Dyapani",
			territory.KE: "Kenya",
			territory.KG: "Kiligizistani",
			territory.KH: "Kambodia",
			territory.KI: "Kilibati",
			territory.KM: "Komolo",
			territory.KN: "Santakitzi na Nevis",
			territory.KP: "Kolea Kasikazini",
			territory.KR: "Kolea Kusini",
			territory.KW: "Kuwaiti",
			territory.KY: "Chisiwa cha Kemen",
			territory.KZ: "Kazachistani",
			territory.LA: "Laosi",
			territory.LB: "Lebanoni",
			territory.LC: "Santalusia",
			territory.LI: "Lishenteni",
			territory.LK: "Sililanka",
			territory.LR: "Libelia",
			territory.LS: "Lesoto",
			territory.LT: "Litwania",
			territory.LU: "Lasembagi",
			territory.LV: "Lativia",
			territory.LY: "Libya",
			territory.MA: "Moloko",
			territory.MC: "Monako",
			territory.MD: "Moldova",
			territory.MG: "Bukini",
			territory.MH: "Chisiwa cha Malushal",
			territory.ML: "Mali",
			territory.MM: "Myama",
			territory.MN: "Mongolia",
			territory.MP: "Chisiwa cha Marian cha Kasikazini",
			territory.MQ: "Malitiniki",
			territory.MR: "Molitania",
			territory.MS: "Monselati",
			territory.MT: "Malta",
			territory.MU: "Molisi",
			territory.MV: "Modivu",
			territory.MW: "Malawi",
			territory.MX: "Meksiko",
			territory.MY: "Malesia",
			territory.MZ: "Msumbiji",
			territory.NA: "Namibia",
			territory.NC: "Nyukaledonia",
			territory.NE: "Nidyeli",
			territory.NF: "Chisiwa cha Nolufok",
			territory.NG: "Nidyelia",
			territory.NI: "Nikalagwa",
			territory.NL: "Uholanzi",
			territory.NO: "Norwe",
			territory.NP: "Nepali",
			territory.NR: "Naulu",
			territory.NU: "Niue",
			territory.NZ: "Nyuzilandi",
			territory.OM: "Omani",
			territory.PA: "Panama",
			territory.PE: "Pelu",
			territory.PF: "Polinesia ya Ufalansa",
			territory.PG: "Papua",
			territory.PH: "Filipino",
			territory.PK: "Pakistani",
			territory.PL: "Polandi",
			territory.PM: "Santapieli na Mikeloni",
			territory.PN: "Pitikeluni",
			territory.PR: "Pwetoliko",
			territory.PS: "Nchingu wa Magalibi wa Mpanda wa kuGaza wa kuPales",
			territory.PT: "Uleno",
			territory.PW: "Palau",
			territory.PY: "Palagwai",
			territory.QA: "Katali",
			territory.RE: "Liyunioni",
			territory.RO: "Lomania",
			territory.RU: "Ulusi",
			territory.RW: "Lwanda",
			territory.SA: "Saudia",
			territory.SB: "Chisiwa cha Solomon",
			territory.SC: "Shelisheli",
			territory.SD: "Sudani",
			territory.SE: "Uswidi",
			territory.SG: "Singapoo",
			territory.SH: "Santahelena",
			territory.SI: "Slovenia",
			territory.SK: "Slovakia",
			territory.SL: "Siela Leoni",
			territory.SM: "Samalino",
			territory.SN: "Senegali",
			territory.SO: "Somalia",
			territory.SR: "Sulinamu",
			territory.ST: "Saotome na Prinsipe",
			territory.SV: "Elsavado",
			territory.SY: "Silia",
			territory.SZ: "Uswazi",
			territory.TC: "Chisiwa cha Tuluchi na Kaiko",
			territory.TD: "Chadi",
			territory.TG: "Togo",
			territory.TH: "Tailandi",
			territory.TJ: "Tadikistani",
			territory.TK: "Tokelau",
			territory.TL: "Timoli ya Mashaliki",
			territory.TM: "Tuluchimenistani",
			territory.TN: "Tunisia",
			territory.TO: "Tonga",
			territory.TR: "Utuluchi",
			territory.TT: "Tilinidad na Tobago",
			territory.TV: "Tuvalu",
			territory.TW: "Taiwani",
			territory.TZ: "Tanzania",
			territory.UA: "Uklaini",
			territory.UG: "Uganda",
			territory.US: "Malekani",
			territory.UY: "Ulugwai",
			territory.UZ: "Uzibechistani",
			territory.VA: "Vatikani",
			territory.VC: "Santavisenti na Glenadini",
			territory.VE: "Venezuela",
			territory.VG: "Chisiwa Chivihi cha Wingalesa",
			territory.VI: "Chisiwa Chivihi cha Malekani",
			territory.VN: "Vietinamu",
			territory.VU: "Vanuatu",
			territory.WF: "Walis na Futuna",
			territory.WS: "Samoa",
			territory.YE: "Yemeni",
			territory.YT: "Maole",
			territory.ZA: "Afilika Kusini",
			territory.ZM: "Zambia",
			territory.ZW: "Zimbabwe",
		},
	}
}
