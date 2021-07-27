package middleware

import (
	"errors"
	"github.com/CoryKelly/Admin_App/db"
	"github.com/CoryKelly/Admin_App/models"
	"github.com/CoryKelly/Admin_App/util"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")
	Id, err := util.ParseJwt(cookie)
	if err != nil {
		return err
	}

	userId, _ := strconv.Atoi(Id)

	user := models.User{
		Id: uint(userId),
	}
	db.Database.Preload("Role").Find(&user)

	role := models.Role{
		Id: user.RoleId,
	}
	db.Database.Preload("Permissions").Find(&role)

	//Loop Permissions
	if c.Method() == "GET" {
		for _, permissions := range role.Permissions {
			if permissions.Name == "view_"+page || permissions.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permissions := range role.Permissions {
			if permissions.Name == "edit_"+page {
				return nil
			}
		}
	}
	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")
}
