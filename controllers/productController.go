package controllers

import (
	"github.com/CoryKelly/Admin_App/db"
	"github.com/CoryKelly/Admin_App/middleware"
	"github.com/CoryKelly/Admin_App/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetAllProducts(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "products"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(db.Database, &models.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "products"); err != nil {
		return err
	}

	var products models.Product

	if err := c.BodyParser(&products); err != nil {
		return err
	}

	db.Database.Create(&products)

	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "products"); err != nil {
		return err
	}

	//Get id from url
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	db.Database.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "products"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	db.Database.Model(&product).Updates(&product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "products"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	db.Database.Delete(&product)

	return nil
}
