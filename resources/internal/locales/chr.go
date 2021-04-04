// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_chr() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "chr",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} ᎤᎾᎢ {0}", Long: "{1} ᎤᎾᎢ {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ᎤᏃ", Feb: "ᎧᎦ", Mar: "ᎠᏅ", Apr: "ᎧᏬ", May: "ᎠᏂ", Jun: "ᏕᎭ", Jul: "ᎫᏰ", Aug: "ᎦᎶ", Sep: "ᏚᎵ", Oct: "ᏚᏂ", Nov: "ᏅᏓ", Dec: "ᎥᏍ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Ꭴ", Feb: "Ꭷ", Mar: "Ꭰ", Apr: "Ꭷ", May: "Ꭰ", Jun: "Ꮥ", Jul: "Ꭻ", Aug: "Ꭶ", Sep: "Ꮪ", Oct: "Ꮪ", Nov: "Ꮕ", Dec: "Ꭵ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ᎤᏃᎸᏔᏅ", Feb: "ᎧᎦᎵ", Mar: "ᎠᏅᏱ", Apr: "ᎧᏬᏂ", May: "ᎠᏂᏍᎬᏘ", Jun: "ᏕᎭᎷᏱ", Jul: "ᎫᏰᏉᏂ", Aug: "ᎦᎶᏂ", Sep: "ᏚᎵᏍᏗ", Oct: "ᏚᏂᏅᏗ", Nov: "ᏅᏓᏕᏆ", Dec: "ᎥᏍᎩᏱ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ᏆᏍᎬ", Mon: "ᏉᏅᎯ", Tue: "ᏔᎵᏁ", Wed: "ᏦᎢᏁ", Thu: "ᏅᎩᏁ", Fri: "ᏧᎾᎩ", Sat: "ᏈᏕᎾ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Ꮖ", Mon: "Ꮙ", Tue: "Ꮤ", Wed: "Ꮶ", Thu: "Ꮕ", Fri: "Ꮷ", Sat: "Ꭴ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ᏍᎬ", Mon: "ᏅᎯ", Tue: "ᏔᎵ", Wed: "ᏦᎢ", Thu: "ᏅᎩ", Fri: "ᏧᎾ", Sat: "ᏕᎾ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ᎤᎾᏙᏓᏆᏍᎬ", Mon: "ᎤᎾᏙᏓᏉᏅᎯ", Tue: "ᏔᎵᏁᎢᎦ", Wed: "ᏦᎢᏁᎢᎦ", Thu: "ᏅᎩᏁᎢᎦ", Fri: "ᏧᎾᎩᎶᏍᏗ", Sat: "ᎤᎾᏙᏓᏈᏕᎾ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ᏌᎾᎴ", PM: "ᏒᎯᏱᎢ"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ᏌᎾᎴ", PM: "ᏒᎯᏱᎢ"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ᏌᎾᎴ", PM: "ᏒᎯᏱᎢᏗᏢ"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "ᏌᏊ ᎢᏳᎾᎵᏍᏔᏅ ᎡᎳᏈ ᎢᎹᎵᏘᏏ ᎠᏕᎳ", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "ᎠᏫᎨᏂᏍᏖᏂ ᎠᏕᎳ", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "ᎠᎵᏇᏂᏯ ᎠᏕᎳ", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "ᎠᎵᎻᏂᎠ ᎠᏕᎳ", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "ᎾᏍᎩᏁᏛᎳᏂ ᎠᏂᏘᎵᏏ ᎠᏕᎳ", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "ᎠᏂᎪᎳ ᎠᏕᎳ", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "ᎠᏥᏂᏘᏂᎠ ᎠᏕᎳ", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ᎡᎳᏗᏜ ᎠᏕᎳ", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "ᎠᎷᏆ ᎠᏕᎳ", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "ᎠᏏᎵᏆᏌᏂ ᎠᏕᎳ", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "ᏉᏏᏂᎠ-ᎲᏤᎪᏫ ᎦᏁᏟᏴᏍᏔᏅ ᎠᏕᎳ", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "ᏆᏇᏙᏍ ᎠᏕᎳ", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "ᏆᏂᎦᎵᏕᏍ ᎠᏕᎳ", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "ᏊᎵᎨᎵᎠ ᎠᏕᎳ", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "ᏆᎭᎴᎢᏂ ᎠᏕᎳ", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "ᏋᎷᏂᏗ ᎠᏕᎳ", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "ᏆᏊᏓ ᎠᏕᎳ", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ᏊᎾᎢ ᎠᏕᎳ", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "ᏉᎵᏫᎠ ᎠᏕᎳ", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "ᏆᏏᎵᎢ ᎠᏕᎳ", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "ᏆᎭᎹ ᎠᏕᎳ", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ᏊᏔᏂ ᎠᏕᎳ", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "ᏆᏣᏩᎾ ᎠᏕᎳ", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "ᏇᎳᎷᏍ ᎠᏕᎳ", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "ᏇᎳᎷᏍ ᎠᏕᎳ (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "ᏇᎵᏍ ᎠᏕᎳ", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "ᎨᎾᏓ ᎠᏕᎳ", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "ᎧᏂᎪ ᎠᏕᎳ", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "ᏍᏫᏏ ᎠᏕᎳ", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "ᏥᎵ ᎠᏕᎳ", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "ᏣᏂᏏ ᎠᏕᎳ (ᏓᎹᏳᏟᏗ)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "ᏓᎶᏂᎨ ᎠᏕᎳ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "ᎪᎸᎻᏈᎢᎠ ᎠᏕᎳ", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "ᎪᏍᏓᎵᎧ ᎠᏕᎳ", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "ᎫᏆ ᎦᏁᏟᏴᏍᏔᏅ ᎠᏕᎳ", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "ᎫᏆ ᎠᏕᎳ", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "ᎢᎬᎾᏕᎾ ᎢᏤᏳᏍᏗ ᎠᏕᎳ", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "ᏤᎩ ᎠᏕᎳ", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "ᏥᏊᏗ ᎠᏕᎳ", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ᏕᏂᏍ ᎠᏕᎳ", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ᏙᎻᏂᎧᏂ ᎠᏕᎳ", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "ᎠᎵᏥᎵᏯ ᎠᏕᎳ", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "ᎢᏥᏈᎢ ᎠᏕᎳ", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "ᎡᎵᏟᏯ ᎠᏕᎳ", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ᎢᏗᎣᏈᎠ ᎠᏕᎳ", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "ᏳᎳᏛ ᎠᏕᎳ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "ᏫᎩ ᎠᏕᎳ", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ᏩᎩᎤ ᏚᎦᏚᏛᎢ ᎠᏕᎳ", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "ᎩᎵᏏᏲ ᎠᏕᎳ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "ᏣᎠᏥᎢ ᎠᏕᎳ", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "ᎦᎠᎾ ᎠᏕᎳ", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "ᏥᏆᎵᏓ ᎠᏕᎳ", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "ᎦᎹᏈᎢᎠ ᎠᏕᎳ", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "ᎩᎢᏂ ᎠᏕᎳ", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ᏆᏖᎹᎳ ᎠᏕᎳ", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "ᎦᏯᎾ ᎠᏕᎳ", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "ᎰᏂᎩ ᎪᏂᎩ ᎠᏕᎳ", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "ᎭᏂᏚᎳᏍ ᎠᏕᎳ", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "ᎧᎶᎡᏏᎠ ᎠᏕᎳ", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "ᎮᏘ ᎠᏕᎳ", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ᎲᏂᎦᎵ ᎠᏕᎳ", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "ᎢᏂᏙᏂᏍᏯ ᎠᏕᎳ", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "ᎢᏏᎵᏱ ᎢᏤ ᎠᏕᎳ", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ᎢᏂᏗᎢᎠ ᎠᏕᎳ", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ᎢᎳᎩ ᎠᏕᎳ", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ᎢᎴᏂ ᎠᏕᎳ", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "ᏧᏁᏍᏓᎸᎯ ᎠᏕᎳ", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "ᏣᎺᎢᎧ ᎠᏕᎳ", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "ᏦᏓᏂ ᎠᏕᎳ", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "ᏣᏩᏂᏏ ᎠᏕᎳ", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ᎨᏂᏯ ᎠᏕᎳ", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "ᎩᎵᏣᎢᏍ ᎠᏕᎳ", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "ᎧᎹᏉᏗᎠᏂ ᎠᏕᎳ", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "ᎪᎼᎳᏍ ᎠᏕᎳ", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "ᏧᏴᏢ ᎪᎵᎠ ᎠᏕᎳ", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "ᏧᎦᎾᏮ ᎪᎵᎠ ᎠᏕᎳ", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "ᎫᏪᎢᏘ ᎠᏕᎳ", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "ᎨᎢᎹᏂ ᏚᎦᏚᏛᎢ ᎠᏕᎳ", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "ᎧᏎᎧᏍᏕᏂ ᎠᏕᎳ", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "ᎳᎣ ᎠᏕᎳ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "ᎴᏆᎾᏂ ᎠᏕᎳ", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "ᏍᎵ ᎳᏂᎧ ᎠᏕᎳ", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "ᎳᏈᎵᏯ ᎠᏕᎳ", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "ᎵᏈᏯ ᎠᏕᎳ", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "ᎼᎶᎪ ᎠᏕᎳ", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "ᎹᎵᏙᏫᎠ ᎠᏕᎳ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ᎹᎳᎦᏏ ᎠᏕᎳ", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "ᎹᏎᏙᏂᎠ ᎠᏕᎳ", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "ᎹᏯᎹᎵ ᎠᏕᎳ", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "ᎹᏂᎪᎵᎠ ᎠᏕᎳ", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "ᎹᎧᎣ ᎠᏕᎳ", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "ᎹᏈᏔᏂᎠ ᎠᏕᎳ (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "ᎹᏈᏔᏂᎠ ᎠᏕᎳ", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "ᎹᏘᎢᏯ ᎠᏕᎳ", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "ᎹᎵᏗᏫᏍ ᎠᏕᎳ", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "ᎹᎳᏫ ᎠᏕᎳ", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "ᏍᏆᏂ ᎠᏕᎳ", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "ᎹᎴᏏᎢᎠ ᎠᏕᎳ", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "ᎼᏎᎻᏇᎩ ᎠᏕᎳ", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "ᎾᎻᏈᎢᏯ ᎠᏕᎳ", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "ᏂᏥᎵᏯ ᎠᏕᎳ", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "ᏂᎧᎳᏆ ᎠᏕᎳ", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "ᏃᏪ ᎠᏕᎳ", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "ᏁᏆᎵ ᎠᏕᎳ", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "ᎢᏤ ᏏᎢᎴᏂᏗ ᎠᏕᎳ", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ᎣᎺᏂ ᎠᏕᎳ", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "ᏆᎾᎹ ᎠᏕᎳ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "ᏇᎷ ᎠᏕᎳ", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "ᏆᏇ ᎢᏤ ᎩᎢᏂ ᎠᏕᎳ", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ᎠᏂᏈᎵᎩᏃ ᎠᏕᎳ", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "ᏆᎩᏍᏖᏂ ᎠᏕᎳ", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "ᏉᎳᏂ ᎠᏕᎳ", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "ᏆᎳᏇᎢᏯ ᎠᏕᎳ", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "ᎧᏔᎵ ᎠᏕᎳ", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "ᎶᎹᏂᏯ ᎠᏕᎳ", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "ᏒᏈᏯ ᎠᏕᎳ", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "ᏲᏂᎢ ᎠᏕᎳ", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ᎶᏩᏂᏓ ᎠᏕᎳ", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "ᏌᎤᏗ ᎠᏕᎳ", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "ᏐᎶᎹᏂ ᏚᎦᏚᏛᎢ ᎠᏕᎳ", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "ᏏᎡᏥᎵᏍ ᎠᏕᎳ", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "ᏑᏕᏂ ᎠᏕᎳ", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "ᏍᏫᏕᏂ ᎠᏕᎳ", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "ᏏᏂᎦᏉᎵ ᎠᏕᎳ", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "ᎤᏓᏅᏘ ᎮᎵᎾ ᎠᏕᎳ", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "ᏏᎡᎳᎴᎣᏂ ᎠᏕᎳ", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "ᏐᎹᎵ ᎠᏕᎳ", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "ᏒᎵᎾᎻ ᎠᏕᎳ", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "ᏧᎦᎾᏮ ᏑᏕᏂ ᎠᏕᎳ", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "ᏌᎣᏙᎺ ᎠᎴ ᏈᏂᏏᏇ ᎠᏕᎳ (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "ᏌᎣᏙᎺ & ᏈᏂᏏᏇ ᎠᏕᎳ", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "ᏏᎵᎠ ᎠᏕᎳ", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "ᏍᏩᏏ ᎠᏕᎳ", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "ᏔᏯᎴᏂ ᎠᏕᎳ", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "ᏔᏥᎩᏍᏕᏂ ᎠᏕᎳ", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "ᏛᎵᎩᎺᏂᏍᏔᏂ ᎠᏕᎳ", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "ᏚᏂᏏᏍᎠ ᎠᏕᎳ", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "ᏔᏂᎪ ᎠᏕᎳ", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "ᎬᏃ ᎠᏕᎳ", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ᏟᏂᏕᏗ & ᏙᏆᎪ ᎠᏕᎳ", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "ᎢᏤ ᏔᎢᏩᏂ ᎠᏕᎳ", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "ᏖᏂᏏᏂᏯ ᎠᏕᎳ", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ᏳᎧᎴᏂ ᎠᏕᎳ", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "ᏳᎦᏂᏓ ᎠᏕᎳ", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "US ᎠᏕᎳ", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "ᏳᎷᏇ ᎠᏕᎳ", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "ᎤᏍᏇᎩᏍᏖᏂ ᎠᏕᎳ", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "ᏪᏁᏑᏪ ᎠᏕᎳ (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "ᏪᏁᏑᏪ ᎠᏕᎳ", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "ᏫᎡᏘᎾᎻᏍ ᎠᏕᎳ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "ᏩᏂᎤᏩᏚ ᎠᏕᎳ", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "ᏌᎼᎠ ᎠᏕᎳ", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "ᎠᏰᏟ ᎬᎿᎨᏍᏛ CFA ᎠᏕᎳ", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "ᏗᎧᎸᎬ ᎨᏆᏙᏯ ᎠᏕᎳ", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "ᏭᏕᎵᎬ ᎬᎿᎨᏍᏛ CFA ᎠᏕᎳ", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP ᎠᏕᎳ", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "ᏄᏬᎵᏍᏛᎾ ᎠᏕᎳ", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "ᏰᎺᏂ ᎠᏕᎳ", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "ᏧᎦᎾᏮ ᎬᎿᎨᏍᏛ ᎠᏕᎳ", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "ᏏᎻᏆᏇ ᎠᏕᎳ", Symbol: "ZMW"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:      "ᎠᏩᎳ",
			language.AB:      "ᎠᏆᏏᎠᏂ",
			language.ACE:     "ᎠᏥᏂᏏ",
			language.ADA:     "ᎠᏓᎾᎦᎺ",
			language.ADY:     "ᎠᏗᎨ",
			language.AF:      "ᎠᎬᎿᎨᏍᏛ",
			language.AGQ:     "ᎠᎨᎹ",
			language.AIN:     "ᎠᏱᏄ",
			language.AK:      "ᎠᎧᎾ",
			language.ALE:     "ᎠᎵᎤᏘ",
			language.ALT:     "ᏧᎦᎾᏮ ᏗᏜ ᎠᎵᏔᎢ",
			language.AM:      "ᎠᎹᎭᎵᎩ",
			language.AN:      "ᎠᏩᎪᏂᏏ",
			language.ANP:     "ᎠᎾᎩᎧ",
			language.AR:      "ᎡᎳᏈ",
			language.AR_001:  "ᎪᎯᏊ ᎢᎬᏥᎩ ᎠᏟᎶᏍᏗ ᎡᎳᏈ",
			language.ARN:     "ᎹᏊᏤ",
			language.ARP:     "ᎠᏩᏈᎰ",
			language.AS:      "ᎠᏌᎻᏏ",
			language.ASA:     "ᎠᏑ",
			language.AST:     "ᎠᏍᏚᎵᎠᏂ",
			language.AV:      "ᎠᏩᎵᎧ",
			language.AWA:     "ᎠᏩᏗ",
			language.AY:      "ᎠᏱᎹᎳ",
			language.AZ:      "ᎠᏎᎵ",
			language.BA:      "ᏆᏍᎯᎩᎠ",
			language.BAN:     "ᏆᎵᏁᏏ",
			language.BAS:     "ᏆᏌᎠ",
			language.BE:      "ᏇᎳᎷᏏ",
			language.BEM:     "ᏇᎹᏆ",
			language.BEZ:     "ᏇᎾ",
			language.BG:      "ᏊᎵᎨᎵᎠᏂ",
			language.BHO:     "ᏉᏣᏊᎵ",
			language.BI:      "ᏈᏍᎳᎹ",
			language.BIN:     "ᏈᏂ",
			language.BLA:     "ᏏᎩᏏᎧ",
			language.BM:      "ᏆᎻᏆᎳ",
			language.BN:      "ᏇᏂᎦᎳ",
			language.BO:      "ᏘᏇᏔᏂ",
			language.BR:      "ᏇᏙᏂ",
			language.BRX:     "ᏉᏙ",
			language.BS:      "ᏆᏍᏂᎠᏂ",
			language.BUG:     "ᏈᎥᎩᏂᏍ",
			language.BYN:     "ᏟᏂ",
			language.CA:      "ᎨᏔᎳᏂ",
			language.CAY:     "ᎧᏳᎦ",
			language.CCP:     "ᏣᎧᎹ",
			language.CE:      "ᏤᏤᏂ",
			language.CEB:     "ᏎᏆᏃ",
			language.CGG:     "ᏥᎦ",
			language.CH:      "ᏣᎼᎶ",
			language.CHK:     "ᏧᎨᏎ",
			language.CHM:     "ᎹᎵ",
			language.CHO:     "ᎠᏣᏓ",
			language.CHR:     "ᏣᎳᎩ",
			language.CHY:     "ᏣᏰᏂ",
			language.CKB:     "ᎠᏰᏟ ᎫᏗᏏ",
			language.CO:      "ᎪᎵᏍᎢᎧᏂ",
			language.CRS:     "ᏎᏎᎵᏩ ᏟᏲᎵ ᎠᏂᎦᎸ",
			language.CS:      "ᏤᎩ",
			language.CU:      "ᏧᏂᎳᏫᏍᏗ ᏍᎳᏫᎪ",
			language.CV:      "ᏧᏩᏏ",
			language.CY:      "ᏪᎵᏏ",
			language.DA:      "ᏕᏂᏍ",
			language.DAK:     "ᏓᎪᏔ",
			language.DAR:     "ᏓᎳᏆ",
			language.DAV:     "ᏔᎢᏔ",
			language.DE:      "ᏙᎢᏥ",
			language.DE_AT:   "ᎠᏟᏯᏂ ᎠᏂᏓᏥ",
			language.DE_CH:   "ᏍᏫᏏ ᎦᎸᎳᏗ ᎠᏂᏓᏥ",
			language.DGR:     "ᎩᏟ ᎤᏄᎳᏥ",
			language.DJE:     "ᏌᎹ",
			language.DSB:     "ᎡᎳᏗ ᏐᏈᎠᏂ",
			language.DUA:     "ᏚᎠᎳ",
			language.DV:      "ᏗᏪᎯ",
			language.DYO:     "ᏦᎳ-ᏬᏱ",
			language.DZ:      "ᏓᏐᏅᎧ",
			language.DZG:     "ᏓᏌᎦ",
			language.EBU:     "ᎡᎻᏊ",
			language.EE:      "ᎡᏪ",
			language.EFI:     "ᎡᏫᎩ",
			language.EKA:     "ᎨᎧᏧᎧ",
			language.EL:      "ᎠᏂᎪᎢ",
			language.EN:      "ᎩᎵᏏ",
			language.EN_AU:   "ᎡᎳᏗᏜ ᎩᎵᏏ",
			language.EN_CA:   "ᎨᎾᏓ ᎩᎵᏏ",
			language.EN_GB:   "UK ᎩᎵᏏ",
			language.EN_US:   "US ᎩᎵᏏ",
			language.EO:      "ᎡᏍᏇᎳᏂᏙ",
			language.ES:      "ᏍᏆᏂ",
			language.ES_419:  "ᏔᏘᏂ ᎠᎹᏰᏟ ᏍᏆᏂ",
			language.ES_ES:   "ᎠᏂᏍᏆᏂᏱ ᏍᏆᏂ",
			language.ES_MX:   "ᏍᏆᏂᏱ ᏍᏆᏂ",
			language.ET:      "ᎡᏍᏙᏂᎠᏂ",
			language.EU:      "ᏆᏍᎨ",
			language.EWO:     "ᎡᏬᏂᏙ",
			language.FA:      "ᏇᏏᎠᏂ",
			language.FA_AF:   "ᏓᎵ",
			language.FF:      "ᏊᎳᏂ",
			language.FI:      "ᏈᏂᏍ",
			language.FIL:     "ᎠᏈᎵᎩ",
			language.FJ:      "ᏫᏥᎠᏂ",
			language.FO:      "ᏇᎶᎡᏍ",
			language.FON:     "ᏠᏂ",
			language.FR:      "ᎦᎸᏥ",
			language.FR_CA:   "ᎨᎾᏓ ᎦᎸᏥ",
			language.FR_CH:   "ᏍᏫᏏ ᎦᎸᏥ",
			language.FUR:     "ᏞᎤᎵᎠᏂ",
			language.FY:      "ᏭᏕᎵᎬ ᏗᏜ ᏟᏏᎠᏂ",
			language.GA:      "ᎨᎵᎩ",
			language.GAA:     "Ꭶ",
			language.GD:      "ᏍᎦᏗ ᎨᎵᎩ",
			language.GEZ:     "ᎩᏏ",
			language.GIL:     "ᎩᏇᏘᏏ",
			language.GL:      "ᎦᎵᏏᎠᏂ",
			language.GN:      "ᏆᎳᏂ",
			language.GOR:     "ᎪᎶᏂᏔᏃ",
			language.GSW:     "ᏍᏫᏏ ᎠᏂᏓᏥ",
			language.GU:      "ᎫᏣᎳᏘ",
			language.GUZ:     "ᎫᏏ",
			language.GV:      "ᎹᎾᎧᏏ",
			language.GWI:     "ᏈᏥᏂ",
			language.HA:      "ᎭᎤᏌ",
			language.HAW:     "ᎭᏩᎼ",
			language.HE:      "ᎠᏂᏈᎷ",
			language.HI:      "ᎯᏂᏗ",
			language.HIL:     "ᎯᎵᎨᎾᏂ",
			language.HMN:     "ᎭᎼᏂᎩ",
			language.HR:      "ᎧᎶᎡᏏᏂ",
			language.HSB:     "ᎦᎸᎳᏗᎨ ᏐᏈᎠᏂ",
			language.HT:      "ᎮᏏᎠᏂ ᏟᏲᎵ",
			language.HU:      "ᎲᏂᎦᎵᎠᏂ",
			language.HUP:     "ᎠᏂᎱᏆ",
			language.HY:      "ᎠᎳᎻᎠᏂ",
			language.HZ:      "ᎮᎴᎶ",
			language.IA:      "ᎠᏰᏟ ᎦᏬᏂᎯᏍᏗ",
			language.IBA:     "ᎢᏆᏂ",
			language.IBB:     "ᎢᏈᏈᎣ",
			language.ID:      "ᎢᏂᏙᏂᏏᎠ",
			language.IG:      "ᎢᎦᎪ",
			language.II:      "ᏏᏧᏩᏂ Ᏹ",
			language.ILO:     "ᎢᎶᎪ",
			language.INH:     "ᎢᏂᎫᏏ",
			language.IO:      "ᎢᏙ",
			language.IS:      "ᏧᏁᏍᏓᎸᎯᎢᎩ",
			language.IT:      "ᎬᏩᎵᏲᏥᎢ",
			language.IU:      "ᎢᏄᎦᏘᏚ",
			language.JA:      "ᏣᏩᏂᏏ",
			language.JBO:     "ᎶᏣᏆᏂ",
			language.JGO:     "ᎾᎪᏆ",
			language.JMC:     "ᎹᏣᎺ",
			language.JV:      "ᏆᏌ ᏣᏩ",
			language.KA:      "ᏦᏥᎠᏂ",
			language.KAB:     "ᎧᏈᎴ",
			language.KAC:     "ᎧᏥᏂ",
			language.KAJ:     "ᏥᏧ",
			language.KAM:     "ᎧᎻᏆ",
			language.KBD:     "ᎧᏆᏗᎠᏂ",
			language.KCG:     "ᏔᏯᏆ",
			language.KDE:     "ᎹᎪᏕ",
			language.KEA:     "ᎧᏊᏪᏗᎠᏄ",
			language.KFO:     "ᎪᎶ",
			language.KHA:     "ᎧᏏ",
			language.KHQ:     "ᎪᏱᎳ ᏥᏂ",
			language.KI:      "ᎩᎫᏳ",
			language.KJ:      "ᎫᏩᏂᎠᎹ",
			language.KK:      "ᎧᏌᎧ",
			language.KKJ:     "ᎧᎪ",
			language.KL:      "ᎧᎳᎵᏑᏘ",
			language.KLN:     "ᎧᎴᏂᏥᏂ",
			language.KM:      "ᎩᎻᎷ",
			language.KMB:     "ᎩᎻᏊᏚ",
			language.KN:      "ᎧᎾᏓ",
			language.KO:      "ᎪᎵᎠᏂ",
			language.KOK:     "ᎧᏂᎧᏂ",
			language.KPE:     "ᏇᎴ",
			language.KR:      "ᎧᏄᎵ",
			language.KRC:     "ᎧᎳᏣᏱ-ᏆᎵᎧᎵ",
			language.KRL:     "ᎧᎴᎵᎠᏂ",
			language.KRU:     "ᎫᎷᎩ",
			language.KS:      "ᎧᏏᎻᎵ",
			language.KSB:     "ᏝᎻᏆᎸ",
			language.KSF:     "ᏆᏫᎠ",
			language.KSH:     "ᎪᎶᏂᎠᏂ",
			language.KU:      "ᎫᏗᏏ",
			language.KUM:     "ᎫᎻᎧ",
			language.KV:      "ᎪᎻ",
			language.KW:      "ᏎᎷᎭ",
			language.KY:      "ᎩᎵᏣᎢᏍ",
			language.LA:      "ᎳᏘᏂ",
			language.LAD:     "ᎳᏗᏃ",
			language.LAG:     "ᎳᏂᎩ",
			language.LB:      "ᎸᎦᏏᎻᏋᎢᏍ",
			language.LEZ:     "ᎴᏏᎦᏂ",
			language.LG:      "ᎦᏂᏓ",
			language.LI:      "ᎴᎹᏊᎵᏏ",
			language.LKT:     "ᎳᎪᏓ",
			language.LN:      "ᎵᏂᎦᎳ",
			language.LO:      "ᎳᎣ",
			language.LOZ:     "ᎶᏏ",
			language.LRC:     "ᏧᏴᏢ ᏗᏜ ᎷᎵ",
			language.LT:      "ᎵᏚᏩᏂᎠᏂ",
			language.LU:      "ᎷᏆ-ᎧᏔᎦ",
			language.LUA:     "ᎷᏆ-ᎷᎷᎠ",
			language.LUN:     "ᎷᎾᏓ",
			language.LUO:     "ᎷᎣ",
			language.LUS:     "ᎻᏐ",
			language.LUY:     "ᎷᏱᎠ",
			language.LV:      "ᎳᏘᏫᎠᏂ",
			language.MAD:     "ᎹᏚᎴᏏ",
			language.MAG:     "ᎹᎦᎯ",
			language.MAI:     "ᎹᏟᎵ",
			language.MAK:     "ᎹᎧᏌ",
			language.MAS:     "ᎹᏌᏱ",
			language.MDF:     "ᎼᎧᏌ",
			language.MEN:     "ᎺᎾᏕ",
			language.MER:     "ᎺᎷ",
			language.MFE:     "ᎼᎵᏏᎡᏂ",
			language.MG:      "ᎹᎳᎦᏏ",
			language.MGH:     "ᎹᎫᏩ-ᎻᏙ",
			language.MGO:     "ᎺᎳ’",
			language.MH:      "ᎹᏌᎵᏏ",
			language.MI:      "ᎹᏫ",
			language.MIC:     "ᎻᎧᎹᎩ",
			language.MIN:     "ᎻᎾᎧᏆᎤ",
			language.MK:      "ᎹᏎᏙᏂᎠᏂ",
			language.ML:      "ᎹᎳᏯᎳᎻ",
			language.MN:      "ᎹᏂᎪᎵᎠᏂ",
			language.MNI:     "ᎺᏂᏉᎵ",
			language.MOH:     "ᎼᎭᎩ",
			language.MOS:     "ᎼᏍᏏ",
			language.MR:      "ᎹᎳᏘ",
			language.MS:      "ᎹᎴ",
			language.MT:      "ᎹᎵᏘᏍ",
			language.MUA:     "ᎽᏂᏓᎩ",
			language.MUL:     "ᏧᏈᏍᏗ ᏗᎦᏬᏂᎯᏍᏗ",
			language.MUS:     "ᎠᎫᏌ",
			language.MWL:     "ᎻᎳᏕᏏ",
			language.MY:      "ᏋᎻᏍ",
			language.MYV:     "ᎡᏏᏯ",
			language.MZN:     "ᎹᏌᏕᎳᏂ",
			language.NA:      "ᏃᎤᎷ",
			language.NAP:     "ᏂᏯᏆᎵᏔᏂ",
			language.NAQ:     "ᎾᎹ",
			language.NB:      "ᏃᎵᏪᏥᏂ ᏉᎧᎹᎵ",
			language.ND:      "ᏧᏴᏢ ᏂᏕᏇᎴ",
			language.NDS:     "ᎡᎳᏗ ᎠᏂᏓᏥ",
			language.NDS_NL:  "ᎡᎳᏗ ᏁᏛᎳᏂ",
			language.NE:      "ᏁᏆᎵ",
			language.NEW:     "ᏁᏩᎵ",
			language.NG:      "ᎾᏙᎦ",
			language.NIA:     "ᏂᎠᏏ",
			language.NIU:     "ᏂᏳᏫᏯᏂ",
			language.NL:      "ᏛᏥ",
			language.NL_BE:   "ᏊᎵᏥᎥᎻ ᏛᏥ",
			language.NMG:     "ᏆᏏᏲ",
			language.NN:      "ᏃᎵᏪᏥᏂ ᎾᎵᏍᎩ",
			language.NNH:     "ᎾᏥᏰᎹᏊᏂ",
			language.NOG:     "ᏃᎦᏱ",
			language.NQO:     "ᎾᎪ",
			language.NR:      "ᏧᎦᎾᏮ ᏂᏕᏇᎴ",
			language.NSO:     "ᏧᏴᏢ ᏗᏜ ᏐᏠ",
			language.NUS:     "ᏄᏪᎵ",
			language.NV:      "ᎾᏩᎰ",
			language.NY:      "ᏂᏯᏂᏣ",
			language.NYN:     "ᏂᏯᎾᎪᎴ",
			language.OC:      "ᎠᏏᏔᏂ",
			language.OM:      "ᎣᎶᎼ",
			language.OR:      "ᎣᏗᎠ",
			language.OS:      "ᎣᏎᏘᎧ",
			language.PA:      "ᏡᏂᏣᏈ",
			language.PAG:     "ᏇᎦᏏᎠᏂ",
			language.PAM:     "ᏆᎹᏆᎾᎦ",
			language.PAP:     "ᏆᏈᏯᎺᎾᏙ",
			language.PAU:     "ᏆᎳᎤᏩᏂ",
			language.PCM:     "ᎾᎩᎵᎠᏂ ᏈᏥᏂ",
			language.PL:      "ᏉᎵᏍ",
			language.PRG:     "ᏡᏏᎠᏂ",
			language.PS:      "ᏆᏍᏙ",
			language.PT:      "ᏉᏧᎩᏍ",
			language.PT_BR:   "ᏆᏏᎵᎢ ᏉᏧᎩᏍ",
			language.PT_PT:   "ᏳᎳᏈ ᏉᏧᎩᏍ",
			language.QU:      "ᎨᏧᏩ",
			language.QUC:     "ᎩᏤ",
			language.RAP:     "ᎳᏆᏄᏫ",
			language.RAR:     "ᎳᎶᏙᎾᎦᏂ",
			language.RM:      "ᎠᏂᎶᎺᏂ",
			language.RN:      "ᎷᏂᏗ",
			language.RO:      "ᎶᎹᏂᎠᏂ",
			language.RO_MD:   "ᎹᎵᏙᏫᎠ ᏣᎹᏂᎠᏂ",
			language.ROF:     "ᎶᎹᏉ",
			language.ROOT:    "ᎤᎾᏍᎦᎸ",
			language.RU:      "ᏲᏅᎯ",
			language.RUP:     "ᎠᏬᎹᏂᎠᏂ",
			language.RW:      "ᎩᏂᏯᏩᏂᏓ",
			language.RWK:     "Ꮖ",
			language.SA:      "ᏍᏂᏍᎩᏗ",
			language.SAD:     "ᏌᏅᏓᏫ",
			language.SAH:     "ᏌᎧᎾ",
			language.SAQ:     "ᏌᎹᏊᎷ",
			language.SAT:     "ᏌᏂᏔᎵ",
			language.SBA:     "ᎾᎦᎹᏇ",
			language.SBP:     "ᏌᏁᎫ",
			language.SC:      "ᏌᏗᏂᎠᏂ",
			language.SCN:     "ᏏᏏᎵᎠᏂ",
			language.SCO:     "ᏍᎦᏗ",
			language.SD:      "ᏏᏂᏗ",
			language.SE:      "ᏧᏴᏢ ᏗᏜ ᏌᎻ",
			language.SEE:     "ᏏᏂᎦ",
			language.SEH:     "ᏎᎾ",
			language.SES:     "ᎪᏱᎳᏈᎶ ᏎᏂ",
			language.SG:      "ᏌᏂᎪ",
			language.SHI:     "ᏔᏤᎵᎯᏘ",
			language.SHN:     "ᏝᏂ",
			language.SI:      "ᏏᎾᎭᎳ",
			language.SK:      "ᏍᎶᏩᎩ",
			language.SL:      "ᏍᎶᏫᏂᎠᏂ",
			language.SM:      "ᏌᎼᏯᏂ",
			language.SMA:     "ᏧᎦᎾᏮ ᏗᏜ ᏌᎻ",
			language.SMJ:     "ᎷᎴ ᏌᎻ",
			language.SMN:     "ᎢᎾᎵ ᏌᎻ",
			language.SMS:     "ᏍᎪᎵᏘ ᏌᎻ",
			language.SN:      "ᏠᎾ",
			language.SNK:     "ᏐᏂᏂᎨ",
			language.SO:      "ᏐᎹᎵ",
			language.SQ:      "ᎠᎵᏇᏂ",
			language.SR:      "ᏒᏈᎠᏂ",
			language.SRN:     "ᏏᎳᎾᏂ ᏙᏃᎪ",
			language.SS:      "ᏍᏩᏘ",
			language.SSY:     "ᏌᎰ",
			language.ST:      "ᏧᎦᎾᏮ ᏗᏜ ᏐᏠ",
			language.SU:      "ᏑᏂᏓᏂᏏ",
			language.SUK:     "ᏑᎫᎹ",
			language.SV:      "ᏍᏫᏗᏏ",
			language.SW:      "ᏍᏩᎯᎵ",
			language.SW_CD:   "ᎧᏂᎪ ᏍᏩᎯᎵ",
			language.SWB:     "ᎪᎼᎵᎠᏂ",
			language.SYR:     "ᏏᎵᎠᎩ",
			language.TA:      "ᏔᎻᎵ",
			language.TE:      "ᏖᎷᎦ",
			language.TEM:     "ᏘᎹᏁ",
			language.TEO:     "ᏖᏐ",
			language.TET:     "ᏖᏚᎼ",
			language.TG:      "ᏔᏥᎩ",
			language.TH:      "ᏔᏱ",
			language.TI:      "ᏘᎩᎵᏂᎠ",
			language.TIG:     "ᏢᏓᏥ",
			language.TK:      "ᎠᏂᎬᎾ",
			language.TLH:     "ᏟᎦᎾ",
			language.TN:      "ᏧᏩᎾ",
			language.TO:      "ᏙᎾᎦᏂ",
			language.TPI:     "ᏙᎩ ᏈᏏᏂ",
			language.TR:      "ᎠᎬᎾ",
			language.TRV:     "ᏔᎶᎪ",
			language.TS:      "ᏦᎾᎦ",
			language.TT:      "ᏔᏔ",
			language.TUM:     "ᏛᎹᏊᎧ",
			language.TVL:     "ᏚᏩᎷ",
			language.TWQ:     "ᏔᏌᏩᎩ",
			language.TY:      "ᏔᎯᏘᎠᏂ",
			language.TYV:     "ᏚᏫᏂᎠᏂ",
			language.TZM:     "ᎠᏰᏟ ᎡᎶᎯ ᏓᏟᎶᏍᏗᏓᏅᎢ ᏔᎹᏏᏘ",
			language.UDM:     "ᎤᏚᎷᏘ",
			language.UG:      "ᏫᎦ",
			language.UK:      "ᏳᎧᎴᏂᎠᏂ",
			language.UMB:     "ᎤᎹᏊᏅᏚ",
			language.UND:     "ᏄᏬᎵᏍᏛᎾ ᎦᏬᏂᎯᏍᏗ",
			language.UR:      "ᎤᎵᏚ",
			language.UZ:      "ᎤᏍᏇᎩ",
			language.VAI:     "ᏩᏱ",
			language.VE:      "ᏫᏂᏓ",
			language.VI:      "ᏫᎡᏘᎾᎻᏍ",
			language.VO:      "ᏬᎳᏊᎩ",
			language.VUN:     "ᏭᎾᏦ",
			language.WA:      "ᏩᎷᎾ",
			language.WAE:     "ᏩᎵᏎᎵ",
			language.WAL:     "ᏬᎳᏱᏔ",
			language.WAR:     "ᏩᎴ",
			language.WO:      "ᏬᎶᏫ",
			language.XAL:     "ᎧᎳᎻᎧ",
			language.XH:      "ᏠᏌ",
			language.XOG:     "ᏐᎦ",
			language.YAV:     "ᏰᎾᎦᏇᏂ",
			language.YBB:     "ᏰᎹᏋ",
			language.YI:      "ᏱᏗᏍ",
			language.YO:      "ᏲᏄᏆ",
			language.YUE:     "ᏓᎶᏂᎨ, ᎨᎾᏙᏂᏏ",
			language.ZGH:     "ᎠᏟᎶᏍᏗ ᎼᎶᎪ ᏔᎹᏏᏘ",
			language.ZH:      "ᏓᎶᏂᎨ, ᎹᏓᏈᏂ",
			language.ZH_HANS: "ᎠᎯᏗᎨ ᎹᏓᏈᏂ ᏓᎶᏂᎨ",
			language.ZH_HANT: "ᎤᏦᏍᏗ ᎹᏓᏈᏂ ᏓᎶᏂᎨ",
			language.ZU:      "ᏑᎷ",
			language.ZUN:     "ᏑᏂ",
			language.ZXX:     "Ꮭ ᎦᏬᏂᎯᏍᏗ ᎦᎸᏛᎢ ᏱᎩ",
			language.ZZA:     "ᏌᏌ",
		},
		Territories: cldr.Territories{
			territory.V_001: "ᎡᎶᎯ",
			territory.V_002: "ᎬᎿᎨᏍᏛ",
			territory.V_003: "ᏧᏴᏢ ᎠᎹᏰᏟ",
			territory.V_005: "ᏧᎦᏃᏮ ᎠᎺᎵᎦ",
			territory.V_009: "ᎣᏏᏰᏂᎠ",
			territory.V_011: "ᏭᏕᎵᎬ ᏗᏜ ᎬᎿᎨᏍᏛ",
			territory.V_013: "ᎠᏰᏟ ᎠᎹᏰᏟ",
			territory.V_014: "ᏗᎧᎸᎬ ᏗᏜ ᎬᎿᎨᏍᏛ",
			territory.V_015: "ᏧᏴᏢ ᏗᏜ ᎬᎿᎨᏍᏛ",
			territory.V_017: "ᎠᏰᏟ ᎬᎿᎨᏍᏛ",
			territory.V_018: "ᏧᎦᎾᏮ ᏗᏜ ᎬᎿᎨᏍᏛ",
			territory.V_019: "ᎠᎺᎵᎦᎢ",
			territory.V_021: "ᏧᏴᏢ ᏗᏜ ᎠᎹᏰᏟ",
			territory.V_029: "ᎨᏆᏙᏯ",
			territory.V_030: "ᏗᎧᎸᎬ ᏗᏜ ᏓᎶᏂᎨᏍᏛ",
			territory.V_034: "ᏧᎦᎾᏮ ᏗᏜ ᏓᎶᏂᎨᏍᏛ",
			territory.V_035: "ᏧᎦᎾᏮ ᏗᎧᎸᎬ ᏓᎶᏂᎨᏍᏛ",
			territory.V_039: "ᏧᎦᎾᏮ ᏗᏜ ᏳᎳᏛ",
			territory.V_053: "ᎠᏍᏔᎴᏏᎠ",
			territory.V_054: "ᎺᎳᏁᏏᎠ",
			territory.V_057: "ᎠᏰᏟ ᏧᎾᎵᎪᎯ ᎾᎿ ᎹᎢᏉᏂᏏᏯ ᎢᎬᎾᏕᎾ",
			territory.V_061: "ᏆᎵᏂᏏᎠ",
			territory.V_142: "ᏓᎶᎾᎨᏍᏛ",
			territory.V_143: "ᎠᏰᏟ ᏓᎶᏂᎨᏍᏛ",
			territory.V_145: "ᏭᏕᎵᎬ ᏗᏜ ᏓᎶᏂᎨᏍᏛ",
			territory.V_150: "ᏳᎳᏛ",
			territory.V_151: "ᏗᎧᎸᎬ ᏗᏜ ᏳᎳᏛ",
			territory.V_154: "ᏧᏴᏢ ᏗᏜ ᏳᎳᏛ",
			territory.V_155: "ᏭᏕᎵᎬ ᏗᏜ ᏳᎳᏛ",
			territory.V_202: "ᎭᏫᏂ-ᏌᎭᏩ ᎬᎿᎨᏍᏛ",
			territory.V_419: "ᎳᏘᏂ ᎠᎹᏰᏟ",
			territory.AC:    "ᎤᎵᏌᎳᏓᏅ ᎤᎦᏚᏛᎢ",
			territory.AD:    "ᎠᏂᏙᎳ",
			territory.AE:    "ᏌᏊ ᎢᏳᎾᎵᏍᏔᏅ ᎡᎳᏈ ᎢᎹᎵᏘᏏ",
			territory.AF:    "ᎠᏫᎨᏂᏍᏖᏂ",
			territory.AG:    "ᎤᏪᏘ & ᏆᏊᏓ",
			territory.AI:    "ᎠᏂᎩᎳ",
			territory.AL:    "ᎠᎵᏇᏂᏯ",
			territory.AM:    "ᎠᎵᎻᏂᎠ",
			territory.AO:    "ᎠᏂᎪᎳ",
			territory.AQ:    "ᏧᏁᏍᏓᎸ",
			territory.AR:    "ᎠᏥᏂᏘᏂᎠ",
			territory.AS:    "ᎠᎺᎵᎧ ᏌᎼᎠ",
			territory.AT:    "ᎠᏍᏟᏯ",
			territory.AU:    "ᎡᎳᏗᏜ",
			territory.AW:    "ᎠᎷᏆ",
			territory.AX:    "ᎣᎴᏅᏓ ᏚᎦᏚᏛᎢ",
			territory.AZ:    "ᎠᏎᏆᏣᏂ",
			territory.BA:    "ᏉᏏᏂᎠ & ᎲᏤᎪᏫᎾ",
			territory.BB:    "ᏆᏇᏙᏍ",
			territory.BD:    "ᏆᏂᎦᎵᏕᏍ",
			territory.BE:    "ᏇᎵᏥᎥᎻ",
			territory.BF:    "ᏋᎩᎾ ᏩᏐ",
			territory.BG:    "ᏊᎵᎨᎵᎠ",
			territory.BH:    "ᏆᎭᎴᎢᏂ",
			territory.BI:    "ᏋᎷᏂᏗ",
			territory.BJ:    "ᏆᏂᎢᏂ",
			territory.BL:    "ᎤᏓᏅᏘ ᏆᏕᎳᎻ",
			territory.BM:    "ᏆᏊᏓ",
			territory.BN:    "ᏊᎾᎢ",
			territory.BO:    "ᏉᎵᏫᎠ",
			territory.BQ:    "ᎧᎵᏈᎢᏂᎯ ᎾᏍᎩᏁᏛᎳᏂ",
			territory.BR:    "ᏆᏏᎵ",
			territory.BS:    "ᎾᏍᎩ ᏆᎭᎹᏍ",
			territory.BT:    "ᏊᏔᏂ",
			territory.BV:    "ᏊᏪ ᎤᎦᏚᏛᎢ",
			territory.BW:    "ᏆᏣᏩᎾ",
			territory.BY:    "ᏇᎳᎷᏍ",
			territory.BZ:    "ᏇᎵᏍ",
			territory.CA:    "ᎨᎾᏓ",
			territory.CC:    "ᎪᎪᏍ (ᎩᎵᏂ) ᏚᎦᏚᏛᎢ",
			territory.CD:    "ᎧᏂᎪ - ᎨᏂᏝᏌ",
			territory.CF:    "ᎬᎿᎨᏍᏛ ᎠᏰᏟ ᏍᎦᏚᎩ",
			territory.CG:    "ᎧᏂᎪ - ᏆᏌᏩᎵ",
			territory.CH:    "ᏍᏫᏍ",
			territory.CI:    "ᎢᏬᎵ ᎾᎿ ᎠᎹᏳᎶᏗ",
			territory.CK:    "ᎠᏓᏍᏓᏴᎲᏍᎩ ᏚᎦᏚᏛᎢ",
			territory.CL:    "ᏥᎵ",
			territory.CM:    "ᎧᎹᎷᏂ",
			territory.CN:    "ᏓᎶᏂᎨᏍᏛ",
			territory.CO:    "ᎪᎸᎻᏈᎢᎠ",
			territory.CP:    "ᎦᏂᏴᏔᏅᎣᏓᎸ ᎤᎦᏚᏛᎢ",
			territory.CR:    "ᎪᏍᏓ ᎵᎧ",
			territory.CU:    "ᎫᏆ",
			territory.CV:    "ᎢᎬᎾᏕᎾ ᎢᏤᏳᏍᏗ",
			territory.CW:    "ᎫᎳᎨᎣ",
			territory.CX:    "ᏓᏂᏍᏓᏲᎯᎲ ᎤᎦᏚᏛᎢ",
			territory.CY:    "ᏌᎢᏆᏍ",
			territory.CZ:    "ᏤᎩᎠ",
			territory.DE:    "ᎠᏂᏛᏥ",
			territory.DG:    "ᏗᏰᎪ ᎦᏏᏯ",
			territory.DJ:    "ᏥᏊᏗ",
			territory.DK:    "ᏗᏂᎹᎦ",
			territory.DM:    "ᏙᎻᏂᎧ",
			territory.DO:    "ᏙᎻᏂᎧᏂ ᏍᎦᏚᎩ",
			territory.DZ:    "ᎠᎵᏥᎵᏯ",
			territory.EA:    "ᏑᏔ ᎠᎴ ᎺᎵᏯ",
			territory.EC:    "ᎡᏆᏙᎵ",
			territory.EE:    "ᎡᏍᏙᏂᏯ",
			territory.EG:    "ᎢᏥᏈᎢ",
			territory.EH:    "ᏭᏕᎵᎬ ᏗᏜ ᏌᎮᎳ",
			territory.ER:    "ᎡᎵᏟᏯ",
			territory.ES:    "ᎠᏂᏍᏆᏂᏱ",
			territory.ET:    "ᎢᏗᎣᏈᎠ",
			territory.EU:    "ᏳᎳᏛ ᎠᏂᎤᎾᏓᏡᎬ",
			territory.EZ:    "ᏳᎶᎠᏍᏓᏅᏅ",
			territory.FI:    "ᏫᏂᎦᏙᎯ",
			territory.FJ:    "ᏫᏥ",
			territory.FK:    "ᏩᎩ ᏚᎦᏚᏛᎢ",
			territory.FM:    "ᎹᎢᏉᏂᏏᏯ",
			territory.FO:    "ᏪᎶ ᏚᎦᏚᏛᎢ",
			territory.FR:    "ᎦᎸᏥᏱ",
			territory.GA:    "ᎦᏉᏂ",
			territory.GB:    "ᎩᎵᏏᏲ",
			territory.GD:    "ᏋᎾᏓ",
			territory.GE:    "ᏣᎠᏥᎢ",
			territory.GF:    "ᎠᏂᎦᎸᏥ ᎩᎠ",
			territory.GG:    "ᎬᏂᏏ",
			territory.GH:    "ᎦᎠᎾ",
			territory.GI:    "ᏥᏆᎵᏓ",
			territory.GL:    "ᎢᏤᏍᏛᏱ",
			territory.GM:    "ᎦᎹᏈᎢᎠ",
			territory.GN:    "ᎩᎢᏂ",
			territory.GP:    "ᏩᏓᎷᏇ",
			territory.GQ:    "ᎡᏆᏙᎵᎠᎵ ᎩᎢᏂ",
			territory.GR:    "ᎪᎢᎯ",
			territory.GS:    "ᏧᎦᏃᏮ ᏣᎠᏥᎢ ᎠᎴ ᎾᏍᎩ ᏧᎦᏃᏮ ᎠᏍᏛᎭᏟ ᏚᎦᏚᏛᎢ",
			territory.GT:    "ᏩᏔᎹᎳ",
			territory.GU:    "ᏆᎻ",
			territory.GW:    "ᎩᎢᏂ-ᏈᏌᎤᏫ",
			territory.GY:    "ᎦᏯᎾ",
			territory.HK:    "ᎰᏂᎩ ᎪᏂᎩ ᎤᏓᏤᎵᏓ ᏧᏂᎸᏫᏍᏓᏁᏗ ᎢᎬᎾᏕᎾ ᏓᎶᏂᎨᏍᏛ",
			territory.HM:    "ᎲᏗ ᎤᎦᏚᏛᎢ ᎠᎴ ᎺᎩᏓᎾᎵᏗ ᏚᎦᏚᏛᎢ",
			territory.HN:    "ᎭᏂᏚᎳᏍ",
			territory.HR:    "ᎧᎶᎡᏏᎠ",
			territory.HT:    "ᎮᎢᏘ",
			territory.HU:    "ᎲᏂᎦᎵ",
			territory.IC:    "ᏥᏍᏆ ᏚᎦᏚᏛᎢ",
			territory.ID:    "ᎢᏂᏙᏂᏍᏯ",
			territory.IE:    "ᎠᏲᎳᏂ",
			territory.IL:    "ᎢᏏᎵᏱ",
			territory.IM:    "ᎤᏍᏗ ᎤᎦᏚᏛᎢ ᎾᎿ ᎠᏍᎦᏯ",
			territory.IN:    "ᎢᏅᏗᎾ",
			territory.IO:    "ᏈᏗᏏ ᏴᏫᏯ ᎠᎺᏉ ᎢᎬᎾᏕᏅ",
			territory.IQ:    "ᎢᎳᎩ",
			territory.IR:    "ᎢᎴᏂ",
			territory.IS:    "ᏧᏁᏍᏓᎸᎯ",
			territory.IT:    "ᎢᏔᎵ",
			territory.JE:    "ᏨᎵᏏ",
			territory.JM:    "ᏣᎺᎢᎧ",
			territory.JO:    "ᏦᏓᏂ",
			territory.JP:    "ᏣᏩᏂᏏ",
			territory.KE:    "ᎨᏂᏯ",
			territory.KG:    "ᎩᎵᏣᎢᏍ",
			territory.KH:    "ᎧᎹᏉᏗᎠᏂ",
			territory.KI:    "ᎧᎵᏆᏘ",
			territory.KM:    "ᎪᎼᎳᏍ",
			territory.KN:    "ᎤᏓᏅᏘ ᎨᏘᏏ ᎠᎴ ᏁᏪᏏ",
			territory.KP:    "ᏧᏴᏢ ᎪᎵᎠ",
			territory.KR:    "ᏧᎦᏃᏮ ᎪᎵᎠ",
			territory.KW:    "ᎫᏪᎢᏘ",
			territory.KY:    "ᎨᎢᎹᏂ ᏚᎦᏚᏛᎢ",
			territory.KZ:    "ᎧᏎᎧᏍᏕᏂ",
			territory.LA:    "ᎴᎣᏍ",
			territory.LB:    "ᎴᏆᎾᏂ",
			territory.LC:    "ᎤᏓᏅᏘ ᎷᏏᏯ",
			territory.LI:    "ᎵᎦᏗᏂᏍᏓᏂ",
			territory.LK:    "ᏍᎵ ᎳᏂᎧ",
			territory.LR:    "ᎳᏈᎵᏯ",
			territory.LS:    "ᎴᏐᏙ",
			territory.LT:    "ᎵᏗᏪᏂᎠ",
			territory.LU:    "ᎸᎧᏎᏋᎩ",
			territory.LV:    "ᎳᏘᏫᎠ",
			territory.LY:    "ᎵᏈᏯ",
			territory.MA:    "ᎼᎶᎪ",
			territory.MC:    "ᎹᎾᎪ",
			territory.MD:    "ᎹᎵᏙᏫᎠ",
			territory.ME:    "ᎼᏂᏔᏁᎦᎶ",
			territory.MF:    "ᎤᏓᏅᏘ ᏡᏡ",
			territory.MG:    "ᎹᏓᎦᏍᎧᎵ",
			territory.MH:    "ᎹᏌᎵ ᏚᎦᏚᏛᎢ",
			territory.MK:    "ᏧᏴᏜ ᎹᏎᏙᏂᏯ",
			territory.ML:    "ᎹᎵ",
			territory.MM:    "ᎹᏯᎹᎵ (ᏇᎵᎹ)",
			territory.MN:    "ᎹᏂᎪᎵᎠ",
			territory.MO:    "ᎹᎧᎣ (ᎤᏓᏤᎵᏓ ᏧᏂᎸᏫᏍᏓᏁᏗ ᎢᎬᎾᏕᎾ) ᏣᎢ",
			territory.MP:    "ᏧᏴᏢ ᏗᏜ ᎹᎵᎠᎾ ᏚᎦᏚᏛᎢ",
			territory.MQ:    "ᎹᏘᏂᎨ",
			territory.MR:    "ᎹᏘᎢᏯ",
			territory.MS:    "ᎹᏂᏘᏌᎳᏗ",
			territory.MT:    "ᎹᎵᏔ",
			territory.MU:    "ᎼᎵᏏᎥᏍ",
			territory.MV:    "ᎹᎵᏗᏫᏍ",
			territory.MW:    "ᎹᎳᏫ",
			territory.MX:    "ᎠᏂᏍᏆᏂ",
			territory.MY:    "ᎹᎴᏏᎢᎠ",
			territory.MZ:    "ᎼᏎᎻᏇᎩ",
			territory.NA:    "ᎾᎻᏈᎢᏯ",
			territory.NC:    "ᎢᏤ ᎧᎵᏙᏂᎠᏂ",
			territory.NE:    "ᎾᎢᏨ",
			territory.NF:    "ᏃᎵᏬᎵᎩ ᎤᎦᏚᏛᎢ",
			territory.NG:    "ᏂᏥᎵᏯ",
			territory.NI:    "ᏂᎧᎳᏆ",
			territory.NL:    "ᏁᏛᎳᏂ",
			territory.NO:    "ᏃᏪ",
			territory.NP:    "ᏁᏆᎵ",
			territory.NR:    "ᏃᎤᎷ",
			territory.NU:    "ᏂᏳ",
			territory.NZ:    "ᎢᏤ ᏏᎢᎴᏂᏗ",
			territory.OM:    "ᎣᎺᏂ",
			territory.PA:    "ᏆᎾᎹ",
			territory.PE:    "ᏇᎷ",
			territory.PF:    "ᎠᏂᎦᎸᏥ ᏆᎵᏂᏏᎠ",
			territory.PG:    "ᏆᏇ ᎢᏤ ᎩᎢᏂ",
			territory.PH:    "ᎠᏂᏈᎵᎩᏃ",
			territory.PK:    "ᏆᎩᏍᏖᏂ",
			territory.PL:    "ᏉᎳᏂ",
			territory.PM:    "ᎤᏓᏅᏘ ᏈᏰ ᎠᎴ ᎻᏇᎶᏂ",
			territory.PN:    "ᏈᎧᎵᏂ ᏚᎦᏚᏛᎢ",
			territory.PR:    "ᏇᎡᏙ ᎵᎢᎪ",
			territory.PS:    "ᏆᎴᏍᏗᏂᎠᏂ ᏄᎬᏫᏳᏌᏕᎩ",
			territory.PT:    "ᏉᏥᎦᎳ",
			territory.PW:    "ᏆᎴᎠᏫ",
			territory.PY:    "ᏆᎳᏇᎢᏯ",
			territory.QA:    "ᎧᏔᎵ",
			territory.QO:    "ᎠᏍᏛ ᎣᏏᏰᏂᎠ",
			territory.RE:    "ᎴᏳᏂᎠᏂ",
			territory.RO:    "ᎶᎹᏂᏯ",
			territory.RS:    "ᏒᏈᏯ",
			territory.RU:    "ᏲᏂᎢ",
			territory.RW:    "ᎶᏩᏂᏓ",
			territory.SA:    "ᏌᎤᏗ ᎡᎴᏈᎠ",
			territory.SB:    "ᏐᎶᎹᏂ ᏚᎦᏚᏛᎢ",
			territory.SC:    "ᏏᎡᏥᎵᏍ",
			territory.SD:    "ᏑᏕᏂ",
			territory.SE:    "ᏍᏫᏕᏂ",
			territory.SG:    "ᏏᏂᎦᏉᎵ",
			territory.SH:    "ᎤᏓᏅᏘ ᎮᎵᎾ",
			territory.SI:    "ᏍᎶᏫᏂᎠ",
			territory.SJ:    "ᏍᏩᎵᏆᎵᏗ ᎠᎴ ᏤᏂ ᎹᏰᏂ",
			territory.SK:    "ᏍᎶᏩᎩᎠ",
			territory.SL:    "ᏏᎡᎳ ᎴᎣᏂ",
			territory.SM:    "ᎤᏓᏅᏘ ᎹᎵᎢᏃ",
			territory.SN:    "ᏏᏂᎦᎵ",
			territory.SO:    "ᏐᎹᎵ",
			territory.SR:    "ᏒᎵᎾᎻ",
			territory.SS:    "ᏧᎦᎾᏮ ᏑᏕᏂ",
			territory.ST:    "ᏌᎣ ᏙᎺ ᎠᎴ ᏈᏂᏏᏇ",
			territory.SV:    "ᎡᎵᏌᎵᏆᏙᎵ",
			territory.SX:    "ᏏᏂᏘ ᎹᏘᏂ",
			territory.SY:    "ᏏᎵᎠ",
			territory.SZ:    "ᎡᏍᏩᏘᏂ",
			territory.TA:    "ᏟᏍᏛᏂ Ꮣ ᎫᎾᎭ",
			territory.TC:    "ᎠᏂᏛᎵᎩ ᎠᎴ ᎨᎢᎪ ᏚᎦᏚᏛᎢ",
			territory.TD:    "ᏣᏗ",
			territory.TF:    "ᎠᏂᎦᎸᏥ ᏧᎦᎾᏮ ᎦᏙᎯ ᎤᎵᏍᏛᎢ",
			territory.TG:    "ᏙᎪ",
			territory.TH:    "ᏔᏯᎴᏂ",
			territory.TJ:    "ᏔᏥᎩᏍᏕᏂ",
			territory.TK:    "ᏙᎨᎳᏭ",
			territory.TL:    "ᏘᎼᎵ-ᎴᏍᏖ",
			territory.TM:    "ᏛᎵᎩᎺᏂᏍᏔᏂ",
			territory.TN:    "ᏚᏂᏏᏍᎠ",
			territory.TO:    "ᏙᎾᎦ",
			territory.TR:    "ᎬᏃ",
			territory.TT:    "ᏟᏂᏕᏗ ᎠᎴ ᏙᏆᎪ",
			territory.TV:    "ᏚᏩᎷ",
			territory.TW:    "ᏔᎢᏩᏂ",
			territory.TZ:    "ᏖᏂᏏᏂᏯ",
			territory.UA:    "ᏳᎧᎴᏂ",
			territory.UG:    "ᏳᎦᏂᏓ",
			territory.UM:    "U.S. ᎠᏍᏛ ᏚᎦᏚᏛᎢ",
			territory.UN:    "ᏌᏊ ᎢᏳᎾᎵᏍᏔᏅ ᎠᏰᎵ ᏚᎾᏙᏢᏒ",
			territory.US:    "ᏌᏊ ᎢᏳᎾᎵᏍᏔᏅ ᏍᎦᏚᎩ",
			territory.UY:    "ᏳᎷᏇ",
			territory.UZ:    "ᎤᏍᏇᎩᏍᏖᏂ",
			territory.VA:    "ᎠᏥᎳᏁᏠ ᎦᏚᎲ",
			territory.VC:    "ᎤᏓᏅᏘ ᏫᏂᏏᏂᏗ ᎠᎴ ᎾᏍᎩ ᏇᎾᏗᏁᏍ",
			territory.VE:    "ᏪᏁᏑᏪᎳ",
			territory.VG:    "ᏈᏗᏍ ᎠᏒᏂᎸ ᏂᎨᏒᎾ ᏚᎦᏚᏛᎢ",
			territory.VI:    "U.S. ᎠᏒᏂᎸ ᏂᎨᏒᎾ ᏚᎦᏚᏛᎢ",
			territory.VN:    "ᏫᎡᏘᎾᎻ",
			territory.VU:    "ᏩᏂᎤᏩᏚ",
			territory.WF:    "ᏩᎵᏍ ᎠᎴ ᏊᏚᎾ",
			territory.WS:    "ᏌᎼᎠ",
			territory.XA:    "ᏡᏙ-ᏄᏍᏛᎢᎥᎧᏁᎬᎢ",
			territory.XB:    "ᏡᏙ-ᏈᏗ",
			territory.XK:    "ᎪᏐᏉ",
			territory.YE:    "ᏰᎺᏂ",
			territory.YT:    "ᎺᏯᏖ",
			territory.ZA:    "ᏧᎦᎾᏮ ᎬᎿᎨᏍᏛ",
			territory.ZM:    "ᏌᎻᏈᏯ",
			territory.ZW:    "ᏏᎻᏆᏇ",
			territory.ZZ:    "ᏄᏬᎵᏍᏛᎾ ᎤᏔᏂᏗᎦᏙᎯ",
		},
	}
}
