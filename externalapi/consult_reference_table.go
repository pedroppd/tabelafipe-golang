package externalapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"tabela-fipe-golang/models"
)

func GetReferenceTables() ([]models.ReferenceTable, error) {
	url := "http://veiculos.fipe.org.br/api/veiculos/ConsultarTabelaDeReferencia"

	response, err := http.Post(url, "application/json", nil)

	if err != nil {
		return nil, errors.New("Error making GET request to reference table: " + err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error reading response from GET reference table: " + err.Error())
	}

	var referenceTable []models.ReferenceTable

	if err := json.Unmarshal(body, &referenceTable); err != nil {
		return nil, errors.New("Error unmarshaling JSON: " + err.Error())
	}
	return referenceTable, nil
}
