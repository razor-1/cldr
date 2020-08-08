// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_ug_CN() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y d-MMMM، EEEE", Long: "d-MMMM، y", Medium: "d-MMM، y", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}، {0}", Short: "{1}، {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "يانۋار", Feb: "فېۋرال", Mar: "مارت", Apr: "ئاپرېل", May: "ماي", Jun: "ئىيۇن", Jul: "ئىيۇل", Aug: "ئاۋغۇست", Sep: "سېنتەبىر", Oct: "ئۆكتەبىر", Nov: "نويابىر", Dec: "دېكابىر"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "يانۋار", Feb: "فېۋرال", Mar: "مارت", Apr: "ئاپرېل", May: "ماي", Jun: "ئىيۇن", Jul: "ئىيۇل", Aug: "ئاۋغۇست", Sep: "سېنتەبىر", Oct: "ئۆكتەبىر", Nov: "نويابىر", Dec: "دېكابىر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "يە", Mon: "دۈ", Tue: "سە", Wed: "چا", Thu: "پە", Fri: "جۈ", Sat: "شە"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ي", Mon: "د", Tue: "س", Wed: "چ", Thu: "پ", Fri: "ج", Sat: "ش"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ي", Mon: "د", Tue: "س", Wed: "چ", Thu: "پ", Fri: "ج", Sat: "ش"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "يەكشەنبە", Mon: "دۈشەنبە", Tue: "سەيشەنبە", Wed: "چارشەنبە", Thu: "پەيشەنبە", Fri: "جۈمە", Sat: "شەنبە"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "چ.ب", PM: "چ.ك"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "چ.ب", PM: "چ.ك"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "چ.ب", PM: "چ.ك"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ug]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "ئاندورران پېسېتاسى", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "ئەرەب بىرلەشمە خەلىپىلىكى دەرھەمى", Symbol: ""},
				currency.AFA: cldr.Currency{DisplayName: "ئافغان ئافغانى (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "ئافغان ئافغانى", Symbol: ""},
				currency.ALK: cldr.Currency{DisplayName: "ئالبانىيە لېكى (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "ئالبانىيە لېكى", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "ئەرمېنىيە دىرامى", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "گوللاندىيەگە قاراشلىق ئانتىللېن گۇلدېنى", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "ئانگولا كۇۋانزاسى", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "ئانگولا كۇۋانزاسى (1977–1991)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "ئانگولا يېڭى كۇۋانزاسى (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "ئانگولا قايتا تەڭشەلگەن كۇۋانزاسى (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "ئارگېنتىنا ئاۋسترالى", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "ئارگېنتىنا پېسو لېيى (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "ئارگېنتىنا پېسوسى (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "ئارگېنتىنا پېسوسى (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "ئارگېنتىنا پېسوسى", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "ئاۋسترىيە شىللىڭى", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "ئاۋسترالىيە دوللىرى", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "ئارۇبان فىلورۇنى", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "ئەزەربەيجان ماناتى (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "ئەزەربەيجان ماناتى", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "بوسنىيە-خېرتسېگوۋىنا دىنارى (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "بوسنىيە-خېرتسېگوۋىنا ئالماشتۇرۇشچان ماركى", Symbol: "KM"},
				currency.BAN: cldr.Currency{DisplayName: "بوسنىيە-خېرتسېگوۋىنا يېڭى دىنارى (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "باربادوس دوللىرى", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "باڭلادىش تاكاسى", Symbol: "৳"},
				currency.BEC: cldr.Currency{DisplayName: "بېلگىيە فرانكى (ئالماشتۇرۇشچان)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "بېلگىيە فرانكى", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "بېلگىيە فرانكى (پۇل–مۇئامىلە)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "بۇلغارىيە قاتتىق لېۋاسى", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "بۇلغارىيە ئىجتىمائىي لېۋاسى", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "بۇلغارىيە لېۋاسى", Symbol: ""},
				currency.BGO: cldr.Currency{DisplayName: "بۇلغارىيە لېۋاسى (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "بەھرەين دىنارى", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "بۇرۇندى فرانكى", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "بېرمۇدا دوللىرى", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "بىرۇنېي دوللىرى", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "بولىۋىيە بولىۋىيانوسى", Symbol: "Bs"},
				currency.BOL: cldr.Currency{DisplayName: "بولىۋىيە بولىۋىيانوسى (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "بولىۋىيە پىسوسى", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "بولىۋىيە مۇدولى", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "بىرازىلىيە يېڭى كرۇزېروسى (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "بىرازىلىيە كرۇزادوسى (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "بىرازىلىيە يېڭى كرۇزېروسى (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "بىرازىلىيە رىيالى", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "بىرازىلىيە يېڭى كرۇزادوسى (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "بىرازىلىيە كرۇزېروسى (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "بىرازىلىيە كرۇزېروسى (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "باھاما دوللىرى", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "بۇتان نگۇلترۇمى", Symbol: ""},
				currency.BUK: cldr.Currency{DisplayName: "بىرما كىياتى", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "بوتسۋانا پۇلاسى", Symbol: "P"},
				currency.BYB: cldr.Currency{DisplayName: "بېلارۇسىيە يېڭى رۇبلىسى (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "بېلارۇسىيە رۇبلىسى", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "بېلارۇسىيە رۇبلىسى (۲۰۰۰–۲۰۱۶)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "بېلىز دوللىرى", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "كانادا دوللىرى", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "كونگو فرانكى", Symbol: ""},
				currency.CHE: cldr.Currency{DisplayName: "WIR ياۋرو", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "شىۋېتسىيە فرانكى", Symbol: ""},
				currency.CHW: cldr.Currency{DisplayName: "WIR فرانكى", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "چىلى ئېسكۇدوسى", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "چىلى ھېسابات بىرلىكى (UF)", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "چىلى پېسوسى", Symbol: "$"},
				currency.CNX: cldr.Currency{DisplayName: "جۇڭگو خەلق بانكىسى دوللىرى", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "جۇڭگو يۈەنى", Symbol: "￥"},
				currency.COP: cldr.Currency{DisplayName: "كولومبىيە پېسوسى", Symbol: "$"},
				currency.COU: cldr.Currency{DisplayName: "كولومبىيە ھەقىقىي قىممەت بىرلىكى", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "كوستارىكا كولونى", Symbol: "₡"},
				currency.CSD: cldr.Currency{DisplayName: "سېربىيە دىنارى (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "چېخسىلوۋاكىيە قاتتىق كورۇناسى", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "كۇبا ئالماشتۇرۇشچان پېسوسى", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "كۇبا پېسوسى", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "يېشىل تۇمشۇق ئېسكۇدوسى", Symbol: ""},
				currency.CYP: cldr.Currency{DisplayName: "سىپرۇس فوند ستېرلىڭى", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "چېخ جۇمھۇرىيىتى كورۇناسى", Symbol: "Kč"},
				currency.DDM: cldr.Currency{DisplayName: "شەرقىي گېرمانىيە ماركى", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "گېرمانىيە ماركى", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "جىبۇتى فرانكى", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "دانىيە كرونى", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "دومىنىكا پېسوسى", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "ئالجىرىيە دىنارى", Symbol: ""},
				currency.ECS: cldr.Currency{DisplayName: "ئېكۋادور سۇكرېسى", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "ئېكۋادور تۇراقلىق قىممەت بىرلىكى", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "ئېستونىيە كرۇنى", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "مىسىر فوند سىتېرلىڭى", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "ئېرىترېيە ناكفاسى", Symbol: ""},
				currency.ESA: cldr.Currency{DisplayName: "ئىسپانىيە پېسېتاسى (A ھېسابات)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "ئىسپانىيە پېسېتاسى (ئالماشتۇرۇش ھېساباتى)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "ئىسپانىيە پېسېتاسى", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ئېفىيوپىيە بىررى", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "ياۋرو", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "فىنلاندىيە مارككاسى", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "فىجى دوللىرى", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "فالكلاند ئاراللىرى فوند سىتېرلىڭى", Symbol: "£"},
				currency.FRF: cldr.Currency{DisplayName: "فىرانسىيە فرانكى", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "ئەنگلىيە فوند سىتېرلىڭى", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "گىرۇزىيە كۇپون لارىتى", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "گىرۇزىيە لارىسى", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "گانا سېدىسى (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "گانا سېدىسى", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "جەبىلتارىق فوند سىتېرلىڭى", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "گامبىيە دالاسى", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "گىۋىنېيە فرانكى", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "گىۋىنېيە سىلىسى", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "ئېكۋاتور گىۋىنېيە ئېكۋېلېسى", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "گىرېتسىيە دراخماسى", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "گىۋاتېمالا كۇۋېتزالى", Symbol: "Q"},
				currency.GWE: cldr.Currency{DisplayName: "پورتۇگالىيە گىۋىنېيە ئېسكۇدوسى", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "گىۋىنېيە-بىسسائۇ پېسوسى", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "گىۋىئانا دوللىرى", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "شياڭگاڭ دوللىرى", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "ھوندۇراس لېمپىراسى", Symbol: "L"},
				currency.HRD: cldr.Currency{DisplayName: "كىرودىيە دىنارى", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "كىرودىيە كۇناسى", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "ھايتى گۇردېسى", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "ۋېنگىرىيە فورېنتى", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "ھىندونېزىيە رۇپىيەسى", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "ئىرېلاندىيە فوندستېرلىڭى", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "ئىسرائىلىيە فوندستېرلىڭى", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "ئىسرائىل شېكېلى (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "ئىسرائىل يېڭى شېكېلى", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ھىندىستان رۇپىسى", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ئىراق دىنارى", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "ئىران رىيالى", Symbol: ""},
				currency.ISJ: cldr.Currency{DisplayName: "ئىسلاندىيە كروناسى (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "ئىسلاندىيە كروناسى", Symbol: "kr"},
				currency.ITL: cldr.Currency{DisplayName: "ئىتالىيە لىراسى", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "يامايكا دوللىرى", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "ئىيوردانىيە دىنارى", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "ياپونىيە يېنى", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "كېنىيە شىللىڭى", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "قىرغىزىستان سومى", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "كامبودژا رىئېلى", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "كومورو فرانكى", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "شىمالىي كورېيە ۋونى", Symbol: "₩"},
				currency.KRH: cldr.Currency{DisplayName: "جەنۇبىي كورېيە خۋانى (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "جەنۇبىي كورېيە ۋونى (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "جەنۇبىي كورېيە ۋونى", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "كۇۋەيت دىنارى", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "كايمان ئاراللىرى دوللىرى", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "قازاقىستان تەڭگىسى", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "لائوس كىپى", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "لىۋان فوند سىتېرلىڭى", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "سىرىلانكا رۇپىسى", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "لىبېرىيە دوللىرى", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "لېسوتو لوتىسى", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "لىتۋا لىتاسى", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "لىتۋا تالوناسى", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "ليۇكسېمبۇرگ ئالماشتۇرۇشچان پېسوسى", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "ليۇكسېمبۇرگ فرانكى", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "لىيۇكسېمبۇرگ پۇل-مۇئامىلە فرانكى", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "لاتۋىيە لاتى", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "لاتۋىيە رۇبلىسى", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "لىۋىيە دىنارى", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "ماراكەش دىرھەمى", Symbol: ""},
				currency.MAF: cldr.Currency{DisplayName: "ماراكەش فرانكى", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "موناكو فرانكى", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "مولدوۋا كۇپونى", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "مولدوۋا لېۋى", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "ماداغاسقار ئارىئارىسى", Symbol: "Ar"},
				currency.MGF: cldr.Currency{DisplayName: "ماداغاسقار فرانكى", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "ماكېدونىيە دىنارى", Symbol: ""},
				currency.MKN: cldr.Currency{DisplayName: "ماكېدونىيە دىنارى (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "مالى فرانكى", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "مىيانمار كىياتى", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "موڭغۇلىيە تۈگرىكى", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "ئاۋمېن پاتاكاسى", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "ماۋرىتانىيە ئۇگىيەسى (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "ماۋرىتانىيە ئۇگىيەسى", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "مالتا لىراسى", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "مالتا فوندستېرلىڭى", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "ماۋرىتىئۇس رۇپىسى", Symbol: "Rs"},
				currency.MVP: cldr.Currency{DisplayName: "مالدىۋى رۇپىسى", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "مالدىۋى رۇفىياسى", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "مالاۋى كۋاچاسى", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "مېكسىكا پېسوسى", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "مېكسىكا كۈمۈش پېسوسى (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "مېكسىكا مەبلەغ بىرلىكى", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "مالايشىيا رىڭگىتى", Symbol: "RM"},
				currency.MZE: cldr.Currency{DisplayName: "موزامبىك ئېسكۇدوسى", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "موزامبىك مېتىكالى (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "موزامبىك مېتىكالى", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "نامىبىيە دوللىرى", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "نىگېرىيە نايراسى", Symbol: "₦"},
				currency.NIC: cldr.Currency{DisplayName: "نىگېرىيە كوردوباسى (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "نىگېرىيە كوردوباسى", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "گوللاندىيە گۈلدىنى", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "نورۋېگىيە كرونى", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "نېپال رۇپىسى", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "يېڭى زېلاندىيە دوللىرى", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ئومان رىيالى", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "پاناما بالبوئاسى", Symbol: ""},
				currency.PEI: cldr.Currency{DisplayName: "پېرۇ ئىنتىسى", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "پېرۇ سولى", Symbol: ""},
				currency.PES: cldr.Currency{DisplayName: "پېرۇ سولى (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "پاپۇئا يېڭى گىۋىنېيە كىناسى", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "فىلىپپىن پېسوسى", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "پاكىستان رۇپىسى", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "پولشا زىلوتى", Symbol: "zł"},
				currency.PLZ: cldr.Currency{DisplayName: "پولشا زىلوتى (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "پورتۇگالىيە ئېسكۇدوسى", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "پاراگۋاي گۇئارانىسى", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "قاتار رىيالى", Symbol: ""},
				currency.RHD: cldr.Currency{DisplayName: "رودېزىيە دوللىرى", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "رۇمىنىيە لېيى (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "رۇمىنىيە لېيى", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "سېربىيە دىنارى", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "رۇسىيە رۇبلىسى", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "رۇسىيە رۇبلىسى (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "رۋاندا فرانكى", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "سەئۇدى رىيالى", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "سولومون ئاراللىرى دوللىرى", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "سېيشېل رۇپىسى", Symbol: ""},
				currency.SDD: cldr.Currency{DisplayName: "سۇدان دىنارى (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "سۇدان فوندستېرلىڭى", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "سۇدان فوندستېرلىڭى (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "شىۋېتسىيە كروناسى", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "سىنگاپور دوللىرى", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "ساينىت-ھېلېنا فوندستېرلىڭى", Symbol: "£"},
				currency.SIT: cldr.Currency{DisplayName: "سىلوۋېنىيە تولارى", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "سىلوۋاكىيە كورۇناسى", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "سېررالېئون لېئونېسى", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "سومالى شىللىڭى", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "سۇرىنام دوللىرى", Symbol: "$"},
				currency.SRG: cldr.Currency{DisplayName: "سۇرىنام گۈلدىنى", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "جەنۇبىي سۇدان فوندستېرلىڭى", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "سان-تومې ۋە پىرىنسىپى دوبراسى (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "سان-تومې ۋە پىرىنسىپى دوبراسى", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "سوۋىت رۇبلىسى", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "سالۋادور كولونى", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "سۈرىيە فوندستېرلىڭى", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "سىۋېزىلاند لىلانگېنى", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "تايلاند باختى", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "تاجىكىستان رۇبلىسى", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "تاجىكىستان سومونىسى", Symbol: ""},
				currency.TMM: cldr.Currency{DisplayName: "تۈركمەنىستان ماناتى (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "تۈركمەنىستان ماناتى", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "تۇنىس دىنارى", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "تونگا پائانگاسى", Symbol: "T$"},
				currency.TPE: cldr.Currency{DisplayName: "تىمور ئېسكۇدوسى", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "تۈركىيە لىراسى (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "تۈركىيە لىراسى", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "تىرىنىداد ۋە توباگو دوللىرى", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "يېڭى تەيۋەن دوللىرى", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "تانزانىيە شىللىڭى", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "ئۇكرائىنا خرىۋناسى", Symbol: "₴"},
				currency.UAK: cldr.Currency{DisplayName: "ئۇكرائىنا كاربوۋانېتسى", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "ئۇگاندا شىللىڭى (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "ئۇگاندا شىللىڭى", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "ئامېرىكا دوللىرى", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "ئامېرىكا دوللىرى (كېيىنكى كۈن)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "ئامېرىكا دوللىرى (ئوخشاش كۈن)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "ئۇرۇگۋاي پېسوسى (ئىندېكىسلاش بىرلىكى)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "ئۇرۇگۋاي پېسوسى (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "ئۇرۇگۋاي پېسوسى", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "ئۆزبېكىستان سومى", Symbol: ""},
				currency.VEB: cldr.Currency{DisplayName: "ۋېنېزۇئېلا بولىۋارى (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "ۋېنېزۇئېلا بولىۋارى (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "ۋېنېزۇئېلا بولىۋارى", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "ۋىيېتنام دوڭى", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "ۋىيېتنام دوڭى (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "ۋانۇئاتۇ ۋاتۇسى", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "ساموئا تالاسى", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "ئافرىقا قىتئەسى پۇل-مۇئامىلە ئىتتىپاقى فرانكى", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "كۈمۈش", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "ئالتۇن", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "ياۋروپا مۇرەككەپ بىرلىكى", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "ياۋروپا پۇل بىرلىكى (XBB)", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "ياۋروپا ھېسابات بىرلىكى (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "ياۋروپا ھېسابات بىرلىكى (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "شەرقىي كارىب دوللىرى", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "ئالاھىدە پۇل ئېلىش ھوقۇقى", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "ياۋروپا پۇل بىرلىكى", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "فىرانسىيە ئالتۇن فرانكى", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "فىرانسىيە UIC فرانكى", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "ئافرىقا قىتئەسى پۇل-مۇئامىلە ئىتتىپاقى فرانكى (BCEAO)", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "پاللادىي", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "تىنچ ئوكيان پۇل-مۇئامىلە ئورتاق گەۋدىسى فرانكى", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "پىلاتىنا", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET فوندى", Symbol: ""},
				currency.XSU: cldr.Currency{DisplayName: "سۇكرې", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "پۇل سىناش بىرلىكى", Symbol: ""},
				currency.XUA: cldr.Currency{DisplayName: "ئاسىيا تەرەققىيات بانكىسى ھېسابات بىرلىكى", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "يوچۇن پۇل", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "يەمەن دىنارى", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "يەمەن رىيالى", Symbol: ""},
				currency.YUD: cldr.Currency{DisplayName: "يۇگوسلاۋىيە قاتتىق دىنارى (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "يۇگوسلاۋىيە يېڭى دىنارى (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "يۇگوسلاۋىيە ئالماشتۇرۇشچان دىنارى (1990–1992)", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "يۇگوسلاۋىيە ئىسلاھات دىنارى (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "جەنۇبىي ئافرىقا راندى (پۇل–مۇئامىلە)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "جەنۇبىي ئافرىقا راندى", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "زامبىيە كۋاچاسى (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "زامبىيە كۋاچاسى", Symbol: "ZK"},
				currency.ZRN: cldr.Currency{DisplayName: "زايىر يېڭى زايىرى (1993–1998)", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "زايىر زايىرى (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "زىمبابۋې دوللىرى (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "زىمبابۋې دوللىرى (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "زىمبابۋې دوللىرى (2008)", Symbol: ""},
			},
		},
	}
}
