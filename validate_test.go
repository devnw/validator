package validator

import (
	"testing"
)

type testStruct struct {
	valid bool
}

func (testobj testStruct) Validate() bool {
	return testobj.valid
}

type noInterfaceimpl struct{}

func TestIsValidNIL(t *testing.T) {
	if IsValid(nil) != false {
		t.Errorf("Test Failed")
	}
}

func TestIsValidEmpty(t *testing.T) {
	if IsValid() != false {
		t.Errorf("Test Failed")
	}
}

func TestIsValidString(t *testing.T) {
	tests := []struct {
		name       string
		toValidate string
		valid      bool
	}{
		{
			"Valid string Test",
			"testify",
			true,
		},
		{
			"Invalid string Test",
			"",
			false,
		},
	}

	for _, test := range tests {
		if IsValid(test.toValidate) != test.valid {
			t.Errorf("Test [%s] Failed", test.name)
		}
	}
}

func TestIsValidStringSlice(t *testing.T) {
	tests := []struct {
		name       string
		toValidate []string
		valid      bool
	}{
		{
			"Valid string Test",
			[]string{"testify", "testify"},
			true,
		},
		{
			"Invalid string Test",
			[]string{"testify", ""},
			false,
		},
	}

	for _, test := range tests {
		if IsValid(test.toValidate) != test.valid {
			t.Errorf("Test [%s] Failed", test.name)
		}
	}
}

func TestIsValidByteSlice(t *testing.T) {
	tests := []struct {
		name       string
		toValidate []byte
		valid      bool
	}{
		{
			"Valid []byte Test",
			[]byte{12, 12, 12},
			true,
		},
		{
			"Invalid []byte Test",
			[]byte{},
			false,
		},
	}

	for _, test := range tests {
		if IsValid(test.toValidate) != test.valid {
			t.Errorf("Test [%s] Failed", test.name)
		}
	}
}

func TestIsValid(t *testing.T) {
	var nilobj *noInterfaceimpl

	tests := []struct {
		name       string
		toValidate interface{}
		valid      bool
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

func TestIsValid_Multiples(t *testing.T) {
	var nilobj *noInterfaceimpl

	tests := []struct {
		name       string
		toValidate []interface{}
		valid      bool
	}{
		{
			"MULTI: Valid test w/validate method passing pointer",
			[]interface{}{&testStruct{true}, &testStruct{true}, &testStruct{true}},
			true,
		},
		{
			"MULTI: Valid test w/base types",
			[]interface{}{"testify", 1223, true},
			true,
		},
		{
			"MULTI: Valid test w/validate method",
			[]interface{}{testStruct{true}, testStruct{true}, testStruct{true}, testStruct{true}},
			true,
		},
		{
			"MULTI: Invalid test w/validate method, passing pointer",
			[]interface{}{&testStruct{false}, &testStruct{false}, &testStruct{false}, &testStruct{false}},
			false,
		},
		{
			"MULTI: Invalid test w/partial valid validate method, passing pointer",
			[]interface{}{&testStruct{true}, &testStruct{true}, &testStruct{true}, &testStruct{false}},
			false,
		},
		{
			"MULTI: Invalid test w/validate method",
			[]interface{}{testStruct{false}, testStruct{false}, testStruct{false}, testStruct{false}},
			false,
		},
		{
			"MULTI: Invalid test w/validate method partial valid",
			[]interface{}{testStruct{true}, testStruct{true}, testStruct{true}, testStruct{false}},
			false,
		},
		{
			"MULTI: Valid test w/out validate method, passing pointer",
			[]interface{}{&noInterfaceimpl{}, &noInterfaceimpl{}, &noInterfaceimpl{}, &noInterfaceimpl{}},
			true,
		},
		{
			"MULTI: Valid test w/out validate method, passing pointer partial valid",
			[]interface{}{&testStruct{true}, &testStruct{true}, &testStruct{true}, &noInterfaceimpl{}},
			true,
		},
		{
			"MULTI: Valid test w/out validate method",
			[]interface{}{noInterfaceimpl{}, noInterfaceimpl{}, noInterfaceimpl{}, noInterfaceimpl{}},
			true,
		},
		{
			"MULTI: Valid test w/out validate method partial diff valid",
			[]interface{}{&testStruct{true}, &testStruct{true}, &testStruct{true}, noInterfaceimpl{}},
			true,
		},
		{
			"MULTI: Invalid nil obj test w/ partial valid",
			[]interface{}{&testStruct{true}, &testStruct{true}, &testStruct{true}, nilobj},
			false,
		},
		{
			"MULTI: Invalid nil obj test",
			[]interface{}{nilobj, nilobj, nilobj, nilobj},
			false,
		},
	}

	for _, test := range tests {
		if IsValid(test.toValidate...) != test.valid {
			t.Errorf("Test [%s] Failed", test.name)
		}
	}
}

func TestAssert_NIL(t *testing.T) {
	result := "empty list passed to asert"
	if err := Assert(); err != nil {
		if err.Error() != result {
			t.Errorf("Test Failed '%s != %s'", err.Error(), result)
		}
	} else {
		t.Errorf("TestAssert_Empty Failed empty error")
	}
}

func TestAssert_Empty(t *testing.T) {
	result := "empty list passed to asert"
	if err := Assert(); err != nil {
		if err.Error() != result {
			t.Errorf("Test Failed '%s != %s'", err.Error(), result)
		}
	} else {
		t.Errorf("TestAssert_Empty Failed empty error")
	}
}

func TestAssert_StringSlice(t *testing.T) {

	result := "parameter at index [0] is invalid | [empty string at index [1] in []string]"
	if err := Assert([]string{"testy index 0", "", "testy index 1"}); err != nil {
		if err.Error() != result {
			t.Errorf("TestAssert_String Failed `%s` != `%s`", err.Error(), result)
		}
	} else {
		t.Errorf("TestAssert_String Failed empty error")
	}
}

func TestAssert_ByteSlice(t *testing.T) {

	result := "parameter at index [0] is invalid | [empty []byte]"
	if err := Assert([]byte{}); err != nil {
		if err.Error() != result {
			t.Errorf("TestAssert_ByteSlice Failed `%s` != `%s`", err.Error(), result)
		}
	} else {
		t.Errorf("TestAssert_ByteSlice Failed empty error")
	}
}

func TestAssert_String(t *testing.T) {

	result := "parameter at index [0] is invalid | [empty string]"
	if err := Assert(""); err != nil {
		if err.Error() != result {
			t.Errorf("TestAssert_String Failed `%s` != `%s`", err.Error(), result)
		}
	} else {
		t.Errorf("TestAssert_String Failed empty error")
	}
}

func TestAssert_NonBase(t *testing.T) {
	var nilobj *noInterfaceimpl

	tests := []struct {
		name       string
		toValidate []interface{}
		result     string
	}{
		{
			"Index 1 - Failed Validator",
			[]interface{}{&testStruct{true}, &testStruct{false}, &testStruct{true}},
			"parameter at index [1] is invalid | [validator failed]",
		},
		{
			"Index 1 - NIL",
			[]interface{}{&testStruct{true}, nil, &testStruct{true}},
			"parameter at index [1] is invalid | [nil]",
		},
		{
			"Index 1 - NIL Object",
			[]interface{}{&testStruct{true}, nilobj, &testStruct{true}},
			"parameter at index [1] is invalid",
		},
	}

	for _, test := range tests {
		if err := Assert(test.toValidate...); err != nil {
			if err.Error() != test.result {
				t.Errorf("Test [%s] Failed `%s` != `%s`", test.name, err.Error(), test.result)
			}
		} else {
			t.Errorf("Test [%s] Failed empty error", test.name)
		}
	}
}

func Benchmark_IsValid_ByteSlice(b *testing.B) {
	var testobj = []byte{12, 12, 12}

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid byte slice, expected valid")
		}
	}
}

func Benchmark_IsValid_StringSlice(b *testing.B) {
	var testobj = []string{"hello", "kitty", "slicey"}

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid string slice, expected valid")
		}
	}
}

func Benchmark_IsValid_String(b *testing.B) {
	var testobj = "hello kitty"

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid string, expected valid")
		}
	}
}

func Benchmark_IsValid_ValidWValidator(b *testing.B) {
	var testobj = testStruct{true}

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func Benchmark_IsValid_ValidWOValidator(b *testing.B) {
	var testobj = noInterfaceimpl{}

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func Benchmark_IsValid_ValidPtrWValidator(b *testing.B) {
	var testobj = &testStruct{true}

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func Benchmark_IsValid_ValidPtrWOValidator(b *testing.B) {
	var testobj = &noInterfaceimpl{}

	for n := 0; n < b.N; n++ {
		if !IsValid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}
}

func Benchmark_IsValid_NILNonInterface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if IsValid(nil) {
			b.Error("Valid struct, expected invalid")
		}
	}
}

func Benchmark_IsValid_NIL(b *testing.B) {
	var testobj *noInterfaceimpl

	for n := 0; n < b.N; n++ {
		if IsValid(testobj) {
			b.Error("Valid struct, expected invalid")
		}
	}

}
