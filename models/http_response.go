package models

import (
	"encoding/json"
	"fmt"
)

type HttpResponse struct {
	StatusCode   int    `json:"statusCode,omitempty"`
	BodyResponse []byte `json:"bodyResponse,omitempty"`
	BodyRequest  []byte `json:"bodyRequest,omitempty"`
	Err          error  `json:"error,omitempty"`
}

func (httpResponse *HttpResponse) IsSuccess() bool {
	return httpResponse.StatusCode == 200
}

func (httpResponse *HttpResponse) GetBodyRequest() *FipeTable {
	var fipeTableRequest FipeTable
	if err := json.Unmarshal(httpResponse.BodyRequest, &fipeTableRequest); err != nil {
		fmt.Printf("Erro ao fazer unmarshal de BodyRequest: %v\n", err)
		return nil
	}
	return &fipeTableRequest
}

func (httpResponse *HttpResponse) GetBodyResponse() *FipeTableResponse {
	var fipeTableResponse FipeTableResponse
	if err := json.Unmarshal(httpResponse.BodyResponse, &fipeTableResponse); err != nil {
		fmt.Printf("Erro ao fazer unmarshal de BodyRequest: %v\n", err)
		return nil
	}
	return &fipeTableResponse
}
