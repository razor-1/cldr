// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_gez_ER() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "gez_ER",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE፥ dd MMMM መዓልት y G", Long: "dd MMMM y", Medium: "dd-MMM-y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ጠ", Feb: "ከ", Mar: "መ", Apr: "አ", May: "ግ", Jun: "ሠ", Jul: "ሐ", Aug: "ነ", Sep: "ከ", Oct: "ጠ", Nov: "ኀ", Dec: "ኀ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ጠሐረ", Feb: "ከተተ", Mar: "መገበ", Apr: "አኀዘ", May: "ግንባት", Jun: "ሠንየ", Jul: "ሐመለ", Aug: "ነሐሰ", Sep: "ከረመ", Oct: "ጠቀመ", Nov: "ኀደረ", Dec: "ኀሠሠ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "እ", Mon: "ሰ", Tue: "ሠ", Wed: "ራ", Thu: "ሐ", Fri: "ዓ", Sat: "ቀ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "እኁድ", Mon: "ሰኑይ", Tue: "ሠሉስ", Wed: "ራብዕ", Thu: "ሐሙስ", Fri: "ዓርበ", Sat: "ቀዳሚት"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ጽባሕ", PM: "ምሴት"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: "ወ", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "", Percent: ""},
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
				currency.BRL: cldr.Currency{DisplayName: "የብራዚል ሪል", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "", Symbol: "P"},
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
				currency.ERN: cldr.Currency{DisplayName: "", Symbol: "Nfk"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "የኢትዮጵያ ብር", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "አውሮ", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "የእንግሊዝ ፓውንድ ስተርሊንግ", Symbol: "£"},
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
				currency.INR: cldr.Currency{DisplayName: "የሕንድ ሩፒ", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "የጃፓን የን", Symbol: "JP¥"},
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
				currency.RUB: cldr.Currency{DisplayName: "የራሻ ሩብል", Symbol: "₽"},
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
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AA:    "አፋርኛ",
			language.AB:    "አብሐዚኛ",
			language.AF:    "አፍሪቃንስኛ",
			language.AM:    "አምሐረኛ",
			language.AR:    "ዐርቢኛ",
			language.AS:    "አሳሜዛዊ",
			language.AY:    "አያማርኛ",
			language.AZ:    "አዜርባይጃንኛ",
			language.BA:    "ባስኪርኛ",
			language.BE:    "ቤላራሻኛ",
			language.BG:    "ቡልጋሪኛ",
			language.BI:    "ቢስላምኛ",
			language.BN:    "በንጋሊኛ",
			language.BO:    "ትበትንኛ",
			language.BR:    "ብሬቶንኛ",
			language.BYN:   "ብሊን",
			language.CA:    "ካታላንኛ",
			language.CO:    "ኮርሲካኛ",
			language.CS:    "ቼክኛ",
			language.CY:    "ወልሽ",
			language.DA:    "ዴኒሽ",
			language.DE:    "ጀርመን",
			language.DZ:    "ድዞንግኻኛ",
			language.EL:    "ግሪክኛ",
			language.EN:    "እንግሊዝኛ",
			language.EO:    "ኤስፐራንቶ",
			language.ES:    "ስፓኒሽ",
			language.ET:    "ኤስቶኒአን",
			language.EU:    "ባስክኛ",
			language.FA:    "ፐርሲያኛ",
			language.FI:    "ፊኒሽ",
			language.FJ:    "ፊጂኛ",
			language.FO:    "ፋሮኛ",
			language.FR:    "ፈረንሳይኛ",
			language.FY:    "ፍሪስኛ",
			language.GA:    "አይሪሽ",
			language.GD:    "እስኮትስ፡ጌልክኛ",
			language.GEZ:   "ግዕዝኛ",
			language.GL:    "ጋለጋኛ",
			language.GN:    "ጓራኒኛ",
			language.GU:    "ጉጃርቲኛ",
			language.HA:    "ሃውሳኛ",
			language.HE:    "ዕብራስጥ",
			language.HI:    "ሐንድኛ",
			language.HR:    "ክሮሽያንኛ",
			language.HU:    "ሀንጋሪኛ",
			language.HY:    "አርመናዊ",
			language.IA:    "ኢንቴርሊንጓ",
			language.ID:    "እንዶኒሲኛ",
			language.IE:    "እንተርሊንግወ",
			language.IK:    "እኑፒያቅኛ",
			language.IS:    "አይስላንድኛ",
			language.IT:    "ጣሊያንኛ",
			language.IU:    "እኑክቲቱትኛ",
			language.JA:    "ጃፓንኛ",
			language.JV:    "ጃቫንኛ",
			language.KA:    "ጊዮርጊያን",
			language.KK:    "ካዛክኛ",
			language.KL:    "ካላሊሱትኛ",
			language.KM:    "ክመርኛ",
			language.KN:    "ካናዳኛ",
			language.KO:    "ኮሪያኛ",
			language.KS:    "ካሽሚርኛ",
			language.KU:    "ኩርድሽኛ",
			language.KY:    "ኪርጊዝኛ",
			language.LA:    "ላቲንኛ",
			language.LN:    "ሊንጋላኛ",
			language.LO:    "ላውስኛ",
			language.LT:    "ሊቱአኒያን",
			language.LV:    "ላትቪያን",
			language.MG:    "ማላጋስኛ",
			language.MI:    "ማዮሪኛ",
			language.MK:    "ማከዶኒኛ",
			language.ML:    "ማላያላምኛ",
			language.MN:    "ሞንጎላዊኛ",
			language.MR:    "ማራዚኛ",
			language.MS:    "ማላይኛ",
			language.MT:    "ማልቲስኛ",
			language.MY:    "ቡርማኛ",
			language.NA:    "ናኡሩ",
			language.NE:    "ኔፓሊኛ",
			language.NL:    "ደች",
			language.NO:    "ኖርዌጂያን",
			language.OC:    "ኦኪታንኛ",
			language.OM:    "ኦሮምኛ",
			language.OR:    "ኦሪያኛ",
			language.PA:    "ፓንጃቢኛ",
			language.PL:    "ፖሊሽ",
			language.PS:    "ፑሽቶኛ",
			language.PT:    "ፖርቱጋሊኛ",
			language.QU:    "ኵቿኛ",
			language.RM:    "ሮማንስ",
			language.RN:    "ሩንዲኛ",
			language.RO:    "ሮማኒያን",
			language.RO_MD: "ሞልዳቫዊና",
			language.RU:    "ራሽኛ",
			language.RW:    "ኪንያርዋንድኛ",
			language.SA:    "ሳንስክሪትኛ",
			language.SD:    "ሲንድሂኛ",
			language.SG:    "ሳንጎኛ",
			language.SI:    "ስንሃልኛ",
			language.SID:   "ሲዳምኛ",
			language.SK:    "ስሎቫክኛ",
			language.SL:    "ስሎቪኛ",
			language.SM:    "ሳሞአኛ",
			language.SN:    "ሾናኛ",
			language.SO:    "ሱማልኛ",
			language.SQ:    "ልቤኒኛ",
			language.SR:    "ሰርቢኛ",
			language.SS:    "ስዋቲኛ",
			language.ST:    "ሶዞኛ",
			language.SU:    "ሱዳንኛ",
			language.SV:    "ስዊድንኛ",
			language.SW:    "ስዋሂሊኛ",
			language.TA:    "ታሚልኛ",
			language.TE:    "ተሉጉኛ",
			language.TG:    "ታጂኪኛ",
			language.TH:    "ታይኛ",
			language.TI:    "ትግርኛ",
			language.TIG:   "ትግረ",
			language.TK:    "ቱርክመንኛ",
			language.TL:    "ታጋሎገኛ",
			language.TN:    "ጽዋናዊኛ",
			language.TO:    "ቶንጋ",
			language.TR:    "ቱርክኛ",
			language.TS:    "ጾንጋኛ",
			language.TT:    "ታታርኛ",
			language.TW:    "ትዊኛ",
			language.UG:    "ኡዊግሁርኛ",
			language.UK:    "ዩክረኒኛ",
			language.UR:    "ኡርዱኛ",
			language.UZ:    "ኡዝበክኛ",
			language.VI:    "ቪትናምኛ",
			language.VO:    "ቮላፑክኛ",
			language.WO:    "ዎሎፍኛ",
			language.XH:    "ዞሳኛ",
			language.YI:    "ይዲሻዊኛ",
			language.YO:    "ዮሩባዊኛ",
			language.ZA:    "ዡዋንግኛ",
			language.ZH:    "ቻይንኛ",
			language.ZU:    "ዙሉኛ",
		},
		Territories: cldr.Territories{
			territory.AD: "አንዶራ",
			territory.AE: "የተባበሩት፡አረብ፡ኤምሬትስ",
			territory.AL: "አልባኒያ",
			territory.AM: "አርሜኒያ",
			territory.AR: "አርጀንቲና",
			territory.AT: "ኦስትሪያ",
			territory.AU: "አውስትሬሊያ",
			territory.AZ: "አዘርባጃን",
			territory.BA: "ቦስኒያ፡እና፡ሄርዞጎቪኒያ",
			territory.BB: "ባርቤዶስ",
			territory.BE: "ቤልጄም",
			territory.BG: "ቡልጌሪያ",
			territory.BH: "ባህሬን",
			territory.BM: "ቤርሙዳ",
			territory.BO: "ቦሊቪያ",
			territory.BR: "ብራዚል",
			territory.BT: "ቡህታን",
			territory.BY: "ቤላሩስ",
			territory.BZ: "ቤሊዘ",
			territory.CF: "የመካከለኛው፡አፍሪካ፡ሪፐብሊክ",
			territory.CH: "ስዊዘርላንድ",
			territory.CL: "ቺሊ",
			territory.CM: "ካሜሩን",
			territory.CN: "ቻይና",
			territory.CO: "ኮሎምቢያ",
			territory.CV: "ኬፕ፡ቬርዴ",
			territory.CY: "ሳይፕረስ",
			territory.CZ: "ቼክ፡ሪፑብሊክ",
			territory.DE: "ጀርመን",
			territory.DK: "ዴንማርክ",
			territory.DM: "ዶሚኒካ",
			territory.DO: "ዶሚኒክ፡ሪፑብሊክ",
			territory.DZ: "አልጄሪያ",
			territory.EC: "ኢኳዶር",
			territory.EE: "ኤስቶኒያ",
			territory.EG: "ግብጽ",
			territory.EH: "ምዕራባዊ፡ሳህራ",
			territory.ER: "ኤርትራ",
			territory.ES: "ስፔን",
			territory.ET: "ኢትዮጵያ",
			territory.FI: "ፊንላንድ",
			territory.FJ: "ፊጂ",
			territory.FM: "ሚክሮኔዢያ",
			territory.FR: "ፈረንሳይ",
			territory.GB: "እንግሊዝ",
			territory.GE: "ጆርጂያ",
			territory.GF: "የፈረንሳይ፡ጉዊአና",
			territory.GM: "ጋምቢያ",
			territory.GN: "ጊኒ",
			territory.GQ: "ኢኳቶሪያል፡ጊኒ",
			territory.GR: "ግሪክ",
			territory.GW: "ቢሳዎ",
			territory.GY: "ጉያና",
			territory.HK: "ሆንግ፡ኮንግ",
			territory.HR: "ክሮኤሽያ",
			territory.HT: "ሀይቲ",
			territory.HU: "ሀንጋሪ",
			territory.ID: "ኢንዶኔዢያ",
			territory.IE: "አየርላንድ",
			territory.IL: "እስራኤል",
			territory.IN: "ህንድ",
			territory.IQ: "ኢራቅ",
			territory.IS: "አይስላንድ",
			territory.IT: "ጣሊያን",
			territory.JM: "ጃማይካ",
			territory.JO: "ጆርዳን",
			territory.JP: "ጃፓን",
			territory.KH: "ካምቦዲያ",
			territory.KM: "ኮሞሮስ",
			territory.KP: "ደቡብ፡ኮሪያ",
			territory.KR: "ሰሜን፡ኮሪያ",
			territory.KW: "ክዌት",
			territory.LB: "ሊባኖስ",
			territory.LT: "ሊቱዌኒያ",
			territory.LV: "ላትቪያ",
			territory.LY: "ሊቢያ",
			territory.MA: "ሞሮኮ",
			territory.MD: "ሞልዶቫ",
			territory.MK: "ማከዶኒያ",
			territory.MN: "ሞንጎሊያ",
			territory.MO: "ማካዎ",
			territory.MR: "ሞሪቴኒያ",
			territory.MT: "ማልታ",
			territory.MU: "ማሩሸስ",
			territory.MX: "ሜክሲኮ",
			territory.MY: "ማሌዢያ",
			territory.NA: "ናሚቢያ",
			territory.NC: "ኒው፡ካሌዶኒያ",
			territory.NG: "ናይጄሪያ",
			territory.NL: "ኔዘርላንድ",
			territory.NO: "ኖርዌ",
			territory.NP: "ኔፓል",
			territory.NZ: "ኒው፡ዚላንድ",
			territory.PE: "ፔሩ",
			territory.PF: "የፈረንሳይ፡ፖሊኔዢያ",
			territory.PG: "ፓፑዋ፡ኒው፡ጊኒ",
			territory.PL: "ፖላንድ",
			territory.PR: "ፖርታ፡ሪኮ",
			territory.RO: "ሮሜኒያ",
			territory.RS: "ሰርቢያ",
			territory.RU: "ራሺያ",
			territory.SA: "ሳውድአረቢያ",
			territory.SD: "ሱዳን",
			territory.SE: "ስዊድን",
			territory.SG: "ሲንጋፖር",
			territory.SI: "ስሎቬኒያ",
			territory.SK: "ስሎቫኪያ",
			territory.SN: "ሴኔጋል",
			territory.SO: "ሱማሌ",
			territory.SY: "ሲሪያ",
			territory.TD: "ቻድ",
			territory.TF: "የፈረንሳይ፡ደቡባዊ፡ግዛቶች",
			territory.TH: "ታይላንድ",
			territory.TJ: "ታጃኪስታን",
			territory.TL: "ምስራቅ፡ቲሞር",
			territory.TN: "ቱኒዚያ",
			territory.TR: "ቱርክ",
			territory.TT: "ትሪኒዳድ፡እና፡ቶባጎ",
			territory.TZ: "ታንዛኒያ",
			territory.UG: "ዩጋንዳ",
			territory.US: "አሜሪካ",
			territory.UZ: "ዩዝበኪስታን",
			territory.VE: "ቬንዙዌላ",
			territory.VG: "የእንግሊዝ፡ድንግል፡ደሴቶች",
			territory.VI: "የአሜሪካ፡ቨርጂን፡ደሴቶች",
			territory.YE: "የመን",
			territory.ZA: "ደቡብ፡አፍሪካ",
			territory.ZM: "ዛምቢያ",
		},
	}
}
