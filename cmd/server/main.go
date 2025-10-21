package main

import (
	"log"
	"net/http"

	"KateBaklanova.com/TIP_PR6/internal/db"
	"KateBaklanova.com/TIP_PR6/internal/httpapi"
	"KateBaklanova.com/TIP_PR6/internal/models"
)

func main() {
	d := db.Connect()

	// Автоматически создаст (или обновит) таблицы под наши модели
	if err := d.AutoMigrate(&models.User{}, &models.Note{}, &models.Tag{}); err != nil {
		log.Fatal("migrate:", err)
	}

	r := httpapi.BuildRouter(d)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
