package ex


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRateDecimal := float64(successRate) / 100
	return float64(productionRate) * successRateDecimal
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionRateInMinutes := float64(productionRate) / 60
	successRateDecimal := float64(successRate) / 100
	return int(productionRateInMinutes * successRateDecimal)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var cost uint

	carsCountDividedBy10 := carsCount / 10
	cost = uint(carsCountDividedBy10 * 95000)

	carsCountRemainder := carsCount % 10
	cost += uint(carsCountRemainder * 10000)

	return cost
}
