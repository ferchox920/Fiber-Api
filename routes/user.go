package routes

import (
	"github.com/ferchox920/fiber-api/database"
	"github.com/ferchox920/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	// This is not the model, more like a serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(user models.User) User {
	return User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func FindUserByID(c *fiber.Ctx) error {
	userID := c.Params("id") // Obtener el ID de los parámetros de la URL

	var user models.User
	result := database.Database.Db.First(&user, userID)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Usuario no encontrado"})
	}

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id") 

	var user models.User
	result := database.Database.Db.First(&user, userID)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Usuario no encontrado"})
	}

	var updateUser models.User
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	// Excluir la modificación del campo ID
	updateUser.ID = user.ID

	// Actualizar otros campos del usuario
	database.Database.Db.Model(&user).Omit("ID").Updates(updateUser) // Utiliza Omit para excluir el campo ID
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id") // Obtener el ID de los parámetros de la URL

	var user models.User
	result := database.Database.Db.First(&user, userID)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Usuario no encontrado"})
	}

	database.Database.Db.Delete(&user)
	return c.Status(204).JSON(nil) // 204 significa que no hay contenido en la respuesta
}
func FindAllUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)

	responseUsers := make([]User, len(users))
	for i, user := range users {
		responseUsers[i] = CreateResponseUser(user)
	}

	return c.Status(200).JSON(responseUsers)
}
