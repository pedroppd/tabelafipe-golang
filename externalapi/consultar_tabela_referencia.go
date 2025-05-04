package externalapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ReferenceTable struct {
	Codigo int    `json:"Codigo"`
	Mes    string `json:"Mes"`
}

func GetReferenceTables() ([]ReferenceTable, error) {
	url := "http://veiculos.fipe.org.br/api/veiculos/ConsultarTabelaDeReferencia"

	response, err := http.Get(url)

	if err != nil {
		return nil, errors.New("Error making GET request to reference table: " + err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error reading response from GET reference table: " + err.Error())
	}

	var referenceTable []ReferenceTable

	if err := json.Unmarshal(body, &referenceTable); err != nil {
		return nil, errors.New("Error unmarshaling JSON: " + err.Error())
	}
	return referenceTable, nil
}
