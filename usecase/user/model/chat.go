package model

type ChatRequest struct {
	Text string `json:"text" validate:"required"`
}

type ChatResponse struct {
	Text string `json:"text"`
}
