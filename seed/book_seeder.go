package seed

import (
	"log"

	"github.com/andijayawizard/go-fiber-api/database"
	"github.com/andijayawizard/go-fiber-api/models"
	"github.com/brianvoe/gofakeit/v6"
)

func SeedBooks() {
	// Generate 10 buku random
	for i := 0; i < 6; i++ {
		book := models.Book{
			Title:  gofakeit.Sentence(3), // 3 kata judul
			Author: gofakeit.Name(),      // Nama random
		}

		result := database.DB.Create(&book)
		if result.Error != nil {
			log.Printf("❌ Failed to insert book: %v\n", result.Error)
		}
	}

	log.Println("✅ 10 Dummy books generated with Faker!")
}
