// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ps_AF() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ps_AF",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE د y د MMMM d", Long: "y MMMM d", Medium: "y MMM d", Short: "y/M/d"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "جنوري", Feb: "فبروري", Mar: "مارچ", Apr: "اپریل", May: "مۍ", Jun: "جون", Jul: "جولای", Aug: "اګست", Sep: "سپتمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ج", Feb: "ف", Mar: "م", Apr: "ا", May: "م", Jun: "ج", Jul: "ج", Aug: "ا", Sep: "س", Oct: "ا", Nov: "ن", Dec: "د"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "جنوري", Feb: "فېبروري", Mar: "مارچ", Apr: "اپریل", May: "مۍ", Jun: "جون", Jul: "جولای", Aug: "اګست", Sep: "سپتمبر", Oct: "اکتوبر", Nov: "نومبر", Dec: "دسمبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "يونۍ", Mon: "دونۍ", Tue: "درېنۍ", Wed: "څلرنۍ", Thu: "پينځنۍ", Fri: "جمعه", Sat: "اونۍ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "غ.م.", PM: "غ.و."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "غ.م.", PM: "غ.و."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "غ.م.", PM: "غ.و."}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "\u200e−", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "متحده عرب امارات درهم", Symbol: ""},
				currency.AFA: cldr.Currency{DisplayName: "افغانۍ (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "افغانۍ", Symbol: "؋"},
				currency.ALL: cldr.Currency{DisplayName: "البانوي لک", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "ارمينيايي ډرام", Symbol: "֏"},
				currency.ANG: cldr.Currency{DisplayName: "هالېنډي انټيليايي ګيلډر", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "انګولي کوانزا", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "ارجنټاين پسو", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "آسترالوي ډالر", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "اروبايي فلورن", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "آزربايجاني منت", Symbol: "₼"},
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
				currency.GHS: cldr.Currency{DisplayName: "ګانين سيډي", Symbol: "GH₵"},
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
				currency.KGS: cldr.Currency{DisplayName: "کرغزستاني سوم", Symbol: "⃀"},
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
				currency.LSL: cldr.Currency{DisplayName: "Lesotho Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "ليبياېي دينار", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "مراکشي درهم", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "مالډوي ليو", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "ملاګاسي ارياري", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "مسيډونايي دينار", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "ميانماري کيات", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "منګوليايي توګريک", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "مکانيس پټاکا", Symbol: ""},
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
				currency.PHP: cldr.Currency{DisplayName: "فلپاينۍ پیسو", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "پاکستانۍ کلداره", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "پولينډي زلوټي", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "پيراګوين ګوراني", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "قطري ريال", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "روماني ليو", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "سربيايي دينار", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "روسي روبل", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "روانډي فرانک", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "سعودي ريال", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "سولومن ټاپوګانو ډالر", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "سيچيليسي روپۍ", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "سوډاني پونډ", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "سويډني کرونا", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "سنګاپور ډالر", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "سينټ هيلينا پونډ", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "سيرا ليوني ليون", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "سيرا ليوني ليون - 1964-2022", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "سومالي شيلنګ", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "سورينيمي ډالر", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "جنوب سوډاني پونډ", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "ساو ټوم او پرينسپي ډوبرا", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "سوريايي پونډ", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "سوازي ليلانګيني", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "تهايي بات", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "تاجکستاني سوموني", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "ترکمانستاني منت", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "تيونسې دينار", Symbol: ""},
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
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "ختيځ افريقايي CFA فرانک", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP فرانک", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "نامعلومه مروجه پېسې", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "يمني ريال", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "جنوبي افريقاېي رنډ", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "زيمبي کواچا", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "افري",
			language.AB:      "ابخازي",
			language.ACE:     "اچيني",
			language.ADA:     "ادانگمي",
			language.ADY:     "اديغي",
			language.AF:      "افریکانسي",
			language.AGQ:     "اغیمي",
			language.AIN:     "اينويي",
			language.AK:      "اکاني",
			language.ALE:     "اليوتي",
			language.ALT:     "سویل الټای",
			language.AM:      "امهاري",
			language.AN:      "اراگونېسي",
			language.ANN:     "Obo",
			language.ANP:     "انگيکي",
			language.AR:      "عربي",
			language.AR_001:  "نوې معياري عربي",
			language.ARN:     "ماپوچه",
			language.ARP:     "اراپاهوي",
			language.ARS:     "نجدی عربی",
			language.AS:      "اسامي",
			language.ASA:     "اسويي",
			language.AST:     "استورياني",
			language.ATJ:     "اتیکامیکو",
			language.AV:      "اواري",
			language.AWA:     "اوادي",
			language.AY:      "ایماري",
			language.AZ:      "اذري",
			language.BA:      "باشکير",
			language.BAL:     "بلوڅي",
			language.BAN:     "بالنی",
			language.BAS:     "باسا",
			language.BE:      "بېلاروسي",
			language.BEM:     "بيمبا",
			language.BEZ:     "بينا",
			language.BG:      "بلغاري",
			language.BGC:     "هریانوی",
			language.BHO:     "بهوجپوري",
			language.BI:      "بسلاما",
			language.BIN:     "بینی",
			language.BLA:     "سکسيکا",
			language.BLO:     "انۍ",
			language.BM:      "بمبارا",
			language.BN:      "بنگالي",
			language.BO:      "تبتي",
			language.BR:      "برېتون",
			language.BRX:     "بودو",
			language.BS:      "بوسني",
			language.BUG:     "بگنيايي",
			language.BYN:     "بلین",
			language.CA:      "کټلاني",
			language.CAY:     "Ca",
			language.CCP:     "چکما",
			language.CE:      "چيچني",
			language.CEB:     "سیبوانوي",
			language.CGG:     "چيگايي",
			language.CH:      "چمورو",
			language.CHK:     "چواوکي",
			language.CHM:     "ماري",
			language.CHO:     "چوکټاوي",
			language.CHP:     "Ch",
			language.CHR:     "چېروکي",
			language.CHY:     "شيني",
			language.CKB:     "منځنۍ کوردي",
			language.CLC:     "چیلکوټین",
			language.CO:      "کورسيکاني",
			language.CRG:     "mc",
			language.CRJ:     "سویل ختیځ کری",
			language.CRK:     "پلینز کری",
			language.CRL:     "شمالي ختیځ کری",
			language.CRM:     "mcr",
			language.CRR:     "Car Alg",
			language.CRS:     "سسيلوا ڪروئل فرانسوي",
			language.CS:      "چېکي",
			language.CSW:     "سومپی کری",
			language.CU:      "د کليسا سلاوي",
			language.CV:      "چوواشي",
			language.CY:      "ويلشي",
			language.DA:      "ډنمارکي",
			language.DAK:     "داکوتا",
			language.DAR:     "درگوا",
			language.DAV:     "ټایټا",
			language.DE:      "الماني",
			language.DE_AT:   "اتريشي آلماني",
			language.DE_CH:   "سویس های جرمن",
			language.DGR:     "داگرب",
			language.DJE:     "زرما",
			language.DOI:     "ډوګري",
			language.DSB:     "کښته سربيايي",
			language.DUA:     "دوالا",
			language.DV:      "ديویهی",
			language.DYO:     "جولا فوني",
			language.DZ:      "ژونگکه",
			language.DZG:     "ډزاګا",
			language.EBU:     "ايمبو",
			language.EE:      "ايو",
			language.EFI:     "افک",
			language.EKA:     "اکجک",
			language.EL:      "یوناني",
			language.EN:      "انګليسي",
			language.EN_AU:   "آسټرالياوي انګليسي",
			language.EN_CA:   "کاناډايي انګلیسي",
			language.EN_GB:   "یو کي انګلیسي",
			language.EN_US:   "د متحده آيالاتو انګليسي",
			language.EO:      "اسپرانتو",
			language.ES:      "هسپانوي",
			language.ES_419:  "لاتيني امريکايي هسپانوي",
			language.ES_ES:   "اروپايي هسپانوي",
			language.ES_MX:   "ميکسيکي هسپانوي",
			language.ET:      "حبشي",
			language.EU:      "باسکي",
			language.EWO:     "اوونڊو",
			language.FA:      "فارسي",
			language.FA_AF:   "دری (افغانستان)",
			language.FF:      "فولاح",
			language.FI:      "فینلنډي",
			language.FIL:     "فلیپیني",
			language.FJ:      "فجیان",
			language.FO:      "فاروئې",
			language.FON:     "فان",
			language.FR:      "فرانسوي",
			language.FR_CA:   "کاناډايي فرانسوي",
			language.FR_CH:   "سويسي فرانسوي",
			language.FRC:     "کاجون فرانسوی",
			language.FRR:     "شمالي فریسیان",
			language.FUR:     "فرائیلیین",
			language.FY:      "لوېديځ فريشي",
			language.GA:      "ائيرلېنډي",
			language.GAA:     "Ga",
			language.GD:      "سکاټلېنډي ګېلک",
			language.GEZ:     "ګیز",
			language.GIL:     "گلبرتي",
			language.GL:      "ګلېشيايي",
			language.GN:      "ګوراني",
			language.GOR:     "ګورن ټالو",
			language.GSW:     "سویس جرمن",
			language.GU:      "ګجراتي",
			language.GUZ:     "ګوسي",
			language.GV:      "مینکس",
			language.GWI:     "ګیچین",
			language.HA:      "هوسا",
			language.HAI:     "ha",
			language.HAW:     "هوایی",
			language.HAX:     "جنوبي هایدا",
			language.HE:      "عبراني",
			language.HI:      "هندي",
			language.HI_LATN: "هنګلش",
			language.HIL:     "ھلیګینون",
			language.HMN:     "همونګ",
			language.HR:      "کروايشيايي",
			language.HSB:     "پورته سربيايي",
			language.HT:      "هيټي کريول",
			language.HU:      "هنګري",
			language.HUP:     "ھوپا",
			language.HUR:     "هلکومیلم",
			language.HY:      "آرمينيايي",
			language.HZ:      "هیرورو",
			language.IA:      "انټرلنګوا",
			language.IBA:     "ابن",
			language.IBB:     "ابیبیو",
			language.ID:      "انډونېزي",
			language.IE:      "آسا نا جبة",
			language.IG:      "اګبو",
			language.II:      "سیچیان یی",
			language.IKT:     "مغربی کینیډین انوکټیټ",
			language.ILO:     "الوکو",
			language.INH:     "انگش",
			language.IO:      "اڊو",
			language.IS:      "ايسلنډي",
			language.IT:      "ایټالوي",
			language.IU:      "انوکتیتوت",
			language.JA:      "جاپاني",
			language.JBO:     "لوجبان",
			language.JGO:     "نګومبا",
			language.JMC:     "ماچمی",
			language.JV:      "جاوايي",
			language.KA:      "جورجيائي",
			language.KAB:     "کیبیل",
			language.KAC:     "کاچین",
			language.KAJ:     "ججو",
			language.KAM:     "کامبا",
			language.KBD:     "کابیرین",
			language.KCG:     "تایپ",
			language.KDE:     "ميکونډي",
			language.KEA:     "کابوورډیانو",
			language.KFO:     "کورو",
			language.KGP:     "کینګا",
			language.KHA:     "خاسې",
			language.KHQ:     "کویرا چینی",
			language.KI:      "ککوؤو",
			language.KJ:      "کواناما",
			language.KK:      "قازق",
			language.KKJ:     "کاکو",
			language.KL:      "کالالیست",
			language.KLN:     "کلینجن",
			language.KM:      "خمر",
			language.KMB:     "کیمبوندو",
			language.KN:      "کناډا",
			language.KO:      "کوریایی",
			language.KOK:     "کونکاني",
			language.KPE:     "کیلي",
			language.KR:      "کنوری",
			language.KRC:     "کراچی بالکر",
			language.KRL:     "کاریلین",
			language.KRU:     "کورخ",
			language.KS:      "کشمیري",
			language.KSB:     "شمبالا",
			language.KSF:     "بفیا",
			language.KSH:     "کولوګنيايي",
			language.KU:      "کردي",
			language.KUM:     "کومک",
			language.KV:      "کومی",
			language.KW:      "کورنيشي",
			language.KWK:     "Vote kwk",
			language.KXV:     "کووئ",
			language.KY:      "کرغيزي",
			language.LA:      "لاتیني",
			language.LAD:     "لاډینو",
			language.LAG:     "لنګی",
			language.LB:      "لوګزامبورګي",
			language.LEZ:     "لیګغیان",
			language.LG:      "ګانده",
			language.LI:      "لمبرگیانی",
			language.LIJ:     "لینګورین",
			language.LIL:     "lill",
			language.LKT:     "لکوټا",
			language.LMO:     "لومبارډ",
			language.LN:      "لنګالا",
			language.LO:      "لاو",
			language.LOU:     "Louis",
			language.LOZ:     "لوزی",
			language.LRC:     "شمالي لوری",
			language.LSM:     "سامیه",
			language.LT:      "ليتواني",
			language.LU:      "لوبا-کټنګا",
			language.LUA:     "لبا لولوا",
			language.LUN:     "لندا",
			language.LUO:     "لو",
			language.LUS:     "ميزو",
			language.LUY:     "لویا",
			language.LV:      "لېټواني",
			language.MAD:     "مدراسی",
			language.MAG:     "مګهي",
			language.MAI:     "مایتھلي",
			language.MAK:     "مکاسار",
			language.MAS:     "ماسائي",
			language.MDF:     "موکشا",
			language.MEN:     "مینڊي",
			language.MER:     "ميرو",
			language.MFE:     "ماریسیسن",
			language.MG:      "ملغاسي",
			language.MGH:     "مکھوامیتو",
			language.MGO:     "ميټا",
			language.MH:      "مارشلیز",
			language.MI:      "ماوري",
			language.MIC:     "ممکق",
			language.MIN:     "مينيگاباو",
			language.MK:      "مقدوني",
			language.ML:      "مالايالم",
			language.MN:      "منګولیایی",
			language.MNI:     "مانی پوری",
			language.MOE:     "mo",
			language.MOH:     "محاواک",
			language.MOS:     "ماسي",
			language.MR:      "مراټهي",
			language.MS:      "ملایا",
			language.MT:      "مالټايي",
			language.MUA:     "مندانګ",
			language.MUL:     "متعدد ژبې",
			language.MUS:     "کريکي",
			language.MWL:     "مرانديز",
			language.MY:      "برمایی",
			language.MYV:     "ارزيا",
			language.MZN:     "مزاندراني",
			language.NA:      "نایرو",
			language.NAP:     "نيپالين",
			language.NAQ:     "ناما",
			language.NB:      "ناروېئي (بوکمال)",
			language.ND:      "شمالي نديبل",
			language.NDS:     "کښته آلماني",
			language.NE:      "نېپالي",
			language.NEW:     "نيواري",
			language.NG:      "ندونگا",
			language.NIA:     "نياس",
			language.NIU:     "نیان",
			language.NL:      "هالېنډي",
			language.NL_BE:   "فلېمېشي",
			language.NMG:     "کواسیو",
			language.NN:      "ناروېئي (نائنورسک)",
			language.NNH:     "نایجیمون",
			language.NO:      "ناروېئي",
			language.NOG:     "نوګی",
			language.NQO:     "نکو",
			language.NR:      "سويلي نديبيل",
			language.NSO:     "شمالي سوتو",
			language.NUS:     "نویر",
			language.NV:      "نواجو",
			language.NY:      "نیانجا",
			language.NYN:     "نینکول",
			language.OC:      "اوکسيټاني",
			language.OJB:     "شمال لویدیځ اوجیبوا",
			language.OJC:     "مرکزي اوجیبوا",
			language.OJS:     "اوجي-کري",
			language.OJW:     "لویدیځ اوجیبوا",
			language.OKA:     "اوکاګان",
			language.OM:      "اورومو",
			language.OR:      "اوڊيا",
			language.OS:      "اوسيټک",
			language.PA:      "پنجابي",
			language.PAG:     "پانګاسین",
			language.PAM:     "پمپانگا",
			language.PAP:     "پاپيامينتو",
			language.PAU:     "پالان",
			language.PCM:     "نائجیریا پیدجن",
			language.PIS:     "پیجین",
			language.PL:      "پولنډي",
			language.PQM:     "mpq",
			language.PRG:     "پروشين",
			language.PS:      "پښتو",
			language.PT:      "پورتګالي",
			language.PT_BR:   "برازیلي پرتګالي",
			language.PT_PT:   "اروپايي پرتګالي",
			language.QU:      "کېچوا",
			language.QUC:     "کچی",
			language.RAJ:     "راجستھانی",
			language.RAP:     "رپانوئي",
			language.RAR:     "راروټانګان",
			language.RHG:     "روهینګیا",
			language.RM:      "رومانیش",
			language.RN:      "رونډی",
			language.RO:      "رومانیایی",
			language.RO_MD:   "مولداویایی",
			language.ROF:     "رومبو",
			language.RU:      "روسي",
			language.RUP:     "اروماني",
			language.RW:      "کینیارونډا",
			language.RWK:     "روا",
			language.SA:      "سنسکریټ",
			language.SAD:     "سنډاوی",
			language.SAH:     "سخا",
			language.SAQ:     "سمبورو",
			language.SAT:     "سنتالي",
			language.SBA:     "نګبای",
			language.SBP:     "سانګوو",
			language.SC:      "سارڊيني",
			language.SCN:     "سیلیسي",
			language.SCO:     "سکاټس",
			language.SD:      "سندهي",
			language.SE:      "شمالي سامي",
			language.SEH:     "سینا",
			language.SES:     "کوییرابورو سینی",
			language.SG:      "سانګو",
			language.SH:      "سرب-کروشيايي",
			language.SHI:     "تاکلهیټ",
			language.SHN:     "شان",
			language.SI:      "سينهالي",
			language.SK:      "سلوواکي",
			language.SL:      "سلوواني",
			language.SLH:     "سویلي لوشوټسید",
			language.SM:      "ساموآن",
			language.SMA:     "سویلي سامی",
			language.SMJ:     "لول سامي",
			language.SMN:     "اناري سميع",
			language.SMS:     "سکولټ سمیع",
			language.SN:      "شونا",
			language.SNK:     "سونینګ",
			language.SO:      "سومالي",
			language.SQ:      "الباني",
			language.SR:      "سربيائي",
			language.SRN:     "سوران ټونګو",
			language.SS:      "سواتی",
			language.SSY:     "سهو",
			language.ST:      "سويلي سوتو",
			language.STR:     "سټریټ سیلش",
			language.SU:      "سوډاني",
			language.SUK:     "سکوما",
			language.SV:      "سویډنی",
			language.SW:      "سواهېلي",
			language.SW_CD:   "کانګو سواهلی",
			language.SWB:     "کومورياني",
			language.SYR:     "سوریاني",
			language.SZL:     "سیلیسیان",
			language.TA:      "تامل",
			language.TCE:     "جنوبي توچون",
			language.TE:      "تېليګو",
			language.TEM:     "تیمني",
			language.TEO:     "تیسو",
			language.TET:     "تتوم",
			language.TG:      "تاجکي",
			language.TGX:     "ټګش",
			language.TH:      "تايلېنډي",
			language.THT:     "طهلتان",
			language.TI:      "تيګريني",
			language.TIG:     "تیګر",
			language.TK:      "ترکمني",
			language.TLH:     "کلينګاني",
			language.TLI:     "ټلینګیت",
			language.TN:      "سووانا",
			language.TO:      "تونګان",
			language.TOK:     "توکی پونا",
			language.TPI:     "توک پیسین",
			language.TR:      "ترکي",
			language.TRV:     "تاروکو",
			language.TS:      "سونګا",
			language.TT:      "تاتار",
			language.TTM:     "شمالي ټچون",
			language.TUM:     "تامبوکا",
			language.TVL:     "تووالو",
			language.TW:      "توی",
			language.TWQ:     "تساواق",
			language.TY:      "تاهیتي",
			language.TYV:     "توینیان",
			language.TZM:     "مرکزی اطلس تمازائيٹ",
			language.UDM:     "ادمورت",
			language.UG:      "اويغوري",
			language.UK:      "اوکرايني",
			language.UMB:     "امبوندو",
			language.UND:     "نامعلومه ژبه",
			language.UR:      "اردو",
			language.UZ:      "اوزبکي",
			language.VAI:     "وای",
			language.VE:      "ویندا",
			language.VEC:     "وینټیان",
			language.VI:      "وېتنامي",
			language.VMW:     "مکوه",
			language.VO:      "والاپوک",
			language.VUN:     "وونجو",
			language.WA:      "والون",
			language.WAE:     "ولسیر",
			language.WAL:     "ولایټا",
			language.WAR:     "وارۍ",
			language.WO:      "ولوف",
			language.WUU:     "وو چینایی",
			language.XAL:     "کالمک",
			language.XH:      "خوسا",
			language.XNR:     "کانګرو",
			language.XOG:     "سوګا",
			language.YAV:     "ینګبین",
			language.YBB:     "یمبا",
			language.YI:      "يديش",
			language.YO:      "یوروبا",
			language.YRL:     "نینګاتو",
			language.YUE:     "چايني، کانټونيز",
			language.ZA:      "ژوانګ",
			language.ZGH:     "معياري مراکشي تمازيټ",
			language.ZH:      "چيني، ماندرين",
			language.ZH_HANS: "چيني ماندرين چيني",
			language.ZH_HANT: "دوديزه ماندرين چيني",
			language.ZU:      "زولو",
			language.ZUN:     "زوني",
			language.ZXX:     "نه ژبني منځپانګه",
			language.ZZA:     "زازا",
		},
		Territories: cldr.Territories{
			territory.V_001: "نړۍ",
			territory.V_002: "افريقا",
			territory.V_003: "شمالی امریکا",
			territory.V_005: "سويلي امريکا",
			territory.V_009: "اوقيانوسيه",
			territory.V_011: "لویدیځ افریقا",
			territory.V_013: "منخنۍ امريکا",
			territory.V_014: "ختیځ افریقا",
			territory.V_015: "شمالي افریقا",
			territory.V_017: "منځنۍ افریقا",
			territory.V_018: "سويلي افريقا",
			territory.V_019: "امريکې",
			territory.V_021: "شمالي امریکا",
			territory.V_029: "کیریبین",
			territory.V_030: "ختیځ آسیا",
			territory.V_034: "سويلي آسيا",
			territory.V_035: "سويلي ختيځ آسيا",
			territory.V_039: "سويلي اروپا",
			territory.V_053: "آسترالیا",
			territory.V_054: "ملانشیا",
			territory.V_057: "د مایکرونیسینین سیمه",
			territory.V_061: "پولنيسيا",
			territory.V_142: "آسيا",
			territory.V_143: "منځنۍ آسيا",
			territory.V_145: "لویدیځ آسیا",
			territory.V_150: "اروپا",
			territory.V_151: "ختيځ اروپا",
			territory.V_154: "شمالي اروپا",
			territory.V_155: "لوېديځ اروپا",
			territory.V_202: "سب سهارن افريقا",
			territory.V_419: "لاتیني امریکا",
			territory.AC:    "اسينشان ټاپو",
			territory.AD:    "اندورا",
			territory.AE:    "متحده عرب امارات",
			territory.AF:    "افغانستان",
			territory.AG:    "انټيګوا او باربودا",
			territory.AI:    "انګیلا",
			territory.AL:    "البانیه",
			territory.AM:    "ارمنستان",
			territory.AO:    "انګولا",
			territory.AQ:    "انتارکتیکا",
			territory.AR:    "ارجنټاين",
			territory.AS:    "امریکایی ساماوا",
			territory.AT:    "اتریش",
			territory.AU:    "آسټرالیا",
			territory.AW:    "آروبا",
			territory.AX:    "الاند ټاپوان",
			territory.AZ:    "اذربايجان",
			territory.BA:    "بوسنيا او هېرزګوينا",
			territory.BB:    "باربادوس",
			territory.BD:    "بنگله دېش",
			territory.BE:    "بیلجیم",
			territory.BF:    "بورکینا فاسو",
			territory.BG:    "بلغاریه",
			territory.BH:    "بحرين",
			territory.BI:    "بروندي",
			territory.BJ:    "بینن",
			territory.BL:    "سينټ بارتيلمي",
			territory.BM:    "برمودا",
			territory.BN:    "برونائي",
			territory.BO:    "بولیویا",
			territory.BQ:    "کیریبین هالینډ",
			territory.BR:    "برازیل",
			territory.BS:    "باهماس",
			territory.BT:    "بهوټان",
			territory.BV:    "بوویټ ټاپو",
			territory.BW:    "بوتسوانه",
			territory.BY:    "بیلاروس",
			territory.BZ:    "بلیز",
			territory.CA:    "کاناډا",
			territory.CC:    "کوکوز (کيلنګ) ټاپوګان",
			territory.CD:    "کانګو - کینشاسا",
			territory.CF:    "وسطي افريقا جمهور",
			territory.CG:    "کانګو - بروزوییل",
			territory.CH:    "سویس",
			territory.CI:    "د عاج ساحل",
			territory.CK:    "کوک ټاپوګان",
			territory.CL:    "چیلي",
			territory.CM:    "کامرون",
			territory.CN:    "چین",
			territory.CO:    "کولمبیا",
			territory.CP:    "د کلپرټون ټاپو",
			territory.CR:    "کوستاریکا",
			territory.CU:    "کیوبا",
			territory.CV:    "کیپ ورد",
			territory.CW:    "کوراکاو",
			territory.CX:    "د کريسمس ټاپو",
			territory.CY:    "قبرس",
			territory.CZ:    "چکیا",
			territory.DE:    "المان",
			territory.DG:    "ډایګو ګارسیا",
			territory.DJ:    "جبوتي",
			territory.DK:    "ډنمارک",
			territory.DM:    "دومینیکا",
			territory.DO:    "جمهوريه ډومينيکن",
			territory.DZ:    "الجزایر",
			territory.EA:    "سيوتا او ماليلا",
			territory.EC:    "اکوادور",
			territory.EE:    "استونیا",
			territory.EG:    "مصر",
			territory.EH:    "لويديځ صحارا",
			territory.ER:    "اریتره",
			territory.ES:    "هسپانیه",
			territory.ET:    "حبشه",
			territory.EU:    "اروپايي اتحاديه",
			territory.EZ:    "اروپايي سيمه",
			territory.FI:    "فنلینډ",
			territory.FJ:    "فجي",
			territory.FK:    "فاکلينډ ټاپوګان",
			territory.FM:    "میکرونیزیا",
			territory.FO:    "فارو ټاپو",
			territory.FR:    "فرانسه",
			territory.GA:    "ګابن",
			territory.GB:    "برتانیه",
			territory.GD:    "ګرنادا",
			territory.GE:    "گورجستان",
			territory.GF:    "فرانسوي ګانا",
			territory.GG:    "ګرنسي",
			territory.GH:    "ګانا",
			territory.GI:    "جبل الطارق",
			territory.GL:    "ګرینلینډ",
			territory.GM:    "ګامبیا",
			territory.GN:    "ګینه",
			territory.GP:    "ګوادلوپ",
			territory.GQ:    "استوایی ګیني",
			territory.GR:    "یونان",
			territory.GS:    "سويلي جارجيا او سويلي سېنډوچ ټاپوګان",
			territory.GT:    "ګواتیمالا",
			territory.GU:    "ګوام",
			territory.GW:    "ګینه بیسو",
			territory.GY:    "ګیانا",
			territory.HK:    "هانګ کانګ SAR چین",
			territory.HM:    "هارډ او ميکډانلډ ټاپوګان",
			territory.HN:    "هانډوراس",
			territory.HR:    "کرواشيا",
			territory.HT:    "هایټي",
			territory.HU:    "مجارستان",
			territory.IC:    "د کناري ټاپوګان",
			territory.ID:    "اندونیزیا",
			territory.IE:    "آيرلېنډ",
			territory.IL:    "اسراييل",
			territory.IM:    "د آئل آف مین",
			territory.IN:    "هند",
			territory.IO:    "د برتانوي هند سمندري سيمه",
			territory.IQ:    "عراق",
			territory.IR:    "ايران",
			territory.IS:    "آیسلینډ",
			territory.IT:    "ایټالیه",
			territory.JE:    "جرسی",
			territory.JM:    "جمیکا",
			territory.JO:    "اردن",
			territory.JP:    "جاپان",
			territory.KE:    "کینیا",
			territory.KG:    "قرغزستان",
			territory.KH:    "کمبودیا",
			territory.KI:    "کیري باتي",
			territory.KM:    "کوموروس",
			territory.KN:    "سینټ کټس او نیویس",
			territory.KP:    "شمالی کوریا",
			territory.KR:    "سویلي کوریا",
			territory.KW:    "کويت",
			territory.KY:    "کیمان ټاپوګان",
			territory.KZ:    "قزاقستان",
			territory.LA:    "لاوس",
			territory.LB:    "لبنان",
			territory.LC:    "سینټ لوسیا",
			territory.LI:    "لیختن اشتاین",
			territory.LK:    "سريلنکا",
			territory.LR:    "لايبيريا",
			territory.LS:    "لسوتو",
			territory.LT:    "لیتوانیا",
			territory.LU:    "لوګزامبورګ",
			territory.LV:    "ليتهويا",
			territory.LY:    "لیبیا",
			territory.MA:    "مراکش",
			territory.MC:    "موناکو",
			territory.MD:    "مولدوا",
			territory.ME:    "مونټینیګرو",
			territory.MF:    "سینټ مارټن",
			territory.MG:    "مدغاسکر",
			territory.MH:    "مارشل ټاپوګان",
			territory.MK:    "شمالي مقدونيه",
			territory.ML:    "مالي",
			territory.MM:    "ميانمار (برما)",
			territory.MN:    "منګوليا",
			territory.MO:    "مکاو SAR چین",
			territory.MP:    "شمالي ماريانا ټاپوګان",
			territory.MQ:    "مارټینیک",
			territory.MR:    "موریتانیا",
			territory.MS:    "مانټیسیرت",
			territory.MT:    "مالټا",
			territory.MU:    "موریشیس",
			territory.MV:    "مالديپ",
			territory.MW:    "مالاوي",
			territory.MX:    "میکسیکو",
			territory.MY:    "مالیزیا",
			territory.MZ:    "موزمبيق",
			territory.NA:    "نیمبیا",
			territory.NC:    "نوی کالیډونیا",
			territory.NE:    "نايجير",
			territory.NF:    "نارفولک ټاپوګان",
			territory.NG:    "نایجیریا",
			territory.NI:    "نکاراګوا",
			territory.NL:    "هالېنډ",
			territory.NO:    "ناروۍ",
			territory.NP:    "نیپال",
			territory.NR:    "نایرو",
			territory.NU:    "نیوو",
			territory.NZ:    "نیوزیلنډ",
			territory.OM:    "عمان",
			territory.PA:    "پاناما",
			territory.PE:    "پیرو",
			territory.PF:    "فرانسوي پولينيسيا",
			territory.PG:    "پاپوا نيو ګيني",
			territory.PH:    "فلپين",
			territory.PK:    "پاکستان",
			territory.PL:    "پولنډ",
			territory.PM:    "سینټ پییر او میکولون",
			territory.PN:    "پيټکيرن ټاپوګان",
			territory.PR:    "پورتو ریکو",
			territory.PS:    "فلسطیني سيمې",
			territory.PT:    "پورتګال",
			territory.PW:    "پلاؤ",
			territory.PY:    "پاراګوی",
			territory.QA:    "قطر",
			territory.QO:    "بهرنۍ اوسيانه",
			territory.RE:    "ریونین",
			territory.RO:    "رومانیا",
			territory.RS:    "سربيا",
			territory.RU:    "روسیه",
			territory.RW:    "روندا",
			territory.SA:    "سعودي عربستان",
			territory.SB:    "سليمان ټاپوګان",
			territory.SC:    "سیچیلیس",
			territory.SD:    "سوډان",
			territory.SE:    "سویډن",
			territory.SG:    "سينگاپور",
			territory.SH:    "سینټ هیلینا",
			territory.SI:    "سلوانیا",
			territory.SJ:    "سوالبارد او جان ميين",
			territory.SK:    "سلواکیا",
			territory.SL:    "سییرا لیون",
			territory.SM:    "سان مارینو",
			territory.SN:    "سينيګال",
			territory.SO:    "سومالیا",
			territory.SR:    "سورینام",
			territory.SS:    "سويلي سوډان",
			territory.ST:    "ساو ټیم او پرنسیپ",
			territory.SV:    "سالوېډور",
			territory.SX:    "سینټ مارټین",
			territory.SY:    "سوریه",
			territory.SZ:    "اسواټيني",
			territory.TA:    "تریستان دا کنها",
			territory.TC:    "د ترکیې او کیکاسو ټاپو",
			territory.TD:    "چاډ",
			territory.TF:    "د فرانسې جنوبي سیمې",
			territory.TG:    "ټوګو",
			territory.TH:    "تهايلنډ",
			territory.TJ:    "تاجکستان",
			territory.TK:    "توکیلو",
			territory.TL:    "تيمور-ليسټ",
			territory.TM:    "تورکمنستان",
			territory.TN:    "تونس",
			territory.TO:    "تونګا",
			territory.TR:    "ترکي",
			territory.TT:    "ټرينيډاډ او ټوباګو",
			territory.TV:    "توالیو",
			territory.TW:    "تائيوان",
			territory.TZ:    "تنزانیا",
			territory.UA:    "اوکراین",
			territory.UG:    "یوګانډا",
			territory.UM:    "د متحده ایالاتو ټاپوګان",
			territory.UN:    "ملگري ملتونه",
			territory.US:    "متحده آيالات",
			territory.UY:    "یوروګوی",
			territory.UZ:    "اوزبکستان",
			territory.VA:    "واتیکان ښار",
			territory.VC:    "سینټ ویسنټینټ او ګرینډینز",
			territory.VE:    "وینزویلا",
			territory.VG:    "بریتانوی ویګور ټاپوګان",
			territory.VI:    "د متحده آيالاتو ورجن ټاپوګان",
			territory.VN:    "وېتنام",
			territory.VU:    "واناتو",
			territory.WF:    "والیس او فوتونا",
			territory.WS:    "ساماوا",
			territory.XA:    "جعلي خج",
			territory.XB:    "سیډو بیډی",
			territory.XK:    "کوسوو",
			territory.YE:    "یمن",
			territory.YT:    "مايوټ",
			territory.ZA:    "سویلي افریقا",
			territory.ZM:    "زیمبیا",
			territory.ZW:    "زیمبابوی",
			territory.ZZ:    "نامعلومه سيمه",
		},
	}
}
