package property

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

//skipped the roman to arabic part - maybe go back to that eventually

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(numToConvert int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for numToConvert >= numeral.Value {
			result.WriteString(numeral.Symbol)
			numToConvert -= numeral.Value
		}
	}
	return result.String()
}
