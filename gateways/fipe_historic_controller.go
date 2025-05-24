package gateways

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"tabela-fipe-golang/externalapi"
	"tabela-fipe-golang/models"
	"tabela-fipe-golang/shared"
)

func GetFipeHistoric(w http.ResponseWriter, r *http.Request) {
	historicFipeTableBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error to try read the body", http.StatusInternalServerError)
		return
	}

	var fipeTable models.FipeTable

	if err = json.Unmarshal(historicFipeTableBody, &fipeTable); err != nil {
		http.Error(w, "Error to try parse the body", http.StatusBadRequest)
		return
	}

	referenceTables, err := externalapi.GetReferenceTables()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Filtering...")
	referenceTableFilteredList := filterByYear(referenceTables, r)

	fmt.Println("Building request object...")
	var fipeTableList []models.FipeTableHistoric
	for _, referenceTableFiltered := range referenceTableFilteredList {
		fipeTableHistoric := buildFipeTableHistoric(fipeTable, referenceTableFiltered.GetCodigo())
		fipeTableList = append(fipeTableList, fipeTableHistoric)
	}

	var wg sync.WaitGroup
	responseChannel := make(chan models.FipeTableResponse, len(fipeTableList))

	for _, fipeTableHistoric := range fipeTableList {
		wg.Add(1)
		externalapi.GetFipeTable(fipeTableHistoric, &wg, responseChannel)
	}
	wg.Wait()
	close(responseChannel)
}

func buildFipeTableHistoric(fipeTable models.FipeTable, referenceTable uint64) models.FipeTableHistoric {
	return models.FipeTableHistoric{CodigoTipoVeiculo: fipeTable.CodigoTipoVeiculo,
		CodigoTabelaReferencia: referenceTable,
		CodigoModelo:           fipeTable.CodigoModelo,
		CodigoMarca:            fipeTable.CodigoMarca,
		CodigoTipoCombustivel:  1,
		AnoModelo:              fipeTable.AnoModelo,
		TipoVeiculo:            "carro",
		ModeloCodigoExterno:    "",
		TipoConsulta:           "tradicional",
	}
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
	monthsSet := shared.ToSet(months)

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
