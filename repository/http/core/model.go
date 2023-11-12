package core

type GetUtilityResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    GetUtilityData `json:"data"`
}

type GetUtilityData struct {
	Value string `json:"value"`
}

type SetUtilityRequest struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Duration int    `json:"duration"`
}

type SetUtilityResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SendEmailRequest struct {
	To       string `json:"to"`
	FromName string `json:"fromName"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

type SendEmailResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UploadImageResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    UploadImageData `json:"data"`
}

type UploadImageData struct {
	Images []string `json:"images"`
}
