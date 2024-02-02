package routes

import (
	"api/src/controllers"
)

var userRoutes = []Route{
	{
		Uri:           "/users",
		Method:        "POST",
		Function:      controllers.CreateUser,
		Authenticated: false,
	},
	{
		Uri:           "/users",
		Method:        "GET",
		Function:      controllers.FindAllUsers,
		Authenticated: true,
	},
	{
		Uri:           "/users/{id}",
		Method:        "GET",
		Function:      controllers.FindUniqueUser,
		Authenticated: true,
	},
	{
		Uri:           "/users/{id}",
		Method:        "PUT",
		Function:      controllers.UpdateUser,
		Authenticated: true,
	},
	{
		Uri:           "/users/{id}",
		Method:        "DELETE",
		Function:      controllers.DeleteUser,
		Authenticated: true,
	},
	{
		Uri:           "/users/{id}/follow",
		Method:        "POST",
		Function:      controllers.FollowUser,
		Authenticated: true,
	},
	{
		Uri:           "/users/{id}/unfollow",
		Method:        "POST",
		Function:      controllers.UnfollowUser,
		Authenticated: true,
	},
	{
		Uri:           "/users/{id}/followers",
		Method:        "GET",
		Function:      controllers.FindFollowers,
		Authenticated: true,
	},
	{
		Uri:           "/users/{id}/following",
		Method:        "GET",
		Function:      controllers.FindFollowing,
		Authenticated: true,
	},
	{
		Uri:           "/users/{id}/update-password",
		Method:        "POST",
		Function:      controllers.UpdatePassword,
		Authenticated: true,
	},
}
