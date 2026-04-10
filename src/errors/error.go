package errors

type AppError struct {
	Message    string  `json:"message"`
	DebugError *string `json:"debug,omitempty"`
	sys        error
}
