package models

import (
	"strings"
	"tabela-fipe-golang/shared"
)

type (
	FipeTableRequestResponse struct {
		ResponseBody *FipeTableResponse `json:"responseBody,omitempty"`
		RequestBody  *FipeTable         `json:"requestBody,omitempty"`
		StatusCode   int                `json:"statusCode,omitempty"`
	}
	FipeTableRequest struct {
		CodigoTipoVeiculo uint64 `json:"codigoTipoVeiculo,omitempty"`
		CodigoModelo      uint64 `json:"codigoModelo,omitempty"`
		CodigoMarca       uint64 `json:"codigoMarca,omitempty"`
		AnoModelo         uint64 `json:"anoModelo,omitempty"`
	}
	FipeTableResponse struct {
		Valor         string `json:"valor,omitempty"`
		Marca         string `json:"marca,omitempty"`
		Modelo        string `json:"modelo,omitempty"`
		AnoModelo     uint64 `json:"anoModelo,omitempty"`
		Combustivel   string `json:"combustivel,omitempty"`
		MesReferencia string `json:"mesReferencia,omitempty"`
		DataConsulta  string `json:"dataConsulta,omitempty"`
	}
	FipeTable struct {
		CodigoTipoVeiculo      uint64 `json:"codigoTipoVeiculo,omitempty"`
		CodigoTabelaReferencia uint64 `json:"codigoTabelaReferencia,omitempty"`
		CodigoModelo           uint64 `json:"codigoModelo,omitempty"`
		CodigoMarca            uint64 `json:"codigoMarca,omitempty"`
		CodigoTipoCombustivel  uint64 `json:"codigoTipoCombustivel,omitempty"`
		AnoModelo              uint64 `json:"anoModelo,omitempty"`
		TipoVeiculo            string `json:"tipoVeiculo,omitempty"`
		ModeloCodigoExterno    string `json:"modeloCodigoExterno,omitempty"`
		TipoConsulta           string `json:"tipoConsulta,omitempty"`
	}
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
