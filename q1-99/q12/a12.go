package q12

var thousandMap = map[int]string{1: "M", 2: "MM", 3: "MMM"}
var hundredMap = map[int]string{1: "C", 2: "CC", 3: "CCC", 4: "CD", 5: "D", 6: "DC", 7: "DCC", 8: "DCCC", 9: "CM", 0: ""}
var tenMap = map[int]string{1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC", 0: ""}
var unitMap = map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 0: ""}

// 输入确保在 1 到 3999 的范围内。
func intToRoman(num int) string {
	var roman string
	if num >= 1000 {
		roman = thousandMap[num%10000/1000]
	}
	if num >= 100 {
		roman += hundredMap[num%1000/100]
	}
	if num >= 10 {
		roman += tenMap[num%100/10]
	}
	roman += unitMap[num%10]

	return roman
}

// 抄来的,  更少的 内存消耗
func intToRomanCopied(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return M[num/1000] + C[num/100%10] + X[num/10%10] + I[num%10]
}

var M = []string{"M", "MM", "MMM"}
var C = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var X = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var I = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

func intToRomanV2(num int) string {
	var roman string
	if num >= 1000 {
		roman = M[num/1000]
	}
	if num >= 100 {
		roman += C[num/100%10]
	}
	if num >= 10 {
		roman += X[num/10%10]
	}
	roman += I[num%10]
	return roman
}
