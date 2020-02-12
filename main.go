/*
	Author:	ezekiel
*/

package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const EarthCirc = 40000
const LatLength = 111

type Coordinates struct {
	Lat, Long float64
}

var location map[string]Coordinates

func main() {
	dist, err := strconv.ParseInt(os.Args[1], 10, 64)
	checkErrors(err)
	place := os.Args[2]

	location = make(map[string]Coordinates)

	location["slc"] = Coordinates{
		45.7386,
		4.4637,
	}

	randomPoint := generateRandomPoint(int(dist))

	convertedPoint := kmToCoordinates(randomPoint, location[place])
	fmt.Printf("%d max from %s\n", dist, place)
	fmt.Printf("%f km ver and %f km hor\n", randomPoint.Lat, randomPoint.Long)
	fmt.Printf("%s : %f %f\n", place, location[place].Lat, location[place].Long)
	fmt.Printf("Converted point : lat = %f && long = %f\n", convertedPoint.Lat, convertedPoint.Long)
	fmt.Printf("%f %f\n", convertedPoint.Lat, convertedPoint.Long)
	fmt.Printf("https://www.google.fr/maps/@%f,%f,12z\n", convertedPoint.Lat, convertedPoint.Long)
}

func checkErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateRandomPoint(dist int) (newPoint Coordinates) {
	rand.Seed(time.Now().UnixNano())
	y := rand.Intn(dist*2-1+1) + 1
	if y <= 150 {
		y = -y
	}

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(dist*2-1+1) + 1
	if x <= 150 {
		x = -x
	}

	newPoint = Coordinates{float64(y), float64(x)}
	return newPoint
}

func kmToCoordinates(point Coordinates, place Coordinates) (RealCoordinates Coordinates) {
	LongLength := EarthCirc * math.Cos(place.Lat*math.Pi/180) / 360
	RealCoordinates.Lat = place.Lat + (point.Lat / LatLength)
	RealCoordinates.Long = place.Long + (point.Long / LongLength)

	return RealCoordinates
}
