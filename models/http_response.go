package models

type HttpResponse struct {
	StatusCode   int    `json:"statusCode,omitempty"`
	BodyResponse []byte `json:"bodyResponse,omitempty"`
	BodyRequest  []byte `json:"bodyRequest,omitempty"`
	Err          error  `json:"error,omitempty"`
}
