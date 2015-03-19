package leap

const TestVersion = 1

func IsLeapYear(year int) bool {
	return (DivisibleByFour(year) && !IsCentury(year)) || IsFourthCentury(year)
}

func DivisibleByFour(year int) bool {
	return year%4 == 0
}

func IsCentury(year int) bool {
	return year%100 == 0
}

func IsFourthCentury(year int) bool {
	return year%400 == 0
}
