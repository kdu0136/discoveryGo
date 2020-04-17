package capter8

import "reflect"

// NewMap returns given type (key, value) map
func NewMap(key, value interface{}) interface{} {
	keyType := reflect.TypeOf(key)
	valueType := reflect.TypeOf(value)
	mapType := reflect.MapOf(keyType, valueType)
	mapValue := reflect.MakeMap(mapType)
	return mapValue.Interface()
}

func FieldNames(s interface{}) ([]string, error) {

}
