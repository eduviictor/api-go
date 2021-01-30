package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	return c.SendString("All Users")
}

func GetUser(c *fiber.Ctx) error {
	return c.SendString("One User")
}

func CreateUser(c *fiber.Ctx) error {
	return c.SendString("Create User")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update User")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Delete User")
}
