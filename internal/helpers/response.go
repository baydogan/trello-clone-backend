package helpers

type Response struct {
	Data    any    `json: "data"`
	Code    int    `json: "code"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

func BuildSuccessResponse(code int, message string, data any) Response {
	return Response{
		Code:    code,
		Error:   false,
		Message: message,
		Data:    data,
	}
}

func BuildErrorResponse(code int, errorMessage string) Response {
	return Response{
		Code:    code,
		Error:   true,
		Message: errorMessage,
	}
}
