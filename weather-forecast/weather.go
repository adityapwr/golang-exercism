// Package weather holds data for current condiotn and location. It returns a log of condition and weather.
package weather

var (
	// CurrentCondition holds data for current weather condition
	CurrentCondition string
	// CurrentLocation holds data for current selected location
	CurrentLocation string
)

// Log return a formatted string consiting of city and condition
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
