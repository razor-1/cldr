// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_en_TV() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "en_TV",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mar", Apr: "Apr", May: "May", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Oct", Nov: "Nov", Dec: "Dec"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "January", Feb: "February", Mar: "March", Apr: "April", May: "May", Jun: "June", Jul: "July", Aug: "August", Sep: "September", Oct: "October", Nov: "November", Dec: "December"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "Sun", Mon: "Mon", Tue: "Tue", Wed: "Wed", Thu: "Thu", Fri: "Fri", Sat: "Sat"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "T", Wed: "W", Thu: "T", Fri: "F", Sat: "S"}, Short: cldr.CalendarDayFormatNameValue{Sun: "Su", Mon: "Mo", Tue: "Tu", Wed: "We", Thu: "Th", Fri: "Fr", Sat: "Sa"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "Sunday", Mon: "Monday", Tue: "Tuesday", Wed: "Wednesday", Thu: "Thursday", Fri: "Friday", Sat: "Saturday"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "a", PM: "p"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "Andorran Peseta", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "United Arab Emirates Dirham", Symbol: ""},
				currency.AFA: cldr.Currency{DisplayName: "Afghan Afghani (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afghan Afghani", Symbol: ""},
				currency.ALK: cldr.Currency{DisplayName: "Albanian Lek (1946–1965)", Symbol: ""},
				currency.ALL: cldr.Currency{DisplayName: "Albanian Lek", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Armenian Dram", Symbol: ""},
				currency.ANG: cldr.Currency{DisplayName: "Netherlands Antillean Guilder", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Angolan Kwanza", Symbol: "Kz"},
				currency.AOK: cldr.Currency{DisplayName: "Angolan Kwanza (1977–1991)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "Angolan New Kwanza (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "Angolan Readjusted Kwanza (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "Argentine Austral", Symbol: ""},
				currency.ARL: cldr.Currency{DisplayName: "Argentine Peso Ley (1970–1983)", Symbol: ""},
				currency.ARM: cldr.Currency{DisplayName: "Argentine Peso (1881–1970)", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "Argentine Peso (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "Argentine Peso", Symbol: "$"},
				currency.ATS: cldr.Currency{DisplayName: "Austrian Schilling", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "Australian Dollar", Symbol: "$"},
				currency.AWG: cldr.Currency{DisplayName: "Aruban Florin", Symbol: ""},
				currency.AZM: cldr.Currency{DisplayName: "Azerbaijani Manat (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Azerbaijani Manat", Symbol: ""},
				currency.BAD: cldr.Currency{DisplayName: "Bosnia-Herzegovina Dinar (1992–1994)", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "Bosnia-Herzegovina Convertible Mark", Symbol: "KM"},
				currency.BAN: cldr.Currency{DisplayName: "Bosnia-Herzegovina New Dinar (1994–1997)", Symbol: ""},
				currency.BBD: cldr.Currency{DisplayName: "Barbadian Dollar", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Bangladeshi Taka", Symbol: "৳"},
				currency.BEC: cldr.Currency{DisplayName: "Belgian Franc (convertible)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "Belgian Franc", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "Belgian Franc (financial)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "Bulgarian Hard Lev", Symbol: ""},
				currency.BGM: cldr.Currency{DisplayName: "Bulgarian Socialist Lev", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "Bulgarian Lev", Symbol: ""},
				currency.BGO: cldr.Currency{DisplayName: "Bulgarian Lev (1879–1952)", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Bahraini Dinar", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Burundian Franc", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Bermudan Dollar", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Brunei Dollar", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Bolivian Boliviano", Symbol: "Bs"},
				currency.BOL: cldr.Currency{DisplayName: "Bolivian Boliviano (1863–1963)", Symbol: ""},
				currency.BOP: cldr.Currency{DisplayName: "Bolivian Peso", Symbol: ""},
				currency.BOV: cldr.Currency{DisplayName: "Bolivian Mvdol", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "Brazilian New Cruzeiro (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "Brazilian Cruzado (1986–1989)", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "Brazilian Cruzeiro (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "Brazilian Real", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "Brazilian New Cruzado (1989–1990)", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "Brazilian Cruzeiro (1993–1994)", Symbol: ""},
				currency.BRZ: cldr.Currency{DisplayName: "Brazilian Cruzeiro (1942–1967)", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "Bahamian Dollar", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Bhutanese Ngultrum", Symbol: ""},
				currency.BUK: cldr.Currency{DisplayName: "Burmese Kyat", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Botswanan Pula", Symbol: "P"},
				currency.BYB: cldr.Currency{DisplayName: "Belarusian Ruble (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "Belarusian Ruble", Symbol: "р."},
				currency.BYR: cldr.Currency{DisplayName: "Belarusian Ruble (2000–2016)", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Belize Dollar", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Canadian Dollar", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Congolese Franc", Symbol: ""},
				currency.CHE: cldr.Currency{DisplayName: "WIR Euro", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Swiss Franc", Symbol: ""},
				currency.CHW: cldr.Currency{DisplayName: "WIR Franc", Symbol: ""},
				currency.CLE: cldr.Currency{DisplayName: "Chilean Escudo", Symbol: ""},
				currency.CLF: cldr.Currency{DisplayName: "Chilean Unit of Account (UF)", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Chilean Peso", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Chinese Yuan (offshore)", Symbol: ""},
				currency.CNX: cldr.Currency{DisplayName: "Chinese People’s Bank Dollar", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Chinese Yuan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Colombian Peso", Symbol: "$"},
				currency.COU: cldr.Currency{DisplayName: "Colombian Real Value Unit", Symbol: ""},
				currency.CRC: cldr.Currency{DisplayName: "Costa Rican Colón", Symbol: "₡"},
				currency.CSD: cldr.Currency{DisplayName: "Serbian Dinar (2002–2006)", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "Czechoslovak Hard Koruna", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "Cuban Convertible Peso", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Cuban Peso", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Cape Verdean Escudo", Symbol: ""},
				currency.CYP: cldr.Currency{DisplayName: "Cypriot Pound", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Czech Koruna", Symbol: "Kč"},
				currency.DDM: cldr.Currency{DisplayName: "East German Mark", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "German Mark", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "Djiboutian Franc", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Danish Krone", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Dominican Peso", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Algerian Dinar", Symbol: ""},
				currency.ECS: cldr.Currency{DisplayName: "Ecuadorian Sucre", Symbol: ""},
				currency.ECV: cldr.Currency{DisplayName: "Ecuadorian Unit of Constant Value", Symbol: ""},
				currency.EEK: cldr.Currency{DisplayName: "Estonian Kroon", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Egyptian Pound", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Eritrean Nakfa", Symbol: ""},
				currency.ESA: cldr.Currency{DisplayName: "Spanish Peseta (A account)", Symbol: ""},
				currency.ESB: cldr.Currency{DisplayName: "Spanish Peseta (convertible account)", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "Spanish Peseta", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Ethiopian Birr", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "Euro", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "Finnish Markka", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "Fijian Dollar", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Falkland Islands Pound", Symbol: "£"},
				currency.FRF: cldr.Currency{DisplayName: "French Franc", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "British Pound", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "Georgian Kupon Larit", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "Georgian Lari", Symbol: "₾"},
				currency.GHC: cldr.Currency{DisplayName: "Ghanaian Cedi (1979–2007)", Symbol: ""},
				currency.GHS: cldr.Currency{DisplayName: "Ghanaian Cedi", Symbol: ""},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltar Pound", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Gambian Dalasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Guinean Franc", Symbol: "FG"},
				currency.GNS: cldr.Currency{DisplayName: "Guinean Syli", Symbol: ""},
				currency.GQE: cldr.Currency{DisplayName: "Equatorial Guinean Ekwele", Symbol: ""},
				currency.GRD: cldr.Currency{DisplayName: "Greek Drachma", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "Guatemalan Quetzal", Symbol: "Q"},
				currency.GWE: cldr.Currency{DisplayName: "Portuguese Guinea Escudo", Symbol: ""},
				currency.GWP: cldr.Currency{DisplayName: "Guinea-Bissau Peso", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "Guyanaese Dollar", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Hong Kong Dollar", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Honduran Lempira", Symbol: "L"},
				currency.HRD: cldr.Currency{DisplayName: "Croatian Dinar", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "Croatian Kuna", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Haitian Gourde", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Hungarian Forint", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Indonesian Rupiah", Symbol: "Rp"},
				currency.IEP: cldr.Currency{DisplayName: "Irish Pound", Symbol: ""},
				currency.ILP: cldr.Currency{DisplayName: "Israeli Pound", Symbol: ""},
				currency.ILR: cldr.Currency{DisplayName: "Israeli Shekel (1980–1985)", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "Israeli New Shekel", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Indian Rupee", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Iraqi Dinar", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Iranian Rial", Symbol: ""},
				currency.ISJ: cldr.Currency{DisplayName: "Icelandic Króna (1918–1981)", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Icelandic Króna", Symbol: "kr"},
				currency.ITL: cldr.Currency{DisplayName: "Italian Lira", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "Jamaican Dollar", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Jordanian Dinar", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Japanese Yen", Symbol: "¥"},
				currency.KES: cldr.Currency{DisplayName: "Kenyan Shilling", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Kyrgystani Som", Symbol: ""},
				currency.KHR: cldr.Currency{DisplayName: "Cambodian Riel", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Comorian Franc", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "North Korean Won", Symbol: "₩"},
				currency.KRH: cldr.Currency{DisplayName: "South Korean Hwan (1953–1962)", Symbol: ""},
				currency.KRO: cldr.Currency{DisplayName: "South Korean Won (1945–1953)", Symbol: ""},
				currency.KRW: cldr.Currency{DisplayName: "South Korean Won", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Kuwaiti Dinar", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Cayman Islands Dollar", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Kazakhstani Tenge", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Laotian Kip", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Lebanese Pound", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Sri Lankan Rupee", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Liberian Dollar", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lesotho Loti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "Lithuanian Litas", Symbol: "Lt"},
				currency.LTT: cldr.Currency{DisplayName: "Lithuanian Talonas", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "Luxembourgian Convertible Franc", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "Luxembourgian Franc", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "Luxembourg Financial Franc", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "Latvian Lats", Symbol: "Ls"},
				currency.LVR: cldr.Currency{DisplayName: "Latvian Ruble", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "Libyan Dinar", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Moroccan Dirham", Symbol: ""},
				currency.MAF: cldr.Currency{DisplayName: "Moroccan Franc", Symbol: ""},
				currency.MCF: cldr.Currency{DisplayName: "Monegasque Franc", Symbol: ""},
				currency.MDC: cldr.Currency{DisplayName: "Moldovan Cupon", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Moldovan Leu", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Malagasy Ariary", Symbol: "Ar"},
				currency.MGF: cldr.Currency{DisplayName: "Malagasy Franc", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "Macedonian Denar", Symbol: ""},
				currency.MKN: cldr.Currency{DisplayName: "Macedonian Denar (1992–1993)", Symbol: ""},
				currency.MLF: cldr.Currency{DisplayName: "Malian Franc", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Myanmar Kyat", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Mongolian Tugrik", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Macanese Pataca", Symbol: ""},
				currency.MRO: cldr.Currency{DisplayName: "Mauritanian Ouguiya (1973–2017)", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Mauritanian Ouguiya", Symbol: ""},
				currency.MTL: cldr.Currency{DisplayName: "Maltese Lira", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "Maltese Pound", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Mauritian Rupee", Symbol: "Rs"},
				currency.MVP: cldr.Currency{DisplayName: "Maldivian Rupee (1947–1981)", Symbol: ""},
				currency.MVR: cldr.Currency{DisplayName: "Maldivian Rufiyaa", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Malawian Kwacha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Mexican Peso", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "Mexican Silver Peso (1861–1992)", Symbol: ""},
				currency.MXV: cldr.Currency{DisplayName: "Mexican Investment Unit", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "Malaysian Ringgit", Symbol: "RM"},
				currency.MZE: cldr.Currency{DisplayName: "Mozambican Escudo", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "Mozambican Metical (1980–2006)", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "Mozambican Metical", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Namibian Dollar", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nigerian Naira", Symbol: "₦"},
				currency.NIC: cldr.Currency{DisplayName: "Nicaraguan Córdoba (1988–1991)", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "Nicaraguan Córdoba", Symbol: "C$"},
				currency.NLG: cldr.Currency{DisplayName: "Dutch Guilder", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "Norwegian Krone", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Nepalese Rupee", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "New Zealand Dollar", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Omani Rial", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Panamanian Balboa", Symbol: ""},
				currency.PEI: cldr.Currency{DisplayName: "Peruvian Inti", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Peruvian Sol", Symbol: ""},
				currency.PES: cldr.Currency{DisplayName: "Peruvian Sol (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Papua New Guinean Kina", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Philippine Piso", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistani Rupee", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Polish Zloty", Symbol: "zł"},
				currency.PLZ: cldr.Currency{DisplayName: "Polish Zloty (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "Portuguese Escudo", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "Paraguayan Guarani", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Qatari Rial", Symbol: ""},
				currency.RHD: cldr.Currency{DisplayName: "Rhodesian Dollar", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "Romanian Leu (1952–2006)", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Romanian Leu", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Serbian Dinar", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Russian Ruble", Symbol: "₽"},
				currency.RUR: cldr.Currency{DisplayName: "Russian Ruble (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "Rwandan Franc", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Saudi Riyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Solomon Islands Dollar", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Seychellois Rupee", Symbol: ""},
				currency.SDD: cldr.Currency{DisplayName: "Sudanese Dinar (1992–2007)", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Sudanese Pound", Symbol: ""},
				currency.SDP: cldr.Currency{DisplayName: "Sudanese Pound (1957–1998)", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Swedish Krona", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Singapore Dollar", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "St. Helena Pound", Symbol: "£"},
				currency.SIT: cldr.Currency{DisplayName: "Slovenian Tolar", Symbol: ""},
				currency.SKK: cldr.Currency{DisplayName: "Slovak Koruna", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Sierra Leonean Leone", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Somali Shilling", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Surinamese Dollar", Symbol: "$"},
				currency.SRG: cldr.Currency{DisplayName: "Surinamese Guilder", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "South Sudanese Pound", Symbol: "£"},
				currency.STD: cldr.Currency{DisplayName: "São Tomé & Príncipe Dobra (1977–2017)", Symbol: ""},
				currency.STN: cldr.Currency{DisplayName: "São Tomé & Príncipe Dobra", Symbol: "Db"},
				currency.SUR: cldr.Currency{DisplayName: "Soviet Rouble", Symbol: ""},
				currency.SVC: cldr.Currency{DisplayName: "Salvadoran Colón", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "Syrian Pound", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Swazi Lilangeni", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Thai Baht", Symbol: "฿"},
				currency.TJR: cldr.Currency{DisplayName: "Tajikistani Ruble", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "Tajikistani Somoni", Symbol: ""},
				currency.TMM: cldr.Currency{DisplayName: "Turkmenistani Manat (1993–2009)", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Turkmenistani Manat", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Tunisian Dinar", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Tongan Paʻanga", Symbol: "T$"},
				currency.TPE: cldr.Currency{DisplayName: "Timorese Escudo", Symbol: ""},
				currency.TRL: cldr.Currency{DisplayName: "Turkish Lira (1922–2005)", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "Turkish Lira", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Trinidad & Tobago Dollar", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "New Taiwan Dollar", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Tanzanian Shilling", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Ukrainian Hryvnia", Symbol: "₴"},
				currency.UAK: cldr.Currency{DisplayName: "Ukrainian Karbovanets", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "Ugandan Shilling (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "Ugandan Shilling", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "US Dollar", Symbol: "$"},
				currency.USN: cldr.Currency{DisplayName: "US Dollar (Next day)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "US Dollar (Same day)", Symbol: ""},
				currency.UYI: cldr.Currency{DisplayName: "Uruguayan Peso (Indexed Units)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "Uruguayan Peso (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "Uruguayan Peso", Symbol: "$"},
				currency.UYW: cldr.Currency{DisplayName: "Uruguayan Nominal Wage Index Unit", Symbol: ""},
				currency.UZS: cldr.Currency{DisplayName: "Uzbekistani Som", Symbol: ""},
				currency.VEB: cldr.Currency{DisplayName: "Venezuelan Bolívar (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "Venezuelan Bolívar (2008–2018)", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Venezuelan Bolívar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Vietnamese Dong", Symbol: "₫"},
				currency.VNN: cldr.Currency{DisplayName: "Vietnamese Dong (1978–1985)", Symbol: ""},
				currency.VUV: cldr.Currency{DisplayName: "Vanuatu Vatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Samoan Tala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Central African CFA Franc", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "Silver", Symbol: ""},
				currency.XAU: cldr.Currency{DisplayName: "Gold", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "European Composite Unit", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "European Monetary Unit", Symbol: ""},
				currency.XBC: cldr.Currency{DisplayName: "European Unit of Account (XBC)", Symbol: ""},
				currency.XBD: cldr.Currency{DisplayName: "European Unit of Account (XBD)", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "East Caribbean Dollar", Symbol: "EC$"},
				currency.XDR: cldr.Currency{DisplayName: "Special Drawing Rights", Symbol: ""},
				currency.XEU: cldr.Currency{DisplayName: "European Currency Unit", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "French Gold Franc", Symbol: ""},
				currency.XFU: cldr.Currency{DisplayName: "French UIC-Franc", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "West African CFA Franc", Symbol: "CFA"},
				currency.XPD: cldr.Currency{DisplayName: "Palladium", Symbol: ""},
				currency.XPF: cldr.Currency{DisplayName: "CFP Franc", Symbol: "CFPF"},
				currency.XPT: cldr.Currency{DisplayName: "Platinum", Symbol: ""},
				currency.XRE: cldr.Currency{DisplayName: "RINET Funds", Symbol: ""},
				currency.XSU: cldr.Currency{DisplayName: "Sucre", Symbol: ""},
				currency.XTS: cldr.Currency{DisplayName: "Testing Currency Code", Symbol: ""},
				currency.XUA: cldr.Currency{DisplayName: "ADB Unit of Account", Symbol: ""},
				currency.XXX: cldr.Currency{DisplayName: "Unknown Currency", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "Yemeni Dinar", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "Yemeni Rial", Symbol: ""},
				currency.YUD: cldr.Currency{DisplayName: "Yugoslavian Hard Dinar (1966–1990)", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "Yugoslavian New Dinar (1994–2002)", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "Yugoslavian Convertible Dinar (1990–1992)", Symbol: ""},
				currency.YUR: cldr.Currency{DisplayName: "Yugoslavian Reformed Dinar (1992–1993)", Symbol: ""},
				currency.ZAL: cldr.Currency{DisplayName: "South African Rand (financial)", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "South African Rand", Symbol: "R"},
				currency.ZMK: cldr.Currency{DisplayName: "Zambian Kwacha (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "Zambian Kwacha", Symbol: "ZK"},
				currency.ZRN: cldr.Currency{DisplayName: "Zairean New Zaire (1993–1998)", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "Zairean Zaire (1971–1993)", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "Zimbabwean Dollar (1980–2008)", Symbol: ""},
				currency.ZWL: cldr.Currency{DisplayName: "Zimbabwean Dollar (2009)", Symbol: ""},
				currency.ZWR: cldr.Currency{DisplayName: "Zimbabwean Dollar (2008)", Symbol: ""},
			},
		},
	}
}
