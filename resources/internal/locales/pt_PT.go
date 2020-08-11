// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_pt_PT() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "pt_PT",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "dd/MM/y", Short: "dd/MM/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'às' {0}", Long: "{1} 'às' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "jan.", Feb: "fev.", Mar: "mar.", Apr: "abr.", May: "mai.", Jun: "jun.", Jul: "jul.", Aug: "ago.", Sep: "set.", Oct: "out.", Nov: "nov.", Dec: "dez."}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "janeiro", Feb: "fevereiro", Mar: "março", Apr: "abril", May: "maio", Jun: "junho", Jul: "julho", Aug: "agosto", Sep: "setembro", Oct: "outubro", Nov: "novembro", Dec: "dezembro"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "domingo", Mon: "segunda", Tue: "terça", Wed: "quarta", Thu: "quinta", Fri: "sexta", Sat: "sábado"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "D", Mon: "S", Tue: "T", Wed: "Q", Thu: "Q", Fri: "S", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "", Mon: "", Tue: "", Wed: "", Thu: "", Fri: "", Sat: ""}, Wide: cldr.CalendarDayFormatNameValue{Sun: "domingo", Mon: "segunda-feira", Tue: "terça-feira", Wed: "quarta-feira", Thu: "quinta-feira", Fri: "sexta-feira", Sat: "sábado"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a.m.", PM: "p.m."}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "manhã", PM: "tarde"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "0 milhão", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤;(#,##0.00\u00a0¤)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Peseta de Andorra", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "dirham dos Emirados Árabes Unidos", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afeghani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "afegâni afegão", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "Lek Albanês (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "lek albanês", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "dram arménio", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "florim das Antilhas Holandesas", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "kwanza angolano", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Cuanza angolano (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Novo cuanza angolano (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Cuanza angolano reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Austral argentino", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Peso lei argentino (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Peso argentino (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Peso argentino (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "peso argentino", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Xelim austríaco", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "dólar australiano", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "florim de Aruba", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Manat azerbaijano (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "manat azeri", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Dinar da Bósnia-Herzegóvina", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "marco bósnio-herzegóvino conversível", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Novo dinar da Bósnia-Herzegovina (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "dólar barbadense", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "taka bengali", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Franco belga (convertível)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Franco belga", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Franco belga (financeiro)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Lev forte búlgaro", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Lev socialista búlgaro", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "lev búlgaro", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Lev búlgaro (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "dinar baremita", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "franco burundiano", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "dólar bermudense", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "dólar bruneano", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "boliviano", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Boliviano (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Peso boliviano", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Mvdol boliviano", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Cruzeiro novo brasileiro (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Cruzado brasileiro (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Cruzeiro brasileiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "real brasileiro", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Cruzado novo brasileiro (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Cruzeiro brasileiro (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Cruzeiro brasileiro (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "dólar das Bahamas", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ngultrum butanês", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Kyat birmanês", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "pula de Botswana", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Rublo novo bielorusso (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "rublo bielorrusso", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Rublo bielorrusso (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "dólar belizense", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "dólar canadiano", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "franco congolês", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "Euro WIR", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "franco suíço", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "Franco WIR", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "Escudo chileno", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Unidades de Fomento chilenas", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "peso chileno", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "yuan offshore", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "Dólar do Banco Popular da China", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "peso colombiano", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Unidade de Valor Real", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "colon costa-riquenho", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Dinar sérvio (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Coroa Forte checoslovaca", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "peso cubano conversível", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "peso cubano", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "escudo cabo-verdiano", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Libra de Chipre", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "coroa checa", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Ostmark da Alemanha Oriental", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Marco alemão", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "franco jibutiano", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "coroa dinamarquesa", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "peso dominicano", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "dinar argelino", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Sucre equatoriano", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Unidad de Valor Constante (UVC) do Equador", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Coroa estoniana", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "libra egípcia", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "nakfa eritreia", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Peseta espanhola (conta A)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Peseta espanhola (conta conversível)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Peseta espanhola", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "birr etíope", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Marca finlandesa", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "dólar fijiano", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "libra das Ilhas Falkland", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Franco francês", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "libra esterlina britânica", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Cupom Lari georgiano", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "lari georgiano", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Cedi do Gana", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "cedi ganês", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "libra de Gibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "dalasi gambiano", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "franco guineense", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Syli da Guiné", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Ekwele da Guiné Equatorial", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Dracma grego", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "quetzal da Guatemala", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Escudo da Guiné Portuguesa", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Peso da Guiné-Bissau", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "dólar da Guiana", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "dólar de Hong Kong", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "lempira das Honduras", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Dinar croata", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "kuna croata", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "gourde haitiano", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "forint húngaro", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "rupia indonésia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Libra irlandesa", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Libra israelita", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "Sheqel antigo israelita", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "sheqel novo israelita", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "rupia indiana", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "dinar iraquiano", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "rial iraniano", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "Coroa antiga islandesa", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "coroa islandesa", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Lira italiana", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "dólar jamaicano", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "dinar jordaniano", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "iene japonês", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "xelim queniano", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "som quirguiz", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "riel cambojano", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "franco comoriano", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "won norte-coreano", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Hwan da Coreia do Sul (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "Won da Coreia do Sul (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "won sul-coreano", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "dinar kuwaitiano", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "dólar das Ilhas Caimão", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "tenge cazaque", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "kip laosiano", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "libra libanesa", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "rupia do Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "dólar liberiano", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Loti do Lesoto", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litas da Lituânia", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Talonas lituano", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Franco conversível de Luxemburgo", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Franco luxemburguês", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Franco financeiro de Luxemburgo", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Lats da Letónia", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Rublo letão", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "dinar líbio", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "dirham marroquino", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Franco marroquino", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Franco monegasco", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "Cupon moldávio", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "leu moldavo", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "ariari malgaxe", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Franco de Madagascar", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "dinar macedónio", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Dinar macedônio (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Franco do Mali", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "kyat de Mianmar", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "tugrik mongol", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "pataca macaense", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "ouguiya mauritana (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "ouguiya mauritana", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Lira maltesa", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Libra maltesa", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "rupia mauriciana", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "rupia maldivana", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "kwacha malauiano", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "peso mexicano", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Peso Plata mexicano (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Unidad de Inversion (UDI) mexicana", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "ringgit malaio", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Escudo de Moçambique", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Metical de Moçambique (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "metical moçambicano", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "dólar namibiano", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "naira nigeriana", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Córdoba nicaraguano (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "córdoba nicaraguano", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Florim holandês", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "coroa norueguesa", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "rupia nepalesa", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "dólar neozelandês", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "rial omanense", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "balboa do Panamá", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Inti peruano", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "sol peruano", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Sol peruano (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "kina papuásia", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "peso filipino", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "rupia paquistanesa", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "zloti polaco", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Zloti polaco (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "escudo português", Symbol: "\u200b"},
				currency.PYG: cldr.Currency{DisplayName: "guarani paraguaio", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "rial catarense", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Dólar rodesiano", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Leu romeno (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "leu romeno", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "dinar sérvio", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "rublo russo", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Rublo russo (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "franco ruandês", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "rial saudita", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "dólar das Ilhas Salomão", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "rupia seichelense", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Dinar sudanês (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "libra sudanesa", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Libra sudanesa (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "coroa sueca", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "dólar singapuriano", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "libra santa-helenense", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Tolar Bons esloveno", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Coroa eslovaca", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "leone de Serra Leoa", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "xelim somali", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "dólar do Suriname", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Florim do Suriname", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "libra sul-sudanesa", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Dobra de São Tomé e Príncipe (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "dobra de São Tomé e Príncipe", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Rublo soviético", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Colom salvadorenho", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "libra síria", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "lilangeni suázi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "baht tailandês", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Rublo do Tadjiquistão", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "somoni tajique", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Manat do Turcomenistão (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "manat turcomeno", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "dinar tunisino", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "paʻanga tonganesa", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Escudo timorense", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Lira turca (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "lira turca", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "dólar de Trindade e Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "novo dólar taiwanês", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "xelim tanzaniano", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "hryvnia ucraniano", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Karbovanetz ucraniano", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Xelim ugandense (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "xelim ugandense", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "dólar dos Estados Unidos", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "Dólar norte-americano (Dia seguinte)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "Dólar norte-americano (Mesmo dia)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Peso uruguaio en unidades indexadas", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Peso uruguaio (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "peso uruguaio", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "som uzbeque", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Bolívar venezuelano (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "bolívar (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "bolívar", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "dong vietnamita", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Dong vietnamita (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "vatu de Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "tala samoano", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "franco CFA (BEAC)", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Prata", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Ouro", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Unidade Composta Europeia", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Unidade Monetária Europeia", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Unidade de Conta Europeia (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Unidade de Conta Europeia (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "dólar das Caraíbas Orientais", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "direito especial de saque", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Unidade da Moeda Europeia", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Franco-ouro francês", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Franco UIC francês", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "franco CFA (BCEAO)", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Paládio", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "franco CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platina", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "Fundos RINET", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Código de Moeda de Teste", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "moeda desconhecida", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Dinar iemenita", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "rial iemenita", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Dinar forte jugoslavo", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Super Dinar jugoslavo", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Dinar conversível jugoslavo", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "Dinar reformado iugoslavo (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Rand sul-africano (financeiro)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "rand sul-africano", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Kwacha zambiano (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "kwacha zambiano", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Zaire Novo zairense (1993–1998)", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zaire zairense (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Dólar do Zimbabwe", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Dólar do Zimbábue (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Dólar do Zimbábue (2008)", Symbol: ""},
			},
		},
	}
}
