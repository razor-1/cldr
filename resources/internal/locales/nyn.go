// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_nyn() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "nyn",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "KBZ", Feb: "KBR", Mar: "KST", Apr: "KKN", May: "KTN", Jun: "KMK", Jul: "KMS", Aug: "KMN", Sep: "KMW", Oct: "KKM", Nov: "KNK", Dec: "KNB"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Okwokubanza", Feb: "Okwakabiri", Mar: "Okwakashatu", Apr: "Okwakana", May: "Okwakataana", Jun: "Okwamukaaga", Jul: "Okwamushanju", Aug: "Okwamunaana", Sep: "Okwamwenda", Oct: "Okwaikumi", Nov: "Okwaikumi na kumwe", Dec: "Okwaikumi na ibiri"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "SAN", Mon: "ORK", Tue: "OKB", Wed: "OKS", Thu: "OKN", Fri: "OKT", Sat: "OMK"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "K", Tue: "R", Wed: "S", Thu: "N", Fri: "T", Sat: "M"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sande", Mon: "Orwokubanza", Tue: "Orwakabiri", Wed: "Orwakashatu", Thu: "Orwakana", Fri: "Orwakataano", Sat: "Orwamukaaga"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham za Buharabu", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza ya Angora", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Doora ya Austureeriya", Symbol: "A$"},
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
				currency.BWP: cldr.Currency{DisplayName: "Pura ya Botswana", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Doora ya Kanada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faranga ya Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faranga ya Swisi", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuan Renminbi ya China", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudo ya Kepuvede", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faranga ya Gyibuti", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinari ya Arigyeriya", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Paundi ya Misiri", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa ya Eritireya", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr ya Ethiopiya", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Yuro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Paundi ya Bungyereza", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Cedi ya Ghana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi ya Gambia", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Faranga ya Guinea", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupiya ya India", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yeni ya Japaani", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shiringi ya Kenya", Symbol: ""},
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
				currency.LRD: cldr.Currency{DisplayName: "Doora ya Liberia", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti ya Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinari ya Libya", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirram ya Moroko", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariari ya Maragariita", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ougwiya ya Mouriteeniya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ougwiya ya Mouriteeniya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupiiha ya Mauritiasi", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwaca ya Malawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikari ya Mozambikwi", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Doora ya Namibiya", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Naira ya Naigyeriya", Symbol: "₦"},
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
				currency.SAR: cldr.Currency{DisplayName: "Riya ya Saudi", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupiiha ya Sherisheri", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dinari ya Sudani", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Paundi ya Sudan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Paundi ya Senti Herena", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Leone", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leone (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Eshiringi ya Somalia", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Purinsipo (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra ya Sao Tome na Purinsipo", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinari ya Tunisia", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Eshiringi ya Tanzania", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Eshiringi ya Uganda", Symbol: "USh"},
				currency.USD: cldr.Currency{DisplayName: "Doora ya America", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Faranga ya CFA BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Faranga ya CFA BCEAO", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Randi ya Sausi Afirika", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha ya Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha ya Zambia", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Doora ya Zimbabwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Orukani",
			language.AM:  "Orumariki",
			language.AR:  "Oruharabu",
			language.BE:  "Oruberarusi",
			language.BG:  "Oruburugariya",
			language.BN:  "Orubengari",
			language.CS:  "Oruceeki",
			language.DE:  "Orugirimaani",
			language.EL:  "Oruguriiki",
			language.EN:  "Orungyereza",
			language.ES:  "Orusupaani",
			language.FA:  "Orupaasiya",
			language.FR:  "Orufaransa",
			language.HA:  "Oruhausa",
			language.HI:  "Oruhindi",
			language.HU:  "Oruhangare",
			language.ID:  "Oruindonezia",
			language.IG:  "Oruibo",
			language.IT:  "Oruyitare",
			language.JA:  "Orujapaani",
			language.JV:  "Orujava",
			language.KM:  "Orukambodiya",
			language.KO:  "Orukoreya",
			language.MS:  "Orumalesiya",
			language.MY:  "Oruburuma",
			language.NE:  "Orunepali",
			language.NL:  "Orudaaki",
			language.NYN: "Runyankore",
			language.PA:  "Orupungyabi",
			language.PL:  "Orupoori",
			language.PT:  "Orupocugo",
			language.RO:  "Oruromania",
			language.RU:  "Orurrasha",
			language.RW:  "Orunyarwanda",
			language.SO:  "Orusomaari",
			language.SV:  "Oruswidi",
			language.TA:  "Orutamiri",
			language.TH:  "Orutailandi",
			language.TR:  "Orukuruki",
			language.UK:  "Orukuraini",
			language.UR:  "Oru-Urudu",
			language.VI:  "Oruviyetinaamu",
			language.YO:  "Oruyoruba",
			language.ZH:  "Oruchaina",
			language.ZU:  "Oruzuru",
		},
		Territories: cldr.Territories{
			territory.AD: "Andora",
			territory.AE: "Amahanga ga Buharabu ageeteereine",
			territory.AF: "Afuganistani",
			territory.AG: "Angiguwa na Babuda",
			territory.AI: "Angwira",
			territory.AL: "Arubania",
			territory.AM: "Arimeniya",
			territory.AO: "Angora",
			territory.AR: "Arigentina",
			territory.AS: "Samowa ya Ameerika",
			territory.AT: "Osituria",
			territory.AU: "Ositureeriya",
			territory.AW: "Aruba",
			territory.AZ: "Azabagyani",
			territory.BA: "Boziniya na Hezegovina",
			territory.BB: "Babadosi",
			territory.BD: "Bangaradeshi",
			territory.BE: "Bubirigi",
			territory.BF: "Bokina Faso",
			territory.BG: "Burugariya",
			territory.BH: "Bahareni",
			territory.BI: "Burundi",
			territory.BJ: "Benini",
			territory.BM: "Berimuda",
			territory.BN: "Burunei",
			territory.BO: "Boriiviya",
			territory.BR: "Buraziiri",
			territory.BS: "Bahama",
			territory.BT: "Butani",
			territory.BW: "Botswana",
			territory.BY: "Bararusi",
			territory.BZ: "Berize",
			territory.CA: "Kanada",
			territory.CD: "Demokoratika Ripaaburika ya Kongo",
			territory.CF: "Eihanga rya Rwagati ya Afirika",
			territory.CG: "Kongo",
			territory.CH: "Swisi",
			territory.CI: "Aivore Kositi",
			territory.CK: "Ebizinga bya Kuuku",
			territory.CL: "Chile",
			territory.CM: "Kameruuni",
			territory.CN: "China",
			territory.CO: "Korombiya",
			territory.CR: "Kositarika",
			territory.CU: "Cuba",
			territory.CV: "Ebizinga bya Kepuvade",
			territory.CY: "Saipurasi",
			territory.CZ: "Ripaaburika ya Zeeki",
			territory.DE: "Bugirimaani",
			territory.DJ: "Gyibuti",
			territory.DK: "Deenimaaka",
			territory.DM: "Dominika",
			territory.DO: "Ripaaburika ya Dominica",
			territory.DZ: "Arigyeriya",
			territory.EC: "Ikweda",
			territory.EE: "Esitoniya",
			territory.EG: "Misiri",
			territory.ER: "Eriteriya",
			territory.ES: "Sipeyini",
			territory.ET: "Ethiyopiya",
			territory.FI: "Bufini",
			territory.FJ: "Figyi",
			territory.FK: "Ebizinga bya Faakilanda",
			territory.FM: "Mikironesiya",
			territory.FR: "Bufaransa",
			territory.GA: "Gabooni",
			territory.GB: "Bungyereza",
			territory.GD: "Gurenada",
			territory.GE: "Gyogiya",
			territory.GF: "Guyana ya Bufaransa",
			territory.GH: "Gana",
			territory.GI: "Giburaata",
			territory.GL: "Guriinirandi",
			territory.GM: "Gambiya",
			territory.GN: "Gine",
			territory.GP: "Gwaderupe",
			territory.GQ: "Guni",
			territory.GR: "Guriisi",
			territory.GT: "Gwatemara",
			territory.GU: "Gwamu",
			territory.GW: "Ginebisau",
			territory.GY: "Guyana",
			territory.HN: "Hondurasi",
			territory.HR: "Korasiya",
			territory.HT: "Haiti",
			territory.HU: "Hangare",
			territory.ID: "Indoneeziya",
			territory.IE: "Irerandi",
			territory.IL: "Isirairi",
			territory.IN: "Indiya",
			territory.IO: "Ebizinga bya Indian ebya Bungyereza",
			territory.IQ: "Iraaka",
			territory.IR: "Iraani",
			territory.IS: "Aisilandi",
			territory.IT: "Itare",
			territory.JM: "Gyamaika",
			territory.JO: "Yorudaani",
			territory.JP: "Gyapaani",
			territory.KE: "Kenya",
			territory.KG: "Kirigizistani",
			territory.KH: "Kambodiya",
			territory.KI: "Kiribati",
			territory.KM: "Koromo",
			territory.KN: "Senti Kittis na Nevisi",
			territory.KP: "Koreya Amatemba",
			territory.KR: "Koreya Amashuuma",
			territory.KW: "Kuweiti",
			territory.KY: "Ebizinga bya Kayimani",
			territory.KZ: "Kazakisitani",
			territory.LA: "Layosi",
			territory.LB: "Lebanoni",
			territory.LC: "Senti Rusiya",
			territory.LI: "Lishenteni",
			territory.LK: "Siriranka",
			territory.LR: "Liberiya",
			territory.LS: "Lesotho",
			territory.LT: "Lithuania",
			territory.LU: "Lakizembaaga",
			territory.LV: "Latviya",
			territory.LY: "Libya",
			territory.MA: "Morocco",
			territory.MC: "Monaco",
			territory.MD: "Moridova",
			territory.MG: "Madagasika",
			territory.MH: "Ebizinga bya Marshaa",
			territory.ML: "Mari",
			territory.MM: "Myanamar",
			territory.MN: "Mongoria",
			territory.MP: "Ebizinga by’amatemba ga Mariana",
			territory.MQ: "Martinique",
			territory.MR: "Mauriteeniya",
			territory.MS: "Montserrati",
			territory.MT: "Marita",
			territory.MU: "Maurishiasi",
			territory.MV: "Maridives",
			territory.MW: "Marawi",
			territory.MX: "Mexico",
			territory.MY: "marayizia",
			territory.MZ: "Mozambique",
			territory.NA: "Namibiya",
			territory.NC: "Niukaredonia",
			territory.NE: "Naigya",
			territory.NF: "Ekizinga Norifoko",
			territory.NG: "Naigyeriya",
			territory.NI: "Nikaragwa",
			territory.NL: "Hoorandi",
			territory.NO: "Noorwe",
			territory.NP: "Nepo",
			territory.NR: "Nauru",
			territory.NU: "Niue",
			territory.NZ: "Niuzirandi",
			territory.OM: "Omaani",
			territory.PA: "Panama",
			territory.PE: "Peru",
			territory.PF: "Polinesia ya Bufaransa",
			territory.PG: "Papua",
			territory.PH: "Firipino",
			territory.PK: "Pakisitaani",
			territory.PL: "Poorandi",
			territory.PM: "Senti Piyerre na Mikweron",
			territory.PN: "Pitkaini",
			territory.PR: "Pwetoriko",
			territory.PS: "Parestiina na Gaza",
			territory.PT: "Pocugo",
			territory.PW: "Palaawu",
			territory.PY: "Paragwai",
			territory.QA: "Kata",
			territory.RE: "Riyuniyoni",
			territory.RO: "Romaniya",
			territory.RU: "Rrasha",
			territory.RW: "Rwanda",
			territory.SA: "Saudi Areebiya",
			territory.SB: "Ebizinga bya Surimaani",
			territory.SC: "Shesheresi",
			territory.SD: "Sudani",
			territory.SE: "Swideni",
			territory.SG: "Singapo",
			territory.SH: "Senti Herena",
			territory.SI: "Sirovaaniya",
			territory.SK: "Sirovaakiya",
			territory.SL: "Sirra Riyooni",
			territory.SM: "Samarino",
			territory.SN: "Senego",
			territory.SO: "Somaariya",
			territory.SR: "Surinaamu",
			territory.ST: "Sawo Tome na Purinsipo",
			territory.SV: "Eri Salivado",
			territory.SY: "Siriya",
			territory.SZ: "Swazirandi",
			territory.TC: "Ebizinga bya Buturuki na Kaiko",
			territory.TD: "Chadi",
			territory.TG: "Togo",
			territory.TH: "Tairandi",
			territory.TJ: "Tajikisitani",
			territory.TK: "Tokerawu",
			territory.TL: "Burugweizooba bwa Timori",
			territory.TM: "Turukimenisitani",
			territory.TN: "Tunizia",
			territory.TO: "Tonga",
			territory.TR: "Buturuki /Take",
			territory.TT: "Turinidad na Tobago",
			territory.TV: "Tuvaru",
			territory.TW: "Tayiwaani",
			territory.TZ: "Tanzania",
			territory.UA: "Ukureini",
			territory.UG: "Uganda",
			territory.US: "Amerika",
			territory.UY: "Urugwai",
			territory.UZ: "Uzibekisitani",
			territory.VA: "Vatikani",
			territory.VC: "Senti Vinsent na Gurenadini",
			territory.VE: "Venezuwera",
			territory.VG: "Ebizinga bya Virigini ebya Bungyereza",
			territory.VI: "Ebizinga bya Virigini ebya Amerika",
			territory.VN: "Viyetinaamu",
			territory.VU: "Vanuatu",
			territory.WF: "Warris na Futuna",
			territory.WS: "Samowa",
			territory.YE: "Yemeni",
			territory.YT: "Mayote",
			territory.ZA: "Sausi Afirika",
			territory.ZM: "Zambia",
			territory.ZW: "Zimbabwe",
		},
	}
}
