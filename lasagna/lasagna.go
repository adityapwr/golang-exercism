package lasagna

// TODO: define the 'OvenTime()' function
func OvenTime() int {
	time := 40
	return time
}

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(elapsed int) int {
	totalOvenTime := OvenTime()
	return totalOvenTime - elapsed
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(totalLayers int) (totalTime int) {
	return totalLayers * 2
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(totalLayers int, timeInOven int) (elapsedTime int) {
	return totalLayers*2 + timeInOven
}
