/*
Created by hlv_trinh
Licensed under GPLv3
*/

package num2words

import (
	"math"
	"strconv"
	"strings"
)

var _smallNumbers = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen",
	"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}
var _tens = []string{
	"", "", "twenty", "thirty", "forty", "fifty",
	"sixty", "seventy", "eighty", "ninety",
}
var _scaleNumbers = []string{
	"", "thousand", "million", "billion", "trillion",
	"quadrillion", "quintillion", "sextillion", "septillion",
	"octillion", "nonillion", "decillion", "undecillion",
	"duodecillion", "tredecillion", "quattuordecillion",
	"quindecillion", "sexdecillion", "septendecillion",
	"octodecillion", "novemdecillion", "vigintillion",
}

/* Converts number into words.
*	Max range:
		- 32bit: int32 range (-2147483648 to 2147483647)
		- 64bit: int64 range (-9223372036854775808 to 9223372036854775807)
*/
func ConvertNumber(number int) string {
	// Parse the positive value
	positive := int(math.Abs(float64(number)))

	//Call digit to string recursive
	ret := convertRecursive(positive, 0)

	//Handle negative int
	if (number) < 0 {
		return "minus " + ret
	} else {
		return ret
	}
}

func ConvertString(number string) (ret string, err error) {
	// Trim first
	number = strings.TrimSpace(number)
	if number == "" {
		return
	}

	//Handle negative int
	if strings.HasPrefix(number, "-") {
		ret = "minus "
		number = number[1:]
	}

	//Remove leading zeros
	number = strings.TrimLeft(number, "0")
	if number == "" {
		number = "0"
	}

	//Call digit to string recursive
	strNum, tmpErr := convertStringRecursive(number, 0)

	if tmpErr != nil {
		err = tmpErr
		ret = ""
		return
	}

	ret += strNum

	return
}

func convertStringRecursive(positiveDigits string, scalPos int) (ret string, err error) {
	hasRemaining := len(positiveDigits) > 3
	threeDigitGroup := positiveDigits

	if hasRemaining {
		ret, err = convertStringRecursive(positiveDigits[:len(positiveDigits)-3], scalPos+1)
		if err != nil {
			ret = ""
			return
		}
		threeDigitGroup = positiveDigits[len(positiveDigits)-3:]
	}

	threeNumGroup, err := strconv.Atoi(threeDigitGroup)
	if err != nil {
		ret = ""
		return
	}
	ret += digitGroup2Text(threeNumGroup, scalPos, hasRemaining)

	return
	return
}

func convertRecursive(positiveNum, scalPos int) (ret string) {
	remainingGroups := positiveNum / 1000
	hasRemaining := remainingGroups > 0

	if hasRemaining {
		ret = convertRecursive(remainingGroups, scalPos+1)
	}

	ret += digitGroup2Text(positiveNum%1000, scalPos, hasRemaining)

	return
}

func digitGroup2Text(threeNumGroup, scalPos int, hasRemaining bool) (ret string) {
	if threeNumGroup > 0 {
		hundreds := threeNumGroup / 100
		tensUnits := threeNumGroup % 100
		tens := tensUnits / 10
		units := tensUnits % 10

		//Delimiter between groups
		if hasRemaining {
			if hundreds == 0 && scalPos == 0 {
				ret += " and "
			} else {
				ret += ", "
			}
		}

		//Hundred number
		if hundreds != 0 {
			ret += _smallNumbers[hundreds] + " hundred"

			if tensUnits != 0 {
				ret += " and "
			}
		}

		//Unit word
		if tens >= 2 {
			ret += _tens[tens]

			if units != 0 {
				ret += "-" + _smallNumbers[units]
			}
		} else if tensUnits != 0 {
			ret += _smallNumbers[tensUnits]
		}

		//Scale word
		if scalPos > 0 {
			ret += " " + _scaleNumbers[scalPos%len(_scaleNumbers)]
		}
	} else if !hasRemaining {
		//Zero rule
		ret = _smallNumbers[0]
	}
	return
}
