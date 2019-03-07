# Validator

[![GoDoc](https://godoc.org/github.com/benji-vesterby/validator?status.svg)](https://godoc.org/github.com/benji-vesterby/validator)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


Validator can be executed against any struct / interface to determine if it is valid. 

The standard check executes a "nil" check against the value being passed into the validator. It then uses reflection to determine if the object is a pointer or not and whether the pointer is nil.

If it's a pointer and the value is nil then it returns invalid, if it's not a pointer or the value is not nil then it checks to see if the type of the object implements the Validator interface. If the object does not implement the interface then it returns valid because the object is non-nil. If it does implement the validator interface then it executes the validate method on the object after typecasting it. At this point the IsValid method will return the value returned from the Validate method.

To install: 

DEP

`dep ensure -add github.com/benji-vesterby/validator`

Go Get

`go get github.com/benji-vesterby/validator`

To use:

    import "github.com/benji-vesterby/validator"
    
    type testStruct struct {
    	valid bool
    }
    
    // Implement the validate method on your struct
    func(this testStruct) Validate() bool {
    	return this.valid
    }
    
    // This accepts an interface. If the interface passed to it is valid it returns true,
    // otherwise it returns false. 
    validator.IsValid(testStruct{true})
    
See unit tests for more examples

https://github.com/benji-vesterby/validator/blob/master/validate_test.go