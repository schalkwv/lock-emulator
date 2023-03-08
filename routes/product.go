package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/schalkwv/lock-emulator/database"
	"github.com/schalkwv/lock-emulator/models"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}
type UpdateProduct struct {
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func createResponseProduct(m models.Product) Product {
	return Product{
		ID:           m.ID,
		Name:         m.Name,
		SerialNumber: m.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var p models.Product
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&p)
	response := createResponseProduct(p)
	return c.Status(200).JSON(response)
}
func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product not found")
	}
	return nil
}
