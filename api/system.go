package api

import (
	"github.com/labstack/echo/v4"
)

func (r *Routes) Hello(c echo.Context) error {
	c.JSON(200, echo.Map{"ans": "hello!"})
	return nil
}
