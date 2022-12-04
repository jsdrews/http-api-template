package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"mongo-test/routes"
)

type config struct {
	// Port is the port number to listen on
	Port string `env:"APP_PORT" envDefault:"8888"`
	DBHostname string `env:"MONGO_HOSTNAME,required"`
	DBUsername string `env:"MONGO_INITDB_ROOT_USERNAME,required"`
	DBPassword string `env:"MONGO_INITDB_ROOT_PASSWORD,required"`
	DBName string `env:"MONGO_INITDB_DATABASE,required"`
}

func main() {
	// ENV parsing
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	// DB connection
	

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// 404 Default handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// V1 group
	v1 := r.Group("/api/v1")
	v1.GET("/ping", routes.Pong)

	// Listen and serve on
	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
