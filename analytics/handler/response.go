package handler

type ResponsePostEvent struct {
	StatusCode int32                  `json:"status"`
	Message    string                 `json:"message"`
	Data       *ResponsePostEventData `json:"data"`
}

type ResponsePostEventData struct {
	Name string
}

type ResponsePostUser struct {
	StatusCode int32  `json:"status"`
	Message    string `json:"message"`
	data       *ResponsePostUserData
}

type ResponsePostUserData struct {
	Username string
	Email    string
}
