package main

import (
	"github.com/isa0-gh/httpshare/template"
	"github.com/isa0-gh/httpshare/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/*", func(c echo.Context) error {
		path := utils.UrlToFilePath(c.Param("*"))
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

	e.Logger.Fatal(e.Start(":8080"))
}
