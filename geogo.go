package geogo

import "math"

// cEarthRadius - average radius of earth in meters 6 371 000 m
// According to https://en.wikipedia.org/wiki/Earth_radius

const cEarthRadius = 6371000.0

type CoordsPoint struct {
	Lat float64
	Lng float64
}

// Distance - find distance in meters between GPS coordinates
func Distance(point1 CoordsPoint, point2 CoordsPoint) float64 {

	lat2 := radians(point2.Lat)
	lat1 := radians(point1.Lat)
	latDelta := radians(point2.Lat - point1.Lat)
	lngDelta := radians(point2.Lng - point1.Lng)

	// Using Haversine formula taken from here: https://www.movable-type.co.uk/scripts/latlong.html

	a := math.Sin(latDelta/2)*math.Sin(latDelta/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(lngDelta/2)*math.Sin(lngDelta/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return cEarthRadius * c
}
