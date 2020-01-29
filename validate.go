package validator

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

// validator interface for checking to see if the struct has a Validate method assigned to it
// NOTE: This implementation should not use a pointer receiver, otherwise it will not work for non-pointer instances
// of structs
type validator interface {
	Validate() (valid bool)
}

// Assert determines the validity of each value passed in to determine
// if the inputs are valid. A non-nil error return indicates an invalid
// assertion and a nil error returns indicates validity
func Assert(objs ...interface{}) error {
	if len(objs) == 0 {
		return errors.New("empty list passed to asert")
	}

	for i, obj := range objs {
		msg := fmt.Sprintf("parameter at index [%v] is invalid", i)
		switch v := obj.(type) {
		case nil:
			return errors.Errorf("%s | [nil]", msg)
		case validator:
			if !v.Validate() {
				return errors.Errorf("%s | [validator failed]", msg)
			}
		case string:
			if v == "" {
				return errors.Errorf("%s | [empty string]", msg)
			}
		case []byte:
			if len(v) == 0 {
				return errors.Errorf("%s | [empty []byte]", msg)
			}
		case []string:
			for j, s := range v {
				if s == "" {
					return errors.Errorf("%s | [empty string at index [%v] in []string]", msg, j)
				}
			}
		case int:
		case int8:
		case int16:
		case int32:
		case int64:
		case uint:
		case uint8:
		case uint16:
		case uint32:
		case uint64:
		case bool:
		case float32:
		case float64:
		default:
			if !valid(obj) {
				return errors.New(msg)
			}
		}
	}

	return nil
}

// IsValid reads in an object, and checks to see if the object implements the validator interface
// if the object does then it executes the objects validate method and returns that
func IsValid(objs ...interface{}) bool {
	if len(objs) == 0 {
		return false
	}

	for _, obj := range objs {
		switch v := obj.(type) {
		case nil:
			return false
		case validator:
			if !v.Validate() {
				return false
			}
		case string:
			if v == "" {
				return false
			}
		case []byte:
			if len(v) == 0 {
				return false
			}
		case []string:
			for _, s := range v {
				if s == "" {
					return false
				}
			}
		case int:
		case int8:
		case int16:
		case int32:
		case int64:
		case uint:
		case uint8:
		case uint16:
		case uint32:
		case uint64:
		case bool:
		case float32:
		case float64:
		default:
			if !valid(v) {
				return false
			}
		}
	}

	return true
}

func valid(obj interface{}) (valid bool) {

	// Using reflection pull the type associated with the object that is passed in. nil types are invalid.
	var tp reflect.Type
	if tp = reflect.TypeOf(obj); tp != nil {

		val := reflect.ValueOf(obj)

		// determine if the value is a pointer or not and whether it's nil if it is a pointer
		if val.Kind() != reflect.Ptr || (val.Kind() == reflect.Ptr && !val.IsNil()) {
			valid = true
		}
	}

	return valid
}
