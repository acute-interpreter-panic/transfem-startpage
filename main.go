package main

import (
	"bytes"
	"fmt"
	"net/http"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/rendering"
	"github.com/labstack/echo/v4"
)

var CurrentRenderingConfig = rendering.DefaultRenderingConfig()

func getIndex(c echo.Context) error {
	var tpl bytes.Buffer
	rendering.IndexTemplate.Execute(&tpl, CurrentRenderingConfig)

	return c.HTML(http.StatusOK, tpl.String())
}

func main() {
	fmt.Println("running transfem startpage")

	e := echo.New()
	e.Static("/assets", "frontend/assets")
	e.GET("/", getIndex)

	e.Logger.Fatal(e.Start(":5500"))
}
