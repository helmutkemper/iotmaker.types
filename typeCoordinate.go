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
	el.DensityValue = float64(el.OriginalValue) * el.DensityFactor
}

func (el *Coordinate) Add(value int) {
	if el.DensityFactor == 0 {
		el.DensityFactor = 1
	}

	el.OriginalValue += value
	el.DensityValue = float64(el.OriginalValue) * el.DensityFactor
}

func (el *Coordinate) Sub(value int) {
	if el.DensityFactor == 0 {
		el.DensityFactor = 1
	}

	el.OriginalValue -= value
	el.DensityValue = float64(el.OriginalValue) * el.DensityFactor
}

func (el *Coordinate) SetDensityFactor(value float64) {
	el.DensityFactor = value
	el.DensityValue = float64(el.OriginalValue) * el.DensityFactor
}

func (el Coordinate) Int() int {
	return int(el.DensityValue)
}

func (el Coordinate) Float64() float64 {
	return el.DensityValue
}

func (el Coordinate) Float32() float32 {
	return float32(el.DensityValue)
}

func (el Coordinate) Float() float64 {
	return el.DensityValue
}

func (el Coordinate) String() string {
	return strconv.FormatInt(int64(el.OriginalValue), 10)
}
