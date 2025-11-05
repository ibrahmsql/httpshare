package main

import (
	_ "embed"
	"flag"
	"fmt"
	"net/url"

	"github.com/isa0-gh/httpshare/template"
	"github.com/isa0-gh/httpshare/utils"
	"github.com/labstack/echo/v4"
)

//go:embed template/tailwind.js
var tailwind string

func main() {
	e := echo.New()
	e.HideBanner = true

	e.GET("/tailwind.js", func(c echo.Context) error {
		return c.String(200, tailwind)
	})

	e.GET("/*", func(c echo.Context) error {
		decoded, _ := url.QueryUnescape(c.Param("*"))
		path := utils.UrlToFilePath(decoded)
		if path == "" {
			path = "."
		}
		if c.QueryParam("download") == "true" {
			return c.File(path)
		}
		data, err := utils.GetFiles(path)
		if err != nil {
			return c.String(500, err.Error())
		}
		output, err := template.Render(data)
		if err != nil {
			return c.String(500, err.Error())
		}

		return c.HTML(200, output)
	})

	port := flag.Int("port", 8080, "Port number to run the server on")
	flag.Parse()
	fmt.Printf("Listening on http://localhost:%d/\n", *port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
