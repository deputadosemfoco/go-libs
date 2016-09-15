package domain

type Error struct {
	Message string   `json:"message,omitempty"`
	Code    int      `json:"code,omitempty"`
	Detail  string   `json:"detail,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}
