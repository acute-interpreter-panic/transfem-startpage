package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/cache"
	"gitea.elara.ws/Hazel/transfem-startpage/internal/rendering"
	"github.com/labstack/echo/v4"
)

var Config = rendering.NewConfig()

func StartFetching() {
	for {
		log.Println("Fetch DiyHrt data...")
		Config.FetchDiyHrt()
		time.Sleep(time.Duration(Config.DiyHrt.FetchIntervals) * time.Second)

		if Config.DiyHrt.FetchIntervals == 0 {
			break
		}
	}
}

func Start(profile string) error {
	err := Config.ScanForConfigFile(profile)
	if err != nil {
		return err
	}

	go StartFetching()

	err = Config.FetchDiyHrt()
	if err != nil {
		log.Println(err)
	}

	e := echo.New()

	// statically serve the file
	cache := cache.NewCache(profile)
	if !cache.Disabled {
		cache.StartStaticServer(e)

		log.Println("downloading website icons...")
		for i, w := range Config.Template.Websites {
			u, err := cache.CacheUrl(w.ImageUrl)
			if err != nil {
				log.Println(err)
			}
			Config.Template.Websites[i].ImageUrl = u
			Config.Template.Websites[i].IsFetched = true
		}
	}

	// https://echo.labstack.com/docs/cookbook/embed-resources
	staticHandler := http.FileServer(getFileSystem())
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/", staticHandler)))
	e.GET("/scripts/*", echo.WrapHandler(http.StripPrefix("/", staticHandler)))
	e.GET("/", getIndex)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(Config.Server.Port)))
	return nil
}
