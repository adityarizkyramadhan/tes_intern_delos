package response

type Response struct {
	Success         bool        `json:"success"`
	Message         string      `json:"message"`
	Count           int         `json:"count"`
	UniqueUserAgent int         `json:"unique_user_agent"`
	Body            interface{} `json:"body"`
}

func ResponseWhenFail(message string, body interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Body:    body,
	}
}

func ResponseWhenSuccess(count int, uniqueUserAgent int, message string, body interface{}) Response {
	return Response{
		Count:           count,
		UniqueUserAgent: uniqueUserAgent,
		Success:         true,
		Message:         message,
		Body:            body,
	}
}
