package csvParser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"

	"github.com/jszwec/csvutil"
)

// parses CSV file and returns a slice of interfaces and erro
func ParseCSV(path string, input reflect.Type) (interface{}, error) {
	// Read file from path
	fileData, err := readFile(path)
	if err != nil {
		return nil, err
	}

	// Find all required fields and store them in array
	requiredFields := []string{}
	numberOfFields := input.NumField()
	for i := 0; i < numberOfFields; i++ {
		field := input.Field(i)
		if strings.Contains(string(field.Tag), "required") {
			requiredFields = append(requiredFields, string(field.Name))
		}
	}

	// Create new decoder
	dec, err := csvutil.NewDecoder(csv.NewReader(strings.NewReader(fileData)))
	if err != nil {
		return nil, err
	}

	// Loop through all rows and parse them into slice
	var rows []interface{}
	for {
		u := reflect.New(input).Interface()

		if err := dec.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		rows = append(rows, reflect.ValueOf(u).Interface())
	}

	// Check if all required fields are filled
	if len(requiredFields) > 0 {
		for _, entry := range rows {
			for _, required := range requiredFields {
				if reflect.ValueOf(entry).Elem().FieldByName(required).String() == "" {
					return nil, fmt.Errorf("required field %s is empty", required)
				}
			}
		}	
	}

	return interfaceArrayToTypeArray(rows, input), nil
}