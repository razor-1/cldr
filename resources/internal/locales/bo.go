// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_bo() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "bo",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y MMMMའི་ཚེས་d, EEEE", Long: "སྤྱི་ལོ་y MMMMའི་ཚེས་d", Medium: "y ལོའི་MMMཚེས་d", Short: "y-MM-dd"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ཟླ་༡", Feb: "ཟླ་༢", Mar: "ཟླ་༣", Apr: "ཟླ་༤", May: "ཟླ་༥", Jun: "ཟླ་༦", Jul: "ཟླ་༧", Aug: "ཟླ་༨", Sep: "ཟླ་༩", Oct: "ཟླ་༡༠", Nov: "ཟླ་༡༡", Dec: "ཟླ་༡༢"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ཟླ་བ་དང་པོ་", Feb: "ཟླ་བ་གཉིས་པ་", Mar: "ཟླ་བ་གསུམ་པ་", Apr: "ཟླ་བ་བཞི་པ་", May: "ཟླ་བ་ལྔ་པ་", Jun: "ཟླ་བ་དྲུག་པ་", Jul: "ཟླ་བ་བདུན་པ་", Aug: "ཟླ་བ་བརྒྱད་པ་", Sep: "ཟླ་བ་དགུ་པ་", Oct: "ཟླ་བ་བཅུ་པ་", Nov: "ཟླ་བ་བཅུ་གཅིག་པ་", Dec: "ཟླ་བ་བཅུ་གཉིས་པ་"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ཉི་མ་", Mon: "ཟླ་བ་", Tue: "མིག་དམར་", Wed: "ལྷག་པ་", Thu: "ཕུར་བུ་", Fri: "པ་སངས་", Sat: "སྤེན་པ་"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ཉི", Mon: "ཟླ", Tue: "མིག", Wed: "ལྷག", Thu: "ཕུར", Fri: "སངས", Sat: "སྤེན"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "གཟའ་ཉི་མ་", Mon: "གཟའ་ཟླ་བ་", Tue: "གཟའ་མིག་དམར་", Wed: "གཟའ་ལྷག་པ་", Thu: "གཟའ་ཕུར་བུ་", Fri: "གཟའ་པ་སངས་", Sat: "གཟའ་སྤེན་པ་"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "སྔ་དྲོ་", PM: "ཕྱི་དྲོ་"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
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
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "ཡུ་ཨན་", Symbol: "¥"},
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
				currency.INR: cldr.Currency{DisplayName: "རྒྱ་གར་སྒོར་", Symbol: "₹"},
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
				currency.RUB: cldr.Currency{DisplayName: "", Symbol: "₽"},
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
				currency.USD: cldr.Currency{DisplayName: "ཨ་རིའི་སྒོར་", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "མ་རྟོགས་པའི་ནུས་མེད་དངུལ་ལོར", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.BN:      "བྷང་ག་ལའི་སྐད།",
			language.BO:      "བོད་སྐད་",
			language.DA:      "ཏེན་མག་གི་སྐད།",
			language.DE:      "འཇར་མན་གྱི།",
			language.DZ:      "རྫོང་ཁ",
			language.EN:      "དབྱིན་ཇིའི་སྐད།",
			language.EN_CA:   "དབྱིན་ཇིའི་སྐད། (ཁེ་ན་ཌ་)",
			language.EN_GB:   "དབྱིན་ཇིའི་སྐད། (དབྱིན་ལན་)",
			language.EN_US:   "དབྱིན་ཇིའི་སྐད། (ཨ་རི་)",
			language.ES:      "ཞི་པན་ཡའི།",
			language.FR:      "ཧྥ་རན་སིའི་།",
			language.GAA:     "གཱ་སྐད།",
			language.GU:      "གུཇ་རཱཏི་སྐད།",
			language.HI:      "ཧིན་དི",
			language.ID:      "ཨིན་དོ་ནི་སི་སྐད།",
			language.IT:      "དབྱི་ཏ་་ལའི་མི།",
			language.JA:      "ཉི་ཧོང་སྐད་",
			language.KN:      "ཀནྣ་ཌ་སྐད།",
			language.KO:      "ཁོ་རེ་ཡའི་སྐད།",
			language.LA:      "ལཱ་ཏིན་སྐད།",
			language.MN:      "སོག་སྐད།",
			language.MR:      "མ་ར་ཐི་སྐད།",
			language.MS:      "མ་ལ་ཡ་སྐད།",
			language.MY:      "འབར་མའི་སྐད།",
			language.NE:      "ནེ་པ་ལི",
			language.NL:      "ཧའོ་ལན་སྐད།",
			language.NN:      "ནོ་ཝེ་སྐད།",
			language.OR:      "ཨཽ་རི་ཡ་སྐད།",
			language.PL:      "པོ་ལན་སྐད།",
			language.PT:      "ཕི་ཐོ་ཡའི།",
			language.PT_BR:   "པ་ཞའི་མི། ཕི་ཐོ་ཡའི་མི།",
			language.RU:      "ཨུ་རུ་སུ་སྐད་",
			language.SA:      "སཾ་སྐྲྀ་ཏ།",
			language.SI:      "ཞི་ལན་སྐད།",
			language.SV:      "ཧྲུའི་ཏན་སྐད།",
			language.TA:      "ཏཱ་མིལ་སྐད།",
			language.TE:      "ཏེ་ལུ་གུ་སྐད།",
			language.TH:      "ཐའའི་ཡུལ་སྐད།",
			language.TR:      "ཐུར་ཁེའི་སྐད།",
			language.UG:      "ཡུ་གུར་སྐད།",
			language.UK:      "ཡུ་ཀྲ་ནི་སྐད།",
			language.UND:     "མིའི་ཤེས་རྟོགས་མ་བྱུང་བ། ཡང་ན་ཆད་ལྷག་ཅན་གྱི་སྐད་བརྡ།",
			language.UR:      "ཝུའུ་ཏུའུ་སྐད།",
			language.VI:      "ཡོ་ནན་སྐད།",
			language.ZEN:     "ཟེ་ན་གཱ་སྐད།",
			language.ZH:      "རྒྱ་སྐད་",
			language.ZH_HANS: "སྟབས་བརྡའི། ཀྲུང་གོའི།",
			language.ZH_HANT: "སྲོལ་རྒྱུན་གྱི།",
			language.ZU:      "ཟུ་ལུ་སྐད།",
			language.ZUN:     "ཟུ་ནི་སྐད།",
			language.ZZA:     "ཟ་ཟའ་སྐད།",
		},
		Territories: cldr.Territories{
			territory.V_001: "འཛམ་གླིང་།",
			territory.V_002: "ཨཕྲི་ཀ།",
			territory.V_053: "ཨསྟྲེ་ལི་ཡ་དང་། ནིའུ་ཛི་ལན྄ཌ།",
			territory.V_142: "ཨེ་ཤི་ཡ།",
			territory.V_150: "ཡུ་རོབ།",
			territory.AD:    "ཨེན་ཌོ་ར།",
			territory.AE:    "ཨ་རབ། ཨི་མི་རཊ྄། ཆིག་སྒྲིལ་རྒྱལ་ཁབ།",
			territory.AF:    "ཨཕ་ག་ནི་སྟཱན།",
			territory.AG:    "ཨེན་ཊི་གུ་དང་། བྷར་བུ་ཌ།",
			territory.AI:    "ཨང་གུའི་ལ།",
			territory.AL:    "ཨལ་བཱ་ནི་ཡ།",
			territory.AM:    "ཨར་མེ་ནི་ཡ།",
			territory.AO:    "ཨང་གཽ་ལ།",
			territory.AQ:    "ལྷོ་རྩེའི་མཐའ་གླིང་།",
			territory.AR:    "ཨར་ཇེན་ཊི་ན།",
			territory.AT:    "ཨསྟྲི་ཡ།",
			territory.AU:    "ཨསྟྲེ་ལི་ཡ།",
			territory.AW:    "ཨ་རུ་བ།",
			territory.AZ:    "ཨཛར་བཡེ་ཇན།",
			territory.BA:    "བོསྣི་ཡ་དང་ཧརྫོ་གོ་ཝི་ན།",
			territory.BB:    "བཱརྦ་ཌོས྄།",
			territory.BD:    "བངྒ་ལ་དེཤ།",
			territory.BE:    "བེལ་ཇི་ཡམ།",
			territory.BF:    "བརཀི་ན། ཕསོ།",
			territory.BG:    "བུལ་ག་རི་ཡ།",
			territory.BH:    "བྷཱ་རེན།",
			territory.BI:    "བུ་རུན་ཌི།",
			territory.BJ:    "བཱེ་ནིན།",
			territory.BM:    "བར་མུ་ཌ།",
			territory.BN:    "བུ་རུ་ནེ།",
			territory.BO:    "བོ་ལི་ཝིཡ།",
			territory.BR:    "བ་རཱ་ཛིལ།",
			territory.BS:    "བྷཱ་མས྄།",
			territory.BT:    "འབྲུག་ཡུལ།",
			territory.BW:    "བོཙ་ཝ་ན།",
			territory.BY:    "བེ་ལུ་རུ་སུ།",
			territory.BZ:    "བེ་ལིཛ།",
			territory.CA:    "ཁེ་ན་ཌ།",
			territory.CH:    "ཧྲུད་ཧྲི།",
			territory.CI:    "ཀོ་ཊེ་ཌི། ཨི་ཝོ་རེ།",
			territory.CK:    "ཀཱུག གླིང་ཕྲེན་རྒྱལ་ཁབ།",
			territory.CL:    "ཅི་ལི།",
			territory.CM:    "ཀ་མེ་རུན།",
			territory.CN:    "རྒྱ་ནག",
			territory.CO:    "ཀོ་ལོམ་བི་ཡ།",
			territory.CR:    "ཀོ་ས྄ཊ་རི་ཀ།",
			territory.CU:    "ཁྱུའུ་བ།",
			territory.CY:    "སཱཡེ་པ་རས྄།",
			territory.CZ:    "ཅཻག་སྤྱི་མཐུན་རྒྱལ་ཁབ།",
			territory.DE:    "འཇར་མན་",
			territory.DJ:    "ཛི་བུ་ཏི།",
			territory.DK:    "ཌེན་མཱརྐ།",
			territory.DM:    "ཌོ་མིན་ནི་ཀ།",
			territory.DO:    "ཌོ་མིནནི་ཀན་སྤྱི་མཐུན་རྒྱལ་ཁབ།",
			territory.DZ:    "ཨལ་ཇི་རི་ཡ།",
			territory.EC:    "ཨི་ཁྭ་ཌོར།",
			territory.EE:    "ཨིསྟོ་ནི་ཡ།",
			territory.EG:    "ཨི་ཇིབྚ།",
			territory.ER:    "ཨེ་རི་ཏྲེ་ཨ།",
			territory.ES:    "སི་པན།",
			territory.ET:    "ཨི་ཐིའོ་པི་ཡ།",
			territory.FI:    "ཕིན་ལན྄ཌ།",
			territory.FJ:    "ཕི་ཇི།",
			territory.FK:    "ཕལྐ་ལནྜ་གླིང་ཕྲན།",
			territory.FR:    "ཕ་རཱན་སི།",
			territory.GA:    "གེ་བཽན།",
			territory.GB:    "དབྱིན་ཇི་",
			territory.GD:    "གྷ་རི་ན་ཌ།",
			territory.GE:    "ཇོར་ཇི་ཡ།",
			territory.GH:    "གྷ་ན།",
			territory.GI:    "ཇིབ་རཱལ་ཊར།",
			territory.GM:    "གྷམ་བི་ཡ།",
			territory.GN:    "གྷི་ནི་ཡ།",
			territory.GR:    "གྷི་རཱི་སི།",
			territory.GT:    "གྷོ་ཊེ་མ་ལ།",
			territory.GW:    "གྷི་ནི་ཡ་བིས྄་སོ།",
			territory.GY:    "གྷུ་ཡཱ་ན།",
			territory.HK:    "ཧོང་ཀོང༌།",
			territory.HN:    "ཧོན་དུ་རས྄།",
			territory.HR:    "ཀུརོ་ཤི་ཡ།",
			territory.HT:    "ཧེ་ཏི།",
			territory.HU:    "ཧངྒ་རི།",
			territory.ID:    "ཨིན་ཌོ་ནེ་ཤི་ཡ།",
			territory.IE:    "ཨ་ཡར་ལནཌ།",
			territory.IL:    "ཨི་ཛ྄་རེལ།",
			territory.IN:    "རྒྱ་གར་",
			territory.IQ:    "ཨི་རག།",
			territory.IR:    "ཨི་རཱན།",
			territory.IS:    "ཨ་ཨི་སི་ལནད།",
			territory.IT:    "ཨི་ཀྲར་ལི་",
			territory.JM:    "ཛ་མེ་ཀ།",
			territory.JO:    "ཇོར་ཌན།",
			territory.JP:    "ཉི་ཧོང་",
			territory.KE:    "ཁེན་ཉི་ཡ།",
			territory.KG:    "ཁིར་གིཛ་སྟཱན།",
			territory.KH:    "ཀམ་བོ་ཌི་ཡ།",
			territory.KI:    "ཀི་རི་བཱ་ཏི།",
			territory.KN:    "སེནྚ། ཀིཊྚས྄། དང༌། ནེ་བིས྄།",
			territory.KR:    "ལྷོ་ཀོ་རི་ཡ།",
			territory.KW:    "ཀུ་ཝེད་རྒྱལ་ཁབ།",
			territory.KY:    "ཁེ་མེན་གླིང་ཕྲན།",
			territory.KZ:    "ཁ་ཛཱག་སྟཱན།",
			territory.LA:    "ལཱ་འོས།",
			territory.LB:    "ལེབ་ནོན།",
			territory.LC:    "སེནྚ། ལུ་ཤི་ཡ།",
			territory.LI:    "ལེག་ཏེན་ཚིན།",
			territory.LK:    "ཤྲཱི་ལངྐ་།",
			territory.LR:    "ལི་བེ་རི་ཡ།",
			territory.LS:    "ལེ་སོ་ཐོ།",
			territory.LT:    "ལི་ཐུ་ཨེ་ནི་ཡ།",
			territory.LU:    "ལཀ་ཛམ་བོརྒ།",
			territory.LV:    "ལཏ་བི་ཡ།",
			territory.LY:    "ལི་བི་ཡ།",
			territory.MA:    "མོ་རོ་ཀྐོ།",
			territory.MC:    "མོ་ན་ཀོ།",
			territory.MG:    "མ་དཱ་གྷསྐཱར།",
			territory.MH:    "མཱར་ཤལ་གླིང་ཕྲེན།",
			territory.ML:    "མ་ལི།",
			territory.MM:    "འབར་མ།",
			territory.MN:    "སོག་ཡུལ།",
			territory.MO:    "མེ་ཀའོ།",
			territory.MR:    "མཽ་རི་ཏ་ནི་ཡ།",
			territory.MT:    "མལ་ཊ།",
			territory.MU:    "མཽ་རིཤས྄།",
			territory.MV:    "མལ་དྭིབ།",
			territory.MW:    "མཱ་ལཱ་ཝི།",
			territory.MX:    "མེཀ་སི་ཀོ།",
			territory.MY:    "མ་ལེ་ཤི་ཡ།",
			territory.MZ:    "མོ་ཛམ་བིག།",
			territory.NA:    "ན་མི་བི་ཡ།",
			territory.NE:    "ནའི་ཇར།",
			territory.NG:    "ནཱའི་ཇི་རི་ཡ།",
			territory.NI:    "ནི་ཀ་ར་གུ་ཨ།",
			territory.NL:    "ཧའོ་ལན།",
			territory.NO:    "ནོར་ཝེ།",
			territory.NP:    "བལ་ཡུལ་",
			territory.NR:    "ནཽ་རུ།",
			territory.NU:    "ནིའུ་ཝ།",
			territory.NZ:    "ནིའུ་ཛི་ལན྄ཌ།",
			territory.OM:    "ཨོ་མན།",
			territory.PA:    "པ་ནཱ་མ།",
			territory.PE:    "པེ་རུ།",
			territory.PG:    "པ་པུ་ཨ། ནིའུ། གྷི་ནི།",
			territory.PH:    "ཕི་ལི་པིནས྄།",
			territory.PK:    "པཀི་སྟཱན།",
			territory.PL:    "པོ་ལནྜ།",
			territory.PT:    "པོར་ཏུ་གྷལ།",
			territory.PW:    "པ་ལཽ།",
			territory.PY:    "པཱ་ར་གེ།",
			territory.QA:    "ཀ་ཏཱར།",
			territory.RO:    "རོ་མཱ་ནིཡ།",
			territory.RS:    "སེར་བི་ཡ།",
			territory.RU:    "ཨུ་རུ་སུ་",
			territory.RW:    "རུ་ཝན་ཌ།",
			territory.SA:    "སཽ་དྷི་ཨ་རཱ་བི་ཡ།",
			territory.SB:    "སོ་ལོ་མོན། གླིང་ཕྲན་ཚོ་ཁག།",
			territory.SC:    "སཱ་ཤཻལ།",
			territory.SD:    "སུ་དཱན།",
			territory.SE:    "ཧྲུའི་ཏན།",
			territory.SG:    "སིངྒ་པུར།",
			territory.SH:    "སེནྚ། ཧེ་ལི་ན།",
			territory.SI:    "ས་ལཽ་ཝེ་ནི་ཡ།",
			territory.SK:    "ས་ལཽ་ཝཀྱ།",
			territory.SL:    "སེ་ཡར་ར། ལིའོན།",
			territory.SM:    "སན་མེ་རི་ནོ།",
			territory.SN:    "སེ་ནི་གྷལ།",
			territory.SO:    "སོ་མཱལི་ཡ།",
			territory.SR:    "སུ་རི་ནཱམ།",
			territory.ST:    "ས་འོ་ཏོད་མད། དང༌། པ྄རིན་སི་པེ།",
			territory.SV:    "ཨེལ། སཱལ་ཝ་ཌོར།",
			territory.TD:    "ཅཻཌ།",
			territory.TR:    "ཏུརཀི།",
			territory.TV:    "ཐུ་ཝ་ལུ།",
			territory.UA:    "ཡུ་ཀྲན།",
			territory.UG:    "ཡུ་གན་ཌ།",
			territory.US:    "ཨ་མེ་རི་ཀ།",
			territory.UY:    "ཨུ་རུ་གྷེ།",
			territory.UZ:    "ཨུཛ་བེ་ཀིསྟཱན།",
			territory.VA:    "ཝེ་ཊི་ཀན།",
			territory.VC:    "སེནྚ། ཝིན་སན། དང༌། གྷིརིན་ཌིན།",
			territory.VE:    "ཝེ་ནི་ཛུའེ་ལ།",
			territory.VN:    "བི་དི་ནམ།",
			territory.VU:    "ཝ་ནུ་ཨ་ཐུ།",
			territory.WS:    "ནུ་བ་ས་མོ་འ།",
			territory.YE:    "ཡེ་མེན།",
			territory.ZA:    "ལྷོ་ ཨཕྲི་ཀ།",
			territory.ZM:    "ཛམ་བི་ཡ།",
			territory.ZW:    "ཛིམ་བྷཱ་བེ།",
			territory.ZZ:    "མིའི་ཤེས་རྟོགས་མ་བྱུང་བའི་ཁོར་ཡུག",
		},
	}
}
