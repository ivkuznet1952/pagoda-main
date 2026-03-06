package ui

func MonthName(monthInt int) string {
	switch monthInt {
	case 1:
		return "Январь"
	case 2:
		return "Февраль"
	case 3:
		return "Март"
	case 4:
		return "Апрель"
	case 5:
		return "Май"
	case 6:
		return "Июнь"
	case 7:
		return "Июль"
	case 8:
		return "Август"
	case 9:
		return "Сентябрь"
	case 10:
		return "Октябрь"
	case 11:
		return "Ноябрь"
	case 12:
		return "Декабрь"
	default:
		return ""
	}

}

// Filter returns a new slice containing only the elements of the input slice that satisfy the predicate function.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
