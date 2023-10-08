package route

import (
	"eksplorasi/controllers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Not Authenticated
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginController)

	// Route Blog
	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)

	secret := os.Getenv("JWT_SECRET")
	// Authenticated
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(secret)))

	// Users
	eAuth.GET("/users", controllers.GetUsersController)
	eAuth.GET("/users/:id", controllers.GetUserController)
	eAuth.DELETE("/users/:id", controllers.DeleteUserController)
	eAuth.PUT("/users/:id", controllers.UpdateUserController)

	// Books
	eAuth.GET("/books", controllers.GetBooksController)
	eAuth.GET("/books/:id", controllers.GetBookController)
	eAuth.POST("/books", controllers.CreateBookController)
	eAuth.DELETE("/books/:id", controllers.DeleteBookController)
	eAuth.PUT("/books/:id", controllers.UpdateBookController)

	return e
}
