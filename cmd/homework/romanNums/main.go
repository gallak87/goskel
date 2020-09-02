package main

import "fmt"

func main() {
	//fmt.Println(intToRoman(4))
	//fmt.Println(intToRoman(14))

	//fmt.Println(intToRoman(37))

	//fmt.Println(intToRoman(594))
	//fmt.Println(intToRoman(769))
	fmt.Println(romanToInt("XIX"))
	fmt.Println(romanToInt("VIII"))
	fmt.Println(romanToInt("XCIV"))
	fmt.Println(romanToInt("CDXCIV"))
}

var lastAsPref = []string{"IV","IX","XL", "XC", "CD", "CM"}

func isLastPairAPrefix(s, cur string) bool {
	s += cur
	for _, v := range lastAsPref {
		if s == v {
			return true
		}
	}
	return false
}

func romanToInt(s string) int {
	// XIX -> 19
	// X -> 10
	// VI -> 6
	// IV -> 4
	// XVIII -> 18
	// DXCIV --> 594
	// CDXCIV --> 494
	// CMXCIV --> 994
	// start from left, checking last As Pref each iteration
	output := 0

	last := ""

	for i := 0; i < len(s); i++ {
		v := string(s[i])
		n := iMap[v]
		if output > 0 && isLastPairAPrefix(last, v) {
			output -= iMap[last] * 2
		}
		output += n
		last = v
	}
	return output
}

var iMap = map[string]int {
	"I":1,
	"V":5,
	"X":10,
	"L":50,
	"C":100,
	"D":500,
	"M":1000,
}

/*

4 -> IV
14 -> XIV
51 -> LI
98 -> XCVIII
 */

var rMap = map[int]string{
	1: "I",
	5: "V",
	10: "X",
	50: "L",
	100: "C",
	500: "D",
	1000: "M",
}
var rList = []int{1, 5, 10, 50, 100, 500, 1000}

func intToRoman(i int) string {
	// 199 --> CXCIX
	// 9 -> IX
	// 9(0) -> XC
	// 1

	// 245
	// 5 -> V
	// 4(0) -> XL
	// 2(00) -> CC
	output := ""
	multiplier := 1
	// bottom up approach, aka in order of 1s, 10s, 100s, 1000s
	for next := i; next > 0; next /= 10 {
		// get remainder
		rem := next % 10
		pl := lowestKey(rem*multiplier)
		// convert remainder and append
		o := convRem(rem, pl, multiplier)
		output = o + output
		// next divides by 10, repeat with multiplier
		multiplier *= 10

	}
	return output
}

func convRem(rem, pl, mul int) string {
	// if we are at PL-1 (ex: 4, 9, 40, 90, 400, 900)
	// prepare the string with a prefix, and bump pL to next
	// index in list
	i := rem*mul
	// 4, 9 -> prev I
	// 40, 90 --> prev X
	if i == (rList[pl+1] - mul) {
		n1 := rMap[mul]
		n2 := rMap[rList[pl+1]]
		return n1 + n2
	}
	output := ""
	for {
		if i == 0 {
			break
		}
		l := lowestKey(i)
		output += rMap[rList[l]]
		i -= rList[l]
	}
	return output
}

func lowestKey(i int) int {
	l := len(rList)-1
	for {
		if rList[l] <= i || l == 0 {
			break
		}
		l--
	}
	return l
}