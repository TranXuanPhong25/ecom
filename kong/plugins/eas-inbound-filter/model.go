package main

type OPAResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	Allow bool `json:"allow"`
}

type OPARequest struct {
	Input OPAInput `json:"input"`
}

type OPAInput struct {
	Method        string `json:"method"`
	Path          string `json:"path"`
	Authenticated bool   `json:"authenticated"`
}
