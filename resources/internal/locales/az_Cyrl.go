// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_az_Cyrl() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "az_Cyrl",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "d MMMM y, EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "јан", Feb: "фев", Mar: "мар", Apr: "апр", May: "май", Jun: "ијн", Jul: "ијл", Aug: "авг", Sep: "сен", Oct: "окт", Nov: "ној", Dec: "дек"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Јанвар", Feb: "Феврал", Mar: "Март", Apr: "Апрел", May: "Май", Jun: "Ијун", Jul: "Ијул", Aug: "Август", Sep: "Сентјабр", Oct: "Октјабр", Nov: "Нојабр", Dec: "Декабр"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Б.", Mon: "Б.Е.", Tue: "Ч.А.", Wed: "Ч.", Thu: "Ҹ.А.", Fri: "Ҹ.", Sat: "Ш."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "7", Mon: "1", Tue: "2", Wed: "3", Thu: "4", Fri: "5", Sat: "6"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "базар", Mon: "базар ертәси", Tue: "чәршәнбә ахшамы", Wed: "чәршәнбә", Thu: "ҹүмә ахшамы", Fri: "ҹүмә", Sat: "шәнбә"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "АМ", PM: "ПМ"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "а", PM: "п"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "манат", Symbol: "₼"},
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
				currency.CNY: cldr.Currency{DisplayName: "", Symbol: "CN¥"},
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
				currency.INR: cldr.Currency{DisplayName: "", Symbol: "₹"},
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
				currency.USD: cldr.Currency{DisplayName: "", Symbol: "US$"},
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
			language.AA:      "афар",
			language.AB:      "абхаз",
			language.ACE:     "акин",
			language.ADA:     "адангме",
			language.ADY:     "адуҝе",
			language.AF:      "африкаанс",
			language.AGQ:     "агһем",
			language.AIN:     "ајну",
			language.AK:      "акан",
			language.ALE:     "алеут",
			language.ALT:     "ҹәнуби алтај",
			language.AM:      "амһар",
			language.AN:      "арагон",
			language.ANP:     "анҝика",
			language.AR:      "әрәб",
			language.AR_001:  "мүасир стандарт әрәб",
			language.ARN:     "арауканҹа",
			language.ARP:     "арапаһо",
			language.AS:      "ассам",
			language.ASA:     "асу",
			language.AST:     "астурија",
			language.AV:      "авар",
			language.AWA:     "авадһи",
			language.AY:      "ајмара",
			language.AZ:      "азәрбајҹан",
			language.BA:      "башгырд",
			language.BAN:     "балли",
			language.BAS:     "баса",
			language.BE:      "беларус",
			language.BEM:     "бемба",
			language.BEZ:     "бена",
			language.BG:      "булгар",
			language.BHO:     "бхочпури",
			language.BI:      "бислама",
			language.BIN:     "бини",
			language.BLA:     "сиксикә",
			language.BM:      "бамбара",
			language.BN:      "бенгал",
			language.BO:      "тибет",
			language.BR:      "бретон",
			language.BRX:     "бодо",
			language.BS:      "босниак",
			language.BUG:     "буҝин",
			language.BYN:     "блин",
			language.CA:      "каталан",
			language.CE:      "чечен",
			language.CEB:     "себуан",
			language.CGG:     "чига",
			language.CH:      "чаморо",
			language.CHK:     "чукиз",
			language.CHM:     "мари",
			language.CHO:     "чоктау",
			language.CHR:     "чероки",
			language.CHY:     "чејен",
			language.CKB:     "соран",
			language.CO:      "корсика",
			language.CRS:     "сејшел креолу",
			language.CS:      "чех",
			language.CU:      "славјан",
			language.CV:      "чуваш",
			language.CY:      "уелс",
			language.DA:      "данимарка",
			language.DAK:     "дакота",
			language.DAR:     "даргва",
			language.DAV:     "таита",
			language.DE:      "алман",
			language.DE_AT:   "Австрија алманҹасы",
			language.DE_CH:   "Исвечрә јүксәк алманҹасы",
			language.DGR:     "догриб",
			language.DJE:     "зарма",
			language.DSB:     "ашағы сорб",
			language.DUA:     "дуала",
			language.DV:      "малдив",
			language.DYO:     "диола",
			language.DZ:      "дзонга",
			language.DZG:     "дазага",
			language.EBU:     "ембу",
			language.EE:      "еве",
			language.EFI:     "ефик",
			language.EKA:     "екаҹук",
			language.EL:      "јунан",
			language.EN:      "инҝилис",
			language.EN_AU:   "Австралија инҝилисҹәси",
			language.EN_CA:   "Канада инҝилисҹәси",
			language.EN_GB:   "инҝилис (Б.К.)",
			language.EN_US:   "инҝилис (АБШ)",
			language.EO:      "есперанто",
			language.ES:      "испан",
			language.ES_419:  "Латын Америкасы испанҹасы",
			language.ES_ES:   "Кастилија испанҹасы",
			language.ES_MX:   "Мексика испанҹасы",
			language.ET:      "естон",
			language.EU:      "баск",
			language.EWO:     "евондо",
			language.FA:      "фарс",
			language.FF:      "фула",
			language.FI:      "фин",
			language.FIL:     "филиппин",
			language.FJ:      "фиҹи",
			language.FO:      "фарер",
			language.FON:     "фон",
			language.FR:      "франсыз",
			language.FR_CA:   "Канада франсызҹасы",
			language.FR_CH:   "Исвечрә франсызҹасы",
			language.FUR:     "фриул",
			language.FY:      "гәрби фриз",
			language.GA:      "ирланд",
			language.GAA:     "га",
			language.GD:      "шотланд келт",
			language.GEZ:     "гез",
			language.GIL:     "гилберт",
			language.GL:      "галисија",
			language.GN:      "гуарани",
			language.GOR:     "горонтало",
			language.GSW:     "Исвечрә алманҹасы",
			language.GU:      "гуҹарат",
			language.GUZ:     "гуси",
			language.GV:      "манкс",
			language.GWI:     "гвичин",
			language.HA:      "һауса",
			language.HAW:     "һавај",
			language.HE:      "иврит",
			language.HI:      "һинд",
			language.HIL:     "һилигајнон",
			language.HMN:     "монг",
			language.HR:      "хорват",
			language.HSB:     "јухары сорб",
			language.HT:      "һаити креол",
			language.HU:      "маҹар",
			language.HUP:     "һупа",
			language.HY:      "ермәни",
			language.HZ:      "һереро",
			language.IA:      "интерлингве",
			language.IBA:     "ибан",
			language.IBB:     "ибибио",
			language.ID:      "индонезија",
			language.IG:      "игбо",
			language.ILO:     "илоко",
			language.INH:     "ингуш",
			language.IO:      "идо",
			language.IS:      "исланд",
			language.IT:      "италјан",
			language.IU:      "инуктитут",
			language.JA:      "јапон",
			language.JBO:     "лоғбан",
			language.JGO:     "нгомба",
			language.JMC:     "мачам",
			language.JV:      "јава",
			language.KA:      "ҝүрҹү",
			language.KAB:     "кабиле",
			language.KAC:     "качин",
			language.KAJ:     "жу",
			language.KAM:     "камба",
			language.KBD:     "кабарда-чәркәз",
			language.KCG:     "тви",
			language.KDE:     "маконде",
			language.KEA:     "кабувердиан",
			language.KFO:     "коро",
			language.KHA:     "хази",
			language.KHQ:     "којра чиини",
			language.KI:      "кикују",
			language.KJ:      "куанјама",
			language.KK:      "газах",
			language.KKJ:     "како",
			language.KL:      "калааллисут",
			language.KLN:     "каленҹин",
			language.KM:      "кхмер",
			language.KMB:     "кимбунду",
			language.KN:      "каннада",
			language.KO:      "кореја",
			language.KOK:     "конкани",
			language.KPE:     "кпелле",
			language.KR:      "канури",
			language.KRC:     "гарачај-балкар",
			language.KRL:     "карел",
			language.KRU:     "курух",
			language.KS:      "кәшмир",
			language.KSB:     "шамбала",
			language.KSF:     "бафиа",
			language.KSH:     "көлн",
			language.KU:      "күрд",
			language.KUM:     "кумык",
			language.KV:      "коми",
			language.KW:      "корн",
			language.KY:      "гырғыз",
			language.LA:      "латын",
			language.LAD:     "сефард",
			language.LAG:     "ланҝи",
			language.LB:      "лүксембург",
			language.LEZ:     "ләзҝи",
			language.LG:      "ганда",
			language.LI:      "лимбург",
			language.LKT:     "лакота",
			language.LN:      "лингала",
			language.LO:      "лаос",
			language.LOZ:     "лози",
			language.LRC:     "шимали лури",
			language.LT:      "литва",
			language.LU:      "луба-катанга",
			language.LUA:     "луба-лулуа",
			language.LUN:     "лунда",
			language.LUO:     "луо",
			language.LUS:     "мизо",
			language.LUY:     "лујиа",
			language.LV:      "латыш",
			language.MAD:     "мадуриз",
			language.MAG:     "магаһи",
			language.MAI:     "маитили",
			language.MAK:     "макасар",
			language.MAS:     "масај",
			language.MDF:     "мокша",
			language.MEN:     "менде",
			language.MER:     "меру",
			language.MFE:     "морисиен",
			language.MG:      "малагас",
			language.MGH:     "махува-меетто",
			language.MGO:     "метаʼ",
			language.MH:      "маршал",
			language.MI:      "маори",
			language.MIC:     "микмак",
			language.MIN:     "минангкабан",
			language.MK:      "македон",
			language.ML:      "малајалам",
			language.MN:      "монгол",
			language.MNI:     "манипүри",
			language.MOH:     "моһавк",
			language.MOS:     "моси",
			language.MR:      "маратһи",
			language.MS:      "малај",
			language.MT:      "малта",
			language.MUA:     "мунданг",
			language.MUL:     "чохсајлы дилләр",
			language.MUS:     "крик",
			language.MWL:     "миранд",
			language.MY:      "бирман",
			language.MYV:     "ерзја",
			language.MZN:     "мазандаран",
			language.NA:      "науру",
			language.NAP:     "неаполитан",
			language.NAQ:     "нама",
			language.NB:      "бокмал норвеч",
			language.ND:      "шимали ндебеле",
			language.NDS_NL:  "ашағы саксон",
			language.NE:      "непал",
			language.NEW:     "невари",
			language.NG:      "ндонга",
			language.NIA:     "ниас",
			language.NIU:     "нијуан",
			language.NL:      "һолланд",
			language.NL_BE:   "фламанд",
			language.NMG:     "квасио",
			language.NN:      "нүнорск норвеч",
			language.NNH:     "нҝиембоон",
			language.NOG:     "ногај",
			language.NQO:     "нго",
			language.NR:      "ҹәнуби ндебеле",
			language.NSO:     "шимали сото",
			language.NUS:     "нуер",
			language.NV:      "навајо",
			language.NY:      "нјанҹа",
			language.NYN:     "нјанкол",
			language.OC:      "окситан",
			language.OM:      "оромо",
			language.OR:      "одија",
			language.OS:      "осетин",
			language.PA:      "пәнҹаб",
			language.PAG:     "пангасинан",
			language.PAM:     "пампанга",
			language.PAP:     "папјаменто",
			language.PAU:     "палајан",
			language.PCM:     "ниҝер креол",
			language.PL:      "полјак",
			language.PRG:     "прусс",
			language.PS:      "пушту",
			language.PT:      "португал",
			language.PT_BR:   "Бразилија португалҹасы",
			language.PT_PT:   "Португалија португалҹасы",
			language.QU:      "кечуа",
			language.QUC:     "киче",
			language.RAP:     "рапануи",
			language.RAR:     "раротонган",
			language.RM:      "романш",
			language.RN:      "рунди",
			language.RO:      "румын",
			language.ROF:     "ромбо",
			language.RU:      "рус",
			language.RUP:     "ароман",
			language.RW:      "кинјарванда",
			language.RWK:     "руа",
			language.SA:      "санскрит",
			language.SAD:     "сандаве",
			language.SAH:     "саха",
			language.SAQ:     "самбуру",
			language.SAT:     "сантал",
			language.SBA:     "нгамбај",
			language.SBP:     "сангу",
			language.SC:      "сардин",
			language.SCN:     "сиҹилија",
			language.SCO:     "скотс",
			language.SD:      "синдһи",
			language.SE:      "шимали сами",
			language.SEH:     "сена",
			language.SES:     "којраборо сенни",
			language.SG:      "санго",
			language.SHI:     "тачелит",
			language.SHN:     "шан",
			language.SI:      "синһала",
			language.SK:      "словак",
			language.SL:      "словен",
			language.SM:      "самоа",
			language.SMA:     "ҹәнуби сами",
			language.SMJ:     "луле сами",
			language.SMN:     "инари сами",
			language.SMS:     "сколт сами",
			language.SN:      "шона",
			language.SNK:     "сонинке",
			language.SO:      "сомали",
			language.SQ:      "албан",
			language.SR:      "серб",
			language.SRN:     "сранан тонго",
			language.SS:      "свати",
			language.SSY:     "саһо",
			language.ST:      "сесото",
			language.SU:      "сундан",
			language.SUK:     "сукума",
			language.SV:      "исвеч",
			language.SW:      "суаһили",
			language.SW_CD:   "Конго суаһилиҹәси",
			language.SWB:     "комор",
			language.SYR:     "сурија",
			language.TA:      "тамил",
			language.TE:      "телугу",
			language.TEM:     "тимне",
			language.TEO:     "тесо",
			language.TET:     "тетум",
			language.TG:      "таҹик",
			language.TH:      "тај",
			language.TI:      "тигрин",
			language.TIG:     "тигре",
			language.TK:      "түркмән",
			language.TLH:     "клингон",
			language.TN:      "свана",
			language.TO:      "тонган",
			language.TPI:     "ток писин",
			language.TR:      "түрк",
			language.TRV:     "тароко",
			language.TS:      "сонга",
			language.TT:      "татар",
			language.TUM:     "тумбука",
			language.TVL:     "тувалу",
			language.TWQ:     "тасаваг",
			language.TY:      "тахити",
			language.TYV:     "тувинјан",
			language.TZM:     "Мәркәзи Атлас тамазиҹәси",
			language.UDM:     "удмурт",
			language.UG:      "ујғур",
			language.UK:      "украјна",
			language.UMB:     "умбунду",
			language.UND:     "намәлум дил",
			language.UR:      "урду",
			language.UZ:      "өзбәк",
			language.VAI:     "ваи",
			language.VE:      "венда",
			language.VI:      "вјетнам",
			language.VO:      "волапүк",
			language.VUN:     "вунјо",
			language.WA:      "валун",
			language.WAE:     "валлес",
			language.WAL:     "валамо",
			language.WAR:     "варај",
			language.WO:      "волоф",
			language.XAL:     "калмык",
			language.XH:      "хоса",
			language.XOG:     "сога",
			language.YAV:     "јангбен",
			language.YBB:     "јемба",
			language.YI:      "идиш",
			language.YO:      "јоруба",
			language.YUE:     "кантон",
			language.ZGH:     "тамази",
			language.ZH:      "чин",
			language.ZH_HANS: "садәләшмиш чин",
			language.ZH_HANT: "әнәнәви чин",
			language.ZU:      "зулу",
			language.ZUN:     "зуни",
			language.ZXX:     "дил мәзмуну јохдур",
			language.ZZA:     "заза",
		},
		Territories: cldr.Territories{
			territory.V_001: "Дүнја",
			territory.V_002: "Африка",
			territory.V_003: "Шимали Америка",
			territory.V_005: "Ҹәнуби Америка",
			territory.V_009: "Океанија",
			territory.V_011: "Гәрби Африка",
			territory.V_013: "Мәркәзи Америка",
			territory.V_014: "Шәрги Африка",
			territory.V_015: "Шимали Африка",
			territory.V_017: "Мәркәзи Африка",
			territory.V_018: "Ҹәнуби Африка",
			territory.V_019: "Америка",
			territory.V_021: "Шимал Америкасы",
			territory.V_029: "Кариб",
			territory.V_030: "Шәрги Асија",
			territory.V_034: "Ҹәнуби Асија",
			territory.V_035: "Ҹәнуб-Шәрги Асија",
			territory.V_039: "Ҹәнуби Авропа",
			territory.V_053: "Австралазија",
			territory.V_054: "Меланезија",
			territory.V_057: "Микронезија Реҝиону",
			territory.V_061: "Полинезија",
			territory.V_142: "Асија",
			territory.V_143: "Мәркәзи Асија",
			territory.V_145: "Гәрби Асија",
			territory.V_150: "Авропа",
			territory.V_151: "Шәрги Авропа",
			territory.V_154: "Шимали Авропа",
			territory.V_155: "Гәрби Авропа",
			territory.V_419: "Латын Америкасы",
			territory.AC:    "Аскенсон адасы",
			territory.AD:    "Андорра",
			territory.AE:    "Бирләшмиш Әрәб Әмирликләри",
			territory.AF:    "Әфганыстан",
			territory.AG:    "Антигуа вә Барбуда",
			territory.AI:    "Анҝилја",
			territory.AL:    "Албанија",
			territory.AM:    "Ермәнистан",
			territory.AO:    "Ангола",
			territory.AQ:    "Антарктика",
			territory.AR:    "Арҝентина",
			territory.AS:    "Америка Самоасы",
			territory.AT:    "Австрија",
			territory.AU:    "Австралија",
			territory.AW:    "Аруба",
			territory.AX:    "Аланд адалары",
			territory.AZ:    "Азәрбајҹан",
			territory.BA:    "Боснија вә Һерсеговина",
			territory.BB:    "Барбадос",
			territory.BD:    "Бангладеш",
			territory.BE:    "Белчика",
			territory.BF:    "Буркина Фасо",
			territory.BG:    "Болгарыстан",
			territory.BH:    "Бәһрејн",
			territory.BI:    "Бурунди",
			territory.BJ:    "Бенин",
			territory.BL:    "Сент-Бартелеми",
			territory.BM:    "Бермуд адалары",
			territory.BN:    "Брунеј",
			territory.BO:    "Боливија",
			territory.BR:    "Бразилија",
			territory.BS:    "Баһам адалары",
			territory.BT:    "Бутан",
			territory.BV:    "Буве адасы",
			territory.BW:    "Ботсвана",
			territory.BY:    "Беларус",
			territory.BZ:    "Белиз",
			territory.CA:    "Канада",
			territory.CC:    "Кокос (Килинг) адалары",
			territory.CD:    "Конго-Киншаса",
			territory.CF:    "Мәркәзи Африка Республикасы",
			territory.CG:    "Конго-Браззавил",
			territory.CH:    "Исвечрә",
			territory.CI:    "Котд’ивуар",
			territory.CK:    "Кук адалары",
			territory.CL:    "Чили",
			territory.CM:    "Камерун",
			territory.CN:    "Чин",
			territory.CO:    "Колумбија",
			territory.CP:    "Клиппертон адасы",
			territory.CR:    "Коста Рика",
			territory.CU:    "Куба",
			territory.CV:    "Кабо-Верде",
			territory.CW:    "Курасао",
			territory.CX:    "Милад адасы",
			territory.CY:    "Кипр",
			territory.CZ:    "Чехија",
			territory.DE:    "Алманија",
			territory.DG:    "Диего Гарсија",
			territory.DJ:    "Ҹибути",
			territory.DK:    "Данимарка",
			territory.DM:    "Доминика",
			territory.DO:    "Доминикан Республикасы",
			territory.DZ:    "Әлҹәзаир",
			territory.EA:    "Сеута вә Мелилја",
			territory.EC:    "Еквадор",
			territory.EE:    "Естонија",
			territory.EG:    "Мисир",
			territory.ER:    "Еритреја",
			territory.ES:    "Испанија",
			territory.ET:    "Ефиопија",
			territory.EU:    "Авропа Бирлији",
			territory.FI:    "Финландија",
			territory.FJ:    "Фиҹи",
			territory.FK:    "Фолкленд адалары",
			territory.FM:    "Микронезија",
			territory.FO:    "Фарер адалары",
			territory.FR:    "Франса",
			territory.GA:    "Габон",
			territory.GB:    "Бирләшмиш Краллыг",
			territory.GD:    "Гренада",
			territory.GE:    "Ҝүрҹүстан",
			territory.GF:    "Франса Гвианасы",
			territory.GG:    "Ҝернси",
			territory.GH:    "Гана",
			territory.GI:    "Ҹәбәллүтариг",
			territory.GL:    "Гренландија",
			territory.GM:    "Гамбија",
			territory.GN:    "Гвинеја",
			territory.GP:    "Гваделупа",
			territory.GQ:    "Екваториал Гвинеја",
			territory.GR:    "Јунаныстан",
			territory.GS:    "Ҹәнуби Ҹорҹија вә Ҹәнуби Сендвич адалары",
			territory.GT:    "Гватемала",
			territory.GU:    "Гуам",
			territory.GW:    "Гвинеја-Бисау",
			territory.GY:    "Гајана",
			territory.HK:    "Һонк Конг Хүсуси Инзибати Әрази Чин",
			territory.HM:    "Һерд вә Макдоналд адалары",
			territory.HN:    "Һондурас",
			territory.HR:    "Хорватија",
			territory.HT:    "Һаити",
			territory.HU:    "Маҹарыстан",
			territory.IC:    "Канар адалары",
			territory.ID:    "Индонезија",
			territory.IE:    "Ирландија",
			territory.IL:    "Исраил",
			territory.IM:    "Мен адасы",
			territory.IN:    "Һиндистан",
			territory.IO:    "Британтјанын Һинд Океаны Әразиси",
			territory.IQ:    "Ираг",
			territory.IR:    "Иран",
			territory.IS:    "Исландија",
			territory.IT:    "Италија",
			territory.JE:    "Ҹерси",
			territory.JM:    "Јамајка",
			territory.JO:    "Иорданија",
			territory.JP:    "Јапонија",
			territory.KE:    "Кенија",
			territory.KG:    "Гырғызыстан",
			territory.KH:    "Камбоҹа",
			territory.KI:    "Кирибати",
			territory.KM:    "Комор адалары",
			territory.KN:    "Сент-Китс вә Невис",
			territory.KP:    "Шимали Кореја",
			territory.KR:    "Ҹәнуби Кореја",
			territory.KW:    "Күвејт",
			territory.KY:    "Кајман адалары",
			territory.KZ:    "Газахыстан",
			territory.LA:    "Лаос",
			territory.LB:    "Ливан",
			territory.LC:    "Сент-Лусија",
			territory.LI:    "Лихтенштејн",
			territory.LK:    "Шри-Ланка",
			territory.LR:    "Либерија",
			territory.LS:    "Лесото",
			territory.LT:    "Литва",
			territory.LU:    "Лүксембург",
			territory.LV:    "Латвија",
			territory.LY:    "Ливија",
			territory.MA:    "Мәракеш",
			territory.MC:    "Монако",
			territory.MD:    "Молдова",
			territory.ME:    "Монтенегро",
			territory.MF:    "Сент Мартин",
			territory.MG:    "Мадагаскар",
			territory.MH:    "Маршал адалары",
			territory.ML:    "Мали",
			territory.MM:    "Мјанма",
			territory.MN:    "Монголустан",
			territory.MO:    "Макао Хүсуси Инзибати Әрази Чин",
			territory.MP:    "Шимали Мариан адалары",
			territory.MQ:    "Мартиник",
			territory.MR:    "Мавританија",
			territory.MS:    "Монсерат",
			territory.MT:    "Малта",
			territory.MU:    "Маврики",
			territory.MV:    "Малдив адалары",
			territory.MW:    "Малави",
			territory.MX:    "Мексика",
			territory.MY:    "Малајзија",
			territory.MZ:    "Мозамбик",
			territory.NA:    "Намибија",
			territory.NC:    "Јени Каледонија",
			territory.NE:    "Ниҝер",
			territory.NF:    "Норфолк адасы",
			territory.NG:    "Ниҝерија",
			territory.NI:    "Никарагуа",
			territory.NL:    "Нидерланд",
			territory.NO:    "Норвеч",
			territory.NP:    "Непал",
			territory.NR:    "Науру",
			territory.NU:    "Ниуе",
			territory.NZ:    "Јени Зеландија",
			territory.OM:    "Оман",
			territory.PA:    "Панама",
			territory.PE:    "Перу",
			territory.PF:    "Франса Полинезијасы",
			territory.PG:    "Папуа-Јени Гвинеја",
			territory.PH:    "Филиппин",
			territory.PK:    "Пакистан",
			territory.PL:    "Полша",
			territory.PM:    "Мүгәддәс Пјер вә Микелон",
			territory.PN:    "Питкерн адалары",
			territory.PR:    "Пуерто Рико",
			territory.PT:    "Португалија",
			territory.PW:    "Палау",
			territory.PY:    "Парагвај",
			territory.QA:    "Гәтәр",
			territory.QO:    "Узаг Океанија",
			territory.RE:    "Рејунјон",
			territory.RO:    "Румынија",
			territory.RS:    "Сербија",
			territory.RU:    "Русија",
			territory.RW:    "Руанда",
			territory.SA:    "Сәудијјә Әрәбистаны",
			territory.SB:    "Соломон адалары",
			territory.SC:    "Сејшел адалары",
			territory.SD:    "Судан",
			territory.SE:    "Исвеч",
			territory.SG:    "Сингапур",
			territory.SH:    "Мүгәддәс Јелена",
			territory.SI:    "Словенија",
			territory.SJ:    "Свалбард вә Јан-Мајен",
			territory.SK:    "Словакија",
			territory.SL:    "Сјерра-Леоне",
			territory.SM:    "Сан-Марино",
			territory.SN:    "Сенегал",
			territory.SO:    "Сомали",
			territory.SR:    "Суринам",
			territory.SS:    "Ҹәнуби Судан",
			territory.ST:    "Сан-Томе вә Принсипи",
			territory.SV:    "Салвадор",
			territory.SX:    "Синт-Мартен",
			territory.SY:    "Сурија",
			territory.SZ:    "Свазиленд",
			territory.TA:    "Тристан да Кунја",
			territory.TC:    "Төркс вә Кајкос адалары",
			territory.TD:    "Чад",
			territory.TF:    "Франсанын Ҹәнуб Әразиләри",
			territory.TG:    "Того",
			territory.TH:    "Таиланд",
			territory.TJ:    "Таҹикистан",
			territory.TK:    "Токелау",
			territory.TL:    "Шәрги Тимор",
			territory.TM:    "Түркмәнистан",
			territory.TN:    "Тунис",
			territory.TO:    "Тонга",
			territory.TR:    "Түркијә",
			territory.TT:    "Тринидад вә Тобаго",
			territory.TV:    "Тувалу",
			territory.TW:    "Тајван",
			territory.TZ:    "Танзанија",
			territory.UA:    "Украјна",
			territory.UG:    "Уганда",
			territory.UM:    "АБШ-а бағлы кичик адаҹыглар",
			territory.US:    "Америка Бирләшмиш Штатлары",
			territory.UY:    "Уругвај",
			territory.UZ:    "Өзбәкистан",
			territory.VA:    "Ватикан",
			territory.VC:    "Сент-Винсент вә Гренадинләр",
			territory.VE:    "Венесуела",
			territory.VG:    "Британијанын Вирҝин адалары",
			territory.VI:    "АБШ Вирҝин адалары",
			territory.VN:    "Вјетнам",
			territory.VU:    "Вануату",
			territory.WF:    "Уоллис вә Футуна",
			territory.WS:    "Самоа",
			territory.XK:    "Косово",
			territory.YE:    "Јәмән",
			territory.YT:    "Мајот",
			territory.ZA:    "Ҹәнуб Африка",
			territory.ZM:    "Замбија",
			territory.ZW:    "Зимбабве",
			territory.ZZ:    "Намәлум Реҝион",
		},
	}
}
