// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_hy_AM() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "hy_AM",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y թ. MMMM d, EEEE", Long: "dd MMMM, y թ.", Medium: "dd MMM, y թ.", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "հնվ", Feb: "փտվ", Mar: "մրտ", Apr: "ապր", May: "մյս", Jun: "հնս", Jul: "հլս", Aug: "օգս", Sep: "սեպ", Oct: "հոկ", Nov: "նոյ", Dec: "դեկ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Հ", Feb: "Փ", Mar: "Մ", Apr: "Ա", May: "Մ", Jun: "Հ", Jul: "Հ", Aug: "Օ", Sep: "Ս", Oct: "Հ", Nov: "Ն", Dec: "Դ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "հունվար", Feb: "փետրվար", Mar: "մարտ", Apr: "ապրիլ", May: "մայիս", Jun: "հունիս", Jul: "հուլիս", Aug: "օգոստոս", Sep: "սեպտեմբեր", Oct: "հոկտեմբեր", Nov: "նոյեմբեր", Dec: "դեկտեմբեր"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "կիր", Mon: "երկ", Tue: "երք", Wed: "չրք", Thu: "հնգ", Fri: "ուր", Sat: "շբթ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Կ", Mon: "Ե", Tue: "Ե", Wed: "Չ", Thu: "Հ", Fri: "Ո", Sat: "Շ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "կր", Mon: "եկ", Tue: "եք", Wed: "չք", Thu: "հգ", Fri: "ու", Sat: "շբ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "կիրակի", Mon: "երկուշաբթի", Tue: "երեքշաբթի", Wed: "չորեքշաբթի", Thu: "հինգշաբթի", Fri: "ուրբաթ", Sat: "շաբաթ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Արաբական Միացյալ Էմիրությունների դիրհամ", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "աֆղանական աֆղանի", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "ալբանական լեկ", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "հայկական դրամ", Symbol: "֏"},
				currency.ANG: cldr.Currency{DisplayName: "նիդեռլանդական անտիլյան գուլդեն", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "անգոլական կվանզա", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "արգենտինական պեսո", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ավստրալիական դոլար", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "արուբական ֆլորին", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "ադրբեջանական մանաթ", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Բոսնիա և Հերցեգովինայի փոխարկվող մարկ", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "բարբադոսյան դոլար", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Բանգլադեշի տակա", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "բուլղարական լև", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Բահրեյնի դինար", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "բուրունդիական ֆրանկ", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "բերմուդյան դոլար", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Բրունեյի դոլար", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "բոլիվիական բոլիվիանո", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "բրազիլական ռեալ", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "բահամյան դոլար", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "բութանական նգուլտրում", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "բոթսվանական պուլա", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "բելառուսական ռուբլի", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Բելառուսական ռուբլի (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Բելիզի դոլար", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "կանադական դոլար", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Կոնգոյի ֆրանկ", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "շվեյցարական ֆրանկ", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "չիլիական պեսո", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "չինական օֆշորային յուան", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "չինական յուան", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "կոլումբիական պեսո", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Կոստա Ռիկայի կոլոն", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "կուբայական փոխարկվող պեսո", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "կուբայական պեսո", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Կաբո Վերդեի էսկուդո", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "չեխական կրոն", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Ջիբութիի ֆրանկ", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "դանիական կրոն", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "դոմինիկյան պեսո", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "ալժիրական դինար", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "եգիպտական ֆունտ", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "էրիթրեական նակվա", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "եթովպիական բիր", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "եվրո", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "ֆիջիական դոլար", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Ֆոլքլենդյան կղզիների ֆունտ", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "բրիտանական ֆունտ ստերլինգ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "վրացական լարի", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "գայանական սեդի", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Ջիբրալթարի ֆունտ", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "գամբիական դալասի", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "գվինեական ֆրանկ", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "գվատեմալական կետսալ", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "գայանական դոլար", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Հոնկոնգի դոլար", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "հոնդուրասական լեմպիրա", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "խորվաթական կունա", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "հայիթյան գուրդ", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "հունգարական ֆորինտ", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "ինդոնեզիական ռուփի", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Իսրայելի նոր շեկել", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "հնդկական ռուփի", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "իրաքյան դինար", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "իրանական ռիալ", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "իսլանդական կրոն", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Ճամայկայի դոլար", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "հորդանանյան դինար", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "ճապոնական իեն", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "քենիական շիլինգ", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "ղրղզական սոմ", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "կամբոջական ռիել", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "կոմորյան ֆրանկ", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "հյուսիսկորեական վոն", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "հարավկորեական վոն", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Քուվեյթի դինար", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Կայմանյան կղզիների դոլար", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "ղազախական տենգե", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "լաոսական կիպ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "լիբանանյան ֆունտ", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Շրի Լանկայի ռուփի", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "լիբերիական դոլար", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "Լիտվական լիտ", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "Լատվիական լատ", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "լիբիական դինար", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Մարոկկոյի դիրհամ", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "մոլդովական լեյ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Մադագասկարի արիարի", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "մակեդոնական դենար", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Մյանմայի կյատ", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "մոնղոլական տուգրիկ", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Մակաոյի պատակա", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "մավրիտանական ուգիյա (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "մավրիտանական ուգիյա", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "մավրիկյան ռուփի", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "մալդիվյան ռուֆիյա", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "մալավիական կվաչա", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "մեքսիկական պեսո", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "մալայզիական ռինգիտ", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "մոզամբիկյան մետիկալ", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "նամիբիական դոլար", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "նիգերիական նայրա", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "նիկարագուական կորդոբա", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "նորվեգական կրոն", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Նեպալի ռուփի", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "նորզելանդական դոլար", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Օմանի ռիալ", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "պանամական բալբոա", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Պերուի սոլ", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Պապուա Նոր Գվինեայի կինա", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ֆիլիպինյան պեսո", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "պակիստանյան ռուփի", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "լեհական զլոտի", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "պարագվայական գուարանի", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Կատարի ռիալ", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "ռումինական լեյ", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "սերբական դինար", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "ռուսական ռուբլի", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ռուանդական ֆրանկ", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Սաուդյան Արաբիայի ռիալ", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Սողոմոնյան կղզիների դոլար", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "սեյշելյան ռուփի", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "սուդանական ֆունտ", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "շվեդական կրոն", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Սինգապուրի դոլար", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Սուրբ Հեղինեի ֆունտ", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Սիեռա Լեոնեի լեոնե", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "սոմալիական շիլինգ", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "սուրինամական դոլար", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "հարավսուդանական ֆունտ", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Սան Տոմե և Փրինսիպիի դոբրա (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Սան Տոմե և Փրինսիպիի դոբրա", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "սիրիական ֆունտ", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "սվազիլենդական լիլանգենի", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "թայլանդական բատ", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "տաջիկական սոմոնի", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "թուրքմենական մանաթ", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "թունիսյան դինար", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Տոնգայի պաանգա", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "թուրքական լիրա", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Տրինիդադ և Տոբագոյի դոլար", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "թայվանական նոր դոլար", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "տանզանիական շիլինգ", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ուկրաինական գրիվնա", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "ուգանդական շիլինգ", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "ԱՄՆ դոլար", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "ուրուգվայական պեսո", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "ուզբեկական սոմ", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "վենեսուելական բոլիվար (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "վենեսուելական բոլիվար", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "վիետնամական դոնգ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Վանուատուի վատու", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "սամոական տալա", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Կենտրոնական Աֆրիկայի ԿՖԱ ֆրանկ", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "արևելակարիբյան դոլար", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Արևմտյան Աֆրիկայի ԿՖԱ ֆրանկ", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "ԿՊՖ ֆրանկ", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "անհայտ արժույթ", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "եմենական ռիալ", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "հարավաֆրիկյան ռանդ", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Զամբիական կվաչա (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "զամբիական կվաչա", Symbol: "ZMW"},
			},
		},
	}
}