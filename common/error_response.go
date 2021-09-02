package common

type ErrorResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
