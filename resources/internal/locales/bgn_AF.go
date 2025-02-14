// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_bgn_AF() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "bgn_AF",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "،", Negative: "-", Percent: "٪", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "0\u00a0هزار\u00a0¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AFN: cldr.Currency{DisplayName: "اوگانستانئ اوگانی", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "بنگلادیشئ ٹاکه", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "بوتانئ انگولٹروم", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
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
				currency.EUR: cldr.Currency{DisplayName: "", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "هندوستانئ روپی", Symbol: "₹"},
				currency.IRR: cldr.Currency{DisplayName: "ایرانئ ریال", Symbol: "ریال"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "", Symbol: "JP¥"},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "سریلانکایی روپی", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.MGA: cldr.Currency{DisplayName: "", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "مالدیوی روپی", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "نیپالین روپی", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "پاکستانئ روپی", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "روسین روبل", Symbol: "₽"},
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
				currency.USD: cldr.Currency{DisplayName: "", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "افار",
			language.AB:      "آبخازی",
			language.ACE:     "آچه\u200cای",
			language.ADA:     "ادانگمی",
			language.ADY:     "ادیغی",
			language.AF:      "افریکانس",
			language.AGQ:     "اغیمی",
			language.AIN:     "آینو",
			language.AK:      "اکانی",
			language.ALE:     "الیوتی",
			language.ALT:     "جهلسرین آلتایی",
			language.AM:      "امهاری",
			language.AN:      "آراگونی",
			language.ANP:     "انگیکای",
			language.AR:      "عربی",
			language.AR_001:  "پیشرفته\u200cاین عربی",
			language.ARN:     "ماپوچه\u200cای",
			language.ARP:     "اراپاهو",
			language.AS:      "آسامی",
			language.ASA:     "آسو",
			language.AST:     "آستوری",
			language.AV:      "آواری",
			language.AWA:     "آوادی",
			language.AY:      "ایمارای",
			language.AZ:      "آذری",
			language.BA:      "باشقیری",
			language.BAN:     "بالینسی",
			language.BAS:     "باسا",
			language.BE:      "بیلاروسی",
			language.BEM:     "بیمبا",
			language.BEZ:     "بینای",
			language.BG:      "بلغاری",
			language.BGN:     "بلوچی (رخشانی)",
			language.BHO:     "بهوجپوری",
			language.BI:      "بیسلامه",
			language.BIN:     "بینی",
			language.BLA:     "سیکسیکا",
			language.BM:      "بمبارای",
			language.BN:      "بنگالی",
			language.BO:      "تبتی",
			language.BR:      "بریتون",
			language.BRX:     "بۆدو",
			language.BS:      "بوسنی",
			language.BUG:     "بوگینسی",
			language.BYN:     "بلین",
			language.CA:      "کاتالانی",
			language.CE:      "چیچینی",
			language.CEB:     "سینوگبانونی",
			language.CGG:     "شیگی",
			language.CH:      "چه\u200cمروری",
			language.CHK:     "تروسکی",
			language.CHM:     "ماری",
			language.CHO:     "چاکتاوی",
			language.CHR:     "چیروکی",
			language.CHY:     "شاینی زبان",
			language.CKB:     "مرکزین کوردی",
			language.CO:      "کرسی",
			language.CRS:     "سیشلی کریولین فرانسوی",
			language.CS:      "چیکی",
			language.CU:      "سلاواکی کلیسایی",
			language.CV:      "چواشی",
			language.CY:      "ولزی",
			language.DA:      "ڈینمارکی",
			language.DAK:     "داکوتی",
			language.DAR:     "دارگوایی",
			language.DAV:     "تایتایی",
			language.DE:      "جرمنی",
			language.DE_AT:   "استرالیاین جرمنی",
			language.DE_CH:   "سویسین جرمنی",
			language.DGR:     "داگریبی",
			language.DJE:     "زرمی",
			language.DSB:     "صُربی سفلی",
			language.DUA:     "دوالی",
			language.DV:      "دیوهی",
			language.DYO:     "جولا فونی",
			language.DZ:      "دزونگخا",
			language.DZG:     "دزازا",
			language.EBU:     "ایمبو",
			language.EE:      "اوه\u200cای",
			language.EFI:     "ایفیکی",
			language.EKA:     "ایکاجوکی",
			language.EL:      "یونانی",
			language.EN:      "انگریزی",
			language.EN_AU:   "استرالیاین انگریزی",
			language.EN_CA:   "کاناڈاین انگریزی",
			language.EN_GB:   "بریتن انگریزی",
			language.EN_US:   "یو اس انگریزی",
			language.EO:      "اسپرانتوی",
			language.ES:      "هسپانوی",
			language.ES_419:  "لاتین امریکایی هسپانوی",
			language.ES_ES:   "اوروپایین هسپانوی",
			language.ES_MX:   "مکسیکوین هسپانوی",
			language.ET:      "استونیایی",
			language.EU:      "باسکی",
			language.EWO:     "اواندویی",
			language.FA:      "پارسی",
			language.FF:      "فولایی",
			language.FI:      "فنلاندی",
			language.FIL:     "فلیپینی",
			language.FJ:      "فیجی",
			language.FO:      "فاروئی",
			language.FON:     "فون",
			language.FR:      "فرانسوی",
			language.FR_CA:   "کاناڈاین فرانسوی",
			language.FR_CH:   "اشاره\u200cاین فرانسوی",
			language.FUR:     "فریولی",
			language.FY:      "روچ\u200cکپتین فریزی",
			language.GA:      "ایرلندی",
			language.GAA:     "گا",
			language.GD:      "اسکاتلندی گیلی",
			language.GEZ:     "گعزی",
			language.GIL:     "گیلبیرتی",
			language.GL:      "گالیسی",
			language.GN:      "گوارانی",
			language.GOR:     "گورونتالو",
			language.GSW:     "جرمنین سوئیسی",
			language.GU:      "گوجراتی",
			language.GUZ:     "گوسی",
			language.GV:      "مانی",
			language.GWI:     "گویچنی",
			language.HA:      "هوسه\u200cای",
			language.HAW:     "هاوایی",
			language.HE:      "عبرانی",
			language.HI:      "هندی",
			language.HIL:     "هیلیگایونی",
			language.HMN:     "همونگی",
			language.HR:      "کراوتی",
			language.HSB:     "علیای سیربی",
			language.HT:      "کریول آییسینی",
			language.HU:      "مجارستانی",
			language.HUP:     "هوپی",
			language.HY:      "ارمنی",
			language.HZ:      "هرویی",
			language.IA:      "اینترلینگوایی",
			language.IBA:     "ایبانگه",
			language.IBB:     "ایبیبیو",
			language.ID:      "ایندونیزیایی",
			language.IG:      "ایگبویی",
			language.II:      "یی سیچوان",
			language.ILO:     "ایلوکانوی",
			language.INH:     "اینگوشی",
			language.IO:      "ایدوی",
			language.IS:      "ایسلندی",
			language.IT:      "ایتالیایی",
			language.IU:      "اینوکتیتوتی",
			language.JA:      "جاپانی",
			language.JBO:     "لوجبانی",
			language.JGO:     "نگومبی",
			language.JMC:     "ماچامه\u200cای",
			language.JV:      "جاوه\u200cای",
			language.KA:      "گرجی",
			language.KAB:     "قبایلی",
			language.KAC:     "کاچینی",
			language.KAJ:     "جیجو",
			language.KAM:     "کامبایی",
			language.KBD:     "کاباردینی",
			language.KCG:     "تیاپی",
			language.KDE:     "ماکوندی",
			language.KEA:     "کابووردیانو",
			language.KFO:     "کورو",
			language.KHA:     "خاسی",
			language.KHQ:     "کوجراچینی",
			language.KI:      "کیکویویی",
			language.KJ:      "کوانیامایی",
			language.KK:      "قزاقی",
			language.KKJ:     "کاکویی",
			language.KL:      "گرینلندی",
			language.KLN:     "کالنجین",
			language.KM:      "خمری",
			language.KMB:     "کیمبوندویی",
			language.KN:      "کانارا",
			language.KO:      "کوریایی",
			language.KOK:     "کونکانی",
			language.KPE:     "کپله\u200cای",
			language.KR:      "کانوری",
			language.KRC:     "قره\u200cچایی‐بالکاری",
			language.KRL:     "کاریلینی",
			language.KRU:     "کوروخی",
			language.KS:      "کشمیری",
			language.KSB:     "شامبالای",
			language.KSF:     "بافیا",
			language.KSH:     "کولوگنی",
			language.KU:      "کوردی",
			language.KUM:     "کومیکی",
			language.KV:      "کومی",
			language.KW:      "کورنی",
			language.KY:      "قیرغیزی",
			language.LA:      "لاتین",
			language.LAD:     "لادینو",
			language.LAG:     "لانگی",
			language.LB:      "لوگزامبورگی",
			language.LEZ:     "لزگی",
			language.LG:      "گاندایی",
			language.LI:      "لیمبورگی",
			language.LKT:     "لاکوتا",
			language.LN:      "لینگالایی",
			language.LO:      "لائوسی",
			language.LOZ:     "لوزی",
			language.LRC:     "بُرزسرین لوری",
			language.LT:      "لیتوانی",
			language.LU:      "لوبایی‐کاتانگا",
			language.LUA:     "لوبایی‐لولوا",
			language.LUN:     "لوندایی",
			language.LUO:     "لوئویی",
			language.LUS:     "لوشه\u200cای",
			language.LUY:     "لویایی",
			language.LV:      "لاتوینی",
			language.MAD:     "مادورایی",
			language.MAG:     "ماگاهی",
			language.MAI:     "مایدیلی",
			language.MAK:     "ماکاساری",
			language.MAS:     "ماسایی",
			language.MDF:     "موکشی",
			language.MEN:     "منده\u200cای",
			language.MER:     "مرویی",
			language.MFE:     "موریسینی",
			language.MG:      "مالاگاسی",
			language.MGH:     "ماکوا متوی",
			language.MGO:     "میٹایی",
			language.MH:      "مارشالی",
			language.MI:      "مائوری",
			language.MIC:     "میکماکی",
			language.MIN:     "مینانگ\u200cکابویی",
			language.MK:      "مقدونی",
			language.ML:      "مالایالامی",
			language.MN:      "منگولی",
			language.MNI:     "مانیپوری",
			language.MOH:     "موهاکی",
			language.MOS:     "موسیی",
			language.MR:      "مراٹی",
			language.MS:      "مالایی",
			language.MT:      "مالته\u200cای",
			language.MUA:     "ماندانگی",
			language.MUL:     "چینکه زبان",
			language.MUS:     "کریکی",
			language.MWL:     "میراندیسی",
			language.MY:      "بورمي",
			language.MYV:     "ارزیای",
			language.MZN:     "مازندرانی",
			language.NA:      "نائورویی",
			language.NAP:     "ناپلی",
			language.NAQ:     "نامایی",
			language.NB:      "نارویی بوک\u200cمولی",
			language.ND:      "بُرزسرین انده\u200cبله\u200cای",
			language.NDS_NL:  "ساکسونی سفلی",
			language.NE:      "نیپالی",
			language.NEW:     "نیواری",
			language.NG:      "اندونگی",
			language.NIA:     "نیاسی",
			language.NIU:     "نیویی",
			language.NL:      "هالنڈی",
			language.NL_BE:   "فلامانی",
			language.NMG:     "کوازیو",
			language.NN:      "نارویی نی\u200cنوشکی",
			language.NNH:     "انگی\u200cایمبونی",
			language.NOG:     "نغایی",
			language.NQO:     "نکوی",
			language.NR:      "جهلسرین انده\u200cبله\u200cای",
			language.NSO:     "بُرزسرین سوتویی",
			language.NUS:     "نویری",
			language.NV:      "ناواهویی",
			language.NY:      "نیانجی",
			language.NYN:     "نیانکوله\u200cای",
			language.OC:      "اوکیتایی",
			language.OM:      "اورومویی",
			language.OR:      "اودیه\u200cای",
			language.OS:      "آسیتینی",
			language.PA:      "پنجاپی",
			language.PAG:     "پانگاسینانی",
			language.PAM:     "پامپانگی",
			language.PAP:     "پاپیامنتوی",
			language.PAU:     "پالائویی",
			language.PCM:     "نایجیریای پیڈگین",
			language.PL:      "پولنڈی",
			language.PRG:     "پروسی",
			language.PS:      "پشتو",
			language.PT:      "پورتگالی",
			language.PT_BR:   "برازیلین پورتگالی",
			language.PT_PT:   "اوروپایین پورتگالی",
			language.QU:      "کچوایی",
			language.QUC:     "کیچه\u200c",
			language.RAP:     "راپانویی",
			language.RAR:     "راروتونگی",
			language.RM:      "رومانش",
			language.RN:      "رونڈی",
			language.RO:      "رومانی",
			language.RO_MD:   "مالداوی",
			language.ROF:     "رومبویی",
			language.RU:      "اوروسی",
			language.RUP:     "آرومانی",
			language.RW:      "کینیارواندی",
			language.RWK:     "روایی",
			language.SA:      "سانسکریٹ",
			language.SAD:     "سانڈاوه\u200cای",
			language.SAH:     "یاقوتی",
			language.SAQ:     "سامبوروی",
			language.SAT:     "سانٹالی",
			language.SBA:     "انگامبی",
			language.SBP:     "سانگویی",
			language.SC:      "ساردینی",
			language.SCN:     "سیسیلی",
			language.SCO:     "اسکاتلندی",
			language.SD:      "سیندی",
			language.SE:      "بُرزسرین سامی",
			language.SEH:     "سینایی",
			language.SES:     "کویرابورا سنی",
			language.SG:      "سانگۆیی",
			language.SHI:     "تاچل\u200cهیتی",
			language.SHN:     "شانی",
			language.SI:      "سینهالی",
			language.SK:      "اسلواکی",
			language.SL:      "اسلوانی",
			language.SM:      "ساموآیی",
			language.SMA:     "جهلسرین سامی",
			language.SMJ:     "لوله سامی",
			language.SMN:     "اناری سمی",
			language.SMS:     "اسکولت سامی",
			language.SN:      "شونی",
			language.SNK:     "سونینکه\u200cای",
			language.SO:      "سومالی",
			language.SQ:      "البانی",
			language.SR:      "سیربی",
			language.SRN:     "تاکی\u200cتاکی",
			language.SS:      "سواتی",
			language.SSY:     "ساهویی",
			language.ST:      "جهلسرین سوتویی",
			language.SU:      "سونڈی",
			language.SUK:     "سوکومایی",
			language.SV:      "سویڈنی",
			language.SW:      "سواحلی",
			language.SW_CD:   "کانگویی سواحلی",
			language.SWB:     "قمرین",
			language.SYR:     "سریانی",
			language.TA:      "تامیلی",
			language.TE:      "تلوگویی",
			language.TEM:     "تیمنه\u200cای",
			language.TEO:     "تیسویی",
			language.TET:     "ٹیٹومی",
			language.TG:      "تاجیکی",
			language.TH:      "تایلندی",
			language.TI:      "ٹیگرینیایی",
			language.TIG:     "ٹایگری",
			language.TK:      "تورکمنی",
			language.TLH:     "کلینگونی",
			language.TN:      "تسوانی",
			language.TO:      "تونگی",
			language.TPI:     "توک\u200cپیسینی",
			language.TR:      "تورکی",
			language.TRV:     "تاروکویی",
			language.TS:      "تسونگی",
			language.TT:      "تاتاری",
			language.TUM:     "تومبوکی",
			language.TVL:     "ٹووالی",
			language.TWQ:     "تسواکی",
			language.TY:      "تاهیتی",
			language.TYV:     "ٹووینی",
			language.TZM:     "مرکزین اتلسین تامازیگتی",
			language.UDM:     "اوڈمورتی",
			language.UG:      "اویغوری",
			language.UK:      "اوکراینی",
			language.UMB:     "امبونڈویی",
			language.UND:     "نازانتین زبان",
			language.UR:      "اوردو",
			language.UZ:      "اوزبکی",
			language.VAI:     "وایی",
			language.VE:      "وینڈایی",
			language.VI:      "ویتنامی",
			language.VO:      "ولاپوکی",
			language.VUN:     "ونجو",
			language.WA:      "والونی",
			language.WAE:     "والسیری",
			language.WAL:     "وولایٹی",
			language.WAR:     "واری",
			language.WO:      "ولوفی",
			language.XAL:     "کالمیکی",
			language.XH:      "خوسی",
			language.XOG:     "سوگایی",
			language.YAV:     "یانگبینی",
			language.YBB:     "یمبی",
			language.YI:      "یدی",
			language.YO:      "یوروبایی",
			language.YUE:     "کانتونیونی",
			language.ZGH:     "آمازیغی مراکشی معیار",
			language.ZH:      "چینایی",
			language.ZH_HANS: "ساده\u200cگین چینایی",
			language.ZH_HANT: "غدیمین چینایی",
			language.ZU:      "زولویی",
			language.ZUN:     "زونی",
			language.ZXX:     "بغیر شه زبانین لڑا",
			language.ZZA:     "زازایی",
		},
		Territories: cldr.Territories{
			territory.V_001: "دونیا/جهان",
			territory.V_002: "افریقا",
			territory.V_009: "اوقیانوسیه",
			territory.V_019: "امریکای براعظم",
			territory.V_053: "استرالیا",
			territory.V_142: "اسیا",
			territory.V_150: "اورورپا",
			territory.AE:    "متحدین عربین امارات",
			territory.AI:    "انگویلا",
			territory.AO:    "انگولا",
			territory.AR:    "ارجنٹاین",
			territory.AU:    "اسٹرالیا",
			territory.AZ:    "آزربایجان",
			territory.BE:    "بیلجیم",
			territory.BG:    "بولغاریه",
			territory.BH:    "بحرین",
			territory.BI:    "بروندی",
			territory.BM:    "بیرمودا",
			territory.BN:    "برونی",
			territory.BO:    "بولیویه",
			territory.BR:    "برازیل",
			territory.BS:    "بهاماس",
			territory.BT:    "بوتان",
			territory.CA:    "کاناڈا",
			territory.CL:    "چیلی",
			territory.CM:    "کامیرون",
			territory.CN:    "چین",
			territory.CO:    "کولومبیا",
			territory.CU:    "کوبا",
			territory.CY:    "قبرس",
			territory.DE:    "جرمنی",
			territory.DJ:    "جیبوتی",
			territory.DZ:    "الجزایر",
			territory.EC:    "اکوادور",
			territory.EG:    "مصر",
			territory.EH:    "روچ\u200cکپتین سحرا",
			territory.ER:    "اریتره",
			territory.ES:    "هسپانیه",
			territory.ET:    "ایتوپیه",
			territory.EU:    "اورورپایی یکویی",
			territory.FJ:    "فیجی",
			territory.FR:    "فرانسه",
			territory.GA:    "گابون",
			territory.GE:    "گرجستان",
			territory.GH:    "گانا",
			territory.GL:    "گرینلاند",
			territory.GM:    "گامبیا",
			territory.GN:    "گوینیا",
			territory.GR:    "یونان",
			territory.GY:    "گویانا",
			territory.HK:    "هانگ کانگ",
			territory.HU:    "هنگری",
			territory.ID:    "ایندونیزیا",
			territory.IL:    "اسرائیل",
			territory.IQ:    "عراق",
			territory.IT:    "ایتالیه",
			territory.JO:    "اردن",
			territory.KE:    "کینیا",
			territory.KG:    "قیرغیزستان",
			territory.KH:    "کمبودیا",
			territory.KM:    "کومورس",
			territory.KW:    "کویٹ",
			territory.KZ:    "قزاقستان",
			territory.LA:    "لاوس",
			territory.LB:    "لیبنان",
			territory.LY:    "لیبیا",
			territory.MA:    "مراکو",
			territory.MD:    "مالداویا",
			territory.MG:    "ماداگاسکار",
			territory.ML:    "مالی",
			territory.MT:    "مالته",
			territory.MU:    "موریتانیا",
			territory.MX:    "مکسیکو",
			territory.MY:    "مالیزیا",
			territory.NE:    "نیجیر",
			territory.NG:    "نایجیریا",
			territory.NZ:    "نیوزلنڈ",
			territory.OM:    "ئومان",
			territory.PA:    "پانامه",
			territory.PE:    "پیرو",
			territory.PH:    "فلیپین",
			territory.PT:    "پورتگال",
			territory.PY:    "پاراگوی",
			territory.QA:    "قطر",
			territory.RO:    "رومانیه",
			territory.RS:    "سیربستان",
			territory.RW:    "روندا",
			territory.SC:    "سیشیل",
			territory.SD:    "سوڈان",
			territory.SG:    "سینگاپور",
			territory.SN:    "سینیگال",
			territory.SO:    "سومالیا",
			territory.SR:    "سورینامی",
			territory.SY:    "سوریه",
			territory.TD:    "چاد",
			territory.TH:    "ٹایلنڈ",
			territory.TJ:    "تاجیکستان",
			territory.TM:    "تورکمنستان",
			territory.TN:    "ٹونیس",
			territory.TZ:    "تانزانیا",
			territory.UG:    "اوگاندا",
			territory.US:    "متحدین ایالات",
			territory.UY:    "اوراگوی",
			territory.UZ:    "اوزبکیستان",
			territory.VE:    "وینزوویلا",
			territory.VN:    "ویتنام",
			territory.YE:    "یمن",
			territory.ZM:    "زامبیا",
			territory.ZW:    "زیمبابوی",
			territory.ZZ:    "نازانتین سیمسر",
		},
	}
}
