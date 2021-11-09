package examples

import (
	"reflect"
	"strings"
	"testing"

	csvParser "github.com/ElmTheDev/go-csv-config-parser"
)

func TestRequiredFields(t *testing.T) {
	type MyConfig struct {
		FirstName string `csv:"first name"`
		LastName  string `csv:"last name,required"`
		Age       int    `csv:"age"`
	}

	_, err := csvParser.ParseCSV("../test_csv.csv", reflect.TypeOf(MyConfig{}))
	if err == nil {
		t.Error("Expected error, got nil")
		t.Fail()
		return
	}

	if !strings.Contains(err.Error(), "required field") {
		t.Error("Expected error to contain 'required field', got", err.Error())
		t.Fail()
		return
	}

	t.Log("Required fields test passed")
}