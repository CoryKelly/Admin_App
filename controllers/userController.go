package controllers

import (
	"github.com/CoryKelly/Admin_App/db"
	"github.com/CoryKelly/Admin_App/middleware"
	"github.com/CoryKelly/Admin_App/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllUsers(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(db.Database, &models.User{}, page))
}

func CreateUser(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.SetPassword("1234")

	db.Database.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	//Get id from url
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	db.Database.Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

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

	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	db.Database.Delete(&user)

	return nil
}
