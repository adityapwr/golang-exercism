package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	switch speed {
	case 0:
		return 0.0
	case 1, 2, 3, 4:
		return 1.0
	case 5, 6, 7, 8:
		return 0.90
	case 9, 10:
		return 0.77
	}
	return 0.0
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return SuccessRate(speed) * float64(speed) * float64(221)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(SuccessRate(speed) * float64(speed) * float64(221) / float64(60))
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	if CalculateProductionRatePerHour(speed) > limit {
		return limit
	}
	return CalculateProductionRatePerHour(speed)
}
