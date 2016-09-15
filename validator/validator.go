package validator

import "reflect"

type (
	Validator interface {
		Validate() (bool, []string)

		NotEmpty(str, msg string) bool
		IsGt(left, right interface{}, msg string) bool
	}

	SimpleValidator struct {
		messages []string
	}
)

func Build() *SimpleValidator {
	return &SimpleValidator{}
}

func (v *SimpleValidator) NotEmpty(str, msg string) bool {
	if str == "" {
		v.messages = append(v.messages, msg)
		return false
	}

	return true
}

func (v *SimpleValidator) IsGt(left, right interface{}, msg string) bool {
	_left := reflect.ValueOf(left)
	_right := reflect.ValueOf(right)
	result := true

	switch _left.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = _left.Int() > _right.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		result = _left.Uint() > _right.Uint()

	case reflect.Float32, reflect.Float64:
		result = _left.Float() > _right.Float()
	}

	if !result {
		v.messages = append(v.messages, msg)
	}

	return result
}

func (v *SimpleValidator) Validate() (bool, []string) {
	return len(v.messages) == 0, v.messages
}
