package main

import (
	"fmt"
	"log"

	"server-api/api"
	"server-api/auth"
	"server-api/controllers"
	"server-api/db"

	"github.com/caarlos0/env"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
)

type AppEnv struct {
	// Port is the port number to listen on
	Port       string `env:"API_PORT" envDefault:"8888"`
	DBHostname string `env:"DB_HOSTNAME,required"`
	DBUsername string `env:"DB_ROOT_USERNAME,required"`
	DBPassword string `env:"DB_ROOT_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	DevMode    bool   `env:"API_DEV_MODE", envDefault:"false"`
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
	validatorOptions := &middleware.Options{}
	// Add authentication middleware for google id token
	if appEnv.DevMode {
		validatorOptions.Options.AuthenticationFunc = auth.DevJWTValidate
	} else {
		validatorOptions.Options.AuthenticationFunc = auth.GoogleJWTValidate
	}
	r.Use(middleware.OapiRequestValidatorWithOptions(swagger, validatorOptions))

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

	// Swagger JSON
	v1.GET("/swagger.json", func(c *gin.Context) {
		c.JSON(200, swagger)
	})

	// Listen and serve on
	r.Run(fmt.Sprintf(":%s", appEnv.Port))
}
