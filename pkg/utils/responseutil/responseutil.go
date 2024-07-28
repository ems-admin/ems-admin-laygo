package responseutil

type CommonResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(data interface{}) *CommonResponse {
	return &CommonResponse{
		Success: true,
		Data:    data,
	}
}

func FailResponse(msg string) *CommonResponse {
	return &CommonResponse{
		Success: false,
		Message: msg,
	}
}
