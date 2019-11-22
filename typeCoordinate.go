package iotmaker_types

import (
	"strconv"
)

type Coordinate struct {
	OriginalValue int
	DensityFactor float64
	DensityValue  float64
}

func (el *Coordinate) Set(value int) {
	if el.DensityFactor == 0 {
		el.DensityFactor = 1
	}

	el.OriginalValue = value
	el.DensityValue = float64(value) * el.DensityFactor
}

func (el *Coordinate) SetDensityFactor(value float64) {
	el.DensityFactor = value
}

func (el Coordinate) Int() int {

}

func (el Coordinate) String() string {
	return strconv.FormatInt(int64(el.OriginalValue), 10)
}
