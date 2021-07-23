package controllers

import (
	"github.com/CoryKelly/Admin_App/db"
	"github.com/CoryKelly/Admin_App/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	db.Database.Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	db.Database.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	//Get id from url
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	db.Database.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	var result interface{}
	db.Database.Table("role_permissions").Where("role_id", id).Delete(result)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	db.Database.Model(&role).Updates(&role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	db.Database.Delete(&role)

	return nil
}
