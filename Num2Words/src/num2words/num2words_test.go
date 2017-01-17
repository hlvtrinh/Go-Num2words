package num2words_test

import (
	"testing"
	"num2words"
	"strings"
)

var num2words_testConvertNumData = []struct {
	number   int
	expected string
}{
	{0, "zero"},
	{1, "one"},
	{12, "twelve"},
	{123, "one hundred and twenty-three"},
	{1234, "one thousand, two hundred and thirty-four"},
	{10, "ten"},
	{100, "one hundred"},
	{1000, "one thousand"},
	{10000, "ten thousand"},
	{100000, "one hundred thousand"},
	{1000000, "one million"},
	{10000000, "ten million"},
	{1000000001, "one billion and one"},
	{1001000001, "one billion, one million and one"},
	{1001000021, "one billion, one million and twenty-one"},
	{1001000100, "one billion, one million, one hundred"},
	{-1234, "minus one thousand, two hundred and thirty-four"},
	//{12345678901243212,"twelve quadrillion, three hundred and forty-five trillion, six hundred and seventy-eight billion, nine hundred and one million, two hundred and forty-three thousand, two hundred and twelve"},
	//{92233720368547758071,"ninety-two quintillion, two hundred and thirty-three quadrillion, seven hundred and twenty trillion, three hundred and sixty-eight billion, five hundred and forty-seven million, seven hundred and fifty-eight thousand and seventy-one"},
}

var num2words_testConvertStringData = []struct {
	number   string
	expected string
}{
	{"", ""},
	{"0", "zero"},
	{"1", "one"},
	{"   01 ", "one"},
	{"12", "twelve"},
	{"123", "one hundred and twenty-three"},
	{"1234", "one thousand, two hundred and thirty-four"},
	{"10", "ten"},
	{"100", "one hundred"},
	{"1000", "one thousand"},
	{"10000", "ten thousand"},
	{"100000", "one hundred thousand"},
	{"1000000", "one million"},
	{"10000000", "ten million"},
	{"1000000001", "one billion and one"},
	{"1001000001", "one billion, one million and one"},
	{"1001000021", "one billion, one million and twenty-one"},
	{"1001000100", "one billion, one million, one hundred"},
	{"-1234", "minus one thousand, two hundred and thirty-four"},
	{"12345678901243212", "twelve quadrillion, three hundred and forty-five trillion, six hundred and seventy-eight billion, nine hundred and one million, two hundred and forty-three thousand, two hundred and twelve"},
	{"92233720368547758071", "ninety-two quintillion, two hundred and thirty-three quadrillion, seven hundred and twenty trillion, three hundred and sixty-eight billion, five hundred and forty-seven million, seven hundred and fifty-eight thousand and seventy-one"},
	{"a12234", "strconv.ParseInt"},
}

func TestConvertNumber(t *testing.T) {
	for _, tc := range num2words_testConvertNumData {
		actual := num2words.ConvertNumber(tc.number)
		if actual != tc.expected {
			t.Errorf("ConvertNumber(%d): \nexpected\t'%s', \nactual\t\t'%s'", tc.number, tc.expected, actual)
		}
	}
}

func TestConvertString(t *testing.T) {
	for _, tc := range num2words_testConvertStringData {
		actual, err := num2words.ConvertString(tc.number)
		if err != nil {
			if !strings.HasPrefix(err.Error(), tc.expected) {
				t.Errorf("ConvertNumber(%s): \nexpected\t'%s', \nactual\t\t'%s'", tc.number, tc.expected, err)
			}
		} else if actual != tc.expected {
			t.Errorf("ConvertNumber(%s): \nexpected\t'%s', \nactual\t\t'%s'", tc.number, tc.expected, actual)
		}
	}
}
