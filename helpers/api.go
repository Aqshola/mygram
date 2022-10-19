package helpers

type ApiInfo struct {
	Status  uint   `json:"status"`
	Message string `json:"message"`
}

type ApiResponse struct {
	Info ApiInfo     `json:"info"`
	Data interface{} `json:"data"`
}

func GenerateApiResponse(status uint, message string, data interface{}) ApiResponse {
	return ApiResponse{
		Info: ApiInfo{
			Status:  status,
			Message: message,
		},
		Data: data,
	}
}
