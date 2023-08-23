package routes

import (
	"github.com/ferchox920/fiber-api/database"
	"github.com/ferchox920/fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product struct {
	// This is not the model, more like a serializer
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
	Description  string `json:"description"`
}

func CreateResponseProduct(product models.Product) Product {
	return Product{ID: product.ID, Name: product.Name, SerialNumber: product.SerialNumber, Description: product.Description}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Verificar si el código de producto ya existe en la base de datos
	existingProduct := models.Product{}
	if err := database.Database.Db.Where("serial_number = ?", product.SerialNumber).First(&existingProduct).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error de base de datos"})
		}
	}

	if existingProduct.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "El código de producto ya está en uso"})
	}

	if err := database.Database.Db.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error de base de datos"})
	}

	responseProduct := CreateResponseProduct(product)
	return c.Status(fiber.StatusCreated).JSON(responseProduct)
}

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	if err := database.Database.Db.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al obtener los productos de la base de datos"})
	}

	responseProducts := make([]Product, len(products))
	for i, product := range products {
		responseProducts[i] = CreateResponseProduct(product)
	}

	return c.Status(fiber.StatusOK).JSON(responseProducts)
}

func FindProductByID(c *fiber.Ctx) error {
	productID := c.Params("id")
	var product models.Product
	result := database.Database.Db.First(&product, productID)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Usuario no encontrado"})
	}

	responseUser := CreateResponseProduct(product)
	return c.Status(200).JSON(responseUser)
}

func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var product models.Product
	result := database.Database.Db.First(&product, productID)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Producto no encontrado"})
	}
	var productUpdate Product
	if err := c.BodyParser(&productUpdate); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	// Verificar si el nuevo número de serie ya existe en la base de datos
	if product.SerialNumber != productUpdate.SerialNumber {
		existingProduct := models.Product{}
		if err := database.Database.Db.Where("serial_number = ?", productUpdate.SerialNumber).First(&existingProduct).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error de base de datos"})
			}
		}
		if existingProduct.ID != 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "El nuevo número de serie ya está en uso"})
		}
	}
	// Actualizar los campos modificables del producto
	if productUpdate.Name != "" {
		product.Name = productUpdate.Name
	}
	if productUpdate.SerialNumber != "" {
		product.SerialNumber = productUpdate.SerialNumber
	}
	if productUpdate.Description != "" {
		product.Description = productUpdate.Description
	}

	if err := database.Database.Db.Save(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar el producto en la base de datos"})
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(fiber.StatusOK).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("id")
	var product models.Product
	result := database.Database.Db.First(&product, productID)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Producto no encontrado"})
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar el producto de la base de datos"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Producto eliminado correctamente"})
}
