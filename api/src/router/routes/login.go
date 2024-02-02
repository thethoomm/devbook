package routes

import (
	"api/src/controllers"
)

var loginRoute = Route{
	Uri:           "/login",
	Method:        "POST",
	Function:      controllers.Login,
	Authenticated: false,
}
