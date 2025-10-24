package response

type DataResponse struct {
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}

type APIResponse struct {
	Data DataResponse `json:"data"`
}
