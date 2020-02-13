package location

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// Coordinates contains latitude & longitude
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Import from a reader containing yaml
func Import(r io.Reader) (map[string]Coordinates, error) {
	b, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	type simpleLoc struct {
		key string
		loc string
	}

	slocations := map[string]string{}
	// var slocations []simpleLoc

	fmt.Printf("parsing yaml: %s\n", b)
	err = yaml.Unmarshal(b, slocations)
	if err != nil {
		return nil, err
	}

	locations := map[string]Coordinates{}

	//For key, value in the simple locations array, the coords are splitted at the coma and insered in a string array called coords.
	//Next the spaces are trimed in order to keep only the values
	for k, v := range slocations {

		coords := strings.Split(v, ",")
		coords[0] = strings.TrimSpace(coords[0])
		coords[1] = strings.TrimSpace(coords[1])

		lat, err := strconv.ParseFloat(coords[0], 64)
		if err != nil {
			return nil, err
		}

		lon, err := strconv.ParseFloat(coords[1], 64)
		if err != nil {
			return nil, err
		}
		//k is set as the key of the returned map
		locations[k] = Coordinates{
			Latitude:  lat,
			Longitude: lon,
		}
	}
	return locations, nil
}
