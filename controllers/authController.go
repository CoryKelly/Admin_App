package controllers

import (
	"github.com/CoryKelly/Admin_App/db"
	"github.com/CoryKelly/Admin_App/models"
	"github.com/CoryKelly/Admin_App/util"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	//Get payload User check for error
	var payload map[string]string
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	//Check if password is correct
	if payload["password"] != payload["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "password does not match",
		})
	}

	//Set User payload
	user := models.User{
		FirstName: payload["first_name"],
		LastName:  payload["last_name"],
		Email:     payload["email"],
		RoleId:    1,
	}

	user.SetPassword(payload["password"])

	//Store User in Database
	db.Database.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	//Get payload User check for error
	var payload map[string]string
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	//Find Users Email
	var user models.User
	db.Database.Where("email=?", payload["email"]).First(&user)

	//Handle if user is not found
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"Message": "Not Found",
		})
	}
	//Handle if users password matches or is incorrect
	if err := user.ComparePassword(payload["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"Message": "Invalid password",
		})
	}

	//JWT Stuff
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//Set Cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"Message": "Success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	var user models.User

	db.Database.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"Message": "Success",
	})
}
