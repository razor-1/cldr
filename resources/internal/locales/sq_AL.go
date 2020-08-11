// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_sq_AL() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "sq_AL",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d.M.yy"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a, zzzz", Long: "h:mm:ss a, z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'në' {0}", Long: "{1} 'në' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan", Feb: "shk", Mar: "mar", Apr: "pri", May: "maj", Jun: "qer", Jul: "korr", Aug: "gush", Sep: "sht", Oct: "tet", Nov: "nën", Dec: "dhj"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "j", Feb: "sh", Mar: "m", Apr: "p", May: "m", Jun: "q", Jul: "k", Aug: "g", Sep: "sh", Oct: "t", Nov: "n", Dec: "dh"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "janar", Feb: "shkurt", Mar: "mars", Apr: "prill", May: "maj", Jun: "qershor", Jul: "korrik", Aug: "gusht", Sep: "shtator", Oct: "tetor", Nov: "nëntor", Dec: "dhjetor"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "die", Mon: "hën", Tue: "mar", Wed: "mër", Thu: "enj", Fri: "pre", Sat: "sht"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "d", Mon: "h", Tue: "m", Wed: "m", Thu: "e", Fri: "p", Sat: "sh"}, Short: cldr.CalendarDayFormatNameValue{Sun: "die", Mon: "hën", Tue: "mar", Wed: "mër", Thu: "enj", Fri: "pre", Sat: "sht"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "e diel", Mon: "e hënë", Tue: "e martë", Wed: "e mërkurë", Thu: "e enjte", Fri: "e premte", Sat: "e shtunë"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "p.d.", PM: "m.d."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "p.d.", PM: "m.d."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "paradite", PM: "pasdite"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤;(#,##0.00\u00a0¤)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Dirhami i Emirateve të Bashkuara Arabe", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "Afgani afgan", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Leku shqiptar", Symbol: "Lekë"},
				currency.AMD: cldr.Currency{DisplayName: "Dramia armene", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Gilderi antilian holandez", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kuanza e Angolës", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "Pesoja argjentinase", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "Dollari australian", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Florini aruban", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "Manata azerbajxhanase", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "Marka e Bosnjë-Hercegovinës [e shkëmbyeshme]", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "Dollari barbadian", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Taka e Bangladeshit", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "Leva bullgare", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "Dinari i Bahreinit", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Franga burundiane", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Dollari i Bermudeve", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Dollari i Bruneit", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviani i Bolivisë", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "Reali brazilian", Symbol: "BRL"},
				currency.BSD: cldr.Currency{DisplayName: "Dollari i Bahamasit", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrumi butanez", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "Pula botsuane", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "Rubla bjelloruse", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Rubla bjelloruse (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Dollari i Ishujve Belize", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Dollari kanadez", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Franga kongole", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "Franga zvicerane", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "Pesoja kiliane", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Juani kinez (për treg të jashtëm)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Juani kinez", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Pesoja kolumbiane", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "Koloni kostarikan", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "Pesoja kubaneze e shkëmbyeshme", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Pesoja kubaneze", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Eskudoja e Kepit të Gjelbër", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "Koruna e Çekisë", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "Franga xhibutiane", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Korona daneze", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Pesoja dominikane", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Dinari algjerian", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "Sterlina egjiptiane", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa e Eritresë", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Bira etiopiane", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euroja", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Dollari i Fixhit", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Stërlina e Ishujve Falkland", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "Sterlina britanike", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Laria gjeorgjiane", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "Sejda ganeze", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Sterlina e Gjibraltarit", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi gambian", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Franga guinease", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "Kuecali i Guatemalës", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "Dollari guajanez", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Dollari i Hong-Kongut", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira hondurase", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "Kuna kroate", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gurdi haitian", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Forinta hungareze", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupia indoneziane", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "Shekeli izrealit", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupia indiane", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinari irakian", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Riali iranian", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Korona islandeze", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "Dollari xhamajkan", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Dinari jordanez", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Jeni japonez", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilinga keniane", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Soma kirgize", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Riali kamboxhian", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Franga komore", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Uoni koreano-verior", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "Uoni koreano-jugor", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinari kuvajtian", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Dollari i Ishujve Kajman", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Tenga kazake", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Kipa e Laosit", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Sterlina libaneze", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Rupia e Sri-Lankës", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Dollari liberian", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "Lita lituaneze", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "Lata letoneze", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "Dinari libian", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Dirhami maroken", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "Leuja moldave", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Arieri malagez", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "Denari maqedonas", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "Kiata e Mianmarit", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrika mongole", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Pataka e Makaos", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ugija mauritane (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Ugija mauritane", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupia mauritiane", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiu i Maldivit", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kuaça malaviane", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Pesoja meksikane", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Ringiti malajzian", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "Metikali i Mozambikut", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Dollari i Namibisë", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naira nigeriane", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "Kordoba nikaraguane", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "Korona norvegjeze", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Rupia nepaleze", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Dollari i Zelandës së Re", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Riali i Omanit", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa panameze", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "Sola peruane", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "Kina e Guinesë së Re-Papua", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Pesoja filipinase", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Rupia pakistaneze", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Zllota polake", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "Guarani paraguaian", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Riali i Katarit", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "Leuja rumune", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Dinari serb", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rubla ruse", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Franga ruandeze", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Riali saudit", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Dollari i Ishujve Solomonë", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupia e Ishujve Sishelë", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "Sterlina sudaneze", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "Korona suedeze", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Dollari i Singaporit", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Sterlina e Ishullit të Shën-Helenës", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "Leoni i Sierra-Leones", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Shilinga somaleze", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Dollari surinamez", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "Sterlina sudanezo-jugore", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Dobra e Sao-Tomes dhe Prinsipes (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Dobra e Sao-Tomes dhe Prinsipes", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "Sterlina siriane", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni i Suazilandës", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Bata tajlandeze", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Somona taxhike", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "Manata turkmene", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinari tunizian", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Panga tongane", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "Lira turke", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Dollari i Trinidadit dhe Tobagos", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Dollari tajvanez", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilinga e Tanzanisë", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Rivnia ukrainase", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "Shilinga ugandeze", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Dollari amerikan", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Pesoja uruguaiane", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Soma uzbeke", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "Bolivari venezuelian (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Bolivari venezuelas", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Donga vietnameze", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Vatuja e Vanuatusë", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Tala samoane", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Franga kamerunase", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Dollari i Karaibeve Lindore", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "Franga e Bregut të Fildishtë", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "Franga franceze e Polinezisë", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "Valutë e panjohur", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Riali i Jemenit", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "Randi afrikano-jugor", Symbol: "ZAR"},
				currency.ZMW: cldr.Currency{DisplayName: "Kuaça e Zambikut", Symbol: "ZMW"},
			},
		},
	}
}
