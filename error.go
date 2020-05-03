// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package validator

import (
	"reflect"
	"strconv"
)

// ValidationError is a custom error type for the
// validator library which provides additional
// information when using the Assert methods for
// validating input values. This error contains
// more information pertinent to what the errors
// were when attempting to validation the arguments.
type ValidationError struct {

	// Message contains the custom message
	// for this validation error
	Message string

	// Index is the argument index which was
	// invalid during validation
	Index int

	// Type is the type of the invalid argument
	// at the index specified
	Type reflect.Type

	// ValidatorFailure indicates if the failure
	// was at the level which executes the Validate
	// method of a struct implementing the validator
	// interface which indicates it was in the users
	// validation logic where the failure occurred
	ValidatorFailure bool
}

// Error implements the error interface for
// validation errors to be returned instead
// of standard lib errors
func (e ValidationError) Error() string {
	return e.String()
}

// String implements the stringer interface
// so that the error can be properly printed
// by the fmt printers, etc...
func (e ValidationError) String() string {

	arg := "argument at index [" + strconv.Itoa(e.Index)

	if e.Type != nil {
		arg += "] | type [" + e.Type.String()
	}

	arg += "] is invalid"

	msg := ""
	if e.Message != "" {
		msg = " -" + e.Message + "-"
	}

	return "validation error" + msg + " " + arg
}
