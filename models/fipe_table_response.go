package models

type Body struct {
	Valor            float64 `json:"valor,omitempty"`
	Marca            string  `json:"marca,omitempty"`
	Modelo           string  `json:"modelo,omitempty"`
	AnoModelo        uint64  `json:"anoModelo,omitempty"`
	Combustivel      string  `json:"combustivel,omitempty"`
	CodigoFipe       string  `json:"codigoFipe,omitempty"`
	MesReferencia    string  `json:"mesReferencia,omitempty"`
	Autenticacao     string  `json:"autenticacao,omitempty"`
	TipoVeiculo      string  `json:"tipoVeiculo,omitempty"`
	SiglaCombustivel string  `json:"siglaCombustivel,omitempty"`
	DataConsulta     string  `json:"dataConsulta,omitempty"`
}
