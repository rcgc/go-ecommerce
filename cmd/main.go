package main

import (
	"log"
	"os"

	"github.com/rcgc/go-ecommerce/infrastructure/handler"
	"github.com/rcgc/go-ecommerce/infrastructure/handler/response"
)

func main(){
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = validateEnvironments()
	if err != nil {
		log.Fatal(err)
	}

	e := newHTTP(response.HTTPErrorHandler)

	dbPool, err := newDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoutes(e, dbPool)
	_ = dbPool

	if os.Getenv("IS_HTTPS") == "true" {
		err = e.StartTLS(":"+os.Getenv("SERVER_PORT"), os.Getenv("CERT_PEM_FILE"), os.Getenv("KEY_PEM"))
	} else {
		err = e.Start(":"+os.Getenv("SERVER_PORT"))
	}
	if err != nil {
		log.Fatal(err)
	}
}
