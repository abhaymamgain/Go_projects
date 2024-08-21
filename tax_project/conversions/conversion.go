package conversions

import (
	"fmt"
	"strconv"
)

func StringsToFloat(strs []string) ([]float64, error) {
	floats := make([]float64, len(strs))

	for i, str := range strs {
		float, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Print("error parsing float")
		}
		floats[i] = float

	}
	return floats, nil
}
