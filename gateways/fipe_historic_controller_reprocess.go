package gateways

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"tabela-fipe-golang/externalapi"
	"tabela-fipe-golang/models"
)

func GetFipeHistoricReprocessed(w http.ResponseWriter, r *http.Request) {
	fipeTableRequestUnparsed, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error to try read the body", http.StatusInternalServerError)
		return
	}

	var fipeTableRequestList []models.FipeTableRequestResponse
	if err := json.Unmarshal(fipeTableRequestUnparsed, &fipeTableRequestList); err != nil {
		http.Error(w, "Error to try parse the body", http.StatusBadRequest)
		return
	}

	fmt.Println("Building request object to fipe table...")
	var wg sync.WaitGroup
	responseChannel := make(chan models.HttpResponse, len(fipeTableRequestList))

	for _, fipeTableRequest := range fipeTableRequestList {
		if fipeTableRequest.StatusCode != 200 {
			wg.Add(1)
			externalapi.GetFipeTable(*fipeTableRequest.RequestBody, &wg, responseChannel)
		}
	}
	wg.Wait()
	close(responseChannel)

	var fipeTableRequestResponse []models.FipeTableRequestResponse
	for res := range responseChannel {
		if res.IsSuccess() {
			fmt.Println("Building request object with status: 200")
			result := models.FipeTableRequestResponse{ResponseBody: res.GetBodyResponse(), RequestBody: res.GetBodyRequest(), StatusCode: res.StatusCode}
			fipeTableRequestResponse = append(fipeTableRequestResponse, result)
		} else {
			fmt.Println("Building request object with status: ", res.StatusCode)
			result := models.FipeTableRequestResponse{ResponseBody: nil, RequestBody: res.GetBodyRequest(), StatusCode: res.StatusCode}
			fipeTableRequestResponse = append(fipeTableRequestResponse, result)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if fipeTableRequestResponse == nil {
		fipeTableRequestResponse = make([]models.FipeTableRequestResponse, 0)
	}
	if err := json.NewEncoder(w).Encode(fipeTableRequestResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
