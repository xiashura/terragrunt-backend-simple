package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.AutomaticEnv()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("stage %s", v.GetString("ENV")))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
