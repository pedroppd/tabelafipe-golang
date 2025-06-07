package externalapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"tabela-fipe-golang/models"
)

func GetFipeTable(fipeTableRequest models.FipeTable, wg *sync.WaitGroup, ch chan<- models.HttpResponse) {
	url := "http://veiculos.fipe.org.br/api/veiculos/ConsultarValorComTodosParametros"
	defer wg.Done()
	jsonData, err := json.Marshal(fipeTableRequest)
	if err != nil {
		ch <- models.HttpResponse{BodyRequest: jsonData, Err: fmt.Errorf("JSON marshal error: %w", err)}
		return
	}
	fmt.Println(string(jsonData))

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		ch <- models.HttpResponse{StatusCode: response.StatusCode, BodyRequest: jsonData, Err: fmt.Errorf("Error making POST request to fipe table: %w", err)}
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ch <- models.HttpResponse{StatusCode: response.StatusCode, BodyRequest: jsonData, Err: fmt.Errorf("Error reading response from POST fipe table: %w", err)}
		return
	}

	ch <- models.HttpResponse{StatusCode: response.StatusCode, BodyResponse: body, BodyRequest: jsonData}
}
