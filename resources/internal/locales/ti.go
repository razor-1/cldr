// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ti() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ti",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE፣ dd MMMM መዓልቲ y G", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ጥሪ", Feb: "ለካ", Mar: "መጋ", Apr: "ሚያ", May: "ግን", Jun: "ሰነ", Jul: "ሓም", Aug: "ነሓ", Sep: "መስ", Oct: "ጥቅ", Nov: "ሕዳ", Dec: "ታሕ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ጥ", Feb: "ለ", Mar: "መ", Apr: "ሚ", May: "ግ", Jun: "ሰ", Jul: "ሓ", Aug: "ነ", Sep: "መ", Oct: "ጥ", Nov: "ሕ", Dec: "ታ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ጥሪ", Feb: "ለካቲት", Mar: "መጋቢት", Apr: "ሚያዝያ", May: "ግንቦት", Jun: "ሰነ", Jul: "ሓምለ", Aug: "ነሓሰ", Sep: "መስከረም", Oct: "ጥቅምቲ", Nov: "ሕዳር", Dec: "ታሕሳስ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "ሰን", Mon: "ሰኑ", Tue: "ሰሉ", Wed: "ረቡ", Thu: "ሓሙ", Fri: "ዓር", Sat: "ቀዳ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ሰ", Mon: "ሰ", Tue: "ሠ", Wed: "ረ", Thu: "ሓ", Fri: "ዓ", Sat: "ቀ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ሰን", Mon: "ሰኑ", Tue: "ሰሉ", Wed: "ረቡ", Thu: "ሓሙ", Fri: "ዓር", Sat: "ቀዳ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ሰንበት", Mon: "ሰኑይ", Tue: "ሠሉስ", Wed: "ረቡዕ", Thu: "ኃሙስ", Fri: "ዓርቢ", Sat: "ቀዳም"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ንጉሆ ሰዓተ", PM: "ድሕር ሰዓት"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ንጉሆ ሰዓተ", PM: "ድሕር ሰዓት"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ንጉሆ ሰዓተ", PM: "ድሕር ሰዓት"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "የብራዚል ሪል", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "", Symbol: "CA$"},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "የቻይና ዩአን ረንሚንቢ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.EGP: cldr.Currency{DisplayName: "", Symbol: "E£"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "የኢትዮጵያ ብር", Symbol: "Br"},
				currency.EUR: cldr.Currency{DisplayName: "አውሮ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "የእንግሊዝ ፓውንድ ስተርሊንግ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
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
				currency.INR: cldr.Currency{DisplayName: "የሕንድ ሩፒ", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "የጃፓን የን", Symbol: "JP¥"},
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
				currency.RUB: cldr.Currency{DisplayName: "የራሻ ሩብል", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
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
				currency.USD: cldr.Currency{DisplayName: "የአሜሪካን ዶላር", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AF:    "አፍሪቃንሰኛ",
			language.AM:    "አምሐረኛ",
			language.AR:    "ዓረበኛ",
			language.AZ:    "አዜርባይጃንኛ",
			language.BE:    "ቤላራሻኛ",
			language.BG:    "ቡልጋሪኛ",
			language.BN:    "በንጋሊኛ",
			language.BR:    "ብሬቶን",
			language.BS:    "ቦስኒያን",
			language.CA:    "ካታላን",
			language.CS:    "ቼክኛ",
			language.CY:    "ወልሽ",
			language.DA:    "ዴኒሽ",
			language.DE:    "ጀርመን",
			language.EL:    "ግሪከኛ",
			language.EN:    "እንግሊዝኛ",
			language.EN_GB: "እንግሊዝኛ (GB)",
			language.EN_US: "እንግሊዝኛ (US)",
			language.EO:    "ኤስፐራንቶ",
			language.ES:    "ስፓኒሽ",
			language.ET:    "ኤስቶኒአን",
			language.EU:    "ባስክኛ",
			language.FA:    "ፐርሲያኛ",
			language.FI:    "ፊኒሽ",
			language.FIL:   "ታጋሎገኛ",
			language.FO:    "ፋሮኛ",
			language.FR:    "ፈረንሳይኛ",
			language.FY:    "ፍሪሰኛ",
			language.GA:    "አይሪሽ",
			language.GD:    "እስኮትስ ጌልክኛ",
			language.GL:    "ጋለቪኛ",
			language.GN:    "ጓራኒ",
			language.GU:    "ጉጃራቲኛ",
			language.HE:    "ዕብራስጥ",
			language.HI:    "ሕንደኛ",
			language.HR:    "ክሮሽያንኛ",
			language.HU:    "ሀንጋሪኛ",
			language.IA:    "ኢንቴር ቋንቋ",
			language.ID:    "እንዶኑሲኛ",
			language.IS:    "አይስላንደኛ",
			language.IT:    "ጣሊያንኛ",
			language.JA:    "ጃፓንኛ",
			language.JV:    "ጃቫንኛ",
			language.KA:    "ጊዮርጊያኛ",
			language.KN:    "ካማደኛ",
			language.KO:    "ኮሪያኛ",
			language.KU:    "ኩርድሽ",
			language.KY:    "ኪሩጋዚ",
			language.LA:    "ላቲንኛ",
			language.LT:    "ሊቱአኒየን",
			language.LV:    "ላቲቪያን",
			language.MK:    "ማክዶኒኛ",
			language.ML:    "ማላያላምኛ",
			language.MR:    "ማራቲኛ",
			language.MS:    "ማላይኛ",
			language.MT:    "ማልቲስኛ",
			language.NE:    "ኔፖሊኛ",
			language.NL:    "ደች",
			language.NN:    "ኖርዌይኛ (ናይ ኝኖርስክ)",
			language.NO:    "ኖርዌጂያን",
			language.OC:    "ኦኪታንኛ",
			language.OR:    "ኦሪያ",
			language.PA:    "ፑንጃቢኛ",
			language.PL:    "ፖሊሽ",
			language.PS:    "ፓሽቶ",
			language.PT:    "ፖርቱጋሊኛ",
			language.PT_BR: "ፖርቱጋልኛ (ናይ ብራዚል)",
			language.PT_PT: "ፖርቱጋልኛ (ናይ ፖርቱጋል)",
			language.RO:    "ሮማኒያን",
			language.RU:    "ራሽኛ",
			language.SH:    "ሰርቦ- ክሮዊታን",
			language.SI:    "ስንሃልኛ",
			language.SK:    "ስሎቨክኛ",
			language.SL:    "ስቁቪኛ",
			language.SQ:    "አልቤኒኛ",
			language.SR:    "ሰርቢኛ",
			language.ST:    "ሰሴቶ",
			language.SU:    "ሱዳንኛ",
			language.SV:    "ስዊድንኛ",
			language.SW:    "ሰዋሂሊኛ",
			language.TA:    "ታሚልኛ",
			language.TE:    "ተሉጉኛ",
			language.TH:    "ታይኛ",
			language.TI:    "ትግርኛ",
			language.TK:    "ናይ ቱርኪ ሰብዓይ (ቱርካዊ)",
			language.TLH:   "ክሊንግኦንኛ",
			language.TR:    "ቱርከኛ",
			language.TW:    "ትዊ",
			language.UK:    "ዩክረኒኛ",
			language.UR:    "ኡርዱኛ",
			language.UZ:    "ኡዝበክኛ",
			language.VI:    "ቪትናምኛ",
			language.XH:    "ዞሳኛ",
			language.YI:    "ዪዲሽ",
			language.ZU:    "ዙሉኛ",
		},
		Territories: cldr.Territories{
			territory.V_001: "ዓለም",
			territory.V_002: "አፍሪካ",
			territory.V_005: "ደቡባዊ አሜሪካ",
			territory.V_009: "ኦሽኒያ",
			territory.V_011: "ምዕራባዊ አፍሪካ",
			territory.V_014: "ምስራቃዊ አፍሪካ",
			territory.V_015: "ሰሜናዊ አፍሪካ",
			territory.V_017: "መካከለኛ አፍሪካ",
			territory.V_018: "ደቡባዊ አፍሪካ",
			territory.V_019: "አሜሪካዎች",
			territory.V_021: "ሰሜናዊ አሜሪካ",
			territory.V_029: "ካሪቢያን",
			territory.V_034: "ምሥራቃዊ እስያ",
			territory.V_039: "ደቡባዊ አውሮፓ",
			territory.V_053: "አውስትራሊያ እና ኒው ዚላንድ",
			territory.V_054: "ሜላኔሲያ",
			territory.V_061: "ፖሊኔዢያ",
			territory.V_142: "እስያ",
			territory.V_145: "ምዕራባዊ እስያ",
			territory.V_150: "አውሮፓ",
			territory.V_151: "ምስራቃዊ አውሮፓ",
			territory.V_154: "ሰሜናዊ አውሮፓ",
			territory.V_155: "ምዕራባዊ አውሮፓ",
			territory.AC:    "አሴንሽን ደሴት",
			territory.AD:    "አንዶራ",
			territory.AE:    "ሕቡራት ኢማራት ዓረብ",
			territory.AF:    "አፍጋኒስታን",
			territory.AG:    "ኣንቲጓን ባሩዳን",
			territory.AI:    "አንጉኢላ",
			territory.AL:    "አልባኒያ",
			territory.AM:    "አርሜኒያ",
			territory.AO:    "አንጐላ",
			territory.AQ:    "አንታርክቲካ",
			territory.AR:    "አርጀንቲና",
			territory.AS:    "ናይ ኣሜሪካ ሳሞኣ",
			territory.AT:    "ኦስትሪያ",
			territory.AU:    "አውስትሬሊያ",
			territory.AW:    "አሩባ",
			territory.AX:    "ደሴታት ኣላንድ",
			territory.AZ:    "አዘርባጃን",
			territory.BA:    "ቦዝንያን ሄርዘጎቪናን",
			territory.BB:    "ባርቤዶስ",
			territory.BD:    "ባንግላዲሽ",
			territory.BE:    "ቤልጄም",
			territory.BF:    "ቡርኪና ፋሶ",
			territory.BG:    "ቡልጌሪያ",
			territory.BH:    "ባህሬን",
			territory.BI:    "ብሩንዲ",
			territory.BJ:    "ቤኒን",
			territory.BL:    "ቅዱስ ባርተለሚይ",
			territory.BM:    "ቤርሙዳ",
			territory.BN:    "ብሩኒ",
			territory.BO:    "ቦሊቪያ",
			territory.BQ:    "ካሪቢያን ኔዘርላንድስ",
			territory.BR:    "ብራዚል",
			territory.BS:    "ባሃማስ",
			territory.BT:    "ቡህታን",
			territory.BV:    "ደሴታት ቦውቬት",
			territory.BW:    "ቦትስዋና",
			territory.BY:    "ቤላሩስ",
			territory.BZ:    "ቤሊዘ",
			territory.CA:    "ካናዳ",
			territory.CC:    "ኮኮስ ኬሊንግ ደሴቶች",
			territory.CD:    "ኮንጎ",
			territory.CF:    "ማእከላይ ኣፍሪቃ ሪፓብሊክ",
			territory.CG:    "ኮንጎ ሪፓብሊክ",
			territory.CH:    "ስዊዘርላንድ",
			territory.CI:    "አይቮሪ ኮስት",
			territory.CK:    "ደሴታት ኩክ",
			territory.CL:    "ቺሊ",
			territory.CM:    "ካሜሩን",
			territory.CN:    "ቻይና",
			territory.CO:    "ኮሎምቢያ",
			territory.CP:    "ክሊፐርቶን ደሴት",
			territory.CR:    "ኮስታ ሪካ",
			territory.CU:    "ኩባ",
			territory.CV:    "ኬፕ ቬርዴ",
			territory.CW:    "ኩራካዎ",
			territory.CX:    "ደሴታት ክሪስትማስ",
			territory.CY:    "ሳይፕረስ",
			territory.CZ:    "ቼክ ሪፓብሊክ",
			territory.DE:    "ጀርመን",
			territory.DG:    "ዲየጎ ጋርሺያ",
			territory.DJ:    "ጂቡቲ",
			territory.DK:    "ዴንማርክ",
			territory.DM:    "ዶሚኒካ",
			territory.DO:    "ዶመኒካ ሪፓብሊክ",
			territory.DZ:    "አልጄሪያ",
			territory.EA:    "ሲውታን ሜሊላን",
			territory.EC:    "ኢኳዶር",
			territory.EE:    "ኤስቶኒያ",
			territory.EG:    "ግብጽ",
			territory.EH:    "ምዕራባዊ ሳህራ",
			territory.ER:    "ኤርትራ",
			territory.ES:    "ስፔን",
			territory.ET:    "ኢትዮጵያ",
			territory.FI:    "ፊንላንድ",
			territory.FJ:    "ፊጂ",
			territory.FK:    "ደሴታት ፎክላንድ",
			territory.FM:    "ሚክሮኔዢያ",
			territory.FO:    "ደሴታት ፋራኦ",
			territory.FR:    "ፈረንሳይ",
			territory.GA:    "ጋቦን",
			territory.GB:    "ዩኬይ",
			territory.GD:    "ግሬናዳ",
			territory.GE:    "ጆርጂያ",
			territory.GF:    "ናይ ፈረንሳይ ጉይና",
			territory.GG:    "ገርንሲ",
			territory.GH:    "ጋና",
			territory.GI:    "ጊብራልታር",
			territory.GL:    "ግሪንላንድ",
			territory.GM:    "ጋምቢያ",
			territory.GN:    "ጊኒ",
			territory.GP:    "ጉዋደሉፕ",
			territory.GQ:    "ኢኳቶሪያል ጊኒ",
			territory.GR:    "ግሪክ",
			territory.GS:    "ደሴታት ደቡብ ጆርጂያን ደቡድ ሳንድዊችን",
			territory.GT:    "ጉዋቲማላ",
			territory.GU:    "ጉዋም",
			territory.GW:    "ቢሳዎ",
			territory.GY:    "ጉያና",
			territory.HK:    "ሆንግ ኮንግ",
			territory.HM:    "ደሴታት ሀርድን ማክዶናልድን",
			territory.HN:    "ሆንዱራስ",
			territory.HR:    "ክሮኤሽያ",
			territory.HT:    "ሀይቲ",
			territory.HU:    "ሀንጋሪ",
			territory.IC:    "ደሴታት ካናሪ",
			territory.ID:    "ኢንዶኔዢያ",
			territory.IE:    "አየርላንድ",
			territory.IL:    "እስራኤል",
			territory.IM:    "አይል ኦፍ ማን",
			territory.IN:    "ህንዲ",
			territory.IO:    "ናይ ብሪጣንያ ህንዳዊ ውቅያኖስ ግዝኣት",
			territory.IQ:    "ኢራቅ",
			territory.IR:    "ኢራን",
			territory.IS:    "አይስላንድ",
			territory.IT:    "ጣሊያን",
			territory.JE:    "ጀርሲ",
			territory.JM:    "ጃማይካ",
			territory.JO:    "ጆርዳን",
			territory.JP:    "ጃፓን",
			territory.KE:    "ኬንያ",
			territory.KG:    "ኪርጂስታን",
			territory.KH:    "ካምቦዲያ",
			territory.KI:    "ኪሪባቲ",
			territory.KM:    "ኮሞሮስ",
			territory.KN:    "ቅዱስ ኪትስን ኔቪስን",
			territory.KP:    "ሰሜን ኮሪያ",
			territory.KR:    "ደቡብ ኮሪያ",
			territory.KW:    "ክዌት",
			territory.KY:    "ካይማን ደሴቶች",
			territory.KZ:    "ካዛኪስታን",
			territory.LA:    "ላኦስ",
			territory.LB:    "ሊባኖስ",
			territory.LC:    "ሴንት ሉቺያ",
			territory.LI:    "ሊችተንስታይን",
			territory.LK:    "ሲሪላንካ",
			territory.LR:    "ላይቤሪያ",
			territory.LS:    "ሌሶቶ",
			territory.LT:    "ሊቱዌኒያ",
			territory.LU:    "ሉክሰምበርግ",
			territory.LV:    "ላትቪያ",
			territory.LY:    "ሊቢያ",
			territory.MA:    "ሞሮኮ",
			territory.MC:    "ሞናኮ",
			territory.MD:    "ሞልዶቫ",
			territory.ME:    "ሞንቴኔግሮ",
			territory.MF:    "ሴንት ማርቲን",
			territory.MG:    "ማዳጋስካር",
			territory.MH:    "ማርሻል አይላንድ",
			territory.MK:    "ሰሜን መቆዶንያ",
			territory.ML:    "ማሊ",
			territory.MM:    "ማያንማር",
			territory.MN:    "ሞንጎሊያ",
			territory.MO:    "ማካው",
			territory.MP:    "ደሴታት ሰሜናዊ ማሪያና",
			territory.MQ:    "ማርቲኒክ",
			territory.MR:    "ሞሪቴኒያ",
			territory.MS:    "ሞንትሴራት",
			territory.MT:    "ማልታ",
			territory.MU:    "ማሩሸስ",
			territory.MV:    "ማልዲቭስ",
			territory.MW:    "ማላዊ",
			territory.MX:    "ሜክሲኮ",
			territory.MY:    "ማሌዢያ",
			territory.MZ:    "ሞዛምቢክ",
			territory.NA:    "ናሚቢያ",
			territory.NC:    "ኒው ካሌዶኒያ",
			territory.NE:    "ኒጀር",
			territory.NF:    "ኖርፎልክ ደሴት",
			territory.NG:    "ናይጄሪያ",
			territory.NI:    "ኒካራጓ",
			territory.NL:    "ኔዘርላንድስ",
			territory.NO:    "ኖርዌ",
			territory.NP:    "ኔፓል",
			territory.NR:    "ናኡሩ",
			territory.NU:    "ኒኡይ",
			territory.NZ:    "ኒው ዚላንድ",
			territory.OM:    "ኦማን",
			territory.PA:    "ፓናማ",
			territory.PE:    "ፔሩ",
			territory.PF:    "ናይ ፈረንሳይ ፖሊነዝያ",
			territory.PG:    "ፓፑዋ ኒው ጊኒ",
			territory.PH:    "ፊሊፒንስ",
			territory.PK:    "ፓኪስታን",
			territory.PL:    "ፖላንድ",
			territory.PM:    "ቅዱስ ፒዬርን ሚኩኤሎን",
			territory.PN:    "ፒትካኢርን",
			territory.PR:    "ፖርታ ሪኮ",
			territory.PS:    "ፍልስጤም",
			territory.PT:    "ፖርቱጋል",
			territory.PW:    "ፓላው",
			territory.PY:    "ፓራጓይ",
			territory.QA:    "ቀጠር",
			territory.QO:    "ወጣ ያለ ኦሽኒያ",
			territory.RE:    "ሪዩኒየን",
			territory.RO:    "ሮሜኒያ",
			territory.RS:    "ሰርቢያ",
			territory.RU:    "ራሺያ",
			territory.RW:    "ሩዋንዳ",
			territory.SA:    "ስዑዲ ዓረብ",
			territory.SB:    "ሰሎሞን ደሴት",
			territory.SC:    "ሲሼልስ",
			territory.SD:    "ሱዳን",
			territory.SE:    "ስዊድን",
			territory.SG:    "ሲንጋፖር",
			territory.SH:    "ሴንት ሄለና",
			territory.SI:    "ስሎቬኒያ",
			territory.SJ:    "ስቫልባርድን ዣን ማየን ደሴታት",
			territory.SK:    "ስሎቫኪያ",
			territory.SL:    "ሴራሊዮን",
			territory.SM:    "ሳን ማሪኖ",
			territory.SN:    "ሴኔጋል",
			territory.SO:    "ሱማሌ",
			territory.SR:    "ሱሪናም",
			territory.SS:    "ደቡብ ሱዳን",
			territory.ST:    "ሳኦ ቶሜን ፕሪንሲፔን",
			territory.SV:    "ኤል ሳልቫዶር",
			territory.SX:    "ሲንት ማርቲን",
			territory.SY:    "ሲሪያ",
			territory.SZ:    "ሱዋዚላንድ",
			territory.TA:    "ትሪስን ዳ ኩንሃ",
			territory.TC:    "ደሴታት ቱርክን ካይኮስን",
			territory.TD:    "ጫድ",
			territory.TF:    "ናይ ፈረንሳይ ደቡባዊ ግዝኣታት",
			territory.TG:    "ቶጐ",
			territory.TH:    "ታይላንድ",
			territory.TJ:    "ታጃኪስታን",
			territory.TK:    "ቶክላው",
			territory.TL:    "ምብራቕ ቲሞር",
			territory.TM:    "ቱርክሜኒስታን",
			territory.TN:    "ቱኒዚያ",
			territory.TO:    "ቶንጋ",
			territory.TR:    "ቱርክ",
			territory.TT:    "ትሪኒዳድን ቶባጎን",
			territory.TV:    "ቱቫሉ",
			territory.TW:    "ታይዋን",
			territory.TZ:    "ታንዛኒያ",
			territory.UA:    "ዩክሬን",
			territory.UG:    "ዩጋንዳ",
			territory.UM:    "ናይ ኣሜሪካ ፍንትት ዝበሉ ደሴታት",
			territory.US:    "ዩኤስ",
			territory.UY:    "ኡራጓይ",
			territory.UZ:    "ዩዝበኪስታን",
			territory.VA:    "ቫቲካን",
			territory.VC:    "ቅዱስ ቪንሴንትን ግሬናዲንስን",
			territory.VE:    "ቬንዙዌላ",
			territory.VG:    "ቨርጂን ደሴታት እንግሊዝ",
			territory.VI:    "ቨርጂን ደሴታት ኣሜሪካ",
			territory.VN:    "ቬትናም",
			territory.VU:    "ቫኑአቱ",
			territory.WF:    "ዋሊስን ፉቱናን",
			territory.WS:    "ሳሞአ",
			territory.XK:    "ኮሶቮ",
			territory.YE:    "የመን",
			territory.YT:    "ሜይኦቴ",
			territory.ZA:    "ደቡብ አፍሪካ",
			territory.ZM:    "ዛምቢያ",
			territory.ZW:    "ዚምቧቤ",
		},
	}
}
