package cldr

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	currencyPlaceholder = "¤"
)

// numberFormat is a struct that contains all the information about number
// formatting for a specific locale that we need to do number, currency, and
// percentage formatting
type numberFormat struct {
	positivePrefix   string
	positiveSuffix   string
	negativePrefix   string
	negativeSuffix   string
	multiplier       int
	minDecimalDigits int
	maxDecimalDigits int
	minIntegerDigits int
	groupSizeFinal   int // only the right-most (least significant) group
	groupSizeMain    int // all other groups
}

var (
	// prefixSuffixRegex is a regular expression that is used to parse number
	// formats
	prefixSuffixRegex = regexp.MustCompile(`(.*?)[#,.0]+(.*)`)
)

// FmtCurrency takes a float number and a currency key and returns a string
// with a properly formatted currency amount with the correct currency symbol.
// If a symbol cannot be found for the reqested currency, the the key is used
// instead. If the currency key requested is not recognized, it is used as the
// symbol, and an error is returned with the formatted string.
func (n Number) FmtCurrency(currency string, number float64) (formatted string, err error) {
	format := n.parseFormat(n.Formats.Currency, true)
	result := n.formatNumber(format, number)
	formatted = strings.ReplaceAll(result, currencyPlaceholder, n.findCurrency(currency).Symbol)
	return
}

// FmtCurrencyAccounting is the same as FmtCurrency, but using the accounting format in the CLDR. If the accounting
// format is missing, it falls back on the standard format (the same one used in FmtCurrency).
func (n Number) FmtCurrencyAccounting(currency string, number float64) (formatted string, err error) {
	var formatToUse string
	if n.Formats.CurrencyAccounting != "" {
		formatToUse = n.Formats.CurrencyAccounting
	} else {
		formatToUse = n.Formats.Currency
	}
	format := n.parseFormat(formatToUse, true)
	result := n.formatNumber(format, number)
	formatted = strings.ReplaceAll(result, currencyPlaceholder, n.findCurrency(currency).Symbol)
	return
}

// FmtCurrencyWhole does exactly what FmtCurrency does, but it leaves off
// any decimal places. AKA, it would return $100 rather than $100.00.
func (n Number) FmtCurrencyWhole(currency string, number float64) (formatted string, err error) {
	format := n.parseFormat(n.Formats.Currency, false)
	result := n.formatNumber(format, number)
	formatted = strings.ReplaceAll(result, currencyPlaceholder, n.findCurrency(currency).Symbol)
	return
}

// FmtCurrencyAccountingWhole is the accounting equivalent of FmtCurrencyWhole. It falls back on the standard format
// if the accounting format is missing.
func (n Number) FmtCurrencyAccountingWhole(currency string, number float64) (formatted string, err error) {
	var formatToUse string
	if n.Formats.CurrencyAccounting != "" {
		formatToUse = n.Formats.CurrencyAccounting
	} else {
		formatToUse = n.Formats.Currency
	}
	format := n.parseFormat(formatToUse, false)
	result := n.formatNumber(format, number)
	formatted = strings.ReplaceAll(result, currencyPlaceholder, n.findCurrency(currency).Symbol)
	return
}

// FmtNumber takes a float number and returns a properly formatted string
// representation of that number according to the locale's number format.
func (n Number) FmtNumber(number float64) string {
	return n.formatNumber(n.parseFormat(n.Formats.Decimal, true), number)
}

// FmtNumberWhole does exactly what FormatNumber does, but it leaves off any
// decimal places. AKA, it would return 100 rather than 100.01.
func (n Number) FmtNumberWhole(number float64) string {
	return n.formatNumber(n.parseFormat(n.Formats.Decimal, false), number)
}

// FmtPercent takes a float number and returns a properly formatted string
// representation of that number as a percentage according to the locale's
// percentage format.
func (n Number) FmtPercent(number float64) string {
	return n.formatNumber(n.parseFormat(n.Formats.Percent, true), number)
}

// parseFormat takes a format string and returns a numberFormat instance
func (n Number) parseFormat(pattern string, includeDecimalDigits bool) *numberFormat {
	format := new(numberFormat)
	patterns := strings.Split(pattern, ";")

	matches := prefixSuffixRegex.FindAllStringSubmatch(patterns[0], -1)
	if len(matches) > 0 {
		if len(matches[0]) > 1 {
			format.positivePrefix = matches[0][1]
		}
		if len(matches[0]) > 2 {
			format.positiveSuffix = matches[0][2]
		}
	}

	// default values for negative prefix & suffix
	format.negativePrefix = n.Symbols.Negative + format.positivePrefix
	format.negativeSuffix = format.positiveSuffix

	// see if they are in the pattern
	if len(patterns) > 1 {
		matches = prefixSuffixRegex.FindAllStringSubmatch(patterns[1], -1)

		if len(matches) > 0 {
			if len(matches[0]) > 1 {
				format.negativePrefix = matches[0][1]
			}
			if len(matches[0]) > 2 {
				format.negativeSuffix = matches[0][2]
			}
		}
	}

	pat := patterns[0]

	format.multiplier = getFormatMultiplier(pat)

	pos := strings.Index(pat, ".")

	if pos != -1 {
		pos2 := strings.LastIndex(pat, "0")
		if pos2 > pos {
			format.minDecimalDigits = pos2 - pos
		}

		pos3 := strings.LastIndex(pat, "#")
		if pos3 >= pos2 {
			format.maxDecimalDigits = pos3 - pos
		} else {
			format.maxDecimalDigits = format.minDecimalDigits
		}

		pat = pat[0:pos]
	}

	p := strings.ReplaceAll(pat, ",", "")
	pos = strings.Index(p, "0")
	if pos != -1 {
		format.minIntegerDigits = strings.LastIndex(p, "0") - pos + 1
	}

	p = strings.ReplaceAll(pat, "#", "0")
	pos = strings.LastIndex(pat, ",")
	if pos != -1 {
		format.groupSizeFinal = strings.LastIndex(p, "0") - pos
		pos2 := strings.LastIndex(p[0:pos], ",")
		if pos2 != -1 {
			format.groupSizeMain = pos - pos2 - 1
		} else {
			format.groupSizeMain = format.groupSizeFinal
		}
	}

	if !includeDecimalDigits {
		format.maxDecimalDigits = 0
		format.minDecimalDigits = 0
	}

	return format
}

// getFormatMultiplier returns the multiplier to use based on the pattern
func getFormatMultiplier(pat string) int {
	if strings.Contains(pat, "%") {
		return 100
	} else if strings.Contains(pat, "‰") {
		return 1000
	}

	return 1
}

// formatNumber takes an arbitrary numberFormat and a number and applies that
// format to that number, returning the resulting string
func (n Number) formatNumber(format *numberFormat, number float64) string {
	negative := number < 0

	// apply the multiplier first - this is mainly used for percents
	value := math.Abs(number * float64(format.multiplier))
	stringValue := ""

	// get the initial string value, with the maximum # decimal digits
	if format.maxDecimalDigits >= 0 {
		stringValue = numberRound(value, format.maxDecimalDigits)
	} else {
		stringValue = fmt.Sprintf("%f", value)
	}

	// separate the integer from the decimal parts
	pos := strings.Index(stringValue, ".")
	integer := stringValue
	decimal := ""
	if pos != -1 {
		integer = stringValue[:pos]
		decimal = stringValue[pos+1:]
	}

	// make sure the minimum # decimal digits are there
	for len(decimal) < format.minDecimalDigits {
		decimal += "0"
	}

	// make sure the minimum # integer digits are there
	for len(integer) < format.minIntegerDigits {
		integer = "0" + integer
	}

	// if there's a decimal portion, prepend the decimal point symbol
	if len(decimal) > 0 {
		decimal = n.Symbols.Decimal + decimal
	}

	// put the integer portion into properly sized groups
	if format.groupSizeFinal > 0 && len(integer) > format.groupSizeFinal {
		if len(integer) > format.groupSizeMain {
			groupFinal := integer[len(integer)-format.groupSizeFinal:]
			groupFirst := integer[:len(integer)-format.groupSizeFinal]
			integer = strings.Join(
				chunkString(groupFirst, format.groupSizeMain),
				n.Symbols.Group,
			) + n.Symbols.Group + groupFinal
		}
	}

	// append/prepend negative/positive prefix/suffix
	formatted := ""
	if negative {
		formatted = format.negativePrefix + integer + decimal + format.negativeSuffix
	} else {
		formatted = format.positivePrefix + integer + decimal + format.positiveSuffix
	}

	// replace percents and permilles with the local symbols (likely to be exactly the same)
	formatted = strings.ReplaceAll(formatted, "%", n.Symbols.Percent)
	formatted = strings.ReplaceAll(formatted, "‰", n.Symbols.PerMille)

	return formatted
}

// chunkString takes a string and chunks it into size-sized pieces in a slice.
// If the length of the string is not divisible by the size, then the first
// chunk in the slice will be padded to compensate.
func chunkString(str string, size int) []string {
	if str == "" {
		return []string{}
	}

	if size == 0 {
		return []string{str}
	}

	chunks := make([]string, int64(math.Ceil(float64(len(str))/float64(size))))

	for len(str) < len(chunks)*size {
		str = " " + str
	}

	for i := 0; i < len(chunks); i++ {
		start := i * size
		stop := int64(math.Min(float64(start+size), float64(len(str))))
		chunks[i] = str[start:stop]
	}

	chunks[0] = strings.TrimLeft(chunks[0], " ")

	return chunks
}

// numberRound takes a number and returns a string containing a rounded to the
// even with the number of decimal places requested.  If this would result in
// the right most decimal place(s) containing "0"s, then all "0"s on the end of
// the decimal portion will be truncated.
func numberRound(number float64, decimals int) string {
	if number == float64(int64(number)) {
		return strconv.FormatInt(int64(number), 10)
	}

	str := fmt.Sprintf("%f", number)
	pos := strings.Index(str, ".")

	if pos != -1 && len(str) > (pos+decimals) {
		str = str[0 : pos+decimals+1]
	}

	backToNum, _ := strconv.ParseFloat(str, 64)
	difference := number - backToNum
	half := 0.5
	for i := 0; i < decimals; i++ {
		half /= 10
	}

	roundUp := false
	if difference > half {
		roundUp = true
	} else if difference == half {
		// for halfs, round to even
		lastDigit := str[:len(str)-1]
		roundUp = lastDigit == "1" || lastDigit == "3" || lastDigit == "5" || lastDigit == "7" || lastDigit == "9"
	}

	if roundUp {
		// multiply, then ceil, then divide
		multiplier := math.Pow(float64(10), float64(decimals))
		multiplied := strconv.FormatFloat(math.Ceil(number*multiplier), 'f', 0, 64)

		if len(multiplied) > decimals {
			str = multiplied[:len(multiplied)-decimals] + "." + multiplied[len(multiplied)-decimals:]
		} else {
			str = "0." + strings.Repeat("0", decimals-len(multiplied)) + multiplied
		}
	}

	str = strings.TrimRight(str, "0.")

	return str
}

func (n Number) findCurrency(currency string) (c Currency) {
	if cur, ok := n.Currencies[currency]; ok {
		c = cur
	}

	return
}
