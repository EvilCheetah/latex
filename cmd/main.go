package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

type Carrier struct {
	Name        string
	MCNumber    string
	USDOTNumber string
	Address     string
	City        string
	State       string
	PostalCode  string
	PhoneNumber string
	Email       string
}

type Vehicle struct {
	Year         string
	Make         string
	Model        string
	LicensePlate string
	VIN          string
}

type Trailer struct {
	Year          string
	Make          string
	TrailerNumber string
	LicensePlate  string
	VIN           string
}

type Commission struct {
	GrossRevenue          string
	CashAdvanceServiceFee string
}

func (t *Templates) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Static("/static", "assets")

	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.POST("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
