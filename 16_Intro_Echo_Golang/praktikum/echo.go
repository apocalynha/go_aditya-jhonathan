package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user by id",
				"user":     user,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"error": "User not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	for i, user := range users {
		if user.Id == id {
			// Delete user from the slice.
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete user by id",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"error": "User not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid user ID",
		})
	}

	var updatedUser User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	for i, user := range users {
		if user.Id == id {
			// Update the user.
			updatedUser.Id = id
			users[i] = updatedUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update user by id",
				"user":     updatedUser,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"error": "User not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)    // Added dynamic route for getting user by ID
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController) // Added dynamic route for updating user by ID
	e.DELETE("/users/:id", DeleteUserController) // Added dynamic route for deleting user by ID

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
