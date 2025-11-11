package main

import (
	"log"
	"net/http"

	"github.com/ShadowNos1/qa-api/internal/app"
)

func main() {
	// Создаем "сервис" в памяти (для теста без БД)
	svc := app.NewInMemoryService()

	// Создаем обработчик HTTP с этим сервисом
	handler := app.NewHandler(svc)

	// Создаем роутер
	router := app.NewRouter(handler)

	// Запускаем сервер на порту 8080
	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
