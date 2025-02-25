package common

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Status:  200,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(status int, message string, err error) APIResponse {
	var errMsg interface{}
	if err != nil {
		errMsg = err.Error()
	}

	return APIResponse{
		Status:  status,
		Message: message,
		Error:   errMsg,
	}
}
