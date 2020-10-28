// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_uz() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "uz",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d-MMMM, y", Long: "d-MMMM, y", Medium: "d-MMM, y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss (z)", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Yan", Feb: "Fev", Mar: "Mar", Apr: "Apr", May: "May", Jun: "Iyn", Jul: "Iyl", Aug: "Avg", Sep: "Sen", Oct: "Okt", Nov: "Noy", Dec: "Dek"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "Y", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "I", Jul: "I", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Yanvar", Feb: "Fevral", Mar: "Mart", Apr: "Aprel", May: "May", Jun: "Iyun", Jul: "Iyul", Aug: "Avgust", Sep: "Sentabr", Oct: "Oktabr", Nov: "Noyabr", Dec: "Dekabr"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Yak", Mon: "Dush", Tue: "Sesh", Wed: "Chor", Thu: "Pay", Fri: "Jum", Sat: "Shan"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Y", Mon: "D", Tue: "S", Wed: "C", Thu: "P", Fri: "J", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Ya", Mon: "Du", Tue: "Se", Wed: "Ch", Thu: "Pa", Fri: "Ju", Sat: "Sh"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "yakshanba", Mon: "dushanba", Tue: "seshanba", Wed: "chorshanba", Thu: "payshanba", Fri: "juma", Sat: "shanba"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "TO", PM: "TK"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "TO", PM: "TK"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "TO", PM: "TK"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "-", Percent: "٪", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Birlashgan Arab Amirliklari dirhami", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afg‘oniston afg‘oniysi", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Albaniya leki", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Armaniston drami", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Niderlandiya antil guldeni", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angola kvanzasi", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Argentina pesosi", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Avstraliya dollari", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba florini", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Ozarbayjon manati", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Bosniya va Gertsegovina ayirboshlash markasi", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados dollari", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladesh takasi", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Bolgariya levi", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Bahrayn dinori", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundi franki", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda dollari", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Bruney dollari", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviya bolivianosi", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Braziliya reali", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Bagama dollari", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Butan ngultrumi", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Botsvana pulasi", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Belarus rubli", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Belarus rubli (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Beliz dollari", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanada dollari", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo franki", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Shveytsariya franki", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Chili pesosi", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Xitoy yuani (ofshor)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Xitoy yuani", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolumbiya pesosi", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Kosta-Rika koloni", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Kuba ayirboshlash pesosi", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Kuba pesosi", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Kabo-Verde eskudosi", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Chexiya kronasi", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Jibuti franki", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Daniya kronasi", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dominikana pesosi", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Jazoir dinori", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Misr funti", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritreya nakfasi", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Efiopiya biri", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Yevro", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Fiji dollari", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Folklend orollari funti", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Angliya funt sterlingi", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Gruziya larisi", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Gana sedisi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltar funti", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambiya dalasisi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Gvineya franki", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Gvatemala ketsali", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Gayana dollari", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Gonkong dollari", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Gonduras lempirasi", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Xorvatiya kunasi", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gaiti gurdi", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Vengriya forinti", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Indoneziya rupiyasi", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Isroil yangi shekeli", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Hindiston rupiyasi", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Iroq dinori", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Eron riyoli", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Islandiya kronasi", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Yamayka dollari", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Iordaniya dinori", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Yaponiya iyenasi", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Keniya shillingi", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Qirg‘iziston somi", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kambodja rieli", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komor orollari franki", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Shimoliy Koreya voni", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Janubiy Koreya voni", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuvayt dinori", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Kayman orollari dollari", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Qozog‘iston tengesi", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laos kipi", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Livan funti", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Shri-Lanka rupiyasi", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberiya dollari", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "Litva liti", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "Latviya lati", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "Liviya dinori", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Marokash dirhami", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Moldova leyi", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Malagasi ariarisi", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Makedoniya dinori", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Myanma kyati", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Mongoliya tugriki", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Makao patakasi", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Mavritaniya uqiyasi (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Mavritaniya uqiyasi", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "Mavritaniya rupiyasi", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Maldiv rupiyasi", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malavi kvachasi", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Meksika pesosi", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Malayziya ringgiti", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Mozambik metikali", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibiya dollari", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nigeriya nayrasi", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Nikaragua kordobasi", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Norvegiya kronasi", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepal rupiyasi", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Yangi Zelandiya dollari", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Ummon riyoli", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panama balboasi", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Peru soli", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Papua – Yangi Gvineya kinasi", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filippin pesosi", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pokiston rupiyasi", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Polsha zlotiyi", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Paragvay guaranisi", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Qatar riyoli", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Ruminiya leyi", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Serbiya dinori", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rossiya rubli", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Ruanda franki", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudiya Arabistoni riyoli", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Solomon orollari dollari", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seyshel rupiyasi", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Sudan funti", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Shvetsiya kronasi", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapur dollari", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Muqaddas Yelena oroli funti", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Syerra-Leone leonesi", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somali shillingi", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Surinam dollari", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Janubiy Sudan funti", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "San-Tome va Prinsipi dobrasi (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "San-Tome va Prinsipi dobrasi", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Suriya funti", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Svazilend lilangenisi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Tailand bati", Symbol: "THB"},
				currency.TJS: cldr.Currency{DisplayName: "Tojikiston somoniysi", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Turkmaniston manati", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunis dinori", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tonga paangasi", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Turk lirasi", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad va Tobago dollari", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Yangi Tayvan dollari", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzaniya shillingi", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukraina grivnasi", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Uganda shillingi", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "AQSH dollari", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Urugvay pesosi", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "O‘zbekiston so‘mi", Symbol: "soʻm"},
				currency.VEF: cldr.Currency{DisplayName: "Venesuela bolivari (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venesuela bolivari", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Vyetnam dongi", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu vatusi", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoa talasi", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Markaziy Afrika CFA franki", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Sharqiy Karib dollari", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "G‘arbiy Afrika CFA franki", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Fransuz Polineziyasi franki", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Noma’lum valyuta", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Yaman riyoli", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Janubiy Afrika rendi", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Zambiya kvachasi", Symbol: "ZMW"},
			},
		},
		Languages: cldr.Languages{
			language.AA:      "afar",
			language.AB:      "abxaz",
			language.ACE:     "achin",
			language.ADA:     "adangme",
			language.ADY:     "adigey",
			language.AF:      "afrikaans",
			language.AGQ:     "agem",
			language.AIN:     "aynu",
			language.AK:      "akan",
			language.ALE:     "aleut",
			language.ALT:     "janubiy oltoy",
			language.AM:      "amxar",
			language.AN:      "aragon",
			language.ANP:     "angika",
			language.AR:      "arab",
			language.AR_001:  "standart arab",
			language.ARN:     "mapuche",
			language.ARP:     "arapaxo",
			language.AS:      "assam",
			language.ASA:     "asu",
			language.AST:     "asturiy",
			language.AV:      "avar",
			language.AWA:     "avadxi",
			language.AY:      "aymara",
			language.AZ:      "ozar",
			language.BA:      "boshqird",
			language.BAN:     "bali",
			language.BAS:     "basa",
			language.BE:      "belarus",
			language.BEM:     "bemba",
			language.BEZ:     "bena",
			language.BG:      "bolgar",
			language.BGN:     "g‘arbiy baluj",
			language.BHO:     "bxojpuri",
			language.BI:      "bislama",
			language.BIN:     "bini",
			language.BLA:     "siksika",
			language.BM:      "bambara",
			language.BN:      "bengal",
			language.BO:      "tibet",
			language.BR:      "breton",
			language.BRX:     "bodo",
			language.BS:      "bosniy",
			language.BUG:     "bugi",
			language.BYN:     "blin",
			language.CA:      "katalan",
			language.CCP:     "chakma",
			language.CE:      "chechen",
			language.CEB:     "sebuan",
			language.CGG:     "chiga",
			language.CH:      "chamorro",
			language.CHK:     "chukot",
			language.CHM:     "mari",
			language.CHO:     "choktav",
			language.CHR:     "cheroki",
			language.CHY:     "cheyenn",
			language.CKB:     "sorani-kurd",
			language.CO:      "korsikan",
			language.CRS:     "kreol (Seyshel)",
			language.CS:      "chex",
			language.CU:      "slavyan (cherkov)",
			language.CV:      "chuvash",
			language.CY:      "valliy",
			language.DA:      "dan",
			language.DAK:     "dakota",
			language.DAR:     "dargva",
			language.DAV:     "taita",
			language.DE:      "nemischa",
			language.DE_AT:   "nemis (Avstriya)",
			language.DE_CH:   "yuqori nemis (Shveytsariya)",
			language.DGR:     "dogrib",
			language.DJE:     "zarma",
			language.DSB:     "quyi sorb",
			language.DUA:     "duala",
			language.DV:      "divexi",
			language.DYO:     "diola-fogni",
			language.DZ:      "dzongka",
			language.DZG:     "dazag",
			language.EBU:     "embu",
			language.EE:      "eve",
			language.EFI:     "efik",
			language.EKA:     "ekajuk",
			language.EL:      "grek",
			language.EN:      "inglizcha",
			language.EN_AU:   "ingliz (Avstraliya)",
			language.EN_CA:   "ingliz (Kanada)",
			language.EN_GB:   "ingliz (Buyuk Britaniya)",
			language.EN_US:   "ingliz (AQSH)",
			language.EO:      "esperanto",
			language.ES:      "ispancha",
			language.ES_419:  "ispan (Lotin Amerikasi)",
			language.ES_ES:   "ispan (Yevropa)",
			language.ES_MX:   "ispan (Meksika)",
			language.ET:      "estoncha",
			language.EU:      "bask",
			language.EWO:     "evondo",
			language.FA:      "fors",
			language.FA_AF:   "dari",
			language.FF:      "fula",
			language.FI:      "fincha",
			language.FIL:     "filipincha",
			language.FJ:      "fiji",
			language.FO:      "farercha",
			language.FON:     "fon",
			language.FR:      "fransuzcha",
			language.FR_CA:   "fransuz (Kanada)",
			language.FR_CH:   "fransuz (Shveytsariya)",
			language.FUR:     "friul",
			language.FY:      "g‘arbiy friz",
			language.GA:      "irland",
			language.GAA:     "ga",
			language.GAG:     "gagauz",
			language.GAN:     "gan",
			language.GD:      "shotland-gel",
			language.GEZ:     "geez",
			language.GIL:     "gilbert",
			language.GL:      "galisiy",
			language.GN:      "guarani",
			language.GOR:     "gorontalo",
			language.GSW:     "nemis (Shveytsariya)",
			language.GU:      "gujarot",
			language.GUZ:     "gusii",
			language.GV:      "men",
			language.GWI:     "gvichin",
			language.HA:      "xausa",
			language.HAW:     "gavaycha",
			language.HE:      "ivrit",
			language.HI:      "hind",
			language.HIL:     "hiligaynon",
			language.HMN:     "xmong",
			language.HR:      "xorvat",
			language.HSB:     "yuqori sorb",
			language.HT:      "gaityan",
			language.HU:      "venger",
			language.HUP:     "xupa",
			language.HY:      "arman",
			language.HZ:      "gerero",
			language.IA:      "interlingva",
			language.IBA:     "iban",
			language.IBB:     "ibibio",
			language.ID:      "indonez",
			language.IG:      "igbo",
			language.II:      "sichuan",
			language.ILO:     "iloko",
			language.INH:     "ingush",
			language.IO:      "ido",
			language.IS:      "island",
			language.IT:      "italyan",
			language.IU:      "inuktitut",
			language.JA:      "yapon",
			language.JBO:     "lojban",
			language.JGO:     "ngomba",
			language.JMC:     "machame",
			language.JV:      "yavan",
			language.KA:      "gruzincha",
			language.KAB:     "kabil",
			language.KAC:     "kachin",
			language.KAJ:     "kaji",
			language.KAM:     "kamba",
			language.KBD:     "kabardin",
			language.KCG:     "tyap",
			language.KDE:     "makonde",
			language.KEA:     "kabuverdianu",
			language.KFO:     "koro",
			language.KHA:     "kxasi",
			language.KHQ:     "koyra-chiini",
			language.KI:      "kikuyu",
			language.KJ:      "kvanyama",
			language.KK:      "qozoqcha",
			language.KKJ:     "kako",
			language.KL:      "grenland",
			language.KLN:     "kalenjin",
			language.KM:      "xmer",
			language.KMB:     "kimbundu",
			language.KN:      "kannada",
			language.KO:      "koreyscha",
			language.KOI:     "komi-permyak",
			language.KOK:     "konkan",
			language.KPE:     "kpelle",
			language.KR:      "kanuri",
			language.KRC:     "qorachoy-bolqor",
			language.KRL:     "karel",
			language.KRU:     "kurux",
			language.KS:      "kashmircha",
			language.KSB:     "shambala",
			language.KSF:     "bafiya",
			language.KSH:     "kyoln",
			language.KU:      "kurdcha",
			language.KUM:     "qo‘miq",
			language.KV:      "komi",
			language.KW:      "korn",
			language.KY:      "qirgʻizcha",
			language.LA:      "lotincha",
			language.LAD:     "ladino",
			language.LAG:     "langi",
			language.LB:      "lyuksemburgcha",
			language.LEZ:     "lezgin",
			language.LG:      "ganda",
			language.LI:      "limburg",
			language.LKT:     "lakota",
			language.LN:      "lingala",
			language.LO:      "laos",
			language.LOZ:     "lozi",
			language.LRC:     "shimoliy luri",
			language.LT:      "litva",
			language.LU:      "luba-katanga",
			language.LUA:     "luba-lulua",
			language.LUN:     "lunda",
			language.LUO:     "luo",
			language.LUS:     "lushay",
			language.LUY:     "luhya",
			language.LV:      "latishcha",
			language.MAD:     "madur",
			language.MAG:     "magahi",
			language.MAI:     "maythili",
			language.MAK:     "makasar",
			language.MAS:     "masay",
			language.MDF:     "moksha",
			language.MEN:     "mende",
			language.MER:     "meru",
			language.MFE:     "morisyen",
			language.MG:      "malagasiy",
			language.MGH:     "maxuva-mitto",
			language.MGO:     "meta",
			language.MH:      "marshall",
			language.MI:      "maori",
			language.MIC:     "mikmak",
			language.MIN:     "minangkabau",
			language.MK:      "makedon",
			language.ML:      "malayalam",
			language.MN:      "mongol",
			language.MNI:     "manipur",
			language.MOH:     "mohauk",
			language.MOS:     "mossi",
			language.MR:      "maratxi",
			language.MS:      "malay",
			language.MT:      "maltiy",
			language.MUA:     "mundang",
			language.MUL:     "bir nechta til",
			language.MUS:     "krik",
			language.MWL:     "miranda",
			language.MY:      "birman",
			language.MYV:     "erzya",
			language.MZN:     "mozandaron",
			language.NA:      "nauru",
			language.NAP:     "neapolitan",
			language.NAQ:     "nama",
			language.NB:      "norveg-bokmal",
			language.ND:      "shimoliy ndebele",
			language.NDS:     "quyi nemis",
			language.NDS_NL:  "quyi sakson",
			language.NE:      "nepal",
			language.NEW:     "nevar",
			language.NG:      "ndonga",
			language.NIA:     "nias",
			language.NIU:     "niue",
			language.NL:      "niderland",
			language.NL_BE:   "flamand",
			language.NMG:     "kvasio",
			language.NN:      "norveg-nyunorsk",
			language.NNH:     "ngiyembun",
			language.NOG:     "no‘g‘ay",
			language.NQO:     "nko",
			language.NR:      "janubiy ndebel",
			language.NSO:     "shimoliy soto",
			language.NUS:     "nuer",
			language.NV:      "navaxo",
			language.NY:      "cheva",
			language.NYN:     "nyankole",
			language.OC:      "oksitan",
			language.OM:      "oromo",
			language.OR:      "oriya",
			language.OS:      "osetin",
			language.PA:      "panjobcha",
			language.PAG:     "pangasinan",
			language.PAM:     "pampanga",
			language.PAP:     "papiyamento",
			language.PAU:     "palau",
			language.PCM:     "kreol (Nigeriya)",
			language.PL:      "polyakcha",
			language.PRG:     "pruss",
			language.PS:      "pushtu",
			language.PT:      "portugalcha",
			language.PT_BR:   "portugal (Braziliya)",
			language.PT_PT:   "portugal (Yevropa)",
			language.QU:      "kechua",
			language.QUC:     "kiche",
			language.RAP:     "rapanui",
			language.RAR:     "rarotongan",
			language.RM:      "romansh",
			language.RN:      "rundi",
			language.RO:      "rumincha",
			language.RO_MD:   "moldovan",
			language.ROF:     "rombo",
			language.ROOT:    "tub aholi tili",
			language.RU:      "ruscha",
			language.RUP:     "arumin",
			language.RW:      "kinyaruanda",
			language.RWK:     "ruanda",
			language.SA:      "sanskrit",
			language.SAD:     "sandave",
			language.SAH:     "saxa",
			language.SAQ:     "samburu",
			language.SAT:     "santali",
			language.SBA:     "ngambay",
			language.SBP:     "sangu",
			language.SC:      "sardin",
			language.SCN:     "sitsiliya",
			language.SCO:     "shotland",
			language.SD:      "sindhi",
			language.SDH:     "janubiy kurd",
			language.SE:      "shimoliy saam",
			language.SEH:     "sena",
			language.SES:     "koyraboro-senni",
			language.SG:      "sango",
			language.SHI:     "tashelxit",
			language.SHN:     "shan",
			language.SI:      "singal",
			language.SK:      "slovakcha",
			language.SL:      "slovencha",
			language.SM:      "samoa",
			language.SMA:     "janubiy saam",
			language.SMJ:     "lule-saam",
			language.SMN:     "inari-saam",
			language.SMS:     "skolt-saam",
			language.SN:      "shona",
			language.SNK:     "soninke",
			language.SO:      "somalicha",
			language.SQ:      "alban",
			language.SR:      "serbcha",
			language.SRN:     "sranan-tongo",
			language.SS:      "svati",
			language.SSY:     "saho",
			language.ST:      "janubiy soto",
			language.SU:      "sundan",
			language.SUK:     "sukuma",
			language.SV:      "shved",
			language.SW:      "suaxili",
			language.SW_CD:   "suaxili (Kongo)",
			language.SWB:     "qamar",
			language.SYR:     "suriyacha",
			language.TA:      "tamil",
			language.TE:      "telugu",
			language.TEM:     "timne",
			language.TEO:     "teso",
			language.TET:     "tetum",
			language.TG:      "tojik",
			language.TH:      "tay",
			language.TI:      "tigrinya",
			language.TIG:     "tigre",
			language.TK:      "turkman",
			language.TLH:     "klingon",
			language.TN:      "tsvana",
			language.TO:      "tongan",
			language.TPI:     "tok-piksin",
			language.TR:      "turk",
			language.TRV:     "taroko",
			language.TS:      "tsonga",
			language.TT:      "tatar",
			language.TUM:     "tumbuka",
			language.TVL:     "tuvalu",
			language.TWQ:     "tasavak",
			language.TY:      "taiti",
			language.TYV:     "tuva",
			language.TZM:     "markaziy atlas tamazigxt",
			language.UDM:     "udmurt",
			language.UG:      "uyg‘ur",
			language.UK:      "ukrain",
			language.UMB:     "umbundu",
			language.UND:     "noma’lum til",
			language.UR:      "urdu",
			language.UZ:      "o‘zbek",
			language.VAI:     "vai",
			language.VE:      "venda",
			language.VI:      "vyetnam",
			language.VO:      "volapyuk",
			language.VUN:     "vunjo",
			language.WA:      "vallon",
			language.WAE:     "valis",
			language.WAL:     "volamo",
			language.WAR:     "varay",
			language.WBP:     "valbiri",
			language.WO:      "volof",
			language.XAL:     "qalmoq",
			language.XH:      "kxosa",
			language.XOG:     "soga",
			language.YAV:     "yangben",
			language.YBB:     "yemba",
			language.YI:      "idish",
			language.YO:      "yoruba",
			language.YUE:     "xitoy, kanton",
			language.ZGH:     "tamazigxt",
			language.ZH:      "xitoy, mandarin",
			language.ZH_HANS: "xitoy (soddalashtirilgan mandarin)",
			language.ZH_HANT: "xitoy (an’anaviy mandarin)",
			language.ZU:      "zulu",
			language.ZUN:     "zuni",
			language.ZXX:     "til tarkibi yo‘q",
			language.ZZA:     "zaza",
		},
		Territories: cldr.Territories{
			territory.V_001: "Dunyo",
			territory.V_002: "Afrika",
			territory.V_003: "Shimoliy Amerika",
			territory.V_005: "Janubiy Amerika",
			territory.V_009: "Okeaniya",
			territory.V_011: "G‘arbiy Afrika",
			territory.V_013: "Markaziy Amerika",
			territory.V_014: "Sharqiy Afrika",
			territory.V_015: "Shimoliy Afrika",
			territory.V_017: "Markaziy Afrika",
			territory.V_018: "Janubiy Afrika",
			territory.V_019: "Amerika",
			territory.V_021: "Shimoliy Amerika – AQSH va Kanada",
			territory.V_029: "Karib havzasi",
			territory.V_030: "Sharqiy Osiyo",
			territory.V_034: "Janubiy Osiyo",
			territory.V_035: "Janubi-sharqiy Osiyo",
			territory.V_039: "Janubiy Yevropa",
			territory.V_053: "Avstralaziya",
			territory.V_054: "Melaneziya",
			territory.V_057: "Mikroneziya mintaqasi",
			territory.V_061: "Polineziya",
			territory.V_142: "Osiyo",
			territory.V_143: "Markaziy Osiyo",
			territory.V_145: "G‘arbiy Osiyo",
			territory.V_150: "Yevropa",
			territory.V_151: "Sharqiy Yevropa",
			territory.V_154: "Shimoliy Yevropa",
			territory.V_155: "G‘arbiy Yevropa",
			territory.V_202: "Sahro janubidagi Afrika",
			territory.V_419: "Lotin Amerikasi",
			territory.AC:    "Me’roj oroli",
			territory.AD:    "Andorra",
			territory.AE:    "Birlashgan Arab Amirliklari",
			territory.AF:    "Afgʻoniston",
			territory.AG:    "Antigua va Barbuda",
			territory.AI:    "Angilya",
			territory.AL:    "Albaniya",
			territory.AM:    "Armaniston",
			territory.AO:    "Angola",
			territory.AQ:    "Antarktida",
			territory.AR:    "Argentina",
			territory.AS:    "Amerika Samoasi",
			territory.AT:    "Avstriya",
			territory.AU:    "Avstraliya",
			territory.AW:    "Aruba",
			territory.AX:    "Aland orollari",
			territory.AZ:    "Ozarbayjon",
			territory.BA:    "Bosniya va Gertsegovina",
			territory.BB:    "Barbados",
			territory.BD:    "Bangladesh",
			territory.BE:    "Belgiya",
			territory.BF:    "Burkina-Faso",
			territory.BG:    "Bolgariya",
			territory.BH:    "Bahrayn",
			territory.BI:    "Burundi",
			territory.BJ:    "Benin",
			territory.BL:    "Sen-Bartelemi",
			territory.BM:    "Bermuda orollari",
			territory.BN:    "Bruney",
			territory.BO:    "Boliviya",
			territory.BQ:    "Boneyr, Sint-Estatius va Saba",
			territory.BR:    "Braziliya",
			territory.BS:    "Bagama orollari",
			territory.BT:    "Butan",
			territory.BV:    "Buve oroli",
			territory.BW:    "Botsvana",
			territory.BY:    "Belarus",
			territory.BZ:    "Beliz",
			territory.CA:    "Kanada",
			territory.CC:    "Kokos (Kiling) orollari",
			territory.CD:    "Kongo (KDR)",
			territory.CF:    "Markaziy Afrika Respublikasi",
			territory.CG:    "Kongo (Respublika)",
			territory.CH:    "Shveytsariya",
			territory.CI:    "Fil suyagi qirg‘og‘i",
			territory.CK:    "Kuk orollari",
			territory.CL:    "Chili",
			territory.CM:    "Kamerun",
			territory.CN:    "Xitoy",
			territory.CO:    "Kolumbiya",
			territory.CP:    "Klipperton oroli",
			territory.CR:    "Kosta-Rika",
			territory.CU:    "Kuba",
			territory.CV:    "Kabo-Verde",
			territory.CW:    "Kyurasao",
			territory.CX:    "Rojdestvo oroli",
			territory.CY:    "Kipr",
			territory.CZ:    "Chexiya Respublikasi",
			territory.DE:    "Germaniya",
			territory.DG:    "Diyego-Garsiya",
			territory.DJ:    "Jibuti",
			territory.DK:    "Daniya",
			territory.DM:    "Dominika",
			territory.DO:    "Dominikan Respublikasi",
			territory.DZ:    "Jazoir",
			territory.EA:    "Seuta va Melilya",
			territory.EC:    "Ekvador",
			territory.EE:    "Estoniya",
			territory.EG:    "Misr",
			territory.EH:    "G‘arbiy Sahroi Kabir",
			territory.ER:    "Eritreya",
			territory.ES:    "Ispaniya",
			territory.ET:    "Efiopiya",
			territory.EU:    "Yevropa Ittifoqi",
			territory.EZ:    "yevrozona",
			territory.FI:    "Finlandiya",
			territory.FJ:    "Fiji",
			territory.FK:    "Folklend (Malvin) orollari",
			territory.FM:    "Mikroneziya",
			territory.FO:    "Farer orollari",
			territory.FR:    "Fransiya",
			territory.GA:    "Gabon",
			territory.GB:    "Britaniya",
			territory.GD:    "Grenada",
			territory.GE:    "Gruziya",
			territory.GF:    "Fransuz Gvianasi",
			territory.GG:    "Gernsi",
			territory.GH:    "Gana",
			territory.GI:    "Gibraltar",
			territory.GL:    "Grenlandiya",
			territory.GM:    "Gambiya",
			territory.GN:    "Gvineya",
			territory.GP:    "Gvadelupe",
			territory.GQ:    "Ekvatorial Gvineya",
			territory.GR:    "Gretsiya",
			territory.GS:    "Janubiy Georgiya va Janubiy Sendvich orollari",
			territory.GT:    "Gvatemala",
			territory.GU:    "Guam",
			territory.GW:    "Gvineya-Bisau",
			territory.GY:    "Gayana",
			territory.HK:    "Gonkong",
			territory.HM:    "Xerd va Makdonald orollari",
			territory.HN:    "Gonduras",
			territory.HR:    "Xorvatiya",
			territory.HT:    "Gaiti",
			territory.HU:    "Vengriya",
			territory.IC:    "Kanar orollari",
			territory.ID:    "Indoneziya",
			territory.IE:    "Irlandiya",
			territory.IL:    "Isroil",
			territory.IM:    "Men oroli",
			territory.IN:    "Hindiston",
			territory.IO:    "Britaniyaning Hind okeanidagi hududi",
			territory.IQ:    "Iroq",
			territory.IR:    "Eron",
			territory.IS:    "Islandiya",
			territory.IT:    "Italiya",
			territory.JE:    "Jersi",
			territory.JM:    "Yamayka",
			territory.JO:    "Iordaniya",
			territory.JP:    "Yaponiya",
			territory.KE:    "Keniya",
			territory.KG:    "Qirgʻiziston",
			territory.KH:    "Kambodja",
			territory.KI:    "Kiribati",
			territory.KM:    "Komor orollari",
			territory.KN:    "Sent-Kits va Nevis",
			territory.KP:    "Shimoliy Koreya",
			territory.KR:    "Janubiy Koreya",
			territory.KW:    "Quvayt",
			territory.KY:    "Kayman orollari",
			territory.KZ:    "Qozogʻiston",
			territory.LA:    "Laos",
			territory.LB:    "Livan",
			territory.LC:    "Sent-Lyusiya",
			territory.LI:    "Lixtenshteyn",
			territory.LK:    "Shri-Lanka",
			territory.LR:    "Liberiya",
			territory.LS:    "Lesoto",
			territory.LT:    "Litva",
			territory.LU:    "Lyuksemburg",
			territory.LV:    "Latviya",
			territory.LY:    "Liviya",
			territory.MA:    "Marokash",
			territory.MC:    "Monako",
			territory.MD:    "Moldova",
			territory.ME:    "Chernogoriya",
			territory.MF:    "Sent-Martin",
			territory.MG:    "Madagaskar",
			territory.MH:    "Marshall orollari",
			territory.MK:    "Shimoliy Makedoniya",
			territory.ML:    "Mali",
			territory.MM:    "Myanma (Birma)",
			territory.MN:    "Mongoliya",
			territory.MO:    "Makao",
			territory.MP:    "Shimoliy Mariana orollari",
			territory.MQ:    "Martinika",
			territory.MR:    "Mavritaniya",
			territory.MS:    "Montserrat",
			territory.MT:    "Malta",
			territory.MU:    "Mavrikiy",
			territory.MV:    "Maldiv orollari",
			territory.MW:    "Malavi",
			territory.MX:    "Meksika",
			territory.MY:    "Malayziya",
			territory.MZ:    "Mozambik",
			territory.NA:    "Namibiya",
			territory.NC:    "Yangi Kaledoniya",
			territory.NE:    "Niger",
			territory.NF:    "Norfolk oroli",
			territory.NG:    "Nigeriya",
			territory.NI:    "Nikaragua",
			territory.NL:    "Niderlandiya",
			territory.NO:    "Norvegiya",
			territory.NP:    "Nepal",
			territory.NR:    "Nauru",
			territory.NU:    "Niue",
			territory.NZ:    "Yangi Zelandiya",
			territory.OM:    "Ummon",
			territory.PA:    "Panama",
			territory.PE:    "Peru",
			territory.PF:    "Fransuz Polineziyasi",
			territory.PG:    "Papua – Yangi Gvineya",
			territory.PH:    "Filippin",
			territory.PK:    "Pokiston",
			territory.PL:    "Polsha",
			territory.PM:    "Sen-Pyer va Mikelon",
			territory.PN:    "Pitkern orollari",
			territory.PR:    "Puerto-Riko",
			territory.PS:    "Falastin",
			territory.PT:    "Portugaliya",
			territory.PW:    "Palau",
			territory.PY:    "Paragvay",
			territory.QA:    "Qatar",
			territory.QO:    "Tashqi Okeaniya",
			territory.RE:    "Reyunion",
			territory.RO:    "Ruminiya",
			territory.RS:    "Serbiya",
			territory.RU:    "Rossiya",
			territory.RW:    "Ruanda",
			territory.SA:    "Saudiya Arabistoni",
			territory.SB:    "Solomon orollari",
			territory.SC:    "Seyshel orollari",
			territory.SD:    "Sudan",
			territory.SE:    "Shvetsiya",
			territory.SG:    "Singapur",
			territory.SH:    "Muqaddas Yelena oroli",
			territory.SI:    "Sloveniya",
			territory.SJ:    "Shpitsbergen va Yan-Mayen",
			territory.SK:    "Slovakiya",
			territory.SL:    "Syerra-Leone",
			territory.SM:    "San-Marino",
			territory.SN:    "Senegal",
			territory.SO:    "Somali",
			territory.SR:    "Surinam",
			territory.SS:    "Janubiy Sudan",
			territory.ST:    "San-Tome va Prinsipi",
			territory.SV:    "Salvador",
			territory.SX:    "Sint-Marten",
			territory.SY:    "Suriya",
			territory.SZ:    "Svazilend",
			territory.TA:    "Tristan-da-Kunya",
			territory.TC:    "Turks va Kaykos orollari",
			territory.TD:    "Chad",
			territory.TF:    "Fransuz Janubiy hududlari",
			territory.TG:    "Togo",
			territory.TH:    "Tailand",
			territory.TJ:    "Tojikiston",
			territory.TK:    "Tokelau",
			territory.TL:    "Sharqiy Timor",
			territory.TM:    "Turkmaniston",
			territory.TN:    "Tunis",
			territory.TO:    "Tonga",
			territory.TR:    "Turkiya",
			territory.TT:    "Trinidad va Tobago",
			territory.TV:    "Tuvalu",
			territory.TW:    "Tayvan",
			territory.TZ:    "Tanzaniya",
			territory.UA:    "Ukraina",
			territory.UG:    "Uganda",
			territory.UM:    "AQSH yondosh orollari",
			territory.UN:    "BMT",
			territory.US:    "AQSH",
			territory.UY:    "Urugvay",
			territory.UZ:    "Oʻzbekiston",
			territory.VA:    "Vatikan",
			territory.VC:    "Sent-Vinsent va Grenadin",
			territory.VE:    "Venesuela",
			territory.VG:    "Britaniya Virgin orollari",
			territory.VI:    "AQSH Virgin orollari",
			territory.VN:    "Vyetnam",
			territory.VU:    "Vanuatu",
			territory.WF:    "Uollis va Futuna",
			territory.WS:    "Samoa",
			territory.XA:    "Qalbaki urg‘u",
			territory.XB:    "Qalbaki Bidi",
			territory.XK:    "Kosovo",
			territory.YE:    "Yaman",
			territory.YT:    "Mayotta",
			territory.ZA:    "Janubiy Afrika Respublikasi",
			territory.ZM:    "Zambiya",
			territory.ZW:    "Zimbabve",
			territory.ZZ:    "Noma’lum mintaqa",
		},
	}
}
