package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/diyhrt"
	"gitea.elara.ws/Hazel/transfem-startpage/internal/rendering"
	"github.com/labstack/echo/v4"
)


var CurrentConfig = rendering.NewConfig()


func FetchDiyHrt() error {
	fmt.Println("Fetch DiyHrt Marketplaces...")

	l, err := diyhrt.GetListings()
	if err != nil {
		return err
	}
	CurrentConfig.Template.LoadDiyHrt(l)
	return nil
}

func getIndex(c echo.Context) error {
	var tpl bytes.Buffer
	rendering.IndexTemplate.Execute(&tpl, CurrentConfig.Template)

	return c.HTML(http.StatusOK, tpl.String())
}

func main() {
	fmt.Println("running transfem startpage")

	profile := "startpage"
	if len(os.Args) > 1 {
		profile = os.Args[1]
	}
	fmt.Println("loading profile " + profile + "...")

	err := CurrentConfig.ScanForConfigFile(profile)
	if err != nil {
		fmt.Println(err)
	}

	err = FetchDiyHrt()
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()
	e.Static("/assets", "frontend/assets")
	e.Static("/scripts", "frontend/scripts")
	e.GET("/", getIndex)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(CurrentConfig.Server.Port)))
}
