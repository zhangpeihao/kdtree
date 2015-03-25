package kdtree

import (
	"math"
	"strconv"
)

type Coordinate struct {
	Values []float64
}

func (co *Coordinate) Dimensions() int {
	return len(co.Values)
}

func (co *Coordinate) String() (str string) {
	first := true
	for _, value := range co.Values {
		if first {
			str = "["
			first = false
		} else {
			str = str + ", "
		}
		str = str + strconv.FormatFloat(value, 'f', 5, 64)
	}
	str = str + "]"
	return
}

func (co *Coordinate) Equal(other *Coordinate) bool {
	if len(co.Values) != len(other.Values) {
		return false
	}
	for index, value := range co.Values {
		if value != other.Values[index] {
			return false
		}
	}
	return true
}

func (co *Coordinate) DistanceTo(other *Coordinate) float64 {
	dim := co.Dimensions()
	if dim != other.Dimensions() {
		return 0.0
	}
	var p, d float64
	for i := 0; i < dim; i++ {
		d = co.Values[i] - other.Values[i]
		p += (d * d)
	}
	//fmt.Printf("%s <-> %s = %.3f\n", a, b, math.Sqrt(p))
	return math.Sqrt(p)
}
