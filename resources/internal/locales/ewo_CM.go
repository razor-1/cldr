// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ewo_CM() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ewo_CM",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "ngo", Feb: "ngb", Mar: "ngl", Apr: "ngn", May: "ngt", Jun: "ngs", Jul: "ngz", Aug: "ngm", Sep: "nge", Oct: "nga", Nov: "ngad", Dec: "ngab"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "o", Feb: "b", Mar: "l", Apr: "n", May: "t", Jun: "s", Jul: "z", Aug: "m", Sep: "e", Oct: "a", Nov: "d", Dec: "b"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ngɔn osú", Feb: "ngɔn bɛ̌", Mar: "ngɔn lála", Apr: "ngɔn nyina", May: "ngɔn tána", Jun: "ngɔn saməna", Jul: "ngɔn zamgbála", Aug: "ngɔn mwom", Sep: "ngɔn ebulú", Oct: "ngɔn awóm", Nov: "ngɔn awóm ai dziá", Dec: "ngɔn awóm ai bɛ̌"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "sɔ́n", Mon: "mɔ́n", Tue: "smb", Wed: "sml", Thu: "smn", Fri: "fúl", Sat: "sér"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "s", Mon: "m", Tue: "s", Wed: "s", Thu: "s", Fri: "f", Sat: "s"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "sɔ́ndɔ", Mon: "mɔ́ndi", Tue: "sɔ́ndɔ məlú mə́bɛ̌", Wed: "sɔ́ndɔ məlú mə́lɛ́", Thu: "sɔ́ndɔ məlú mə́nyi", Fri: "fúladé", Sat: "séradé"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "kíkíríg", PM: "ngəgógəle"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirám yá Emirá Aráb Uní", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "Kwánəza yá Angolá", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dolár yá Osətəralí", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinár yá Bahərɛ́n", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Fəláŋ yá Burundí", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Púlá yá Botswána", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dolár yá Kanáda", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Fəláŋ yá Kongo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Fəláŋ yá Suís", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuán Renəminəbí yá Tsainís", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Esəkúdo yá Kápə́vɛ́rə", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Fəláŋ yá dzibutí", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinár yá Alehérí", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Lívə́lə yá Ehíbətía", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Náfəka yá Eritelé", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Bír yá Etsiópia", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "əró", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Lívə́lə Sətərəlíŋ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Tzedí yá Ganá", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasí yá Gámbía", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Síli yá Giné", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupí yá ɛ́ndía", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yɛ́n yá Hapɔ́n", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Silíŋ yá Keniá", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Fəláŋ yá Komória", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dolár yá Libéria", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lotí yá Lesotó", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinár yá Libí", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Dirám yá Maróg", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariari yá Maləgás", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugiya yá Moritaní (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugiya yá Moritaní", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupí yá Morís", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kwatsa yá Malawí", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikal yá Mozambíg", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dolár yá Namibí", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Náíra yá Nihéria", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Fəláŋ yá Ruwandá", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riál yá Arabí Saudí", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupí yá Sɛsɛ́l", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Lívələ yá Sudán", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Lívələ yá Sudán (1956–2007)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Lívələ yá Ǹfúfúb Elɛ́n", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Leóne yá Sierá-leónə", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leóne yá Sierá-leónə (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Silíŋ yá Somalí", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dóbə́ra yá Saó Tomé ai Pəlinəsípe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dóbə́ra yá Saó Tomé ai Pəlinəsípe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni yá Swazí", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinár yá Tunisí", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Silíŋ yá Tanazaní", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Silíŋ yá Ugandá (1966–1987)", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dolár yá amɛ́rəkə", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Fəláŋ CFA (BEAC)", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Fəláŋ CFA (BCEAO)", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Ránədə yá Afiríka", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwatsa yá Zambí (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwatsa yá Zambí", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dolár yá Zimbabwé", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK:  "Ǹkɔ́bɔ akán",
			language.AM:  "Ǹkɔ́bɔ amária",
			language.AR:  "Ǹkɔ́bɔ arábia",
			language.BE:  "Ǹkɔ́bɔ belarúsian",
			language.BG:  "Ǹkɔ́bɔ buləgárian",
			language.BN:  "Ǹkɔ́bɔ bɛngalí",
			language.CS:  "Ǹkɔ́bɔ tsɛ́g",
			language.DE:  "Ǹkɔ́bɔ ndzáman",
			language.EL:  "Ǹkɔ́bɔ gəlɛ́g",
			language.EN:  "Ǹkɔ́bɔ éngəlís",
			language.ES:  "ǹkɔ́bɔ kpənyá",
			language.EWO: "ewondo",
			language.FA:  "ǹkɔ́bɔ fɛ́rəsian",
			language.FR:  "Ǹkɔ́bɔ fulɛnsí",
			language.HA:  "Ǹkɔ́bɔ aúsá",
			language.HI:  "Ǹkɔ́bɔ hindí",
			language.HU:  "Ǹkɔ́bɔ ungárían",
			language.ID:  "Ǹkɔ́bɔ ɛndonésian",
			language.IG:  "Ǹkɔ́bɔ ibó",
			language.IT:  "Ǹkɔ́bɔ etáliɛn",
			language.JA:  "Ǹkɔ́bɔ hapɔ́n",
			language.JV:  "Ǹkɔ́bɔ havanís",
			language.KM:  "Ǹkɔ́bɔ kəmɛ́r",
			language.KO:  "Ǹkɔ́bɔ koréan",
			language.MS:  "Ǹkɔ́bɔ malɛ́sian",
			language.MY:  "Ǹkɔ́bɔ birəmán",
			language.NE:  "ǹkɔ́bɔ nefálian",
			language.NL:  "Ǹkɔ́bɔ nɛrəlándía",
			language.PA:  "ǹkɔ́bɔ funəhábia",
			language.PL:  "ǹkɔ́bɔ fólis",
			language.PT:  "ǹkɔ́bɔ fɔtugɛ́s",
			language.RO:  "ńkɔ́bɔ románía",
			language.RU:  "ǹkɔ́bɔ rúsian",
			language.RW:  "ǹkɔ́bɔ ruwandá",
			language.SO:  "ǹkɔ́bɔ somália",
			language.SV:  "ǹkɔ́bɔ suwɛ́d",
			language.TA:  "ǹkɔ́bɔ tamíl",
			language.TH:  "ǹkɔ́bɔ táilan",
			language.TR:  "ǹkɔ́bɔ túrəki",
			language.UK:  "ǹkɔ́bɔ ukelénia",
			language.UR:  "ǹkɔ́bɔ urudú",
			language.VI:  "ǹkɔ́bɔ hiɛdənám",
			language.YO:  "ǹkɔ́bɔ yorúba",
			language.ZH:  "Ǹkɔ́bɔ tsainís",
			language.ZU:  "ǹkɔ́bɔ zulú",
		},
		Territories: cldr.Territories{
			territory.AD: "Andór",
			territory.AE: "Bemirá yá Arábə uní",
			territory.AF: "Afəganisətán",
			territory.AG: "Antígwa ai Barəbúda",
			territory.AI: "Angíyə",
			territory.AL: "Aləbánia",
			territory.AM: "Arəménia",
			territory.AO: "Angolá",
			territory.AR: "Arəhenətína",
			territory.AS: "Bəsamóa yá Amə́rəka",
			territory.AT: "Osətəlía",
			territory.AU: "Osətəlalí",
			territory.AW: "Arúba",
			territory.AZ: "Azɛrəbaidzáŋ",
			territory.BA: "Bosəní ai ɛrəzegovín",
			territory.BB: "Barəbád",
			territory.BD: "Bangaladɛ́s",
			territory.BE: "Bɛləhíg",
			territory.BF: "Buləkiná Fasó",
			territory.BG: "Buləgarí",
			territory.BH: "Bahərɛ́n",
			territory.BI: "Burundí",
			territory.BJ: "Bəníŋ",
			territory.BM: "Bɛrəmúd",
			territory.BN: "Buluné",
			territory.BO: "Bolívia",
			territory.BR: "Bəlazíl",
			territory.BS: "Bahámas",
			territory.BT: "Butáŋ",
			territory.BW: "Botswaná",
			territory.BY: "Bəlarús",
			territory.BZ: "Bəlís",
			territory.CA: "kanadá",
			territory.CD: "ǹnam Kongó Demokəlatíg",
			territory.CF: "ǹnam Zǎŋ Afiriká",
			territory.CG: "Kongó",
			territory.CH: "Suís",
			territory.CI: "Kód Divɔ́r",
			territory.CK: "Minlán Mí kúg",
			territory.CL: "Tsilí",
			territory.CM: "Kamərún",
			territory.CN: "Tsáina",
			territory.CO: "Kolɔmbí",
			territory.CR: "Kosta Ríka",
			territory.CU: "Kubá",
			territory.CV: "Minlán Mí Káb Vɛr",
			territory.CY: "Sipəlús",
			territory.CZ: "Ǹnam Tsɛ́g",
			territory.DE: "Ndzáman",
			territory.DJ: "Dzibutí",
			territory.DK: "Danəmárəg",
			territory.DM: "Dómənika",
			territory.DO: "République dominicaine",
			territory.DZ: "Aləyéria",
			territory.EC: "Ekwatór",
			territory.EE: "Esetoní",
			territory.EG: "Ehíbətɛn",
			territory.ER: "Elitəlé",
			territory.ES: "Kpənyá",
			territory.ET: "Etiopí",
			territory.FI: "Finəlán",
			territory.FJ: "Fidzí",
			territory.FK: "Minlán Mi Fóləkəlan",
			territory.FM: "Mikoronésia",
			territory.FR: "Fulɛnsí",
			territory.GA: "Gabóŋ",
			territory.GB: "Ǹnam Engəlis",
			territory.GD: "Gələnádə",
			territory.GE: "Horə́yia",
			territory.GF: "Guyán yá Fulɛnsí",
			territory.GH: "Ganá",
			territory.GI: "Yiləbalatár",
			territory.GL: "Goelán",
			territory.GM: "Gambí",
			territory.GN: "Giné",
			territory.GP: "Guadəlúb",
			territory.GQ: "Giné Ekwató",
			territory.GR: "Gəlɛ́s",
			territory.GT: "Guatemalá",
			territory.GU: "Guám",
			territory.GW: "Giné Bisaó",
			territory.GY: "Guyán",
			territory.HN: "Ondurás",
			territory.HR: "Kəlowásia",
			territory.HT: "Aití",
			territory.HU: "Ongirí",
			territory.ID: "ɛndonésia",
			territory.IE: "Irəlándə",
			territory.IL: "Isəraɛ́l",
			territory.IN: "ɛ́ndə",
			territory.IO: "ǹnam ɛngəlís yá Máŋ mə́ ɛ́ndə",
			territory.IQ: "Irág",
			territory.IR: "Irán",
			territory.IS: "Isəlándə",
			territory.IT: "Itáliɛn",
			territory.JM: "Hamaíka",
			territory.JO: "Horədaní",
			territory.JP: "Hapɔ́n",
			territory.KE: "Keniá",
			territory.KG: "Kirigisətán",
			territory.KH: "kambodía",
			territory.KI: "Kiribatí",
			territory.KM: "Komɔ́r",
			territory.KN: "Ǹfúfúb-Kilisətóv-ai-Nevis",
			territory.KP: "Koré yá Nór",
			territory.KR: "Koré yá Súd",
			territory.KW: "Kowɛ́d",
			territory.KY: "Minlán Mí Kalimáŋ",
			territory.KZ: "Kazakətáŋ",
			territory.LA: "Laós",
			territory.LB: "Libáŋ",
			territory.LC: "Ǹfúfúb-Lúsia",
			territory.LI: "Lísə́sə́táin",
			territory.LK: "Səri Laŋká",
			territory.LR: "Libéria",
			territory.LS: "Ləsotó",
			territory.LT: "Lituaní",
			territory.LU: "Lukəzambúd",
			territory.LV: "Lətoní",
			territory.LY: "Libí",
			territory.MA: "Marɔ́g",
			territory.MC: "Mɔnakó",
			territory.MD: "Molədaví",
			territory.MG: "Madagasəkárə",
			territory.MH: "Minlán Mí Maresál",
			territory.ML: "Malí",
			territory.MM: "Mianəmár",
			territory.MN: "Mɔngɔ́lia",
			territory.MP: "Minlán Mi Marián yá Nór",
			territory.MQ: "Marətiníg",
			territory.MR: "Moritaní",
			territory.MS: "Mɔ́ntserád",
			territory.MT: "Málətə",
			territory.MU: "Morís",
			territory.MV: "Malədívə",
			territory.MW: "Malawí",
			territory.MX: "Mɛkəsíg",
			territory.MY: "Malɛ́zia",
			territory.MZ: "Mozambíg",
			territory.NA: "Namibí",
			territory.NC: "Ǹkpámɛn Kaledónia",
			territory.NE: "Nihɛ́r",
			territory.NF: "Minlán Nɔrəfɔ́ləkə",
			territory.NG: "Nihéria",
			territory.NI: "Nikarágua",
			territory.NL: "Pɛíbá",
			territory.NO: "Nɔrəvɛ́s",
			territory.NP: "Nepál",
			territory.NR: "Naurú",
			territory.NU: "Niué",
			territory.NZ: "Ǹkpámɛn Zeláŋ",
			territory.OM: "Omán",
			territory.PA: "Panamá",
			territory.PE: "Perú",
			territory.PF: "Polinesí yá Fulɛnsí",
			territory.PG: "Papwazi yá Ǹkpámɛ́n Giné",
			territory.PH: "Filipín",
			territory.PK: "Pakisətán",
			territory.PL: "fólis",
			territory.PM: "Ǹfúfúb-Píɛr-ai-Mikəlɔ́ŋ",
			territory.PN: "Pítə́kɛ́rɛnə",
			territory.PR: "Pwɛrəto Ríko",
			territory.PS: "Ǹnam Palɛsətín",
			territory.PT: "fɔrətugɛ́s",
			territory.PW: "Palau",
			territory.PY: "Paragué",
			territory.QA: "Katár",
			territory.RE: "Reuniɔ́ŋ",
			territory.RO: "Rumaní",
			territory.RU: "Rúsian",
			territory.RW: "Ruwandá",
			territory.SA: "Arabí Saudí",
			territory.SB: "Minlán Mí Solomɔ́n",
			territory.SC: "Sɛsɛ́l",
			territory.SD: "Sudáŋ",
			territory.SE: "Suwɛ́d",
			territory.SG: "Singapúr",
			territory.SH: "Ǹfúfúb-Ɛlɛ́na",
			territory.SI: "Səlovénia",
			territory.SK: "Səlovakí",
			territory.SL: "Sierá-leónə",
			territory.SM: "Ǹfúfúb Maríno",
			territory.SN: "Senegál",
			territory.SO: "Somália",
			territory.SR: "Surinám",
			territory.ST: "Saó Tomé ai Pəlinəsípe",
			territory.SV: "Saləvadór",
			territory.SY: "Sirí",
			territory.SZ: "Swazilándə",
			territory.TC: "Minlán Mí túrə́g-ai-Kaíg",
			territory.TD: "Tsád",
			territory.TG: "Togó",
			territory.TH: "Tailán",
			territory.TJ: "Tadzikisətáŋ",
			territory.TK: "Tokeló",
			territory.TL: "Timôr",
			territory.TM: "Turəkəmənisətáŋ",
			territory.TN: "Tunisí",
			territory.TO: "Tɔngá",
			territory.TR: "Turəkí",
			territory.TT: "Təlinité-ai-Tobágo",
			territory.TV: "Tuvalú",
			territory.TW: "Taiwán",
			territory.TZ: "Taŋəzaní",
			territory.UA: "Ukərɛ́n",
			territory.UG: "Ugandá",
			territory.US: "Ǹnam Amɛrəkə",
			territory.UY: "Urugué",
			territory.UZ: "Uzubekisətán",
			territory.VA: "Ǹnam Vatikán",
			territory.VC: "Ǹfúfúb-Vɛngəsáŋ-ai-Bə Gələnadín",
			territory.VE: "Venezuéla",
			territory.VG: "ńnam Minlán ɛ́ngəlís",
			territory.VI: "Minlán Mi Amɛrəkə",
			territory.VN: "Viɛdənám",
			territory.VU: "Vanuátu",
			territory.WF: "Walís-ai-Futúna",
			territory.WS: "Samoá",
			territory.YE: "Yemɛ́n",
			territory.YT: "Mayɔ́d",
			territory.ZA: "Afiríka yá Súd",
			territory.ZM: "Zambí",
			territory.ZW: "Zimbabwé",
		},
	}
}
