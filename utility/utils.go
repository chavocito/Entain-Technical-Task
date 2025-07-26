package utility

import "slices"

var validHeaders = []string{
	"game",
	"server",
	"payment",
	"slices",
}

func IsHeaderValid(header string) bool {
	return slices.Contains(validHeaders, header)
}

func IsStateValid(state string) bool {
	return state == "win" || state == "lose"
}

func ConvertAmountToDBRepresentation(originalAmount float64) int64 {
	return int64(originalAmount * 100.00)
}

func ConvertAmountToDecimalForDisplay(intAmount int64) float64 {
	return float64(intAmount) / 100.00
}
