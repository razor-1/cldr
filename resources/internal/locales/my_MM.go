// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_my_MM() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "my_MM",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y၊ MMMM d၊ EEEE", Long: "y၊ d MMMM", Medium: "y၊ MMM d", Short: "dd-MM-yy"}, Time: cldr.CalendarDateFormat{Full: "zzzz HH:mm:ss", Long: "z HH:mm:ss", Medium: "B HH:mm:ss", Short: "B H:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ဇန်", Feb: "ဖေ", Mar: "မတ်", Apr: "ဧ", May: "မေ", Jun: "ဇွန်", Jul: "ဇူ", Aug: "ဩ", Sep: "စက်", Oct: "အောက်", Nov: "နို", Dec: "ဒီ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ဇ", Feb: "ဖ", Mar: "မ", Apr: "ဧ", May: "မ", Jun: "ဇ", Jul: "ဇ", Aug: "ဩ", Sep: "စ", Oct: "အ", Nov: "န", Dec: "ဒ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ဇန်နဝါရီ", Feb: "ဖေဖော်ဝါရီ", Mar: "မတ်", Apr: "ဧပြီ", May: "မေ", Jun: "ဇွန်", Jul: "ဇူလိုင်", Aug: "ဩဂုတ်", Sep: "စက်တင်ဘာ", Oct: "အောက်တိုဘာ", Nov: "နိုဝင်ဘာ", Dec: "ဒီဇင်ဘာ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "တနင်္ဂနွေ", Mon: "တနင်္လာ", Tue: "အင်္ဂါ", Wed: "ဗုဒ္ဓဟူး", Thu: "ကြာသပတေး", Fri: "သောကြာ", Sat: "စနေ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "တ", Mon: "တ", Tue: "အ", Wed: "ဗ", Thu: "က", Fri: "သ", Sat: "စ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "တနင်္ဂနွေ", Mon: "တနင်္လာ", Tue: "အင်္ဂါ", Wed: "ဗုဒ္ဓဟူး", Thu: "ကြာသပတေး", Fri: "သောကြာ", Sat: "စနေ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "တနင်္ဂနွေ", Mon: "တနင်္လာ", Tue: "အင်္ဂါ", Wed: "ဗုဒ္ဓဟူး", Thu: "ကြာသပတေး", Fri: "သောကြာ", Sat: "စနေ"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "နံနက်", PM: "ညနေ"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "နံနက်", PM: "ညနေ"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "နံနက်", PM: "ညနေ"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "¤\u00a0#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "အာရပ်စော်ဘွားများ ပြည်ထောင်စု ဒါဟမ်း", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "အာဖဂန် အာဖဂါနီ", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "အယ်ဘေးနီးယား လီခ်", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "အာမေးနီးယား ဒရမ်", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "နယ်သာလန် အန်တီလန် ဂင်းဒါး", Symbol: "NAf"},
				currency.AOA: cldr.Currency{DisplayName: "အန်ဂိုလာ ကွမ်ဇာ", Symbol: "AOA"},
				currency.ARP: cldr.Currency{DisplayName: "အာဂျင်တီးနား ပီဆို (၁၉၈၃–၁၉၈၅)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "အာဂျင်တီးနား ပီဆို", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ဩစတြေးလျ ဒေါ်လာ", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "အရူးဗာ ဖလိုရင်း", Symbol: "Afl"},
				currency.AZN: cldr.Currency{DisplayName: "အဇာဘိုင်ဂျန် မာနတ်", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "ဘော့စနီးယားနှင့် ဟာဇီဂိုဘီးနား ငွေလဲနိုင်သော မတ်က်", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "ဘာဘေးဒီယန်း ဒေါ်လာ", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "ဘင်္ဂလားဒေ့ရှ် တာကာ", Symbol: "BDT"},
				currency.BEF: cldr.Currency{DisplayName: "ဘယ်လ်ဂျီယမ် ဖရန့်", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "ဘူလ်ဂေးရီးယား လက်ဖ်", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "ဘာရိန်း ဒီနား", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "ဘူရွန်ဒီ ဖရန့်", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "ဘာမြူဒါ ဒေါ်လာ", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ဘရူနိုင်း ဒေါ်လာ", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "ဘိုလီးဗီးယား ဘိုလီးဗီယားနို", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "ဘိုလီးဘီးယား ပီဆို", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "ဘရာဇီး ရီးယဲ", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "ဘဟားမား ဒေါ်လာ", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ဘူတန် အံဂါလ်ထရန်", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "ဗမာ ကျပ်", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "ဘော့ဆွာနာ ပုလ", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "ဘီလာရုစ် ရူဘယ်အသစ် (၁၉၉၄–၁၉၉၉)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "ဘီလာရုစ် ရူဘယ်", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "ဘီလာရုဇ် ရူဘယ် (၂၀၀၀–၂၀၁၆)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "ဘလိဇ် ဒေါ်လာ", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "ကနေဒါ ဒေါ်လာ", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "ကွန်ဂို ဖရန့်", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "ဆွစ် ဖရန့်", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "ချီလီ ပီဆို", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "တရုတ် ယွမ် (ဟောင်ကောင်)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "တရုတ် ယွမ်", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "ကိုလံဘီယာ ပီဆို", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "ကို့စတာရီကာ ကိုလွန်", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "နိုင်ငံခြားငွေလဲလှယ်နိုင်သော ကျူးဘားပီဆို", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "ကျူးဘား ပီဆို", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "ကိတ်ပ်ဗာဒီ အက်စ်ခူဒို", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "ဆိုက်ပရက်စ် ပေါင်", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "ချက် ခိုရိုနာ", Symbol: "CZK"},
				currency.DEM: cldr.Currency{DisplayName: "ဂျာမဏီ မတ်", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "ဂျီဘူတီ ဖရန့်", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ဒိန်းမတ် ခရိုဏာ", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ဒိုမီနီကန် ပီဆို", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "အယ်လ်ဂျီးရီးယား ဒီနာ", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "အီဂျစ် ပေါင်", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "အီရီထရီးယား နာ့ခ်ဖာ", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "စပိန် ပယ်စေးတာ", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "အီသီယိုးပီးယား ဘီးယာ", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "ယူရို", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "ဖီဂျီ ဒေါ်လာ", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ဖော့ကလန်ကျွန်းစု ပေါင်", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "ပြင်သစ် ဖရန့်", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "ဗြိတိသျှ ပေါင်", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "ဂျော်ဂျီယာ လားရီ", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "ဂါနာ ဆဲဒီ", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "ဂျီဘရော်လ်တာ ပေါင်", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "ဂမ်ဘီယာ ဒါလာစီ", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "ဂီနီရာ ဖရန့်", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ဂွါတီမာလာ ခက်ဇော်လ်", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "ဂိုင်ယာနာ ဒေါ်လာ", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "ဟောင်ကောင် ဒေါ်လာ", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "ဟွန်ဒူးရပ်စ် လမ်းပီရာ", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "ခရိုအေးရှား ခူးနာ", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "ဟေတီဂူးအော်ဒ်", Symbol: "G"},
				currency.HUF: cldr.Currency{DisplayName: "ဟန်ဂေရီယံ ဖော်ရင့်တ်", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "အင်ဒိုနီးရှား ရူပီးယား", Symbol: "IDR"},
				currency.ILP: cldr.Currency{DisplayName: "အစ္စရေး ပေါင်", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "အစ္စရေး ရှဲကလ်အသစ်", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "အိန္ဒိယ ရူပီး", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "အီရတ် ဒီနာ", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "အီရန် ရီအော်လ်", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "အိုက်စလန် ခရိုဏာ", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "ဂျမေကာ ဒေါ်လာ", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "ဂျော်ဒန် ဒီနာ", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "ဂျပန် ယန်း", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "ကင်ညာ သျှီလင်", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "ကာဂျစ္စတန် ဆော်မ်", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "ကမ္ဘောဒီးယား ရီးယဲ", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "ကိုမိုရိုစ် ဖရန့်", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "မြောက်ကိုရီးယား ဝမ်", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "တောင်ကိုရီးယား ဝမ်", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "ကူဝိတ် ဒီနာ", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "ကေမန် ကျွန်းစု ဒေါ်လာ", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "ကာဇက်စတန် ထိန်ဂျီ", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "လာအို ကစ်", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "လက်ဘနွန် ပေါင်", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "သီရိလင်္ကာ ရူပီး", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "လိုက်ဘေးရီးယား ဒေါ်လာ", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "လစ်သူယေးနီးယားလီတားစ်", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "လတ်ဗီးယားလတ်", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "လစ်ဗျား ဒိုင်နာ", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "မိုရိုကို ဒရမ်", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "မောလ်ဒိုဗာ လယ်အို", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "မာလာဂါစီ အရီရရီ", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "မက်ဆီဒိုးနီးယား ဒီနာ", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "မြန်မာ ကျပ်", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "မွန်ဂိုးလီးယား ထူးဂရခ်", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "မကာအို ပါတားကား", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "မော်ရီတေးနီးယား အူဂီးယာ (၁၉၇၃–၂၀၁၇)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "မော်ရီတေးနီးယန်း အူဂီးယာ", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "မောရစ်ရှ ရူပီး", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "မော်လ်ဒိုက် ရူးဖီရာ", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "မာလာဝီ ခွါးချာ", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "မက်ကဆီကို ပီဆို", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "မလေးရှား ရင်းဂစ်", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "မိုဇမ်ဘစ် မက်တီခယ်လ်", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "နမီးဘီးယား ဒေါ်လာ", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "နိုင်ဂျီးရီးယား နိုင်းရာ", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "နီကာရာဂွါ ခိုးဒိုဘာ", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "နော်ဝေ ခရိုဏာ", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "နီပေါ ရူပီး", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "နယူးဇီလန် ဒေါ်လာ", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "အိုမန်နီ ရီရယ်", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "ပနားမား ဘလျဘိုးအာ", Symbol: "B/."},
				currency.PEN: cldr.Currency{DisplayName: "ပီရူး ဆိုးလ်", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "ပါပူအာ နယူးဂီနီ ခီးနာ", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ဖိလစ်ပိုင် ပီဆို", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "ပါကစ္စတန် ရူပီး", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "ပိုလန် ဇလော့တီ", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "ပါရာဂွေး ဂွါးအ်နီး", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "ကာတာရီ ရီရယ်", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "ရိုမေးနီးယား လယ်အို", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "ဆားဘီးယား ဒယ်နား", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "ရုရှ ရူဘယ်", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "ရုရှ ရူဘယ် (၁၉၉၁–၁၉၉၈)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "ရဝန်ဒါ ဖရန့်", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "ဆော်ဒီအာရေးဗီးယား ရီယော်လ်", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "ဆော်လမွန်ကျွန်းစု ဒေါ်လာ", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "ဆေးရှဲ ရူပီး", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "ဆူဒန် ပေါင်", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "ဆူဒန် ပေါင်အဟောင်း", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "ဆွီဒင် ခရိုဏာ", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "စင်္ကာပူ ဒေါ်လာ", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "စိန့်ဟယ်လယ်နာ ပေါင်", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "ဆီယာရာလီယွန်း လီအိုနီ", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "ဆိုမာလီ သျှီလင်", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "ဆူရီနမ်း ဒေါ်လာ", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "တောင်ဆူဒန် ပေါင်", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "ဆောင်တူမေးနှင့် ပရင်စီပီ ဒိုဘရာ", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "ဆောင်တူမေးနှင့် ပရင်စီပီ ဒိုဘရာ (၂၀၁၈)", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "ဆိုဗီယက် ရူဗယ်", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "ဆီးရီးယား ပေါင်", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "ဆွာဇီလန် လီလန်းဂီနီ", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "ထိုင်း ဘတ်", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "တာဂျစ်ကစ္စတန် ဆိုမိုနီ", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "တာ့ခ်မင်နစ္စတန် မာနတ်", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "တူနီးရှား ဒိုင်နာ", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "တွန်ဂါ ဗန်ဂါ", Symbol: "TOP"},
				currency.TRL: cldr.Currency{DisplayName: "ရှေးဟောင်းတူရကီ လိုင်ရာ", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "တူရကီ လိုင်ရာ", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ထရီနီဒတ်နှင့် တိုဘက်ဂို ဒေါ်လာ", Symbol: "TT$"},
				currency.TWD: cldr.Currency{DisplayName: "ထိုင်ဝမ် ဒေါ်လာအသစ်", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "တန်ဇန်းနီးယား သျှီလင်", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "ယူကရိန်း ဟီရီဗင်းညား", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "ယူဂန္ဒာ သျှီလင်", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "အမေရိကန် ဒေါ်လာ", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "အမေရိကန် ဒေါ်လာ (နောက်နေ့)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "အမေရိကန် ဒေါ်လာ (တနေ့တည်း)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "ဥရုဂွေး ပီဆို", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "ဥဇဘက်ကစ္စတန် ဆော်မ်", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "ဗင်နီဇွဲလား ဘိုလီဗာ (၂၀၀၈–၂၀၁၈)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "ဗင်နီဇွဲလန်း ဘိုလီဗာ", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "ဗီယက်နမ် ဒေါင်", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "ဗနွားတူ ဗားထူ", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "ဆမိုအား ထားလာ", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "ကင်မရွန်း ဖရန့်", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "ငွေ", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "ရွှေ", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "ဥရောပငွေကြေးစံနစ်", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "အရှေ့ကာရစ်ဘီယံ ဒေါ်လာ", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "အထူးထုတ်ယူခွင့်", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "အနောက် အာဖရိက CFA ဖရန့်", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP ဖရန့်", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "ပလက်တီနမ်", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "စမ်းသပ် ငွေကြေး ကုဒ်", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "မသိသို့မဟုတ်မရှိသောငွေကြေး", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "ယီမင်နီ ရီရယ်", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "တောင်အာဖရိက ရန်း", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "ဇင်ဘာဘွေ ခွါးချာ", Symbol: "ZMW"},
				currency.ZWD: cldr.Currency{DisplayName: "ဇင်ဘာဘွေ ဒေါ်လာ", Symbol: ""},
			},
		},
	}
}
