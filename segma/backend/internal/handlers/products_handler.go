package handlers

import "github.com/gofiber/fiber/v2"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

func Products(c *fiber.Ctx) error {
	products := []Product{
		{
			ID:          1,
			Name:        "Paracetamol",
			Description: "Pain relief medicine",
			Price:       120.00,
			Stock:       50,
		},
		{
			ID:          2,
			Name:        "Vitamin C",
			Description: "Immune support supplement",
			Price:       350.00,
			Stock:       25,
		},
	}

	return c.JSON(fiber.Map{
		"data": products,
	})
}
