// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_ff_Latn_NG() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ff_Latn_NG",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "sii", Feb: "col", Mar: "mbo", Apr: "see", May: "duu", Jun: "kor", Jul: "mor", Aug: "juk", Sep: "slt", Oct: "yar", Nov: "jol", Dec: "bow"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "s", Feb: "c", Mar: "m", Apr: "s", May: "d", Jun: "k", Jul: "m", Aug: "j", Sep: "s", Oct: "y", Nov: "j", Dec: "b"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "siilo", Feb: "colte", Mar: "mbooy", Apr: "seeɗto", May: "duujal", Jun: "korse", Jul: "morso", Aug: "juko", Sep: "siilto", Oct: "yarkomaa", Nov: "jolal", Dec: "bowte"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "dew", Mon: "aaɓ", Tue: "maw", Wed: "nje", Thu: "naa", Fri: "mwd", Sat: "hbi"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "d", Mon: "a", Tue: "m", Wed: "n", Thu: "n", Fri: "m", Sat: "h"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "dewo", Mon: "aaɓnde", Tue: "mawbaare", Wed: "njeslaare", Thu: "naasaande", Fri: "mawnde", Sat: "hoore-biir"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "subaka", PM: "kikiiɗe"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "", Percent: ""},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirham Emiraati Araab Dentuɗi", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "", Symbol: "؋"},
				currency.AMD: cldr.Currency{DisplayName: "", Symbol: "֏"},
				currency.AOA: cldr.Currency{DisplayName: "Kwansaa Anngolaa", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Dolaar Ostaraalii", Symbol: "A$"},
				currency.AZN: cldr.Currency{DisplayName: "", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "", Symbol: "৳"},
				currency.BHD: cldr.Currency{DisplayName: "Dinaar Bahrayn", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Mbuuɗu Burunndi", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.BWP: cldr.Currency{DisplayName: "Pulaa Botwanaa", Symbol: "P"},
				currency.BZD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Dolaar Kandaaa", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Faraa Konngo", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Faraa Suwiis", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CNY: cldr.Currency{DisplayName: "Yuam Siin", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudoo Kap Weer", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Faraa Jibutii", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Dinaar Alaseri", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Liibar Ejipt", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Nafka Eriteree", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Biir Ecoppi", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Oroo", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Liibar Sterling", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Sedi Ganaa", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi Gammbi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Faraa Gine", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "", Symbol: "kn"},
				currency.HUF: cldr.Currency{DisplayName: "", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupii Enndo", Symbol: "₹"},
				currency.ISK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.JPY: cldr.Currency{DisplayName: "Yen Sapoo", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Siling Keñaa", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Faraa Komoor", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "", Symbol: "₩"},
				currency.KYD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Dolaar Liberiyaa", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Loti Lesotoo", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Dinaar Libi", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Diraham Maruk", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Ariyari Madagaskaar", Symbol: "Ar"},
				currency.MMK: cldr.Currency{DisplayName: "", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "", Symbol: "₮"},
				currency.MRO: cldr.Currency{DisplayName: "Ugiyya Muritani (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ugiyya Muritani", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupii Moriis", Symbol: "Rs"},
				currency.MWK: cldr.Currency{DisplayName: "Kuwacca Malaawi", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "", Symbol: "RM"},
				currency.MZM: cldr.Currency{DisplayName: "Metikal Mosammbik", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Dolaar Namibii", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nayraa Nijeriyaa", Symbol: "₦"},
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
				currency.RWF: cldr.Currency{DisplayName: "Faraa Ruwanndaa", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyaal Arabi Sawdit", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Rupii Seysel", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Liibar Sudaan", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Liibar Sent Helen", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Lewoon Seraa Liyon", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Lewoon Seraa Liyon (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Siling Soomali", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "Dobra Sawo Tome e Prensipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra Sawo Tome e Prensipe", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni Swaasilannda", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "", Symbol: "฿"},
				currency.TND: cldr.Currency{DisplayName: "Dinaar Tunisii", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Siling Tansanii", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Siling Uganndaa", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Dolaar Dowlaaji Dentuɗi", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "", Symbol: "$"},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VND: cldr.Currency{DisplayName: "", Symbol: "₫"},
				currency.XAF: cldr.Currency{DisplayName: "Mbuuɗi Seefaa BEAC", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Mbuuɗu Seefaa BCEAO", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "", Symbol: "¤"},
				currency.ZAR: cldr.Currency{DisplayName: "Rannda Afrik Bŋ Worgo", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Kuwacca Sammbi (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kuwacca Sammbi", Symbol: "ZK"},
				currency.ZWD: cldr.Currency{DisplayName: "Dolaar Simbaabuwe", Symbol: ""},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0}: {1}"},
		Languages: cldr.Languages{
			language.AK: "Akaan",
			language.AM: "Amarik",
			language.AR: "Aarabeere",
			language.BE: "Belaruuse",
			language.BG: "Bulgariire",
			language.BN: "Bengali",
			language.CS: "Cekkere",
			language.DE: "Docceere",
			language.EL: "Gerke",
			language.EN: "Engeleere",
			language.ES: "Español",
			language.FA: "Perseere",
			language.FF: "Pulaar",
			language.FR: "Farayseere",
			language.HA: "Hawsaŋkoore",
			language.HI: "Hinndi",
			language.HU: "Hongariire",
			language.ID: "Endonesiire",
			language.IG: "Igiboore",
			language.IT: "Italiyeere",
			language.JA: "Saponeere",
			language.JV: "Sawaneere",
			language.KM: "Kemeere",
			language.KO: "Koreere",
			language.MS: "Malayeere",
			language.MY: "Burmeese",
			language.NE: "Nepaaleere",
			language.NL: "Dacceere",
			language.PA: "Punjabeere",
			language.PL: "Poloneere",
			language.PT: "Purtugeere",
			language.RO: "Romaneere",
			language.RU: "Riis",
			language.RW: "Ruwaanndeere",
			language.SO: "Somalii",
			language.SV: "Sweedeere",
			language.TA: "Tamil",
			language.TH: "Taay",
			language.TR: "Turkeere",
			language.UK: "Ukereneere",
			language.UR: "Urdu",
			language.VI: "Wiyetnameere",
			language.YO: "Yorrubaa",
			language.ZH: "Sinuwaare",
			language.ZU: "Suluŋkoore",
		},
		Territories: cldr.Territories{
			territory.AD: "Anndoora",
			territory.AE: "Emiraat Araab Denntuɗe",
			territory.AF: "Afganistaan",
			territory.AG: "Antiguwaa e Barbudaa",
			territory.AI: "Anngiyaa",
			territory.AL: "Albanii",
			territory.AM: "Armenii",
			territory.AO: "Anngolaa",
			territory.AR: "Arjantiin",
			territory.AS: "Samowa Amerik",
			territory.AT: "Otiriis",
			territory.AU: "Ostaraalii",
			territory.AW: "Aruuba",
			territory.AZ: "Ajerbayjaan",
			territory.BA: "Bosnii Hersegowiin",
			territory.BB: "Barbadoos",
			territory.BD: "Banglaadees",
			territory.BE: "Beljik",
			territory.BF: "Burkibaa Faaso",
			territory.BG: "Bulgarii",
			territory.BH: "Bahreyn",
			territory.BI: "Burunndi",
			territory.BJ: "Benee",
			territory.BM: "Bermudaa",
			territory.BN: "Burnaay",
			territory.BO: "Boliwii",
			territory.BR: "Beresiil",
			territory.BS: "Bahamaas",
			territory.BT: "Butaan",
			territory.BW: "Botswaana",
			territory.BY: "Belaruus",
			territory.BZ: "Beliise",
			territory.CA: "Kanadaa",
			territory.CD: "Ndenndaandi Demokaraasiire Konngo",
			territory.CF: "Ndenndaandi Santarafrik",
			territory.CG: "Konngo",
			territory.CH: "Suwiis",
			territory.CI: "Kodduwaar",
			territory.CK: "Duuɗe Kuuk",
			territory.CL: "Cilii",
			territory.CM: "Kameruun",
			territory.CN: "Siin",
			territory.CO: "Kolombiya",
			territory.CR: "Kosta Rikaa",
			territory.CU: "Kubaa",
			territory.CV: "Duuɗe Kap Weer",
			territory.CY: "Siipar",
			territory.CZ: "Ndenndaandi Cek",
			territory.DE: "Almaañ",
			territory.DJ: "Jibutii",
			territory.DK: "Danmark",
			territory.DM: "Dominika",
			territory.DO: "Ndenndanndi Dominika",
			territory.DZ: "Alaseri",
			territory.EC: "Ekuwatoor",
			territory.EE: "Estoni",
			territory.EG: "Ejipt",
			territory.ER: "Eriteree",
			territory.ES: "Espaañ",
			territory.ET: "Ecoppi",
			territory.FI: "Fenland",
			territory.FJ: "Fijji",
			territory.FK: "Duuɗe Falkland",
			territory.FM: "Mikoronesii",
			territory.FR: "Farayse",
			territory.GA: "Gaboo",
			territory.GB: "Laamateeri Rentundi",
			territory.GD: "Garnaad",
			territory.GE: "Jeorgii",
			territory.GF: "Giyaan Farayse",
			territory.GH: "Ganaa",
			territory.GI: "Jibraltaar",
			territory.GL: "Gorwendland",
			territory.GM: "Gammbi",
			territory.GN: "Gine",
			territory.GP: "Gwaadalup",
			territory.GQ: "Ginee Ekuwaatoriyaal",
			territory.GR: "Gerees",
			territory.GT: "Gwaatemalaa",
			territory.GU: "Guwam",
			territory.GW: "Gine-Bisaawo",
			territory.GY: "Giyaan",
			territory.HN: "Onnduraas",
			territory.HR: "Korwasii",
			territory.HT: "Haytii",
			territory.HU: "Onngiri",
			territory.ID: "Enndonesii",
			territory.IE: "Irlannda",
			territory.IL: "Israa’iila",
			territory.IN: "Enndo",
			territory.IO: "Keeriindi britaani to maayo enndo",
			territory.IQ: "Iraak",
			territory.IR: "Iraan",
			territory.IS: "Islannda",
			territory.IT: "Itali",
			territory.JM: "Jamayka",
			territory.JO: "Jordani",
			territory.JP: "Sapoo",
			territory.KE: "Keñaa",
			territory.KG: "Kirgistaan",
			territory.KH: "Kambodso",
			territory.KI: "Kiribari",
			territory.KM: "Komoor",
			territory.KN: "Sent Kits e Newis",
			territory.KP: "Koree Rewo",
			territory.KR: "Koree Worgo",
			territory.KW: "Kuweyti",
			territory.KY: "Duuɗe Kaymaa",
			territory.KZ: "Kasakstaan",
			territory.LA: "Lawoos",
			territory.LB: "Libaa",
			territory.LC: "Sent Lusiyaa",
			territory.LI: "Lincenstayn",
			territory.LK: "Siri Lanka",
			territory.LR: "Liberiyaa",
			territory.LS: "Lesoto",
			territory.LT: "Lituaanii",
			territory.LU: "Liksembuur",
			territory.LV: "Letonii",
			territory.LY: "Libi",
			territory.MA: "Maruk",
			territory.MC: "Monaakoo",
			territory.MD: "Moldawii",
			territory.MG: "Madagaskaar",
			territory.MH: "Duuɗe Marsaal",
			territory.ML: "Maali",
			territory.MM: "Miyamaar",
			territory.MN: "Monngolii",
			territory.MP: "Duuɗe Mariyaana Rewo",
			territory.MQ: "Martinik",
			territory.MR: "Muritani",
			territory.MS: "Monseraat",
			territory.MT: "Malte",
			territory.MU: "Moriis",
			territory.MV: "Maldiiwe",
			territory.MW: "Malaawi",
			territory.MX: "Meksik",
			territory.MY: "Malesii",
			territory.MZ: "Mosammbik",
			territory.NA: "Namibii",
			territory.NC: "Nuwel Kaledonii",
			territory.NE: "Nijeer",
			territory.NF: "Duuɗe Norfolk",
			territory.NG: "Nijeriyaa",
			territory.NI: "Nikaraguwaa",
			territory.NL: "Nederlannda",
			territory.NO: "Norwees",
			territory.NP: "Nepaal",
			territory.NR: "Nawuru",
			territory.NU: "Niuwe",
			territory.NZ: "Nuwel Selannda",
			territory.OM: "Omaan",
			territory.PA: "Panamaa",
			territory.PE: "Peru",
			territory.PF: "Polinesii Farayse",
			territory.PG: "Papuwaa Nuwel Gine",
			territory.PH: "Filipiin",
			territory.PK: "Pakistaan",
			territory.PL: "Poloñ",
			territory.PM: "See Piyeer e Mikeloo",
			territory.PN: "Pitkern",
			territory.PR: "Porto Rikoo",
			territory.PS: "Palestiin Sisjordani e Gaasaa",
			territory.PT: "Purtugaal",
			territory.PW: "Palawu",
			territory.PY: "Paraguwaay",
			territory.QA: "Kataar",
			territory.RE: "Rewiñoo",
			territory.RO: "Rumanii",
			territory.RU: "Riisii",
			territory.RW: "Ruwanndaa",
			territory.SA: "Arabii Sawdit",
			territory.SB: "Duuɗe Solomon",
			territory.SC: "Seysel",
			territory.SD: "Sudaan",
			territory.SE: "Suweed",
			territory.SG: "Sinngapuur",
			territory.SH: "Sent Helen",
			territory.SI: "Slowenii",
			territory.SK: "Slowakii",
			territory.SL: "Seraa liyon",
			territory.SM: "See Maree",
			territory.SN: "Senegaal",
			territory.SO: "Somalii",
			territory.SR: "Surinaam",
			territory.ST: "Sawo Tome e Perensipe",
			territory.SV: "El Salwador",
			territory.SY: "Sirii",
			territory.SZ: "Swaasilannda",
			territory.TC: "Duuɗe Turke e Keikoos",
			territory.TD: "Caad",
			territory.TG: "Togoo",
			territory.TH: "Taylannda",
			territory.TJ: "Tajikistaan",
			territory.TK: "Tokelaaw",
			territory.TL: "Timoor Fuɗnaange",
			territory.TM: "Turkmenistaan",
			territory.TN: "Tunisii",
			territory.TO: "Tonngaa",
			territory.TR: "Turkii",
			territory.TT: "Tirnidaad e Tobaago",
			territory.TV: "Tuwaluu",
			territory.TW: "Taywaan",
			territory.TZ: "Tansanii",
			territory.UA: "Ukereen",
			territory.UG: "Unganndaa",
			territory.US: "Dowlaaji Dentuɗi Amerik",
			territory.UY: "Uruguwaay",
			territory.UZ: "Usbekistaan",
			territory.VA: "Dowla Waticaan",
			territory.VC: "See Weesaa e Garnadiin",
			territory.VE: "Wenesuwelaa",
			territory.VG: "duuɗe kecce britanii",
			territory.VI: "Duuɗe Kecce Amerik",
			territory.VN: "Wiyetnaam",
			territory.VU: "Wanuwaatuu",
			territory.WF: "Walis e Futuna",
			territory.WS: "Samowaa",
			territory.YE: "Yemen",
			territory.YT: "Mayoot",
			territory.ZA: "Afrik bŋ Worgo",
			territory.ZM: "Sammbi",
			territory.ZW: "Simbaabuwe",
		},
	}
}
