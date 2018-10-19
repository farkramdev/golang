package main

import (
	// "github.com/farkramdev/golang/api"
	// "github.com/farkramdev/golang/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

const projectID = "jwt-authen"

func main() {
	// serviceAccount, err := ioutil.ReadFile("service-account.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = api.Init(api.Config{
	// 	ServiceAccountJSON: serviceAccount,
	// 	ProjectID:          projectID,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	e := echo.New()

	e.Use(
		middleware.Recover(),
		middleware.Secure(),
		middleware.Logger(),
		middleware.Gzip(),
		middleware.BodyLimit("2M"),
		middleware.CSRFWithConfig(middleware.CSRFConfig{
			TokenLookup: "header:X-XSRF-TOKEN",
		}),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{
				"http://localhost:8080",
			},
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderContentType,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				echo.GET,
				echo.POST,
			},
			MaxAge: 3600,
		}),
	)
	// Health check
	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK KKKKK")
	})

	// Register services
	// service.Auth(e.Group("/auth"))

	e.Start(":9000")

	// e.Start(standard.WithConfig(engine.Config{
	// 	Address:      ":9000",
	// 	ReadTimeout:  30 * time.Second,
	// 	WriteTimeout: 30 * time.Second,
	// }))
}
