package data

import (
	"fmt"
	"reflect"
)

const (
	isRequiredError = "%s is required"
	minimumError    = "%s must be greater than %d"
)

type Validator interface {
	Valid() error
}

func requiredFields(f reflect.StructField, v reflect.Value) error {
	if f.Tag.Get("required") == "true" {

		if v.IsZero() {
			return fmt.Errorf(isRequiredError, f.Name)
		}
	}
	return nil
}

func minFields(f reflect.StructField, v reflect.Value) error {
	if min := f.Tag.Get("min"); min != "" {
		minInt := 0
		if _, err := fmt.Sscanf(min, "%d", &minInt); err != nil {
			return err
		}

		if v.Int() < int64(minInt) {
			return fmt.Errorf(minimumError, f.Name, minInt)
		}
	}
	return nil
}

func readStructFields(s interface{}) error {
	v := reflect.ValueOf(s).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		vf := v.Field(i)
		if tag := f.Tag.Get("validate"); tag == "" {
			continue
		}

		if err := requiredFields(f, vf); err != nil {
			return err
		}

		if err := minFields(f, vf); err != nil {
			return err
		}
	}
	return nil
}

func Valid(s interface{}) error {
	return readStructFields(s)
}
