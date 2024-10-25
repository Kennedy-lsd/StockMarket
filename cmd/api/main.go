package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Kennedy-lsd/StockMarket/config"
	"github.com/Kennedy-lsd/StockMarket/internal/handlers"
	"github.com/Kennedy-lsd/StockMarket/internal/repos"
	"github.com/Kennedy-lsd/StockMarket/internal/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializeDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
		return
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to the database!")
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}

func main() {
	initEnv()
	apiConfig := config.LoadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		apiConfig.DB_HOST, apiConfig.DB_USER, apiConfig.DB_PASSWORD, apiConfig.DB_NAME, apiConfig.DB_PORT, apiConfig.DB_SSLMODE)

	InitializeDB(dsn)

	api := echo.New()

	stockRepository := repos.NewStockRepository(DB)
	stockService := services.NewStockService(stockRepository)
	stockHandler := handlers.NewStockHandler(stockService)

	commentRepository := repos.NewCommentRepository(DB)
	commentService := services.NewCommentService(commentRepository)
	commentHandler := handlers.NewCommentHandler(commentService)

	stockR := api.Group("/api")
	stockR.GET("/stocks", stockHandler.GetAllStocks)
	stockR.GET("/stocks/:id", stockHandler.GetStockById)
	stockR.POST("/stocks", stockHandler.CreateStock)
	stockR.DELETE("/stocks/:id", stockHandler.DeleteStockById)

	stockC := api.Group("/api")
	stockC.GET("/comments", commentHandler.GetAllComments)
	stockC.GET("/comments/:id", commentHandler.GetCommentById)
	stockC.POST("/comments", commentHandler.CreateComment)
	stockC.DELETE("/comments/:id", commentHandler.DeleteCommentById)
	stockC.PATCH("/comments/:id", commentHandler.UpdateCommentById)

	api.Start(fmt.Sprintf(":%v", apiConfig.SERVER_PORT))

}
