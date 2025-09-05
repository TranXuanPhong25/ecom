package main

type AuthzResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	Allow bool `json:"allow"`
}

type AuthzRequest struct {
	Input OPAInput `json:"input"`
}

type ExtAuthzResponse struct {
	Allow  bool   `json:"allow"`
	Reason string `json:"reason"`
}
type OPAInput struct {
	Method        string   `json:"method"`
	Path          string   `json:"path"`
	Roles         []string `json:"roles"`
	Resource      string   `json:"resource"`
	Action        string   `json:"action"`
	Authenticated bool     `json:"authenticated"`
}
