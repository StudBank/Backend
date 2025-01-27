package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

func (a *API) applyMiddlewares() {
	a._rmTrailingSlash()

	a._recover()
	a._otelTracing()

	a._bodyLimit()
	a._logger()
	a._requestId()

	a._cors()
	a._csrf()
	a._secure()
}

func (a *API) _cors() {
	a.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
}

func (a *API) _csrf() {
	a.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "cookie:_csrf",
		CookiePath:     "/",
		CookieDomain:   "example.com",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))
}

func (a *API) _secure() {
	a.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))
}

func (a *API) _recover() {
	a.Use(middleware.Recover())
}

func (a *API) _bodyLimit() {
	a.Use(middleware.BodyLimit("2M"))
}

func (a *API) _otelTracing() {
	a.Use(otelecho.Middleware("service-name"))
}

func (a *API) _requestId() {
	a.Use(middleware.RequestID())
}
func (a *API) _rmTrailingSlash() {
	a.Pre(middleware.RemoveTrailingSlash())
}

func (a *API) _logger() {
	a.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURIPath: true,
		LogStatus:  true,
		LogLatency: true,
		LogError:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			a.MLog.Info().
				Str("URI", v.URIPath).
				Int("status", v.Status).
				Int64("latency", int64(v.Latency)).
				Err(v.Error).
				Msg("request")

			return nil
		},
	}))
}

func (a *API) _timeout() {
	a.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			a.MLog.Warn().Msg("Timeout on route")
		},
		Timeout: 30 * time.Second,
	}))
}
