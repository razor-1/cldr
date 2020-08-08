// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_uz_Cyrl() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, dd MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss (zzzz)", Long: "HH:mm:ss (z)", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "янв", Feb: "фев", Mar: "мар", Apr: "апр", May: "май", Jun: "июн", Jul: "июл", Aug: "авг", Sep: "сен", Oct: "окт", Nov: "ноя", Dec: "дек"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Я", Feb: "Ф", Mar: "М", Apr: "А", May: "М", Jun: "И", Jul: "И", Aug: "А", Sep: "С", Oct: "О", Nov: "Н", Dec: "Д"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "январ", Feb: "феврал", Mar: "март", Apr: "апрел", May: "май", Jun: "июн", Jul: "июл", Aug: "август", Sep: "сентябр", Oct: "октябр", Nov: "ноябр", Dec: "декабр"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "якш", Mon: "душ", Tue: "сеш", Wed: "чор", Thu: "пай", Fri: "жум", Sat: "шан"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Я", Mon: "Д", Tue: "С", Wed: "Ч", Thu: "П", Fri: "Ж", Sat: "Ш"}, Short: cldr.CalendarDayFormatNameValue{Sun: "як", Mon: "ду", Tue: "се", Wed: "чо", Thu: "па", Fri: "жу", Sat: "ша"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "якшанба", Mon: "душанба", Tue: "сешанба", Wed: "чоршанба", Thu: "пайшанба", Fri: "жума", Sat: "шанба"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ТО", PM: "ТК"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ТО", PM: "ТК"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ТО", PM: "ТК"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_root]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "-", Percent: "٪", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ANG: cldr.Currency{DisplayName: "Голланд Антил гульдени", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Аргентина песоси", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Аруба флорини", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Барбадос доллари", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "Бермуда доллари", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Боливия болвиани", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Бразил реали", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Багама доллари", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "Белиз доллари", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Канада доллари", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "Чили песоси", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Хитой юани", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Колумбия песоси", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Коста-Рика колони", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Куба Айирбошлаш песоси", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Куба песоси", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Доминикан песоси", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Жазоир динори", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Миср фунти", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "Евро", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Фолькленд ороли фунти", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Инглиз фунт", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Гватемала кветзали", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Гаяна доллари", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Гондурас лемпираси", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Гаити гурдаси", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Ҳинд рупияси", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Ямайка доллари", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Япон йенаси", Symbol: "JP¥"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "Кайман ороли Доллари", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Ливия динори", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Марокаш дирҳами", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MXN: cldr.Currency{DisplayName: "Мексика песоси", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Никарагуа кордобаси", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "", Symbol: "NZ$"},
				currency.PAB: cldr.Currency{DisplayName: "Панама бальбоаси", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Перу сол", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Парагвай гуарани", Symbol: "₲"},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "RON"},
				currency.RUB: cldr.Currency{DisplayName: "Рус рубли", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "Суринам доллари", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Тринидад ва Тобаго доллари", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.USD: cldr.Currency{DisplayName: "АҚШ доллари", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Уругвай песоси", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Ўзбекистон сўм", Symbol: "сўм"},
				currency.VEF: cldr.Currency{DisplayName: "Венесуэла боливари (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Венесуэла боливари", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Шарқий Кариб доллари", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
	}
}
