// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_be_BY() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "be_BY",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d.MM.y", Short: "d.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss, zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'у' {0}", Long: "{1} 'у' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "сту", Feb: "лют", Mar: "сак", Apr: "кра", May: "май", Jun: "чэр", Jul: "ліп", Aug: "жні", Sep: "вер", Oct: "кас", Nov: "ліс", Dec: "сне"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "с", Feb: "л", Mar: "с", Apr: "к", May: "м", Jun: "ч", Jul: "л", Aug: "ж", Sep: "в", Oct: "к", Nov: "л", Dec: "с"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "студзень", Feb: "люты", Mar: "сакавік", Apr: "красавік", May: "май", Jun: "чэрвень", Jul: "ліпень", Aug: "жнівень", Sep: "верасень", Oct: "кастрычнік", Nov: "лістапад", Dec: "снежань"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "аў", Wed: "ср", Thu: "чц", Fri: "пт", Sat: "сб"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "н", Mon: "п", Tue: "а", Wed: "с", Thu: "ч", Fri: "п", Sat: "с"}, Short: cldr.CalendarDayFormatNameValue{Sun: "нд", Mon: "пн", Tue: "аў", Wed: "ср", Thu: "чц", Fri: "пт", Sat: "сб"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "нядзеля", Mon: "панядзелак", Tue: "аўторак", Wed: "серада", Thu: "чацвер", Fri: "пятніца", Sat: "субота"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "дырхам ААЭ", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "афганскі афгані", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "албанскі лек", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "армянскі драм", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "нідэрландскі антыльскі гульдэн", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "ангольская кванза", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "аргенцінскае песа", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "аўстралійскі долар", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "арубанскі фларын", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "азербайджанскі манат", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "канверсоўная марка Босніі і Герцагавіны", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "барбадаскі долар", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "бангладэшская така", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "балгарскі леў", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "бахрэйнскі дынар", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "бурундзійскі франк", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "бермудскі долар", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "брунейскі долар", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "балівіяна", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "бразільскі рэал", Symbol: "BRL"},
				currency.BSD: cldr.Currency{DisplayName: "багамскі долар", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "бутанскі нгултрум", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "батсванская пула", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "беларускі рубель", Symbol: "Br"},
				currency.BYR: cldr.Currency{DisplayName: "беларускі рубель (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "белізскі долар", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "канадскі долар", Symbol: "CAD"},
				currency.CDF: cldr.Currency{DisplayName: "кангалезскі франк", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "швейцарскі франк", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "чылійскае песа", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "афшорны кітайскі юань", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "кітайскі юань", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "калумбійскае песа", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "коста-рыканскі калон", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "кубінскае канверсоўнае песа", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "кубінскае песа", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "эскуда Каба-Вердэ", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "чэшская крона", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "джыбуційскі франк", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "дацкая крона", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "дамініканскае песа", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "алжырскі дынар", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "егіпецкі фунт", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "эрытрэйская накфа", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "эфіопскі быр", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "еўра", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "фіджыйскі долар", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "фунт Фалклендскіх астравоў", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "брытанскі фунт стэрлінгаў", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "грузінскі лары", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "ганскі седзі", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "гібралтарскі фунт", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "гамбійскі даласі", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "гвінейскі франк", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "гватэмальскі кетсаль", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "гаянскі долар", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "ганконгскі долар", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "гандураская лемпіра", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "харвацкая куна", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "гаіцянскі гурд", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "венгерскі форынт", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "інданезійская рупія", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "новы ізраільскі шэкель", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "індыйская рупія", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "іракскі дынар", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "іранскі рыял", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "ісландская крона", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "ямайскі долар", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "іарданскі дынар", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "японская іена", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "кенійскі шылінг", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "кіргізскі сом", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "камбаджыйскі рыель", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "каморскі франк", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "паўночнакарэйская вона", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "паўднёвакарэйская вона", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "кувейцкі дынар", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "долар Кайманавых астравоў", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "казахстанскі тэнге", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "лаоскі кіп", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "ліванскі фунт", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "шры-ланкійская рупія", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "ліберыйскі долар", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "лівійскі дынар", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "мараканскі дырхам", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "малдаўскі лей", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "малагасійскі арыяры", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "македонскі дэнар", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "м’янманскі к’ят", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "мангольскі тугрык", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "патака Макаа", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "маўрытанская ўгія (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "маўрытанская угія", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "маўрыкійская рупія", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "мальдыўская руфія", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "малавійская квача", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "мексіканскае песа", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "малайзійскі рынгіт", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "мазамбікскі метыкал", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "намібійскі долар", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "нігерыйская наіра", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "нікарагуанская кордаба", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "нарвежская крона", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "непальская рупія", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "новазеландскі долар", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "аманскі рыял", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "панамскае бальбоа", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "перуанскі соль", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "кіна Папуа-Новай Гвінеі", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "філіпінскае песа", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "пакістанская рупія", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "польскі злоты", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "парагвайскі гуарані", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "катарскі рыял", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "румынскі лей", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "сербскі дынар", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "расійскі рубель", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "руандыйскі франк", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "саудаўскі рыял", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "долар Саламонавых астравоў", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "сейшэльская рупія", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "суданскі фунт", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "шведская крона", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "сінгапурскі долар", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "фунт в-ва Святой Алены", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "сьера-леонскі леонэ", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "самалійскі шылінг", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "сурынамскі долар", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "паўднёвасуданскі фунт", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "добра Сан-Тамэ і Прынсіпі (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "добра Сан-Тамэ і Прынсіпі", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "сірыйскі фунт", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "свазілендскі лілангені", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "тайскі бат", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "таджыкскі самані", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "туркменскі манат", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "туніскі дынар", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "танганская паанга", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "турэцкая ліра", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "долар Трынідада і Табага", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "новы тайваньскі долар", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "танзанійскі шылінг", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "украінская грыўня", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "угандыйскі шылінг", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "долар ЗША", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "уругвайскае песа", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "узбекскі сум", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "венесуальскі балівар (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "венесуэльскі балівар", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "в’етнамскі донг", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "вануацкі вату", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "самаанская тала", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "цэнтральнаафрыканскі франк КФА", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "усходнекарыбскі долар", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "заходнеафрыканскі франк КФА", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "французскі ціхаакіянскі франк", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "невядомая валюта", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "еменскі рыал", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "паўднёваафрыканскі рэнд", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "замбійская квача", Symbol: "ZMW"},
			},
		},
	}
}
