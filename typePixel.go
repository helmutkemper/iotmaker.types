package iotmaker_types

import (
	"strconv"
)

type Pixel int

func (el Pixel) String() string {
	return strconv.FormatInt(int64(el), 10)
}
