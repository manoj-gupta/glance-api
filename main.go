package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/manoj-gupta/glance/internal/db"
	"github.com/manoj-gupta/glance/internal/routes"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Initialize app
	app := App{}
	err := app.Initialize()
	if err != nil {
		log.Fatal("App Initialization failed")
		return
	}

	// Clean up app
	defer app.DeInitialize()

	// Start app
	app.Run()
}

// App .. structure to keep app constructs
type App struct {
	router *gin.Engine
	db     *sql.DB
}

// Initialize ... init app
func (a *App) Initialize() error {
	var err error

	a.db, err = db.Initialize()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Create the router
	a.router = gin.Default()

	// Initialize the routes
	routes.Init(a.router)
	routes.InitializeRoutes(a.router)

	return nil
}

// Run ... run app
func (a *App) Run() {
	// Run router
	a.router.Run()
}

// DeInitialize ... deinit app
func (a *App) DeInitialize() {
	a.db.Close()
}
