package main

import (
	"fmt"

	"github.com/gin-contrib/cors"

	// a/app/crud"
	// "github.com/ppphp/yayagf/pkg/model"
	"a/app/config"
	"a/app/router"
	"log"

	"github.com/gin-gonic/gin"
)

// @title "a API
// @version master
// @description This is a a server

// @contact.name your name
// @contact.url https://a
// @contact.email your email

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	router.AddRoute(r)

	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	//drv, err := model.Open("mysql", config.GetConfig().DB)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//crud.C = crud.NewClient(crud.Driver(drv))
	//if err := crud.C.Schema.Create(context.Background()); err != nil {
	//	log.Fatal(err)
	//}

	if err := r.Run(fmt.Sprintf(":%v", config.GetConfig().Port)); err != nil {
		log.Fatal(err)
	}
}
