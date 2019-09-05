package geogo

import (
	"math"
	"testing"
)

const cPermissibleInaccuracy = 0.005                            // allow 0.5% inaccuracy
var coordsMinsk = CoordsPoint{Lat: 53.8847608, Lng: 27.4532865} // Minsk coordinates, center of the world :)

func TestShortDistance(t *testing.T) {
	const distanceNewEurope2McDonalds = 329 // actual distance between these places is 329 meters
	// Coordinated for short distance check
	var coordsMinskNewEurope = CoordsPoint{Lat: 53.9292509, Lng: 27.5811346}
	var coordsMinskMcDonalds = CoordsPoint{Lat: 53.9293472, Lng: 27.5797196}
	checkInaccuracy(t, coordsMinskNewEurope, coordsMinskMcDonalds, distanceNewEurope2McDonalds)
}

func TestDistanceMinskNewYork(t *testing.T) {
	const distanceMinsk2NewYork = 7124.0 // km
	var coordsNewYork = CoordsPoint{Lat: 40.6976701, Lng: -74.2598637}
	checkInaccuracy(t, coordsMinsk, coordsNewYork, distanceMinsk2NewYork*1000)

}

func TestDistanceMinskTokyo(t *testing.T) {
	const distanceMinsk2Tokyo = 8127.89 // km
	var coordsTokyo = CoordsPoint{Lat: 35.5062909, Lng: 138.6486705}
	checkInaccuracy(t, coordsMinsk, coordsTokyo, distanceMinsk2Tokyo*1000)
}

func checkInaccuracy(t *testing.T, point1 CoordsPoint, point2 CoordsPoint, actualDistance float64) {
	distance := Distance(point1, point2)
	difference := math.Abs(distance - actualDistance)
	inaccuracy := cPermissibleInaccuracy * actualDistance
	if difference > inaccuracy {
		t.Errorf("Expected difference less than '%.2f', actual difference is '%.2f'",
			inaccuracy, difference)
	}
}
