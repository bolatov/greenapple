package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AlgoIndex",
		"GET",
		"/algos",
		AlgoIndex,
	},
	Route{
		"AlgoShow",
		"GET",
		"/algos/{algoId}",
		AlgoShow,
	},
	Route{
		"AlgoCreate",
		"POST",
		"/algos",
		AlgoCreate,
	},
	Route{
		"AlgoRandom",
		"GET",
		"/random",
		AlgoRandom,
	},
}
