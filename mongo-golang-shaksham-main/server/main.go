package main

import (
	"fmt"
	"os"
	// "net/http"
	// "net/http"
	// "os"
	// "log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gorilla/handlers"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pranshult25/queriesportalbackend/common"
	"github.com/pranshult25/queriesportalbackend/router"
	// "github.com/rs/cors"
)

func main() {
	err := run()

	if err != nil {
		fmt.Println("oh no")
		panic(err)
	}
}

func run() error {

	// mux := http.NewServeMux()
	// init env
	err := common.LoadEnv()
	if err != nil {
		fmt.Println("hello")
		return err
	}

	// init db
	err = common.InitDB()
	if err != nil {
		fmt.Println("hello2")
		return err
	}

	// defer closing db
	defer common.CloseDB()

	// create app
	app := fiber.New()

	// add basic middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		// AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "Post, GET, DELETE, OPTIONS",
		AllowCredentials: true,
		MaxAge: 389312748127,
	}))
	fmt.Println("hello3")
	// app.Use(c.Handler()))

	
	// app.Use(cors.ConfigDefault.AllowCredentials == true)
	
	// add routes
	router.Router(app)
	
	

	// start server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "4000"
	}
	fmt.Println("hello4")
	app.Listen(":" + port)



	return nil
}