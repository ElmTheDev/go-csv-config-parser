package csvParser

import (
	"io/ioutil"
	"reflect"
)

// reads file and returns contents as string
func readFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// converts interface array to regular interface through Reflect
func interfaceArrayToTypeArray(arr []interface{}, t reflect.Type) interface{} {
	result := reflect.SliceOf(t)
	arrVal := reflect.MakeSlice(result, 0, 0)

	for _, v := range arr {
		if reflect.TypeOf(v).Kind() == reflect.Ptr {
			arrVal = reflect.Append(arrVal, reflect.ValueOf(v).Elem())
		} else {
			arrVal = reflect.Append(arrVal, reflect.ValueOf(v))
		}
	}

	return arrVal.Interface()
}