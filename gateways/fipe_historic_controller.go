package gateways

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tabela-fipe-golang/externalapi"
	"tabela-fipe-golang/models"
)

func GetFipeHistoric(w http.ResponseWriter, r *http.Request) {
	historicFipeTableBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error to try read the body", http.StatusInternalServerError)
		return
	}

	var fipeTable models.FipeTable

	if err = json.Unmarshal(historicFipeTableBody, &fipeTable); err != nil {
		http.Error(w, "Error to try parse the body", http.StatusInternalServerError)
		return
	}

	referenceTables, err := externalapi.GetReferenceTables()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, referenceTable := range referenceTables {
		fmt.Println(referenceTable)
	}
}
