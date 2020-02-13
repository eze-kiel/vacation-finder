/*
	Author:	ezekiel
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/eze-kiel/vacation-finder/location"
)

const EarthCirc = 40000
const LatLength = 111

func main() {
	var place string
	var dist int64
	var n int
	var lat float64
	var long float64

	flag.StringVar(&place, "place", "slc", "the name of your starting point")
	flag.Int64Var(&dist, "dist", 100, "the distance in kilometers from your start point")
	flag.IntVar(&n, "n", 1, "number of points generated")

	// useless for now
	flag.Float64Var(&lat, "lat", 0, "your starting latitude")
	flag.Float64Var(&long, "long", 0, "your starting longitude")
	//

	flag.Parse()

	f, err := os.Open("locations.yml")
	if err != nil {
		log.Fatal(err)
	}

	places, err := location.Import(f)
	if err != nil {
		log.Fatal(err)
	}

	for x := 0; x < n; x++ {

		fmt.Printf("\nPoint #%d/%d :\n", x+1, n)
		randomPoint := generateRandomPoint(int(dist))

		convertedPoint := kmToCoordinates(randomPoint, places[place])
		fmt.Printf("%d max from %s\n", dist, place)
		fmt.Printf("%f km ver and %f km hor\n", randomPoint.Latitude, randomPoint.Longitude)
		fmt.Printf("%s : %f %f\n", place, places[place].Latitude, places[place].Longitude)
		fmt.Printf("Converted point : lat = %f && long = %f\n", convertedPoint.Latitude, convertedPoint.Longitude)
		fmt.Printf("%f %f\n", convertedPoint.Latitude, convertedPoint.Longitude)
		fmt.Printf("https://www.google.fr/maps/@%f,%f,12z\n", convertedPoint.Latitude, convertedPoint.Longitude)
	}
}

func checkErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateRandomPoint(dist int) (newPoint location.Coordinates) {
	rand.Seed(time.Now().UnixNano())
	y := rand.Intn((dist * 2) + 1)
	y = dist - y

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn((dist * 2) + 1)
	x = dist - x

	newPoint = location.Coordinates{Latitude: float64(y), Longitude: float64(x)}
	return newPoint
}

func kmToCoordinates(point location.Coordinates, place location.Coordinates) (RealCoordinates location.Coordinates) {
	LongLength := EarthCirc * math.Cos(place.Latitude*math.Pi/180) / 360
	RealCoordinates.Latitude = place.Latitude + (point.Latitude / LatLength)
	RealCoordinates.Longitude = place.Longitude + (point.Longitude / LongLength)

	return RealCoordinates
}
