// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/localizer/resources/currency"
)

func getLocale_gu() *cldr.Locale {
	return &cldr.Locale{
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, Time: cldr.CalendarDateFormat{Full: "hh:mm:ss a zzzz", Long: "hh:mm:ss a z", Medium: "hh:mm:ss a", Short: "hh:mm a"}, DateTime: cldr.CalendarDateFormat{Full: "{1} એ {0} વાગ્યે", Long: "{1} એ {0} વાગ્યે", Medium: "{1} {0}", Short: "{1} {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "જાન્યુ", Feb: "ફેબ્રુ", Mar: "માર્ચ", Apr: "એપ્રિલ", May: "મે", Jun: "જૂન", Jul: "જુલાઈ", Aug: "ઑગસ્ટ", Sep: "સપ્ટે", Oct: "ઑક્ટો", Nov: "નવે", Dec: "ડિસે"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "જા", Feb: "ફે", Mar: "મા", Apr: "એ", May: "મે", Jun: "જૂ", Jul: "જુ", Aug: "ઑ", Sep: "સ", Oct: "ઑ", Nov: "ન", Dec: "ડિ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "જાન્યુઆરી", Feb: "ફેબ્રુઆરી", Mar: "માર્ચ", Apr: "એપ્રિલ", May: "મે", Jun: "જૂન", Jul: "જુલાઈ", Aug: "ઑગસ્ટ", Sep: "સપ્ટેમ્બર", Oct: "ઑક્ટોબર", Nov: "નવેમ્બર", Dec: "ડિસેમ્બર"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "રવિ", Mon: "સોમ", Tue: "મંગળ", Wed: "બુધ", Thu: "ગુરુ", Fri: "શુક્ર", Sat: "શનિ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "ર", Mon: "સો", Tue: "મં", Wed: "બુ", Thu: "ગુ", Fri: "શુ", Sat: "શ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "ર", Mon: "સો", Tue: "મં", Wed: "બુ", Thu: "ગુ", Fri: "શુ", Sat: "શ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "રવિવાર", Mon: "સોમવાર", Tue: "મંગળવાર", Wed: "બુધવાર", Thu: "ગુરુવાર", Fri: "શુક્રવાર", Sat: "શનિવાર"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Plural:   cldr.Plural{Cardinal: LocalePlural[tag_gu]()},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: "٫", Group: "", Negative: "", Percent: "", PerMille: ""},
			Formats: cldr.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤#,##0.00", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "યુનાઈટેડ આરબ અમિરાત દિરહામ", Symbol: "AED"},
				currency.AFN: cldr.Currency{DisplayName: "અફ્ગાન અફ્ગાની", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "અલ્બેનિયન લેક", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "અર્મેનિયન ડ્રેમ", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "નેધરલેંડ એંટિલિન ગિલ્ડર", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "અંગોલિયન ક્વાન્ઝા", Symbol: "AOA"},
				currency.ARS: cldr.Currency{DisplayName: "અર્જેન્ટીના પેસો", Symbol: "ARS"},
				currency.AUD: cldr.Currency{DisplayName: "ઑસ્ટ્રેલિયન ડૉલર", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "અરુબન ફ્લોરિન", Symbol: "AWG"},
				currency.AZN: cldr.Currency{DisplayName: "અઝરબૈજાની મનાત", Symbol: "AZN"},
				currency.BAM: cldr.Currency{DisplayName: "બોસ્નિયા અને હર્ઝેગોવિના રૂપાંતર યોગ્ય માર્ક", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "બાર્બાડિયન ડોલર", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "બાંગ્લાદેશી ટાકા", Symbol: "BDT"},
				currency.BGN: cldr.Currency{DisplayName: "બલ્ગેરીયન લેવ", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "બેહરિની દિનાર", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "બુરુન્ડિયન ફ્રેંક", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "બર્મુડન ડોલર", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "બ્રુનેઇ ડોલર", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "બોલિવિયન બોલિવિયાનો", Symbol: "BOB"},
				currency.BRL: cldr.Currency{DisplayName: "બ્રાઝિલીયન રિઆલ", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "બહામિયન ડોલર", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ભુતાનિઝ એંગુલ્ત્રમ", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "બોત્સવાનન પુલા", Symbol: "BWP"},
				currency.BYN: cldr.Currency{DisplayName: "બેલારુશિયન રૂબલ", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "બેલારુશિયન રૂબલ (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "બેલિઝ ડોલર", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "કેનેડિયન ડૉલર", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "કોંગોલિઝ ફ્રેંક", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "સ્વિસ ફ્રેંક", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "ચિલિઅન પેસો", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "ચાઇનીઝ યુઆન (ઑફશોર)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "ચાઇનિઝ યુઆન", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "કોલમ્બિયન પેસો", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "કોસ્ટા રિકન કોલોન", Symbol: "CRC"},
				currency.CUC: cldr.Currency{DisplayName: "ક્યુબન રૂપાંતર યોગ્ય પેસો", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "ક્યુબન પેસો", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "કેપ વર્દિયન એસ્કુડો", Symbol: "CVE"},
				currency.CZK: cldr.Currency{DisplayName: "ચેક રીપબ્લિક કોરુના", Symbol: "CZK"},
				currency.DJF: cldr.Currency{DisplayName: "જિબુટિયન ફ્રેંક", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "ડેનિશ ક્રોન", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "ડોમિનિકન પેસો", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "અલ્જિરિયન દિનાર", Symbol: "DZD"},
				currency.EGP: cldr.Currency{DisplayName: "ઇજિપ્તિયન પાઉન્ડ", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "એરિટ્રેયન નક્ફા", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ઇથિયોપીયન બિર", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "યુરો", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "ફિજિઅન ડોલર", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ફૉકલેન્ડ આઇલેંડ્સ પાઉન્ડ", Symbol: "FKP"},
				currency.GBP: cldr.Currency{DisplayName: "બ્રિટિશ પાઉન્ડ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "જ્યોર્જિઅન લારી", Symbol: "GEL"},
				currency.GHS: cldr.Currency{DisplayName: "ઘાનાઇયન સેડી", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "જીબ્રાલ્ટર પાઉન્ડ", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "ગેમ્બિયન દલાસી", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "ગિનીયન ફ્રેંક", Symbol: "GNF"},
				currency.GTQ: cldr.Currency{DisplayName: "ગ્વાટેમાલા કુઇટ્ઝલ", Symbol: "GTQ"},
				currency.GYD: cldr.Currency{DisplayName: "ગયાનિઝ ડોલર", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "હોંગ કોંગ ડૉલર", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "હોન્ડ્યુરન લેમ્પિરા", Symbol: "HNL"},
				currency.HRK: cldr.Currency{DisplayName: "ક્રોએશિયન ક્યુના", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "હાઇટિઇન ગોર્ડ", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "હંગેરીયન ફોરિન્ત", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "ઇન્ડોનેશિયન રૂપિયા", Symbol: "IDR"},
				currency.ILS: cldr.Currency{DisplayName: "ઇઝરાયેલી ન્યુ શેકલ", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "ભારતીય રૂપિયા", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "ઇરાકી દિનાર", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ઇરાનિયન રિયાલ", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "આઇસલેન્ડિક ક્રોના", Symbol: "ISK"},
				currency.JMD: cldr.Currency{DisplayName: "જમૈકિયન ડોલર", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "જોર્ડનિયન દિનાર", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "જાપાનીઝ યેન", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "કેન્યેન શિલિંગ", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "કિર્ગિસ્તાની સોમ", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "કેમ્બોડિયન રીઅલ", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "કોમોરિઅન ફ્રેંક", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "ઉત્તર કોરિયન વન", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "દક્ષિણ કોરિયન વન", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "કુવૈતી દિનાર", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "કેયમેન આઇલેંડ્સ ડોલર", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "કઝાકિસ્તાની ટેંગ", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "લાઓશિયન કિપ", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "લેબેનિઝ પાઉન્ડ", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "શ્રી લંકન રૂપી", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "લિબેરિયન ડોલર", Symbol: "LRD"},
				currency.LSL: cldr.Currency{DisplayName: "લેસોથો લોતી", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "લિથુએનિયન લિતાસ", Symbol: "LTL"},
				currency.LVL: cldr.Currency{DisplayName: "લાતવિયન લેત્સ", Symbol: "LVL"},
				currency.LYD: cldr.Currency{DisplayName: "લિબ્યન દિનાર", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "મોરોક્કન દિરહામ", Symbol: "MAD"},
				currency.MDL: cldr.Currency{DisplayName: "મોલડોવેન લિયુ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "માલાગેસી અરીઆરી", Symbol: "MGA"},
				currency.MKD: cldr.Currency{DisplayName: "મેસેડોનિયન દિનાર", Symbol: "MKD"},
				currency.MMK: cldr.Currency{DisplayName: "મ્યાંમાર ક્યાત", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "મોંગોલિયન ટગરિક", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "માકાનિઝ પતાકા", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "મોરીશેનિયન ઓગુયા (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "મોરીશેનિયન ઓગુયા", Symbol: "MRU"},
				currency.MUR: cldr.Currency{DisplayName: "મોરેશીઅન રૂપી", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "માલ્દિવિયન રુફિયા", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "માલાવિયન ક્વાચા", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "મેક્સિકન પેસો", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "મલેશિયન રિંગ્ગેટ", Symbol: "MYR"},
				currency.MZN: cldr.Currency{DisplayName: "મોઝામ્બિકન મેટિકલ", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "નામિબિયન ડોલર", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "નાઇજીરિયન નૈરા", Symbol: "NGN"},
				currency.NIO: cldr.Currency{DisplayName: "નિકારાગુઅન કોર્ડોબા", Symbol: "NIO"},
				currency.NOK: cldr.Currency{DisplayName: "નૉર્વેજિયન ક્રોન", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "નેપાલિઝ રૂપી", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "ન્યૂઝિલેંડ ડૉલર", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "ઓમાની રિયાલ", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "પનામેનિયન બાલ્બોઆ", Symbol: "PAB"},
				currency.PEN: cldr.Currency{DisplayName: "પેરુવિયન સોલ", Symbol: "PEN"},
				currency.PGK: cldr.Currency{DisplayName: "પાપુઆ ન્યૂ ગિનીયન કિના", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ફિલીપાઇન પેસો", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "પાકિસ્તાની રૂપી", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "પોલિસ ઝ્લોટી", Symbol: "PLN"},
				currency.PYG: cldr.Currency{DisplayName: "પરાગ્વેયન ગુઆરાની", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "કતારી રિયાલ", Symbol: "QAR"},
				currency.RON: cldr.Currency{DisplayName: "રોમાનિયન લેઉ", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "સર્બિયન દિનાર", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "રશિયન રૂબલ", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "રવાંડન ફ્રેંક", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "સાઉદી રિયાલ", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "સોલોમન આઇલેંડ્સ ડોલર", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "સેશેલોઈ રૂપી", Symbol: "SCR"},
				currency.SDG: cldr.Currency{DisplayName: "સુદાનિઝ પાઉન્ડ", Symbol: "SDG"},
				currency.SEK: cldr.Currency{DisplayName: "સ્વીડિશ ક્રોના", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "સિંગાપુર ડૉલર", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "સેંટ હેલેના પાઉન્ડ", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "સિએરા લિઓનિઅન લિઓન", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "સોમાલી શિલિંગ", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "સૂરીનામિઝ ડોલર", Symbol: "SRD"},
				currency.SSP: cldr.Currency{DisplayName: "દક્ષિણ સુદાનિઝ પાઉન્ડ", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "સાઓ ટૉમ એન્ડ પ્રિંસાઇપ ડોબ્રા (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "સાઓ ટૉમ એન્ડ પ્રિંસાઇપ ડોબ્રા", Symbol: "STN"},
				currency.SYP: cldr.Currency{DisplayName: "સાઇરિયન પાઉન્ડ", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "સ્વાઝી લિલાન્ગેની", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "થાઇ બાહ્ત", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "તાજિકિસ્તાની સોમોની", Symbol: "TJS"},
				currency.TMT: cldr.Currency{DisplayName: "તુર્કમેનિસ્તાની મનત", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "ટ્યુનિશિયન દિનાર", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "ટોંગન પ’અંગા", Symbol: "TOP"},
				currency.TRY: cldr.Currency{DisplayName: "તુર્કિશ લિરા", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ત્રિનિદાદ અને ટોબેગો ડોલર", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "ન્યુ તાઇવાન ડૉલર", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "તાન્ઝાનિયન શિલિંગ", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "યુક્રેનિયન હ્રિવિનિયા", Symbol: "UAH"},
				currency.UGX: cldr.Currency{DisplayName: "યુગાંડન શિલિંગ", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "યુઍસ ડોલર", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "ઉરુગ્વેયન પેસો", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "ઉઝ્બેકિસ્તાન સોમ", Symbol: "UZS"},
				currency.VEF: cldr.Currency{DisplayName: "વેનેઝુએલન બોલિવર (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "વેનેઝુએલન બોલિવર", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "વિયેતનામીસ ડોંગ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "વનૌતુ વાતુ", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "સમોઅન તાલા", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "મધ્ય આફ્રિકન [CFA] ફ્રેંક", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "ઇસ્ટ કેરિબિયન ડોલર", Symbol: "EC$"},
				currency.XOF: cldr.Currency{DisplayName: "પશ્ચિમી આફ્રિકન [CFA] ફ્રેંક", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "[CFP] ફ્રેંક", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "અજ્ઞાત ચલણ", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "યેમેની રિઆલ", Symbol: "YER"},
				currency.ZAR: cldr.Currency{DisplayName: "દક્ષિણ આફ્રિકી રેંડ", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "ઝામ્બિયન ક્વાચા (1968–2012)", Symbol: ""},
				currency.ZMW: cldr.Currency{DisplayName: "ઝામ્બિયન ક્વાચા", Symbol: "ZMW"},
			},
		},
	}
}
