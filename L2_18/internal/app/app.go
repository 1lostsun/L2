package app

import (
	"github.com/1lostsun/L2/tree/main/L2_18/internal/handler"
	logger "github.com/1lostsun/L2/tree/main/L2_18/internal/pkg"
	"github.com/1lostsun/L2/tree/main/L2_18/internal/repo"
	"github.com/1lostsun/L2/tree/main/L2_18/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

// Run : Запускает приложение
func Run() {

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file")
		os.Exit(1)
	}

	port := os.Getenv("PORT")

	Repo := repo.New()
	UseCase := usecase.New(Repo)
	Handler := handler.New(UseCase)
	r := gin.New()

	Handler.InitRoutes(r)

	logger.Info("Running server on port %s", port)
	if err := r.Run(port); err != nil {
		logger.Error("Error loading .env file %v", err)
		os.Exit(1)
	}
}
