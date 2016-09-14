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

	// ResponseBuilder is a builder for Response struct
	ResponseBuilder struct{ Response }

	// Error is the base error struct to http responses
	Error struct {
		Message            string `json:"message"`
		ValidationMessages string `json:"validationMessages"`
	}
)

func ResBuilder() *ResponseBuilder {
	return &ResponseBuilder{}
}

func (builder *ResponseBuilder) AsSuccess(data interface{}) *ResponseBuilder {
	builder.Success = true
	builder.Data = data

	return builder
}

func (builder *ResponseBuilder) AsErr(msg string) *ResponseBuilder {
	builder.Success = false
	builder.Message = msg

	return builder
}

func (builder *ResponseBuilder) WithValMsgs(msgs []string) *ResponseBuilder {
	builder.ValidationMessages = msgs

	return builder
}

func (builder *ResponseBuilder) Build() Response {
	return builder.Response
}
