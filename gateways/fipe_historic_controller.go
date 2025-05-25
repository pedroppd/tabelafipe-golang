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
	fipeTableRequestUnparsed, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error to try read the body", http.StatusInternalServerError)
		return
	}

	var fipeTableRequest models.FipeTableRequest

	if err = json.Unmarshal(fipeTableRequestUnparsed, &fipeTableRequest); err != nil {
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
	var fipeTableRequestList []models.FipeTableHistoric
	for _, referenceTableFiltered := range referenceTableFilteredList {
		fipeTableHistoric := buildFipeTableHistoric(fipeTableRequest, referenceTableFiltered.GetCodigo())
		fipeTableRequestList = append(fipeTableRequestList, fipeTableHistoric)
	}

	fmt.Println("Building request object to fipe table...")
	var wg sync.WaitGroup
	responseChannel := make(chan models.HttpResponse, len(fipeTableRequestList))

	for _, fipeTableRequest := range fipeTableRequestList {
		wg.Add(1)
		externalapi.GetFipeTable(fipeTableRequest, &wg, responseChannel)
	}
	wg.Wait()
	close(responseChannel)

	for res := range responseChannel {
		if res.StatusCode == 200 {
			fmt.Printf("[OK] Status %d \n\n", res.StatusCode)
		} else {
			fmt.Printf("[ERROR] %v\n", res.Err)
		}
	}
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
