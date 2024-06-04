package romantoint

// Leetcode #13 - Convert roman numeral to int
/* Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000 */

func romanToInt(s string) int {
	result := 0
	prev_num := 0
	for _, numeral := range s {

		num := convert((numeral))
		if prev_num != 0 && prev_num < num {
			result = result + num - 2*prev_num
		} else {
			result += num
		}

		prev_num = num
	}
	return result
}

func convert(s string) int {
	num := 0
	if s == "I" {
		num = 1
	} else if s == "V" {
		num = 5
	} else if s == "X" {
		num = 10
	} else if s == "L" {
		num = 50
	} else if s == "C" {
		num = 100
	} else if s == "D" {
		num = 500
	} else if s == "M" {
		num = 1000
	}
	return num
}
