// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_ky() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y-'ж'., d-MMMM, EEEE", Long: "y-'ж'., d-MMMM", Medium: "y-'ж'., d-MMM", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Янв", Feb: "Фев", Mar: "Мар", Apr: "Апр", May: "Май", Jun: "Июн", Jul: "Июл", Aug: "Авг", Sep: "Сен", Oct: "Окт", Nov: "Ноя", Dec: "Дек"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Январь", Feb: "Февраль", Mar: "Март", Apr: "Апрель", May: "Май", Jun: "Июнь", Jul: "Июль", Aug: "Август", Sep: "Сентябрь", Oct: "Октябрь", Nov: "Ноябрь", Dec: "Декабрь"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "жек.", Mon: "дүй.", Tue: "шейш.", Wed: "шарш.", Thu: "бейш.", Fri: "жума", Sat: "ишм."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Ж", Mon: "Д", Tue: "Ш", Wed: "Ш", Thu: "Б", Fri: "Ж", Sat: "И"}, Short: cldr.CalendarDayFormatNameValue{Sun: "жш.", Mon: "дш.", Tue: "шш.", Wed: "шр.", Thu: "бш.", Fri: "жм.", Sat: "иш."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "жекшемби", Mon: "дүйшөмбү", Tue: "шейшемби", Wed: "шаршемби", Thu: "бейшемби", Fri: "жума", Sat: "ишемби"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "тң", PM: "тк"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "тң", PM: "тк"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "таңкы", PM: "түштөн кийинки"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_ky]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Бириккен Араб Эмираттарынын дирхамы", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Афганстан афганиси", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "албан леги", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Армения драмы", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "нидерланд-антил гулдени", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Ангола кванзасы", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "аргентина песосу", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Австралия доллары", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "аруба флорини", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Азербайжан манаты", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "босния-герцоговина конвертациялануучу маркасы", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "барбадос доллары", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Бангладеш такасы", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "болгар левиси", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Бахрейн динары", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Бурунди франкы", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "бермуд доллары", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Бруней доллары", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "боливия боливианосу", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "бразилия реалы", Symbol: "BRL"},
				currency.BSD: cldr.Currency{DisplayName: "багама доллары", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Бутан нгултруму", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Ботсвана пуласы", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "беларусь рублу", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "беларусь рублу (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "белиз доллары", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "канада доллары", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "Конго франкы", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "швейцария франкы", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "чили песосу", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Кытай юаны (офшор)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Кытай юаны", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "колумбия песосу", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "коста-рика колону", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "кубанын конвертациялануучу песосу", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "куба песосу", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Кабо-Верде эскудосу", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "чех кронасы", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Жибути франкы", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "дания крону", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "доминикан песосу", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "алжир динары", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "египет фунту", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Эритреа накфасы", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Эфиопия бирри", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Евро", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Фижи доллары", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "фолкленд аралдарынын фунту", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "британия фунт стерлинги", Symbol: "GBP"},
				currency.GEL: cldr.Currency{DisplayName: "Грузия лариси", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Гана седиси", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "гибралтар фунту", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Гамбия даласиси", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Гине франкы", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "гватемала кетсалы", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "гуйана доллары", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Гонконг доллары", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "гондурас лемпирасы", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "хорват кунасы", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "гаити гурдусу", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "венгр форинти", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Индонезия рупийасы", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Израилдин жаңы шекели", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "Индия руписи", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "Ирак динары", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Иран риалы", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "исландия крону", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "ямайка доллары", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Йордания динары", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Жапан йени", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Кения шиллинги", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Кыргызстан сому", Symbol: "сом"},
				currency.KHR: cldr.Currency{DisplayName: "Камбожа риели", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Коморос франкы", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Түндүк Корея уону", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Түштүк Корея уону", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "Кувейт динары", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "кайман доллары", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Казакстан теңгеси", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Лаос киби", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Ливан фунту", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Шри Ланка руписи", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Либерия доллары", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "литва литасы", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "латвия латы", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "ливия динары", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "марокко дирхамы", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "молдован лейи", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Мадагаскар ариариси", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "македон денары", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Мйанмар кйаты", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Монгол тугриги", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Макау патакасы", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Мавритания угиясы (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Мавритания угиясы", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Мавританий руписи", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Малдив руфийасы", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Малави квачасы", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "мексика песосу", Symbol: "MXN"},
				currency.MYR: cldr.Currency{DisplayName: "Малайзия ринггити", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Мозамбик метикалы", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Намибия доллары", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Нигерия найрасы", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "никарагуа кордобасы", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "норвегия крону", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Непал руписи", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Жаңы Зеландия доллары", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "Оман риалы", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "панама балбоасы", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "перу солу", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Папуа Жаңы Гвинея кинасы", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Филиппин песосу", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Пакистан руписи", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "польша злотыйы", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "парагвай гуараниси", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Катар риалы", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "румын лейи", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "серб динары", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "орус рублу", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Руанда франкы", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Сауд риалы", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Соломон аралдарынын доллары", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Сейшел руписи", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "судан фунту", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "швеция крону", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Сингапур доллары", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Ыйык Елена аралынын фунту", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Сиерра-Леоне леонеси", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Сомали шиллинги", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "суринам доллары", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Түштүк Судан фунту", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Сао Томе жана Принсипе добрасы (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Сао Томе жана Принсипе добрасы", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Сирия фунту", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Свази лилангени", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Тай баты", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Тажикстан сомониси", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Түркмөнстан манаты", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "тунис динары", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Тонга паангасы", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Түркия лирасы", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "тринидад жана тобаго доллары", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Тайвань жаңы доллары", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "Танзания шиллинги", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "украин гривени", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Уганда шиллинги", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "АКШ доллары", Symbol: "USD"},
				currency.UYU: cldr.Currency{DisplayName: "уругвай песосу", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Өзбекстан суму", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "венесуэла боливары (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Венесуэла боливары", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Вьетнам доңу", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Вануату ватусу", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Самоа таласы", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Борбордук Африка КФА франкы", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "чыгыш кариб доллары", Symbol: "XCD"},
				currency.XOF: cldr.Currency{DisplayName: "КФА франкы", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "КФП франкы", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Белгисиз акча", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Йемен риалы", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Түштүк Африка ранды", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Замбия квачасы", Symbol: "ZMW"},
			},
		},
	}
}
