package validator

import (
	"testing"
)

type testStruct struct {
	valid bool
}

func(this testStruct) Validate() bool {
	return this.valid
}

type noInterfaceimpl struct {}

func TestIsValid(t *testing.T) {
	var nilobj *noInterfaceimpl

	tests := []struct {
		name string
		toValidate interface{}
		valid bool
	}{
		{
			"Valid test w/validate method passing pointer",
			&testStruct{true},
			true,
		},
		{
			"Valid test w/validate method",
			testStruct{true},
			true,
		},
		{
			"Invalid test w/validate method, passing pointer",
			&testStruct{false},
			false,
		},
		{
			"Invalid test w/validate method",
			testStruct{false},
			false,
		},
		{
			"Valid test w/out validate method, passing pointer",
			&noInterfaceimpl{},
			true,
		},
		{
			"Valid test w/out validate method",
			noInterfaceimpl{},
			true,
		},
		{
			"Invalid nil test",
			nil,
			false,
		},
		{
			"Invalid nil obj test",
			nilobj,
			false,
		},
	}

	for _, test := range tests {
		if IsValid(test.toValidate) != test.valid {
			t.Errorf("Test [%s] Failed", test.name)
		}
	}
}

func BenchmarkIsValidValidWValidator(b *testing.B) {
	var testobj = testStruct{true}

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func BenchmarkIsValidValidWOValidator(b *testing.B) {
	var testobj = noInterfaceimpl{}

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func BenchmarkIsValidValidPtrWValidator(b *testing.B) {
	var testobj = &testStruct{true}

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func BenchmarkIsValidValidPtrWOValidator(b *testing.B) {
	var testobj = &noInterfaceimpl{}

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}
}

func BenchmarkIsValidNILNonInterface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if IsValid(nil) {
			b.Error("Valid struct, expected invalid")
		}
	}
}

func BenchmarkIsValidNIL(b *testing.B) {
	var testobj *noInterfaceimpl

	for n := 0; n < b.N; n++ {
		if IsValid(testobj) {
			b.Error("Valid struct, expected invalid")
		}
	}

}