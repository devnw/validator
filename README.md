# Validator

[![CI](https://github.com/benjivesterby/validator/workflows/CI/badge.svg)](https://github.com/benjivesterby/validator/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/benjivesterby/validator)](https://goreportcard.com/report/github.com/benjivesterby/validator)
[![codecov](https://codecov.io/gh/benjivesterby/validator/branch/master/graph/badge.svg)](https://codecov.io/gh/benjivesterby/validator)
[![GoDoc](https://godoc.org/github.com/benjivesterby/validator?status.svg)](https://pkg.go.dev/github.com/benjivesterby/validator)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

Validator can be executed against any struct / interface to determine if it is valid.

To install:

`go get -u github.com/benjivesterby/validator`

To use validator

1. import `github.com/benjivesterby/validator`
2. execute boolean validation: `validator.Valid(obj1, obj2, ..., objN)`
3. execute validation assertion: `validator.Assert(obj1, obj2, ..., objN)`

`Valid()` returns a boolean indicating validity

`Assert()` returns a *nil* error if the inputs are valid, and returns an error for invalid arguments, and the error specifies the index of the erroneous value

To implement your own validator use the interface method `Validate() bool` as shown below:

```go
import "github.com/benjivesterby/validator"

type testStruct struct {
    valid bool
}

// Implement the validate method on your struct
func(this testStruct) Validate() bool {
    return this.valid
}

// This accepts an interface. If the interface passed to it is valid it returns true,
// otherwise it returns false.
validator.Valid(testStruct{true})
```

This libary will check for *nil* first on any nillable type, then it uses a type swtich to check for validity on known types.

* For slices it will indicate which element of the slice that is passed in is invalid when using `Assert`
* Valid will also check individual slice indexes but will not indicate which is invalid since it only returns a `bool`
