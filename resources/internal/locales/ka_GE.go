// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
)

func Get_ka_GE() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "ka_GE",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, dd MMMM, y", Long: "d MMMM, y", Medium: "d MMM. y", Short: "dd.MM.yy"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "GMT{0}"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "იან", Feb: "თებ", Mar: "მარ", Apr: "აპრ", May: "მაი", Jun: "ივნ", Jul: "ივლ", Aug: "აგვ", Sep: "სექ", Oct: "ოქტ", Nov: "ნოე", Dec: "დეკ"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "ი", Feb: "თ", Mar: "მ", Apr: "ა", May: "მ", Jun: "ი", Jul: "ი", Aug: "ა", Sep: "ს", Oct: "ო", Nov: "ნ", Dec: "დ"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "იანვარი", Feb: "თებერვალი", Mar: "მარტი", Apr: "აპრილი", May: "მაისი", Jun: "ივნისი", Jul: "ივლისი", Aug: "აგვისტო", Sep: "სექტემბერი", Oct: "ოქტომბერი", Nov: "ნოემბერი", Dec: "დეკემბერი"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "კვი", Mon: "ორშ", Tue: "სამ", Wed: "ოთხ", Thu: "ხუთ", Fri: "პარ", Sat: "შაბ"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "კ", Mon: "ო", Tue: "ს", Wed: "ო", Thu: "ხ", Fri: "პ", Sat: "შ"}, Short: cldr.CalendarDayFormatNameValue{Sun: "კვ", Mon: "ორ", Tue: "სმ", Wed: "ოთ", Thu: "ხთ", Fri: "პრ", Sat: "შბ"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "კვირა", Mon: "ორშაბათი", Tue: "სამშაბათი", Wed: "ოთხშაბათი", Thu: "ხუთშაბათი", Fri: "პარასკევი", Sat: "შაბათი"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "AM", PM: "PM"}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: "#,##0%"},
			Currencies: cldr.Currencies{
				currency.ADP: cldr.Currency{DisplayName: "ანდორული პესეტა", Symbol: ""},
				currency.AED: cldr.Currency{DisplayName: "არაბთა გაერთიანებული საამიროების დირჰამი", Symbol: "AED"},
				currency.AFA: cldr.Currency{DisplayName: "ავღანი (1927–2002)", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "ავღანური ავღანი", Symbol: "AFN"},
				currency.ALL: cldr.Currency{DisplayName: "ალბანური ლეკი", Symbol: "ALL"},
				currency.AMD: cldr.Currency{DisplayName: "სომხური დრამი", Symbol: "AMD"},
				currency.ANG: cldr.Currency{DisplayName: "ნიდერლანდების ანტილების გულდენი", Symbol: "ANG"},
				currency.AOA: cldr.Currency{DisplayName: "ანგოლური კვანზა", Symbol: "AOA"},
				currency.AOK: cldr.Currency{DisplayName: "ანგოლური კვანზა (1977–1990)", Symbol: ""},
				currency.AON: cldr.Currency{DisplayName: "ანგოლური ახალი კვანზა (1990–2000)", Symbol: ""},
				currency.AOR: cldr.Currency{DisplayName: "ანგოლური მიტოლებული კვანზა (1995–1999)", Symbol: ""},
				currency.ARA: cldr.Currency{DisplayName: "არგენტინული აუსტრალი", Symbol: ""},
				currency.ARP: cldr.Currency{DisplayName: "არგენტინული პესო (1983–1985)", Symbol: ""},
				currency.ARS: cldr.Currency{DisplayName: "არგენტინული პესო", Symbol: "ARS"},
				currency.ATS: cldr.Currency{DisplayName: "ავსტრიული შილინგი", Symbol: ""},
				currency.AUD: cldr.Currency{DisplayName: "ავსტრალიური დოლარი", Symbol: "AUD"},
				currency.AWG: cldr.Currency{DisplayName: "არუბანული გულდენი", Symbol: "AWG"},
				currency.AZM: cldr.Currency{DisplayName: "აზერბაიჯანული მანათი (1993–2006)", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "აზერბაიჯანული მანათი", Symbol: "AZN"},
				currency.BAD: cldr.Currency{DisplayName: "ბოსნია-ჰერცოგოვინას დინარი", Symbol: ""},
				currency.BAM: cldr.Currency{DisplayName: "ბოსნია და ჰერცოგოვინას კონვერტირებადი მარკა", Symbol: "BAM"},
				currency.BBD: cldr.Currency{DisplayName: "ბარბადოსული დოლარი", Symbol: "BBD"},
				currency.BDT: cldr.Currency{DisplayName: "ბანგლადეშური ტაკა", Symbol: "BDT"},
				currency.BEC: cldr.Currency{DisplayName: "ბელგიური ფრანკი (კოვერტირებადი)", Symbol: ""},
				currency.BEF: cldr.Currency{DisplayName: "ბელგიური ფრანკი", Symbol: ""},
				currency.BEL: cldr.Currency{DisplayName: "ბელგიური ფრანკი (ფინანსური)", Symbol: ""},
				currency.BGL: cldr.Currency{DisplayName: "ბულგარული მყარი ლევი", Symbol: ""},
				currency.BGN: cldr.Currency{DisplayName: "ბულგარული ლევი", Symbol: "BGN"},
				currency.BHD: cldr.Currency{DisplayName: "ბაჰრეინული დინარი", Symbol: "BHD"},
				currency.BIF: cldr.Currency{DisplayName: "ბურუნდიული ფრანკი", Symbol: "BIF"},
				currency.BMD: cldr.Currency{DisplayName: "ბერმუდული დოლარი", Symbol: "BMD"},
				currency.BND: cldr.Currency{DisplayName: "ბრუნეული დოლარი", Symbol: "BND"},
				currency.BOB: cldr.Currency{DisplayName: "ბოლივიური ბოლივიანო", Symbol: "BOB"},
				currency.BOP: cldr.Currency{DisplayName: "ბოლივიური პესო", Symbol: ""},
				currency.BRB: cldr.Currency{DisplayName: "ბრაზილიური კრუზეირო ნოვო (1967–1986)", Symbol: ""},
				currency.BRC: cldr.Currency{DisplayName: "ბრაზილიური კრუზადო", Symbol: ""},
				currency.BRE: cldr.Currency{DisplayName: "ბრაზილიური კრუზეირო (1990–1993)", Symbol: ""},
				currency.BRL: cldr.Currency{DisplayName: "ბრაზილიური რეალი", Symbol: "R$"},
				currency.BRN: cldr.Currency{DisplayName: "ბრაზილიური კრუზადო ნოვო", Symbol: ""},
				currency.BRR: cldr.Currency{DisplayName: "ბრაზილიური კრუზეირო", Symbol: ""},
				currency.BSD: cldr.Currency{DisplayName: "ბაჰამური დოლარი", Symbol: "BSD"},
				currency.BTN: cldr.Currency{DisplayName: "ბუტანური ნგულტრუმი", Symbol: "BTN"},
				currency.BWP: cldr.Currency{DisplayName: "ბოცვანური პულა", Symbol: "BWP"},
				currency.BYB: cldr.Currency{DisplayName: "ახალი ბელარუსიული რუბლი (1994–1999)", Symbol: ""},
				currency.BYN: cldr.Currency{DisplayName: "ბელორუსული რუბლი", Symbol: "BYN"},
				currency.BYR: cldr.Currency{DisplayName: "ბელორუსული რუბლი (2000–2016)", Symbol: "BYR"},
				currency.BZD: cldr.Currency{DisplayName: "ბელიზის დოლარი", Symbol: "BZD"},
				currency.CAD: cldr.Currency{DisplayName: "კანადური დოლარი", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "კონგოს ფრანკი", Symbol: "CDF"},
				currency.CHF: cldr.Currency{DisplayName: "შვეიცარიული ფრანკი", Symbol: "CHF"},
				currency.CLP: cldr.Currency{DisplayName: "ჩილეს პესო", Symbol: "CLP"},
				currency.CNH: cldr.Currency{DisplayName: "ჩინური იუანი (ოფშორი)", Symbol: "CNH"},
				currency.CNY: cldr.Currency{DisplayName: "ჩინური იუანი", Symbol: "CNY"},
				currency.COP: cldr.Currency{DisplayName: "კოლუმბიური პესო", Symbol: "COP"},
				currency.CRC: cldr.Currency{DisplayName: "კოსტა-რიკული კოლონი", Symbol: "CRC"},
				currency.CSD: cldr.Currency{DisplayName: "ძველი სერბიული დინარი", Symbol: ""},
				currency.CSK: cldr.Currency{DisplayName: "ჩეხოსლოვაკიის მყარი კრონა", Symbol: ""},
				currency.CUC: cldr.Currency{DisplayName: "კუბური კონვერტირებადი პესო", Symbol: "CUC"},
				currency.CUP: cldr.Currency{DisplayName: "კუბური პესო", Symbol: "CUP"},
				currency.CVE: cldr.Currency{DisplayName: "კაბო-ვერდეს ესკუდო", Symbol: "CVE"},
				currency.CYP: cldr.Currency{DisplayName: "კვიპროსის გირვანქა", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "ჩეხური კრონა", Symbol: "CZK"},
				currency.DDM: cldr.Currency{DisplayName: "აღმოსავლეთ გერმანული მარკა", Symbol: ""},
				currency.DEM: cldr.Currency{DisplayName: "გერმანული მარკა", Symbol: ""},
				currency.DJF: cldr.Currency{DisplayName: "ჯიბუტის ფრანკი", Symbol: "DJF"},
				currency.DKK: cldr.Currency{DisplayName: "დანიური კრონა", Symbol: "DKK"},
				currency.DOP: cldr.Currency{DisplayName: "დომინიკური პესო", Symbol: "DOP"},
				currency.DZD: cldr.Currency{DisplayName: "ალჟირული დინარი", Symbol: "DZD"},
				currency.EEK: cldr.Currency{DisplayName: "ესტონური კრუნა", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "ეგვიპტური გირვანქა", Symbol: "EGP"},
				currency.ERN: cldr.Currency{DisplayName: "ერიტრეის ნაკფა", Symbol: "ERN"},
				currency.ESP: cldr.Currency{DisplayName: "ესპანური პესეტა", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "ეთიოპიური ბირი", Symbol: "ETB"},
				currency.EUR: cldr.Currency{DisplayName: "ევრო", Symbol: "€"},
				currency.FIM: cldr.Currency{DisplayName: "ფინური მარკა", Symbol: ""},
				currency.FJD: cldr.Currency{DisplayName: "ფიჯის დოლარი", Symbol: "FJD"},
				currency.FKP: cldr.Currency{DisplayName: "ფოლკლენდის კუნძულების ფუნტი", Symbol: "FKP"},
				currency.FRF: cldr.Currency{DisplayName: "ფრანგული ფრანკი", Symbol: ""},
				currency.GBP: cldr.Currency{DisplayName: "ბრიტანული გირვანქა სტერლინგი", Symbol: "£"},
				currency.GEK: cldr.Currency{DisplayName: "ქართული კუპონი ლარით", Symbol: ""},
				currency.GEL: cldr.Currency{DisplayName: "ქართული ლარი", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "განური სედი", Symbol: "GHS"},
				currency.GIP: cldr.Currency{DisplayName: "გიბრალტარული ფუნტი", Symbol: "GIP"},
				currency.GMD: cldr.Currency{DisplayName: "გამბიური დალასი", Symbol: "GMD"},
				currency.GNF: cldr.Currency{DisplayName: "გვინეური ფრანკი", Symbol: "GNF"},
				currency.GRD: cldr.Currency{DisplayName: "ბერძნული დრაჰმა", Symbol: ""},
				currency.GTQ: cldr.Currency{DisplayName: "გვატემალური კეტსალი", Symbol: "GTQ"},
				currency.GWE: cldr.Currency{DisplayName: "პორტუგალიური გინეა ესკუდო", Symbol: ""},
				currency.GYD: cldr.Currency{DisplayName: "გაიანური დოლარი", Symbol: "GYD"},
				currency.HKD: cldr.Currency{DisplayName: "ჰონკონგის დოლარი", Symbol: "HKD"},
				currency.HNL: cldr.Currency{DisplayName: "ჰონდურასული ლემპირა", Symbol: "HNL"},
				currency.HRD: cldr.Currency{DisplayName: "ხორვატიული დინარი", Symbol: ""},
				currency.HRK: cldr.Currency{DisplayName: "ხორვატული კუნა", Symbol: "HRK"},
				currency.HTG: cldr.Currency{DisplayName: "ჰაიტური გურდი", Symbol: "HTG"},
				currency.HUF: cldr.Currency{DisplayName: "უნგრული ფორინტი", Symbol: "HUF"},
				currency.IDR: cldr.Currency{DisplayName: "ინდონეზიური რუპია", Symbol: "IDR"},
				currency.IEP: cldr.Currency{DisplayName: "ირლანდიური გირვანქა", Symbol: ""},
				currency.ILS: cldr.Currency{DisplayName: "ისრაელის ახალი შეკელი", Symbol: "ILS"},
				currency.INR: cldr.Currency{DisplayName: "ინდური რუპია", Symbol: "INR"},
				currency.IQD: cldr.Currency{DisplayName: "ერაყული დინარი", Symbol: "IQD"},
				currency.IRR: cldr.Currency{DisplayName: "ირანული რიალი", Symbol: "IRR"},
				currency.ISK: cldr.Currency{DisplayName: "ისლანდიური კრონა", Symbol: "ISK"},
				currency.ITL: cldr.Currency{DisplayName: "იტალიური ლირა", Symbol: ""},
				currency.JMD: cldr.Currency{DisplayName: "იამაიკური დოლარი", Symbol: "JMD"},
				currency.JOD: cldr.Currency{DisplayName: "იორდანიული დოლარი", Symbol: "JOD"},
				currency.JPY: cldr.Currency{DisplayName: "იაპონური იენი", Symbol: "JPY"},
				currency.KES: cldr.Currency{DisplayName: "კენიური შილინგი", Symbol: "KES"},
				currency.KGS: cldr.Currency{DisplayName: "ყირგიზული სომი", Symbol: "KGS"},
				currency.KHR: cldr.Currency{DisplayName: "კამბოჯური რიელი", Symbol: "KHR"},
				currency.KMF: cldr.Currency{DisplayName: "კომორული ფრანკი", Symbol: "KMF"},
				currency.KPW: cldr.Currency{DisplayName: "ჩრდილოეთ კორეული ვონი", Symbol: "KPW"},
				currency.KRW: cldr.Currency{DisplayName: "სამხრეთ კორეული ვონი", Symbol: "KRW"},
				currency.KWD: cldr.Currency{DisplayName: "ქუვეითური დინარი", Symbol: "KWD"},
				currency.KYD: cldr.Currency{DisplayName: "კაიმანის კუნძულების დოლარი", Symbol: "KYD"},
				currency.KZT: cldr.Currency{DisplayName: "ყაზახური ტენგე", Symbol: "KZT"},
				currency.LAK: cldr.Currency{DisplayName: "ლაოსური კიპი", Symbol: "LAK"},
				currency.LBP: cldr.Currency{DisplayName: "ლიბანური ფუნტი", Symbol: "LBP"},
				currency.LKR: cldr.Currency{DisplayName: "შრი-ლანკური რუპია", Symbol: "LKR"},
				currency.LRD: cldr.Currency{DisplayName: "ლიბერიული დოლარი", Symbol: "LRD"},
				currency.LTL: cldr.Currency{DisplayName: "ლიტვური ლიტა", Symbol: "LTL"},
				currency.LTT: cldr.Currency{DisplayName: "ლიტვური ტალონი", Symbol: ""},
				currency.LUC: cldr.Currency{DisplayName: "ლუქსემბურგის კონვერტირებადი ფრანკი", Symbol: ""},
				currency.LUF: cldr.Currency{DisplayName: "ლუქსემბურგის ფრანკი", Symbol: ""},
				currency.LUL: cldr.Currency{DisplayName: "ლუქსემბურგის ფინანსური ფრანკი", Symbol: ""},
				currency.LVL: cldr.Currency{DisplayName: "ლატვიური ლატი", Symbol: "LVL"},
				currency.LVR: cldr.Currency{DisplayName: "ლატვიური რუბლი", Symbol: ""},
				currency.LYD: cldr.Currency{DisplayName: "ლიბიური დინარი", Symbol: "LYD"},
				currency.MAD: cldr.Currency{DisplayName: "მაროკოს დირჰამი", Symbol: "MAD"},
				currency.MAF: cldr.Currency{DisplayName: "მაროკოს ფრანკი", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "მოლდოვური ლეუ", Symbol: "MDL"},
				currency.MGA: cldr.Currency{DisplayName: "მადაგასკარის არიარი", Symbol: "MGA"},
				currency.MGF: cldr.Currency{DisplayName: "მადაგასკარის ფრანკი", Symbol: ""},
				currency.MKD: cldr.Currency{DisplayName: "მაკედონიური დინარი", Symbol: "MKD"},
				currency.MLF: cldr.Currency{DisplayName: "მალის ფრანკი", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "მიანმარის კიატი", Symbol: "MMK"},
				currency.MNT: cldr.Currency{DisplayName: "მონღოლური ტუგრიკი", Symbol: "MNT"},
				currency.MOP: cldr.Currency{DisplayName: "მაკაუს პატაკა", Symbol: "MOP"},
				currency.MRO: cldr.Currency{DisplayName: "მავრიტანული უგია (1973–2017)", Symbol: "MRO"},
				currency.MRU: cldr.Currency{DisplayName: "მავრიტანული უგია", Symbol: "MRU"},
				currency.MTL: cldr.Currency{DisplayName: "მალტის ლირა", Symbol: ""},
				currency.MTP: cldr.Currency{DisplayName: "მალტის გირვანქა", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "მავრიტანული რუპია", Symbol: "MUR"},
				currency.MVR: cldr.Currency{DisplayName: "მალდივური რუფია", Symbol: "MVR"},
				currency.MWK: cldr.Currency{DisplayName: "მალავიური კვაჩა", Symbol: "MWK"},
				currency.MXN: cldr.Currency{DisplayName: "მექსიკური პესო", Symbol: "MX$"},
				currency.MXP: cldr.Currency{DisplayName: "მექსიკური ვერცხლის პესო (1861–1992)", Symbol: ""},
				currency.MYR: cldr.Currency{DisplayName: "მალაიზიური რინგიტი", Symbol: "MYR"},
				currency.MZE: cldr.Currency{DisplayName: "მოზამბიკური ესკუდო", Symbol: ""},
				currency.MZM: cldr.Currency{DisplayName: "ძველი მოზამბიკური მეტიკალი", Symbol: ""},
				currency.MZN: cldr.Currency{DisplayName: "მოზამბიკური მეტიკალი", Symbol: "MZN"},
				currency.NAD: cldr.Currency{DisplayName: "ნამიბიური დოლარი", Symbol: "NAD"},
				currency.NGN: cldr.Currency{DisplayName: "ნიგერიული ნაირა", Symbol: "NGN"},
				currency.NIC: cldr.Currency{DisplayName: "ნიკარაგუას კორდობა", Symbol: ""},
				currency.NIO: cldr.Currency{DisplayName: "ნიკარაგუული კორდობა", Symbol: "NIO"},
				currency.NLG: cldr.Currency{DisplayName: "ჰოლანდიური გულდენი", Symbol: ""},
				currency.NOK: cldr.Currency{DisplayName: "ნორვეგიული კრონა", Symbol: "NOK"},
				currency.NPR: cldr.Currency{DisplayName: "ნეპალური რუპია", Symbol: "NPR"},
				currency.NZD: cldr.Currency{DisplayName: "ახალი ზელანდიის დოლარი", Symbol: "NZD"},
				currency.OMR: cldr.Currency{DisplayName: "ომანის რიალი", Symbol: "OMR"},
				currency.PAB: cldr.Currency{DisplayName: "პანამური ბალბოა", Symbol: "PAB"},
				currency.PEI: cldr.Currency{DisplayName: "პერუს ინტი", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "პერუს სოლი", Symbol: "PEN"},
				currency.PES: cldr.Currency{DisplayName: "პერუს სოლი (1863–1965)", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "პაპუა-ახალი გვინეის კინა", Symbol: "PGK"},
				currency.PHP: cldr.Currency{DisplayName: "ფილიპინური პესო", Symbol: "PHP"},
				currency.PKR: cldr.Currency{DisplayName: "პაკისტანური რუპია", Symbol: "PKR"},
				currency.PLN: cldr.Currency{DisplayName: "პოლონური ზლოტი", Symbol: "PLN"},
				currency.PLZ: cldr.Currency{DisplayName: "პოლონური ზლოტი (1950–1995)", Symbol: ""},
				currency.PTE: cldr.Currency{DisplayName: "პორტუგალიური ესკუდო", Symbol: ""},
				currency.PYG: cldr.Currency{DisplayName: "პარაგვაული გუარანი", Symbol: "PYG"},
				currency.QAR: cldr.Currency{DisplayName: "კატარის რიალი", Symbol: "QAR"},
				currency.RHD: cldr.Currency{DisplayName: "როდეზიული დოლარი", Symbol: ""},
				currency.ROL: cldr.Currency{DisplayName: "ძველი რუმინული ლეუ", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "რუმინული ლეუ", Symbol: "RON"},
				currency.RSD: cldr.Currency{DisplayName: "სერბული დინარი", Symbol: "RSD"},
				currency.RUB: cldr.Currency{DisplayName: "რუსული რუბლი", Symbol: "RUB"},
				currency.RUR: cldr.Currency{DisplayName: "რუსული რუბლი (1991–1998)", Symbol: "р."},
				currency.RWF: cldr.Currency{DisplayName: "რუანდული ფრანკი", Symbol: "RWF"},
				currency.SAR: cldr.Currency{DisplayName: "საუდის არაბეთის რიალი", Symbol: "SAR"},
				currency.SBD: cldr.Currency{DisplayName: "სოლომონის კუნძულების დოლარი", Symbol: "SBD"},
				currency.SCR: cldr.Currency{DisplayName: "სეიშელური რუპია", Symbol: "SCR"},
				currency.SDD: cldr.Currency{DisplayName: "სუდანის დინარი", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "სუდანური ფუნტი", Symbol: "SDG"},
				currency.SDP: cldr.Currency{DisplayName: "სუდანის გირვანქა", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "შვედური კრონა", Symbol: "SEK"},
				currency.SGD: cldr.Currency{DisplayName: "სინგაპურის დოლარი", Symbol: "SGD"},
				currency.SHP: cldr.Currency{DisplayName: "წმ. ელენეს კუნძულის ფუნტი", Symbol: "SHP"},
				currency.SLL: cldr.Currency{DisplayName: "სიერა-ლეონეს ლეონე", Symbol: "SLL"},
				currency.SOS: cldr.Currency{DisplayName: "სომალური შილინგი", Symbol: "SOS"},
				currency.SRD: cldr.Currency{DisplayName: "სურინამული დოლარი", Symbol: "SRD"},
				currency.SRG: cldr.Currency{DisplayName: "სურინამის გულდენი", Symbol: ""},
				currency.SSP: cldr.Currency{DisplayName: "სამხრეთ სუდანური ფუნტი", Symbol: "SSP"},
				currency.STD: cldr.Currency{DisplayName: "სან-ტომე და პრინსიპის დობრა (1977–2017)", Symbol: "STD"},
				currency.STN: cldr.Currency{DisplayName: "სან-ტომე და პრინსიპის დობრა", Symbol: "STN"},
				currency.SUR: cldr.Currency{DisplayName: "საბჭოთა რუბლი", Symbol: ""},
				currency.SYP: cldr.Currency{DisplayName: "სირიული ფუნტი", Symbol: "SYP"},
				currency.SZL: cldr.Currency{DisplayName: "სვაზილენდის ლილანგენი", Symbol: "SZL"},
				currency.THB: cldr.Currency{DisplayName: "ტაილანდური ბატი", Symbol: "THB"},
				currency.TJR: cldr.Currency{DisplayName: "ტაჯიკური რუბლი", Symbol: ""},
				currency.TJS: cldr.Currency{DisplayName: "ტაჯიკური სომონი", Symbol: "TJS"},
				currency.TMM: cldr.Currency{DisplayName: "თურქმენული მანათი", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "თურქმენეთის მანათი", Symbol: "TMT"},
				currency.TND: cldr.Currency{DisplayName: "ტუნისური დინარი", Symbol: "TND"},
				currency.TOP: cldr.Currency{DisplayName: "ტონგანური პაანგა", Symbol: "TOP"},
				currency.TRL: cldr.Currency{DisplayName: "თურქული ლირა", Symbol: ""},
				currency.TRY: cldr.Currency{DisplayName: "ახალი თურქული ლირა", Symbol: "TRY"},
				currency.TTD: cldr.Currency{DisplayName: "ტრინიდად და ტობაგოს დოლარი", Symbol: "TTD"},
				currency.TWD: cldr.Currency{DisplayName: "ტაივანური ახალი დოლარი", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "ტანზანიური შილინგი", Symbol: "TZS"},
				currency.UAH: cldr.Currency{DisplayName: "უკრაინული გრივნა", Symbol: "UAH"},
				currency.UAK: cldr.Currency{DisplayName: "უკრაინული კარბოვანეცი", Symbol: ""},
				currency.UGS: cldr.Currency{DisplayName: "უგანდური შილინგი (1966–1987)", Symbol: ""},
				currency.UGX: cldr.Currency{DisplayName: "უგანდური შილინგი", Symbol: "UGX"},
				currency.USD: cldr.Currency{DisplayName: "აშშ დოლარი", Symbol: "US$"},
				currency.USN: cldr.Currency{DisplayName: "აშშ დოლარი (შემდეგი დღე)", Symbol: ""},
				currency.USS: cldr.Currency{DisplayName: "აშშ დოლარი (იგივე დღე)", Symbol: ""},
				currency.UYP: cldr.Currency{DisplayName: "ურუგვაის პესო (1975–1993)", Symbol: ""},
				currency.UYU: cldr.Currency{DisplayName: "ურუგვაის პესო", Symbol: "UYU"},
				currency.UZS: cldr.Currency{DisplayName: "უზბეკური სუმი", Symbol: "UZS"},
				currency.VEB: cldr.Currency{DisplayName: "ვენესუელის ბოლივარი (1871–2008)", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "ვენესუელის ბოლივარი (2008–2018)", Symbol: "VEF"},
				currency.VES: cldr.Currency{DisplayName: "ვენესუელის ბოლივარი", Symbol: "VES"},
				currency.VND: cldr.Currency{DisplayName: "ვიეტნამური დონგი", Symbol: "VND"},
				currency.VUV: cldr.Currency{DisplayName: "ვანუატუს ვატუ", Symbol: "VUV"},
				currency.WST: cldr.Currency{DisplayName: "სამოური ტალა", Symbol: "WST"},
				currency.XAF: cldr.Currency{DisplayName: "ცენტრალურ აფრიკული CFA ფრანკი", Symbol: "FCFA"},
				currency.XAG: cldr.Currency{DisplayName: "ვერცხლი", Symbol: ""},
				currency.XBA: cldr.Currency{DisplayName: "ევროპული კომპპოზიტური ერთეული", Symbol: ""},
				currency.XBB: cldr.Currency{DisplayName: "ევროპული ფულადი ერთეული", Symbol: ""},
				currency.XCD: cldr.Currency{DisplayName: "აღმოსავლეთ კარიბიული დოლარი", Symbol: "EC$"},
				currency.XEU: cldr.Currency{DisplayName: "ევროპული სავალუტო ერთეული", Symbol: ""},
				currency.XFO: cldr.Currency{DisplayName: "ფრანგული ოქროს ფრანკი", Symbol: ""},
				currency.XOF: cldr.Currency{DisplayName: "დასავლეთ აფრიკული CFA ფრანკი", Symbol: "CFA"},
				currency.XPF: cldr.Currency{DisplayName: "CFP ფრანკი", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "უცნობი ვალუტა", Symbol: "¤"},
				currency.YDD: cldr.Currency{DisplayName: "იემენის დინარი", Symbol: ""},
				currency.YER: cldr.Currency{DisplayName: "იემენის რეალი", Symbol: "YER"},
				currency.YUD: cldr.Currency{DisplayName: "იუგოსლავიური მყარი დინარი", Symbol: ""},
				currency.YUM: cldr.Currency{DisplayName: "იუგოსლავიური ახალი დინარი", Symbol: ""},
				currency.YUN: cldr.Currency{DisplayName: "იუგოსლავიური კონვერტირებადი დინარი", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "სამხრეთ აფრიკული რანდი", Symbol: "ZAR"},
				currency.ZMK: cldr.Currency{DisplayName: "ზამბიური კვაჭა (1968–2012)", Symbol: "ZMK"},
				currency.ZMW: cldr.Currency{DisplayName: "ზამბიური კვაჭა", Symbol: "ZMW"},
				currency.ZRN: cldr.Currency{DisplayName: "ზაირის ახალი ზაირი", Symbol: ""},
				currency.ZRZ: cldr.Currency{DisplayName: "ზაირის ზაირი", Symbol: ""},
				currency.ZWD: cldr.Currency{DisplayName: "ზიმბაბვეს დოლარი", Symbol: ""},
			},
		},
	}
}
