package externalapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"tabela-fipe-golang/models"
)

func GetFipeTable(fipeTableHistoric models.FipeTableHistoric, wg *sync.WaitGroup, ch chan<- models.FipeTableResponse) {
	url := "http://veiculos.fipe.org.br/api/veiculos/ConsultarValorComTodosParametros"

	jsonData, err := json.Marshal(fipeTableHistoric)
	if err != nil {
		ch <- models.FipeTableResponse{Err: fmt.Errorf("JSON marshal error: %w", err)}
		return
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		ch <- models.FipeTableResponse{Err: fmt.Errorf("Error making POST request to fipe table: %w", err)}
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ch <- models.FipeTableResponse{Err: fmt.Errorf("Error reading response from POST fipe table: %w", err)}
		return
	}

	ch <- models.FipeTableResponse{}
	var fipeTableResponse models.FipeTableResponse

	if err := json.Unmarshal(body, &fipeTableResponse); err != nil {
		log.Printf("Error unmarshaling JSON: %v", err)
		return
	}
}
