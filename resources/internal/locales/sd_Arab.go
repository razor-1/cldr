// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_sd_Arab() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "sd_Arab",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "جنوري", Feb: "فيبروري", Mar: "مارچ", Apr: "اپريل", May: "مئي", Jun: "جون", Jul: "جولاءِ", Aug: "آگسٽ", Sep: "سيپٽمبر", Oct: "آڪٽوبر", Nov: "نومبر", Dec: "ڊسمبر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "آچر", Mon: "سو", Tue: "اڱارو", Wed: "اربع", Thu: "خم", Fri: "جمعو", Sat: "ڇنڇر"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "آچر", Mon: "سومر", Tue: "اڱارو", Wed: "اربع", Thu: "خميس", Fri: "جمعو", Sat: "ڇنڇر"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "صبح،\u202fمنجهند", PM: "منجهند،\u202fشام"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "صبح،\u202fمنجهند", PM: "منجهند،\u202fشام"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "صبح، منجهند", PM: "منجهند، شام"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "گڏيل عرب امارات درهم", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "افغاني افغاني", Symbol: "؋"},
				currency.ALL: cldr.Currency{DisplayName: "الباني ليڪ", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "ارماني ڊرم", Symbol: "֏"},
				currency.ANG: cldr.Currency{DisplayName: "نيڌرلينڊ انٽليئن گلڊر", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "انگوليائي ڪوانزا", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "ارجنٽائن پيسو", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "آسٽريلوي ڊالر", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "اروبن فلورن", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "آذربائيجاني منت", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "بوسنيا هرزگوينا ڪنورٽبل مارڪ", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "باربيڊين ڊالر", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "بنگلاديشي ٽڪا", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "بلغارین لیو", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "بحريني دينار", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "برونڊي فرينڪ", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "برمودي ڊالر", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "برونائي ڊالر", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "بولیوین بولیویانو", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "برازيلي ريل", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "بهاماني ڊادلر", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "ڀوٽاني گلٽرم", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "بوستواني پولا", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "بیلاروسی ربل", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "بيليز ڊالر", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "ڪئينڊيائي ڊالر", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "ڪانگو فرينڪ", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "سوئس فرينڪ", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "چلي پيسو", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "چيني يوآن (غير ملڪي)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "چيني يوآن", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "ڪولمبيائي پيسو", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "ڪوسٽا ريڪا ڪولن", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "ڪيوبن ڪنورٽيبل پيسو", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "ڪيوبن پيسو", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "ڪيپ وردي ايسڪوڊو", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "چيڪ ڪرونا", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "جبوتي فرينڪ", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "دانش ڪرون", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "ڊومينيڪن پيسو", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "الجيريائي دينار", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "مصري پائونڊ", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "ايريٽيريائي ناڪفا", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ايٿوپيائي بر", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "يورو", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "فجي ڊالر", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "فاڪلينڊ ٻيٽ پائونڊ", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "برطانوي پائونڊ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "جارجيائي لاري", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "گهانين سيدي", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "جبرالٽر پائونڊ", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "گيمبيا دلاسائي", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "گني فرينڪ", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "گوئٽي مالا ڪٽزل", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "گيانا ڊالر", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "هانگ ڪانگ ڊالر", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "هونڊوراس ليمپرا", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "ڪروشيائي ڪونا", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "هيٽي گورڊي", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "هنگيرين فورنٽ", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "انڊونيشيائي رپيه", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "اسرائيلي نيو شيڪل", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "انڊين رپي", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "عراقي دينار", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "ايراني ريال", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "آئيس لينڊي ڪرونا", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "جميڪائي ڊالر", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "اردني دينار", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "جاپاني يين", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ڪينيائي سلنگ", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "ڪرغزستاني سوم", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "ڪمبوڊيائي ريال", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "ڪوموريائي فرينڪ", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "اتر ڪوريائي ون", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "ڏکڻ ڪوريائي ون", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "ڪويتي دينار", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "ڪيمين ٻيٽ ڊالر", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "قازقستان ٽينگا", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "لائوشيائي ڪپ", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "لبناني پائونڊ", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "سري لنڪن رپي", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "لائبیریائی ڊالر", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "ليسوٿو لوٽي", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "لبيائي دينار", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "موروڪيائي درهم", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "مالديپ ليو", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "ملاگاسي اریاری", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "ميسي ڊوني دينار", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "ميانمار ڪياٽ", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "منگولي تجرڪ", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "ميڪانيز پٽاڪا", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "موريشيائي اوگوئیا (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "موريشيائي اوگوئیا", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "ماريشيائي رپي", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "مالديپ روفيا", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "ملاوي ڪواچا", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "ميڪسيڪو پيسو", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "ملائيشيائي رنگٽ", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "موزمبيق ميٽيڪل", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "نميبائي ڊالر", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "نائجريائي نائرا", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "نڪارا گوا ڪارڊوبا", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "نارويائي ڪرون", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "نيپالي رپي", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "نيوزي لينڊي ڊالر", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "عماني ريال", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "پاناما پالبوا", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "پيرو سول", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "پاپوا نيو گني ڪنا", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "فلپائني پيسو", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "پاڪستاني رپي", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "پولش زلاٽي", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "پيراگوئي گاراني", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "قطري ريال", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "رومانیائي لیو", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "سربيا دينار", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "روسي ربل", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "روانڊا فرينڪ", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "سعودي ريال", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "سولومان ٻيٽ ڊالر", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "سشلي رپي", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "سوڊاني پائونڊ", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "سويڊني ڪرونا", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "سنگاپوري ڊالر", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "سينٽ هيلنا پائونڊ", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "سیرا لیونيائي لیون", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "سیرا لیونيائي لیون (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "سومالي شلنگ", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "سرينامي ڊالر", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "ڏکڻ سوڊاني پائونڊ", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "سائو ٽوم ۽ پرنسپي دوبرا (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "سائو ٽوم ۽ پرنسپي دوبرا", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "سيريائي پائونڊ", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "سوازي للانگيني", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "ٿائي باهت", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "تاجڪستاني سوموني", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "ترڪمانستان منت", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "تیونس دینار", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "تونگن پانگا", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "ترڪي لرا", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "ٽرينڊيڊ ۽ ٽوباگو ڊالر", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "نيو تائيوان ڊالر", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "تنزانيائي شلنگ", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "یوڪرائن هریونیا", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "يگانڊا شلنگ", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "آمريڪي ڊالر", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "يوروگوئي پيسو", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "ازبڪستاني سوم", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Venezuelan Bolívar (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "وینزویلا بولیور", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "ويٽنامي ڊونگ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "وانواتو واتو", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "ساموآن ٽالا", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "وچ آفريڪا فرينڪ", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "اوڀر ڪيريبين ڊالر", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "اولهه آفريڪا فرينڪ", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP فرينڪ", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "اڻڄاتل سڪو", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "يمني ريال", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "ڏکڻ آفريقي رانڊ", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "زمبائي ڪواچا", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "افار",
			language.AB:      "ابقازیان",
			language.ACE:     "اچائينيز",
			language.ADA:     "ادنگمي",
			language.ADY:     "اديگهي",
			language.AF:      "آفريڪي",
			language.AGQ:     "اگهيم",
			language.AIN:     "آئينو",
			language.AK:      "اڪان",
			language.ALE:     "اليوٽ",
			language.ALT:     "ڏکڻ التائي",
			language.AM:      "امهاري",
			language.AN:      "ارگني",
			language.ANN:     "اوبولو",
			language.ANP:     "انجيڪا",
			language.AR:      "عربي",
			language.AR_001:  "جديد معياري عربي",
			language.ARN:     "ماپوچي",
			language.ARP:     "اراپائو",
			language.ARS:     "نجدي عربي",
			language.AS:      "آسامي",
			language.ASA:     "اسو",
			language.AST:     "اسٽورين",
			language.ATJ:     "اٽيڪاميڪو",
			language.AV:      "اويرس",
			language.AWA:     "اواڌي",
			language.AY:      "ایمارا",
			language.AZ:      "ازري",
			language.BA:      "ڪينيڊا",
			language.BAN:     "بالينيس",
			language.BAS:     "باسا",
			language.BE:      "بيلاروسي",
			language.BEM:     "بيمبا",
			language.BEZ:     "بينا",
			language.BG:      "بلغاريائي",
			language.BGC:     "ھريانوي",
			language.BHO:     "ڀوجپوري",
			language.BI:      "بسلاما",
			language.BIN:     "بني",
			language.BLA:     "سڪسڪا",
			language.BLO:     "آنيائي",
			language.BM:      "بمبارا",
			language.BN:      "بنگلا",
			language.BO:      "تبيتائي",
			language.BR:      "بريٽن",
			language.BRX:     "بودو",
			language.BS:      "بوسنيائي",
			language.BUG:     "بگنيز",
			language.BYN:     "بلن",
			language.CA:      "ڪيٽالان",
			language.CAY:     "ڪايوگا",
			language.CCP:     "چمڪا",
			language.CE:      "چیچن",
			language.CEB:     "سبوانو",
			language.CGG:     "چگا",
			language.CH:      "چمورو",
			language.CHK:     "چڪيز",
			language.CHM:     "ماري",
			language.CHO:     "چوڪ تو",
			language.CHP:     "چائپائن",
			language.CHR:     "چروڪي",
			language.CHY:     "چايان",
			language.CKB:     "مرڪزي ڪردش",
			language.CLC:     "چلڪوٽن",
			language.CO:      "ڪارسيڪائي",
			language.CRG:     "ميچيف",
			language.CRJ:     "ڏکڻ اڀرندو ڪري",
			language.CRK:     "پلينز ڪري",
			language.CRL:     "اترين اوڀر ڪري",
			language.CRM:     "موس ڪري",
			language.CRR:     "ڪيرولينا الگانڪويئن",
			language.CRS:     "سيسلوا ڪريئول فرانسي",
			language.CS:      "چيڪ",
			language.CSW:     "سوامپي ڪري",
			language.CU:      "چرچ سلاوی",
			language.CV:      "چو واش",
			language.CY:      "ويلش",
			language.DA:      "ڊينش",
			language.DAK:     "ڊڪوٽا",
			language.DAR:     "ڊارگوا",
			language.DAV:     "تائيتا",
			language.DE:      "جرمن",
			language.DE_AT:   "آسٽريائي جرمن",
			language.DE_CH:   "سوئس هائي جرمن",
			language.DGR:     "داگرب",
			language.DJE:     "زارما",
			language.DOI:     "ڊوگري",
			language.DSB:     "لوئر سوربين",
			language.DUA:     "ڊيولا",
			language.DV:      "دويهي",
			language.DYO:     "جولا فوني",
			language.DZ:      "زونخا",
			language.DZG:     "دزاگا",
			language.EBU:     "ايمبيو",
			language.EE:      "ايو",
			language.EFI:     "ايفڪ",
			language.EKA:     "ايڪاجڪ",
			language.EL:      "يوناني",
			language.EN:      "انگريزي",
			language.EN_AU:   "آسٽريليائي انگريزي",
			language.EN_CA:   "ڪينيڊيائي انگريزي",
			language.EN_GB:   "برطانوي انگريزي",
			language.EN_US:   "انگريزي (آمريڪا)",
			language.EO:      "ايسپرانٽو",
			language.ES:      "هسپانوي",
			language.ES_419:  "لاطيني آمريڪي اسپينش",
			language.ES_ES:   "يورپي اسپيني",
			language.ES_MX:   "ميڪسيڪين اسپيني",
			language.ET:      "ايستونائي",
			language.EU:      "باسق",
			language.EWO:     "اوانڊو",
			language.FA:      "فارسي",
			language.FA_AF:   "دري",
			language.FF:      "فلاهه",
			language.FI:      "فنش",
			language.FIL:     "فلپائني",
			language.FJ:      "فجي",
			language.FO:      "فيروايس",
			language.FON:     "فون",
			language.FR:      "فرانسيسي",
			language.FR_CA:   "ڪينيڊيائي فرانسيسي",
			language.FR_CH:   "سوئس فرانسيسي",
			language.FRC:     "ڪيجن فرانسيسي",
			language.FRR:     "اترين فريسين",
			language.FUR:     "فرائي لئين",
			language.FY:      "مغربي فريشن",
			language.GA:      "آئرش",
			language.GAA:     "گا",
			language.GD:      "اسڪاٽش گيلڪ",
			language.GEZ:     "جيز",
			language.GIL:     "گلبرٽيز",
			language.GL:      "گليشئين",
			language.GN:      "گواراني",
			language.GOR:     "گورنٽلو",
			language.GSW:     "سوئس جرمن",
			language.GU:      "گجراتي",
			language.GUZ:     "گشي",
			language.GV:      "مينڪس",
			language.GWI:     "گوچن",
			language.HA:      "هوسا",
			language.HAI:     "ھائيڊا",
			language.HAW:     "هوائي",
			language.HAX:     "ڏاکڻي ھائڊا",
			language.HE:      "عبراني",
			language.HI:      "هندي",
			language.HIL:     "هلي گيانان",
			language.HMN:     "مونگ",
			language.HR:      "ڪروشيائي",
			language.HSB:     "اپر سربيائي",
			language.HT:      "هيٽي ڪرولي",
			language.HU:      "هنگري",
			language.HUP:     "هوپا",
			language.HUR:     "ھاڪملم",
			language.HY:      "ارماني",
			language.HZ:      "هريرو",
			language.IA:      "انٽرلنگئا",
			language.IBA:     "ايبن",
			language.IBB:     "ابيبيو",
			language.ID:      "انڊونيشي",
			language.IE:      "انٽرلنگئي",
			language.IG:      "اگبو",
			language.II:      "سچوان يي",
			language.IKT:     "مغربي ڪينيڊين انوڪٽيٽ",
			language.ILO:     "الوڪو",
			language.INH:     "انگش",
			language.IO:      "ادو",
			language.IS:      "آئيس لينڊڪ",
			language.IT:      "اطالوي",
			language.IU:      "انو ڪتوت",
			language.JA:      "جاپاني",
			language.JBO:     "لوجبين",
			language.JGO:     "نغومبا",
			language.JMC:     "ميڪم",
			language.JV:      "جاونيز",
			language.KA:      "جارجيائي",
			language.KAB:     "ڪبائل",
			language.KAC:     "ڪچن",
			language.KAJ:     "پوڪيپسي",
			language.KAM:     "ڪئمبا",
			language.KBD:     "ڪبارڊيئن",
			language.KCG:     "تياپ",
			language.KDE:     "مڪوندي",
			language.KEA:     "ڪيبيو ويرڊيانو",
			language.KFO:     "ڪورو",
			language.KGP:     "ڪئينگينگ",
			language.KHA:     "خاسي",
			language.KHQ:     "ڪيورا چني",
			language.KI:      "اڪويو",
			language.KJ:      "ڪنياما",
			language.KK:      "قازق",
			language.KKJ:     "ڪڪو",
			language.KL:      "ڪالا ليسٽ",
			language.KLN:     "ڪيلين جن",
			language.KM:      "خمر",
			language.KMB:     "ڪنمبونڊو",
			language.KN:      "ڪناڊا",
			language.KO:      "ڪوريائي",
			language.KOK:     "ڪونڪي",
			language.KPE:     "ڪپيل",
			language.KR:      "ڪنوري",
			language.KRC:     "ڪراچي بالڪر",
			language.KRL:     "ڪريلئين",
			language.KRU:     "ڪورخ",
			language.KS:      "ڪشميري",
			language.KSB:     "شمبالا",
			language.KSF:     "بافيا",
			language.KSH:     "ڪلونئين",
			language.KU:      "ڪردي",
			language.KUM:     "ڪومڪ",
			language.KV:      "ڪومي",
			language.KW:      "ڪورنش",
			language.KWK:     "ڪئاڪ ولا",
			language.KXV:     "ڪووي",
			language.KY:      "ڪرغيز",
			language.LA:      "لاطيني",
			language.LAD:     "لڊينو",
			language.LAG:     "لانگي",
			language.LB:      "لگزمبرگ",
			language.LEZ:     "ليزگهين",
			language.LG:      "گاندا",
			language.LI:      "لمبرگش",
			language.LIJ:     "لگيوريئن",
			language.LIL:     "ليلوئيٽ",
			language.LKT:     "لڪوٽا",
			language.LMO:     "لامبارڊ",
			language.LN:      "لنگالا",
			language.LO:      "لائو",
			language.LOU:     "لوئيزيانا ڪريئول",
			language.LOZ:     "لوزي",
			language.LRC:     "اتر لوري",
			language.LSM:     "ساميا",
			language.LT:      "ليٿونيائي",
			language.LU:      "لوبا-ڪتانگا",
			language.LUA:     "لوبا-لولوا",
			language.LUN:     "لنڊا",
			language.LUO:     "لو",
			language.LUS:     "ميزو",
			language.LUY:     "لوهيا",
			language.LV:      "لاتوين",
			language.MAD:     "مدورائي",
			language.MAG:     "مگاهي",
			language.MAI:     "ميٿلي",
			language.MAK:     "مڪاسر",
			language.MAS:     "مسائي",
			language.MDF:     "موڪشا",
			language.MEN:     "مينڊي",
			language.MER:     "ميرو",
			language.MFE:     "موریسیین",
			language.MG:      "ملاگاسي",
			language.MGH:     "مخووا ميتو",
			language.MGO:     "ميتا",
			language.MH:      "مارشليز",
			language.MI:      "مائوري",
			language.MIC:     "ميڪ مڪ",
			language.MIN:     "مناڪابوا",
			language.MK:      "ميسي ڊونيائي",
			language.ML:      "مليالم",
			language.MN:      "منگولي",
			language.MNI:     "ماني پوري",
			language.MOE:     "انو آئيمن",
			language.MOH:     "موهاڪ",
			language.MOS:     "موسي",
			language.MR:      "مراٺي",
			language.MS:      "ملي",
			language.MT:      "مالٽي",
			language.MUA:     "من دانگ",
			language.MUL:     "هڪ کان وڌيڪ ٻوليون",
			language.MUS:     "ڪريڪ",
			language.MWL:     "مرانڊيز",
			language.MY:      "برمي",
			language.MYV:     "ايريزيا",
			language.MZN:     "مزيندراني",
			language.NA:      "نائو",
			language.NAP:     "نيپولٽن",
			language.NAQ:     "ناما",
			language.NB:      "نارويائي بوڪمال",
			language.ND:      "اتر دبيلي",
			language.NDS:     "لو جرمن",
			language.NE:      "نيپالي",
			language.NEW:     "نيواري",
			language.NG:      "ڊونگا",
			language.NIA:     "نياس",
			language.NIU:     "نووي",
			language.NL:      "ڊچ",
			language.NL_BE:   "فليمش",
			language.NMG:     "ڪويسيو",
			language.NN:      "نارويائي نيوناسڪ",
			language.NNH:     "نغيمبون",
			language.NO:      "نارويجيائي",
			language.NOG:     "نوگائي",
			language.NQO:     "نڪو",
			language.NR:      "ڏکڻ دبيلي",
			language.NSO:     "اتر سوٿو",
			language.NUS:     "نيور",
			language.NV:      "نواجو",
			language.NY:      "نيانجا",
			language.NYN:     "نايانڪول",
			language.OC:      "آڪسيٽن",
			language.OJB:     "اتر الھندي اوجيبوا",
			language.OJC:     "وچولي اوجيبوي",
			language.OJS:     "اوجي ڪري",
			language.OJW:     "مغربي اوجيبو",
			language.OKA:     "اوڪاناگن",
			language.OM:      "اورومو",
			language.OR:      "اوڊيا",
			language.OS:      "اوسيٽڪ",
			language.PA:      "پنجابي",
			language.PAG:     "پانگا سينان",
			language.PAM:     "پيم پينگا",
			language.PAP:     "پاپي امينٽو",
			language.PAU:     "پلون",
			language.PCM:     "نائيجرين پجن",
			language.PIS:     "پائجن",
			language.PL:      "پولش",
			language.PQM:     "ماليسيٽ پاسماڪئوڊي",
			language.PRG:     "پرشن",
			language.PS:      "پشتو",
			language.PT:      "پورٽگليز",
			language.PT_BR:   "برازيلي پرتگالي",
			language.PT_PT:   "يورپي پرتگالي",
			language.QU:      "ڪيچوا",
			language.QUC:     "ڪچي",
			language.RAJ:     "راجستاني",
			language.RAP:     "ريپنوئي",
			language.RAR:     "ريرو ٽينگو",
			language.RHG:     "روھنگيا",
			language.RM:      "رومانش",
			language.RN:      "رونڊي",
			language.RO:      "روماني",
			language.RO_MD:   "مالديوي",
			language.ROF:     "رومبو",
			language.RU:      "روسي",
			language.RUP:     "ارومينين",
			language.RW:      "ڪنيار وانڊا",
			language.RWK:     "روا",
			language.SA:      "سنسڪرت",
			language.SAD:     "سنداوي",
			language.SAH:     "ساخا",
			language.SAQ:     "سيمبورو",
			language.SAT:     "سنتالي",
			language.SBA:     "نغمبي",
			language.SBP:     "سانگوو",
			language.SC:      "سارڊيني",
			language.SCN:     "سسلي",
			language.SCO:     "اسڪاٽس",
			language.SD:      "سنڌي",
			language.SE:      "اتر سامي",
			language.SEH:     "سينا",
			language.SES:     "ڪيورابورو سيني",
			language.SG:      "سانگو",
			language.SHI:     "تيچل هاتي",
			language.SHN:     "شان",
			language.SI:      "سنهالا",
			language.SK:      "سلواڪي",
			language.SL:      "سلوويني",
			language.SLH:     "ڏاکڻي لشوٽسيڊ",
			language.SM:      "سموئا",
			language.SMA:     "ڏکڻ سامي",
			language.SMJ:     "لولي سامي",
			language.SMN:     "اناري سامي",
			language.SMS:     "اسڪاٽ سامي",
			language.SN:      "شونا",
			language.SNK:     "سونينڪي",
			language.SO:      "سومالي",
			language.SQ:      "الباني",
			language.SR:      "سربيائي",
			language.SRN:     "سرانن تانگو",
			language.SS:      "سواتي",
			language.SSY:     "سهو",
			language.ST:      "ڏکڻ سوٿي",
			language.STR:     "اسٽريٽ سليش",
			language.SU:      "سوڊاني",
			language.SUK:     "سڪوما",
			language.SV:      "سويڊش",
			language.SW:      "سواحيلي",
			language.SW_CD:   "ڪونگو سواحيلي",
			language.SWB:     "ڪمورين",
			language.SYR:     "شامي",
			language.SZL:     "سليسيئن",
			language.TA:      "تامل",
			language.TCE:     "ڏاکڻي ٽچون",
			language.TE:      "تلگو",
			language.TEM:     "تمني",
			language.TEO:     "تيسو",
			language.TET:     "تيتم",
			language.TG:      "تاجڪ",
			language.TGX:     "ٽئگِش",
			language.TH:      "ٿائي",
			language.THT:     "ٽهلٽن",
			language.TI:      "تگرينيائي",
			language.TIG:     "تگري",
			language.TK:      "ترڪمين",
			language.TLH:     "ڪلون",
			language.TLI:     "ٽِلنگٽ",
			language.TN:      "تسوانا",
			language.TO:      "تونگن",
			language.TOK:     "توڪي پونا",
			language.TPI:     "تاڪ پسن",
			language.TR:      "ترڪي",
			language.TRV:     "تاروڪو",
			language.TS:      "سونگا",
			language.TT:      "تاتار",
			language.TTM:     "اترين ٽچون",
			language.TUM:     "تمبوڪا",
			language.TVL:     "توالو",
			language.TWQ:     "تساوڪي",
			language.TY:      "تاهيتي",
			language.TYV:     "تووينيائي",
			language.TZM:     "وچ اٽلس تمازائيٽ",
			language.UDM:     "ادمرت",
			language.UG:      "يوغور",
			language.UK:      "يوڪراني",
			language.UMB:     "اومبنڊو",
			language.UND:     "اڻڄاتل ٻولي",
			language.UR:      "اردو",
			language.UZ:      "ازبڪ",
			language.VAI:     "يا",
			language.VE:      "وينڊا",
			language.VEC:     "ونيشن",
			language.VI:      "ويتنامي",
			language.VMW:     "مکووا",
			language.VO:      "والپڪ",
			language.VUN:     "ونجو",
			language.WA:      "ولون",
			language.WAE:     "والسر",
			language.WAL:     "وولايٽا",
			language.WAR:     "واري",
			language.WO:      "وولوف",
			language.WUU:     "وو چيني",
			language.XAL:     "ڪيلمڪ",
			language.XH:      "زھوسا",
			language.XNR:     "ڪينگري",
			language.XOG:     "سوگا",
			language.YAV:     "يانگ بين",
			language.YBB:     "ييمبا",
			language.YI:      "يدش",
			language.YO:      "يوروبا",
			language.YRL:     "نھين گاٽو",
			language.YUE:     "چيني، ڪينٽونيز",
			language.ZA:      "جوئنگ",
			language.ZGH:     "معياري مراڪشي تامازائيٽ",
			language.ZH:      "چيني، مندارن",
			language.ZH_HANS: "سادي مندارن چيني",
			language.ZH_HANT: "رواجي مندارن چيني",
			language.ZU:      "زولو",
			language.ZUN:     "زوني",
			language.ZXX:     "ڪوئي ٻولي جو مواد ڪونهي",
			language.ZZA:     "زازا",
		},
		Territories: cldr.Territories{
			territory.V_001: "دنيا",
			territory.V_002: "آفريڪا",
			territory.V_003: "اتر آمريڪا",
			territory.V_005: "ڏکڻ آمريڪا",
			territory.V_009: "اوشنيا",
			territory.V_011: "اولهه آفريقا",
			territory.V_013: "وچ آمريڪا",
			territory.V_014: "اوڀر آفريڪا",
			territory.V_015: "اتر آفريڪا",
			territory.V_017: "وچ آفريڪا",
			territory.V_018: "ڏاکڻي آمريڪا",
			territory.V_019: "آمريڪا",
			territory.V_021: "اترين آمريڪا",
			territory.V_029: "ڪيريبين",
			territory.V_030: "اوڀر ايشيا",
			territory.V_034: "ڏکڻ ايشيا",
			territory.V_035: "ڏکڻ اوڀر ايشيا",
			territory.V_039: "ڏکڻ يورپ",
			territory.V_053: "آسٽریلیشیا",
			territory.V_054: "میلانیشیا",
			territory.V_057: "مائڪرونيشائي خطو",
			territory.V_061: "پولینیشیا",
			territory.V_142: "ايشيا",
			territory.V_143: "وچ ايشيا",
			territory.V_145: "اولهه ايشيا",
			territory.V_150: "يورپ",
			territory.V_151: "اوڀر يورپ",
			territory.V_154: "اترين يورپ",
			territory.V_155: "اولهه يورپ",
			territory.V_202: "سب-سهارا آفريڪا",
			territory.V_419: "لاطيني آمريڪا",
			territory.AC:    "طلوع ٻيٽ",
			territory.AD:    "اندورا",
			territory.AE:    "متحده عرب امارات",
			territory.AF:    "افغانستان",
			territory.AG:    "انٽيگا ۽ باربد",
			territory.AI:    "انگويلا",
			territory.AL:    "البانيا",
			territory.AM:    "ارمینیا",
			territory.AO:    "انگولا",
			territory.AQ:    "انٽارڪٽيڪا",
			territory.AR:    "ارجنٽينا",
			territory.AS:    "آمريڪي ساموا",
			territory.AT:    "آسٽريا",
			territory.AU:    "آسٽريليا",
			territory.AW:    "عروبا",
			territory.AX:    "الند ٻيٽ",
			territory.AZ:    "آذربائيجان",
			territory.BA:    "بوسنيا ۽ هرزوگووينا",
			territory.BB:    "باربڊوس",
			territory.BD:    "بنگلاديش",
			territory.BE:    "بيلجيم",
			territory.BF:    "برڪينا فاسو",
			territory.BG:    "بلغاريا",
			territory.BH:    "بحرين",
			territory.BI:    "برونڊي",
			territory.BJ:    "بينن",
			territory.BL:    "سینٽ برٿلیمی",
			territory.BM:    "برمودا",
			territory.BN:    "برونائي",
			territory.BO:    "بوليويا",
			territory.BQ:    "ڪيريبين نيدرلينڊ",
			territory.BR:    "برازيل",
			territory.BS:    "باهاماس",
			territory.BT:    "ڀوٽان",
			territory.BV:    "بووٽ ٻيٽ",
			territory.BW:    "بوٽسوانا",
			territory.BY:    "بیلارس",
			territory.BZ:    "بيليز",
			territory.CA:    "ڪينيڊا",
			territory.CC:    "ڪوڪوس ٻيٽ",
			territory.CD:    "ڪانگو -ڪنشاسا",
			territory.CF:    "وچ آفريقي جمهوريه",
			territory.CG:    "ڪانگو - برازاویل",
			territory.CH:    "سوئزرلينڊ",
			territory.CI:    "ڪوٽي ويرا",
			territory.CK:    "ڪوڪ ٻيٽ",
			territory.CL:    "چلي",
			territory.CM:    "ڪيمرون",
			territory.CN:    "چين",
			territory.CO:    "ڪولمبيا",
			territory.CP:    "ڪلپرٽن ٻيٽ",
			territory.CR:    "ڪوسٽا ريڪا",
			territory.CU:    "ڪيوبا",
			territory.CV:    "ڪيپ وردي",
			territory.CW:    "ڪيوراسائو",
			territory.CX:    "ڪرسمس ٻيٽ",
			territory.CY:    "سائپرس",
			territory.CZ:    "چيڪيا",
			territory.DE:    "جرمني",
			territory.DG:    "ڊئيگو گارسيا",
			territory.DJ:    "ڊجبيوتي",
			territory.DK:    "ڊينمارڪ",
			territory.DM:    "ڊومينيڪا",
			territory.DO:    "ڊومينيڪن جمهوريه",
			territory.DZ:    "الجيريا",
			territory.EA:    "سیوٽا ۽ میلیلا",
			territory.EC:    "ايڪواڊور",
			territory.EE:    "ايسٽونيا",
			territory.EG:    "مصر",
			territory.EH:    "اولهه صحارا",
			territory.ER:    "ايريٽيريا",
			territory.ES:    "اسپين",
			territory.ET:    "ايٿوپيا",
			territory.EU:    "يورپين يونين",
			territory.EZ:    "يورو زون",
			territory.FI:    "فن لينڊ",
			territory.FJ:    "فجي",
			territory.FK:    "فاڪ لينڊ ٻيٽ",
			territory.FM:    "مائڪرونيشيا",
			territory.FO:    "فارو ٻيٽ",
			territory.FR:    "فرانس",
			territory.GA:    "گبون",
			territory.GB:    "برطانيہ",
			territory.GD:    "گريناڊا",
			territory.GE:    "جارجيا",
			territory.GF:    "فرانسيسي گيانا",
			territory.GG:    "گورنسي",
			territory.GH:    "گهانا",
			territory.GI:    "جبرالٽر",
			territory.GL:    "گرين لينڊ",
			territory.GM:    "گيمبيا",
			territory.GN:    "گني",
			territory.GP:    "گواڊیلوپ",
			territory.GQ:    "ايڪوٽوريل گائينا",
			territory.GR:    "يونان",
			territory.GS:    "ڏکڻ جارجيا ۽ ڏکڻ سينڊوچ ٻيٽ",
			territory.GT:    "گوئٽي مالا",
			territory.GU:    "گوام",
			territory.GW:    "گني بسائو",
			territory.GY:    "گيانا",
			territory.HK:    "هانگ ڪانگ SAR",
			territory.HM:    "هرڊ ۽ مڪڊونلڊ ٻيٽ",
			territory.HN:    "هنڊورس",
			territory.HR:    "ڪروئيشيا",
			territory.HT:    "هيٽي",
			territory.HU:    "هنگري",
			territory.IC:    "ڪينري ٻيٽ",
			territory.ID:    "انڊونيشيا",
			territory.IE:    "آئرلينڊ",
			territory.IL:    "اسرائيل",
			territory.IM:    "انسانن جو ٻيٽ",
			territory.IN:    "ڀارت",
			territory.IO:    "برطانوي هندي سمنڊ خطو",
			territory.IQ:    "عراق",
			territory.IR:    "ايران",
			territory.IS:    "آئس لينڊ",
			territory.IT:    "اٽلي",
			territory.JE:    "جرسي",
			territory.JM:    "جميڪا",
			territory.JO:    "اردن",
			territory.JP:    "جاپان",
			territory.KE:    "ڪينيا",
			territory.KG:    "ڪرغستان",
			territory.KH:    "ڪمبوڊيا",
			territory.KI:    "ڪرباتي",
			territory.KM:    "ڪوموروس",
			territory.KN:    "سينٽ ڪٽس و نيوس",
			territory.KP:    "اتر ڪوريا",
			territory.KR:    "ڏکڻ ڪوريا",
			territory.KW:    "ڪويت",
			territory.KY:    "ڪي مين ٻيٽ",
			territory.KZ:    "قازقستان",
			territory.LA:    "لائوس",
			territory.LB:    "لبنان",
			territory.LC:    "سينٽ لوسيا",
			territory.LI:    "لچي ٽينسٽين",
			territory.LK:    "سري لنڪا",
			territory.LR:    "لائبیریا",
			territory.LS:    "ليسوٿو",
			territory.LT:    "لٿونيا",
			territory.LU:    "لگزمبرگ",
			territory.LV:    "لاتويا",
			territory.LY:    "لبيا",
			territory.MA:    "مراڪش",
			territory.MC:    "موناڪو",
			territory.MD:    "مالدووا",
			territory.ME:    "مونٽي نيگرو",
			territory.MF:    "سينٽ مارٽن",
			territory.MG:    "مدگاسڪر",
			territory.MH:    "مارشل ٻيٽ",
			territory.MK:    "اتر مقدونيا",
			territory.ML:    "مالي",
			territory.MM:    "ميانمار (برما)",
			territory.MN:    "منگوليا",
			territory.MO:    "مڪائو SAR چين",
			territory.MP:    "اتريان ماريانا ٻيٽ",
			territory.MQ:    "مارتينڪ",
			territory.MR:    "موريتانيا",
			territory.MS:    "مونٽسراٽ",
			territory.MT:    "مالٽا",
			territory.MU:    "موريشس",
			territory.MV:    "مالديپ",
			territory.MW:    "مالاوي",
			territory.MX:    "ميڪسيڪو",
			territory.MY:    "ملائيشيا",
			territory.MZ:    "موزمبیق",
			territory.NA:    "نيميبيا",
			territory.NC:    "نیو ڪالیڊونیا",
			territory.NE:    "نائيجر",
			territory.NF:    "نورفوڪ ٻيٽ",
			territory.NG:    "نائيجيريا",
			territory.NI:    "نڪراگوا",
			territory.NL:    "نيدرلينڊ",
			territory.NO:    "ناروي",
			territory.NP:    "نيپال",
			territory.NR:    "نائورو",
			territory.NU:    "نووي",
			territory.NZ:    "نيو زيلينڊ",
			territory.OM:    "عمان",
			territory.PA:    "پناما",
			territory.PE:    "پيرو",
			territory.PF:    "فرانسيسي پولينيشيا",
			territory.PG:    "پاپوا نیو گني",
			territory.PH:    "فلپائن",
			territory.PK:    "پاڪستان",
			territory.PL:    "پولينڊ",
			territory.PM:    "سینٽ پیئر و میڪوئیلون",
			territory.PN:    "پٽڪئرن ٻيٽ",
			territory.PR:    "پيوئرٽو ريڪو",
			territory.PS:    "فلسطيني علائقا",
			territory.PT:    "پرتگال",
			territory.PW:    "پلائو",
			territory.PY:    "پيراگوءِ",
			territory.QA:    "قطر",
			territory.QO:    "بيروني سامونڊي",
			territory.RE:    "ري يونين",
			territory.RO:    "رومانيا",
			territory.RS:    "سربيا",
			territory.RU:    "روس",
			territory.RW:    "روانڊا",
			territory.SA:    "سعودي عرب",
			territory.SB:    "سولومون ٻيٽَ",
			territory.SC:    "شي شلز",
			territory.SD:    "سوڊان",
			territory.SE:    "سوئيڊن",
			territory.SG:    "سنگاپور",
			territory.SH:    "سينٽ ھيلينا",
			territory.SI:    "سلوینیا",
			territory.SJ:    "سوالبارڊ ۽ جان ماین",
			territory.SK:    "سلوواڪيا",
			territory.SL:    "سيرا ليون",
			territory.SM:    "سین مرینو",
			territory.SN:    "سينيگال",
			territory.SO:    "سوماليا",
			territory.SR:    "سورينام",
			territory.SS:    "ڏکڻ سوڊان",
			territory.ST:    "سائو ٽوم ۽ پرنسپیي",
			territory.SV:    "ال سلواڊور",
			territory.SX:    "سنٽ مارٽن",
			territory.SY:    "شام",
			territory.SZ:    "ايسواٽني",
			territory.TA:    "ٽرسٽن دا ڪوها",
			territory.TC:    "ترڪ ۽ ڪيڪوس ٻيٽ",
			territory.TD:    "چاڊ",
			territory.TF:    "فرانسيسي ڏاکڻي علائقا",
			territory.TG:    "ٽوگو",
			territory.TH:    "ٿائيليند",
			territory.TJ:    "تاجڪستان",
			territory.TK:    "ٽوڪلائو",
			territory.TL:    "تيمور ليستي",
			territory.TM:    "ترڪمانستان",
			territory.TN:    "تيونيسيا",
			territory.TO:    "ٽونگا",
			territory.TR:    "ترڪييي",
			territory.TT:    "ٽريني ڊيڊ ۽ ٽوباگو ٻيٽ",
			territory.TV:    "توالو",
			territory.TW:    "تائیوان",
			territory.TZ:    "تنزانيا",
			territory.UA:    "يوڪرين",
			territory.UG:    "يوگنڊا",
			territory.UM:    "آمريڪي خارجي ٻيٽ",
			territory.UN:    "گڏيل قومون",
			territory.US:    "آمريڪا جون گڏيل رياستون",
			territory.UY:    "يوروگوءِ",
			territory.UZ:    "ازبڪستان",
			territory.VA:    "ويٽڪين سٽي",
			territory.VC:    "سینٽ ونسنت ۽ گریناڊینز",
			territory.VE:    "وينزويلا",
			territory.VG:    "برطانوي ورجن ٻيٽ",
			territory.VI:    "آمريڪي ورجن ٻيٽ",
			territory.VN:    "ويتنام",
			territory.VU:    "وينيٽيو",
			territory.WF:    "والس ۽ فتونا",
			territory.WS:    "ساموا",
			territory.XA:    "سوڊو-لهجا",
			territory.XB:    "سوڊو-بي ڊي",
			territory.XK:    "ڪوسووو",
			territory.YE:    "يمن",
			territory.YT:    "مياتي",
			territory.ZA:    "ڏکڻ آفريقا",
			territory.ZM:    "زيمبيا",
			territory.ZW:    "زمبابوي",
			territory.ZZ:    "اڻڄاتل خطو",
		},
	}
}
