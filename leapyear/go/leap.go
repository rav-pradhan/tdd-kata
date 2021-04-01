package leap

func IsLeapYear(year int) bool {
	return divisibleBy(4, year) && !(divisibleBy(100, year) && notDivisibleBy(400, year))
}

func divisibleBy(number, year int) bool {
	return year%number == 0
}

func notDivisibleBy(number, year int) bool {
	return year%number != 0
}
