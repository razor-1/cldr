// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_lrc() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "lrc",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMM d, EEEE", Long: "y MMMM d", Medium: "y MMM d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "جانڤیە", Feb: "فئڤریە", Mar: "مارس", Apr: "آڤریل", May: "مئی", Jun: "جوٙأن", Jul: "جوٙلا", Aug: "آگوست", Sep: "سئپتامر", Oct: "ئوکتوڤر", Nov: "نوڤامر", Dec: "دئسامر"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "رئال بئرئزیل", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "یوان چین", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "یورو", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "پوند بئریتانیا", Symbol: "£"},
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
				currency.INR: cldr.Currency{DisplayName: "روٙپیه هئن", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "", Symbol: "د.ع.\u200f"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "یئن جاپوٙن", Symbol: "JP¥"},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
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
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "روٙبل روٙسیه", Symbol: "₽"},
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
				currency.XXX: cldr.Currency{DisplayName: "پیل نادیار", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AB:      "آذأربایئجانی",
			language.AF:      "آفریکانس",
			language.AGQ:     "آقئم",
			language.AK:      "آکان",
			language.AM:      "أمھأری",
			language.AR:      "أرأڤی",
			language.AR_001:  "عروی مدرن",
			language.ARN:     "ماپوٙچئ",
			language.AS:      "آسامی",
			language.ASA:     "آسوٙ",
			language.AZ:      "آذأری",
			language.AZ_ARAB: "آذأری ھارگە",
			language.BA:      "باشکیری",
			language.BE:      "بئلاروٙسی",
			language.BEM:     "بیما",
			language.BEZ:     "بئنا",
			language.BG:      "بولغاری",
			language.BGN:     "بألوٙچی أقتوٙنئشین",
			language.BM:      "بامبارا",
			language.BN:      "بأنگالی",
			language.BO:      "تأبأتی",
			language.BR:      "بئرئتون",
			language.BRX:     "بودو",
			language.BS:      "بوسنیایی",
			language.CA:      "کاتالان",
			language.CE:      "چئچئنی",
			language.CGG:     "چیگا",
			language.CHR:     "چوروٙکی",
			language.CKB:     "کوردی سوٙرانی",
			language.CO:      "کوریسکان",
			language.CV:      "چواشی",
			language.CY:      "ڤئلزی",
			language.DA:      "دانمارکی",
			language.DAV:     "تایتا",
			language.DE:      "آلمانی",
			language.DE_AT:   "آلمانی ئوتریشی",
			language.DE_CH:   "آلمانی سوٙییسی",
			language.DJE:     "زارما",
			language.DSB:     "سوربی ھاری",
			language.DUA:     "دوٙالا",
			language.DYO:     "جولا فوٙنیی",
			language.DZ:      "زوٙنگخا",
			language.EBU:     "ئمبو",
			language.EE:      "ئڤئ",
			language.EL:      "یوٙنانی",
			language.EN:      "ئینگیلیسی",
			language.EN_AU:   "ئینگیلیسی ئوستارالیایی",
			language.EN_CA:   "ئینگیلیسی کانادایی",
			language.EN_GB:   "ئینگیلیسی بئریتانیا گأپ",
			language.EN_US:   "ئینگیلیسی ئمریکایی",
			language.EO:      "ئسپئرانتو",
			language.ES:      "ئسپانیایی",
			language.ES_419:  "ئسپانیایی ئمریکا لاتین",
			language.ES_ES:   "ئسپانیایی ئوروٙپا",
			language.ES_MX:   "ئسپانیایی مئکزیک",
			language.ET:      "ئستونیایی",
			language.EU:      "باسکی",
			language.FA:      "فارسی",
			language.FI:      "فأنلاندی",
			language.FIL:     "فیلیپینی",
			language.FJ:      "فیجی",
			language.FO:      "فاروٙسی",
			language.FR:      "فآرانسئ ئی",
			language.FR_CA:   "فآرانسئ ئی کانادا",
			language.FR_CH:   "فآرانسئ ئی سوٙییس",
			language.FY:      "فئریسی أفتونئشین",
			language.GA:      "ئیرلأندی",
			language.GAG:     "گاگائوز",
			language.GL:      "گالیسی",
			language.GN:      "گوٙآرانی",
			language.GSW:     "آلمانی سوٙئیسی",
			language.GU:      "گوجأراتی",
			language.GUZ:     "گوٙسی",
			language.GV:      "مانکس",
			language.HA:      "ھائوسا",
			language.HAW:     "ھاڤایی",
			language.HE:      "عئبری",
			language.HI:      "ھئنی",
			language.HR:      "کوروڤاتی",
			language.HSB:     "سوربی ڤارو",
			language.HT:      "ھاییتی",
			language.HU:      "مأجاری",
			language.HY:      "أرمأنی",
			language.ID:      "أندونئزیایی",
			language.IG:      "ئیگبو",
			language.II:      "سی چوان یی",
			language.IS:      "ئیسلأندی",
			language.IT:      "ئیتالیایی",
			language.IU:      "ئینوکتیتوٙت",
			language.JA:      "جاپوٙنی",
			language.JGO:     "نئگوٙمبا",
			language.JMC:     "ماچامئ",
			language.JV:      "جاڤئ یی",
			language.KA:      "گورجی",
			language.KAB:     "کابیلئ",
			language.KAM:     "کامبا",
			language.KDE:     "ماکوٙندئ",
			language.KEA:     "کاباردینو",
			language.KHQ:     "کی یورا چینی",
			language.KI:      "کیکیوٙ",
			language.KK:      "قأزاق",
			language.KL:      "کالالیسوٙت",
			language.KLN:     "کالئجین",
			language.KM:      "خئمئر",
			language.KN:      "کاناد",
			language.KO:      "کورئ یی",
			language.KOI:     "کومی پئرمیاک",
			language.KOK:     "کوٙنکانی",
			language.KS:      "کأشمیری",
			language.KSB:     "شامبالا",
			language.KSF:     "بافیا",
			language.KU:      "کوردی کورمانجی",
			language.KW:      "کورنیش",
			language.KY:      "قئرقیزی",
			language.LA:      "لاتین",
			language.LAG:     "لانگی",
			language.LB:      "لوٙکزامبوٙرگی",
			language.LG:      "گاندا",
			language.LKT:     "لاکوٙتا",
			language.LN:      "لینگالا",
			language.LO:      "لاو",
			language.LRC:     "لۊری شومالی",
			language.LT:      "لیتوڤانیایی",
			language.LU:      "لوٙبا کاتانگا",
			language.LUO:     "لوٙ",
			language.LUY:     "لوٙئیا",
			language.LV:      "لاتوڤیایی",
			language.MAS:     "ماسایی",
			language.MER:     "مئرو",
			language.MFE:     "موٙریسی",
			language.MG:      "مالاگاشی",
			language.MGH:     "ماخوڤا میتو",
			language.MGO:     "مئتاٛ",
			language.MI:      "مائوری",
			language.MK:      "مأقدوٙنی",
			language.ML:      "مالایام",
			language.MN:      "موغولی",
			language.MOH:     "موٙھاڤک",
			language.MR:      "مأراتی",
			language.MS:      "مالایی",
			language.MT:      "مالتی",
			language.MUA:     "موٙندانگ",
			language.MY:      "بئرمئ یی",
			language.MZN:     "مازأندأرانی",
			language.NAQ:     "ناما",
			language.NB:      "نورڤئجی بوٙکمال",
			language.ND:      "نئدئبئلئ شومالی",
			language.NDS:     "آلمانی ھاری",
			language.NDS_NL:  "آلمانی ھارگە جا",
			language.NE:      "نئپالی",
			language.NL:      "ھولأندی",
			language.NL_BE:   "فئلاماندی",
			language.NMG:     "کئڤاسیوٙ",
			language.NN:      "نورڤئجی نینورسک",
			language.NQO:     "نئکوٙ",
			language.NUS:     "نیوٙئر",
			language.NYN:     "نیان کوٙلئ",
			language.OM:      "ئوروموٙ",
			language.OR:      "ئوریا",
			language.PA:      "پأنجابی",
			language.PL:      "لأھئستانی",
			language.PS:      "پأشتوٙ",
			language.PT:      "پورتئغالی",
			language.PT_BR:   "پورتئغالی بئرئزیل",
			language.PT_PT:   "پورتئغالی ئوروٙپایی",
			language.QU:      "کوچوٙا",
			language.QUC:     "کیچی",
			language.RM:      "رومانش",
			language.RN:      "راندی",
			language.RO:      "رومانیایی",
			language.RO_MD:   "رومانیایی مولداڤی",
			language.ROF:     "رومبو",
			language.RU:      "روٙسی",
			language.RW:      "کینیاروآندا",
			language.RWK:     "رئڤا",
			language.SA:      "سانسکئریت",
			language.SAQ:     "سامبوٙروٙ",
			language.SBP:     "سانگوٙ",
			language.SD:      "سئندی",
			language.SDH:     "کوردی ھارگە",
			language.SE:      "سامی شومالی",
			language.SEH:     "سئنا",
			language.SES:     "کیارابورو سئنی",
			language.SG:      "سانگو",
			language.SHI:     "تاچئلھیت",
			language.SI:      "سینھالا",
			language.SK:      "ئسلوڤاکی",
			language.SL:      "ئسلوڤئنیایی",
			language.SMA:     "سامی ھارگە",
			language.SMJ:     "لۉلئ سامی",
			language.SMN:     "ئیناری سامی",
			language.SMS:     "ئسکولت سامی",
			language.SN:      "شونا",
			language.SO:      "سوٙمالی",
			language.SQ:      "آلبانی",
			language.SR:      "سئربی",
			language.SU:      "سوٙدانی",
			language.SV:      "سوٙئدی",
			language.SW:      "سأڤاحیلی",
			language.SW_CD:   "سأڤاحیلی کونگو",
			language.TA:      "تامیل",
			language.TE:      "تئلئگو",
			language.TEO:     "تئسو",
			language.TG:      "تاجیکی",
			language.TH:      "تایلأندی",
			language.TI:      "تیگرینیا",
			language.TK:      "تورکأمأنی",
			language.TO:      "توٙنگان",
			language.TR:      "تورکی",
			language.TT:      "تاتار",
			language.TWQ:     "تاساڤاق",
			language.TZM:     "تامازیغ مینجایی",
			language.UG:      "ئویغوٙر",
			language.UK:      "ئوکراینی",
			language.UND:     "زوٙن نادیار",
			language.UR:      "ئوردوٙ",
			language.UZ:      "ئوزبأکی",
			language.VAI:     "ڤای",
			language.VI:      "ڤییئتنامی",
			language.VUN:     "ڤوٙنجوٙ",
			language.WBP:     "ڤارلپیری",
			language.WO:      "ڤولوف",
			language.XH:      "خوٙسا",
			language.XOG:     "سوٙگا",
			language.YO:      "یوروبا",
			language.ZGH:     "تامازیغ مأراکئشی",
			language.ZH:      "چینی",
			language.ZH_HANS: "چینی سادە بیە",
			language.ZH_HANT: "چینی سونأتی",
			language.ZU:      "زولو",
			language.ZXX:     "بی نئشوٙ",
		},
		Territories: cldr.Territories{
			territory.V_001: "دونیا",
			territory.V_002: "ئفریقا",
			territory.V_003: "ئمریکا شومالی",
			territory.V_005: "ئمریکا ھارگە",
			territory.V_009: "ھوم پئڤأند جأھوٙن آڤ",
			territory.V_013: "مینجا ئمریکا",
			territory.V_019: "ئمریکا",
			territory.V_021: "ئمریکا ڤارو",
			territory.V_029: "کارائیب",
			territory.V_142: "آسیا",
			territory.V_150: "ئوروٙپا",
			territory.V_419: "ئمریکا لاتین",
			territory.BR:    "بئرئزیل",
			territory.CN:    "چین",
			territory.DE:    "آلمان",
			territory.FR:    "فأرانسە",
			territory.GB:    "بیریتانیا گأپ",
			territory.IN:    "ھئن",
			territory.IT:    "ئیتالیا",
			territory.JP:    "جاپوٙن",
			territory.RU:    "روٙسیە",
			territory.US:    "ڤولاتیا یأکاگئرتە",
			territory.ZZ:    "راساگە نادیار",
		},
	}
}
