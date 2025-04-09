package routes

import (
	"net/http"

	"github.com/book-wise/controllers"
)

var LoginRoute = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	Function:    controllers.Login,
	RequireAuth: false,
}
