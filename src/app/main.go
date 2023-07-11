package main

import (
	"fmt"
	"log"

	"mongo-test/api"
	"mongo-test/controllers"
	"mongo-test/db"

	"github.com/caarlos0/env"
	"github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
)

type AppEnv struct {
	// Port is the port number to listen on
	Port       string `env:"APP_PORT" envDefault:"8888"`
	DBHostname string `env:"MONGO_HOSTNAME,required"`
	DBUsername string `env:"MONGO_INITDB_ROOT_USERNAME,required"`
	DBPassword string `env:"MONGO_INITDB_ROOT_PASSWORD,required"`
	DBName     string `env:"MONGO_INITDB_DATABASE,required"`
}

func main() {
	// ENV parsing
	appEnv := AppEnv{}
	err := env.Parse(&appEnv)
	if err != nil {
		log.Fatal(err)
	}

	// DB connection
	clientUri := fmt.Sprintf("mongodb://%s:%s@%s:27017", appEnv.DBUsername, appEnv.DBPassword, appEnv.DBHostname)
	client, err := db.GetClient(clientUri)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(appEnv.DBName)

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Validate server per the swagger spec
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal("Error loading swagger spec: ", err)
	}
	r.Use(middleware.OapiRequestValidator(swagger))

	// 404 Default handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// V1 API
	v1 := r.Group("/api/v1")
	server := controllers.Server{DB: db}
	options := api.GinServerOptions{}
	// Create a StrictServer with the API config
	ssi := api.NewStrictHandler(server, []api.StrictMiddlewareFunc{})
	// Registers the handlers per the config
	api.RegisterHandlersWithOptions(v1, ssi, options)
	// api.RegisterHandlersWithOptions(r, ssi, options)

	// Listen and serve on
	r.Run(fmt.Sprintf(":%s", appEnv.Port))
}
