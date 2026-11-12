package dto

type URL struct {
	Path string `json:"path"`
}

type ProblemResponse struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}
