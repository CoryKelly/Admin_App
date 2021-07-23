package controllers

import (
	"github.com/CoryKelly/Admin_App/db"
	"github.com/CoryKelly/Admin_App/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	db.Database.Preload("Role").Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.SetPassword("1234")

	db.Database.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	//Get id from url
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	db.Database.Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	db.Database.Model(&user).Updates(&user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	db.Database.Delete(&user)

	return nil
}
