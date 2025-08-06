package Models

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

func SuccessResponse(data interface{}, message string) APIResponse {
	return APIResponse{
		Code:    200,
		Message: message,
		Data:    data,
		Success: true,
	}
}

func ErrorResponse(code int, message string) APIResponse {
	return APIResponse{
		Code:    code,
		Message: message,
		Success: false,
	}
}

func ValidationErrorResponse(message string) APIResponse {
	return APIResponse{
		Code:    400,
		Message: message,
		Success: false,
	}
}

func NotFoundResponse(message string) APIResponse {
	return APIResponse{
		Code:    404,
		Message: message,
		Success: false,
	}
}

func UnauthorizedResponse(message string) APIResponse {
	return APIResponse{
		Code:    401,
		Message: message,
		Success: false,
	}
}

func InternalErrorResponse(message string) APIResponse {
	return APIResponse{
		Code:    500,
		Message: message,
		Success: false,
	}
}