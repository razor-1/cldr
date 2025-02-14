// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_bas_CM() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "bas_CM",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "kɔn", Feb: "mac", Mar: "mat", Apr: "mto", May: "mpu", Jun: "hil", Jul: "nje", Aug: "hik", Sep: "dip", Oct: "bio", Nov: "may", Dec: "liɓ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "k", Feb: "m", Mar: "m", Apr: "m", May: "m", Jun: "h", Jul: "n", Aug: "h", Sep: "d", Oct: "b", Nov: "m", Dec: "l"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Kɔndɔŋ", Feb: "Màcɛ̂l", Mar: "Màtùmb", Apr: "Màtop", May: "M̀puyɛ", Jun: "Hìlòndɛ̀", Jul: "Njèbà", Aug: "Hìkaŋ", Sep: "Dìpɔ̀s", Oct: "Bìòôm", Nov: "Màyɛsèp", Dec: "Lìbuy li ńyèe"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "nɔy", Mon: "nja", Tue: "uum", Wed: "ŋge", Thu: "mbɔ", Fri: "kɔɔ", Sat: "jon"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "n", Mon: "n", Tue: "u", Wed: "ŋ", Thu: "m", Fri: "k", Sat: "j"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "ŋgwà nɔ̂y", Mon: "ŋgwà njaŋgumba", Tue: "ŋgwà ûm", Wed: "ŋgwà ŋgê", Thu: "ŋgwà mbɔk", Fri: "ŋgwà kɔɔ", Sat: "ŋgwà jôn"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "I\u202fbikɛ̂glà", PM: "I\u202fɓugajɔp"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "I bikɛ̂glà", PM: "I ɓugajɔp"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirhàm èmìrâ", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "Kwànza àŋgolà", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dɔ̀lâr òstralìà", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
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
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
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
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
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
				currency.RWF: cldr.Currency{DisplayName: "Frǎŋ Rùandà", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Rìal sàudì", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rùpiɛ̀ sèsɛ̂l", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Dìnâr sùdân", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Paùnd sùdân", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Paùnd hèlenà", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Lèonɛ̀", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Lèonɛ̀ (1964—2022)", Symbol: ""},
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
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Frǎŋ CFA (BCEAO)", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Rân àfrǐkàsɔ̀", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwàca sàmbià (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwàca sàmbià", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dɔ̀lâr sìmbàbwê", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Hɔp u akan",
			language.AM:  "Hɔp u amhārìk",
			language.AR:  "Hɔp u arâb",
			language.BAS: "Ɓàsàa",
			language.BE:  "Hɔp u bièlòrûs",
			language.BG:  "Hɔp u bûlgâr",
			language.BN:  "Hɔp u bɛŋgàli",
			language.CS:  "Hɔp u cɛ̂k",
			language.DE:  "Hɔp u jamân",
			language.EL:  "Hɔp u gri ᷇kyà",
			language.EN:  "Hɔp u ŋgisì",
			language.ES:  "Hɔp u panyā",
			language.FA:  "Hɔp u pɛrsìà",
			language.FR:  "Hɔp u pulàsi",
			language.HA:  "Hɔp u ɓausa",
			language.HI:  "Hɔp u hindì",
			language.HU:  "Hɔp u hɔŋgrìi",
			language.ID:  "Hɔp u indònesìà",
			language.IG:  "Hɔp u iɓò",
			language.IT:  "Hɔp u italìà",
			language.JA:  "Hɔp u yapàn",
			language.JV:  "Hɔp u yavà",
			language.KM:  "Hɔp u kmɛ̂r",
			language.KO:  "Hɔp u kɔrēà",
			language.MS:  "Hɔp u makɛ᷆",
			language.MY:  "Hɔp u birmàn",
			language.NE:  "Hɔp u nepa᷆l",
			language.NL:  "Hɔp u nlɛ̀ndi",
			language.PA:  "Hɔp u pɛnjàbi",
			language.PL:  "Hɔp u pɔlɔ̄nà",
			language.PT:  "Hɔp u pɔtɔ̄kì",
			language.RO:  "Hɔp u rùmanìà",
			language.RU:  "Hɔp u ruslànd",
			language.RW:  "Hɔp u ruāndà",
			language.SO:  "Hɔp u somàlî",
			language.SV:  "Hɔp u suɛ᷆d",
			language.TA:  "Hɔp u tamu᷆l",
			language.TH:  "Hɔp u tây",
			language.TR:  "Hɔp u tûrk",
			language.UK:  "Hɔp u ukrǎnìà",
			language.UR:  "Hɔp u urdù",
			language.VI:  "Hɔp u vyɛ̄dnàm",
			language.YO:  "Hɔp u yorūbà",
			language.ZH:  "Hɔp u kinà",
			language.ZU:  "Hɔp u zulù",
		},
		Territories: cldr.Territories{
			territory.AD: "Àŋdɔ̂r",
			territory.AE: "Àdnà i Bilɔ̀ŋ bi Arābìà",
			territory.AF: "Àfgànìstâŋ",
			territory.AG: "Àŋtigà ɓɔ Bàrbudà",
			territory.AI: "Àŋgiyà",
			territory.AL: "Àlbanìà",
			territory.AM: "Àrmenìà",
			territory.AO: "Àŋgolà",
			territory.AR: "Àrgàŋtinà",
			territory.AT: "Òstrǐk",
			territory.AU: "Òstralìà",
			territory.AW: "Àrubà",
			territory.AZ: "Àzɛ̀rbajàŋ",
			territory.BA: "Bòhnià Ɛrzègòvinà",
			territory.BB: "Bàrbadò",
			territory.BD: "Bàŋglàdɛ̂s",
			territory.BE: "Bɛlgyùm",
			territory.BF: "Bùrkìnà Fasò",
			territory.BG: "Bùlgarìà",
			territory.BH: "Bàraìn",
			territory.BI: "Bùrundì",
			territory.BJ: "Bènɛ̂ŋ",
			territory.BM: "Bɛ̀rmudà",
			territory.BN: "Brunei",
			territory.BO: "Bòlivìà",
			territory.BR: "Bràsîl",
			territory.BS: "Bàhamàs",
			territory.BT: "Bùtân",
			territory.BW: "Bòdsùanà",
			territory.BY: "Bèlarùs",
			territory.BZ: "Bèlîs",
			territory.CA: "Kànadà",
			territory.CD: "Kòŋgo ìkɛŋi",
			territory.CF: "Ŋ̀ɛm Afrīkà",
			territory.CG: "Kòŋgo",
			territory.CH: "Sùwîs",
			territory.CI: "Màŋ mi Njɔ̂k",
			territory.CK: "Bìòn bi Kook",
			territory.CL: "Kìlî",
			territory.CM: "Kàmɛ̀rûn",
			territory.CN: "Kinà",
			territory.CO: "Kɔ̀lɔmbìà",
			territory.CR: "Kòstà Rikà",
			territory.CU: "Kubà",
			territory.CV: "Kabwɛ᷆r",
			territory.CY: "Kiprò",
			territory.DE: "Jamân",
			territory.DJ: "Jìbutì",
			territory.DK: "Dànmârk",
			territory.DM: "Dòmnîk",
			territory.DO: "Dòmnikà",
			territory.DZ: "Àlgerìà",
			territory.EC: "Èkwàtorìà",
			territory.EE: "Èstonìà",
			territory.EG: "Ègîptò",
			territory.ER: "Èrìtrěà",
			territory.ES: "Pànya",
			territory.ET: "Ètìopìà",
			territory.FI: "Fìnlând",
			territory.FJ: "Fiji",
			territory.FK: "Bìòn bi Falkland",
			territory.FM: "Mìkrònesìà",
			territory.FR: "Pùlàsi / Fɛ̀lɛ̀nsi /",
			territory.GA: "Gàbɔ̂ŋ",
			territory.GB: "Àdnà i Lɔ̂ŋ",
			territory.GD: "Grènadà",
			territory.GE: "Gèɔrgìà",
			territory.GF: "Gùyanà Pùlàsi",
			territory.GH: "Ganà",
			territory.GI: "Gìlbràtâr",
			territory.GL: "Grǐnlànd",
			territory.GM: "Gàmbià",
			territory.GN: "Gìnê",
			territory.GP: "Gwàdèlûp",
			territory.GQ: "Gìne Èkwàtorìà",
			territory.GR: "Grǐkyà",
			territory.GT: "Gwàtèmalà",
			territory.GU: "Gùâm",
			territory.GW: "Gìne Bìsàô",
			territory.GY: "Gùyanà",
			territory.HN: "Ɔ̀ŋduràs",
			territory.HR: "Kròasìà",
			territory.HT: "Àitì",
			territory.HU: "Ɔ̀ŋgriì",
			territory.ID: "Indònèsià",
			territory.IE: "Ìrlând",
			territory.IL: "Isràɛ̂l",
			territory.IN: "Indìà",
			territory.IO: "Bìtèk bi Ŋgisì i Tūyɛ Īndìà",
			territory.IQ: "Ìrâk",
			territory.IR: "Ìrâŋ",
			territory.IS: "Ìslandìà",
			territory.IT: "Ìtalìà",
			territory.JM: "Jàmàikà",
			territory.JO: "Yɔ̀rdanià",
			territory.KE: "Kenìà",
			territory.KG: "Kìrgìzìstàŋ",
			territory.KH: "Kàmbodìà",
			territory.KI: "Kìrìbatì",
			territory.KM: "Kɔ̀mɔ̂r",
			territory.KN: "Nûmpubi Kîts nì Nevìs",
			territory.KP: "Kɔ̀re ì Ŋ̀ɔmbɔk",
			territory.KR: "Kɔ̀re ì Ŋ̀wɛ̀lmbɔk",
			territory.KW: "Kòwêt",
			territory.KY: "Bìòn bi Kaymàn",
			territory.KZ: "Kàzàkstâŋ",
			territory.LA: "Làôs",
			territory.LB: "Lèbanòn",
			territory.LC: "Nûmpubi Lusì",
			territory.LI: "Ligstɛntàn",
			territory.LK: "Srìlaŋkà",
			territory.LR: "Lìberìà",
			territory.LS: "Lesòtò",
			territory.LT: "Lìtùanìà",
			territory.LU: "Lùgsàmbûr",
			territory.LV: "Làdvià",
			territory.LY: "Libìà",
			territory.MA: "Màrokò",
			territory.MC: "Mònakò",
			territory.MD: "Moldavìà",
			territory.MG: "Màdàgàskâr",
			territory.MH: "Bìòn bi Marcàl",
			territory.ML: "Màli",
			territory.MM: "Myànmâr",
			territory.MN: "Mòŋgolìà",
			territory.MP: "Bìòn bi Marìanà ŋ̀ɔmbɔk",
			territory.MQ: "Màrtìnîk",
			territory.MR: "Mòrìtanìà",
			territory.MS: "Mɔ̀ŋseràt",
			territory.MT: "Maltà",
			territory.MU: "Mòrîs",
			territory.MV: "Màldîf",
			territory.MW: "Màlàwi",
			territory.MX: "Mɛ̀gsîk",
			territory.MY: "Màlɛ̀sìà",
			territory.MZ: "Mòsàmbîk",
			territory.NA: "Nàmibìà",
			territory.NC: "Kàlèdonìà Yɔ̀ndɔ",
			territory.NE: "Nìjɛ̂r",
			territory.NF: "Òn i Nɔrfɔ̂k",
			territory.NG: "Nìgerìà",
			territory.NI: "Nìkàragwà",
			territory.NL: "Ǹlɛndi",
			territory.NO: "Nɔ̀rvegìà",
			territory.NP: "Nèpâl",
			territory.NR: "Nerù",
			territory.NU: "Nìuɛ̀",
			territory.NZ: "Sìlând Yɔ̀ndɔ",
			territory.OM: "Òmân",
			territory.PA: "Pànàma",
			territory.PE: "Pèrû",
			territory.PF: "Pòlìnesìà Pùlàsi",
			territory.PG: "Gìne ì Pàpu",
			territory.PH: "Fìlìpîn",
			territory.PK: "Pàkìstân",
			territory.PL: "Pòlànd",
			territory.PM: "Nûmpubi Petrò nì Mikèlôn",
			territory.PN: "Pìdkaìrn",
			territory.PR: "Pɔ̀rtò Rikò",
			territory.PS: "Pàlɛ̀htinà Hyɔ̀ŋg nì Gazà",
			territory.PT: "Pɔ̀tɔkì",
			territory.PW: "Pàlaù",
			territory.PY: "Pàràgwê",
			territory.QA: "Kàtâr",
			territory.RE: "Rèunyɔ̂ŋ",
			territory.RO: "Rùmanìà",
			territory.RU: "Ruslànd",
			territory.RW: "Rùandà",
			territory.SA: "Sàudi Àrabìà",
			territory.SB: "Bìòn bi Salōmò",
			territory.SC: "Sèsɛ̂l",
			territory.SD: "Sùdâŋ",
			territory.SE: "Swedɛ̀n",
			territory.SG: "Sìŋgàpûr",
			territory.SH: "Nûmpubi Ɛlēnà",
			territory.SI: "Slòvanìà",
			territory.SK: "Slòvakìà",
			territory.SL: "Sièra Lèɔ̂n",
			territory.SM: "Nûmpubi Māatìn",
			territory.SN: "Sènègâl",
			territory.SO: "Sòmalìà",
			territory.SR: "Sùrinâm",
			territory.ST: "Sào Tòme ɓɔ Prɛ̀ŋcipè",
			territory.SV: "Sàlvàdɔ̂r",
			territory.SY: "Sirìà",
			territory.SZ: "Swàzìlând",
			territory.TC: "Bìòn bi Tûrks nì Kalkòs",
			territory.TD: "Câd",
			territory.TG: "Tògo",
			territory.TH: "Taylànd",
			territory.TJ: "Tàjìkìstaŋ",
			territory.TK: "Tòkèlaò",
			territory.TL: "Tìmɔ̂r lìkòl",
			territory.TM: "Tùrgmènìstân",
			territory.TN: "Tùnisìà",
			territory.TO: "Tɔŋgà",
			territory.TR: "Tùrkây",
			territory.TT: "Trìnidàd ɓɔ Tòbagò",
			territory.TV: "Tùvàlù",
			territory.TW: "Tàywân",
			territory.TZ: "Tànzànià",
			territory.UA: "Ùkrɛ̌n",
			territory.UG: "Ùgandà",
			territory.US: "Àdnà i Bilɔ̀ŋ bi Amerkà",
			territory.UY: "Ùrùgwêy",
			territory.UZ: "Ùzbèkìstân",
			territory.VA: "Vàtìkâŋ",
			territory.VC: "Nûmpubi Vɛ̂ŋsâŋ nì grènàdîn",
			territory.VE: "Vènèzùelà",
			territory.VG: "Bìòn bi kɔnji bi Ŋgisì",
			territory.VI: "Bìòn bi kɔnji bi U.S.",
			territory.VN: "Vìɛ̀dnâm",
			territory.VU: "Vànùatù",
			territory.WF: "Wàlîs nì Fùtunà",
			territory.WS: "Sàmoà",
			territory.YE: "Yèmɛ̂n",
			territory.YT: "Màyɔ̂t",
			territory.ZA: "Àfrǐkà Sɔ̀",
			territory.ZM: "Zàmbià",
			territory.ZW: "Zìmbàbwê",
		},
	}
}
