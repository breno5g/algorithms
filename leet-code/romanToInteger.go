package leetcode

func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var sum int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman[string(s[i])] < roman[string(s[i+1])] {
			sum -= roman[string(s[i])]
		} else {
			sum += roman[string(s[i])]
		}
	}
	return sum
}
