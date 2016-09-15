package check

import "testing"

func TestNotEmptyShouldPanicWithEmptyValue(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	c := C()

	c.Str().NotEmpty("", "string value cannot be empty")
}

func TestNotEmptyShouldNotPanicWithValue(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code is in panic")
		}
	}()

	c := C()

	c.Str().NotEmpty("value", "string value cannot be empty")
}
