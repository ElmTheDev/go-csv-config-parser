package examples

import (
	"fmt"
	"reflect"
	"testing"

	csvParser "github.com/ElmTheDev/go-csv-config-parser"
)

func TestRegularFields(t *testing.T) {
	type MyConfig struct {
		FirstName string `csv:"first name"`
		LastName  string `csv:"last name"`
		Age       int    `csv:"age"`
	}

	result, err := csvParser.ParseCSV("../test_csv.csv", reflect.TypeOf(MyConfig{}))
	if err != nil {
		panic(err)
	}

	resultArr := result.([]MyConfig)

	validData := fmt.Sprintf("%v", resultArr)
	if validData != "[{elm dev 19} {john  24}]" {
		t.Errorf("Expected %v, got %v", "[{elm dev 19} {john  24}]", validData)
		t.Fail()
		return
	}

	t.Log("Regular fields test passed")
}