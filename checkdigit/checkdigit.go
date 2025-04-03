package checkdigit

// CalculateModulus10All calculates the check digit using the Luhn algorithm. supports weight 2 and 3.
// 末尾のcheckdigitを除いた配列を渡す
func CalculateModulus10All(weight int, digits []int) int {
	sum := 0
	calcWeight := weight

	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * calcWeight
		if calcWeight == weight {
			calcWeight = 1
		} else {
			calcWeight = weight
		}
	}

	modResult := sum % 10
	if modResult == 0 {
		return 0
	}
	return 10 - modResult
}

// ValidateModulus10All validates the check digit using the Luhn algorithm. supports weight 2 and 3.
func ValidateModulus10All(weight int, digits []int) bool {
	if len(digits) < 2 {
		return false
	}
	// 末尾を除いた配列でチェックディジットを計算
	inputDigits := digits[:len(digits)-1]
	lastDigit := digits[len(digits)-1]

	checkDigit := CalculateModulus10All(weight, inputDigits)
	return checkDigit == lastDigit
}
