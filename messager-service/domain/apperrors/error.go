package apperrors

type Error struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	LogMessage string `json:"logMessage"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewInternalServerError(message string) *Error {
	return &Error{
		Code:       500,
		Message:    "Internal Server Error",
		LogMessage: message,
	}
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Code:    400,
		Message: "Bad Request: " + message,
	}
}

func NewNotFoundError(message string) *Error {
	return &Error{
		Code:    404,
		Message: "Not Found: " + message,
	}
}

func NewUnauthorizedError(message string) *Error {
	return &Error{
		Code:    401,
		Message: "Unauthorized: " + message,
	}
}

func NewForbiddenError(message string) *Error {
	return &Error{
		Code:    403,
		Message: "Forbidden: " + message,
	}
}

func NewBadGatewayError(message string) *Error {
	return &Error{
		Code:    502,
		Message: "Bad Gateway: " + message,
	}
}

func NewServiceUnavailableError(message string) *Error {
	return &Error{
		Code:    503,
		Message: "Service Unavailable: " + message,
	}
}

func NewUnknownError(message string) *Error {
	return &Error{
		Code:    520,
		Message: "Unknown Error: " + message,
	}
}

func NewConflictError(message string) *Error {
	return &Error{
		Code:    409,
		Message: "Conflict Error: " + message,
	}
}
