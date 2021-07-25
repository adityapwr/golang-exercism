//Package space tell how old someone would be on earth
package space

type Planet string

var orbitEarthYears = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

//Age takes current age of someone on planet and returns age on earth
func Age(seconds float64, planet Planet) (age float64) {
	const earthSecondsPerYear float64 = 31557600
	period, ok := orbitEarthYears[planet]
	if !ok {
		panic("Planet doesn't exists")
	}
	return seconds / (period * earthSecondsPerYear)
}
