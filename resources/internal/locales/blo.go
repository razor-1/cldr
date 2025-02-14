// Code generated by make_resources.go; DO NOT EDIT.
package locales

import (
	"github.com/razor-1/cldr"
	"github.com/razor-1/cldr/resources/currency"
	"github.com/razor-1/cldr/resources/language"
	"github.com/razor-1/cldr/resources/territory"
)

func Get_blo() *cldr.Locale {
	return &cldr.Locale{
		Locale:   "blo",
		Calendar: cldr.Calendar{Formats: cldr.CalendarFormats{Date: cldr.CalendarDateFormat{Full: "EEEE, MMMM d y", Long: "MMMM d y", Medium: "MMM d y", Short: "M/d/y"}, Time: cldr.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: cldr.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}, GMT: "{0} Gk"}, FormatNames: cldr.CalendarFormatNames{Months: cldr.CalendarMonthFormatNames{Abbreviated: cldr.CalendarMonthFormatNameValue{Jan: "kaw", Feb: "kpa", Mar: "ci", Apr: "ɖʊ", May: "ɖu5", Jun: "ɖu6", Jul: "la", Aug: "kǝu", Sep: "fʊm", Oct: "cim", Nov: "pom", Dec: "bʊn"}, Narrow: cldr.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"}, Short: cldr.CalendarMonthFormatNameValue{Jan: "", Feb: "", Mar: "", Apr: "", May: "", Jun: "", Jul: "", Aug: "", Sep: "", Oct: "", Nov: "", Dec: ""}, Wide: cldr.CalendarMonthFormatNameValue{Jan: "ɩjikawǝrka kaŋɔrɔ", Feb: "ɩjikpaka kaŋɔrɔ", Mar: "arɛ́cika kaŋɔrɔ", Apr: "njɩbɔ nɖʊka kaŋɔrɔ", May: "acafʊnɖuka kaŋɔrɔ", Jun: "anɔɔɖuka kaŋɔrɔ", Jul: "alàlaka kaŋɔrɔ", Aug: "ɩjikǝuka kaŋɔrɔ", Sep: "abofʊmka kaŋɔrɔ", Oct: "ɩjicimka kaŋɔrɔ", Nov: "acapomka kaŋɔrɔ", Dec: "anɔɔbʊnka kaŋɔrɔ"}}, Days: cldr.CalendarDayFormatNames{Abbreviated: cldr.CalendarDayFormatNameValue{Sun: "alah", Mon: "aɖɩt", Tue: "atal", Wed: "alar", Thu: "alam", Fri: "arɩs", Sat: "asib"}, Narrow: cldr.CalendarDayFormatNameValue{Sun: "lh", Mon: "ɖt", Tue: "tl", Wed: "lr", Thu: "lm", Fri: "rs", Sat: "sb"}, Short: cldr.CalendarDayFormatNameValue{Sun: "alh", Mon: "aɖt", Tue: "atl", Wed: "alr", Thu: "alm", Fri: "ars", Sat: "asb"}, Wide: cldr.CalendarDayFormatNameValue{Sun: "alahaɖɩ", Mon: "aɖɩtɛnɛɛ", Tue: "atalaata", Wed: "alaarba", Thu: "alaamɩshɩ", Fri: "arɩsǝma", Sat: "asiibi"}}, Periods: cldr.CalendarPeriodFormatNames{Abbreviated: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Narrow: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Short: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}, Wide: cldr.CalendarPeriodFormatNameValue{AM: "", PM: ""}}}},
		Number: cldr.Number{
			Symbols: cldr.Symbols{Decimal: ",", Group: "\u00a0", Negative: "-", Percent: "%", PerMille: "‰"},
			Formats: cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00;¤\u00a0-#,##0.00", CurrencyAccounting: "", Percent: "%\u00a0#,#0;%\u00a0-#,#0"},
			Currencies: cldr.Currencies{
				currency.AED: cldr.Currency{DisplayName: "Emiir baGanɔ gaɖɔŋkɔnɔ kaAlaaributǝna kaɖiram", Symbol: ""},
				currency.AFN: cldr.Currency{DisplayName: "Afganistan kahafganii", Symbol: "؋"},
				currency.ALL: cldr.Currency{DisplayName: "Albanii kalɛɛkɩ", Symbol: ""},
				currency.AMD: cldr.Currency{DisplayName: "Armenii kaɖram", Symbol: "֏"},
				currency.ANG: cldr.Currency{DisplayName: "Holanɖ kaKarayiib kafɔlɔrɛŋ", Symbol: ""},
				currency.AOA: cldr.Currency{DisplayName: "Aŋgolaa kakʊwansa", Symbol: "Kz"},
				currency.ARS: cldr.Currency{DisplayName: "Arjantin kapɛsoo", Symbol: "$"},
				currency.AUD: cldr.Currency{DisplayName: "Ɔstraliya kaɖala", Symbol: "A$"},
				currency.AWG: cldr.Currency{DisplayName: "Arubaa kafɔlɔrɛŋ", Symbol: ""},
				currency.AZN: cldr.Currency{DisplayName: "Asɛrbaɩjaŋ kamanaatɩ", Symbol: "₼"},
				currency.BAM: cldr.Currency{DisplayName: "Bɔsniya na Hɛrsegɔfina kamarkɩ", Symbol: "KM"},
				currency.BBD: cldr.Currency{DisplayName: "Barbaɖɔɔsɩ kaɖala", Symbol: "$"},
				currency.BDT: cldr.Currency{DisplayName: "Baŋglaɖɛɛshɩ kataka", Symbol: "৳"},
				currency.BGN: cldr.Currency{DisplayName: "Bulgarii kalɛɛfʊ", Symbol: ""},
				currency.BHD: cldr.Currency{DisplayName: "Barɛɛn kaɖinaa", Symbol: ""},
				currency.BIF: cldr.Currency{DisplayName: "Burunɖii kafaraŋ", Symbol: ""},
				currency.BMD: cldr.Currency{DisplayName: "Bɛrmuɖaa kaɖala", Symbol: "$"},
				currency.BND: cldr.Currency{DisplayName: "Brunɛɩ kaɖala", Symbol: "$"},
				currency.BOB: cldr.Currency{DisplayName: "Bolifiya kabolifiyano", Symbol: "Bs"},
				currency.BRL: cldr.Currency{DisplayName: "Bresil kareyal", Symbol: "R$"},
				currency.BSD: cldr.Currency{DisplayName: "Bahamaasɩ kaɖala", Symbol: "$"},
				currency.BTN: cldr.Currency{DisplayName: "Butan kaŋgulturɔm", Symbol: ""},
				currency.BWP: cldr.Currency{DisplayName: "Bɔsʊwanaa kapula", Symbol: "P"},
				currency.BYN: cldr.Currency{DisplayName: "Belaruus karubǝl", Symbol: ""},
				currency.BZD: cldr.Currency{DisplayName: "Beliis kaɖala", Symbol: "$"},
				currency.CAD: cldr.Currency{DisplayName: "Kanaɖaa kaɖala", Symbol: "CA$"},
				currency.CDF: cldr.Currency{DisplayName: "Koŋgoo kafaraŋ", Symbol: ""},
				currency.CHF: cldr.Currency{DisplayName: "Suwis kafaraŋ", Symbol: ""},
				currency.CLP: cldr.Currency{DisplayName: "Shilii kapɛsoo", Symbol: "$"},
				currency.CNH: cldr.Currency{DisplayName: "Caɩna kayuwan ba sǝ̂ra afʊba ma", Symbol: ""},
				currency.CNY: cldr.Currency{DisplayName: "Caɩna kayuwan", Symbol: "CN¥"},
				currency.COP: cldr.Currency{DisplayName: "Kolɔmbii kapɛsoo", Symbol: "$"},
				currency.CRC: cldr.Currency{DisplayName: "Kɔsta Rikaa kakolɔn", Symbol: "₡"},
				currency.CUC: cldr.Currency{DisplayName: "Kubaa kapɛsoo ba sǝ̂ra afʊba ma", Symbol: "$"},
				currency.CUP: cldr.Currency{DisplayName: "Kubaa kapɛsoo", Symbol: "$"},
				currency.CVE: cldr.Currency{DisplayName: "Kapfɛɛr kahɛskuɖoo", Symbol: ""},
				currency.CZK: cldr.Currency{DisplayName: "Cɛk kakrona", Symbol: "Kč"},
				currency.DJF: cldr.Currency{DisplayName: "Jibutii kafaraŋ", Symbol: ""},
				currency.DKK: cldr.Currency{DisplayName: "Ɖanǝmark kakrona", Symbol: "kr"},
				currency.DOP: cldr.Currency{DisplayName: "Ɖominikaa kapɛsoo", Symbol: "$"},
				currency.DZD: cldr.Currency{DisplayName: "Aljerii kaɖinaa", Symbol: ""},
				currency.EGP: cldr.Currency{DisplayName: "Ejipti kapɔŋ", Symbol: "E£"},
				currency.ERN: cldr.Currency{DisplayName: "Eritree kanafka", Symbol: ""},
				currency.ESP: cldr.Currency{DisplayName: "", Symbol: "₧"},
				currency.ETB: cldr.Currency{DisplayName: "Etiyopii kabiir", Symbol: ""},
				currency.EUR: cldr.Currency{DisplayName: "eroo", Symbol: "€"},
				currency.FJD: cldr.Currency{DisplayName: "Fiji kaɖala", Symbol: "$"},
				currency.FKP: cldr.Currency{DisplayName: "Fɔklanɖ kaBʊtukǝltǝna kapɔŋ", Symbol: "£"},
				currency.GBP: cldr.Currency{DisplayName: "Gagɛɛshɩtǝna kapɔŋ", Symbol: "£"},
				currency.GEL: cldr.Currency{DisplayName: "Jɔrjiya kalari", Symbol: "₾"},
				currency.GHS: cldr.Currency{DisplayName: "Gana kasiɖi", Symbol: "GH₵"},
				currency.GIP: cldr.Currency{DisplayName: "Gibraltaa kapɔŋ", Symbol: "£"},
				currency.GMD: cldr.Currency{DisplayName: "Gambii kaɖalaasi", Symbol: ""},
				currency.GNF: cldr.Currency{DisplayName: "Ginee kafaraŋ", Symbol: "FG"},
				currency.GTQ: cldr.Currency{DisplayName: "Guwatemalaa kakesaal", Symbol: "Q"},
				currency.GYD: cldr.Currency{DisplayName: "Guyanaa kaɖala", Symbol: "$"},
				currency.HKD: cldr.Currency{DisplayName: "Hɔŋ Kɔŋ kaɖala", Symbol: "HK$"},
				currency.HNL: cldr.Currency{DisplayName: "Hɔnɖuraasɩ kalampira", Symbol: "L"},
				currency.HRK: cldr.Currency{DisplayName: "Krowasii kakuna", Symbol: "kn"},
				currency.HTG: cldr.Currency{DisplayName: "Hayitii kaguurɖi", Symbol: ""},
				currency.HUF: cldr.Currency{DisplayName: "Ɔŋgrii kafɔrɩntɩ", Symbol: "Ft"},
				currency.IDR: cldr.Currency{DisplayName: "Ɛnɖonosii karupiyaa", Symbol: "Rp"},
				currency.ILS: cldr.Currency{DisplayName: "Yishraɛl kashekɛl afɔlɩ", Symbol: "₪"},
				currency.INR: cldr.Currency{DisplayName: "Inɖiya karupii", Symbol: "₹"},
				currency.IQD: cldr.Currency{DisplayName: "Ɩraakɩ kaɖinaa", Symbol: ""},
				currency.IRR: cldr.Currency{DisplayName: "Iraŋ kariyal", Symbol: ""},
				currency.ISK: cldr.Currency{DisplayName: "Islanɖ kakrona", Symbol: "kr"},
				currency.JMD: cldr.Currency{DisplayName: "Jamaɩka kaɖala", Symbol: "$"},
				currency.JOD: cldr.Currency{DisplayName: "Jɔrɖanii kaɖinaa", Symbol: ""},
				currency.JPY: cldr.Currency{DisplayName: "Japaŋ kayɛn", Symbol: "JP¥"},
				currency.KES: cldr.Currency{DisplayName: "Keniya kashílè", Symbol: ""},
				currency.KGS: cldr.Currency{DisplayName: "Kirgistan kasɔm", Symbol: "⃀"},
				currency.KHR: cldr.Currency{DisplayName: "Kamboɖiya kariyɛl", Symbol: "៛"},
				currency.KMF: cldr.Currency{DisplayName: "Komɔɔr kafaraŋ", Symbol: "CF"},
				currency.KPW: cldr.Currency{DisplayName: "Koree gʊnyɩpɛnɛlaŋ kawɔn", Symbol: "₩"},
				currency.KRW: cldr.Currency{DisplayName: "Koree gʊnyɩsonolaŋ kawɔn", Symbol: "₩"},
				currency.KWD: cldr.Currency{DisplayName: "Koweeti kaɖinaa", Symbol: ""},
				currency.KYD: cldr.Currency{DisplayName: "Kayimaan kaBʊtukǝltǝna kaɖala", Symbol: "$"},
				currency.KZT: cldr.Currency{DisplayName: "Kasastan katɛŋgɛ", Symbol: "₸"},
				currency.LAK: cldr.Currency{DisplayName: "Lawɔs kakip", Symbol: "₭"},
				currency.LBP: cldr.Currency{DisplayName: "Liibaaŋ kapɔŋ", Symbol: "L£"},
				currency.LKR: cldr.Currency{DisplayName: "Siri Laŋkaa karupii", Symbol: "Rs"},
				currency.LRD: cldr.Currency{DisplayName: "Liberiya kaɖala", Symbol: "$"},
				currency.LSL: cldr.Currency{DisplayName: "Lesotoo kaloti", Symbol: ""},
				currency.LTL: cldr.Currency{DisplayName: "", Symbol: "Lt"},
				currency.LVL: cldr.Currency{DisplayName: "", Symbol: "Ls"},
				currency.LYD: cldr.Currency{DisplayName: "Libii kaɖinaa", Symbol: ""},
				currency.MAD: cldr.Currency{DisplayName: "Morooko kaɖiram", Symbol: ""},
				currency.MDL: cldr.Currency{DisplayName: "Molɖafiya kalewu", Symbol: ""},
				currency.MGA: cldr.Currency{DisplayName: "Maɖagaskaa kaharɩyaarɩ", Symbol: "Ar"},
				currency.MKD: cldr.Currency{DisplayName: "Maseɖoniya kaɖenaa", Symbol: ""},
				currency.MMK: cldr.Currency{DisplayName: "Miyanmaa kakiyaatɩ", Symbol: "K"},
				currency.MNT: cldr.Currency{DisplayName: "Mɔŋgolii katugiriiki", Symbol: "₮"},
				currency.MOP: cldr.Currency{DisplayName: "Makawoo kapataka", Symbol: ""},
				currency.MRU: cldr.Currency{DisplayName: "Moritanii kahugiya", Symbol: ""},
				currency.MUR: cldr.Currency{DisplayName: "Imoris karupii", Symbol: "Rs"},
				currency.MVR: cldr.Currency{DisplayName: "Malɖiifu karufiyaa", Symbol: ""},
				currency.MWK: cldr.Currency{DisplayName: "Malawii kakʊwaasha", Symbol: ""},
				currency.MXN: cldr.Currency{DisplayName: "Mɛsik kapɛsoo", Symbol: "MX$"},
				currency.MYR: cldr.Currency{DisplayName: "Malɛsii kariŋgiiti", Symbol: "RM"},
				currency.MZN: cldr.Currency{DisplayName: "Mosambii kametikal", Symbol: ""},
				currency.NAD: cldr.Currency{DisplayName: "Namibii kaɖala", Symbol: "$"},
				currency.NGN: cldr.Currency{DisplayName: "Nanjiiriya kanɛɛra", Symbol: "₦"},
				currency.NIO: cldr.Currency{DisplayName: "Nikaraguwaa kakɔrɖoba", Symbol: "C$"},
				currency.NOK: cldr.Currency{DisplayName: "Nɔrfɛsh kakrona", Symbol: "kr"},
				currency.NPR: cldr.Currency{DisplayName: "Neepal karupii", Symbol: "Rs"},
				currency.NZD: cldr.Currency{DisplayName: "Selanɖ afɔlɩ kaɖala", Symbol: "NZ$"},
				currency.OMR: cldr.Currency{DisplayName: "Oman kariyal", Symbol: ""},
				currency.PAB: cldr.Currency{DisplayName: "Panamaa kabalbowa", Symbol: ""},
				currency.PEN: cldr.Currency{DisplayName: "Peruu kasol", Symbol: ""},
				currency.PGK: cldr.Currency{DisplayName: "Papuasii Ginee afɔlɩ kakina", Symbol: ""},
				currency.PHP: cldr.Currency{DisplayName: "Filipiin kapɛsoo", Symbol: "₱"},
				currency.PKR: cldr.Currency{DisplayName: "Pakistan karupii", Symbol: "Rs"},
				currency.PLN: cldr.Currency{DisplayName: "Polanɖ kasǝlɔɔtɩ", Symbol: "zł"},
				currency.PYG: cldr.Currency{DisplayName: "Paraguwee kaguwarani", Symbol: "₲"},
				currency.QAR: cldr.Currency{DisplayName: "Kataa kariyal", Symbol: ""},
				currency.RON: cldr.Currency{DisplayName: "Romanii kalewu", Symbol: "lei"},
				currency.RSD: cldr.Currency{DisplayName: "Sɛrbii kaɖinaa", Symbol: ""},
				currency.RUB: cldr.Currency{DisplayName: "Rɔɔshɩya karubǝl", Symbol: "₽"},
				currency.RWF: cldr.Currency{DisplayName: "Rʊwanɖaa kafaraŋ", Symbol: "RF"},
				currency.SAR: cldr.Currency{DisplayName: "Sauɖiya kariyal", Symbol: ""},
				currency.SBD: cldr.Currency{DisplayName: "Salomɔɔn kaBʊtukǝltǝna kaɖala", Symbol: "$"},
				currency.SCR: cldr.Currency{DisplayName: "Seshɛl karupii", Symbol: ""},
				currency.SDG: cldr.Currency{DisplayName: "Suɖaŋ kapɔŋ", Symbol: ""},
				currency.SEK: cldr.Currency{DisplayName: "Sʊwɛɖ kakrona", Symbol: "kr"},
				currency.SGD: cldr.Currency{DisplayName: "Siŋgapuur kaɖala", Symbol: "$"},
				currency.SHP: cldr.Currency{DisplayName: "Sɛŋ Elenaa kapɔŋ", Symbol: "£"},
				currency.SLE: cldr.Currency{DisplayName: "Seraleyɔn kaleyɔn", Symbol: ""},
				currency.SLL: cldr.Currency{DisplayName: "Seraleyɔn kaleyɔn (1964—2022)", Symbol: ""},
				currency.SOS: cldr.Currency{DisplayName: "Somalii kashílè", Symbol: ""},
				currency.SRD: cldr.Currency{DisplayName: "Surinam kaɖala", Symbol: "$"},
				currency.SSP: cldr.Currency{DisplayName: "Suɖaŋ gʊnyɩsonolaŋ kapɔŋ", Symbol: "£"},
				currency.STN: cldr.Currency{DisplayName: "Saotomee kaɖobra", Symbol: "Db"},
				currency.SYP: cldr.Currency{DisplayName: "Sirii kapɔŋ", Symbol: "£"},
				currency.SZL: cldr.Currency{DisplayName: "Sʊwasilanɖ kalilaŋgenii", Symbol: ""},
				currency.THB: cldr.Currency{DisplayName: "Taɩlanɖ kabaatɩ", Symbol: "฿"},
				currency.TJS: cldr.Currency{DisplayName: "Tajikistan kasomooni", Symbol: ""},
				currency.TMT: cldr.Currency{DisplayName: "Turkmenistan kamanaatɩ", Symbol: ""},
				currency.TND: cldr.Currency{DisplayName: "Tunisii kaɖinaa", Symbol: ""},
				currency.TOP: cldr.Currency{DisplayName: "Tɔŋga kapaŋga", Symbol: "T$"},
				currency.TRY: cldr.Currency{DisplayName: "Turkii kalira", Symbol: "₺"},
				currency.TTD: cldr.Currency{DisplayName: "Triniɖaaɖ na Tobagoo kaɖala", Symbol: "$"},
				currency.TWD: cldr.Currency{DisplayName: "Taɩwan kaɖala afɔlɩ", Symbol: "NT$"},
				currency.TZS: cldr.Currency{DisplayName: "Taŋsanii kashílè", Symbol: ""},
				currency.UAH: cldr.Currency{DisplayName: "Ikrɛɛn karifniya", Symbol: "₴"},
				currency.UGX: cldr.Currency{DisplayName: "Uganɖaa kashílè", Symbol: ""},
				currency.USD: cldr.Currency{DisplayName: "Amalɩka kaɖala", Symbol: "US$"},
				currency.UYU: cldr.Currency{DisplayName: "Uruguwee kapɛsoo", Symbol: "$"},
				currency.UZS: cldr.Currency{DisplayName: "Usbeekistan kasɔm", Symbol: ""},
				currency.VEF: cldr.Currency{DisplayName: "", Symbol: "Bs"},
				currency.VES: cldr.Currency{DisplayName: "Fenesuwelaa kabolifar", Symbol: ""},
				currency.VND: cldr.Currency{DisplayName: "Fɛtnam kaɖɔŋgɩ", Symbol: "₫"},
				currency.VUV: cldr.Currency{DisplayName: "Fanuwatu kafatu", Symbol: ""},
				currency.WST: cldr.Currency{DisplayName: "Samowa katala", Symbol: ""},
				currency.XAF: cldr.Currency{DisplayName: "Garɩɖontǝna gɩcɩɩca kasɛɛfa", Symbol: "FCFA"},
				currency.XCD: cldr.Currency{DisplayName: "Karayiib gajakalaŋ kaɖala", Symbol: "EC$"},
				currency.XCG: cldr.Currency{DisplayName: "", Symbol: "Cg."},
				currency.XOF: cldr.Currency{DisplayName: "Garɩɖontǝna gɩteŋshilelaŋ kasɛɛfa", Symbol: "F\u202fCFA"},
				currency.XPF: cldr.Currency{DisplayName: "Polinesiya Gafɔntǝna kaja kafaraŋ", Symbol: "CFPF"},
				currency.XXX: cldr.Currency{DisplayName: "gɩtanɩɩ kʊyɔʊ ʊ mana ma", Symbol: "¤"},
				currency.YER: cldr.Currency{DisplayName: "Yemɛn kariyal", Symbol: ""},
				currency.ZAR: cldr.Currency{DisplayName: "Sautafrika karanɖɩ", Symbol: "R"},
				currency.ZMW: cldr.Currency{DisplayName: "Sambii kakʊwaasha", Symbol: "ZK"},
			},
		},
		Display: cldr.LocaleDisplayPattern{Pattern: "{0} ({1})", Separator: "{0}, {1}", KeyTypePattern: "{0} : {1}"},
		Languages: cldr.Languages{
			language.AR:      "gɩlaaribuja",
			language.AR_001:  "gɩlaaribuja ŋgɩɖee kǝ ba na fʊba na",
			language.BLO:     "anii kagɩja",
			language.BN:      "baŋglaa kagɩja",
			language.DE:      "gɩjaamaja",
			language.EN:      "gɛɛshɩ",
			language.EN_US:   "gɛɛshɩ (GKA)",
			language.ES:      "gɩspaŋja",
			language.ES_419:  "gɩspaŋja (latɛŋ kaAmalɩkatǝna)",
			language.FR:      "gɩfɔnɔ",
			language.HI_LATN: "hiŋgliishɩ kagɩja",
			language.ID:      "Ɛnɖonosii kagɩja",
			language.IT:      "gɩtaliija",
			language.JA:      "gɩjapaŋja",
			language.KO:      "Koree kagɩja",
			language.NL:      "Holanɖ kagɩja",
			language.PL:      "Polanɖ kagɩja",
			language.PT:      "gɩpɔrtigalja",
			language.RU:      "gɩrɔɔshɩyaja",
			language.TH:      "taɩ kagɩja",
			language.TR:      "gɩturkiija",
			language.UND:     "gɩkrǝ ŋgɩɖee kʊyɔʊ ʊ mana ma",
			language.ZH:      "gɩcaɩnaja, manɖarɛŋ",
			language.ZH_HANS: "gɩcaɩnaja manɖarɛŋ gɩburoka",
			language.ZH_HANT: "gɩcaɩnaja manɖarɛŋ tututu",
		},
		Territories: cldr.Territories{
			territory.V_001: "nɖulinya",
			territory.V_002: "Garɩɖontǝna",
			territory.V_003: "Gamalɩkatǝna gʊnyɩpɛnɛlaŋ",
			territory.V_005: "Gamalɩkatǝna gʊnyɩsonolaŋ",
			territory.V_009: "Oseyanii",
			territory.V_011: "Garɩɖontǝna gɩteŋshilelaŋ",
			territory.V_013: "Gamalɩkatǝna gɩcɩɩca",
			territory.V_014: "Garɩɖontǝna gajakalaŋ",
			territory.V_015: "Garɩɖontǝna gʊnyɩpɛnɛlaŋ",
			territory.V_017: "Garɩɖontǝna gɩcɩɩca",
			territory.V_018: "Garɩɖontǝna gʊnyɩsonolaŋ",
			territory.V_019: "Gamalɩkatǝna",
			territory.V_021: "Gamalɩkatǝna kagʊnyɩpɛnɛlaŋ",
			territory.V_029: "Karayiib",
			territory.V_030: "Gacǝlǝŋtǝna gajakalaŋ",
			territory.V_034: "Gacǝlǝŋtǝna gʊnyɩsonolaŋ",
			territory.V_035: "Gacǝlǝŋtǝna gʊsono na gajakayɛlaŋ kʊfɔɔ nɩ",
			territory.V_039: "Garɩfɔntǝna gʊnyɩsonolaŋ",
			territory.V_053: "Ɔstracǝlǝŋtǝna",
			territory.V_054: "Melanesiya",
			territory.V_057: "Mikronesiya kagʊsaʊ",
			territory.V_061: "Polinesiya",
			territory.V_142: "Gacǝlǝŋtǝna",
			territory.V_143: "Gacǝlǝŋtǝna gɩcɩɩca",
			territory.V_145: "Gacǝlǝŋtǝna gɩteŋshilelaŋ",
			territory.V_150: "Garɩfɔntǝna",
			territory.V_151: "Garɩfɔntǝna gajakalaŋ",
			territory.V_154: "Garɩfɔntǝna gʊnyɩpɛnɛlaŋ",
			territory.V_155: "Garɩfɔntǝna gɩteŋshilelaŋ",
			territory.V_202: "Garɩɖontǝna Sahara katǝntǝn",
			territory.V_419: "Latɛŋ kaAmalɩkatǝna",
			territory.AC:    "Asɛnsiyɔɔn kaAtukǝltǝna",
			territory.AD:    "Anɖɔraa",
			territory.AE:    "Emiir baGanɔ gaɖɔŋkɔnɔ kaAlaaributǝna",
			territory.AF:    "Afganistan",
			territory.AG:    "Antiguwaa na Barbuɖaa",
			territory.AI:    "Aŋguwilaa",
			territory.AL:    "Albanii",
			territory.AM:    "Armenii",
			territory.AO:    "Aŋgolaa",
			territory.AQ:    "Gatutaltǝna",
			territory.AR:    "Arjantin",
			territory.AS:    "Samowa Amalɩka kaja",
			territory.AT:    "Otrish",
			territory.AU:    "Ɔstraliya",
			territory.AW:    "Arubaa",
			territory.AX:    "Ɔɔlanɖ kaBʊtǝlǝltǝna",
			territory.AZ:    "Asɛrbaɩjaŋ",
			territory.BA:    "Bɔsniya na Hɛrsegɔfina",
			territory.BB:    "Barbaɖɔɔsɩ",
			territory.BD:    "Baŋglaɖɛɛshɩ",
			territory.BE:    "Bɛljiiki",
			territory.BF:    "Burkinaa",
			territory.BG:    "Bulgarii",
			territory.BH:    "Barɛɛn",
			territory.BI:    "Burunɖii",
			territory.BJ:    "Benɛɛ",
			territory.BL:    "Sɛŋ-Batolomayɔ",
			territory.BM:    "Bɛrmuɖaa",
			territory.BN:    "Brunɛɩ",
			territory.BO:    "Bolifiya",
			territory.BQ:    "Holanɖ kaKarayiib",
			territory.BR:    "Bresil",
			territory.BS:    "Bahamaasɩ",
			territory.BT:    "Butan",
			territory.BV:    "Bufee kaAtukǝltǝna",
			territory.BW:    "Bɔsʊwanaa",
			territory.BY:    "Belaruus",
			territory.BZ:    "Beliis",
			territory.CA:    "Kanaɖaa",
			territory.CC:    "Kokoos (Kiiliŋ) kaBʊtukǝltǝna",
			territory.CD:    "Koŋgoo Kinshasaa",
			territory.CF:    "Santrafrika",
			territory.CG:    "Koŋgoo Brasafil",
			territory.CH:    "Suwis",
			territory.CI:    "Koɖifʊaa",
			territory.CK:    "Kʊkʊ kaBʊtukǝltǝna",
			territory.CL:    "Shilii",
			territory.CM:    "Kamerun",
			territory.CN:    "Caɩna",
			territory.CO:    "Kolɔmbii",
			territory.CP:    "Klipɛɛtɔn kaAtukǝltǝna",
			territory.CQ:    "Sark",
			territory.CR:    "Kɔsta Rikaa",
			territory.CU:    "Kubaa",
			territory.CV:    "Kapfɛɛr",
			territory.CW:    "Kurasawuu",
			territory.CX:    "Nowɛl kaAtukǝltǝna",
			territory.CY:    "Ciprɔs",
			territory.CZ:    "Cɛk",
			territory.DE:    "Gajaamatǝna",
			territory.DG:    "Ɖiyego Garsiya",
			territory.DJ:    "Jibutii",
			territory.DK:    "Ɖanǝmark",
			territory.DM:    "Ɖominikaa",
			territory.DO:    "Ɖominikaa kaRepibliiki",
			territory.DZ:    "Aljerii",
			territory.EA:    "Seyuta na Meliliya",
			territory.EC:    "Ekuwaɖɔɔr",
			territory.EE:    "Ɛstoniya",
			territory.EG:    "Ejipti",
			territory.EH:    "Sarawii",
			territory.ER:    "Eritree",
			territory.ES:    "Ɛspanyǝ",
			territory.ET:    "Etiyopii",
			territory.EU:    "Ganɔ gaɖɔŋkɔnɔ kaBʊtǝna Garɩfɔntǝna nɩ",
			territory.EZ:    "Eroo kaBʊtǝna",
			territory.FI:    "Fɛnlanɖ",
			territory.FJ:    "Fiji",
			territory.FK:    "Fɔklanɖ kaBʊtukǝltǝna",
			territory.FM:    "Mikronesiya",
			territory.FO:    "Faroi kaBʊtukǝltǝna",
			territory.FR:    "Gafɔntǝna",
			territory.GA:    "Gabɔŋ",
			territory.GB:    "Gagɛɛshɩtǝna",
			territory.GD:    "Grenaɖaa",
			territory.GE:    "Jɔrjiya",
			territory.GF:    "Guyanaa Gafɔntǝna kaja",
			territory.GG:    "Gǝrǝnsɛɩ",
			territory.GH:    "Gana",
			territory.GI:    "Gibraltaa",
			territory.GL:    "Grinlanɖ",
			territory.GM:    "Gambii",
			territory.GN:    "Ginee",
			territory.GP:    "Guwaɖeluupu",
			territory.GQ:    "Ginee Malabo",
			territory.GR:    "Grɛs",
			territory.GS:    "Jɔrjiya gʊnyɩsonolaŋ kaja na Sanɖuush gʊnyɩsonolaŋ kaBʊtukǝltǝna",
			territory.GT:    "Guwatemalaa",
			territory.GU:    "Guwam",
			territory.GW:    "Ginee Bisoo",
			territory.GY:    "Guyanaa",
			territory.HK:    "Hɔŋ Kɔŋ Caɩna kaja",
			territory.HM:    "Hɛɛrɖ na Mɛkɖɔnalɖ kaBʊtukǝltǝna",
			territory.HN:    "Hɔnɖuraasɩ",
			territory.HR:    "Krowasii",
			territory.HT:    "Hayitii",
			territory.HU:    "Ɔŋgrii",
			territory.IC:    "Kanarii kaBʊtukǝltǝna",
			territory.ID:    "Ɛnɖonosii",
			territory.IE:    "Irlanɖ",
			territory.IL:    "Yishraɛl",
			territory.IM:    "Man kaAtukǝltǝna",
			territory.IN:    "Inɖiya",
			territory.IO:    "Gɛɛshɩ kaAtǝna Inɖiya kaTeŋku nɩ",
			territory.IQ:    "Ɩraakɩ",
			territory.IR:    "Iraŋ",
			territory.IS:    "Islanɖ",
			territory.IT:    "Italii",
			territory.JE:    "Jersei",
			territory.JM:    "Jamaɩka",
			territory.JO:    "Jɔrɖanii",
			territory.JP:    "Japaŋ",
			territory.KE:    "Keniya",
			territory.KG:    "Kirgistan",
			territory.KH:    "Kamboɖiya",
			territory.KI:    "Kiribatii",
			territory.KM:    "Komɔɔr",
			territory.KN:    "Sɛŋ Kits na Nefis",
			territory.KP:    "Koree gʊnyɩpɛnɛlaŋ",
			territory.KR:    "Koree gʊnyɩsonolaŋ",
			territory.KW:    "Koweeti",
			territory.KY:    "Kayimaan kaBʊtukǝltǝna",
			territory.KZ:    "Kasastan",
			territory.LA:    "Lawɔs",
			territory.LB:    "Liibaaŋ",
			territory.LC:    "Sɛŋ Lusiya",
			territory.LI:    "Liishtɛntaɩn",
			territory.LK:    "Siri Laŋkaa",
			territory.LR:    "Liberiya",
			territory.LS:    "Lesotoo",
			territory.LT:    "Lituwaniya",
			territory.LU:    "Lusɛmbuur",
			territory.LV:    "Lɛtfiya",
			territory.LY:    "Libii",
			territory.MA:    "Morooko",
			territory.MC:    "Monakoo",
			territory.MD:    "Mɔlɖafiya",
			territory.ME:    "Mɔntenegroo",
			territory.MF:    "Sɛŋ Martɛɛŋ",
			territory.MG:    "Maɖagaskaa",
			territory.MH:    "Marshal kaBʊtukǝltǝna",
			territory.MK:    "Maseɖoniya gʊnyɩpɛnɛlaŋ kaja",
			territory.ML:    "Malii",
			territory.MM:    "Miyanmaa (Birmanii)",
			territory.MN:    "Mɔŋgolii",
			territory.MO:    "Makawoo Caɩna kaja",
			territory.MP:    "Mariyan kaBʊtukǝltǝna gʊnyɩpɛnɛlaŋ",
			territory.MQ:    "Martiniiki",
			territory.MR:    "Moritanii",
			territory.MS:    "Mɔnsɛraatɩ",
			territory.MT:    "Malta",
			territory.MU:    "Imoris",
			territory.MV:    "Malɖiifu",
			territory.MW:    "Malawii",
			territory.MX:    "Mɛsik",
			territory.MY:    "Malɛsii",
			territory.MZ:    "Mosambii",
			territory.NA:    "Namibii",
			territory.NC:    "Kaleɖonii afɔlɩ",
			territory.NE:    "Nijɛr",
			territory.NF:    "Nɔrfook kaAtukǝltǝna",
			territory.NG:    "Nanjiiriya",
			territory.NI:    "Nikaraguwaa",
			territory.NL:    "Holanɖ",
			territory.NO:    "Nɔrfɛsh",
			territory.NP:    "Neepal",
			territory.NR:    "Nawuru",
			territory.NU:    "Niwuye",
			territory.NZ:    "Selanɖ afɔlɩ",
			territory.OM:    "Oman",
			territory.PA:    "Panamaa",
			territory.PE:    "Peruu",
			territory.PF:    "Polinesiya Gafɔntǝna kaja",
			territory.PG:    "Papuasii Ginee afɔlɩ",
			territory.PH:    "Filipiin",
			territory.PK:    "Pakistan",
			territory.PL:    "Polanɖ",
			territory.PM:    "Sɛŋ-Petrɔs na Mikelɔŋ",
			territory.PN:    "Pɩtkɛɛn kaBʊtukǝltǝna",
			territory.PR:    "Pɔrto Rikoo",
			territory.PS:    "Palɛstiin kAsàʊ",
			territory.PT:    "Pɔrtigal",
			territory.PW:    "Palawoo",
			territory.PY:    "Paraguwee",
			territory.QA:    "Kataa",
			territory.QO:    "Oseyanii kasaʊlǝŋka",
			territory.RE:    "Reeniyɔŋ",
			territory.RO:    "Romanii",
			territory.RS:    "Sɛrbii",
			territory.RU:    "Rɔɔshɩya",
			territory.RW:    "Rʊwanɖaa",
			territory.SA:    "Sauɖiya",
			territory.SB:    "Salomɔɔn kaBʊtukǝltǝna",
			territory.SC:    "Seshɛl",
			territory.SD:    "Suɖaŋ",
			territory.SE:    "Sʊwɛɖ",
			territory.SG:    "Siŋgapuur",
			territory.SH:    "Sɛŋ Elenaa (kaAtukǝltǝna)",
			territory.SI:    "Slofeniya",
			territory.SJ:    "Sǝfalbaaɖ na Yan Mayɛn",
			territory.SK:    "Slofakii",
			territory.SL:    "Seraleyɔn",
			territory.SM:    "Sɛŋ Marinoo",
			territory.SN:    "Senegal",
			territory.SO:    "Somalii",
			territory.SR:    "Surinam",
			territory.SS:    "Suɖaŋ gʊnyɩsonolaŋ",
			territory.ST:    "Saotomee",
			territory.SV:    "Ɛl Salfaɖɔɔr",
			territory.SX:    "Sɛŋ Martɛɛŋ (Holanɖ kaja)",
			territory.SY:    "Sirii",
			territory.SZ:    "Ɛsʊwatinii",
			territory.TA:    "Tristan ɖa Kuna",
			territory.TC:    "Turkisii na Kayɩkɔɔsɩ kaBʊtukǝltǝna",
			territory.TD:    "Caɖ",
			territory.TF:    "Gafɔntǝna kaBʊtǝna gʊnyɩsonolaŋ kabʊja",
			territory.TG:    "Togoo",
			territory.TH:    "Taɩlanɖ",
			territory.TJ:    "Tajikistan",
			territory.TK:    "Tokelaʊ",
			territory.TL:    "Timɔɔ gajakalaŋ",
			territory.TM:    "Turkmenistan",
			territory.TN:    "Tunisii",
			territory.TO:    "Tɔŋga",
			territory.TR:    "Turkii",
			territory.TT:    "Triniɖaaɖ na Tobagoo",
			territory.TV:    "Tufalu",
			territory.TW:    "Taɩwan",
			territory.TZ:    "Taŋsanii",
			territory.UA:    "Ikrɛɛn",
			territory.UG:    "Uganɖaa",
			territory.UM:    "Ganɔ gaɖɔŋkɔnɔ kaBʊtǝna Amalɩka nɩ kaBʊtukǝltǝna bʊlǝŋka",
			territory.UN:    "Ganɔ gaɖɔŋkɔnɔ kaBʊtǝna nɖulinya nɩ",
			territory.US:    "Ganɔ gaɖɔŋkɔnɔ kaBʊtǝna Amalɩka nɩ",
			territory.UY:    "Uruguwee",
			territory.UZ:    "Usbeekistan",
			territory.VA:    "Fatikaŋ kaMpá",
			territory.VC:    "Sɛŋ Fɩnsaŋ na Grenaɖiniisi",
			territory.VE:    "Fenesuwelaa",
			territory.VG:    "Fɩrjɩɩn kǝBʊtukǝltǝna Gɛɛshɩ kabʊja",
			territory.VI:    "Fɩrjɩɩn kaBʊtukǝltǝna Amalɩka kabʊja",
			territory.VN:    "Fɛtnam",
			territory.VU:    "Fanuwatu",
			territory.WF:    "Walis na Futuna",
			territory.WS:    "Samowa",
			territory.XA:    "sǝɖoo-aksaŋ",
			territory.XB:    "sǝɖoo-biɖi",
			territory.XK:    "Kɔsofoo",
			territory.YE:    "Yemɛn",
			territory.YT:    "Mayɔɔtɩ",
			territory.ZA:    "Sautafrika",
			territory.ZM:    "Sambii",
			territory.ZW:    "Simbabʊwee",
			territory.ZZ:    "gʊsaʊɩ kʊyɔʊ ʊ mana ma",
		},
	}
}
