package utils

import (
	"math"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func GetDistanceBetweenCoordinates(from, to Coordinate) float64 {
	if from.Latitude == to.Latitude && from.Longitude == to.Longitude {
		return 0
	}

	fromRadian := (math.Pi * from.Latitude) / 180
	toRadian := (math.Pi * to.Latitude) / 180

	theta := from.Longitude - to.Longitude
	radTheta := (math.Pi * theta) / 180

	dist := math.Sin(fromRadian)*math.Sin(toRadian) +
		math.Cos(fromRadian)*math.Cos(toRadian)*math.Cos(radTheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = (dist * 180) / math.Pi
	dist = dist * 60 * 1.1515
	dist = dist * 1.609344

	return dist
}

// How to use:
// 	from := Coordinate{Latitude: 52.5200, Longitude: 13.4050} // Berlin
// 	to := Coordinate{Latitude: 48.8566, Longitude: 2.3522}    // Paris

// 	distance := GetDistanceBetweenCoordinates(from, to)
// 	println("Distância entre as coordenadas (km):", distance)
