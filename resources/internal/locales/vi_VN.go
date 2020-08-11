// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_vi_VN() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "vi_VN",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "dd/MM/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{0} {1}", Long: "{0} {1}", Medium: "{0}, {1}", Short: "{0}, {1}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Thg 1", Feb: "Thg 2", Mar: "Thg 3", Apr: "Thg 4", May: "Thg 5", Jun: "Thg 6", Jul: "Thg 7", Aug: "Thg 8", Sep: "Thg 9", Oct: "Thg 10", Nov: "Thg 11", Dec: "Thg 12"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "Tháng 1", Feb: "Tháng 2", Mar: "Tháng 3", Apr: "Tháng 4", May: "Tháng 5", Jun: "Tháng 6", Jul: "Tháng 7", Aug: "Tháng 8", Sep: "Tháng 9", Oct: "Tháng 10", Nov: "Tháng 11", Dec: "Tháng 12"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "CN", Mon: "Th 2", Tue: "Th 3", Wed: "Th 4", Thu: "Th 5", Fri: "Th 6", Sat: "Th 7"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "CN", Mon: "T2", Tue: "T3", Wed: "T4", Thu: "T5", Fri: "T6", Sat: "T7"}, Short: cldr.CalendarDayFormatNameValue{Sun: "CN", Mon: "T2", Tue: "T3", Wed: "T4", Thu: "T5", Fri: "T6", Sat: "T7"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Chủ Nhật", Mon: "Thứ Hai", Tue: "Thứ Ba", Wed: "Thứ Tư", Thu: "Thứ Năm", Fri: "Thứ Sáu", Sat: "Thứ Bảy"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "SA", PM: "CH"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "SA", PM: "CH"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "SA", PM: "CH"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Đồng Peseta của Andora", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "Dirham UAE", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "Đồng Afghani của Afghanistan (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghani Afghanistan", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "Lek Albania", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "Dram Armenia", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "Guilder Antille Hà Lan", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "Kwanza Angola", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "Đồng Kwanza của Angola (1977–1991)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Đồng Kwanza Mới của Angola (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Đồng Kwanza Điều chỉnh lại của Angola (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Đồng Austral của Argentina", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Đồng Peso Ley của Argentina (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Đồng Peso Argentina (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Đồng Peso Argentina (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Peso Argentina", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "Đồng Schiling Áo", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Đô la Australia", Symbol: "AU$"},
				currency.AWG: cldr.Currency{DisplayName: "Florin Aruba", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "Đồng Manat của Azerbaijan (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Manat Azerbaijan", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "Đồng Dinar của Bosnia-Herzegovina (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Mark Bosnia-Herzegovina có thể chuyển đổi", Symbol: "BAM"},
				currency.BAN: cldr.Currency{DisplayName: "Đồng Dinar Mới của Bosnia-Herzegovina (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Đô la Barbados", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "Taka Bangladesh", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "Đồng Franc Bỉ (có thể chuyển đổi)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Đồng Franc Bỉ", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Đồng Franc Bỉ (tài chính)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Đồng Lev Xu của Bun-ga-ri", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Đồng Lev Xã hội chủ nghĩa của Bun-ga-ri", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Lev Bulgaria", Symbol: "BGN"},
				currency.BGO: cldr.Currency{DisplayName: "Đồng Lev của Bun-ga-ri (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Dinar Bahrain", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "Franc Burundi", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "Đô la Bermuda", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "Đô la Brunei", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "Boliviano Bolivia", Symbol: "BOB"},
				currency.BOL: cldr.Currency{DisplayName: "Đồng Boliviano của Bolivia (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Đồng Peso Bolivia", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Đồng Mvdol Bolivia", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Đồng Cruzerio Mới của Braxin (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Đồng Cruzado của Braxin (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Đồng Cruzerio của Braxin (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Real Braxin", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Đồng Cruzado Mới của Braxin (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Đồng Cruzeiro của Braxin (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Đồng Cruzeiro của Braxin (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Đô la Bahamas", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "Ngultrum Bhutan", Symbol: "BTN"},
				currency.BUK: cldr.Currency{DisplayName: "Đồng Kyat Miến Điện", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Pula Botswana", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "Đồng Rúp Mới của Belarus (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Rúp Belarus", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "Rúp Belarus (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Đô la Belize", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "Đô la Canada", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Franc Congo", Symbol: "CDF"},
				currency.CHE: cldr.Currency{DisplayName: "Đồng Euro WIR", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Franc Thụy sĩ", Symbol: "CHF"},
				currency.CHW: cldr.Currency{DisplayName: "Đồng France WIR", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "Đồng Escudo của Chile", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Đơn vị Kế toán của Chile (UF)", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Peso Chile", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "Nhân dân tệ (hải ngoại)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "Nhân dân tệ", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Peso Colombia", Symbol: "COP"},
				currency.COU: cldr.Currency{DisplayName: "Đơn vị Giá trị Thực của Colombia", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Colón Costa Rica", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "Đồng Dinar của Serbia (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Đồng Koruna Xu của Czechoslovakia", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Peso Cuba có thể chuyển đổi", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "Peso Cuba", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "Escudo Cape Verde", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "Đồng Bảng Síp", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Koruna Cộng hòa Séc", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "Đồng Mark Đông Đức", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "Đồng Mark Đức", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Franc Djibouti", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "Krone Đan Mạch", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "Peso Dominica", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "Dinar Algeria", Symbol: "DZD"},
				currency.ECS: cldr.Currency{DisplayName: "Đồng Scure Ecuador", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Đơn vị Giá trị Không đổi của Ecuador", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Crun Extônia", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Bảng Ai Cập", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "Nakfa Eritrea", Symbol: "ERN"},
				currency.ESA: cldr.Currency{DisplayName: "Đồng Peseta Tây Ban Nha (Tài khoản)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Đồng Peseta Tây Ban Nha (tài khoản có thể chuyển đổi)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Đồng Peseta Tây Ban Nha", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Birr Ethiopia", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Đồng Markka Phần Lan", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Đô la Fiji", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "Bảng Quần đảo Falkland", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "Franc Pháp", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "Bảng Anh", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Đồng Kupon Larit của Georgia", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Lari Georgia", Symbol: "GEL"},
				currency.GHC: cldr.Currency{DisplayName: "Cedi Ghana (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Cedi Ghana", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "Bảng Gibraltar", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "Dalasi Gambia", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "Franc Guinea", Symbol: "GNF"},
				currency.GNS: cldr.Currency{DisplayName: "Syli Guinea", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Đồng Ekwele của Guinea Xích Đạo", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Drachma Hy Lạp", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Quetzal Guatemala", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "Đồng Guinea Escudo Bồ Đào Nha", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Peso Guinea-Bissau", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Đô la Guyana", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "Đô la Hồng Kông", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Lempira Honduras", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "Đồng Dinar Croatia", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Kuna Croatia", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "Gourde Haiti", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "Forint Hungary", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "Rupiah Indonesia", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "Pao Ai-len", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Pao Ixraen", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Sheqel Israel mới", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Rupee Ấn Độ", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Dinar Iraq", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "Rial Iran", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "Króna Iceland", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "Lia Ý", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Đô la Jamaica", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "Dinar Jordan", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "Yên Nhật", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Shilling Kenya", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "Som Kyrgyzstan", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "Riel Campuchia", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "Franc Comoros", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "Won Triều Tiên", Symbol: "KPW"},
				currency.KRH: cldr.Currency{DisplayName: "Đồng Hwan Hàn Quốc (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "Đồng Won Hàn Quốc (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "Won Hàn Quốc", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Dinar Kuwait", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "Đô la Quần đảo Cayman", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "Tenge Kazakhstan", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "Kip Lào", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "Bảng Li-băng", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "Rupee Sri Lanka", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "Đô la Liberia", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "Ioti Lesotho", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Litas Lít-va", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "Đồng Talonas Litva", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Đồng Franc Luxembourg có thể chuyển đổi", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Đồng Franc Luxembourg", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Đồng Franc Luxembourg tài chính", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Lats Latvia", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "Đồng Rúp Latvia", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Dinar Libi", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "Dirham Ma-rốc", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "Đồng Franc Ma-rốc", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Đồng Franc Monegasque", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "Đồng Cupon Moldova", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Leu Moldova", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "Ariary Malagasy", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "Đồng Franc Magalasy", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Denar Macedonia", Symbol: "MKD"},
				currency.MKN: cldr.Currency{DisplayName: "Đồng Denar Macedonia (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Đồng Franc Mali", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Kyat Myanma", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "Tugrik Mông Cổ", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "Pataca Ma Cao", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "Ouguiya Mauritania (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Ouguiya Mauritania", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "Lia xứ Man-tơ", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Đồng Bảng Malta", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Rupee Mauritius", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "Rufiyaa Maldives", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "Kwacha Malawi", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "Peso Mexico", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Đồng Peso Bạc Mê-hi-cô (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Đơn vị Đầu tư Mê-hi-cô", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Ringgit Malaysia", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "Escudo Mozambique", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Đồng Metical Mozambique (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Metical Mozambique", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "Đô la Namibia", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "Naira Nigeria", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "Đồng Córdoba Nicaragua (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Córdoba Nicaragua", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "Đồng Guilder Hà Lan", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Krone Na Uy", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "Rupee Nepal", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "Đô la New Zealand", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Rial Oman", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "Balboa Panama", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "Đồng Inti Peru", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Sol Peru", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "Đồng Sol Peru (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Kina Papua New Guinea", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "Peso Philipin", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "Rupee Pakistan", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "Zloty Ba Lan", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "Đồng Zloty Ba Lan (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Đồng Escudo Bồ Đào Nha", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Guarani Paraguay", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "Rial Qatar", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "Đồng Đô la Rhode", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Đồng Leu Rumani (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Leu Romania", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "Dinar Serbia", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "Rúp Nga", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "Đồng Rúp Nga (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Franc Rwanda", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "Riyal Ả Rập Xê-út", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "Đô la quần đảo Solomon", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "Rupee Seychelles", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "Đồng Dinar Sudan (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Bảng Sudan", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "Đồng Bảng Sudan (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Krona Thụy Điển", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "Đô la Singapore", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "Bảng St. Helena", Symbol: "SHP"},
				currency.SIT: cldr.Currency{DisplayName: "Tôla Xlôvênia", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Cuaron Xlôvác", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Leone Sierra Leone", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "Schilling Somali", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "Đô la Suriname", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "Đồng Guilder Surinam", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "Bảng Nam Sudan", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "Dobra São Tomé và Príncipe (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "Dobra São Tomé và Príncipe", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "Đồng Rúp Sô viết", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Colón El Salvador", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Bảng Syria", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "Lilangeni Swaziland", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "Bạt Thái Lan", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Đồng Rúp Tajikistan", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Somoni Tajikistan", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "Đồng Manat Turkmenistan (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Manat Turkmenistan", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "Dinar Tunisia", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "Paʻanga Tonga", Symbol: "TOP"},
				currency.TPE: cldr.Currency{DisplayName: "Đồng Escudo Timor", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Lia Thổ Nhĩ Kỳ (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Lia Thổ Nhĩ Kỳ", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "Đô la Trinidad và Tobago", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "Đô la Đài Loan mới", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Shilling Tanzania", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "Hryvnia Ukraina", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "Đồng Karbovanets Ucraina", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Đồng Shilling Uganda (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Shilling Uganda", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "Đô la Mỹ", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "Đô la Mỹ (Ngày tiếp theo)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "Đô la Mỹ (Cùng ngày)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Đồng Peso Uruguay (Đơn vị Theo chỉ số)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Đồng Peso Uruguay (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Peso Uruguay", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "Som Uzbekistan", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "Đồng bolívar của Venezuela (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Bolívar Venezuela (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Bolívar Venezuela", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "Đồng Việt Nam", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Đồng Việt Nam (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Vatu Vanuatu", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "Tala Samoa", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "Franc CFA Trung Phi", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Bạc", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Vàng", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "Đơn vị Tổng hợp Châu Âu", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "Đơn vị Tiền tệ Châu Âu", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "Đơn vị Kế toán Châu Âu (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "Đơn vị Kế toán Châu Âu (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "Đô la Đông Caribê", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Quyền Rút vốn Đặc biệt", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "Đơn vị Tiền Châu Âu", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "Đồng France Pháp Vàng", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "Đồng UIC-Franc Pháp", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "Franc CFA Tây Phi", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Paladi", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "Franc CFP", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Bạch kim", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "Quỹ RINET", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Mã Tiền tệ Kiểm tra", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Tiền tệ chưa biết", Symbol: "XXX"},
				currency.YDD: cldr.Currency{DisplayName: "Đồng Dinar Yemen", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Rial Yemen", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "Đồng Dinar Nam Tư Xu (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Đồng Dinar Nam Tư Mới (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Đồng Dinar Nam Tư Có thể chuyển đổi (1990–1992)", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "Đồng Dinar Nam Tư Tái cơ cấu (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "Đồng Rand Nam Phi (tài chính)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Rand Nam Phi", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "Đồng kwacha của Zambia (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Kwacha Zambia", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "Đồng Zaire Mới (1993–1998)", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Đồng Zaire (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Đồng Đô la Zimbabwe (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Đồng Đô la Zimbabwe (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Đồng Đô la Zimbabwe (2008)", Symbol: ""},
			},
		},
	}
}
