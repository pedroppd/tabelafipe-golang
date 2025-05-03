package models

type FipeTable struct {
	CodigoTipoVeiculo uint64 `json:"codigoTipoVeiculo,omitempty"`
	CodigoModelo      uint64 `json:"codigoModelo,omitempty"`
	CodigoMarca       uint64 `json:"codigoMarca,omitempty"`
	AnoModelo         uint64 `json:"anoModelo,omitempty"`
}
