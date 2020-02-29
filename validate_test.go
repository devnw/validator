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

func getValids() (funcs []func(objs ...interface{}) bool) {
	return []func(objs ...interface{}) bool{
		IsValid,
		Valid,
	}
}

func TestValidNIL(t *testing.T) {

	for _, f := range getValids() {
		if f(nil) != false {
			t.Errorf("Test Failed")
		}
	}
}

func TestValidEmpty(t *testing.T) {
	for _, f := range getValids() {
		if f() != false {
			t.Errorf("Test Failed")
		}
	}
}

func TestValidString(t *testing.T) {
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

	for _, f := range getValids() {
		for _, test := range tests {
			if f(test.toValidate) != test.valid {
				t.Errorf("Test [%s] Failed", test.name)
			}
		}
	}
}

func TestValidStringSlice(t *testing.T) {
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

	for _, f := range getValids() {
		for _, test := range tests {
			if f(test.toValidate) != test.valid {
				t.Errorf("Test [%s] Failed", test.name)
			}
		}
	}
}

func TestValidByteSlice(t *testing.T) {
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

	for _, f := range getValids() {
		for _, test := range tests {
			if f(test.toValidate) != test.valid {
				t.Errorf("Test [%s] Failed", test.name)
			}
		}
	}
}

func TestValid(t *testing.T) {
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

	for _, f := range getValids() {
		for _, test := range tests {
			if f(test.toValidate) != test.valid {
				t.Errorf("Test [%s] Failed", test.name)
			}
		}
	}
}

func TestValid_NotCoveredBaseTypes(t *testing.T) {
	tests := []struct {
		name       string
		toValidate interface{}
		valid      bool
	}{
		{
			"int",
			int(12),
			true,
		},
		{
			"int8",
			int8(12),
			true,
		},
		{
			"int16",
			int16(12),
			true,
		},
		{
			"int32",
			int32(12),
			true,
		},
		{
			"int64",
			int64(12),
			true,
		},
		{
			"uint",
			uint(12),
			true,
		},
		{
			"uint8",
			uint8(12),
			true,
		},
		{
			"uint16",
			uint16(12),
			true,
		},
		{
			"uint32",
			uint32(12),
			true,
		},
		{
			"uint64",
			uint64(12),
			true,
		},
		{
			"bool",
			bool(true),
			true,
		},
		{
			"float32",
			float32(12.5),
			true,
		},
		{
			"float64",
			float64(12.5),
			true,
		},
	}

	for _, f := range getValids() {
		for _, test := range tests {
			if f(test.toValidate) != test.valid {
				t.Errorf("Test [%s] Failed", test.name)
			}
		}
	}
}

func TestAssert_NotCoveredBaseTypes(t *testing.T) {
	tests := []struct {
		name       string
		toValidate interface{}
		valid      bool
	}{
		{
			"int",
			int(12),
			true,
		},
		{
			"int8",
			int8(12),
			true,
		},
		{
			"int16",
			int16(12),
			true,
		},
		{
			"int32",
			int32(12),
			true,
		},
		{
			"int64",
			int64(12),
			true,
		},
		{
			"uint",
			uint(12),
			true,
		},
		{
			"uint8",
			uint8(12),
			true,
		},
		{
			"uint16",
			uint16(12),
			true,
		},
		{
			"uint32",
			uint32(12),
			true,
		},
		{
			"uint64",
			uint64(12),
			true,
		},
		{
			"bool",
			bool(true),
			true,
		},
		{
			"float32",
			float32(12.5),
			true,
		},
		{
			"float64",
			float64(12.5),
			true,
		},
	}

	for _, test := range tests {
		if Assert(test.toValidate) != nil {
			t.Errorf("Test [%s] Failed", test.name)
		}
	}
}

func TestValid_Multiples(t *testing.T) {
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

	for _, f := range getValids() {
		for _, test := range tests {
			if f(test.toValidate...) != test.valid {
				t.Errorf("Test [%s] Failed", test.name)
			}
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

func Benchmark_Valid_ByteSlice(b *testing.B) {
	var testobj = []byte{12, 12, 12}

	for n := 0; n < b.N; n++ {
		if !Valid(testobj) {
			b.Error("Invalid byte slice, expected valid")
		}
	}
}

func Benchmark_Valid_StringSlice(b *testing.B) {
	var testobj = []string{"hello", "kitty", "slicey"}

	for n := 0; n < b.N; n++ {
		if !Valid(testobj) {
			b.Error("Invalid string slice, expected valid")
		}
	}
}

func Benchmark_Valid_String(b *testing.B) {
	var testobj = "hello kitty"

	for n := 0; n < b.N; n++ {
		if !Valid(testobj) {
			b.Error("Invalid string, expected valid")
		}
	}
}

func Benchmark_Valid_ValidWValidator(b *testing.B) {
	var testobj = testStruct{true}

	for n := 0; n < b.N; n++ {
		if !Valid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func Benchmark_Valid_ValidWOValidator(b *testing.B) {
	var testobj = noInterfaceimpl{}

	for n := 0; n < b.N; n++ {
		if !Valid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func Benchmark_Valid_ValidPtrWValidator(b *testing.B) {
	var testobj = &testStruct{true}

	for n := 0; n < b.N; n++ {
		if !Valid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func Benchmark_Valid_ValidPtrWOValidator(b *testing.B) {
	var testobj = &noInterfaceimpl{}

	for n := 0; n < b.N; n++ {
		if !Valid(testobj) {
			b.Error("Invalid struct, expected valid")
		}
	}
}

func Benchmark_Valid_NILNonInterface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if Valid(nil) {
			b.Error("Valid struct, expected invalid")
		}
	}
}

func Benchmark_Valid_NIL(b *testing.B) {
	var testobj *noInterfaceimpl

	for n := 0; n < b.N; n++ {
		if Valid(testobj) {
			b.Error("Valid struct, expected invalid")
		}
	}

}

func Benchmark_Assert_ByteSlice(b *testing.B) {
	var testobj = []byte{12, 12, 12}

	for n := 0; n < b.N; n++ {
		if Assert(testobj) != nil {
			b.Error("Invalid byte slice, expected valid")
		}
	}
}

func Benchmark_Assert_StringSlice(b *testing.B) {
	var testobj = []string{"hello", "kitty", "slicey"}

	for n := 0; n < b.N; n++ {
		if Assert(testobj) != nil {
			b.Error("Invalid string slice, expected valid")
		}
	}
}

func Benchmark_Assert_String(b *testing.B) {
	var testobj = "hello kitty"

	for n := 0; n < b.N; n++ {
		if Assert(testobj) != nil {
			b.Error("Invalid string, expected valid")
		}
	}
}

func Benchmark_Assert_AssertWValidator(b *testing.B) {
	var testobj = testStruct{true}

	for n := 0; n < b.N; n++ {
		if Assert(testobj) != nil {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func Benchmark_Assert_AssertWOValidator(b *testing.B) {
	var testobj = noInterfaceimpl{}

	for n := 0; n < b.N; n++ {
		if Assert(testobj) != nil {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func Benchmark_Assert_AssertPtrWAssertator(b *testing.B) {
	var testobj = &testStruct{true}

	for n := 0; n < b.N; n++ {
		if Assert(testobj) != nil {
			b.Error("Invalid struct, expected valid")
		}
	}

}

func Benchmark_Assert_AssertPtrWOAssertator(b *testing.B) {
	var testobj = &noInterfaceimpl{}

	for n := 0; n < b.N; n++ {
		if Assert(testobj) != nil {
			b.Error("Invalid struct, expected valid")
		}
	}
}

func Benchmark_Assert_NILNonInterface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if Assert(nil) == nil {
			b.Error("Assert struct, expected invalid")
		}
	}
}

func Benchmark_Assert_NIL(b *testing.B) {
	var testobj *noInterfaceimpl

	for n := 0; n < b.N; n++ {
		if Assert(testobj) == nil {
			b.Error("Assert struct, expected invalid")
		}
	}

}
