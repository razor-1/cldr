// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_pcm_NG() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "pcm_NG",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "H:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'fọ' {0}", Long: "{1} 'fọ' {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jén", Feb: "Fẹ́b", Mar: "Mach", Apr: "Épr", May: "Mee", Jun: "Jun", Jul: "Jul", Aug: "Ọ́gọ", Sep: "Sẹp", Oct: "Ọkt", Nov: "Nọv", Dec: "Dis"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Jénúári", Feb: "Fẹ́búári", Mar: "Mach", Apr: "Éprel", May: "Mee", Jun: "Jun", Jul: "Julai", Aug: "Ọgọst", Sep: "Sẹptẹ́mba", Oct: "Ọktóba", Nov: "Nọvẹ́mba", Dec: "Disẹ́mba"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sọ́n", Mon: "Mọ́n", Tue: "Tiú", Wed: "Wẹ́n", Thu: "Tọ́z", Fri: "Fraí", Sat: "Sát"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Sọ́n", Mon: "Mọ́n", Tue: "Tiú", Wed: "Wẹ́n", Thu: "Tọ́z", Fri: "Fraí", Sat: "Sát"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sọ́ndè", Mon: "Mọ́ndè", Tue: "Tiúzdè", Wed: "Wẹ́nẹ́zdè", Thu: "Tọ́zdè", Fri: "Fraídè", Sat: "Sátọdè"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "Fọ mọ́nin", PM: "Fọ ívnin"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Yunaítẹ́d Áráb Ẹ́míréts Dírham", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afgán Afgáni", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Albéniá Lẹk", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armẹ́niá Dram", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Nẹ́dalánds Antílián Gílda", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angólá Kwánza", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Ajẹntína Pẹ́so", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Ọstréliá Dọ́la", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Arúba Flọ́rin", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Azẹrbaiján Mánat", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Bọ́sniá an Hẹzẹgovína Mak Wé Pẹ́sin Fít Chenj", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbédọs Dọ́la", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladẹ́sh Táka", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Bọlgériá Lẹv", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Baréin Dínar", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burúndí Frank", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bẹmiúda Dọ́la", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Brunẹí Dọ́la", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Bolíviá Boliviáno", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Brazíl Rẹal", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Bahámas Dọ́la", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Bután Ngúltrum", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Botswáná Púla", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Bẹlarús Rúbul", Symbol: "BYN"},
				currency.BZD: cldr.Currency{DisplayName: "Bẹliz Dọ́la", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kánádá Dọ́la", Symbol: "KA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kóngó Frank", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Swís Frank", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Chílí Pẹ́so", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Chaíná Yuan (ples-dẹm aúsaíd chaína)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Chaíná Yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolómbiá Pẹ́so", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Kósta Ríka Kólọn", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Kiúbá Pẹ́so Wé Pẹ́sin Fít Chenj", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kiúbá Pẹ́so", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Kép Vẹ́d Ẹskúdo", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Chẹ́k Kórúna", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Jibútí Frank", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Dẹ́nmák Króna", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dọmíníkan Pẹ́so", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Aljíria Dínar", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Íjípt Paund", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Ẹritrẹá Nákfa", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ẹtiópiá Berr", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Yúro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Fíjí Dọ́la", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Fọlkland Aílands Paund", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Brítísh Páund", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Jọ́jiá Lári", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Ganá Sídi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Jibrọ́lta Páund", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gámbiá Dalási", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Gíní Frank", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Guátẹmála Kwuẹ́tzal", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Giyána Dọ́la", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Họng Kọ́ng Dọ́la", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Họndúrán Lẹmpíra", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kroéshia Kúna", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haíti Gourd", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Họngériá Fọ́rint", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indoníshiá Rupia", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Ízrẹ́l Niú Shẹ́kẹl", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Índiá Rúpi", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Irák Dínar", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Irán Rial", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Aíslánd Króna", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaíka Dọla", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Jọ́dán Dínar", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Japán Yẹn", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Kẹ́nyá Shílin", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kẹjístan Som", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kambódiá Riẹl", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Kọ́mọ́ros Frank", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Nọ́t Koriá Wọn", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Saút Koriá Wọn", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuwét Dínar", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Kéman Aílands Dọla", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kazakstan Tẹ́nj", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laós Kip", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Lẹ́bánọ́n Paund", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lánká Rúpi", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Laibẹ́riá Dọ́la", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Líbia Dínar", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Morọko Dírham", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Mọldóva Lu", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Malagásí Ariári", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Masẹdónia Dínar", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Miánmá Kiat", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mọngóliá Túgrik", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Makáo Pátáka", Symbol: "MOP"},
				currency.MRU: cldr.Currency{DisplayName: "Mọriténiá Uguíya", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Mọríshọ́s Rúpi", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Mọ́ldívs Rúfíya", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Maláwi ́Kwácha", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Mẹ́ksíko Pẹ́so", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Maléshiá Ríngit", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Mozámbík Métíkal", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namíbiá Dọ́la", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naijíriá Naíra", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Nikarágwua Kordóba", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Nọ́wé Króna", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nẹ́pál Rúpi", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Niú Zílánd Dọ́las", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Omán Rial", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Pánáma Balbóa", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Pẹrúvián Sol", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Pápuá Niú Gíni Kína", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Fílípíns Píso", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakístán Rúpi", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Pólánd Zílọ́ti", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Páragwuá Guaráni", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Kata Ríal", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Roméniá Lu", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Sẹrbia Dínar", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rọ́shiá Rúbul", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ruwándá Frank", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saúdí Arébiá Riyal", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Sólómọ́n Aílands Dọ́la", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Sẹ́chẹ́ls Rúpi", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Sudan Paund", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Swídẹ́n Króna", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapọ́ Dọ́la", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Sent Hẹlẹ́ná Paund", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Siẹ́ra Líoniá Liọn", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Sọmáliá Shílin", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Súrínám Dọla", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Saút Sudán Paund", Symbol: "SSP"},
				currency.STN: cldr.Currency{DisplayName: "Sao Tómẹ & Prínsípẹ Dóbra", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Síriá Paund", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Swází Lilánjẹ́ni", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Taílánd Baht", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Tajíkstan Sómóni", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Tọkmẹ́nístán Mánat", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tuníshia Dínar", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tónga Pánga", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Tọ́kí Líra", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trínídad & Tobágo Dọ́la", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Niú Taiwán Dọ́la", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzániá Shílin", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Yukrẹín Rívnia", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Yugándá Shílin", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "US Dọ́la", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Yurugwaí Pẹ́so", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Uzbẹ́kistan Som", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Vẹnẹzuẹlá Bolívar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Viẹ́tnám Dọng", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuátú Vátu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samóa Tála", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Sẹ́ntrál Áfríká Frank", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Íst Karíbián Dọla", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Wẹ́st Afríká Sẹ́fa Frank", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Frẹ́nch Poliníshiá Frank", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Mọní Wé Pípol Nọ No", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Yẹ́mẹ́n Rial", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Saút Áfríká Rand", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Zámbiá Kwácha", Symbol: "ZMW"},
			},
		},
	}
}
