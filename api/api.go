package api

import (
	"errors"
	"net/http"

	_ "gitea.repetitra.ru/StudBank/Backend/docs"
	"gitea.repetitra.ru/StudBank/Backend/etc"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

type API struct {
	*echo.Echo
	MLog *logrus.Logger
}

func (a *API) Init() (*API, error) {
	a.Echo = echo.New()
	a.MLog = etc.GetLogger("api", logrus.TraceLevel)

	a.applyMiddlewares()

	routes := Routes{Log: a.MLog}
	initRoutes(a, routes)

	return a, nil
}

type Routes struct {
	Log *logrus.Logger
	//DB *db.DB
}

func initRoutes(a *API, routes Routes) {
	a.GET("/hello", routes.Hello)
}

func (a *API) MRun() error {
	a.Use(echoprometheus.NewMiddleware("myapp"))
	go func() {
		metrics := echo.New()
		metrics.GET("/metrics", echoprometheus.NewHandler())
		if err := metrics.Start(":8081"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.MLog.WithError(err).Fatal("Metrics server stopped due to the error!")
		}
	}()

	go func() {
		swagger := echo.New()
		swagger.Any("/*", echoSwagger.WrapHandler)
		if err := swagger.Start(":8082"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.MLog.WithError(err).Fatal("Swagger server stopped due to the error!")
		}
	}()

	return a.Start(":8080")
}
