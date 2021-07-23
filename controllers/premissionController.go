package controllers

import (
	"github.com/CoryKelly/Admin_App/db"
	"github.com/CoryKelly/Admin_App/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	db.Database.Preload("Role").Find(&permissions)

	return c.JSON(permissions)
}
