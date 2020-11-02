package common

// BaseHTTPResponse is the base response for all objects
type BaseHTTPResponse struct {
	Meta             *Meta             `json:"meta,omitempty"`
	Data             interface{}       `json:"data"`
	ValidationErrors map[string]string `json:"validation_errors,omitempty"`
	Error            *ErrorResponse    `json:"error,omitempty"`
}

// Meta meta object
type Meta struct {
	Service string `json:"service,omitempty"`
	Version string `json:"version,omitempty"`
}

// ErrorResponse for generic response error detail
type ErrorResponse struct {
	Message   string `json:"message,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
}
