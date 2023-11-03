package users

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mcarello/user_service/internal/app/users/configuration"
	"github.com/mcarello/user_service/internal/app/users/database"
	"github.com/mcarello/user_service/internal/app/users/restapi"
)

func init() {
	// read env
	configuration.ReadConfiguration()
}

func RunServer() {
	// Connect to Database
	db := database.ConnectDatabase()

	log.Println("Run Server")
	e := echo.New()

	// Initial Beans
	restapi.NewRoutes(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
