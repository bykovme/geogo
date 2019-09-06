package geogo

import "math"

// cEarthRadius - average radius of earth in meters 6 371 000 m
// According to https://en.wikipedia.org/wiki/Earth_radius

const cEarthRadius = 6371000.0
const cShortRadiusLimit = 5000.0

var ShortRadiusLimit float64

func init() {
	ShortRadiusLimit = cShortRadiusLimit
}

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

	// Using Haversine formula

	a := math.Sin(latDelta/2)*math.Sin(latDelta/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(lngDelta/2)*math.Sin(lngDelta/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := cEarthRadius * c

	// If distance is small, use more precise alg considering the Earth is flat :)
	if distance < ShortRadiusLimit {
		return ShortDistance(point1, point2)
	}

	return distance
}

// ShortDistance - calculate short range distance
func ShortDistance(point1 CoordsPoint, point2 CoordsPoint) float64 {

	lat1 := radians(point1.Lat)
	latDelta := radians(point2.Lat - point1.Lat)
	lngDelta := radians(point2.Lng - point1.Lng)

	// Using Cartesian formula
	deltaX := cEarthRadius * lngDelta * math.Cos(lat1)
	deltaY := cEarthRadius * latDelta

	distance := math.Sqrt(deltaX*deltaX + deltaY*deltaY)
	return distance
}
