package main

import (
	"embed"

	"gitea.elara.ws/Hazel/transfem-startpage/internal/cli"
	"gitea.elara.ws/Hazel/transfem-startpage/internal/server"
)

//go:embed frontend/*
var FrontendFiles embed.FS

func main() {
	server.FrontendFiles = FrontendFiles
	cli.Cli()
}
