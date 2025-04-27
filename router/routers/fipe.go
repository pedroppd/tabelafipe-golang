package routers

import "net/http"

var fipeRoutes = []Router{
	{
		URI:    "/fipe-historic",
		Method: http.MethodPost,
		Func:   func(http.ResponseWriter, *http.Request) {},
	},
}
