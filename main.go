package main

import (
	"context"
	"github.com/nodias/golang-oauth2.0-client/router"
	"github.com/nodias/golang-oauth2.0-common/models"
	"github.com/nodias/golang-oauth2.0-common/shared/logger"
	"github.com/nodias/golang-oauth2.0-common/shared/repository"

	"github.com/urfave/negroni"
)

var config models.TomlConfig

func init() {
	models.Load("config/%s/config.toml")
	config = *models.GetConfig()
	logger.Init()
	repository.Init()
	repository.NewOpenDB()
}

func main() {
	log := logger.New(context.Background())
	n := negroni.New()
	n.UseHandler(router.NewRouter())
	log.Info("Client - Server On!")
	n.Run(config.Servers["Client"].PORT)
}
