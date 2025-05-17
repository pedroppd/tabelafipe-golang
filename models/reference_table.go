package models

import (
	"strings"
	"tabela-fipe-golang/shared"
)

type ReferenceTable struct {
	Codigo uint64 `json:"Codigo"`
	Mes    string `json:"Mes"`
}

func (referenceTable *ReferenceTable) GetCodigo() uint64 {
	return referenceTable.Codigo
}

func (referenceTable *ReferenceTable) GetYear() int {
	year := strings.Split(referenceTable.Mes, "/")[1]
	value, _ := shared.ParseStringToInt(year)
	return value
}

func (referenceTable *ReferenceTable) GetMonth() string {
	return strings.Split(referenceTable.Mes, "/")[0]
}
