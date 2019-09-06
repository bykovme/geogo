package main

import (
	"fmt"
	"github.com/bykovme/geogo"
)

func main() {
	coordsTokyo := geogo.CoordsPoint{Lat: 35.5062909, Lng: 138.6486705}
	coordsMinsk := geogo.CoordsPoint{Lat: 53.8847608, Lng: 27.4532865}

	distance := geogo.Distance(coordsTokyo, coordsMinsk)

	fmt.Printf("Distance between Minsk and Tokyo: %.2f meters", distance)
}
