package gateways

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"tabela-fipe-golang/models"
)

func GetFipeHistoric(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		errors.New("Error to try read the body")
		return
	}

	var fipeTable models.FipeTable

	if err = json.Unmarshal(requestBody, &fipeTable); err != nil {
		errors.New("Error to try parse the body")
		return
	}
}
