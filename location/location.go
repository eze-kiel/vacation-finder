package location

import (
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

	err = yaml.Unmarshal(b, slocations)
	if err != nil {
		return nil, err
	}

	locations := map[string]Coordinates{}

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

		locations[k] = Coordinates{
			Latitude:  lat,
			Longitude: lon,
		}
	}
	return locations, nil
}
