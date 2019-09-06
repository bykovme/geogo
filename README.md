[![Build Status](https://travis-ci.com/bykovme/geogo.svg?branch=master)](https://travis-ci.com/bykovme/geogo)
# GeoGo
Go (golang) library for distance calculation

## Installation

Install the package with the following command line

```
go get github.com/bykovme/geogo
```

## Quick documentation

There are 2 function

```
geogo.Distance
geogo.ShortDistance
```

Function **Distance** calculates the distance between any coordinates using [Haversine](https://en.wikipedia.org/wiki/Haversine_formula) formula

**ShortDistance** calculates the distance using usual [Cartesian](https://en.wikipedia.org/wiki/Cartesian_coordinate_system) formula but it should be more precise on short ranges

Both functions return result in **meters**

Depending on the coordinates the function **Distance** decides on its own if long range or short range calculation should be used.
By default the short range calculation works for distances less than 5km (5000 meters).

## Usage example 
This example is located within the package here: 


```go
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
```
