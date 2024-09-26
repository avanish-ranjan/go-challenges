package Controllers

import (
	"fmt"
	"unicode"
)

// Diners club starts with 36, 38, 39 and length 14
func ValidateCard(cardNumber string) bool {
	if len(cardNumber) == 15 || len(cardNumber) == 16 {
		if cardNumber[0] == '4' || cardNumber[0] == '3' {
			for _, num := range cardNumber {
				if unicode.IsDigit(num) {
					return true
				}
			}
		}
	}

	fmt.Println(cardNumber)
	return false
}

func ValidateCardV2(cardNumber string) bool {
	// acceptedLength := []string{"15", "16", "14"}
	// acceptedPrefix := []string{"4", "3", "36", "38", "39"}
	// acceptedExpression := `^[0-9]{15}`

	fmt.Println(cardNumber)
	return false
}
