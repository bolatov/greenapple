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
		"/algo",
		AlgoIndex,
	},
	Route{
		"AlgoShow",
		"GET",
		"/algo/{algoId}",
		AlgoShow,
	},
	Route{
		"AlgoCreate",
		"POST",
		"/algo",
		AlgoCreate,
	},
	Route{
		"AlgoRandom",
		"GET",
		"/random",
		AlgoRandom,
	},
	Route{
		"AlgoList",
		"GET",
		"/list",
		AlgoList,
	}
}
