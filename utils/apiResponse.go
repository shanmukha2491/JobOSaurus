package utils

type ApiResponse struct {
	StatusCode int
	Message    string
	Payload    interface{}
	Success    bool
}

func NewApiResponse(StatusCode int, Message string, Payload interface{}) ApiResponse {
	return ApiResponse{StatusCode: StatusCode, Message: Message, Payload: Payload, Success: StatusCode < 400}
}
