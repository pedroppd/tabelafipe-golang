package routers

import (
	"net/http"
	"tabela-fipe-golang/gateways"
)

var fipeRoutes = []Router{
	{
		URI:    "/fipe-historic",
		Method: http.MethodPost,
		Func:   gateways.GetFipeHistoric,
	},
	{
		URI:    "/fipe-historic/reprocess",
		Method: http.MethodPost,
		Func:   gateways.GetFipeHistoricReprocessed,
	},
}
