package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponse(message string, code int, data interface{}) Response {

	jsonResponse := Response{
		Status:  code,
		Message: message,
		Data:    data,
	}

	return jsonResponse
}
