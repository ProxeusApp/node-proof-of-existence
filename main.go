package main

import (
	"context"

	"github.com/ProxeusApp/node-proof-of-existence/controller"
	"github.com/ProxeusApp/node-proof-of-existence/service"
	"github.com/ProxeusApp/proxeus-core/externalnode"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())

	configuration, err := ReadConfiguration()
	if err != nil {
		panic("reading configuration: " + err.Error())
	}

	twitterService := service.NewDefaultTwitterService(context.Background(), configuration.Get("TWITTER_CONSUMER_KEY"), configuration.Get("TWITTER_CONSUMER_SECRET"))
	controller := controller.NewController(twitterService)

	g := e.Group("/node/:id")
	conf := middleware.DefaultJWTConfig
	conf.SigningKey = []byte(configuration.Get("SERVICE_SECRET"))
	conf.TokenLookup = "query:" + configuration.Get("AUTH_KEY")
	g.Use(middleware.JWTWithConfig(conf))

	g.POST("/next", controller.GetTweetByURLOrID)
	g.GET("/config", externalnode.Nop)
	g.POST("/config", externalnode.Nop)
	g.POST("/remove", externalnode.Nop)
	g.POST("/close", externalnode.Nop)
	g.GET("/health", externalnode.Health)

	err = externalnode.Register(configuration.Get("PROXEUS_INSTANCE_URL"), configuration.Get("SERVICE_NAME"), configuration.Get("SERVICE_URL"), configuration.Get("SERVICE_SECRET"), "Retrieves token balances of an address", 2)
	if err != nil {
		panic("couldn't register: " + err.Error())
	}

	err = e.Start(":" + configuration.Get("SERVICE_PORT"))
	if err != nil {
		panic(err.Error())
	}
}
