package main

import (
	"log"
	"net/http"

	c "github.com/AsrofunNiam/lets-code-elastic-search/configuration"

	"github.com/AsrofunNiam/lets-code-elastic-search/app"
	"github.com/AsrofunNiam/lets-code-elastic-search/helper"
	"github.com/go-playground/validator/v10"
)

func main() {
	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	port := configuration.Port
	db := app.ConnectDatabase(configuration.User, configuration.Host, configuration.Password, configuration.PortDB, configuration.Db)
	elasticClient := app.ConnectionElastic(configuration.ElasticHost, configuration.ElasticPort, configuration.ElasticUser, configuration.ElasticPassword)

	validate := validator.New()
	router := app.NewRouter(elasticClient, db, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
