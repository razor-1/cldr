// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_km_KH() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "km_KH",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} នៅ\u200bម៉ោង {0}", Long: "{1} នៅ\u200bម៉ោង {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "ម៉ោង\u200bសកល {0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "មករា", Feb: "កុម្ភៈ", Mar: "មីនា", Apr: "មេសា", May: "ឧសភា", Jun: "មិថុនា", Jul: "កក្កដា", Aug: "សីហា", Sep: "កញ្ញា", Oct: "តុលា", Nov: "វិច្ឆិកា", Dec: "ធ្នូ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ម", Feb: "ក", Mar: "ម", Apr: "ម", May: "ឧ", Jun: "ម", Jul: "ក", Aug: "ស", Sep: "ក", Oct: "ត", Nov: "វ", Dec: "ធ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "មករា", Feb: "កុម្ភៈ", Mar: "មីនា", Apr: "មេសា", May: "ឧសភា", Jun: "មិថុនា", Jul: "កក្កដា", Aug: "សីហា", Sep: "កញ្ញា", Oct: "តុលា", Nov: "វិច្ឆិកា", Dec: "ធ្នូ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "អាទិត្យ", Mon: "ចន្ទ", Tue: "អង្គារ", Wed: "ពុធ", Thu: "ព្រហ", Fri: "សុក្រ", Sat: "សៅរ៍"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "អ", Mon: "ច", Tue: "អ", Wed: "ព", Thu: "ព", Fri: "ស", Sat: "ស"}, Short: cldr.CalendarDayFormatNameValue{Sun: "អា", Mon: "ច", Tue: "អ", Wed: "ពុ", Thu: "ព្រ", Fri: "សុ", Sat: "ស"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "អាទិត្យ", Mon: "ចន្ទ", Tue: "អង្គារ", Wed: "ពុធ", Thu: "ព្រហស្បតិ៍", Fri: "សុក្រ", Sat: "សៅរ៍"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00¤", CurrencyAccounting: "#,##0.00¤;(#,##0.00¤)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "ឌៀរហាំ\u200bអារ៉ាប់រួម", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "អាហ្វហ្គានី\u200bអាហ្វហ្គានីស្ថាន", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "លិក\u200bអាល់បានី", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "ដ្រាំ\u200bអាមេនី", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "ហ្គីឌិន\u200bហុល្លង់\u200bអង់ទីលៀន", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "ក្វាន់ហ្សា\u200bអង់ហ្គោឡា", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "ប៉េសួអាហ្សង់ទីន", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ដុល្លារ\u200bអូស្ត្រាលី", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "ហ្វ្រ័ររិញ\u200bអារ៉ូបា", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "ម៉ាណាត\u200bអាស៊ែបៃហ្សង់", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "ម៉ាក\u200bអាច\u200bបម្លែង\u200bបាន\u200bបូស្នី", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "ដុល្លារ\u200bបាបាដុស", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "តាកា\u200bបង់ក្លាដែស", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "លីវ\u200bប៊ុលហ្គារី", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "ឌីណា\u200bបារ៉ែន", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "ហ្វ្រង់\u200bប៊ូរុនឌី", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "ដុល្លារ\u200bប៊ឺមុយដា", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ដុល្លារ\u200bប្រុយណេ", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "បូលីវីណូ\u200bបូលីវី", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "រៀល\u200bប្រេស៊ីល", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "ដុល្លារ\u200bបាហាម៉ា", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ញូលត្រឹម\u200bប៊ូតាន", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "ពូឡា\u200bបុតស្វាណា", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "រ៉ូបល\u200bបេឡារុស", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "រ៉ូបល\u200bបេឡារុស (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "ដុល្លារ\u200bបេលី", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "ដុល្លារ\u200bកាណាដា", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "ហ្វ្រង់\u200bកុងហ្គោ", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "ហ្វ្រង់ស្វីស", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "ប៉េសូឈីលី", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "យ៉ន់ចិន (ក្រៅប្រទេស)", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "យ៉ន់\u200bចិន", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "ប៉េសូកូឡុំប៊ី", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "កូឡុង\u200bកូស្តារីកា", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "ប៉េសូ\u200bគុយបាអាច\u200bបម្លែង\u200bបាន", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "ប៉េសូគុយបា", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "អ៊ីស្កូឌូ\u200bកាប់វែរ", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "កូរុណា\u200bសាធារណៈ\u200bឆេក", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "ហ្វ្រង់\u200bជីប៊ូទី", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ក្រូណេ\u200bដាណាម៉ាក់", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ប៉េសូដូមីនីក", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "ឌីណា\u200bអាល់ស៊េរី", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "ផោនអេហ្ស៊ីប", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "ណាក់ហ្វាអេរីទ្រា", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ប៊័រ\u200bអេត្យូពី", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "អឺរ៉ូ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "ដុល្លារ\u200bហ្វីជី", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ផោន\u200bកោះ\u200bហ្វក់ឡែន", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "ផោនចក្រភពអង់គ្លេស", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "ឡារី\u200b\u200bហ្សកហ្ស៊ី", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "ស៊ីឌី\u200bហ្គាណា", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "ផោន\u200bហ្ស៊ីប្រាល់តា", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "ដាឡាស៊ី\u200bហ្គាំប៊ី", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "ហ្វ្រង់\u200bហ្គីណេ", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ហ្គីស្សាល\u200bក្វាតេម៉ាឡា", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "ដុល្លារ\u200bហ្គីយ៉ាន", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "ដុល្លារ\u200bហុងកុង", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "លិមពីរ៉ា\u200bហុងឌូរ៉ាស", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "គូណា\u200bក្រូអាត", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "គោឌី\u200bហៃទី", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "ហ្វូរីន\u200bហុងគ្រី", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "រូពីឥណ្ឌូណេស៊ី", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "ស៊ីគែលថ្មីអ៊ីស្រាអែល", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "រូពីឥណ្ឌា", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ឌីណា\u200bអ៊ីរ៉ាក់", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "រៀល\u200bអ៊ីរ៉ង់", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "ក្រូណា\u200bអ៊ីស្លង់", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "ដុល្លារ\u200bហ្សាម៉ាអ៊ីក", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "ឌីណា\u200bហ្ស៊កដានី", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "យេន\u200bជប៉ុន", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ស៊ីលិញ\u200bកេនយ៉ា", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "សុម\u200bកៀហ្ស៊ីស៊ីស្ថាន", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "រៀល\u200bកម្ពុជា", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "ហ្វ្រង់\u200bកូម័រ", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "វ៉ុនកូរ៉េខាងជើង", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "វ៉ុនកូរ៉េខាងត្បូង", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "ឌីណា\u200bគុយវ៉ែត", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "ដុល្លារ\u200bកោះ\u200bកៃម៉ែន", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "តង់ហ្គី\u200bកាហ្សាក់ស្ថាន", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "គីប\u200bឡាវ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "ផោន\u200bលីបង់", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "រូពីស្រីលង្កា", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "ដុល្លារ\u200bលីប៊ី", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "លីតា\u200bលីទុយអានី", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "ឡាត់\u200bឡេតូនី", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "ឌីណា\u200bលីប៊ី", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "ឌៀរហាំ\u200bម៉ារ៉ុក", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "លូ\u200bម៉ុលដាវី", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "អារៀរី\u200bម៉ាឡាហ្គាស៊ី", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "ឌីណាម៉ាសេដូនី", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "គីយ៉ាត\u200bភូមា", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "ទូរីក\u200bម៉ុងហ្គោលី", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "ប៉ាតាកា\u200bម៉ាកាវ", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "អ៊ូហ្គីយ៉ា\u200bម៉ូរីតានី (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "អ៊ូហ្គីយ៉ា\u200bម៉ូរីតានី", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "រូពីម៉ូរីតានី", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "រ៉ូហ្វីយ៉ា\u200bម៉ាល់ឌីវ", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "ក្វាចា\u200bម៉ាឡាវី", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "ប៉េសូម៉ិកសិក", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "រីងហ្គីត\u200bម៉ាឡេស៊ី", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "មីទីខល\u200bម៉ូសំប៊ិក", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "ដុល្លារ\u200bណាមីប៊ី", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "ណៃរ៉ា\u200bនីហ្សេរីយ៉ា", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "ខឌូបា\u200bនីការ៉ាហ្គា", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "ក្រូណេ\u200bន័រវ៉េ", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "រូពីនេប៉ាល់", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "ដុល្លារ\u200bនូវែលសេឡង់", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "រៀល\u200bរូម៉ានី", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "បាល់ប៉ៅ\u200bប៉ាណាម៉ា", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "ញូវ៉ូសូល\u200bប៉េរូ", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "គីណាប៉ាពួញូហ្គីណេ", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ប៉េសូហ្វីលីពីន", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "រូពីប៉ាគីស្ថាន", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "ហ្សូទី\u200bប៉ូឡូញ", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "ហ្គូរីនី\u200bប៉ារ៉ាហ្គាយ", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "រៀល\u200bកាតា", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "លូ\u200bរូម៉ានី", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "ឌីណាស៊ែប", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "រ៉ូបល\u200bរុស្ស៊ី", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ហ្វ្រង់\u200bរវ៉ាន់ដា", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "រីយ៉ាល\u200bអារ៉ាប៊ីសាអូឌីត", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "ដុល្លារ\u200bកោះ\u200bសូឡូម៉ុង", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "រូពី\u200bសីស្ហែល", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "ផោន\u200bស៊ូដង់", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "ក្រូណា\u200bស៊ុយអែត", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "ដុល្លារ\u200b\u200bសិង្ហបូរី", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "ផោន\u200bសាំងហេឡេណា", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "លីអ៊ុន\u200bសៀរ៉ាឡេអូន", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "ស៊ីលិញ\u200bសូម៉ាលី", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "ដុល្លារ\u200bសូរីណាម", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "ផោន\u200bស៊ូដង់\u200bខាង\u200bត្បូង", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "ឌូប្រា\u200bសៅតូម៉េ និងប្រាំងស៊ីប (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "ឌូប្រា\u200bសៅតូម៉េ និងប្រាំងស៊ីប", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "ផោន\u200bស៊ីរី", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "លីឡាងហ្គីនី\u200bស្វាស៊ីឡង់", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "បាត\u200bថៃ", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "សូមុនី\u200bតាហ្ស៊ីគីស្ថាន", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "ម៉ាណាត\u200bតួកម៉េនីស្ថាន", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "ឌីណាទុយនេស៊ី", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "ប៉ាអង់កា\u200bតុងហ្គា", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "លីរ៉ាទួរគី", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ដុល្លារ\u200bទ្រីនីដាដ និងតូបាហ្គោ", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "ដុល្លារ\u200bតៃវ៉ាន់", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "ស៊ីលិញ\u200bតង់សានី", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ហ៊ូនីយ៉ា\u200bអ៊ុយក្រែន", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "ស៊ីលិញ\u200bអ៊ូហ្គង់ដា", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "ដុល្លារ\u200bអាមេរិក", Symbol: "$"},
				currency.UYU: cldr.Currency{DisplayName: "ប៉េសូអ៊ុយរូហ្គាយ", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "សុមអ៊ូបេគីស្ថាន", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "ប៊ូលីវ៉ា\u200bវ៉េណេស៊ុយអេឡា (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "ប៊ូលីវ៉ា\u200bវ៉េណេស៊ុយអេឡា", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "ដុង\u200bវៀតណាម", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "វ៉ាទូវ៉ានូអាទូ", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "តាឡា\u200bសាម័រ", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "ហ្វ្រង់ CFA អាហ្វ្រិកកណ្តាល", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "ដុល្លារ\u200bការ៉ាប៊ីន\u200bខាង\u200bកើត", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "ហ្វ្រង់ CFA អាហ្វ្រិកខាងលិច", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "ហ្វ្រង់ CFP", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "រូបិយប័ណ្ណ\u200bមិនស្គាល់", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "រៀល\u200bយេម៉ែន", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "រ៉ង់អាហ្វ្រិកខាងត្បូង", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "ក្វាចា សំប៊ី (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "ក្វាចាហ្សំប៊ី", Symbol: "ZMW"},
			},
		},
	}
}
