/*
Created by hlv_trinh
Licensed under GPLv3
*/

package main 

import (
	"math"
	"fmt"
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
	"thousand", "million", "billion", "trillion", 
	"quadrillion", "quintillion", "sextillion", "septillion",
	"octillion", "nonillion", "decillion", "undecillion", 
	"duodecillion", "tredecillion", "quattuordecillion",
	"quindecillion", "sexdecillion", "septendecillion", 
	"octodecillion", "novemdecillion", "vigintillion",
}

// Converts number into words.
func Convert(number int) (strResult string) {
	// Zero rule
	if number == 0 {
		return _smallNumbers[0]
	}

	// Parse the positive value
	positive := int(math.Abs(float64(number)))

	//Call digit to string recursive
	strResult = convertRecursive(positive, 0)

	//Handle negative int
	if (number) < 0{
		return "minus " + strResult
	} else {
		return strResult
	}
}

func convertRecursive(positiveNum, scalPos int) (ret string){
	if (positiveNum >= 1000){
		ret = convertRecursive(positiveNum / 1000, (scalPos + 1) % len(_scaleNumbers)) + " " + _scaleNumbers[scalPos] + ", "
	}
	
	ret += digitGroup2Text(positiveNum%1000)
	
	return
}

func digitGroup2Text(threeNumGroup int) (ret string) {
	hundreds := threeNumGroup / 100
	tensUnits := threeNumGroup % 100

	if hundreds != 0 {
		ret += _smallNumbers[hundreds] + " hundred and"

		if tensUnits != 0 {
			ret += " "
		}
	}

	tens := tensUnits / 10
	units := tensUnits % 10

	if tens >= 2 {
		ret += _tens[tens]

		if units != 0 {
			ret += "-" + _smallNumbers[units]
		}
	} else if tensUnits != 0 {
		ret += _smallNumbers[tensUnits]
	}

	return
}

func main(){
	fmt.Println(12345678901243212, " -> ", Convert(12345678901243212))
	fmt.Println(-1234, " -> ", Convert(-1234))
	fmt.Println(0, " -> ", Convert(0))
	fmt.Println(1, " -> ", Convert(1))
	fmt.Println(12, " -> ", Convert(12))
	fmt.Println(123, " -> ", Convert(123))
	fmt.Println(1234, " -> ", Convert(1234))
}