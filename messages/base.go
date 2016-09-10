package messages

type (
	// Request is the base struct for all requests
	Request struct {
	}

	// Response is the base struct for all responses
	Response struct {
		Success            bool        `json:"success"`
		Message            string      `json:"message"`
		ValidationMessages []string    `json:"validationMessages"`
		Data               interface{} `json:"data"`
	}
)
