// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_tr_CY() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "tr_CY",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "d MMMM y EEEE", Long: "d MMMM y", Medium: "d MMM y", Short: "d.MM.y"}, Time: cldr.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Oca", Feb: "Şub", Mar: "Mar", Apr: "Nis", May: "May", Jun: "Haz", Jul: "Tem", Aug: "Ağu", Sep: "Eyl", Oct: "Eki", Nov: "Kas", Dec: "Ara"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "O", Feb: "Ş", Mar: "M", Apr: "N", May: "M", Jun: "H", Jul: "T", Aug: "A", Sep: "E", Oct: "E", Nov: "K", Dec: "A"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Ocak", Feb: "Şubat", Mar: "Mart", Apr: "Nisan", May: "Mayıs", Jun: "Haziran", Jul: "Temmuz", Aug: "Ağustos", Sep: "Eylül", Oct: "Ekim", Nov: "Kasım", Dec: "Aralık"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Paz", Mon: "Pzt", Tue: "Sal", Wed: "Çar", Thu: "Per", Fri: "Cum", Sat: "Cmt"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "P", Mon: "P", Tue: "S", Wed: "Ç", Thu: "P", Fri: "C", Sat: "C"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Pa", Mon: "Pt", Tue: "Sa", Wed: "Ça", Thu: "Pe", Fri: "Cu", Sat: "Ct"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Pazar", Mon: "Pazartesi", Tue: "Salı", Wed: "Çarşamba", Thu: "Perşembe", Fri: "Cuma", Sat: "Cumartesi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "ÖÖ", PM: "ÖS"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "ÖÖ", PM: "ÖS"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "ÖÖ", PM: "ÖS"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "%#,##0"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorra Pezetası", Symbol: "ADP"},
				currency.AED: cldr.Currency{DisplayName: "Birleşik Arap Emirlikleri Dirhemi", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Afganistan Afganisi (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afganistan Afganisi", Symbol: "AFN"},
				currency.ALK: cldr.Currency{DisplayName: "Arnavutluk Leki (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Arnavutluk Leki", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Ermenistan Dramı", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Hollanda Antilleri Guldeni", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Angola Kvanzası", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Angola Kvanzası (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Yeni Angola Kvanzası (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Angola Kvanzası Reajustado (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Arjantin Australi", Symbol: "ARA"},
				currency.ARL: cldr.Currency{DisplayName: "Arjantin Peso Leyi (1970–1983)", Symbol: "ARL"},
				currency.ARM: cldr.Currency{DisplayName: "Arjantin Pesosu (1881–1970)", Symbol: "ARM"},
				currency.ARP: cldr.Currency{DisplayName: "Arjantin Pezosu (1983–1985)", Symbol: "ARP"},
				currency.ARS: cldr.Currency{DisplayName: "Arjantin Pesosu", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Avusturya Şilini", Symbol: "ATS"},
				currency.AUD: cldr.Currency{DisplayName: "Avustralya Doları", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruba Florini", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Azerbaycan Manatı (1993–2006)", Symbol: "AZM"},
				currency.AZN: cldr.Currency{DisplayName: "Azerbaycan Manatı", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Bosna Hersek Dinarı", Symbol: "BAD"},
				currency.BAM: cldr.Currency{DisplayName: "Konvertibl Bosna Hersek Markı", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Yeni Bosna Hersek Dinarı (1994–1997)", Symbol: "BAN"},
				currency.BBD: cldr.Currency{DisplayName: "Barbados Doları", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladeş Takası", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Belçika Frangı (konvertibl)", Symbol: "BEC"},
				currency.BEF: cldr.Currency{DisplayName: "Belçika Frangı", Symbol: "BEF"},
				currency.BEL: cldr.Currency{DisplayName: "Belçika Frangı (finansal)", Symbol: "BEL"},
				currency.BGL: cldr.Currency{DisplayName: "Bulgar Levası (Hard)", Symbol: "BGL"},
				currency.BGM: cldr.Currency{DisplayName: "Sosyalist Bulgaristan Levası", Symbol: "BGM"},
				currency.BGN: cldr.Currency{DisplayName: "Bulgar Levası", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Bulgar Levası (1879–1952)", Symbol: "BGO"},
				currency.BHD: cldr.Currency{DisplayName: "Bahreyn Dinarı", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Burundi Frangı", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Bermuda Doları", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Brunei Doları", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Bolivya Bolivyanosu", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Bolivya Bolivyanosu (1863–1963)", Symbol: "BOL"},
				currency.BOP: cldr.Currency{DisplayName: "Bolivya Pezosu", Symbol: "BOP"},
				currency.BOV: cldr.Currency{DisplayName: "Bolivya Mvdolu", Symbol: "BOV"},
				currency.BRB: cldr.Currency{DisplayName: "Yeni Brezilya Kruzeirosu (1967–1986)", Symbol: "BRB"},
				currency.BRC: cldr.Currency{DisplayName: "Brezilya Kruzadosu", Symbol: "BRC"},
				currency.BRE: cldr.Currency{DisplayName: "Brezilya Kruzeirosu (1990–1993)", Symbol: "BRE"},
				currency.BRL: cldr.Currency{DisplayName: "Brezilya Reali", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Yeni Brezilya Kruzadosu", Symbol: "BRN"},
				currency.BRR: cldr.Currency{DisplayName: "Brezilya Kruzeirosu", Symbol: "BRR"},
				currency.BRZ: cldr.Currency{DisplayName: "Brezilya Kruzeirosu (1942–1967)", Symbol: "BRZ"},
				currency.BSD: cldr.Currency{DisplayName: "Bahama Doları", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Butan Ngultrumu", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Burma Kyatı", Symbol: "BUK"},
				currency.BWP: cldr.Currency{DisplayName: "Botsvana Pulası", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Yeni Beyaz Rusya Rublesi (1994–1999)", Symbol: "BYB"},
				currency.BYN: cldr.Currency{DisplayName: "Belarus Rublesi", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Beyaz Rusya Rublesi (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "Belize Doları", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Kanada Doları", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Kongo Frangı", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "WIR Avrosu", Symbol: "CHE"},
				currency.CHF: cldr.Currency{DisplayName: "İsviçre Frangı", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "WIR Frangı", Symbol: "CHW"},
				currency.CLE: cldr.Currency{DisplayName: "Şili Esküdosu", Symbol: "CLE"},
				currency.CLF: cldr.Currency{DisplayName: "Şili Unidades de Fomento", Symbol: "CLF"},
				currency.CLP: cldr.Currency{DisplayName: "Şili Pesosu", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Çin Yuanı (offshore)", Symbol: "CNH"},
				currency.CNX: cldr.Currency{DisplayName: "Çin Halk Cumhuriyeti Merkez Bankası Doları", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Çin Yuanı", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolombiya Pesosu", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Unidad de Valor Real", Symbol: "COU"},
				currency.CRC: cldr.Currency{DisplayName: "Kosta Rika Kolonu", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Eski Sırbistan Dinarı", Symbol: "CSD"},
				currency.CSK: cldr.Currency{DisplayName: "Çekoslavak Korunası (Hard)", Symbol: "CSK"},
				currency.CUC: cldr.Currency{DisplayName: "Konvertibl Küba Pesosu", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Küba Pesosu", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Cape Verde Esküdosu", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Güney Kıbrıs Lirası", Symbol: "CYP"},
				currency.CZK: cldr.Currency{DisplayName: "Çek Korunası", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Doğu Alman Markı", Symbol: "DDM"},
				currency.DEM: cldr.Currency{DisplayName: "Alman Markı", Symbol: "DEM"},
				currency.DJF: cldr.Currency{DisplayName: "Cibuti Frangı", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Danimarka Kronu", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Dominik Pesosu", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Cezayir Dinarı", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Ekvador Sukresi", Symbol: "ECS"},
				currency.ECV: cldr.Currency{DisplayName: "Ekvador Unidad de Valor Constante (UVC)", Symbol: "ECV"},
				currency.EEK: cldr.Currency{DisplayName: "Estonya Krunu", Symbol: "EEK"},
				currency.EGP: cldr.Currency{DisplayName: "Mısır Lirası", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Eritre Nakfası", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "İspanyol Pezetası (A hesabı)", Symbol: "ESA"},
				currency.ESB: cldr.Currency{DisplayName: "İspanyol Pezetası (konvertibl hesap)", Symbol: "ESB"},
				currency.ESP: cldr.Currency{DisplayName: "İspanyol Pezetası", Symbol: "ESP"},
				currency.ETB: cldr.Currency{DisplayName: "Etiyopya Birri", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Fin Markkası", Symbol: "FIM"},
				currency.FJD: cldr.Currency{DisplayName: "Fiji Doları", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Falkland Adaları Lirası", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Fransız Frangı", Symbol: "FRF"},
				currency.GBP: cldr.Currency{DisplayName: "İngiliz Sterlini", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Gürcistan Kupon Larisi", Symbol: "GEK"},
				currency.GEL: cldr.Currency{DisplayName: "Gürcistan Larisi", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Gana Sedisi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Gana Sedisi", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Cebelitarık Lirası", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Gambiya Dalasisi", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Gine Frangı", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Gine Sylisi", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Ekvator Ginesi Ekuelesi", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Yunan Drahmisi", Symbol: "GRD"},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemala Quetzalı", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Portekiz Ginesi Esküdosu", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Gine-Bissau Pezosu", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Guyana Doları", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Hong Kong Doları", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Honduras Lempirası", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Hırvatistan Dinarı", Symbol: "HRD"},
				currency.HRK: cldr.Currency{DisplayName: "Hırvatistan Kunası", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Haiti Gurdu", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Macar Forinti", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Endonezya Rupisi", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "İrlanda Lirası", Symbol: "IEP"},
				currency.ILP: cldr.Currency{DisplayName: "İsrail Lirası", Symbol: "ILP"},
				currency.ILR: cldr.Currency{DisplayName: "İsrail Şekeli (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Yeni İsrail Şekeli", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Hindistan Rupisi", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Irak Dinarı", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "İran Riyali", Symbol: "IRR"},
				currency.ISJ: cldr.Currency{DisplayName: "İzlanda Kronu (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "İzlanda Kronu", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "İtalyan Lireti", Symbol: "ITL"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaika Doları", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Ürdün Dinarı", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Japon Yeni", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenya Şilini", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Kırgızistan Somu", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Kamboçya Rieli", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Komorlar Frangı", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Kuzey Kore Wonu", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Güney Kore Hwanı (1953–1962)", Symbol: "KRH"},
				currency.KRO: cldr.Currency{DisplayName: "Güney Kore Wonu (1945–1953)", Symbol: "KRO"},
				currency.KRW: cldr.Currency{DisplayName: "Güney Kore Wonu", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuveyt Dinarı", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Cayman Adaları Doları", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Kazakistan Tengesi", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Laos Kipi", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Lübnan Lirası", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lanka Rupisi", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Liberya Doları", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Lesotho Lotisi", Symbol: "LSL"},
				currency.LTL: cldr.Currency{DisplayName: "Litvanya Litası", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "Litvanya Talonu", Symbol: "LTT"},
				currency.LUC: cldr.Currency{DisplayName: "Konvertibl Lüksemburg Frangı", Symbol: "LUC"},
				currency.LUF: cldr.Currency{DisplayName: "Lüksemburg Frangı", Symbol: "LUF"},
				currency.LUL: cldr.Currency{DisplayName: "Finansal Lüksemburg Frangı", Symbol: "LUL"},
				currency.LVL: cldr.Currency{DisplayName: "Letonya Latı", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "Letonya Rublesi", Symbol: "LVR"},
				currency.LYD: cldr.Currency{DisplayName: "Libya Dinarı", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Fas Dirhemi", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Fas Frangı", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Monako Frangı", Symbol: "MCF"},
				currency.MDC: cldr.Currency{DisplayName: "Moldova Kuponu", Symbol: "MDC"},
				currency.MDL: cldr.Currency{DisplayName: "Moldova Leyi", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Madagaskar Ariarisi", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Madagaskar Frangı", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Makedonya Dinarı", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Makedonya Dinarı (1992–1993)", Symbol: "MKN"},
				currency.MLF: cldr.Currency{DisplayName: "Mali Frangı", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Myanmar Kyatı", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Moğolistan Tugriki", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Makao Patakası", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Moritanya Ugiyası (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "Moritanya Ugiyası", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Malta Lirası", Symbol: "MTL"},
				currency.MTP: cldr.Currency{DisplayName: "Malta Sterlini", Symbol: "MTP"},
				currency.MUR: cldr.Currency{DisplayName: "Mauritius Rupisi", Symbol: "MUR"},
				currency.MVP: cldr.Currency{DisplayName: "Maldiv Rupisi", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "Maldiv Rufiyaası", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Malavi Kvaçası", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Meksika Pesosu", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Gümüş Meksika Pezosu (1861–1992)", Symbol: "MXP"},
				currency.MXV: cldr.Currency{DisplayName: "Meksika Unidad de Inversion (UDI)", Symbol: "MXV"},
				currency.MYR: cldr.Currency{DisplayName: "Malezya Ringgiti", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Mozambik Esküdosu", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Eski Mozambik Metikali", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mozambik Metikali", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Namibya Doları", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Nijerya Nairası", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Nikaragua Kordobası (1988–1991)", Symbol: "NIC"},
				currency.NIO: cldr.Currency{DisplayName: "Nikaragua Kordobası", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Hollanda Florini", Symbol: "NLG"},
				currency.NOK: cldr.Currency{DisplayName: "Norveç Kronu", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Nepal Rupisi", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Yeni Zelanda Doları", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Umman Riyali", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Panama Balboası", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Peru İnti", Symbol: "PEI"},
				currency.PEN: cldr.Currency{DisplayName: "Peru Solü", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Peru Solü (1863–1965)", Symbol: "PES"},
				currency.PGK: cldr.Currency{DisplayName: "Papua Yeni Gine Kinası", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Filipinler Pesosu", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistan Rupisi", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Polonya Zlotisi", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Polonya Zlotisi (1950–1995)", Symbol: "PLZ"},
				currency.PTE: cldr.Currency{DisplayName: "Portekiz Esküdosu", Symbol: "PTE"},
				currency.PYG: cldr.Currency{DisplayName: "Paraguay Guaranisi", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Katar Riyali", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Rodezya Doları", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Eski Romen Leyi", Symbol: "ROL"},
				currency.RON: cldr.Currency{DisplayName: "Romen Leyi", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Sırp Dinarı", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rus Rublesi", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Rus Rublesi (1991–1998)", Symbol: "RUR"},
				currency.RWF: cldr.Currency{DisplayName: "Ruanda Frangı", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Suudi Arabistan Riyali", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Solomon Adaları Doları", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Seyşeller Rupisi", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Eski Sudan Dinarı", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Sudan Lirası", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Eski Sudan Lirası", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "İsveç Kronu", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Singapur Doları", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Saint Helena Lirası", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Slovenya Toları", Symbol: "SIT"},
				currency.SKK: cldr.Currency{DisplayName: "Slovak Korunası", Symbol: "SKK"},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leone Leonesi", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Somali Şilini", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Surinam Doları", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Surinam Guldeni", Symbol: "SRG"},
				currency.SSP: cldr.Currency{DisplayName: "Güney Sudan Lirası", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "São Tomé ve Príncipe Dobrası (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "Sao Tome ve Principe Dobrası", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Sovyet Rublesi", Symbol: "SUR"},
				currency.SVC: cldr.Currency{DisplayName: "El Salvador Kolonu", Symbol: "SVC"},
				currency.SYP: cldr.Currency{DisplayName: "Suriye Lirası", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Svaziland Lilangenisi", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Tayland Bahtı", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tacikistan Rublesi", Symbol: "TJR"},
				currency.TJS: cldr.Currency{DisplayName: "Tacikistan Somonisi", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Türkmenistan Manatı (1993–2009)", Symbol: "TMM"},
				currency.TMT: cldr.Currency{DisplayName: "Türkmenistan Manatı", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Tunus Dinarı", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Tonga Paʻangası", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Timor Esküdosu", Symbol: "TPE"},
				currency.TRL: cldr.Currency{DisplayName: "Eski Türk Lirası", Symbol: "TRL"},
				currency.TRY: cldr.Currency{DisplayName: "Türk Lirası", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad ve Tobago Doları", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Yeni Tayvan Doları", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzanya Şilini", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Ukrayna Grivnası", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrayna Karbovanetz", Symbol: "UAK"},
				currency.UGS: cldr.Currency{DisplayName: "Uganda Şilini (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Uganda Şilini", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "ABD Doları", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "ABD Doları (Ertesi gün)", Symbol: "USN"},
				currency.USS: cldr.Currency{DisplayName: "ABD Doları (Aynı gün)", Symbol: "USS"},
				currency.UYI: cldr.Currency{DisplayName: "Uruguay Peso en Unidades Indexadas", Symbol: "UYI"},
				currency.UYP: cldr.Currency{DisplayName: "Uruguay Pezosu (1975–1993)", Symbol: "UYP"},
				currency.UYU: cldr.Currency{DisplayName: "Uruguay Pesosu", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Özbekistan Somu", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Venezuela Bolivarı (1871–2008)", Symbol: "VEB"},
				currency.VEF: cldr.Currency{DisplayName: "Venezuela Bolivarı (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "Venezuela Bolivarı", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Vietnam Dongu", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Vietnam Dongu (1978–1985)", Symbol: "VNN"},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu Vatusu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Samoa Talası", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Orta Afrika CFA Frangı", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Gümüş", Symbol: "XAG"},
				currency.XAU: cldr.Currency{DisplayName: "Altın", Symbol: "XAU"},
				currency.XBA: cldr.Currency{DisplayName: "Birleşik Avrupa Birimi", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Avrupa Para Birimi (EMU)", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Avrupa Hesap Birimi (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Avrupa Hesap Birimi (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Doğu Karayip Doları", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Özel Çekme Hakkı (SDR)", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Avrupa Para Birimi", Symbol: "XEU"},
				currency.XFO: cldr.Currency{DisplayName: "Fransız Altın Frangı", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Fransız UIC-Frangı", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "Batı Afrika CFA Frangı", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Paladyum", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP Frangı", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platin", Symbol: "XPT"},
				currency.XRE: cldr.Currency{DisplayName: "RINET Fonları", Symbol: "XRE"},
				currency.XSU: cldr.Currency{DisplayName: "Sucre", Symbol: "XSU"},
				currency.XTS: cldr.Currency{DisplayName: "Test Para Birimi Kodu", Symbol: "XTS"},
				currency.XUA: cldr.Currency{DisplayName: "ADB Hesap Birimi", Symbol: "XUA"},
				currency.XXX: cldr.Currency{DisplayName: "Bilinmeyen Para Birimi", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Yemen Dinarı", Symbol: "YDD"},
				currency.YER: cldr.Currency{DisplayName: "Yemen Riyali", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Yugoslav Dinarı (Hard)", Symbol: "YUD"},
				currency.YUM: cldr.Currency{DisplayName: "Yeni Yugoslav Dinarı", Symbol: "YUM"},
				currency.YUN: cldr.Currency{DisplayName: "Konvertibl Yugoslav Dinarı", Symbol: "YUN"},
				currency.YUR: cldr.Currency{DisplayName: "İyileştirilmiş Yugoslav Dinarı (1992–1993)", Symbol: "YUR"},
				currency.ZAL: cldr.Currency{DisplayName: "Güney Afrika Randı (finansal)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Güney Afrika Randı", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambiya Kvaçası (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "Zambiya Kvaçası", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Yeni Zaire Zairesi", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zaire Zairesi", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabve Doları", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabve Doları (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabve Doları (2008)", Symbol: "ZWR"},
			},
		},
	}
}
