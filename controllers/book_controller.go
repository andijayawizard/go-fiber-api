package controllers

import (
	"strconv"
	"strings"

	"github.com/andijayawizard/go-fiber-api/database"
	"github.com/andijayawizard/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	// Ambil query params
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")
	search := c.Query("search", "")

	// Convert string ke integer
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var books []models.Book
	query := database.DB

	// Kalau ada search query
	if search != "" {
		search = strings.ToLower(search)
		query = query.Where("LOWER(title) LIKE ? OR LOWER(author) LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var total int64
	query.Model(&models.Book{}).Count(&total)

	// Ambil data books sesuai limit + offset
	result := query.Limit(limit).Offset(offset).Find(&books)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return c.JSON(fiber.Map{
		"page":        page,
		"limit":       limit,
		"total_data":  total,
		"total_pages": totalPages,
		"search":      search,
		"data":        books,
	})
}

func GetBook(c *fiber.Ctx) error {
	var book models.Book
	id := c.Params("id")
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	database.DB.Create(book)
	return c.Status(fiber.StatusCreated).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	update := new(models.Book)
	if err := c.BodyParser(update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	book.Title = update.Title
	book.Author = update.Author
	database.DB.Save(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if err := database.DB.Delete(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
