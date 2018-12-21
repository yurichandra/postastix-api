package handler

type errorResponse struct {
	HTTPStatusCode int `json:"-"`
	Error          error
	ErrorText      string `json:"errors"`
}
