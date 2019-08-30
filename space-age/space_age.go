package space

import (
	"math"
)

// take in the seconds and planet from the test and send the result back as float64
func Age(seconds float64, planet string) float64 {
	// declare the variable that we will be dividing by
	var planet_secs float64
	// set the number of seconds for Earth year on each planet
	if planet == "Earth" {
		planet_secs = 31557600
	} else if planet == "Mercury" {
		planet_secs = 7600543.81992
	} else if planet == "Venus" {
		planet_secs = 19414149.052176
	} else if planet == "Mars" {
		planet_secs = 59354032.69008
	} else if planet == "Jupiter" {
		planet_secs = 374355659.124
	} else if planet == "Saturn" {
		planet_secs = 929292362.884
	} else if planet == "Uranus" {
		planet_secs = 2651370019.3296
	} else if planet == "Neptune" {
		planet_secs = 5200418560.032
	}
	// do the calculation of seconds alive divided by number of seconds on that planet equal to Earth year, then round it
	return (rounding(seconds / planet_secs))
}

// round the result of the calculation to the nearest hundredth and pass the value back as float64
func rounding(in_round float64) float64 {
	return (math.Round(in_round*100) / 100)
}
