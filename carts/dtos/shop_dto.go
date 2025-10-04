package dtos

type Shop struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetShopsResponse struct {
	Shops       []Shop   `json:"shops"`
	NotFoundIDs []string `json:"notFoundIds,omitempty"`
}
