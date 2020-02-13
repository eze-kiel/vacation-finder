package location

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Import(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output map[string]Coordinates
	}{
		{
			name:  "single",
			input: `paris: 99.99,1.475`,
			output: map[string]Coordinates{
				"paris": Coordinates{
					Latitude:  99.99,
					Longitude: 1.475,
				},
			},
		},
		{
			name: "dual",
			input: `slc: 1.2,3.4
someplace: 5.6,7.8
`,
			output: map[string]Coordinates{
				"slc": Coordinates{
					Latitude:  1.2,
					Longitude: 3.4,
				},
				"someplace": Coordinates{
					Latitude:  5.6,
					Longitude: 7.8,
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			b := strings.NewReader(tt.input)

			l, err := Import(b)

			assert.NoError(t, err, "error occured but shouldnt")
			assert.NotEmpty(t, l, "no locations returned")
			assert.Equal(t, tt.output, l, "result doesn't match expectations")
		})

	}
}
