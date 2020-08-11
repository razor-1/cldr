// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_ps() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ps",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE د y د MMMM d", Long: "د y د MMMM d", Medium: "y MMM d", Short: "y/M/d"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جنوري", Feb: "فبروري", Mar: "مارچ", Apr: "اپریل", May: "مۍ", Jun: "جون", Jul: "جولای", Aug: "اګست", Sep: "سپتمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "جنوري", Feb: "فېبروري", Mar: "مارچ", Apr: "اپریل", May: "مۍ", Jun: "جون", Jul: "جولای", Aug: "اګست", Sep: "سپتمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "يونۍ", Mon: "دونۍ", Tue: "درېنۍ", Wed: "څلرنۍ", Thu: "پينځنۍ", Fri: "جمعه", Sat: "اونۍ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "يونۍ", Mon: "دونۍ", Tue: "درېنۍ", Wed: "څلرنۍ", Thu: "پينځنۍ", Fri: "جمعه", Sat: "اونۍ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "يونۍ", Mon: "دونۍ", Tue: "درېنۍ", Wed: "څلرنۍ", Thu: "پينځنۍ", Fri: "جمعه", Sat: "اونۍ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "غ.م.", PM: "غ.و."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "غ.م.", PM: "غ.و."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "غ.م.", PM: "غ.و."}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "٪", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "0K", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "متحده عرب امارات درهم", Symbol: ""},
				currency.AFA: cldr.Currency{DisplayName: "افغانۍ (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "افغانۍ", Symbol: "؋"},
				currency.ALL: cldr.Currency{DisplayName: "البانوي لک", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "ارمينيايي ډرام", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "هالېنډي انټيليايي ګيلډر", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "انګولي کوانزا", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "ارجنټاين پسو", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "آسترالوي ډالر", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "اروبايي فلورن", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "آزربايجاني منت", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "بوسنيا هرزګوينيايي بدلېدونکې مارک", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "باربيډين ډالر", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "بنګالۍ ټاکه", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "بلغاري ليو", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "بحريني دينار", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "برونډي فرانک", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "برمودا ډالر", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "برونايي ډالر", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "بوليوي بوليويانو", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "برازيلي ريل", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "بهاماسي ډالر", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "بهوټانۍ انګولټرم", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "بوټسواني پولا", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "بلاروسي روبل", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "بليز ډالر", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "کاناډايي ډالر", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "کانګولي فرانک", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "سويسي فرانک", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "چلي پسو", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "چيني يوان (آف شور)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "چيني يوان", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "کولمبين پسو", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "کوسټا ريکن کولون", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "کيوبايي بدلېدونکي پسو", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "کيوبايي پسو", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "کيپ وردين اسکوډو", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "چيک کرونا", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "جبوتي فرانک", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "ډنمارکي کرون", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "دومينيکا پسو", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "الجيرين دينار", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "مصري پونډ", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "ايريټرين نکفا", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ايتهوپيايي بر", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "يورو", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "فجي ډالر", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "پاکلېنډ ټاپوګانو پونډ", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "برتانوې پونډ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "جارجیاېي لارې", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "ګانين سيډي", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "جبل الطارقي پونډ", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "ګيمبين دلاسې", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "ګنې فرانک", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "ګويټيمالن کوټزل", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "ګايانيز ډالر", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "هانګ کانګ ډالر", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "هونډوران ليمپيرا", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "کروشيايي کونا", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "هيټي ګورډ", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "هنګري فورنټ", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "انډونيشي روپيا", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "اسرايلي نيو شيکل", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "هندي روپۍ", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "عراقي دينار", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "ايراني ريال", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "آيسلېنډي کرونا", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "جمايکايي ډالر", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "اردني دينار", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "جاپاني ين", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "کينيايي شيلنګ", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "کرغزستاني سوم", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "کمبوډي ريل", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "کوموري فرانک", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "شمالي کوريايي وان", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "جنوبي کوريايي وان", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "کويتي دينار", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "کيمن ټاپوګانو ډالر", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "قازقستاني ټينج", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "لاشې کپ", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "لبناني پونډ", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "سري لنکن روپۍ", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "لايبيرين ډالر", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "ليبياېي دينار", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "مراکشي درهم", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "مالډوي ليو", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "ملاګاسي ارياري", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "مسيډونايي دينار", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "ميانماري کيات", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "منګوليايي توګريک", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "مکانيز پټاکا", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "موريشيسي ډالر", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "موريشيسي روپۍ", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "مالديپي روپيا", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "ملاوي کواچا", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "ميکسيکن پيسو", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "ملايشي رنګټ", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "موزمبيقي ميټيکل", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "نيميبيايي ډالر", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "نايجيري نايرا", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "نيکاراګون کورډوبا", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "نارويجين کرون", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "نيپالي روپۍ", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "نيوزي لينډي ډالر", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "عماني ريال", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "پانامۍ بالبوا", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "پيروين سول", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "پاپوا نيوګاني کينا", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "فلپاينۍ پسو", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "پاکستانۍ کلداره", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "پولينډي زلوټي", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "پيراګوين ګوراني", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "قطري ريال", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "روماني ليو", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "سربيايي دينار", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "روسي روبل", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "روانډي فرانک", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "سعودي ريال", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "سولومن ټاپوګانو ډالر", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "سيچيليسي روپۍ", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "سوډاني پونډ", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "سويډني کرونا", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "سنګاپور ډالر", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "سينټ هيلينا پونډ", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "سيرا ليوني ليون", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "سومالي شيلنګ", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "سورينيمي ډالر", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "جنوب سوډاني پونډ", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "ساو ټوم او پرينسپي ډوبرا", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "سوريايي پونډ", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "سوازي ليلانګيني", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "تهايي بات", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "تاجکستاني سوموني", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "ترکمانستاني منت", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "تونسي دينار", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "ټونګن پانګا", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "ترکي ليرا", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "ټرينيډاډ او ټوباګو ډالر", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "نيو تائيواني ډالر", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "تنزاني شيلنګ", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "اوکرايني هريونيا", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "يوګانډي شيلنګ", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "امريکايي ډالر", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "يوراګوي پسو", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "ازبکستاني سوم", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "وينزويلي بوليوار", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "ويتنامي ډونګ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "ونواتو واتو", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "سموون تالا", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "مرکزي افريقايي CFA فرانک", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "ختيځ کربين ډالر", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "ختيځ افريقايي CFA فرانک", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP فرانک", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "نامعلومه مروجه پېسې", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "يمني ريال", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "جنوبي افريقاېي رنډ", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "زيمبي کواچا", Symbol: "ZK"},
			},
		},
	}
}
