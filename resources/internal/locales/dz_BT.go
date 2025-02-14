// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_dz_BT() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "dz_BT",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, སྤྱི་ལོ་y MMMM ཚེས་dd", Long: "སྤྱི་ལོ་y MMMM ཚེས་ dd", Medium: "སྤྱི་ལོ་y ཟླ་MMM ཚེས་dd", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "ཆུ་ཚོད་ h སྐར་མ་ mm:ss a zzzz", Long: "ཆུ་ཚོད་ h སྐར་མ་ mm:ss a z", Medium: "ཆུ་ཚོད་h:mm:ss a", Short: "ཆུ་ཚོད་ h སྐར་མ་ mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "ཇི་ཨེམ་ཏི་{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ཟླ་༡", Feb: "ཟླ་༢", Mar: "ཟླ་༣", Apr: "ཟླ་༤", May: "ཟླ་༥", Jun: "ཟླ་༦", Jul: "ཟླ་༧", Aug: "ཟླ་༨", Sep: "ཟླ་༩", Oct: "ཟླ་༡༠", Nov: "ཟླ་༡༡", Dec: "ཟླ་༡༢"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "༡", Feb: "༢", Mar: "༣", Apr: "༤", May: "༥", Jun: "༦", Jul: "༧", Aug: "༨", Sep: "༩", Oct: "༡༠", Nov: "༡༡", Dec: "༡༢"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "སྤྱི་ཟླ་དངཔ་", Feb: "སྤྱི་ཟླ་གཉིས་པ་", Mar: "སྤྱི་ཟླ་གསུམ་པ་", Apr: "སྤྱི་ཟླ་བཞི་པ", May: "སྤྱི་ཟླ་ལྔ་པ་", Jun: "སྤྱི་ཟླ་དྲུག་པ", Jul: "སྤྱི་ཟླ་བདུན་པ་", Aug: "སྤྱི་ཟླ་བརྒྱད་པ་", Sep: "སྤྱི་ཟླ་དགུ་པ་", Oct: "སྤྱི་ཟླ་བཅུ་པ་", Nov: "སྤྱི་ཟླ་བཅུ་གཅིག་པ་", Dec: "སྤྱི་ཟླ་བཅུ་གཉིས་པ་"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ཟླ་", Mon: "མིར་", Tue: "ལྷག་", Wed: "ཕུར་", Thu: "སངས་", Fri: "སྤེན་", Sat: "ཉི་"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ཟླ", Mon: "མིར", Tue: "ལྷག", Wed: "ཕུར", Thu: "སངྶ", Fri: "སྤེན", Sat: "ཉི"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "གཟའ་ཟླ་བ་", Mon: "གཟའ་མིག་དམར་", Tue: "གཟའ་ལྷག་པ་", Wed: "གཟའ་ཕུར་བུ་", Thu: "གཟའ་པ་སངས་", Fri: "གཟའ་སྤེན་པ་", Sat: "གཟའ་ཉི་མ་"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "སྔ་ཆ་", PM: "ཕྱི་ཆ་"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤#,##,##0.00", CurrencyAccounting: "", Percent: "#,##,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "ཡུ་ནཱའི་ཊེཌ་ ཨ་རབ་ ཨེ་མེ་རེཊས་ཀྱི་དངུལ་ ཌིར་ཧཱམ", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "ཨཕ་གཱན་གྱི་དངུལ་ ཨཕ་ག་ནི", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "ཨཱོས་ཊྲེ་ལི་ཡ་གི་དངུལ་ ཌོ་ལར", Symbol: "AU$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "བྷང་ལ་དེཤ་གི་དངུལ་ ཏ་ཀ", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "བར་མུ་ཌ་གི་དངུལ་ ཌོ་ལར", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "བྲ་ཛིལ་གྱི་དངུལ་ རེ་ཡལ", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "དངུལ་ཀྲམ", Symbol: "Nu."},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "ཀེ་ན་ཌ་གི་དངུལ་ ཌོ་ལར", Symbol: "CA$"},
				currency.CHF: cldr.Currency{DisplayName: "སུ་ཡིས་ཀྱི་དངུལ་ ཕྲངཀ", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "ཅི་ལི་གི་དངུལ་ པེ་སོ", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "རྒྱ་ནག་གི་དངུལ་ ཡུ་ཝཱན", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "ཀོ་ལོམ་བྷི་ཡ་གི་དངུལ་ པེ་སོ", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "ཀིའུ་བྷ་གི་དངུལ་ པེ་སོ", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "ཌེན་མཱཀ་གི་དངུལ་ ཀྲོན", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "ཨཱལ་ཇི་རི་ཡ་གི་དངུལ་ ཌའི་ནར", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "ཨི་ཇིབཊ་གི་དངུལ་ པ་འུནཌ", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.EUR: cldr.Currency{DisplayName: "ཡུ་རོ༌དངུལ་", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "བྲི་ཊིཤ་ པ་འུནཌ་ ཨིས་ཊར་ལིང", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "ཧོང་ཀོང་གི་དངུལ་ ཌོ་ལར", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "ཨིན་ཌོ་ནེ་ཤི་ཡ་གི་དངུལ་ རུ་པི་ཡ", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "ཨིས་རེལ་གྱི་དངུལ་གསརཔ་ ཤེ་ཀེལ", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "རྒྱ་གར་གྱི་དངུལ་ རུ་པི", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ཨི་རཱཀ་གི་དངུལ་ ཌི་ན", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "ཨི་རཱན་གྱི་དངུལ་ རི་ཨཱལ", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "ཨཱཡིས་ལེནཌ་གི་དངུལ་ ཀྲོ་ན", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "ཇཱ་མཻ་ཀ་གི་དངུལ་ ཌོ་ལར", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "ཇོར་ཌན་གྱི་དངུལ་ ཌི་ན", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "ཇཱ་པཱན་གྱི་དངུལ་ ཡེན", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ཀེན་ཡ་གི་དངུལ་ ཤི་ལིང", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "ཀེམ་བྷོ་ཌི་ཡ་གི་དངུལ་ རི་ཨཱལ", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "ནོརཐ་ ཀོ་རི་ཡ་གི་དངུལ་ ཝོན", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "སཱའུཐ་ ཀོ་རི་ཡ་གི་དངུལ་ ཝོན", Symbol: "KR₩"},
				currency.KWD: cldr.Currency{DisplayName: "ཀུ་ཝེཊ་གི་དངུལ་ ཌི་ན", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "ཀ་ཛགས་ཏཱན་གྱི་དངུལ་ ཏེང་གེ", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "ལཱ་ཝོས་ཀྱི་དངུལ་ ཀིཔ", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "ལེ་བ་ནོན་གྱི་དངུལ་ པ་འུནཌ", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "ཤྲི་ ལང་ཀ་གི་དངུལ་ རུ་པི", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "ལཱའི་བེ་རི་ཡ་གི་དངུལ་ ཌོ་ལར", Symbol: "$"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "ལི་བི་ཡ་གི་དངུལ་ ཌི་ན", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "མོ་རོ་ཀོ་གི་དངུལ་ ཌིར་ཧཱམ", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "མི་ཡཱན་མར་གྱི་དངུལ་ ཅཱཏ", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "སོག་པོའི་དངུལ་ ཏུ་གྲིཀ", Symbol: "₮"},
				currency.MUR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "མཱལ་དིབས་ཀྱི་དངུལ་ རུ་ཕི་ཡ", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "མེཀ་སི་ཀོ་གི་དངུལ་ པེ་སོ", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "མ་ལེ་ཤི་ཡ་གི་དངུལ་ རིང་གིཊ", Symbol: "RM"},
				currency.NAD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "ནོར་ཝེ་གི་དངུལ་ ཀྲོ་ན", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "བལ་པོའི་དངུལ་ རུ་པི", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "ནིའུ་ཛི་ལེནཌ་གི་དངུལ་ ཌོ་ལར", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ཨོ་མཱན་གྱི་དངུལ་ རི་ཨཱལ", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "པ་ན་མ་གི་དངུལ་ བཱལ་བོ་ཝ", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "པ་རུ་གི་དངུལ་ ནུ་བོ་ སཱོལ", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "ཕི་ལི་པིནས་གྱི་དངུལ་ པེ་སོ", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "པ་ཀིས་ཏཱན་གྱི་དངུལ་ རུ་པི", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "པོ་ལེནཌ་ཀྱི་དངུལ ཛ྄ལོ་ཊི", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "ཀ་ཊར་གྱི་དངུལ་ རི་ཨཱལ", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "", Symbol: "lei"},
				currency.RUB: cldr.Currency{DisplayName: "ཨུ་རུ་སུ་གི་དངུལ་ རུ་བཱལ", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "སཱཝ་དིའི་དངུལ་ རི་ཡཱལ", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "སེ་ཤཱལས་ཀྱི་དངུལ་ རུ་པི", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "སུའི་ཌེན་གྱི་དངུལ་ ཀྲོ་ན", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "སིང་ག་པོར་གྱི་དངུལ་ ཌོ་ལར", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "སི་རི་ཡ་གི་དངུལ་ པ་འུནཌ", Symbol: "£"},
				currency.THB: cldr.Currency{DisplayName: "ཐཱའི་ལེནཌ་གི་དངུལ་ བཱཏ", Symbol: "TH฿"},
				currency.TJS: cldr.Currency{DisplayName: "ཏ་ཇི་ཀིས་ཏཱན་གྱི་དངུལ་ སོ་མོ་ནི", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "ཊཱར་ཀི་གི་དངུལ་ ལི་ར", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "ཊཱའི་ཝཱན་གི་དངུལ ཌོ་ལར", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "ཊཱན་ཛཱ་ནི་ཡ་གི་དངུལ་ ཤི་ལིང", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "ཡུ་གྷེན་ཌ་གི་དངུལ་ ཤི་ལིང", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "ཡུ་ཨེས་ ཌོ་ལར", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "ཡུ་རུ་གུ་ཝའི་གི་དངུལ་ པེ་སོ", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "ཨུས་བེ་ཀིས་ཏཱན་གྱི་དངུལ་ སོམ", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "བེ་ནི་ཛུ་ཝེ་ལ་གི་དངུལ་ བོ་ལི་བར (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "བེ་ནི་ཛུ་ཝེ་ལ་གི་དངུལ་ བོ་ལི་བར", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "བེཊ་ནཱམ་གྱི་དངུལ་ ཌོང", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "XAF"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "མ་ཤེས་པའི་དངུལ", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "སཱའུཐ་ ཨཕ་རི་ཀ་གི་དངུལ་ རཱནད", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0}། ({1}།)", Separator: "{0}་, {1}", KeyTypePattern: "{0}།: {1}"},
		Languages: cldr.Languages{
			language.AA:      "ཨ་ཕར་ཁ",
			language.AB:      "ཨཱབ་ཁ་ཟི་ཡ་ཁ",
			language.AF:      "ཨཕ་རི་ཀཱནས་ཁ",
			language.AM:      "ཨམ་ཧ་རིཀ་ཁ",
			language.AR:      "ཨེ་ར་བིཀ་ཁ",
			language.AS:      "ཨ་ས་མིས་ཁ",
			language.AZ:      "ཨ་ཛར་བྷའི་ཇཱན་ཁ",
			language.BE:      "བེལ་ཨ་རུས་ཁ",
			language.BG:      "བཱལ་གེ་རི་ཡཱན་ཁ",
			language.BN:      "བངྒ་ལ་ཁ",
			language.BO:      "བོད་ཁ",
			language.BS:      "བྷོས་ནི་ཡཱན་ཁ",
			language.CA:      "ཀེ་ཊ་ལཱན་ཁ",
			language.CS:      "ཅེཀ་ཁ",
			language.CY:      "ཝེལཤ་ཁ",
			language.DA:      "ཌེ་ནིཤ་ཁ",
			language.DAK:     "ད་ཀོ་ཏ་ཁ",
			language.DE:      "ཇཱར་མཱན་ཁ",
			language.DE_AT:   "ཨཱོས་ཊྲི་ཡཱན་ཇཱར་མཱན་ཁ",
			language.DE_CH:   "སུ་ཡིས་གི་མཐོ་སའི་ཇཱར་མཱན་ཁ",
			language.DV:      "དི་བེ་ཧི་ཁ",
			language.DZ:      "རྫོང་ཁ",
			language.EGY:     "ཨི་ཇིཔ་ཤཱན (སྔ་དུས་ཀྱི)",
			language.EL:      "གྲིཀ་ཁ",
			language.EN:      "ཨིང་ལིཤ་ཁ",
			language.EN_AU:   "ཨཱོས་ཊྲེ་ལི་ཡཱན་ཨིང་ལིཤ་ཁ",
			language.EN_CA:   "ཀེ་ན་ཌི་ཡཱན་ཨིང་ལིཤ་ཁ",
			language.EN_GB:   "བྲི་ཊིཤ་ཨིང་ལིཤ་ཁ",
			language.EN_US:   "ཡུ་ཨེས་ཨིང་ལིཤ་ཁ",
			language.EO:      "ཨེས་པ་རཱན་ཏོ་ཁ",
			language.ES:      "ཨིས་པེ་ནིཤ་ཁ",
			language.ES_419:  "ལེ་ཊིན་ཨ་མེ་རི་ཀཱན་གི་ཨིས་པེ་ནིཤ་ཁ",
			language.ES_ES:   "ཡུ་རོབ་ཀྱི་ཨིས་པེ་ནིཤ་ཁ",
			language.ET:      "ཨེས་ཊོ་ནི་ཡཱན་ཁ",
			language.EU:      "བཱསཀ་ཁ",
			language.FA:      "པར་ཤི་ཡཱན་ཁ",
			language.FI:      "ཕི་ནིཤ་ཁ",
			language.FIL:     "ཕི་ལི་པི་ནོ་ཁ",
			language.FJ:      "ཕི་ཇི་ཡཱན་ཁ",
			language.FO:      "ཕཱ་རོ་ཨིས་ཁ",
			language.FR:      "ཕྲནཅ་ཁ",
			language.FR_CA:   "ཀེ་ན་ཌི་ཡཱན་ཕྲནཅ་ཁ",
			language.FR_CH:   "སུ་ཡིས་ཕྲནཅ་ཁ",
			language.FY:      "ནུབ་ཕྼི་སི་ཡན་ཁ",
			language.GA:      "ཨཱའི་རིཤ་ཁ",
			language.GL:      "གལ་ཨིས་ཨི་ཡན་ཁ",
			language.GN:      "གུ་ཝ་ར་ནི་ཁ",
			language.GRC:     "གིརིཀ, སྔ་དུས་ཀྱི (༡༤༥༣)",
			language.GSW:     "སུ་ཡིས་ཇཱར་མཱན་ཁ",
			language.GU:      "གུ་ཇ་ར་ཏི་ཁ",
			language.HA:      "ཧཝ་ས་ཁ",
			language.HAW:     "ཧ་ཝ་ཡིའི་ཁ",
			language.HE:      "ཧེ་བྲུ་ཁ",
			language.HI:      "ཧིན་དི་ཁ",
			language.HR:      "ཀྲོ་ཨེ་ཤི་ཡཱན་ཁ",
			language.HT:      "ཧེ་ཏི་ཡཱན་ཁ",
			language.HU:      "ཧཱང་གྷ་རི་ཡཱན་ཁ",
			language.HY:      "ཨར་མི་ནི་ཡཱན་ཁ",
			language.ID:      "ཨིན་ཌོ་ནེ་ཤི་ཡཱན་ཁ",
			language.IG:      "ཨིག་བོ་ཁ",
			language.IS:      "ཨ་ཡིས་ལེན་ཌིཀ་ཁ",
			language.IT:      "ཨི་ཊ་ལི་ཡཱན་ཁ",
			language.JA:      "ཇཱ་པཱ་ནིས་ཁ",
			language.JV:      "ཇཱ་བ་ནིས་ཁ",
			language.KA:      "ཇཽ་ཇི་ཡཱན་ཁ",
			language.KAC:     "ཀ་ཆིན་ཁ",
			language.KFO:     "ཀོ་རོ་ཁ",
			language.KK:      "ཀ་ཛགས་ཁ",
			language.KM:      "ཁེ་མེར་ཁ",
			language.KN:      "ཀ་ན་ཌ་ཁ",
			language.KO:      "ཀོ་རི་ཡཱན་ཁ",
			language.KS:      "ཀཱཤ་མི་རི་ཁ",
			language.KU:      "ཀར་ཌིཤ་ཁ",
			language.KY:      "ཀིར་གིས་ཁ",
			language.LA:      "ལེ་ཊིན་ཁ",
			language.LB:      "ལག་ཛམ་བོརྒ་ཁ",
			language.LO:      "ལཱ་ཝོས་ཁ",
			language.LT:      "ལི་ཐུ་ཝེ་ནི་ཡཱན་ཁ",
			language.LV:      "ལཊ་བི་ཡཱན་ཁ",
			language.MG:      "མ་ལ་ག་སི་ཁ",
			language.MI:      "མ་ཨོ་རི་ཁ",
			language.MK:      "མ་སེ་ཌོ་ནི་ཡཱན་ཁ",
			language.ML:      "མ་ལ་ཡ་ལམ་ཁ",
			language.MN:      "སོག་པོའི་ཁ",
			language.MNC:     "མན་ཇུ་ཁ",
			language.MNI:     "མ་ནི་པུ་རི",
			language.MR:      "མ་ར་ཐི་ཁ",
			language.MS:      "མ་ལེ་ཁ",
			language.MT:      "མཱལ་ཊ་ཁ",
			language.MY:      "བར་མིས་ཁ",
			language.NB:      "ནོར་ཝེ་ཇི་ཡཱན་བོཀ་མཱལ་ཁ",
			language.NE:      "ནེ་པཱལི་ཁ",
			language.NEW:     "ནི་ཝ་རི",
			language.NL:      "ཌཆ་ཁ",
			language.NL_BE:   "ཕྷེལེ་མིཤ་ཁ",
			language.NN:      "ནོར་ཝེ་ཇི་ཡཱན་ནོརསཀ་ཁ",
			language.NO:      "ནོར་ཝི་ཇི་ཡན་ཁ",
			language.OR:      "ཨོ་རི་ཡ་ཁ",
			language.PA:      "པཱན་ཇ་བི་ཁ",
			language.PI:      "པ་ལི",
			language.PL:      "པོ་ལིཤ་ཁ",
			language.PS:      "པཱཤ་ཏོ་ཁ",
			language.PT:      "པོར་ཅུ་གིས་ཁ",
			language.PT_BR:   "བྲ་ཛི་ལི་ཡཱན་པོར་ཅུ་གིས་ཁ",
			language.PT_PT:   "ཨི་བེ་རི་ཡཱན་པོར་ཅུ་གིས་ཁ",
			language.QU:      "ཀྭེ་ཆུ་ཨ་ཁ",
			language.RM:      "རོ་མེ་ནིཤ་ཁ",
			language.RO:      "རོ་མེ་ནི་ཡཱན་ཁ",
			language.RU:      "ཨུ་རུ་སུའི་ཁ",
			language.SA:      "སཾསྐྲྀཏ་ཁ",
			language.SD:      "སིན་དཱི་ཁ",
			language.SHN:     "ཤཱན་ཁ",
			language.SI:      "སིང་ཧ་ལ་ཁ",
			language.SK:      "སུ་ལོ་བཱཀ་ཁ",
			language.SL:      "སུ་ལོ་བི་ནི་ཡཱན་ཁ",
			language.SO:      "སོ་མ་ལི་ཁ",
			language.SQ:      "ཨཱལ་བེ་ནི་ཡཱན་ཁ",
			language.SR:      "སཱར་བྷི་ཡཱན་ཁ",
			language.SU:      "སཱུན་ད་ནིས་ཁ",
			language.SV:      "སུའི་ཌིཤ་ཁ",
			language.SW:      "སྭཱ་ཧི་ལི་ཁ",
			language.TA:      "ཏ་མིལ་ཁ",
			language.TE:      "ཏེ་ལུ་གུ་ཁ",
			language.TG:      "ཏ་ཇིཀ་ཁ",
			language.TH:      "ཐཱའི་ཁ",
			language.TI:      "ཏིག་རི་ཉ་ཁ",
			language.TK:      "ཊཱརཀ་མེན་ཁ",
			language.TO:      "ཊོང་གྷན་ཁ",
			language.TR:      "ཊཱར་ཀིཤ་ཁ",
			language.TT:      "ཊ་ཊར་ཁ",
			language.UG:      "ཝི་གུར་ཁ",
			language.UK:      "ཡུ་ཀེ་རེ་ནི་ཡཱན་ཁ",
			language.UND:     "ཁ་ངོ་མ་ཤེསཔ",
			language.UR:      "ཨུར་དུ་ཁ",
			language.UZ:      "ཨུས་བེཀ་ཁ",
			language.VI:      "བེཊ་ནཱ་མིས་ཁ",
			language.WO:      "ཝོ་ལོཕ་ཁ",
			language.XH:      "ཞོ་ས་ཁ",
			language.YO:      "ཡོ་རུ་བ་ཁ",
			language.ZH:      "རྒྱ་མི་ཁ",
			language.ZH_HANS: "རྒྱ་མི་ཁ་འཇམ་སངམ",
			language.ZH_HANT: "སྔ་དུས་ཀྱི་རྒྱ་མི་ཁ",
			language.ZU:      "ཟུ་ལུ་ཁ",
			language.ZXX:     "སྐད་རིག་ནང་དོན་མེདཔ",
		},
		Territories: cldr.Territories{
			territory.V_001: "འཛམ་གླིང༌",
			territory.V_002: "ཨཕ་རི་ཀ",
			territory.V_003: "བྱང་ཨ་མི་རི་ཀ",
			territory.V_005: "ལྷོ་ཨ་མི་རི་ཀ",
			territory.V_009: "ཨོཤི་ཡཱན་ན",
			territory.V_011: "ནུབ་ཕྱོགས་ཀྱི་ཨཕ་རི་ཀ",
			territory.V_013: "བར་ཕྱོགས་ཨ་མི་རི་ཀ",
			territory.V_014: "ཤར་ཕྱོགས་ཀྱི་ཨཕ་རི་ཀ",
			territory.V_015: "བྱང་ཕྱོགས་ཀྱི་ཨཕ་རི་ཀ",
			territory.V_017: "སྦུག་ཕྱོགས་ཀྱི་ཨཕ་རི་ཀ",
			territory.V_018: "ལྷོའི་ཨཕ་རི་ཀ",
			territory.V_019: "ཨ་མི་རི་ཀ་ཚུ",
			territory.V_021: "བྱང་ཕྱོགས་ཀྱི་ཨ་མི་རི་ཀ",
			territory.V_029: "ཀེ་རི་བི་ཡེན",
			territory.V_030: "ཤར་ཕྱོགས་ཀྱི་ཨེ་ཤི་ཡ",
			territory.V_034: "ལྷོའི་ཨེ་ཤི་ཡ",
			territory.V_035: "ལྷོ་ཤར་ཕྱོགས་ཀྱི་ཨེ་ཤི་ཡ",
			territory.V_039: "ལྷོའི་ཡུ་རོབ",
			territory.V_053: "ཨཱོས་ཊྲེལ་ཨེ་ཤི་ཡ",
			territory.V_054: "མེ་ལ་ནི་ཤི་ཡ",
			territory.V_057: "ལུང་ཕྱོགས་མའི་ཀྲོ་ནི་ཤི་ཡ",
			territory.V_061: "པོ་ལི་ནི་ཤི་ཡ",
			territory.V_142: "ཨེ་ཤི་ཡ",
			territory.V_143: "སྦུག་ཕྱོགས་ཀྱི་ཨེ་ཤི་ཡ",
			territory.V_145: "ནུབ་ཕྱོགས་ཀྱི་ཨེ་ཤི་ཡ",
			territory.V_150: "ཡུ་རོབ",
			territory.V_151: "ཤར་ཕྱོགས་ཀྱི་ཡུ་རོབ",
			territory.V_154: "བྱང་ཕྱོགས་ཀྱི་ཡུ་རོབ",
			territory.V_155: "ནུབ་ཕྱོགས་ཀྱི་ཡུ་རོབ",
			territory.V_419: "ལེ་ཊིནཨ་མི་རི་ཀ",
			territory.AC:    "ཨེ་སེན་ཤུན་ཚོ་གླིང༌",
			territory.AD:    "ཨཱན་དོ་ར",
			territory.AE:    "ཡུ་ནཱའི་ཊེཌ་ ཨ་རབ་ ཨེ་མེ་རེཊས",
			territory.AF:    "ཨཕ་གྷ་ནི་སཏཱན",
			territory.AG:    "ཨན་ཊི་གུ་ཝ་ ཨེནཌ་ བྷར་བྷུ་ཌ",
			territory.AI:    "ཨང་གི་ལ",
			territory.AL:    "ཨཱལ་བེ་ནི་ཡ",
			territory.AM:    "ཨར་མི་ནི་ཡ",
			territory.AO:    "ཨང་གྷོ་ལ",
			territory.AQ:    "འཛམ་གླིང་ལྷོ་མཐའི་ཁྱགས་གླིང",
			territory.AR:    "ཨར་ཇེན་ཊི་ན",
			territory.AS:    "ས་མོ་ཨ་ཡུ་ཨེས་ཨེ་མངའ་ཁོངས",
			territory.AT:    "ཨཱོས་ཊྲི་ཡ",
			territory.AU:    "ཨཱོས་ཊྲེལ་ལི་ཡ",
			territory.AW:    "ཨ་རུ་བཱ",
			territory.AX:    "ཨ་ལནཌ་གླིང་ཚོམ",
			territory.AZ:    "ཨ་ཛར་བྷའི་ཇཱན",
			territory.BA:    "བྷོས་ནི་ཡ་ ཨེནཌ་ ཧར་ཛི་གྷོ་བི་ན",
			territory.BB:    "བྷར་བེ་ཌོས",
			territory.BD:    "བངྒ་ལ་དེཤ",
			territory.BE:    "བྷེལ་ཇམ",
			territory.BF:    "བྷར་ཀི་ན་ ཕེ་སོ",
			territory.BG:    "བུལ་ག་རི་ཡ",
			territory.BH:    "བྷ་རེན",
			territory.BI:    "བྷུ་རུན་ཌི",
			territory.BJ:    "བྷེ་ནིན",
			territory.BL:    "སེནཊ་ བར་ཐོ་ལོམ་མིའུ",
			territory.BM:    "བར་མུ་ཌ",
			territory.BN:    "བྷྲུ་ནའི",
			territory.BO:    "བྷེ་ལི་བི་ཡ",
			territory.BQ:    "ཀེ་རི་བི་ཡེན་ནེ་དར་ལནཌས྄",
			territory.BR:    "བྲ་ཛིལ",
			territory.BS:    "བྷ་ཧ་མས྄",
			territory.BT:    "འབྲུག",
			territory.BV:    "བོའུ་ཝེཊ་མཚོ་གླིང",
			territory.BW:    "བྷོཙ་ཝ་ན",
			territory.BY:    "བེལ་ཨ་རུ་སུ",
			territory.BZ:    "བྷེ་ལིཛ",
			territory.CA:    "ཀེ་ན་ཌ",
			territory.CC:    "ཀོ་ཀོས་གླིང་ཚོམ",
			territory.CD:    "ཀོང་གྷོ ཀིན་ཤ་ས",
			territory.CF:    "སེན་ཊལ་ ཨཕ་རི་ཀཱན་ རི་པབ་ལིཀ",
			territory.CG:    "ཀོང་གྷོ བྷྲ་ཛ་བིལ",
			territory.CH:    "སུ་ཝིཊ་ཛར་ལེནཌ",
			territory.CI:    "ཀོ་ཊེ་ ཌི་ཨི་ཝོ་རེ",
			territory.CK:    "ཀུག་གླིང་ཚོམ",
			territory.CL:    "ཅི་ལི",
			territory.CM:    "ཀེ་མ་རུན",
			territory.CN:    "རྒྱ་ནག",
			territory.CO:    "ཀོ་ལོམ་བྷི་ཡ",
			territory.CP:    "ཀི་ལི་པེར་ཊོན་མཚོ་གླིང་",
			territory.CR:    "ཀོས་ཊ་རི་ཀ",
			territory.CU:    "ཀིའུ་བྷ",
			territory.CV:    "ཀེཔ་བཱཌ",
			territory.CW:    "ཀྱཱུར་ར་ཀོ",
			territory.CX:    "ཁི་རིསྟ་མེས་མཚོ་གླིང",
			territory.CY:    "སཱའི་པྲས",
			territory.CZ:    "ཅེཀ་ རི་པབ་ལིཀ",
			territory.DE:    "ཇཱར་མ་ནི",
			territory.DG:    "ཌི་ཡེ་གོ་གར་སིའོ",
			territory.DJ:    "ཇི་བྷུ་ཊི",
			territory.DK:    "ཌེན་མཱཀ",
			territory.DM:    "ཌོ་མི་ནི་ཀ",
			territory.DO:    "ཌོ་མི་ནི་ཀཱན་ རི་པབ་ལིཀ",
			territory.DZ:    "ཨཱལ་ཇི་རི་ཡ",
			territory.EA:    "སེ་ཨུ་ཏ་ ཨེནཌ་ མེལ་ལི་ལ",
			territory.EC:    "ཨེ་ཁྭ་ཌོར",
			territory.EE:    "ཨེས་ཊོ་ནི་ཡ",
			territory.EG:    "ཨི་ཇིབཊ",
			territory.EH:    "ནུབ་ཕྱོགས་ ས་ཧཱ་ར",
			territory.ER:    "ཨེ་རི་ཊྲེ་ཡ",
			territory.ES:    "ཨིས་པེན",
			territory.ET:    "ཨི་ཐི་ཡོ་པི་ཡ",
			territory.EU:    "ཡུ་རོབ་གཅིག་བསྡོམས་ཚོགས་པ",
			territory.FI:    "ཕིན་ལེནཌ",
			territory.FJ:    "ཕི་ཇི",
			territory.FK:    "ཕལྐ་ལནྜ་གླིང་ཚོམ",
			territory.FM:    "མའི་ཀྲོ་ནི་ཤི་ཡ",
			territory.FO:    "ཕཱའེ་རོ་གླིང་ཚོམ",
			territory.FR:    "ཕྲཱནས",
			territory.GA:    "གྷ་བྷོན",
			territory.GB:    "ཡུ་ནཱའི་ཊེཌ་ ཀིང་ཌམ",
			territory.GD:    "གྲྀ་ན་ཌ",
			territory.GE:    "ཇཽར་ཇཱ",
			territory.GF:    "གུའི་ཡ་ན་ ཕྲནས྄་མངའ་ཁོངས",
			territory.GG:    "གུ་ཨེརྣ་སི",
			territory.GH:    "གྷ་ན",
			territory.GI:    "ཇིབ་རཱལ་ཊར",
			territory.GL:    "གིརཱིན་ལནཌ྄",
			territory.GM:    "གྷེམ་བི་ཡ",
			territory.GN:    "གྷི་ནི",
			territory.GP:    "གོ་ཌེ་ལུ་པེ",
			territory.GQ:    "ཨེ་ཀུ་ཊོ་རེལ་ གི་ནི",
			territory.GR:    "གིརིས྄",
			territory.GS:    "སཱའུཐ་ཇཽར་ཇཱ་ དང་ སཱའུཐ་སེནཌ྄་ཝིཅ་གླིང་ཚོམ",
			territory.GT:    "གྷོ་ཊ་མ་ལ",
			territory.GU:    "གུ་འམ་ མཚོ་གླིང",
			territory.GW:    "གྷི་ནི་ བྷི་སཱའུ",
			territory.GY:    "གྷ་ཡ་ན",
			territory.HK:    "ཧོང་ཀོང་ཅཱའི་ན",
			territory.HM:    "ཧཱརཌ་མཚོ་གླིང་ དང་ མེཀ་ཌོ་ནལཌ྄་གླིང་ཚོམ",
			territory.HN:    "ཧཱན་ཌུ་རཱས྄",
			territory.HR:    "ཀྲོ་ཨེ་ཤ",
			territory.HT:    "ཧེ་ཊི",
			territory.HU:    "ཧཱང་གྷ་རི",
			territory.IC:    "ཀ་ནེ་རི་གླིང་ཚོམ",
			territory.ID:    "ཨིན་ཌོ་ནེ་ཤི་ཡ",
			territory.IE:    "ཨཱ་ཡ་ལེནཌ",
			territory.IL:    "ཨིས་ར་ཡེལ",
			territory.IM:    "ཨ་ཡུལ་ ཨོཕ་ མཱན",
			territory.IN:    "རྒྱ་གར",
			territory.IO:    "བྲི་ཊིཤ་རྒྱ་གར་གྱི་རྒྱ་མཚོ་ས་ཁོངས",
			territory.IQ:    "ཨི་རཱཀ",
			territory.IR:    "ཨི་རཱན",
			territory.IS:    "ཨཱའིས་ལེནཌ",
			territory.IT:    "ཨི་ཊ་ལི",
			territory.JE:    "ཇེར་སི",
			territory.JM:    "ཇཱ་མཻ་ཀ",
			territory.JO:    "ཇོར་ཌན",
			territory.JP:    "ཇ་པཱན",
			territory.KE:    "ཀེན་ཡ",
			territory.KG:    "ཀིར་གིས་སཏཱན",
			territory.KH:    "ཀམ་བྷོ་ཌི་ཡ",
			territory.KI:    "ཀི་རི་བ་ཏི་མཚོ་གླིང",
			territory.KM:    "ཀོ་མོ་རོས",
			territory.KN:    "སེནཊ་ ཀིཊས་ དང་ ནེ་བིས",
			territory.KP:    "བྱང་ ཀོ་རི་ཡ",
			territory.KR:    "ལྷོ་ ཀོ་རི་ཡ",
			territory.KW:    "ཀུ་ཝེཊ",
			territory.KY:    "ཁེ་མེན་གླིང་ཚོམ",
			territory.KZ:    "ཀ་ཛགས་སཏཱན",
			territory.LA:    "ལཱ་ཝོས",
			territory.LB:    "ལེ་བ་ནོན",
			territory.LC:    "སེནཊ་ ལུ་སི་ཡ",
			territory.LI:    "ལིཀ་ཏནས་ཏ་ཡིན",
			territory.LK:    "ཤྲཱི་ལང་ཀ",
			territory.LR:    "ལཱའི་བེ་རི་ཡ",
			territory.LS:    "ལཻ་སོ་ཐོ",
			territory.LT:    "ལི་ཐུ་ཝེ་ནི་ཡ",
			territory.LU:    "ལག་ཛམ་བོརྒ",
			territory.LV:    "ལཊ་བི་ཡ",
			territory.LY:    "ལི་བི་ཡ",
			territory.MA:    "མོ་རོ་ཀོ",
			territory.MC:    "མོ་ན་ཀོ",
			territory.MD:    "མོལ་དོ་བཱ",
			territory.ME:    "མོན་ཊི་ནེག་རོ",
			territory.MF:    "སེནཊ་ མཱར་ཊིན",
			territory.MG:    "མ་དཱ་གེས་ཀར",
			territory.MH:    "མར་ཤེལ་གླིང་ཚོམ",
			territory.ML:    "མཱ་ལི",
			territory.MM:    "མི་ཡཱན་མར་ (བྷར་མ)",
			territory.MN:    "སོག་པོ་ཡུལ",
			territory.MO:    "མཀ་ཨའུ་ཅཱའི་ན",
			territory.MP:    "བྱང་ཕྱོགས་ཀྱི་མ་ར་ཡ་ན་གླིང་ཚོམ",
			territory.MQ:    "མཱར་ཊི་ནིཀ",
			territory.MR:    "མོ་རི་ཊེ་ནི་ཡ",
			territory.MS:    "མོན་ས་རཊ",
			territory.MT:    "མཱལ་ཊ",
			territory.MU:    "མོ་རི་ཤཱས",
			territory.MV:    "མཱལ་དིབས",
			territory.MW:    "མ་ལ་ཝི",
			territory.MX:    "མེཀ་སི་ཀོ",
			territory.MY:    "མ་ལེ་ཤི་ཡ",
			territory.MZ:    "མོ་ཛམ་བྷིཀ",
			territory.NA:    "ན་མི་བི་ཡ",
			territory.NC:    "ནིའུ་ཀ་ལི་དོ་ནི་ཡ",
			territory.NE:    "ནཱའི་ཇཱ",
			territory.NF:    "ནོར་ཕོལཀ་མཚོ་གླིང༌",
			territory.NG:    "ནཱའི་ཇི་རི་ཡ",
			territory.NI:    "ནི་ཀྲ་ཝ་ག",
			territory.NL:    "ནེ་དར་ལནཌས྄",
			territory.NO:    "ནོར་ཝེ",
			territory.NP:    "བལ་ཡུལ",
			territory.NR:    "ནའུ་རུ་",
			territory.NU:    "ནི་ཨུ་ཨཻ",
			territory.NZ:    "ནིའུ་ཛི་ལེནཌ",
			territory.OM:    "ཨོ་མཱན",
			territory.PA:    "པ་ན་མ",
			territory.PE:    "པེ་རུ",
			territory.PF:    "ཕྲཱནས྄་ཀྱི་པོ་ལི་ནི་ཤི་ཡ",
			territory.PG:    "པ་པུ་ ནིའུ་གི་ནི",
			territory.PH:    "ཕི་ལི་པིནས",
			territory.PK:    "པ་ཀི་སཏཱན",
			territory.PL:    "པོ་ལེནཌ",
			territory.PM:    "སིནཊ་པི་ཡེར་ ཨེནཌ་ མིཀོ་ལེན",
			territory.PN:    "པིཊ་ཀེ་ཡེརན་གླིང་ཚོམ",
			territory.PR:    "པུ་འེར་ཊོ་རི་ཁོ",
			territory.PS:    "པེ་ལིསི་ཊི་ནི་ཡན་ཊེ་རི་ཐོ་རི",
			territory.PT:    "པོར་ཅུ་གཱལ",
			territory.PW:    "པ་ལའུ",
			territory.PY:    "པ་ར་གུ་ཝའི",
			territory.QA:    "ཀ་ཊར",
			territory.QO:    "ཨོཤི་ཡཱན་ན་གྱི་མཐའ་མཚམས",
			territory.RE:    "རེ་ཡུ་ནི་ཡོན",
			territory.RO:    "རོ་མེ་ནི་ཡ",
			territory.RS:    "སཱར་བྷི་ཡ",
			territory.RU:    "ཨུ་རུ་སུ",
			territory.RW:    "རུ་ཝན་ཌ",
			territory.SA:    "སཱཝ་དི་ ཨ་རེ་བྷི་ཡ",
			territory.SB:    "སོ་ལོ་མོན་ གླིང་ཚོམ",
			territory.SC:    "སེ་ཤཱལས",
			territory.SD:    "སུ་ཌཱན",
			territory.SE:    "སུའི་ཌེན",
			territory.SG:    "སིང་ག་པོར",
			territory.SH:    "སེནཊ་ ཧེ་ལི་ན",
			territory.SI:    "སུ་ལོ་བི་ནི་ཡ",
			territory.SJ:    "སྭཱལ་བྷརྡ་ ཨེནཌ་ ཇཱན་མ་ཡེན",
			territory.SK:    "སུ་ལོ་བཱ་ཀི་ཡ",
			territory.SL:    "སི་ར་ ལི་འོན",
			territory.SM:    "སཱན་མ་རི་ནོ",
			territory.SN:    "སེ་ནི་གྷལ",
			territory.SO:    "སོ་མ་ལི་ཡ",
			territory.SR:    "སུ་རི་ནཱམ",
			territory.SS:    "སཱའུཐ་ སུ་ཌཱན",
			territory.ST:    "སཝ་ ཊོ་མེ་ ཨེནཌ་ པྲྀན་སི་པེ",
			territory.SV:    "ཨེལ་སལ་བ་ཌོར",
			territory.SX:    "སིནཊ་ མཱར་ཊེན",
			territory.SY:    "སི་རི་ཡ",
			territory.SZ:    "སུ་ཝ་ཛི་ལེནཌ",
			territory.TA:    "ཏྲིས་ཏན་ད་ཀུན་ཧ",
			territory.TC:    "ཏུརྐས྄་ ཨེནཌ་ ཀ་ཀོས་གླིང་ཚོམ",
			territory.TD:    "ཅཱཌ",
			territory.TF:    "ཕྲནཅ་གི་ལྷོ་ཕྱོགས་མངའ་ཁོངས",
			territory.TG:    "ཊོ་གྷོ",
			territory.TH:    "ཐཱའི་ལེནཌ",
			territory.TJ:    "ཏ་ཇིག་གི་སཏཱན",
			territory.TK:    "ཏོ་ཀེ་ལའུ་ མཚོ་གླིང",
			territory.TL:    "ཏི་་མོར་ལེ་ཨེསཊ",
			territory.TM:    "ཊཱརཀ་མེནའི་སཏཱན",
			territory.TN:    "ཊུ་ནི་ཤི་ཡ",
			territory.TO:    "ཊོང་གྷ",
			territory.TR:    "ཊཱར་ཀི",
			territory.TT:    "ཊི་ནི་ཌཱཌ་ ཨེནཌ་ ཊོ་བྷེ་གྷོ",
			territory.TV:    "ཏུ་ཝ་ལུ",
			territory.TW:    "ཊཱའི་ཝཱན",
			territory.TZ:    "ཊཱན་ཛཱ་ནི་ཡ",
			territory.UA:    "ཡུ་ཀརེན",
			territory.UG:    "ཡུ་གྷན་ཌ",
			territory.UM:    "ཡུ་ཨེས་གྱི་མཐའ་མཚམས་མཚོ་གླིང་",
			territory.US:    "ཡུ་ཨེས་ཨེ",
			territory.UY:    "ཡུ་རུ་གུ་ཝའི",
			territory.UZ:    "ཨུས་བེག་གི་སཏཱན",
			territory.VA:    "བ་ཊི་ཀཱན་ སི་ཊི",
			territory.VC:    "སེནཊ་ཝིན་སེནཌ྄ ཨེནཌ་ གི་རེ་ན་དིནས྄",
			territory.VE:    "བེ་ནི་ཛུ་ཝེ་ལ",
			territory.VG:    "ཝརཇིན་གླིང་ཚོམ་ བྲཱི་ཊིཤ་མངའ་ཁོངས",
			territory.VI:    "ཝརཇིན་གླིང་ཚོམ་ ཡུ་ཨེས་ཨེ་མངའ་ཁོངས",
			territory.VN:    "བེཊ་ནཱམ",
			territory.VU:    "ཝ་ནུ་ཨ་ཏུ",
			territory.WF:    "ཝལ་ལིས྄་ ཨེནཌ་ ཕུ་ཏུ་ན་",
			territory.WS:    "ས་མོ་ཨ",
			territory.YE:    "ཡེ་མེན",
			territory.YT:    "མེ་ཡོཊ",
			territory.ZA:    "སཱའུཐ་ ཨཕ་རི་ཀ",
			territory.ZM:    "ཛམ་བྷི་ཡ",
			territory.ZW:    "ཛིམ་བྷབ་ཝེ",
			territory.ZZ:    "ངོ་མ་ཤེས་པའི་ལུང་ཕྱོགས",
		},
	}
}
