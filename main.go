package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cbodonnell/til-api/config"
	"github.com/cbodonnell/til-api/handlers"
	"github.com/cbodonnell/til-api/logging"
	"github.com/cbodonnell/til-api/repositories"
	"github.com/cbodonnell/til-api/services"
)

func main() {
	// Get configuration
	ENV := os.Getenv("ENV")
	conf, err := config.ReadConfig(ENV)
	if err != nil {
		log.Fatal(err)
	}

	// create cache layer
	// cache := cache.NewRedisCache(conf)
	// err = cache.FlushDB()
	// if err != nil {
	// 	log.Println(err)
	// }
	// create repository
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Db.Host, conf.Db.Port, conf.Db.User, conf.Db.Password, conf.Db.Dbname)
	repo := repositories.NewPSQLTilRepository(connStr)
	defer repo.Close()
	// create service
	service := services.NewStandardTilService(repo)
	// create handler
	handler := handlers.NewMuxHandler(handlers.MuxHandlerOptions{
		TilService:   service,
		AuthEndpoint: conf.Auth,
	})
	if conf.AllowedOrigins != "" {
		handler.AllowOrigins(strings.Split(conf.AllowedOrigins, ","))
	}
	r := handler.GetRouter()

	// create worker
	// worker := workers.NewExampleWorker(conf)
	// go worker.Start()

	// Set log file
	if conf.LogFile != "" {
		logFile := logging.SetLogFile(conf.LogFile)
		defer logFile.Close()
	}

	// Run server
	log.Println(fmt.Sprintf("Serving on port %d", conf.Port))

	// TLS
	if conf.SSLCert == "" {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), r))
	}
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", conf.Port), conf.SSLCert, conf.SSLKey, r))
}
