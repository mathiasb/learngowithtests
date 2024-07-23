package proptesting

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Description string
	Arabic      int
	Want        string
}{
	{"1 to I", 1, "I"},
	{"2 to II", 2, "II"},
	{"3 to III", 3, "III"},
	{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
	{"5 gets converted to V", 5, "V"},
	{"8 gets converted to VIII", 8, "VIII"},
	{"9 gets converted to IX", 9, "IX"},
	{"10 gets converted to X", 10, "X"},
	{"14 gets converted to XIV", 14, "XIV"},
	{"18 gets converted to XVIII", 18, "XVIII"},
	{"20 gets converted to XX", 20, "XX"},
	{"39 gets converted to XXXIX", 39, "XXXIX"},
	{"40 gets converted to XL", 40, "XL"},
	{"47 gets converted to XLVII", 47, "XLVII"},
	{"49 gets converted to XLIX", 49, "XLIX"},
	{Description: "50 gets converted to L", Arabic: 50, Want: "L"},
	{Description: "100 gets converted to C", Arabic: 100, Want: "C"},
	{Description: "90 gets converted to XC", Arabic: 90, Want: "XC"},
	{Description: "400 gets converted to CD", Arabic: 400, Want: "CD"},
	{Description: "500 gets converted to D", Arabic: 500, Want: "D"},
	{Description: "900 gets converted to CM", Arabic: 900, Want: "CM"},
	{Description: "1000 gets converted to M", Arabic: 1000, Want: "M"},
	{Description: "1984 gets converted to MCMLXXXIV", Arabic: 1984, Want: "MCMLXXXIV"},
	{Description: "3999 gets converted to MMMCMXCIX", Arabic: 3999, Want: "MMMCMXCIX"},
	{Description: "2014 gets converted to MMXIV", Arabic: 2014, Want: "MMXIV"},
	{Description: "1006 gets converted to MVI", Arabic: 1006, Want: "MVI"},
	{Description: "798 gets converted to DCCXCVIII", Arabic: 798, Want: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Want {
				t.Errorf("got %s, want %s", got, test.Want)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Want, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Want)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {

		if arabic > 3999 {
			log.Println(arabic)
			return true
		}
		roman := ConvertToRoman(int(arabic))
		fromRoman := ConvertToArabic(roman)
		return uint16(fromRoman) == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
