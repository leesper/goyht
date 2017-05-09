package goyht

import (
	"fmt"
	"reflect"
)

func toMap(st interface{}, extras map[string]string) (map[string]string, error) {
	val := reflect.ValueOf(st)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("need a struct type, got %T", st)
	}

	typ := val.Type()
	result := map[string]string{}

	for i := 0; i < val.NumField(); i++ {
		sf := typ.Field(i)
		if tag, ok := sf.Tag.Lookup("param"); ok && tag != "" {
			result[tag] = val.Field(i).String()
		}
	}

	for k, v := range extras {
		result[k] = v
	}
	return result, nil
}

func checkErr(code, subcode int, message string) error {
	const success = 200
	if code != success || subcode != success {
		return fmt.Errorf("code %d subcode %d msg %s", code, subcode, message)
	}
	return nil
}
