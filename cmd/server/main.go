package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Kikozai/Moon/config"
	"github.com/Kikozai/Moon/controllers"
	dbConn "github.com/Kikozai/Moon/db/sqlc"
	"github.com/Kikozai/Moon/routes"

	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	db     *dbConn.Queries

	AuthController controllers.AuthController
	AuthRoutes     routes.AuthRoutes
)

func init() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	conn, err := sql.Open(config.PostgreDriver, config.PostgresSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	db = dbConn.New(conn)

	fmt.Println("PostgreSQL connected successfully...")

	AuthController = *controllers.NewAuthController(db)
	AuthRoutes = routes.NewAuthRoutes(AuthController)

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	router := server.Group("/api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Welcome to Golang with PostgreSQL"})
	})

	AuthRoutes.AuthRoute(router)
        server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": 
                       fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})
	log.Fatal(server.Run(":" + config.Port))
}

