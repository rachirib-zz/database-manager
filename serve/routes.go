package serve

import (
	"net/http"
)

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
		"DatabaseIndex",
		"GET",
		"/databases",
		DatabaseIndex,
	},
	Route{
		"DatabaseShow",
		"GET",
		"/databases/{databaseId}",
		DatabaseShow,
	},
}
