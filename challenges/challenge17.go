package challenges

import "fmt"

// If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there
// are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words,
// how many letters would be used?

var zeroTo9 = [10]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var tenTo19 = [10]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var zeroTo100 = [10]string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func intToText(n int) string {
	intHundreds := n / 100
	intTens := n / 10 % 10
	intUnits := n % 10
	strHundreds := ""
	if n == 1000 {
		return "one thousand"
	}
	if n > 100 {
		strHundreds = zeroTo9[intHundreds] + " hundred and "
	}
	if n%100 == 0 {
		strHundreds = zeroTo9[intHundreds] + " hundred"
	}

	strTens := zeroTo100[intTens]
	var strUnits string
	if intTens == 1 {
		strUnits = tenTo19[intUnits]
		strTens = ""
	} else if intTens == 0 || intUnits == 0 {
		strUnits = zeroTo9[intUnits]
	} else {
		strUnits = "-" + zeroTo9[intUnits]
	}
	return strHundreds + strTens + strUnits
}

func Challenge17() {
	text := ""
	count := 0
	for n := 1; n <= 1000; n++ {
		text += (intToText(n) + "\n")
	}

	for _, r := range text {
		if r != ' ' && r != '\n' && r != '-' {
			count++
		}
	}
	fmt.Print(text)
	fmt.Println(count)
}
