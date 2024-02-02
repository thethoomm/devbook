package routes

import "api/src/controllers"

var postRoutes = []Route{
	{
		Uri:           "/posts",
		Method:        "POST",
		Function:      controllers.CreatePost,
		Authenticated: true,
	},
	{
		Uri:           "/posts",
		Method:        "GET",
		Function:      controllers.FindAllPosts,
		Authenticated: true,
	},
	{
		Uri:           "/posts/{id}",
		Method:        "GET",
		Function:      controllers.FindUniquePost,
		Authenticated: true,
	},
	{
		Uri:           "/posts/{id}",
		Method:        "PUT",
		Function:      controllers.UpdatePost,
		Authenticated: true,
	},
	{
		Uri:           "/posts/{id}",
		Method:        "DELETE",
		Function:      controllers.DeletePost,
		Authenticated: true,
	},
	{
		Uri:           "/users/{id}/posts",
		Method:        "GET",
		Function:      controllers.FindPostByUser,
		Authenticated: true,
	},
	{
		Uri:           "/posts/{id}/like",
		Method:        "POST",
		Function:      controllers.LikePost,
		Authenticated: true,
	},
	{
		Uri:           "/posts/{id}/dislike",
		Method:        "POST",
		Function:      controllers.DislikePost,
		Authenticated: true,
	},
}
