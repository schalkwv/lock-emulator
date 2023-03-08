package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/schalkwv/lock-emulator/database"
	"github.com/schalkwv/lock-emulator/models"
	"time"
)

type Order struct {
	ID        uint      `json:"id"`
	Product   Product   `json:"product"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

func createResponseOrder(m models.Order, u User, p Product) Order {
	return Order{
		ID:        m.ID,
		Product:   p,
		User:      u,
		CreatedAt: m.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	if err := findUser(int(order.UserId), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var product models.Product
	if err := findProduct(int(order.ProductId), &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)
	responseUser := createResponseUser(user)
	responseProduct := createResponseProduct(product)

	responseOrder := createResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}
