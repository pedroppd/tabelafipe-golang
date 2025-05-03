package gateways

import (
	"fmt"
	"net/http"
)

func GetFipeHistoric(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Cheguei no controller GetFipeHistoric")
}
