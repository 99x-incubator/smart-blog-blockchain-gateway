package routes

import (
	"net/http"

	"github.com/Blogchain-Gateway/api/handlers"
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
		"CreateAccount",
		"GET",
		"/createaccount",
		handlers.CreateAccount,
	},
	Route{
		"Transactions",
		"POST",
		"/transaction/type/{TType}",
		handlers.Transaction,
	},
}
