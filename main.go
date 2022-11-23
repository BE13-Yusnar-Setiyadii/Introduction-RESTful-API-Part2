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
	// your solution here
	userid, _ := strconv.Atoi(c.Param("id"))
	var data int
	for x, z := range users {
		if z.Id == userid {
			data = x
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"users":   users[data],
	})
}

// // delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	userid, _ := strconv.Atoi(c.Param("id"))
	var data int
	for x, z := range users {
		if z.Id == userid {
			data = x
		}
	}

	users = append(users[:data], users[data+1:]...)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
		"users":   users,
	})
}

// // update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	userid, _ := strconv.Atoi(c.Param("id"))
	var data int
	for x, z := range users {
		if z.Id == userid {
			data = x
		}
	}
	user := User{}
	c.Bind(&user)

	x := users

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id",
		"users":   x[data],
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"users":    user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
