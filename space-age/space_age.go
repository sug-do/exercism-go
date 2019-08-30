package space

import "math"

const eY float64 = 31557600

// Age will take in the seconds and planet from the test and send the result back as float64
func Age(seconds float64, planet string) float64 {
	// map all values to one variable that we can reuse in the calculation
	planetSecs := map[string]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	// do calculation against the result of the planet (from the main function) returned back from map
	return (math.Round((seconds/(eY*planetSecs[planet]))*100) / 100)
}
