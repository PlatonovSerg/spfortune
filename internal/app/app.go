package app

import (
	"SparFortuneDDD/config"
	"SparFortuneDDD/internal/db"
	"SparFortuneDDD/internal/modules/authentication"
	"SparFortuneDDD/internal/routes"
	"SparFortuneDDD/pkg"
	"log"
	"net/http"
	"time"
)

// App - структура для управления зависимостями
type App struct {
	Config      config.Config
	JWTService  *pkg.JWTService
	AuthService *authentication.Service
	AuthHandler *authentication.Handler
	Router      http.Handler
}

// NewApp - инициализация приложения
func NewApp() *App {
	// Загружаем конфигурацию
	config.LoadEnv()

	// Создаём экземпляр JWTService
	jwtService := pkg.NewJWTService(
		config.AppConfig.JWTSecret,
		"SparFortuneApp",
		time.Hour*72,
	)

	// Инициализируем базу данных
	db.InitDB(config.AppConfig.DSN)

	// Создаём сервисы и хендлеры
	authService := authentication.NewService(jwtService)
	authHandler := authentication.NewHandler(authService)

	// Настраиваем маршрутизатор
	router := routes.SetupRouter(jwtService, authHandler)

	return &App{
		Config:      config.AppConfig,
		JWTService:  jwtService,
		AuthService: authService,
		AuthHandler: authHandler,
		Router:      router,
	}
}

// Run - запуск приложения
func (a *App) Run() error {
	log.Println("Starting server on port", a.Config.Port)
	return http.ListenAndServe(":"+a.Config.Port, a.Router)
}
