package gateways

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"tabela-fipe-golang/externalapi"
	"tabela-fipe-golang/models"
	"tabela-fipe-golang/shared"
)

func GetFipeHistoric(w http.ResponseWriter, r *http.Request) []models.ReferenceTable {
	historicFipeTableBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error to try read the body", http.StatusInternalServerError)
		return []models.ReferenceTable{}
	}

	var fipeTable models.FipeTable

	if err = json.Unmarshal(historicFipeTableBody, &fipeTable); err != nil {
		http.Error(w, "Error to try parse the body", http.StatusBadRequest)
		return []models.ReferenceTable{}
	}

	referenceTables, err := externalapi.GetReferenceTables()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return []models.ReferenceTable{}
	}

	return filterByYear(referenceTables, r)
}

func filterByYear(referenceTables []models.ReferenceTable, r *http.Request) []models.ReferenceTable {
	var newReferenceTables []models.ReferenceTable
	query := r.URL.Query()
	beginYear, err := shared.ParseStringToInt(query.Get("beginYear"))
	if err != nil {
		beginYear = 0
	}

	endYear, err := shared.ParseStringToInt(query.Get("endYear"))
	if err != nil {
		endYear = 9999
	}

	months := getMonths(query)
	monthsSet := shared.ToSet(months) // opcional: cria um map[string]bool pra lookup rápido

	for _, referenceTable := range referenceTables {
		year := referenceTable.GetYear()
		month := referenceTable.GetMonth()

		isInYearRange := year >= beginYear && year <= endYear
		isInMonthList := len(months) == 0 || monthsSet[month]

		if isInYearRange && isInMonthList {
			newReferenceTables = append(newReferenceTables, referenceTable)
		}
	}
	return newReferenceTables
}

func getMonths(query url.Values) []string {
	months := query.Get("months")
	if months != "" {
		return strings.Split(months, ",")
	}
	return []string{}
}
