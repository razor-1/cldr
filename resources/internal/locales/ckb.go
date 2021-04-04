// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ckb() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ckb",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "dی MMMMی y", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "کانوونی دووەم", Feb: "شوبات", Mar: "ئازار", Apr: "نیسان", May: "ئایار", Jun: "حوزەیران", Jul: "تەمووز", Aug: "ئاب", Sep: "ئەیلوول", Oct: "تشرینی یەکەم", Nov: "تشرینی دووەم", Dec: "کانونی یەکەم"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ک", Feb: "ش", Mar: "ئ", Apr: "ن", May: "ئ", Jun: "ح", Jul: "ت", Aug: "ئ", Sep: "ئ", Oct: "ت", Nov: "ت", Dec: "ک"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "کانوونی دووەم", Feb: "شوبات", Mar: "ئازار", Apr: "نیسان", May: "ئایار", Jun: "حوزەیران", Jul: "تەمووز", Aug: "ئاب", Sep: "ئەیلوول", Oct: "تشرینی یەکەم", Nov: "تشرینی دووەم", Dec: "کانونی یەکەم"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "یەکشەممە", Mon: "دووشەممە", Tue: "سێشەممە", Wed: "چوارشەممە", Thu: "پێنجشەممە", Fri: "ھەینی", Sat: "شەممە"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ی", Mon: "د", Tue: "س", Wed: "چ", Thu: "پ", Fri: "ھ", Sat: "ش"}, Short: cldr.CalendarDayFormatNameValue{Sun: "١ش", Mon: "٢ش", Tue: "٣ش", Wed: "٤ش", Thu: "٥ش", Fri: "ھ", Sat: "ش"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "یەکشەممە", Mon: "دووشەممە", Tue: "سێشەممە", Wed: "چوارشەممە", Thu: "پێنجشەممە", Fri: "ھەینی", Sat: "شەممە"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ب.ن", PM: "د.ن"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ب.ن", PM: "د.ن"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ب.ن", PM: "د.ن"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u200f-", Percent: "٪", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AFN: cldr.Currency{DisplayName: "ئەفغانیی ئەفغانستان", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "دیناری بەحرەینی", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "دۆلاری بەلیزی", Symbol: "$"},
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
				currency.DZD: cldr.Currency{DisplayName: "دیناری جەزائیری", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "یورۆ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
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
				currency.INR: cldr.Currency{DisplayName: "", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "دیناری عێراقی", Symbol: "د.ع.\u200f"},
				currency.IRR: cldr.Currency{DisplayName: "ڕیاڵی ئێرانی", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "دیناری ئوردنی", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "", Symbol: "JP¥"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "دیناری کووەیتی", Symbol: ""},
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
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ڕیاڵی عومانی", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "ڕیاڵی قەتەری", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "ڕیاڵی سەعوودی", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "دیناری توونس", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "لیرەی تورکیا", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "دۆلاری ترینیداد و تۆباگۆ", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XAU: cldr.Currency{DisplayName: "زێڕ", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "ئەفار",
			language.AB:      "ئەبخازی",
			language.ACE:     "ئاچەیی",
			language.ADA:     "دانگمێ",
			language.ADY:     "ئادیگی",
			language.AF:      "ئەفریکانس",
			language.AGQ:     "ئاگێم",
			language.AIN:     "ئاینوو",
			language.AK:      "ئاکان",
			language.ALE:     "ئالیوت",
			language.ALT:     "ئاڵتایی باشوور",
			language.AM:      "ئەمھەری",
			language.AN:      "ئاراگۆنی",
			language.ANP:     "ئەنگیکا",
			language.AR:      "عەرەبی",
			language.AR_001:  "عەرەبیی ستاندارد",
			language.ARN:     "ماپووچە",
			language.ARP:     "ئاراپاهۆ",
			language.AS:      "ئاسامی",
			language.ASA:     "ئاسوو",
			language.AST:     "ئاستۆری",
			language.AV:      "ئەڤاری",
			language.AWA:     "ئاوادهی",
			language.AY:      "ئایمارا",
			language.AZ:      "ئازەربایجانی",
			language.AZ_ARAB: "ئازەربایجانی باشووری",
			language.BA:      "باشکیەر",
			language.BAN:     "بالی",
			language.BAS:     "باسا",
			language.BE:      "بیلاڕووسی",
			language.BEM:     "بێمبا",
			language.BEZ:     "بێنا",
			language.BG:      "بۆلگاری",
			language.BHO:     "بوجپووری",
			language.BI:      "بیسلاما",
			language.BIN:     "بینی",
			language.BLA:     "سیکسیکا",
			language.BM:      "بامبارا",
			language.BN:      "بەنگلادێشی",
			language.BO:      "تەبەتی",
			language.BR:      "برێتونی",
			language.BRX:     "بۆدۆ",
			language.BS:      "بۆسنی",
			language.BUG:     "بووگی",
			language.BYN:     "بلین",
			language.CA:      "كاتالۆنی",
			language.CE:      "چیچانی",
			language.CEB:     "سێبوانۆ",
			language.CGG:     "کیگا",
			language.CH:      "چامۆرۆ",
			language.CHK:     "چووکی",
			language.CHM:     "ماری",
			language.CHO:     "چۆکتاو",
			language.CHR:     "چێرۆکی",
			language.CHY:     "شایان",
			language.CKB:     "کوردیی ناوەندی",
			language.CO:      "کۆرسیکی",
			language.CRS:     "فەرەنسیی سیشێلی",
			language.CS:      "چێکی",
			language.CU:      "سلاویی کلیسەیی",
			language.CV:      "چووڤاشی",
			language.CY:      "وێلزی",
			language.DA:      "دانماركی",
			language.DAK:     "داکۆتایی",
			language.DAR:     "دارگینی",
			language.DAV:     "تایتا",
			language.DE:      "ئەڵمانی",
			language.DGR:     "دۆگریب",
			language.DJE:     "زارما",
			language.DSB:     "سربیی خوارین",
			language.DUA:     "دووالا",
			language.DV:      "دیڤێهی",
			language.DYO:     "جۆلافۆنی",
			language.DZ:      "دزوونگخا",
			language.DZG:     "دازا",
			language.EBU:     "ئێمبوو",
			language.EE:      "ئێوێیی",
			language.EFI:     "ئێفیک",
			language.EKA:     "ئێکاجووک",
			language.EL:      "یۆنانی",
			language.EN:      "ئینگلیزی",
			language.EN_AU:   "ئینگلیزیی ئۆسترالیایی",
			language.EN_CA:   "ئینگلیزیی کەنەدایی",
			language.EN_GB:   "ئینگلیزی (GB)",
			language.EN_US:   "ئینگلیزیی ئەمەریکایی",
			language.EO:      "ئێسپیرانتۆ",
			language.ES:      "ئیسپانی",
			language.ET:      "ئیستۆنی",
			language.EU:      "باسکی",
			language.EWO:     "ئێوۆندۆ",
			language.FA:      "فارسی",
			language.FF:      "فوولایی",
			language.FI:      "فینلەندی",
			language.FIL:     "فیلیپینی",
			language.FJ:      "فیجی",
			language.FO:      "فەرۆیی",
			language.FON:     "فۆنی",
			language.FR:      "فەرەنسی",
			language.FUR:     "فریئوولی",
			language.FY:      "فریسیی ڕۆژاوا",
			language.GA:      "ئیرلەندی",
			language.GAA:     "گایی",
			language.GD:      "گه\u200cلیكی سكۆتله\u200cندی",
			language.GEZ:     "گیزی",
			language.GIL:     "گیلبێرتی",
			language.GL:      "گالیسی",
			language.GN:      "گووارانی",
			language.GOR:     "گۆرۆنتالی",
			language.GSW:     "ئەڵمانیی سویسڕا",
			language.GU:      "گوجاراتی",
			language.GUZ:     "گووسی",
			language.GV:      "مانکی",
			language.GWI:     "گویچین",
			language.HA:      "هائووسا",
			language.HAW:     "هاوایی",
			language.HE:      "عیبری",
			language.HI:      "هیندی",
			language.HIL:     "هیلیگاینۆن",
			language.HMN:     "همۆنگ",
			language.HR:      "كرواتی",
			language.HSB:     "سربیی سەروو",
			language.HT:      "کریولی هائیتی",
			language.HU:      "هەنگاری",
			language.HUP:     "هووپا",
			language.HY:      "ئەرمەنی",
			language.HZ:      "هێرێرۆ",
			language.IA:      "ئینترلینگووا",
			language.IBA:     "ئیبان",
			language.IBB:     "ئیبیبۆ",
			language.ID:      "ئیندۆنیزی",
			language.IE:      "ئینتەرلیگ",
			language.IG:      "ئیگبۆ",
			language.II:      "سیچوان یی",
			language.ILO:     "ئیلۆکۆ",
			language.INH:     "ئینگووش",
			language.IO:      "ئیدۆ",
			language.IS:      "ئیسلەندی",
			language.IT:      "ئیتالی",
			language.IU:      "ئینوکتیتوت",
			language.JA:      "ژاپۆنی",
			language.JBO:     "لۆژبان",
			language.JGO:     "نگۆمبا",
			language.JMC:     "ماچامێ",
			language.JV:      "جاڤایی",
			language.KA:      "گۆرجستانی",
			language.KAB:     "کبائیلی",
			language.KAC:     "کاچین",
			language.KAJ:     "کیجوو",
			language.KAM:     "کامبا",
			language.KBD:     "کاباردی",
			language.KCG:     "تیاپ",
			language.KDE:     "ماکۆندە",
			language.KEA:     "کابووڤێردیانۆ",
			language.KFO:     "کۆرۆ",
			language.KHA:     "کهاسی",
			language.KHQ:     "کۆیرا چینی",
			language.KI:      "کیکوویوو",
			language.KJ:      "کوانیاما",
			language.KK:      "کازاخی",
			language.KKJ:     "کاکۆ",
			language.KL:      "کالالیسووت",
			language.KLN:     "کالێنجین",
			language.KM:      "خمێر",
			language.KMB:     "کیمبووندوو",
			language.KN:      "کاننادا",
			language.KO:      "كۆری",
			language.KOK:     "کۆنکانی",
			language.KPE:     "کپێلێ",
			language.KR:      "کانووری",
			language.KRC:     "کاراچای بالکار",
			language.KRL:     "کارێلی",
			language.KRU:     "کوورووخ",
			language.KS:      "کەشمیری",
			language.KSB:     "شامابالا",
			language.KSF:     "بافیا",
			language.KSH:     "کۆلۆنی",
			language.KU:      "کوردی",
			language.KUM:     "کوومیک",
			language.KV:      "کۆمی",
			language.KW:      "کۆڕنی",
			language.KY:      "كرگیزی",
			language.LA:      "لاتینی",
			language.LAD:     "لادینۆ",
			language.LAG:     "لانگی",
			language.LB:      "لوکسەمبورگی",
			language.LEZ:     "لەزگی",
			language.LG:      "گاندا",
			language.LI:      "لیمبورگی",
			language.LKT:     "لاکۆتا",
			language.LN:      "لينگالا",
			language.LO:      "لائۆیی",
			language.LOZ:     "لۆزی",
			language.LRC:     "لوڕیی باکوور",
			language.LT:      "لیتوانی",
			language.LU:      "لووبا کاتانگا",
			language.LUA:     "لووبا لوولووا",
			language.LUN:     "لووندا",
			language.LUO:     "لووئۆ",
			language.LUS:     "میزۆ",
			language.LUY:     "لوویا",
			language.LV:      "لێتۆنی",
			language.MAD:     "مادووری",
			language.MAG:     "ماگاهی",
			language.MAI:     "مائیتیلی",
			language.MAK:     "ماکاسار",
			language.MAS:     "ماسایی",
			language.MDF:     "مۆکشا",
			language.MEN:     "مێندێ",
			language.MER:     "مێروو",
			language.MFE:     "مۆریسی",
			language.MG:      "مالاگاسی",
			language.MGH:     "ماخوامیتۆ",
			language.MGO:     "مێتە",
			language.MH:      "مارشاڵی",
			language.MI:      "مائۆری",
			language.MIC:     "میکماک",
			language.MIN:     "مینانکاباو",
			language.MK:      "ماكێدۆنی",
			language.ML:      "مالایالام",
			language.MN:      "مەنگۆلی",
			language.MNI:     "مانیپووری",
			language.MOH:     "مۆهاوک",
			language.MOS:     "مۆسی",
			language.MR:      "ماراتی",
			language.MS:      "مالیزی",
			language.MT:      "ماڵتی",
			language.MUA:     "موندانگ",
			language.MUL:     "چەند زمان",
			language.MUS:     "کریک",
			language.MWL:     "میراندی",
			language.MY:      "میانماری",
			language.MYV:     "ئێرزیا",
			language.MZN:     "مازەندەرانی",
			language.NA:      "نائوروو",
			language.NAP:     "ناپۆلی",
			language.NAQ:     "ناما",
			language.NB:      "نەرویژیی بۆکمال",
			language.ND:      "ئندێبێلێی باکوور",
			language.NE:      "نیپالی",
			language.NEW:     "نێواری",
			language.NG:      "ندۆنگا",
			language.NIA:     "نیاس",
			language.NIU:     "نیئوویی",
			language.NL:      "هۆڵەندی",
			language.NL_BE:   "فلێمی",
			language.NMG:     "کواسیۆ",
			language.NN:      "نەرویژیی نینۆرسک",
			language.NNH:     "نگیمبوون",
			language.NO:      "نۆروێژی",
			language.NOG:     "نۆگای",
			language.NQO:     "نکۆ",
			language.NR:      "ئندێبێلێی باشوور",
			language.NSO:     "سۆتۆی باکوور",
			language.NUS:     "نوێر",
			language.NV:      "ناڤاجۆ",
			language.NY:      "نیانجا",
			language.NYN:     "نیانکۆلێ",
			language.OC:      "ئۆکسیتانی",
			language.OM:      "ئۆرۆمۆ",
			language.OR:      "ئۆدیا",
			language.OS:      "ئۆسێتی",
			language.PA:      "پەنجابی",
			language.PAG:     "پانگاسینان",
			language.PAM:     "پامپانگا",
			language.PAP:     "پاپیامێنتۆ",
			language.PAU:     "پالائوویی",
			language.PCM:     "پیجینی نیجریا",
			language.PL:      "پۆڵەندی",
			language.PRG:     "پڕووسی",
			language.PS:      "پەشتوو",
			language.PT:      "پورتوگالی",
			language.QU:      "کێچوا",
			language.QUC:     "کیچەیی",
			language.RAP:     "ڕاپانوویی",
			language.RAR:     "ڕاڕۆتۆنگان",
			language.RM:      "ڕۆمانش",
			language.RN:      "ڕووندی",
			language.RO:      "ڕۆمانی",
			language.RO_MD:   "مۆڵداڤی",
			language.ROF:     "ڕۆمبۆ",
			language.ROOT:    "ڕووت",
			language.RU:      "ڕووسی",
			language.RUP:     "ئارمۆمانی",
			language.RW:      "کینیارواندا",
			language.RWK:     "ڕوا",
			language.SA:      "سانسکريت",
			language.SAD:     "سانداوێ",
			language.SAH:     "ساخا",
			language.SAQ:     "سامبووروو",
			language.SAT:     "سانتالی",
			language.SBA:     "نگامبای",
			language.SBP:     "سانگوو",
			language.SC:      "ساردینی",
			language.SCN:     "سیسیلی",
			language.SCO:     "سکۆتس",
			language.SD:      "سيندی",
			language.SDH:     "کوردیی باشووری",
			language.SE:      "سامیی باکوور",
			language.SEH:     "سێنا",
			language.SES:     "کۆیرابۆرۆ سێنی",
			language.SG:      "سانگۆ",
			language.SH:      "سێربۆكرواتی",
			language.SHI:     "شیلها",
			language.SHN:     "شان",
			language.SI:      "سینهالی",
			language.SK:      "سلۆڤاكی",
			language.SL:      "سلۆڤێنی",
			language.SM:      "سامۆیی",
			language.SMA:     "سامیی باشوور",
			language.SMJ:     "لوولێ سامی",
			language.SMN:     "ئیناری سامی",
			language.SMS:     "سامیی سکۆڵت",
			language.SN:      "شۆنا",
			language.SNK:     "سۆنینکێ",
			language.SO:      "سۆمالی",
			language.SQ:      "ئەڵبانی",
			language.SR:      "سربی",
			language.SRN:     "سرانان تۆنگۆ",
			language.SS:      "سواتی",
			language.SSY:     "ساهۆ",
			language.ST:      "سۆتۆی باشوور",
			language.SU:      "سوندانی",
			language.SUK:     "سووکووما",
			language.SV:      "سویدی",
			language.SW:      "سواهیلی",
			language.SW_CD:   "سواهیلیی کۆنگۆ",
			language.SWB:     "کۆمۆری",
			language.SYR:     "سریانی",
			language.TA:      "تامیلی",
			language.TE:      "تێلووگوو",
			language.TEM:     "تیمنێ",
			language.TEO:     "تێسوو",
			language.TET:     "تێتووم",
			language.TG:      "تاجیکی",
			language.TH:      "تایلەندی",
			language.TI:      "تیگرینیا",
			language.TIG:     "تیگرێ",
			language.TK:      "تورکمانی",
			language.TLH:     "كلینگۆن",
			language.TN:      "تسوانا",
			language.TO:      "تۆنگان",
			language.TPI:     "تۆکپیسین",
			language.TR:      "تورکی",
			language.TRV:     "تارۆکۆ",
			language.TS:      "تسۆنگا",
			language.TT:      "تاتاری",
			language.TUM:     "تومبووکا",
			language.TVL:     "تووڤالوو",
			language.TW:      "توی",
			language.TWQ:     "تاساواک",
			language.TY:      "تاهیتی",
			language.TYV:     "تووڤینی",
			language.TZM:     "ئەمازیغی ناوەڕاست",
			language.UDM:     "ئوودموورت",
			language.UG:      "ئۆیخۆری",
			language.UK:      "ئۆكراینی",
			language.UMB:     "ئومبووندوو",
			language.UND:     "زمانی نەناسراو",
			language.UR:      "ئۆردوو",
			language.UZ:      "ئوزبەکی",
			language.VAI:     "ڤایی",
			language.VE:      "ڤێندا",
			language.VI:      "ڤیەتنامی",
			language.VO:      "ڤۆلاپووک",
			language.VUN:     "ڤوونجوو",
			language.WA:      "والوون",
			language.WAE:     "والسێر",
			language.WAL:     "وۆلایتا",
			language.WAR:     "وارای",
			language.WO:      "وۆلۆف",
			language.XAL:     "کالمیک",
			language.XH:      "سسوسا",
			language.XOG:     "سۆگا",
			language.YAV:     "یانگبێن",
			language.YBB:     "یێمبا",
			language.YI:      "ییدیش",
			language.YO:      "یۆرووبا",
			language.YUE:     "کانتۆنی",
			language.ZGH:     "ئەمازیغیی مەغریب",
			language.ZH:      "چینی",
			language.ZH_HANS: "چینی (چینیی ئاسانکراو)",
			language.ZH_HANT: "چینی (چینیی دێرین)",
			language.ZU:      "زوولوو",
			language.ZUN:     "زوونی",
			language.ZXX:     "هیچ ناوەرۆکی زمانی نیە",
			language.ZZA:     "زازا",
		},
		Territories: cldr.Territories{
			territory.V_001: "جیهان",
			territory.V_002: "ئەفریقا",
			territory.V_003: "ئەمەریکای باکوور",
			territory.V_005: "ئەمەریکای باشوور",
			territory.V_009: "ئۆقیانووسیا",
			territory.V_011: "ڕۆژاوای ئەفریقا",
			territory.V_013: "ئەمریکای ناوەڕاست",
			territory.V_014: "ڕۆژھەڵاتی ئەفریقا",
			territory.V_015: "باکووری ئەفریقا",
			territory.V_017: "ناوەڕاستی ئەفریقا",
			territory.V_018: "باشووری ئەفریقا",
			territory.V_019: "ئەمەریکای باکوور و باشوور",
			territory.V_021: "ئەمریکای باکوور",
			territory.V_029: "کاریبی",
			territory.V_030: "ڕۆژهەڵاتی ئاسیا",
			territory.V_034: "باشووری ئاسیا",
			territory.V_035: "باشووری ڕۆژھەڵاتی ئاسیا",
			territory.V_039: "باشووری ئەورووپا",
			territory.V_053: "ئۆسترالیا",
			territory.V_054: "میلانێزیا",
			territory.V_057: "ناوچەی مایکرۆنیزیا",
			territory.V_061: "پۆلینیزیا",
			territory.V_142: "ئاسیا",
			territory.V_143: "ناوەڕاستی ئاسیا",
			territory.V_145: "ڕۆژاوای ئاسیا",
			territory.V_150: "ئەورووپا",
			territory.V_151: "ڕۆژهەڵاتی ئەورووپا",
			territory.V_154: "باکووری ئەورووپا",
			territory.V_155: "ڕۆژاوای ئەورووپا",
			territory.V_202: "ئەفریقای ژێر سەحرا",
			territory.V_419: "ئەمەریکای لاتین",
			territory.AC:    "دوورگەی ئاسینسیۆن",
			territory.AD:    "ئاندۆرا",
			territory.AE:    "میرنشینە یەکگرتووە عەرەبییەکان",
			territory.AF:    "ئەفغانستان",
			territory.AG:    "ئانتیگوا و باربودا",
			territory.AI:    "ئانگویلا",
			territory.AL:    "ئەڵبانیا",
			territory.AM:    "ئەرمەنستان",
			territory.AO:    "ئەنگۆلا",
			territory.AQ:    "ئانتارکتیکا",
			territory.AR:    "ئەرژەنتین",
			territory.AS:    "ساموای ئەمەریکایی",
			territory.AT:    "نەمسا",
			territory.AU:    "ئوسترالیا",
			territory.AW:    "ئارووبا",
			territory.AX:    "دوورگەکانی ئالاند",
			territory.AZ:    "ئازەربایجان",
			territory.BA:    "بۆسنیا و ھەرزەگۆڤینا",
			territory.BB:    "باربادۆس",
			territory.BD:    "بەنگلادیش",
			territory.BE:    "بەلژیک",
			territory.BF:    "بورکینافاسۆ",
			territory.BG:    "بولگاریا",
			territory.BH:    "بەحرەین",
			territory.BI:    "بوروندی",
			territory.BJ:    "بێنین",
			territory.BL:    "سەن بارتێلێمی",
			territory.BM:    "بێرموودا",
			territory.BN:    "بروونای",
			territory.BO:    "بۆلیڤیا",
			territory.BQ:    "دوورگە کاریبیەکانی هۆڵەندا",
			territory.BR:    "برازیل",
			territory.BS:    "بەھاما",
			territory.BT:    "بووتان",
			territory.BV:    "دوورگەی بووڤێ",
			territory.BW:    "بۆتسوانا",
			territory.BY:    "بیلاڕووس",
			territory.BZ:    "بەلیز",
			territory.CA:    "کەنەدا",
			territory.CC:    "دوورگەکانی کیلینگ",
			territory.CD:    "کۆماری دیموکراتیی کۆنگۆ",
			territory.CF:    "کۆماری ئەفریقای ناوەڕاست",
			territory.CG:    "کۆماری کۆنگۆ",
			territory.CH:    "سویسڕا",
			territory.CI:    "کۆتدیڤوار",
			territory.CK:    "دوورگەکانی کوک",
			territory.CL:    "چیلی",
			territory.CM:    "کامیرۆن",
			territory.CN:    "چین",
			territory.CO:    "کۆلۆمبیا",
			territory.CP:    "دوورگەی کلیپێرتۆن",
			territory.CR:    "کۆستاریکا",
			territory.CU:    "کووبا",
			territory.CV:    "کەیپڤەرد",
			territory.CW:    "کوراچاو",
			territory.CX:    "دوورگەی کریسمس",
			territory.CY:    "قیبرس",
			territory.CZ:    "کۆماری چیک",
			territory.DE:    "ئەڵمانیا",
			territory.DG:    "دیەگۆ گارسیا",
			territory.DJ:    "جیبووتی",
			territory.DK:    "دانمارک",
			territory.DM:    "دۆمینیکا",
			territory.DO:    "کۆماری دۆمینیکا",
			territory.DZ:    "جەزایر",
			territory.EA:    "سێئووتا و مێلییا",
			territory.EC:    "ئیکوادۆر",
			territory.EE:    "ئیستۆنیا",
			territory.EG:    "میسر",
			territory.EH:    "سەحرای ڕۆژاوا",
			territory.ER:    "ئەریتریا",
			territory.ES:    "ئیسپانیا",
			territory.ET:    "ئەتیۆپیا",
			territory.EU:    "یەکێتیی ئەورووپا",
			territory.EZ:    "ناوچەی یۆرۆ",
			territory.FI:    "فینلاند",
			territory.FJ:    "فیجی",
			territory.FK:    "دوورگەکانی مالڤیناس (دوورگەکانی فاڵکلاند)",
			territory.FM:    "مایکرۆنیزیا",
			territory.FO:    "دوورگەکانی فارەو",
			territory.FR:    "فەڕەنسا",
			territory.GA:    "گابۆن",
			territory.GB:    "شانشینی یەکگرتوو",
			territory.GD:    "گرینادا",
			territory.GE:    "گورجستان",
			territory.GF:    "گیانای فەرەنسا",
			territory.GG:    "گێرنزی",
			territory.GH:    "غەنا",
			territory.GI:    "گیبرالتار",
			territory.GL:    "گرینلاند",
			territory.GM:    "گامبیا",
			territory.GN:    "گینێ",
			territory.GP:    "گوادێلۆپ",
			territory.GQ:    "گینێی ئیستوایی",
			territory.GR:    "یۆنان",
			territory.GS:    "دوورگەکانی جۆرجیا و ساندویچی باشوور",
			territory.GT:    "گواتیمالا",
			territory.GU:    "گوام",
			territory.GW:    "گینێ بیساو",
			territory.GY:    "گویانا",
			territory.HK:    "هۆنگ کۆنگ",
			territory.HM:    "دوورگەکانی هێرد و مەکدانڵد",
			territory.HN:    "ھۆندووراس",
			territory.HR:    "کرۆواتیا",
			territory.HT:    "ھایتی",
			territory.HU:    "هەنگاریا",
			territory.IC:    "دوورگەکانی کەناری",
			territory.ID:    "ئیندۆنیزیا",
			territory.IE:    "ئیرلەند",
			territory.IL:    "ئیسرائیل",
			territory.IM:    "دوورگەی مان",
			territory.IN:    "ھیندستان",
			territory.IO:    "ھەرێمی بەریتانی لە ئۆقیانووسی ھیند",
			territory.IQ:    "عێراق",
			territory.IR:    "ئێران",
			territory.IS:    "ئایسلەند",
			territory.IT:    "ئیتالیا",
			territory.JE:    "جێرسی",
			territory.JM:    "جامایکا",
			territory.JO:    "ئوردن",
			territory.JP:    "ژاپۆن",
			territory.KE:    "کینیا",
			territory.KG:    "کرگیزستان",
			territory.KH:    "کەمبۆدیا",
			territory.KI:    "کیریباس",
			territory.KM:    "دوورگەکانی کۆمۆر",
			territory.KN:    "سەن کیتس و نیڤیس",
			territory.KP:    "کۆریای باکوور",
			territory.KR:    "کۆریای باشوور",
			territory.KW:    "کوەیت",
			territory.KY:    "دوورگەکانی کایمان",
			territory.KZ:    "کازاخستان",
			territory.LA:    "لاوس",
			territory.LB:    "لوبنان",
			territory.LC:    "سەن لووسیا",
			territory.LI:    "لیختنشتاین",
			territory.LK:    "سریلانکا",
			territory.LR:    "لیبەریا",
			territory.LS:    "لەسۆتۆ",
			territory.LT:    "لیتوانایا",
			territory.LU:    "لوکسەمبورگ",
			territory.LV:    "لاتڤیا",
			territory.LY:    "لیبیا",
			territory.MA:    "مەغریب",
			territory.MC:    "مۆناکۆ",
			territory.MD:    "مۆلدۆڤا",
			territory.ME:    "مۆنتینیگرۆ",
			territory.MF:    "سەن مارتین",
			territory.MG:    "ماداگاسکار",
			territory.MH:    "دوورگەکانی مارشاڵ",
			territory.MK:    "مەکدۆنیای باکوور",
			territory.ML:    "مالی",
			territory.MM:    "میانمار",
			territory.MN:    "مەنگۆلیا",
			territory.MO:    "ماکائۆ",
			territory.MP:    "دوورگەکانی ماریانای باکوور",
			territory.MQ:    "مارتینیک",
			territory.MR:    "مۆریتانیا",
			territory.MS:    "مۆنتسێرات",
			territory.MT:    "ماڵتا",
			territory.MU:    "مووریتیووس",
			territory.MV:    "مالدیڤ",
			territory.MW:    "مالاوی",
			territory.MX:    "مەکسیک",
			territory.MY:    "مالیزیا",
			territory.MZ:    "مۆزامبیک",
			territory.NA:    "نامیبیا",
			territory.NC:    "نیووکالێدۆنیا",
			territory.NE:    "نیجەر",
			territory.NF:    "دوورگەی نۆرفۆڵک",
			territory.NG:    "نیجریا",
			territory.NI:    "نیکاراگوا",
			territory.NL:    "ھۆڵەندا",
			territory.NO:    "نۆرویژ",
			territory.NP:    "نیپال",
			territory.NR:    "نائوروو",
			territory.NU:    "نیووئی",
			territory.NZ:    "نیوزیلاند",
			territory.OM:    "عومان",
			territory.PA:    "پاناما",
			territory.PE:    "پێروو",
			territory.PF:    "پۆلینیسیای فەرەنسا",
			territory.PG:    "پاپوا گینێی نوێ",
			territory.PH:    "فلیپین",
			territory.PK:    "پاکستان",
			territory.PL:    "پۆڵەندا",
			territory.PM:    "سەن پیێر و میکێلۆن",
			territory.PN:    "دوورگەکانی پیتکەرن",
			territory.PR:    "پۆرتۆڕیکۆ",
			territory.PS:    "فەلەستین",
			territory.PT:    "پورتوگال",
			territory.PW:    "پالاو",
			territory.PY:    "پاراگوای",
			territory.QA:    "قەتەر",
			territory.QO:    "دەرەوەی ئۆقیانووسیا",
			territory.RE:    "ڕییوونیەن",
			territory.RO:    "ڕۆمانیا",
			territory.RS:    "سربیا",
			territory.RU:    "ڕووسیا",
			territory.RW:    "ڕواندا",
			territory.SA:    "عەرەبستانی سەعوودی",
			territory.SB:    "دوورگەکانی سلێمان",
			territory.SC:    "سیشێل",
			territory.SD:    "سوودان",
			territory.SE:    "سوید",
			territory.SG:    "سینگاپور",
			territory.SH:    "سەن هێلێنا",
			territory.SI:    "سلۆڤێنیا",
			territory.SJ:    "سڤالبارد و یان مایەن",
			territory.SK:    "سلۆڤاکیا",
			territory.SL:    "سیەرالیۆن",
			territory.SM:    "سان مارینۆ",
			territory.SN:    "سێنێگاڵ",
			territory.SO:    "سۆمالیا",
			territory.SR:    "سورینام",
			territory.SS:    "سوودانی باشوور",
			territory.ST:    "ساوتۆمێ و پرینسیپی",
			territory.SV:    "ئێلسالڤادۆر",
			territory.SX:    "سینت مارتن",
			territory.SY:    "سووریا",
			territory.SZ:    "سوازیلاند",
			territory.TA:    "تریستێن دا کوونا",
			territory.TC:    "دوورگەکانی تورکس و کایکۆس",
			territory.TD:    "چاد",
			territory.TF:    "هەرێمە باشووریەکانی فەرەنسا",
			territory.TG:    "تۆگۆ",
			territory.TH:    "تایلەند",
			territory.TJ:    "تاجیکستان",
			territory.TK:    "تۆکێلاو",
			territory.TL:    "تیمۆری ڕۆژھەڵات",
			territory.TM:    "تورکمانستان",
			territory.TN:    "توونس",
			territory.TO:    "تۆنگا",
			territory.TR:    "تورکیا",
			territory.TT:    "ترینیداد و تۆباگو",
			territory.TV:    "تووڤالوو",
			territory.TW:    "تایوان",
			territory.TZ:    "تانزانیا",
			territory.UA:    "ئۆکرانیا",
			territory.UG:    "ئوگاندا",
			territory.UM:    "دوورگەکانی دەرەوەی ئەمریکا",
			territory.UN:    "نەتەوە یەکگرتووەکان",
			territory.US:    "ویلایەتە یەکگرتووەکان",
			territory.UY:    "ئوروگوای",
			territory.UZ:    "ئوزبەکستان",
			territory.VA:    "ڤاتیکان",
			territory.VC:    "سەینت ڤینسەنت و گرینادینز",
			territory.VE:    "ڤەنزوێلا",
			territory.VG:    "دوورگەکانی ڤیرجنی بەریتانیا",
			territory.VI:    "دوورگەکانی ڤیرجنی ئەمەریکا",
			territory.VN:    "ڤیەتنام",
			territory.VU:    "ڤانوواتوو",
			territory.WF:    "والیس و فوتونا",
			territory.WS:    "ساموا",
			territory.XK:    "کۆسۆڤۆ",
			territory.YE:    "یەمەن",
			territory.YT:    "مایۆت",
			territory.ZA:    "ئەفریقای باشوور",
			territory.ZM:    "زامبیا",
			territory.ZW:    "زیمبابوی",
			territory.ZZ:    "ناوچەی نەناسراو",
		},
	}
}
