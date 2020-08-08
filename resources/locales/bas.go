// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_bas() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "kɔn", Feb: "mac", Mar: "mat", Apr: "mto", May: "mpu", Jun: "hil", Jul: "nje", Aug: "hik", Sep: "dip", Oct: "bio", Nov: "may", Dec: "liɓ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "k", Feb: "m", Mar: "m", Apr: "m", May: "m", Jun: "h", Jul: "n", Aug: "h", Sep: "d", Oct: "b", Nov: "m", Dec: "l"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Kɔndɔŋ", Feb: "Màcɛ̂l", Mar: "Màtùmb", Apr: "Màtop", May: "M̀puyɛ", Jun: "Hìlòndɛ̀", Jul: "Njèbà", Aug: "Hìkaŋ", Sep: "Dìpɔ̀s", Oct: "Bìòôm", Nov: "Màyɛsèp", Dec: "Lìbuy li ńyèe"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "nɔy", Mon: "nja", Tue: "uum", Wed: "ŋge", Thu: "mbɔ", Fri: "kɔɔ", Sat: "jon"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "n", Tue: "u", Wed: "ŋ", Thu: "m", Fri: "k", Sat: "j"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ŋgwà nɔ̂y", Mon: "ŋgwà njaŋgumba", Tue: "ŋgwà ûm", Wed: "ŋgwà ŋgê", Thu: "ŋgwà mbɔk", Fri: "ŋgwà kɔɔ", Sat: "ŋgwà jôn"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "I bikɛ̂glà", PM: "I ɓugajɔp"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "I bikɛ̂glà", PM: "I ɓugajɔp"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_root]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirhàm èmìrâ", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Kwànza àŋgolà", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dɔ̀lâr òstralìà", Symbol: "A$"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinâr Bàraìn", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Frǎŋ bùrundì", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pùla Bòtswanà", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dɔ̀lâr kànadà", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Frǎŋ kòŋgo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Frǎŋ sùwîs", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yùan kinà", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Èskudò kabwe᷆r", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Frǎŋ jìbutì", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dìnâr àlgerìà", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Paùnd ègîptò", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nafkà èrìtrěà", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Bîr ètìopìà", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Èrô", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Stɛrlìŋ ŋgìsì", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sèdi gānà", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasì gambìà", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Frǎŋ gìnê", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rùpi īndìà", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yɛ̂n yàpân", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Silîŋ kenìà", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Frǎŋ kòmorà", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dɔ̀lâr lìberìà", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lotì lèsòtò", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dìnâr libìà", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dìrham màrôk", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Frǎŋ màlàgasì", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ùgwiya mòrìtanìa (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ùgwiya mòrìtanìa", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupìɛ̀ mòrîs", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwaca màlawì", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Mètìkal mòsàmbîk", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dɔ̀lâr nàmibìà", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nayrà nìgerìà", Symbol: "₦"},
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
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Frǎŋ Rùandà", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Rìal sàudì", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rùpiɛ̀ sèsɛ̂l", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dìnâr sùdân", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Paùnd sùdân", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Paùnd hèlenà", Symbol: "£"},
				currency.SLL: cldr.Currency{DisplayName: "Lèonɛ̀", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Silîŋ sòmàli", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobrà sàotòme (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobrà sàotòme", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lìlàŋgeni swàzì", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dìnâr tùnîs", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Silîŋ tànzànià", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Silîŋ ùgàndà", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dɔla àmerkà", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Frǎŋ CFA (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Frǎŋ CFA (BCEAO)", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Rân àfrǐkàsɔ̀", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwàca sàmbià (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwàca sàmbià", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dɔ̀lâr sìmbàbwê", Symbol: ""},
			},
		},
	}
}
