// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_tt_RU() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "tt_RU",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "d MMMM, y\u202f'ел', EEEE", Long: "d MMMM, y\u202f'ел'", Medium: "d MMM, y\u202f'ел'", Short: "dd.MM.y"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "гыйн.", Feb: "фев.", Mar: "мар.", Apr: "апр.", May: "май", Jun: "июнь", Jul: "июль", Aug: "авг.", Sep: "сент.", Oct: "окт.", Nov: "нояб.", Dec: "дек."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "гыйнвар", Feb: "февраль", Mar: "март", Apr: "апрель", May: "май", Jun: "июнь", Jul: "июль", Aug: "август", Sep: "сентябрь", Oct: "октябрь", Nov: "ноябрь", Dec: "декабрь"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "якш.", Mon: "дүш.", Tue: "сиш.", Wed: "чәр.", Thu: "пәнҗ.", Fri: "җом.", Sat: "шим."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Я", Mon: "Д", Tue: "С", Wed: "Ч", Thu: "П", Fri: "Җ", Sat: "Ш"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "якшәмбе", Mon: "дүшәмбе", Tue: "сишәмбе", Wed: "чәршәмбе", Thu: "пәнҗешәмбе", Fri: "җомга", Sat: "шимбә"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00¤", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Берләшкән Гарәп Әмирлекләре дирхамы", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Әфган әфганы", Symbol: "؋"},
				currency.ALL: cldr.Currency{DisplayName: "Албания Леке", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Армения драмы", Symbol: "֏"},
				currency.ANG: cldr.Currency{DisplayName: "Нидерланд Антиль утраулары гульдены", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Ангола кванзасы", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Аргентина песосы", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Австралия доллары", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Арубан Флорины", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Әзербайҗан манаты", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "Босния һәм Герцеговинаның конвертацияләнә торган маркасы", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Барбадос доллары", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Бангладеш такасы", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Болгар Левы", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Бахрейн динары", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Бурунди франкы", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Бермуд доллары", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Бруней доллары", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Боливиа боливианосы", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Бразилия реалы", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Багам доллары", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Бутан нгултрумы", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Ботсвана пуласы", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Беларус сумы", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Белиз доллары", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Канада доллары", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Конголез франкы", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Швейцария франкы", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Чили песосы", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Кытай Юане (оффшор)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "кытай юане", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Колумбия песосы", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Коста-Рика колоны", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Куба конвертацияләнә торган песосы", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Куба песосы", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Кабо-Верде эскудосы", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Чех кронасы", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Җибути франкы", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Дания Кронасы", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Доминикана песосы", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Алжир динары", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Мисыр фунты", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Эритрея накфасы", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Эфиопия быры", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "евро", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Фиджи доллары", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Фолкленд утраулары фунты", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "фунт стерлинг", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Грузия ларие", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "Гана седие", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "Гибралтар фунты", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Гамбия даласие", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Гвинея франкы", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Гватемала кетсалы", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Гайана доллары", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Гонконг доллары", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Гондурас лемпиры", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Хорватия кунасы", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Гаити гурды", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Венгрия форинты", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Индонезия рупиясе", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Израиль яңа шекеле", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Индия рупиясе", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Ирак динары", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Иран риалы", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Исландия Кронасы", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Ямайка доллары", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Иордания динары", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "япон иенасы", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Кения шиллингы", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Ккргызстан сомы", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "Камбоджа риелы", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Комор утраулары франкы", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Төньяк Корея Вонасы", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Көньяк Корея Вонасы", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Кувейт динары", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Кайман утраулары доллары", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Казахстан тәңкәсе", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Лаос кипы", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Ливан фунты", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Шри-Ланка рупиясе", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Либерия доллары", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Лесото лотисы", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Ливия динары", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Марокко дирхамы", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Молдавия Лее", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Малагаси ариариясе", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Македония денары", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Мьянма кьяты", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Монголия тугрикы", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Макао патакы", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Мавритан угиясы", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Маврикий рупие", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Мальдив руфиясе", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Малавия квачие", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Мексика песосы", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Малайзия ринггиты", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Мозамбик метикалы", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Намибия доллары", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Нигерия найрасы", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Никарагуа кордовасы", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Норвегия Кронасы", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Непал рупиясе", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Яңа Зеландия доллары", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Оман риалы", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Панама бальбоасы", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Перу солы", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Папуа Яңа Гвинея Кинасы", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Филиппин песосы", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Пакистан рупиясе", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Польша злотые", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Парагвай гуаранисы", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Катар риалы", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Румыния Лее", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Сербия динары", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Россия сумы", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Руанда франкы", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Согуд Гарәбстаны риалы", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Соломон утраулары доллары", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Сейшел утраулары рупиясы", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Судан фунты", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Швеция Кронасы", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Сингапур доллары", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Изге Елена утравы фунты", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Сьерра-Леон леоны", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Сьерра-Леоне леоны (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Сомали шиллингы", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Суринам доллары", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Көньяк Судан фунты", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "Сан-Томе һәм Принсипи добрасы", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Сурия фунты", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Свази эмалангенисы", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Тайвань Баты", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Таҗикстан сомонисы", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Төркмәнстан Манаты", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Тунис динары", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Тонга Паангасы", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Төркия Лирасы", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Тринидад һәм Тобаго доллары", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Яңа Тайвань доллары", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Танзания шиллингы", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Украина гривнасы", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Уганда шиллингы", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "АКШ доллары", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Уругвай песосы", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Үзбәкстан Сомы", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Венесуэла боливары", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Вьетнам Донгы", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Вануату ваты", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Самоа Таласы", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Үзәк Африка франкы КФА", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Көнчыгыш Кариб доллары", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Көнбатыш Африка КФА франкы", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "Франциянең диңгез аръягы җәмгыятьләре франкы", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "билгесез валюта", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Йәмән риалы", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Көньяк Африка Рэнды", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Замбия квачасы", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AF:      "африкаанс",
			language.AM:      "амхар",
			language.AR:      "гарәп",
			language.AR_001:  "Заманча стандарт гарәп",
			language.ARN:     "мапуче",
			language.AS:      "ассам",
			language.AZ:      "әзәрбайҗан",
			language.BA:      "башкорт",
			language.BAN:     "бали",
			language.BE:      "белорус",
			language.BEM:     "бемба",
			language.BG:      "болгар",
			language.BN:      "бенгали",
			language.BO:      "тибет",
			language.BR:      "бретон",
			language.BS:      "босния",
			language.CA:      "каталан",
			language.CEB:     "себуано",
			language.CHM:     "мари",
			language.CHR:     "чероки",
			language.CKB:     "үзәк көрд",
			language.CO:      "корсика",
			language.CS:      "чех",
			language.CY:      "уэльс",
			language.DA:      "дания",
			language.DE:      "алман",
			language.DE_CH:   "югары алман (Швейцария)",
			language.DSB:     "түбән сорб",
			language.DV:      "мальдив",
			language.DZ:      "дзонг-кха",
			language.EL:      "грек",
			language.EN:      "инглиз",
			language.EN_GB:   "инглиз (Берләшкән Корольлек)",
			language.EN_US:   "инглиз (АКШ)",
			language.EO:      "эсперанто",
			language.ES:      "испан",
			language.ES_419:  "испан (Латин Америкасы)",
			language.ES_ES:   "испан (Европа)",
			language.ET:      "эстон",
			language.EU:      "баск",
			language.FA:      "фарсы",
			language.FF:      "фула",
			language.FI:      "фин",
			language.FIL:     "филиппин",
			language.FO:      "фарер",
			language.FR:      "француз",
			language.GA:      "ирланд",
			language.GD:      "шотланд гэль",
			language.GL:      "галисия",
			language.GN:      "гуарани",
			language.GU:      "гуҗарати",
			language.HA:      "хауса",
			language.HAW:     "гавайи",
			language.HE:      "яһүд",
			language.HI:      "һинд",
			language.HI_LATN: "Һинглиш",
			language.HIL:     "хилигайнон",
			language.HR:      "хорват",
			language.HSB:     "югары сорб",
			language.HT:      "гаити креол",
			language.HU:      "венгр",
			language.HY:      "әрмән",
			language.HZ:      "гереро",
			language.IBB:     "ибибио",
			language.ID:      "индонезия",
			language.IG:      "игбо",
			language.IS:      "исланд",
			language.IT:      "итальян",
			language.IU:      "инуктикут",
			language.JA:      "япон",
			language.KA:      "грузин",
			language.KK:      "казакъ",
			language.KM:      "кхмер",
			language.KN:      "каннада",
			language.KO:      "корея",
			language.KOK:     "конкани",
			language.KR:      "канури",
			language.KRU:     "курух",
			language.KS:      "кашмири",
			language.KU:      "көрд",
			language.KY:      "кыргыз",
			language.LA:      "латин",
			language.LB:      "люксембург",
			language.LO:      "лаос",
			language.LT:      "литва",
			language.LV:      "латыш",
			language.MEN:     "менде",
			language.MG:      "малагаси",
			language.MI:      "маори",
			language.MK:      "македон",
			language.ML:      "малаялам",
			language.MN:      "монгол",
			language.MNI:     "манипури",
			language.MOH:     "могаук",
			language.MR:      "маратхи",
			language.MS:      "малай",
			language.MT:      "мальта",
			language.MY:      "бирма",
			language.NE:      "непали",
			language.NIU:     "ниуэ",
			language.NL:      "голланд",
			language.NL_BE:   "фламандча",
			language.NY:      "ньянҗа",
			language.OC:      "окситан",
			language.OM:      "оромо",
			language.OR:      "ория",
			language.PA:      "пәнҗаби",
			language.PAP:     "папьяменто",
			language.PL:      "поляк",
			language.PS:      "пушту",
			language.PT:      "португал",
			language.PT_PT:   "португал (Европа)",
			language.QU:      "кечуа",
			language.QUC:     "киче",
			language.RM:      "ретороман",
			language.RO:      "румын",
			language.RU:      "рус",
			language.RW:      "руанда",
			language.SA:      "санскрит",
			language.SAH:     "саха",
			language.SAT:     "сантали",
			language.SD:      "синдһи",
			language.SE:      "төньяк саам",
			language.SI:      "сингал",
			language.SK:      "словак",
			language.SL:      "словен",
			language.SMA:     "көньяк саам",
			language.SMJ:     "луле-саам",
			language.SMN:     "инари-саам",
			language.SMS:     "колтта-саам",
			language.SO:      "сомали",
			language.SQ:      "албан",
			language.SR:      "серб",
			language.SV:      "швед",
			language.SYR:     "сүрия",
			language.TA:      "тамил",
			language.TE:      "телугу",
			language.TG:      "таҗик",
			language.TH:      "тай",
			language.TI:      "тигринья",
			language.TK:      "төрекмән",
			language.TO:      "тонга",
			language.TR:      "төрек",
			language.TT:      "татар",
			language.TZM:     "үзәк атлас тамазигт",
			language.UG:      "уйгыр",
			language.UK:      "украин",
			language.UND:     "билгесез тел",
			language.UR:      "урду",
			language.UZ:      "үзбәк",
			language.VE:      "венда",
			language.VI:      "вьетнам",
			language.WO:      "волоф",
			language.YI:      "идиш",
			language.YO:      "йоруба",
			language.ZH:      "мандарин кытайчасы",
			language.ZH_HANS: "гадиләштерелгән мандарин кытайчасы",
			language.ZH_HANT: "традицион мандарин кытайчасы",
		},
		Territories: cldr.Territories{
			territory.V_001: "дөнья",
			territory.V_002: "Африка",
			territory.V_003: "Төньяк Америка",
			territory.V_005: "Көньяк Америка",
			territory.V_009: "Океания",
			territory.V_011: "Көнбатыш Африка",
			territory.V_013: "Үзәк Америка",
			territory.V_014: "Көнчыгыш Африка",
			territory.V_015: "Төньяк Африка",
			territory.V_017: "Урта Африка",
			territory.V_018: "Көньяктагы Африка",
			territory.V_019: "Америка",
			territory.V_021: "Төньяктагы Америка",
			territory.V_029: "Кариб бассейны",
			territory.V_030: "Көнчыгыш Азия",
			territory.V_034: "Көньяк Азия",
			territory.V_035: "Көньяк-Көнчыгыш Азия",
			territory.V_039: "Көньяк Европа",
			territory.V_053: "Австралазия",
			territory.V_054: "Меланезия",
			territory.V_057: "Микронезия төбәге",
			territory.V_061: "Полинезия",
			territory.V_142: "Азия",
			territory.V_143: "Үзәк Азия",
			territory.V_145: "Көнбатыш Азия",
			territory.V_150: "Европа",
			territory.V_151: "Көнчыгыш Европа",
			territory.V_154: "Төньяк Европа",
			territory.V_155: "Көнбатыш Европа",
			territory.V_202: "Сахарадан көньякта Африка",
			territory.V_419: "Латин Америка",
			territory.AC:    "Вознесение утравы",
			territory.AD:    "Андорра",
			territory.AE:    "Берләшкән Гарәп Әмирлекләре",
			territory.AF:    "Әфганстан",
			territory.AG:    "Антигуа һәм Барбуда",
			territory.AI:    "Ангилья",
			territory.AL:    "Албания",
			territory.AM:    "Әрмәнстан",
			territory.AO:    "Ангола",
			territory.AQ:    "Антарктика",
			territory.AR:    "Аргентина",
			territory.AS:    "Америка Самоасы",
			territory.AT:    "Австрия",
			territory.AU:    "Австралия",
			territory.AW:    "Аруба",
			territory.AX:    "Аланд утраулары",
			territory.AZ:    "Әзәрбайҗан",
			territory.BA:    "Босния һәм Герцеговина",
			territory.BB:    "Барбадос",
			territory.BD:    "Бангладеш",
			territory.BE:    "Бельгия",
			territory.BF:    "Буркина-Фасо",
			territory.BG:    "Болгария",
			territory.BH:    "Бәхрәйн",
			territory.BI:    "Бурунди",
			territory.BJ:    "Бенин",
			territory.BL:    "Сен-Бартельми",
			territory.BM:    "Бермуд утраулары",
			territory.BN:    "Бруней",
			territory.BO:    "Боливия",
			territory.BQ:    "Кариб Нидерландлары",
			territory.BR:    "Бразилия",
			territory.BS:    "Багам утраулары",
			territory.BT:    "Бутан",
			territory.BV:    "Буве утравы",
			territory.BW:    "Ботсвана",
			territory.BY:    "Беларусь",
			territory.BZ:    "Белиз",
			territory.CA:    "Канада",
			territory.CC:    "Кокос (Килинг) утраулары",
			territory.CD:    "Конго (КДР)",
			territory.CF:    "Үзәк Африка Республикасы",
			territory.CG:    "Конго - Браззавиль",
			territory.CH:    "Швейцария",
			territory.CI:    "Кот-д’Ивуар",
			territory.CK:    "Кук утраулары",
			territory.CL:    "Чили",
			territory.CM:    "Камерун",
			territory.CN:    "Кытай",
			territory.CO:    "Колумбия",
			territory.CP:    "Клиппертон утравы",
			territory.CR:    "Коста-Рика",
			territory.CU:    "Куба",
			territory.CV:    "Кабо-Верде",
			territory.CW:    "Кюрасао",
			territory.CX:    "Раштуа утравы",
			territory.CY:    "Кипр",
			territory.CZ:    "Чехия Республикасы",
			territory.DE:    "Германия",
			territory.DG:    "Диего Гарсия",
			territory.DJ:    "Җибүти",
			territory.DK:    "Дания",
			territory.DM:    "Доминика",
			territory.DO:    "Доминикана Республикасы",
			territory.DZ:    "Алжир",
			territory.EA:    "Сеута һәм Мелилья",
			territory.EC:    "Эквадор",
			territory.EE:    "Эстония",
			territory.EG:    "Мисыр",
			territory.EH:    "Көнбатыш Сахара",
			territory.ER:    "Эритрея",
			territory.ES:    "Испания",
			territory.ET:    "Эфиопия",
			territory.EU:    "Европа Берлеге",
			territory.EZ:    "Еврозона",
			territory.FI:    "Финляндия",
			territory.FJ:    "Фиджи",
			territory.FK:    "Фолкленд утраулары",
			territory.FM:    "Микронезия",
			territory.FO:    "Фарер утраулары",
			territory.FR:    "Франция",
			territory.GA:    "Габон",
			territory.GB:    "Берләшкән Корольлек",
			territory.GD:    "Гренада",
			territory.GE:    "Грузия",
			territory.GF:    "Француз Гвианасы",
			territory.GG:    "Гернси",
			territory.GH:    "Гана",
			territory.GI:    "Гибралтар",
			territory.GL:    "Гренландия",
			territory.GM:    "Гамбия",
			territory.GN:    "Гвинея",
			territory.GP:    "Гваделупа",
			territory.GQ:    "Экваториаль Гвинея",
			territory.GR:    "Греция",
			territory.GS:    "Көньяк Георгия һәм Көньяк Сандвич утраулары",
			territory.GT:    "Гватемала",
			territory.GU:    "Гуам",
			territory.GW:    "Гвинея-Бисау",
			territory.GY:    "Гайана",
			territory.HK:    "Гонконг Махсус Идарәле Төбәге",
			territory.HM:    "Херд утравы һәм Макдональд утраулары",
			territory.HN:    "Гондурас",
			territory.HR:    "Хорватия",
			territory.HT:    "Гаити",
			territory.HU:    "Венгрия",
			territory.IC:    "Канар утраулары",
			territory.ID:    "Индонезия",
			territory.IE:    "Ирландия",
			territory.IL:    "Израиль",
			territory.IM:    "Мэн утравы",
			territory.IN:    "Индия",
			territory.IO:    "Британиянең Һинд Океанындагы Территориясе",
			territory.IQ:    "Гыйрак",
			territory.IR:    "Иран",
			territory.IS:    "Исландия",
			territory.IT:    "Италия",
			territory.JE:    "Джерси",
			territory.JM:    "Ямайка",
			territory.JO:    "Иордания",
			territory.JP:    "Япония",
			territory.KE:    "Кения",
			territory.KG:    "Кыргызстан",
			territory.KH:    "Камбоджа",
			territory.KI:    "Кирибати",
			territory.KM:    "Комор утраулары",
			territory.KN:    "Сент-Китс һәм Невис",
			territory.KP:    "Төньяк Корея",
			territory.KR:    "Көньяк Корея",
			territory.KW:    "Күвәйт",
			territory.KY:    "Кайман утраулары",
			territory.KZ:    "Казахстан",
			territory.LA:    "Лаос",
			territory.LB:    "Ливан",
			territory.LC:    "Сент-Люсия",
			territory.LI:    "Лихтенштейн",
			territory.LK:    "Шри-Ланка",
			territory.LR:    "Либерия",
			territory.LS:    "Лесото",
			territory.LT:    "Литва",
			territory.LU:    "Люксембург",
			territory.LV:    "Латвия",
			territory.LY:    "Ливия",
			territory.MA:    "Марокко",
			territory.MC:    "Монако",
			territory.MD:    "Молдова",
			territory.ME:    "Черногория",
			territory.MF:    "Сент-Мартин",
			territory.MG:    "Мадагаскар",
			territory.MH:    "Маршалл утраулары",
			territory.MK:    "Төньяк Македония",
			territory.ML:    "Мали",
			territory.MM:    "Мьянма (Бирма)",
			territory.MN:    "Монголия",
			territory.MO:    "Макао Махсус Идарәле Төбәге",
			territory.MP:    "Төньяк Мариана утраулары",
			territory.MQ:    "Мартиника",
			territory.MR:    "Мавритания",
			territory.MS:    "Монтсеррат",
			territory.MT:    "Мальта",
			territory.MU:    "Маврикий",
			territory.MV:    "Мальдив утраулары",
			territory.MW:    "Малави",
			territory.MX:    "Мексика",
			territory.MY:    "Малайзия",
			territory.MZ:    "Мозамбик",
			territory.NA:    "Намибия",
			territory.NC:    "Яңа Каледония",
			territory.NE:    "Нигер",
			territory.NF:    "Норфолк утравы",
			territory.NG:    "Нигерия",
			territory.NI:    "Никарагуа",
			territory.NL:    "Нидерланд",
			territory.NO:    "Норвегия",
			territory.NP:    "Непал",
			territory.NR:    "Науру",
			territory.NU:    "Ниуэ",
			territory.NZ:    "Яңа Зеландия",
			territory.OM:    "Оман",
			territory.PA:    "Панама",
			territory.PE:    "Перу",
			territory.PF:    "Француз Полинезиясе",
			territory.PG:    "Папуа - Яңа Гвинея",
			territory.PH:    "Филиппин",
			territory.PK:    "Пакистан",
			territory.PL:    "Польша",
			territory.PM:    "Сен-Пьер һәм Микелон",
			territory.PN:    "Питкэрн утраулары",
			territory.PR:    "Пуэрто-Рико",
			territory.PS:    "Фәләстин территорияләре",
			territory.PT:    "Португалия",
			territory.PW:    "Палау",
			territory.PY:    "Парагвай",
			territory.QA:    "Катар",
			territory.QO:    "Ерак Океания",
			territory.RE:    "Реюньон",
			territory.RO:    "Румыния",
			territory.RS:    "Сербия",
			territory.RU:    "Россия",
			territory.RW:    "Руанда",
			territory.SA:    "Согуд Гарәбстаны",
			territory.SB:    "Сөләйман утраулары",
			territory.SC:    "Сейшел утраулары",
			territory.SD:    "Судан",
			territory.SE:    "Швеция",
			territory.SG:    "Сингапур",
			territory.SH:    "Изге Елена утравы",
			territory.SI:    "Словения",
			territory.SJ:    "Шпицберген һәм Ян-Майен",
			territory.SK:    "Словакия",
			territory.SL:    "Сьерра-Леоне",
			territory.SM:    "Сан-Марино",
			territory.SN:    "Сенегал",
			territory.SO:    "Сомали",
			territory.SR:    "Суринам",
			territory.SS:    "Көньяк Судан",
			territory.ST:    "Сан-Томе һәм Принсипи",
			territory.SV:    "Сальвадор",
			territory.SX:    "Синт-Мартен",
			territory.SY:    "Сүрия",
			territory.SZ:    "Свазиленд",
			territory.TA:    "Тристан-да-Кунья",
			territory.TC:    "Теркс һәм Кайкос утраулары",
			territory.TD:    "Чад",
			territory.TF:    "Франциянең Көньяк Территорияләре",
			territory.TG:    "Того",
			territory.TH:    "Тайланд",
			territory.TJ:    "Таҗикстан",
			territory.TK:    "Токелау",
			territory.TL:    "Тимор-Лесте",
			territory.TM:    "Төркмәнстан",
			territory.TN:    "Тунис",
			territory.TO:    "Тонга",
			territory.TR:    "Төркия",
			territory.TT:    "Тринидад һәм Тобаго",
			territory.TV:    "Тувалу",
			territory.TW:    "Тайвань",
			territory.TZ:    "Танзания",
			territory.UA:    "Украина",
			territory.UG:    "Уганда",
			territory.UM:    "АКШ Кече Читтәге утраулары",
			territory.UN:    "Берләшкән Милләтләр",
			territory.US:    "АКШ",
			territory.UY:    "Уругвай",
			territory.UZ:    "Үзбәкстан",
			territory.VA:    "Ватикан",
			territory.VC:    "Сент-Винсент һәм Гренадин",
			territory.VE:    "Венесуэла",
			territory.VG:    "Британия Виргин утраулары",
			territory.VI:    "АКШ Виргин утраулары",
			territory.VN:    "Вьетнам",
			territory.VU:    "Вануату",
			territory.WF:    "Уоллис һәм Футуна",
			territory.WS:    "Самоа",
			territory.XA:    "Псевдоакцентлар",
			territory.XB:    "Псевдо-Биди",
			territory.XK:    "Косово",
			territory.YE:    "Йәмән",
			territory.YT:    "Майотта",
			territory.ZA:    "Көньяк Африка",
			territory.ZM:    "Замбия",
			territory.ZW:    "Зимбабве",
			territory.ZZ:    "билгесез төбәк",
		},
	}
}
