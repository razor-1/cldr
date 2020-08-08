// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_zh_Hans_SG() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "y年M月d日EEEE", Long: "y年M月d日", Medium: "y年M月d日", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "zzzz ah:mm:ss", Long: "z ah:mm:ss", Medium: "ah:mm:ss", Short: "ah:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "1月", Feb: "2月", Mar: "3月", Apr: "4月", May: "5月", Jun: "6月", Jul: "7月", Aug: "8月", Sep: "9月", Oct: "10月", Nov: "11月", Dec: "12月"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "一月", Feb: "二月", Mar: "三月", Apr: "四月", May: "五月", Jun: "六月", Jul: "七月", Aug: "八月", Sep: "九月", Oct: "十月", Nov: "十一月", Dec: "十二月"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "周日", Mon: "周一", Tue: "周二", Wed: "周三", Thu: "周四", Fri: "周五", Sat: "周六"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "日", Mon: "一", Tue: "二", Wed: "三", Thu: "四", Fri: "五", Sat: "六"}, Short: cldr.CalendarDayFormatNameValue{Sun: "周日", Mon: "周一", Tue: "周二", Wed: "周三", Thu: "周四", Fri: "周五", Sat: "周六"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "星期日", Mon: "星期一", Tue: "星期二", Wed: "星期三", Thu: "星期四", Fri: "星期五", Sat: "星期六"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "上午", PM: "下午"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "上午", PM: "下午"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "上午", PM: "下午"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_zh]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "", Percent: "", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "安道尔比塞塔", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "阿联酋迪拉姆", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "阿富汗尼 (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "阿富汗尼", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "阿尔巴尼亚列克(1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "阿尔巴尼亚列克", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "亚美尼亚德拉姆", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "荷兰安的列斯盾", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "安哥拉宽扎", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "安哥拉宽扎 (1977–1990)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "安哥拉新宽扎 (1990–2000)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "安哥拉重新调整宽扎 (1995–1999)", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "阿根廷奥斯特拉尔", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "阿根廷法定比索 (1970–1983)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "阿根廷比索 (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "阿根廷比索 (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "阿根廷比索", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "奥地利先令", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "澳大利亚元", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "阿鲁巴弗罗林", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "阿塞拜疆马纳特 (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "阿塞拜疆马纳特", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "波士尼亚-赫塞哥维纳第纳尔 (1992–1994)", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "波斯尼亚-黑塞哥维那可兑换马克", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "波士尼亚-赫塞哥维纳新第纳尔 (1994–1997)", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "巴巴多斯元", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "孟加拉塔卡", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "比利时法郎（可兑换）", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "比利时法郎", Symbol: "BEF"},
				currency.BEL: cldr.Currency{DisplayName: "比利时法郎（金融）", Symbol: "BEL"},
				currency.BGL: cldr.Currency{DisplayName: "保加利亚硬列弗", Symbol: "BGL"},
				currency.BGM: cldr.Currency{DisplayName: "保加利亚社会党列弗", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "保加利亚列弗", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "保加利亚列弗 (1879–1952)", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "巴林第纳尔", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "布隆迪法郎", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "百慕大元", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "文莱元", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "玻利维亚诺", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "玻利维亚诺 (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "玻利维亚比索", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "玻利维亚 Mvdol（资金）", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "巴西新克鲁赛罗 (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "巴西克鲁扎多 (1986–1989)", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "巴西克鲁塞罗 (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "巴西雷亚尔", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "巴西新克鲁扎多 (1989–1990)", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "巴西克鲁塞罗 (1993–1994)", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "巴西克鲁塞罗 (1942–1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "巴哈马元", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "不丹努尔特鲁姆", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "缅元", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "博茨瓦纳普拉", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "白俄罗斯新卢布 (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "白俄罗斯卢布", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "白俄罗斯卢布 (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "伯利兹元", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "加拿大元", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "刚果法郎", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "欧元 (WIR)", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "瑞士法郎", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "法郎 (WIR)", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "智利埃斯库多", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "智利（资金）", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "智利比索", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "人民币（离岸）", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "中国人民银行元", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "人民币", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "哥伦比亚比索", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "哥伦比亚币", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "哥斯达黎加科朗", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "旧塞尔维亚第纳尔", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "捷克硬克朗", Symbol: "CSK"},
				currency.CUC: cldr.Currency{DisplayName: "古巴可兑换比索", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "古巴比索", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "佛得角埃斯库多", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "塞浦路斯镑", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "捷克克朗", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "东德奥斯特马克", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "德国马克", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "吉布提法郎", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "丹麦克朗", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "多米尼加比索", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "阿尔及利亚第纳尔", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "厄瓜多尔苏克雷", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "厄瓜多尔 (UVC)", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "爱沙尼亚克朗", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "埃及镑", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "厄立特里亚纳克法", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "西班牙比塞塔（帐户 A）", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "西班牙比塞塔（兑换帐户）", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "西班牙比塞塔", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "埃塞俄比亚比尔", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "欧元", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "芬兰马克", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "斐济元", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "福克兰群岛镑", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "法国法郎", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "英镑", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "乔治亚库蓬拉瑞特", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "格鲁吉亚拉里", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "加纳塞第", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "加纳塞地", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "直布罗陀镑", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "冈比亚达拉西", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "几内亚法郎", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "几内亚西里", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "赤道几内亚埃奎勒", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "希腊德拉克马", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "危地马拉格查尔", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "葡萄牙几内亚埃斯库多", Symbol: "GWE"},
				currency.GWP: cldr.Currency{DisplayName: "几内亚比绍比索", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "圭亚那元", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "港元", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "洪都拉斯伦皮拉", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "克罗地亚第纳尔", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "克罗地亚库纳", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "海地古德", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "匈牙利福林", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "印度尼西亚盾", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "爱尔兰镑", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "以色列镑", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "以色列谢克尔(1980–1985)", Symbol: "ILS"},
				currency.ILS: cldr.Currency{DisplayName: "以色列新谢克尔", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "印度卢比", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "伊拉克第纳尔", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "伊朗里亚尔", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "冰岛克朗(1918–1981)", Symbol: "ISJ"},
				currency.ISK: cldr.Currency{DisplayName: "冰岛克朗", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "意大利里拉", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "牙买加元", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "约旦第纳尔", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "日元", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "肯尼亚先令", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "吉尔吉斯斯坦索姆", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "柬埔寨瑞尔", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "科摩罗法郎", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "朝鲜元", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "韩元 (1953–1962)", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "韩元 (1945–1953)", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "韩元", Symbol: "￦"},
				currency.KWD: cldr.Currency{DisplayName: "科威特第纳尔", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "开曼元", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "哈萨克斯坦坚戈", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "老挝基普", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "黎巴嫩镑", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "斯里兰卡卢比", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "利比里亚元", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "莱索托洛蒂", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "立陶宛立特", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "立陶宛塔咯呐司", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "卢森堡可兑换法郎", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "卢森堡法郎", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "卢森堡金融法郎", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "拉脱维亚拉特", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "拉脱维亚卢布", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "利比亚第纳尔", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "摩洛哥迪拉姆", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "摩洛哥法郎", Symbol: "MAF"},
				currency.MCF: cldr.Currency{DisplayName: "摩纳哥法郎", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "摩尔多瓦库邦", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "摩尔多瓦列伊", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "马达加斯加阿里亚里", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "马达加斯加法郎", Symbol: "MGF"},
				currency.MKD: cldr.Currency{DisplayName: "马其顿第纳尔", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "马其顿第纳尔 (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "马里法郎", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "缅甸元", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "蒙古图格里克", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "澳门币", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "毛里塔尼亚乌吉亚 (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "毛里塔尼亚乌吉亚", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "马耳他里拉", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "马耳他镑", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "毛里求斯卢比", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "马尔代夫卢比(1947–1981)", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "马尔代夫卢菲亚", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "马拉维克瓦查", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "墨西哥比索", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "墨西哥银比索 (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "墨西哥（资金）", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "马来西亚林吉特", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "莫桑比克埃斯库多", Symbol: "MZE"},
				currency.MZM: cldr.Currency{DisplayName: "旧莫桑比克美提卡", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "莫桑比克美提卡", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "纳米比亚元", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "尼日利亚奈拉", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "尼加拉瓜科多巴 (1988–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "尼加拉瓜科多巴", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "荷兰盾", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "挪威克朗", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "尼泊尔卢比", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "新西兰元", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "阿曼里亚尔", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "巴拿马巴波亚", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "秘鲁印第", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "秘鲁索尔", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "秘鲁索尔 (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "巴布亚新几内亚基那", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "菲律宾比索", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "巴基斯坦卢比", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "波兰兹罗提", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "波兰兹罗提 (1950–1995)", Symbol: "PLZ"},
				currency.PTE: cldr.Currency{DisplayName: "葡萄牙埃斯库多", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "巴拉圭瓜拉尼", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "卡塔尔里亚尔", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "罗得西亚元", Symbol: "RHD"},
				currency.ROL: cldr.Currency{DisplayName: "旧罗马尼亚列伊", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "罗马尼亚列伊", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "塞尔维亚第纳尔", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "俄罗斯卢布", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "俄国卢布 (1991–1998)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "卢旺达法郎", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "沙特里亚尔", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "所罗门群岛元", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "塞舌尔卢比", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "苏丹第纳尔 (1992–2007)", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "苏丹镑", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "旧苏丹镑", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "瑞典克朗", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "新加坡元", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "圣赫勒拿群岛磅", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "斯洛文尼亚托拉尔", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "斯洛伐克克朗", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "塞拉利昂利昂", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "索马里先令", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "苏里南元", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "苏里南盾", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "南苏丹镑", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "圣多美和普林西比多布拉 (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "圣多美和普林西比多布拉", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "苏联卢布", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "萨尔瓦多科朗", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "叙利亚镑", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "斯威士兰里兰吉尼", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "泰铢", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "塔吉克斯坦卢布", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "塔吉克斯坦索莫尼", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "土库曼斯坦马纳特 (1993–2009)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "土库曼斯坦马纳特", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "突尼斯第纳尔", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "汤加潘加", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "帝汶埃斯库多", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "土耳其里拉 (1922–2005)", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "土耳其里拉", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "特立尼达和多巴哥元", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "新台币", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "坦桑尼亚先令", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "乌克兰格里夫纳", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "乌克兰币", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "乌干达先令 (1966–1987)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "乌干达先令", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "美元", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "美元（次日）", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "美元（当日）", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "乌拉圭比索（索引单位）", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "乌拉圭比索 (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "乌拉圭比索", Symbol: "UYU"},
				currency.UYW: cldr.Currency{DisplayName: "乌拉圭票面工资指数单位", Symbol: ""},
				currency.UZS: cldr.Currency{DisplayName: "乌兹别克斯坦苏姆", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "委内瑞拉玻利瓦尔 (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "委内瑞拉玻利瓦尔 (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "委内瑞拉玻利瓦尔", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "越南盾", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "越南盾 (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "瓦努阿图瓦图", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "萨摩亚塔拉", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "中非法郎", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "白银", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "黄金", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "欧洲复合单位", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "欧洲货币联盟", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "欧洲计算单位 (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "欧洲计算单位 (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "东加勒比元", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "特别提款权", Symbol: "XDR"},
				currency.XEU: cldr.Currency{DisplayName: "欧洲货币单位", Symbol: "XEU"},
				currency.XFO: cldr.Currency{DisplayName: "法国金法郎", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "法国法郎 (UIC)", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "西非法郎", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "钯", Symbol: "XPD"},
				currency.XPF: cldr.Currency{DisplayName: "太平洋法郎", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "铂", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "RINET 基金", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "苏克雷", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "测试货币代码", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "非洲开发银行记账单位", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "未知货币", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "也门第纳尔", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "也门里亚尔", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "南斯拉夫硬第纳尔 (1966–1990)", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "南斯拉夫新第纳尔 (1994–2002)", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "南斯拉夫可兑换第纳尔 (1990–1992)", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "南斯拉夫改良第纳尔 (1992–1993)", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "南非兰特 (金融)", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "南非兰特", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "赞比亚克瓦查 (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "赞比亚克瓦查", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "新扎伊尔 (1993–1998)", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "扎伊尔 (1971–1993)", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "津巴布韦元 (1980–2008)", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "津巴布韦元 (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "津巴布韦元 (2008)", Symbol: "ZWR"},
			},
		},
	}
}
