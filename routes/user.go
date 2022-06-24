package routes

import (
	"github.com/jaykeHarrison/go-rest-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/go-rest-api/models"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func CreateResponseUser (userModel models.User) User {
	return User{
		ID: userModel.ID,
		FirstName: userModel.FirstName,
		LastName: userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	
	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []User{}

	database.Database.Db.Find(&users)

	return c.Status(200).JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Database.Db

	var user User
	db.First(&user, id)

	if user.FirstName == "" {
		return c.Status(404).JSON("No user found by that ID")
	}

	return c.Status(200).JSON(user)
}

func DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Database.Db

	var user models.User

	db.Find(&user, id)
	
	if user.FirstName == "" {
		return c.Status(404).JSON("No user found by that ID")
	}

	db.Delete(&user, id)

	return c.Status(200).JSON("User deleted")
}