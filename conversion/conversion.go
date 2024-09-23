package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloats(vals []string) ([]float64, error) {
	var floats = make([]float64, 0)
	for _, val := range vals {
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Printf("Unable to parse %s float64 value\n", val, err.Error())
			return nil, err
		}
		floats = append(floats, f)
	}

	return floats, nil
}
