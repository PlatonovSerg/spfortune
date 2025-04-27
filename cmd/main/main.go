package main

import (
	"log"

	"SparFortuneDDD/internal/app"
)

func main() {
	application := app.NewApp()               // Создаём экземпляр приложения
	if err := application.Run(); err != nil { // Запускаем приложение
		log.Fatalf("Failed to run app: %v", err)
	}
}
