package leetcode

// import "strconv"

// func palindromeNumber(x int) bool {
// 	if x < 0 {
// 		return false
// 	}

// 	if x < 10 {
// 		return true
// 	}

// 	if x%10 == 0 {
// 		return false
// 	}

// 	return reverseString(strconv.Itoa(x)) == strconv.Itoa(x)

// }

// func reverseString(s string) string {
// 	r := ""

// 	for i := len(s) - 1; i >= 0; i-- {
// 		r += string(s[i])
// 	}

// 	return string(r)
// }

func palindromeNumber(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	rev := 0

	for x > rev {
		rev = rev*10 + x%10

		x = x / 10
	}

	return x == rev || x == rev/10
}
