package cli

import (
	"log"
	"os"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/server"
)

func Start() error {
	profile := "default"
	if len(os.Args) > 2 {
		profile = os.Args[2]
	}
	log.Println("starting server with profile " + profile)

	return server.Start(profile)
}
