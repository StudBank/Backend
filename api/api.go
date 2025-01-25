package api

import (
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type APIConfig struct {
}

type API struct {
	*echo.Echo
	MLog zerolog.Logger
}

func (a *API) Init(conf APIConfig) (*API, error) {
	a.Echo = echo.New()
	a.MLog = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	a.applyMiddlewares()

	routes := Routes{Log: a.MLog}
	initRoutes(a, routes)

	return a, nil
}

type Routes struct {
	Log zerolog.Logger
	//DB *db.DB
}

func initRoutes(a *API, routes Routes) {
	a.GET("/hello", routes.Hello)
}

func (a *API) MRun() {
	c := jaegertracing.New(a.Echo, nil)
	defer c.Close()
	a.Use(echoprometheus.NewMiddleware("myapp"))

	go func() {
		metrics := echo.New()
		metrics.GET("/metrics", echoprometheus.NewHandler())
		if err := metrics.Start(":8081"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.MLog.Err(err).Msg("Metrics server stopped due to the error!")
		}
	}()

	a.Logger.Fatal(a.Start(":1323"))
}
