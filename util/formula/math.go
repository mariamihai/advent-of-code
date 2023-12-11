package formula

// GCD - Greatest Common Divisor [CMMDC]
func GCD(nr1, nr2 int) int {
	for nr2 != 0 {
		temp := nr2
		nr2 = nr1 % nr2
		nr1 = temp
	}

	return nr1
}

// LCM - Least Common Multiple [CMMMC]
func LCM(nrs ...int) int {
	result := nrs[0] * nrs[1] / GCD(nrs[0], nrs[1])

	for i := 2; i < len(nrs); i++ {
		result = LCM(result, nrs[i])
	}

	return result
}
