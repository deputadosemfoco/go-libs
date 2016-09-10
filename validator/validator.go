package validator

type (
	// Validator interface for all validation operations
	Validator interface {
		Validate() (bool, []string)

		NotEmpty(str string, msg string) bool
		IsEmpty(str string) bool
		IsFltGreaterThanZero(value float64, msg string) bool
		IsIntGreaterThanZero(value int, msg string) bool
	}

	// SimpleValidator is the implementation for Validator
	SimpleValidator struct {
		messages []string
	}
)

// Build and returns new Validator instance
func Build() *SimpleValidator {
	return &SimpleValidator{}
}

// NotEmpty checks if string is empty and invalid message to messages array
func (v *SimpleValidator) NotEmpty(str string, msg string) bool {
	if str == "" {
		v.messages = append(v.messages, msg)
		return false
	}

	return true
}

// IsEmpty simple checks if string is empty
func (v *SimpleValidator) IsEmpty(str string) bool {
	if str == "" {
		return true
	}

	return false
}

// IsFltGreaterThanZero checks if float value is greater than 0
func (v *SimpleValidator) IsFltGreaterThanZero(value float64, msg string) bool {
	if value == 0 {
		v.messages = append(v.messages, msg)
		return false
	}

	return true
}

// IsIntGreaterThanZero checks if int value is greater than 0
func (v *SimpleValidator) IsIntGreaterThanZero(value int, msg string) bool {
	if value == 0 {
		v.messages = append(v.messages, msg)
		return false
	}

	return true
}

// Validate checks if messages array contains any value and returns it with a bool value indicating if is valid or not
func (v *SimpleValidator) Validate() (bool, []string) {
	return len(v.messages) == 0, v.messages
}
