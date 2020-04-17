package capter8

import (
	"reflect"
)

func AssertEqual(expected, actual interface{}) bool{
	return reflect.DeepEqual(expected, actual)
}