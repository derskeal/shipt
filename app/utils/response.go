package utils

type ResponseStatus string

const (
	SUCCESS ResponseStatus = "success"
	ERROR   ResponseStatus = "error"
)

type Response struct {
	Status  ResponseStatus `json:"status"`
	Message interface{}    `json:"message"`
	Data    interface{}    `json:"data"`
}

func SuccessResponse(body interface{}) Response {
	res := Response{
		Status:  SUCCESS,
		Message: nil,
		Data:    body,
	}
	return res
}
