package models

import "encoding/json"

type GeneratePresignedURLUploadImagePayload struct {
	FileSize   json.Number `json:"fileSize"`
	FileName   string      `json:"fileName"`
	HttpMethod string      `json:"httpMethod"`
	MimeType   string      `json:"mimeType"`
	Resource   string      `json:"resource"`
}
