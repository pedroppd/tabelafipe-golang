package models

type FipeTableHistoric struct {
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
