package restapi

import (
	"github.com/labstack/echo/v4"
	"github.com/mcarello/user_service/internal/app/users/database"
	"github.com/mcarello/user_service/internal/app/users/restapi/user"
)

func NewRoutes(e *echo.Echo, db *database.DB) {
	user.RegisterHandlers(e, user.NewService(user.NewRepository(db)))
}
