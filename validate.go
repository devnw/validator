// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

// Package validator sets up several methods for simple validation as well
// as setting up an interface which when implemented allows for custom validation
// logic by the implementor. `Valid()` returns a boolean indicating if the value(s)
// passed into it are valid based on the different types. `Assert()` returns an error
// indicating `nil` if the values are valid. Otherwise `Assert()` returns an error
// indicating the index of the value(s) passed into it as to which was determined
// to be invalid.
package validator // import "github.com/devnw/validator"

import (
	"reflect"
	"strconv"
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
		return ValidationError{
			Message: "empty argument list passed to assert",
			Index:   0,
		}
	}

	for i, obj := range objs {
		switch v := obj.(type) {
		case nil:
			return ValidationError{
				Message: "nil value",
				Index:   i,
				Type:    reflect.TypeOf(i),
			}
		case validator:
			if !v.Validate() {
				return ValidationError{
					Message:          "nil value",
					Index:            i,
					Type:             reflect.TypeOf(obj),
					ValidatorFailure: true,
				}
			}
		case string:
			if v == "" {
				return ValidationError{
					Message: "empty string",
					Index:   i,
					Type:    reflect.TypeOf(obj),
				}
			}
		case []byte:
			if len(v) == 0 {
				return ValidationError{
					Message: "empty slice",
					Index:   i,
					Type:    reflect.TypeOf(obj),
				}
			}
		case []string:
			for j, s := range v {
				if s == "" {
					return ValidationError{
						Message: "empty string in slice at index " + strconv.Itoa(j),
						Index:   i,
						Type:    reflect.TypeOf(obj),
					}
				}
			}
		case int, int8, int16, int32, int64, uint, uint8,
			uint16, uint32, uint64, uintptr, bool,
			float32, float64, complex128, complex64:
			continue
		default:
			if !valid(obj) { // TODO: return index of error
				return ValidationError{
					Message: "invalid at reflection check",
					Index:   i,
					Type:    reflect.TypeOf(obj),
				}
			}
		}
	}

	return nil
}

// IsValid is DEPRECATED, replaced by Valid. This method will be available until version 1.1.0
func IsValid(objs ...interface{}) bool {
	return Valid(objs...)
}

// Valid reads in an object, and checks to see if the object implements the validator interface
// if the object does then it executes the objects validate method and returns that
func Valid(objs ...interface{}) bool {
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
		case int, int8, int16, int32, int64, uint, uint8,
			uint16, uint32, uint64, uintptr, bool,
			float32, float64, complex128, complex64:
			continue
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
