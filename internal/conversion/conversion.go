package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats = make([]float64, len(strings))

	for i, s := range strings {
		float, err := strconv.ParseFloat(s, 64)

		if err != nil {
			return nil, fmt.Errorf("error converting string to float: %w", err)
		}
		floats[i] = float
	}

	return floats, nil
}
