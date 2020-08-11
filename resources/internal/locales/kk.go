// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_kk() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "kk",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y 'ж'. d MMMM, EEEE", Long: "y 'ж'. d MMMM", Medium: "y 'ж'. dd MMM", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "қаң.", Feb: "ақп.", Mar: "нау.", Apr: "сәу.", May: "мам.", Jun: "мау.", Jul: "шіл.", Aug: "там.", Sep: "қыр.", Oct: "қаз.", Nov: "қар.", Dec: "жел."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Қ", Feb: "А", Mar: "Н", Apr: "С", May: "М", Jun: "М", Jul: "Ш", Aug: "Т", Sep: "Қ", Oct: "Қ", Nov: "Қ", Dec: "Ж"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Қаңтар", Feb: "Ақпан", Mar: "Наурыз", Apr: "Сәуір", May: "Мамыр", Jun: "Маусым", Jul: "Шілде", Aug: "Тамыз", Sep: "Қыркүйек", Oct: "Қазан", Nov: "Қараша", Dec: "Желтоқсан"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "жс", Mon: "дс", Tue: "сс", Wed: "ср", Thu: "бс", Fri: "жм", Sat: "сб"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Ж", Mon: "Д", Tue: "С", Wed: "С", Thu: "Б", Fri: "Ж", Sat: "С"}, Short: cldr.CalendarDayFormatNameValue{Sun: "жс", Mon: "дс", Tue: "сс", Wed: "ср", Thu: "бс", Fri: "жм", Sat: "сб"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "жексенбі", Mon: "дүйсенбі", Tue: "сейсенбі", Wed: "сәрсенбі", Thu: "бейсенбі", Fri: "жұма", Sat: "сенбі"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Біріккен Араб Әмірліктерінің дирхамы", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Ауғанстан афганиі", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Албания легі", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Армения драмы", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Нидерланд антиль гульдені", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Ангола кванзасы", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Аргентина песосы", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Австралия доллары", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Аруба флорині", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Азербайджан манаты (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Әзірбайжан манаты", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Босния және Герцеговина айырбасталмалы маркасы", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Барбадос доллары", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Бангладеш такасы", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Болгария леві", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Бахрейн динары", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Бурунди франкі", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Бермуд доллары", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Бруней доллары", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Боливия боливианосы", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Бразилия реалы", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Багам доллары", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Бутан нгултрумы", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Ботсвана пуласы", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Беларусь рублі", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Беларусь рублі (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Белиз доллары", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Канада доллары", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Конго франкі", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Швейцария франкі", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Чили песосы", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Қытай юані (офшор)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Қытай юані", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Колумбия песосы", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Коста-Рика колоны", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Куба айырбасталмалы песосы", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Куба песосы", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Кабо-Верде эскудосы", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Чехия кронасы", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Джибути франкі", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Дат кроны", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Доминикан песосы", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Алжир динары", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Мысыр фунты", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Эритрея накфасы", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Эфиопия быры", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Еуро", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Фиджи доллары", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Фолкленд аралдарының фунты", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Британдық фунт", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Грузия лариі", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Гана седиі", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Гибралтар фунты", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Гамбия даласиі", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Гвинея франкі", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Гватемала кетсалі", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Гайана доллары", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Гонконг доллары", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Гондурас лемпирасы", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Хорватия кунасы", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Гаити гурды", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Венгрия форинті", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Индонезия рупиясы", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Израиль жаңа шекелі", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Үндістан рупиясы", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Ирак динары", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Иран риалы", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Исландия кронасы", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Ямайка доллары", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Иордания динары", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Жапония иенасы", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Кения шиллингі", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Қырғызстан сомы", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Камбоджа риелі", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Комор аралдары франкі", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Солтүстік Корея воны", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Оңтүстік Корея воны", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Кувейт динары", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Кайман аралдары доллары", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Қазақстан теңгесі", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Лаос кипі", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Ливан фунты", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Шри-Ланка рупиясы", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Либерия доллары", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "Литва литы", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "Латвия латы", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Ливия динары", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Марокко дирхамы", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Молдова лейі", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Малагаси ариариі", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Македония динары", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Мьянма кьяты", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Моңғолия тугригі", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Макао патакасы", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Мавритания угиясы (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Мавритания угиясы", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Маврикий рупиясы", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Мальдив руфиясы", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Малави квачасы", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Мексика песосы", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Малайзия ринггиті", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Мозамбик метикалы", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Намибия доллары", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Нигерия найрасы", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Никарагуа кордобасы", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Норвегия кроны", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Непал рупиясы", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Жаңа Зеландия доллары", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Оман риалы", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Панама бальбоасы", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Перу солі", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Папуа - Жаңа Гвинея кинасы", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Филиппин песосы", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Пәкістан рупиясы", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Польша злотасы", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Парагвай гуараниі", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Катар риалы", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Румыния лейі", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Сербия динары", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Ресей рублі", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Руанда франкі", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Сауд Арабиясының риалы", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Соломон аралдары доллары", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Сейшель рупиясы", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Судан фунты", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Швеция кроны", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Сингапур доллары", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Әулие Елена аралы фунты", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Сьерра-Леоне леонесі", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Сомали шиллингі", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Суринам доллары", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Оңтүстік Судан фунты", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Сант-Томе мен Принсипи добрасы (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Сант-Томе мен Принсипи добрасы", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Сирия фунты", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Свазиленд лилангениі", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Тай баты", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Тәжікстан сомониі", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Түрікменстан манаты", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Тунис динары", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Тонга паангасы", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Түрік лирасы", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Тринидад және Тобаго доллары", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Жаңа Тайван доллары", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Танзания шиллингі", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Украина гривнасы", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Уганда шиллингі", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "АҚШ доллары", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Уругвай песосы", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Өзбекстан сомы", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Венесуэла боливары (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Венесуэла боливары", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Вьетнам донгі", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Вануату ватуы", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Самоа таласы", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "КФА ВЕАС франкі", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Шығыс Кариб доллары", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "КФА ВСЕАО франкі", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "КФП франкі", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Белгісіз валюта", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Йемен риалы", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Оңтүстік Африка рэнді", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Замбия квачасы", Symbol: "ZMW"},
			},
		},
	}
}
