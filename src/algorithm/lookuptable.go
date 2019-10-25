package algorithm

//import "fmt"
//import "math"

type LookupTable_t struct {

	DBKey string
}

func (o* LookupTable_t)Compute(currentTemperature float64) float64 {
	val := currentTemperature
	if val < 25 {
		return 20
	} else if (val >= 25) && (val < 35) {
		return 40
	} else {
		return 60
	}
}

