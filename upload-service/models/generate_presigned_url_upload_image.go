package models

type GeneratePresignedURLUploadImagePayload struct {
	FileSize   string `json:"fileSize"`
	Filename   string `json:"filename"`
	HttpMethod string `json:"httpMethod"`
	MimeType   string `json:"mimeType"`
	Resource   string `json:"resource"`
}
