package controllers

import (
	"encoding/csv"
	"github.com/CoryKelly/Admin_App/db"
	"github.com/CoryKelly/Admin_App/middleware"
	"github.com/CoryKelly/Admin_App/models"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
)

func GetAllOrders(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "orders"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(db.Database, &models.Order{}, page))
}

func Export(c *fiber.Ctx) error {
	filePath := "./csv/orders.csv"
	if err := CreateFile(filePath); err != nil {
		return err
	}
	return c.Download(filePath)
}

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	var orders []models.Order

	db.Database.Preload("OrderItems").Find(&orders)

	writer.Write([]string{
		"ID", "Name", "Email", "Product Title", "Price", "Quantity",
	})

	for _, order := range orders {
		data := []string{
			strconv.Itoa(int(order.Id)),
			order.FirstName + " " + order.LastName,
			order.Email,
			"", //Product Title
			"", //Price
			"", //Quantity
		}
		if err := writer.Write(data); err != nil {
			return err
		}

		for _, orderItem := range order.OrderItems {
			data := []string{
				"", //Product Title
				"", //Price
				"", //Quantity
				orderItem.ProductTitle,
				strconv.Itoa(int(orderItem.Price)),
				strconv.Itoa(int(orderItem.Quantity)),
			}
			if err := writer.Write(data); err != nil {
				return err
			}
		}
	}
	return nil
}

type Sales struct {
	Date string `json:"date"`
	Sum  string `json:"sum"`
}

func Chart(c *fiber.Ctx) error {
	var sales []Sales

	db.Database.Raw(`
			SELECT DATE_FORMAT(o.create_at, '%Y-%m-%d') as date, SUM(oi.price * oi.quantity) as sum
			FROM orders o
			JOIN order_items oi on o.id = oi.order_id
			GROUP BY date
			`).Scan(&sales)

	return c.JSON(sales)
}
