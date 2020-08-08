// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_mn_MN() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y 'оны' MMMM'ын' d, EEEE 'гараг'", Long: "y 'оны' MMMM'ын' d", Medium: "y 'оны' MMM'ын' d", Short: "y.MM.dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss (zzzz)", Long: "HH:mm:ss (z)", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "1-р сар", Feb: "2-р сар", Mar: "3-р сар", Apr: "4-р сар", May: "5-р сар", Jun: "6-р сар", Jul: "7-р сар", Aug: "8-р сар", Sep: "9-р сар", Oct: "10-р сар", Nov: "11-р сар", Dec: "12-р сар"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "I", Feb: "II", Mar: "III", Apr: "IV", May: "V", Jun: "VI", Jul: "VII", Aug: "VIII", Sep: "IX", Oct: "X", Nov: "XI", Dec: "XII"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Нэгдүгээр сар", Feb: "Хоёрдугаар сар", Mar: "Гуравдугаар сар", Apr: "Дөрөвдүгээр сар", May: "Тавдугаар сар", Jun: "Зургаадугаар сар", Jul: "Долоодугаар сар", Aug: "Наймдугаар сар", Sep: "Есдүгээр сар", Oct: "Аравдугаар сар", Nov: "Арван нэгдүгээр сар", Dec: "Арван хоёрдугаар сар"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Ня", Mon: "Да", Tue: "Мя", Wed: "Лх", Thu: "Пү", Fri: "Ба", Sat: "Бя"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Ня", Mon: "Да", Tue: "Мя", Wed: "Лх", Thu: "Пү", Fri: "Ба", Sat: "Бя"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Ня", Mon: "Да", Tue: "Мя", Wed: "Лх", Thu: "Пү", Fri: "Ба", Sat: "Бя"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Ням", Mon: "Даваа", Tue: "Мягмар", Wed: "Лхагва", Thu: "Пүрэв", Fri: "Баасан", Sat: "Бямба"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ү.ө.", PM: "ү.х."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ү.ө.", PM: "ү.х."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ү.ө.", PM: "ү.х."}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_mn]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "арабын нэгдсэн эмиратын дирхам", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Афганистаны афгани", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Албанийн лек", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Арменийн драм", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Нидерландын Антиллийн гулдер", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Анголын кванза", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Аргентины песо", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Австралийн доллар", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Арубын флорин", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Азербайжаны манат", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Босни-Герцеговины хөрвөгч марк", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "барбадос доллар", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Бангладешийн така", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Болгарын лев", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Бахрейн динар", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Бурундийн франк", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Бермудын доллар", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Брунейн доллар", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Боливийн боливиано", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Бразилийн реал", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Багамын доллар", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Бутаны нгултрум", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Ботсванийн пула", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Беларусийн рубль", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "беларусь рубль (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Белизийн доллар", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "канад доллар", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Конгогийн франк", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Швейцарийн франк", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Чилийн песо", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Хятадын юань (офшор)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Хятадын юань", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Колумбын песо", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Коста-Рикагийн колон", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Кубын хөрвөгч песо", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Кубын песо", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Кабо-Вердийн эскудо", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Чехийн крон", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Жибутийн франк", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Данийн крон", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Доминиканы песо", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Алжирийн доллар", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Египетийн фунт", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Эритрейн накфа", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Этиопын бирр", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "евро", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Фижигийн доллар", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Фолклендийн арлуудын паунд", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Британийн фунт", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Гүржийн лари", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Ганагийн седи", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Гибралтарын фунт", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Гамбийн даласи", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Гвинейн франк", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Гватемалын кецал", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Гайанын доллар", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Хонг Конгийн доллар", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Гондурасын лемпира", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Хорватын куна", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Гаитийн гурд", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Унгарын форинт", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Индонезийн рупи", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Израилийн шинэ шекел", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Энэтхэгийн рупи", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Иракийн динар", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Ираны риял", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Исландын крон", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Ямайкийн доллар", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Йорданы динар", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Японы иен", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Кенийн шиллинг", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Кыргызын сом", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Камбожийн риел", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Коморын франк", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Хойд Солонгосын вон", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Өмнөд Солонгосын вон", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Кувейтийн динар", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Кайманы арлуудын доллар", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Казахстаны тэнгэ", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Лаосын кип", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Ливаны фунт", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Шри-Ланкийн рупи", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Либерийн доллар", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "литвийн литас", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "латвийн лац", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "Ливийн доллар", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Мороккогийн дирхэм", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Молдовын лей", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Малагасийн ариари", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Македонийн динар", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Мьянмарын киат", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Монгол төгрөг", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Макаогийн патака", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "мавритан угия (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Мавританийн угия", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Маврикийн рупи", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Мальдивийн руфия", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Малавийн квача", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Мексикийн песо", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Малайзын рингит", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Мозамбикийн метикал", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Намибийн доллар", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Нигерийн найра", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Никарагуагийн кордоба", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Норвегийн крон", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Балбын рупи", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Шинэ Зеландын доллар", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Оманийн риал", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Панамын бальбоа", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Перугийн соль", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Папуа-Шинэ Гвинейн кина", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Филиппиний песо", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Пакистаны рупи", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Польшийн злот", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Парагвайн гуарани", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Катарын риал", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Румыны лей", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Сербийн динар", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Оросын рубль", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Руандагийн франк", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Саудын риял", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Соломоны арлуудын доллар", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Сейшелийн рупи", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Суданы фунт", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Шведийн крон", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Сингапурын доллар", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Сент Хеленагийн фунт", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Сьерра-Леоны леон", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Сомалийн шиллинг", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Суринамын доллар", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Өмнөд Суданы фунт", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "сан-томе ба принсипи добра (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Сан-Томе ба Принсипигийн добра", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Сирийн фунт", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Свазиландын лилангени", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Тайландын бат", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Тажикийн сомон", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Туркмены манат", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Тунисын доллар", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Тонгагийн панга", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Туркийн лира", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Тринидад ба Тобагогийн доллар", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Шинэ Тайванийн доллар", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Танзанийн шиллинг", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Украины гривна", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Угандагийн шиллинг", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "америк доллар", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "Уругвайн песо", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Узбекийн сом", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "венесуэлийн боливар (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Венесуэлийн боливар", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Вьетнамын донг", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Вануатугийн вату", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Самоагийн тала", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Төв Африкийн франк", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Зүүн Карибийн доллар", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Баруун Африкийн франк", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Францын колонийн франк", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "тодорхойгүй мөнгөн тэмдэгт", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Йемений риял", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Өмнөд Африкийн ранд", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Замби квача (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Замбийн квача", Symbol: "ZMW"},
			},
		},
	}
}