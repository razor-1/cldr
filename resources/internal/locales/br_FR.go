// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_br_FR() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "br_FR",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'da' {0}", Long: "{1} 'da' {0}", Medium: "{1}, {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Gen.", Feb: "Cʼhwe.", Mar: "Meur.", Apr: "Ebr.", May: "Mae", Jun: "Mezh.", Jul: "Goue.", Aug: "Eost", Sep: "Gwen.", Oct: "Here", Nov: "Du", Dec: "Kzu."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "01", Feb: "02", Mar: "03", Apr: "04", May: "05", Jun: "06", Jul: "07", Aug: "08", Sep: "09", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Genver", Feb: "Cʼhwevrer", Mar: "Meurzh", Apr: "Ebrel", May: "Mae", Jun: "Mezheven", Jul: "Gouere", Aug: "Eost", Sep: "Gwengolo", Oct: "Here", Nov: "Du", Dec: "Kerzu"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Lun", Tue: "Meu.", Wed: "Mer.", Thu: "Yaou", Fri: "Gwe.", Sat: "Sad."}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "Su", Mon: "L", Tue: "Mz", Wed: "Mc", Thu: "Y", Fri: "G", Sat: "Sa"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Lun", Tue: "Meu.", Wed: "Mer.", Thu: "Yaou", Fri: "Gwe.", Sat: "Sad."}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sul", Mon: "Lun", Tue: "Meurzh", Wed: "Mercʼher", Thu: "Yaou", Fri: "Gwener", Sat: "Sadorn"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "A.M.", PM: "G.M."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "A.M.", PM: "G.M."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "A.M.", PM: "G.M."}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "٬", Negative: "\u061c-", Percent: "٪\u061c", PerMille: "؉"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "peseta Andorra", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "dirham EAU", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "afghani Afghanistan (1927–2002)", Symbol: "AFA"},
				currency.AFN: cldr.Currency{DisplayName: "afghani Afghanistan", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "lek Albania (1946–1965)", Symbol: "ALK"},
				currency.ALL: cldr.Currency{DisplayName: "lek Albania", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "dram Armenia", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "florin Antilhez nederlandat", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "kwanza Angola", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "kwanza Angola (1977–1991)", Symbol: "AOK"},
				currency.AON: cldr.Currency{DisplayName: "kwanza nevez Angola (1990–2000)", Symbol: "AON"},
				currency.AOR: cldr.Currency{DisplayName: "", Symbol: "AOR"},
				currency.ARA: cldr.Currency{DisplayName: "", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "peso Arcʼhantina (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "peso Arcʼhantina (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "peso Arcʼhantina", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "schilling Aostria", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "dollar Aostralia", Symbol: "$A"},
				currency.AWG: cldr.Currency{DisplayName: "florin Aruba", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "manat Azerbaidjan (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "manat Azerbaidjan", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "dinar Bosnia ha Herzegovina (1992–1994)", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "mark kemmadus Bosnia ha Herzegovina", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "dinar nevez Bosnia ha Herzegovina (1994–1997)", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "dollar Barbados", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "taka Bangladesh", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "lur Belgia (kemmadus)", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "lur Belgia", Symbol: "BEF"},
				currency.BGL: cldr.Currency{DisplayName: "", Symbol: "BGL"},
				currency.BGM: cldr.Currency{DisplayName: "lev sokialour Bulgaria", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "lev Bulgaria", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "lev Bulgaria (1879–1952)", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "dinar Bahrein", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "lur Burundi", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "dollar Bermuda", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "dollar Brunei", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano Bolivia", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "boliviano Bolivia (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "peso Bolivia", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "real Brazil", Symbol: "BRL"},
				currency.BRN: cldr.Currency{DisplayName: "", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "dollar Bahamas", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ngultrum Bhoutan", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "kyat Birmania", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "pula Botswana", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "roubl nevez Belarus (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "roubl Belarus", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "roubl Belarus (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "dollar Belize", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "dollar Kanada", Symbol: "$CA"},
				currency.CDF: cldr.Currency{DisplayName: "lur Kongo", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "euro WIR", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "lur Suis", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "lur WIR", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "escudo Chile", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "unanenn jediñ Chile", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "peso Chile", Symbol: "CLP"},
				currency.CNX: cldr.Currency{DisplayName: "dollar Bank poblel Sina", Symbol: "CNX"},
				currency.CNY: cldr.Currency{DisplayName: "yuan Sina", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "peso Kolombia", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "colón Costa Rica", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "dinar Serbia (2002–2006)", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "", Symbol: "CSK"},
				currency.CUC: cldr.Currency{DisplayName: "peso kemmadus Kuba", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "peso Kuba", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "escudo Kab Glas", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "lur Kiprenez", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "kurunenn Tchek", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "mark Alamagn ar Reter", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "mark Alamagn", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "lur Djibouti", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "kurunenn Danmark", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "peso Dominikan", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "dinar Aljeria", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "kurunenn Estonia", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "lur Egipt", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa Eritrea", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "peseta gemmadus Spagn", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "peseta Spagn", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "birr Etiopia", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "mark Finland", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "dollar Fidji", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "lur Inizi Falkland", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "lur gall", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "lur Breizh-Veur", Symbol: "£ RU"},
				currency.GEK: cldr.Currency{DisplayName: "", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "lari Jorjia", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "", Symbol: "GHC"},
				currency.GHS: cldr.Currency{DisplayName: "cedi Ghana", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "lur Jibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi Gambia", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "lur Ginea", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "syli Ginea", Symbol: "GNS"},
				currency.GQE: cldr.Currency{DisplayName: "ekwele Ginea ar Cʼheheder", Symbol: "GQE"},
				currency.GRD: cldr.Currency{DisplayName: "drakm Gres", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "quetzal Guatemala", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "", Symbol: "GWE"},
				currency.GWP: cldr.Currency{DisplayName: "peso Ginea-Bissau", Symbol: "GWP"},
				currency.GYD: cldr.Currency{DisplayName: "dollar Guyana", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "dollar Hong Kong", Symbol: "$ HK"},
				currency.HNL: cldr.Currency{DisplayName: "lempira Honduras", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "dinar Kroatia", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "kuna Kroatia", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "gourde Haiti", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "forint Hungaria", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "roupi Indonezia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "lur Iwerzhon", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "lur Israel", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "shekel Israel (1980–1985)", Symbol: "ILR"},
				currency.ILS: cldr.Currency{DisplayName: "shekel nevez Israel", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "roupi India", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "dinar Iraq", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "rial Iran", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "kurunenn Island (1918–1981)", Symbol: "ISJ"},
				currency.ISK: cldr.Currency{DisplayName: "kurunenn Island", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "lur Italia", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "dollar Jamaika", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "dinar Jordania", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "yen Japan", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "shilling Kenya", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "som Kyrgyzstan", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "riel Kambodja", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "lur Komorez", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "won Korea an Norzh", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "hwan Korea ar Su (1953–1962)", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "won Korea ar Su (1945–1953)", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "won Korea ar Su", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "dinar Koweit", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "dollar Inizi Cayman", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "tenge Kazakstan", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "kip Laos", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "lur Liban", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "roupi Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "dollar Liberia", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "loti Lesotho", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "litas Lituania", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "talonas Lituania", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "lur kemmadus Luksembourg", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "lur Luksembourg", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "lats Latvia", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "roubl Latvia", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "dinar Libia", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "dirham Maroko", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "lur Maroko", Symbol: "MAF"},
				currency.MCF: cldr.Currency{DisplayName: "lur Monaco", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "leu Moldova", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ariary Madagaskar", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "lur Madagaskar", Symbol: "MGF"},
				currency.MKD: cldr.Currency{DisplayName: "denar Makedonia", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "denar Makedonia (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "lur Mali", Symbol: "MLF"},
				currency.MMK: cldr.Currency{DisplayName: "kyat Myanmar", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "tugrik Mongolia", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "pataca Macau", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "ouguiya Maouritania (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "ouguiya Maouritania", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "lira Malta", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "lur Malta", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "roupi Moris", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "roupi Maldivez", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "rufiyaa Maldivez", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha Malawi", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "peso Mecʼhiko", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "peso arcʼhant Mecʼhiko (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "ringgit Malaysia", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "escudo Mozambik", Symbol: "MZE"},
				currency.MZM: cldr.Currency{DisplayName: "metical Mozambik (1980–2006)", Symbol: "MZM"},
				currency.MZN: cldr.Currency{DisplayName: "metical Mozambik", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "dollar Namibia", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "naira Nigeria", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "cordoba Nicaragua (1988–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "cordoba Nicaragua", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "florin an Izelvroioù", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "kurunenn Norvegia", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "roupi Nepal", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "dollar Zeland-Nevez", Symbol: "$ ZN"},
				currency.OMR: cldr.Currency{DisplayName: "rial Oman", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "balboa Panamá", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "sol Perou", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "sol Perou (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "kina Papoua Ginea-Nevez", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "peso Filipinez", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "roupi Pakistan", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "zloty Polonia", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "zloty Polonia (1950–1995)", Symbol: "PLZ"},
				currency.PTE: cldr.Currency{DisplayName: "escudo Portugal", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "guarani Paraguay", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "rial Qatar", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "dollar Rodezia", Symbol: "RHD"},
				currency.ROL: cldr.Currency{DisplayName: "leu Roumania (1952–2006)", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "leu Roumania", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "dinar Serbia", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "roubl Rusia", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "roubl Rusia (1991–1998)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "lur Rwanda", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "riyal Arabia Saoudat", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "dollar Inizi Salomon", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "roupi Sechelez", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "dinar Soudan (1992–2007)", Symbol: "SDD"},
				currency.SDG: cldr.Currency{DisplayName: "lur Soudan", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "lur Soudan (1957–1998)", Symbol: "SDP"},
				currency.SEK: cldr.Currency{DisplayName: "kurunenn Sveden", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "dollar Singapour", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "lur Saint-Helena", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "tolar Slovenia", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "kurunenn Slovakia", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "leone Sierra Leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "shilling Somalia", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "dollar Surinam", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "florin Surinam", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "lur Susoudan", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "dobra São Tomé ha Príncipe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "dobra São Tomé ha Príncipe", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "roubl soviedel", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "colón Salvador", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "lur Siria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni Swaziland", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "baht Thailand", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "roubl Tadjikistan", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "somoni Tadjikistan", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "manat Turkmenistan (1993–2009)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "manat Turkmenistan", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "dinar Tunizia", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "paʻanga Tonga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "escudo Timor", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "lur Turkia (1922–2005)", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "lur Turkia", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "dollar Trinidad ha Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "dollar nevez Taiwan", Symbol: "TWD"},
				currency.TZS: cldr.Currency{DisplayName: "shilling Tanzania", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "hryvnia Ukraina", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "shilling Ouganda (1966–1987)", Symbol: "UGS"},
				currency.UGX: cldr.Currency{DisplayName: "shilling Ouganda", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "dollar SU", Symbol: "$ SU"},
				currency.USN: cldr.Currency{DisplayName: "", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "peso Uruguay (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "peso Uruguay", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "som Ouzbekistan", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "bolivar Venezuela (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "bolivar Venezuela (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "bolivar Venezuela", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "dong Viêt Nam", Symbol: "VND"},
				currency.VNN: cldr.Currency{DisplayName: "dong Viêt Nam (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "vatu Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "tala Samoa", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "lur CFA Kreizafrika", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "arcʼhant", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "aour", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "unanenn genaoz europat", Symbol: "XBA"},
				currency.XBB: cldr.Currency{DisplayName: "unanenn voneiz europat", Symbol: "XBB"},
				currency.XBC: cldr.Currency{DisplayName: "unanenn jediñ europat (XBC)", Symbol: "XBC"},
				currency.XBD: cldr.Currency{DisplayName: "unanenn jediñ europat (XBD)", Symbol: "XBD"},
				currency.XCD: cldr.Currency{DisplayName: "dollar Karib ar reter", Symbol: "XCD"},
				currency.XDR: cldr.Currency{DisplayName: "gwirioù tennañ arbennik", Symbol: "XDR"},
				currency.XEU: cldr.Currency{DisplayName: "unanenn jediñ europat", Symbol: "XEU"},
				currency.XFO: cldr.Currency{DisplayName: "lur aour Frañs", Symbol: "XFO"},
				currency.XFU: cldr.Currency{DisplayName: "lur Unaniezh etrebroadel an hentoù-houarn", Symbol: "XFU"},
				currency.XOF: cldr.Currency{DisplayName: "lur CFA Afrika ar Cʼhornôg", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "palladiom", Symbol: "XPD"},
				currency.XPF: cldr.Currency{DisplayName: "lur CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "platin", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "", Symbol: "XSU"},
				currency.XTS: cldr.Currency{DisplayName: "kod moneiz amprouiñ", Symbol: "XTS"},
				currency.XXX: cldr.Currency{DisplayName: "moneiz dianav", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "dinar Yemen", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "rial Yemen", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "dinar nevez Yougoslavia (1994–2002)", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "dinar kemmadus Yougoslavia (1990–1992)", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "dinar adreizhet Yougoslavia (1992–1993)", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "", Symbol: "ZAL"},
				currency.ZAR: cldr.Currency{DisplayName: "rand Suafrika", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "kwacha Zambia (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha Zambia", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "", Symbol: "ZRN"},
				currency.ZRZ: cldr.Currency{DisplayName: "", Symbol: "ZRZ"},
				currency.ZWD: cldr.Currency{DisplayName: "dollar Zimbabwe (1980–2008)", Symbol: "ZWD"},
				currency.ZWL: cldr.Currency{DisplayName: "dollar Zimbabwe (2009)", Symbol: "ZWL"},
				currency.ZWR: cldr.Currency{DisplayName: "dollar Zimbabwe (2008)", Symbol: "ZWR"},
			},
		},
	}
}
