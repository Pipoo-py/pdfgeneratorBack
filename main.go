package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	generatepdf "pdfgen/generator"
	pdfstruct "pdfgen/struct"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://izpdfgenerator.netlify.app/"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", getInit)
	e.POST("/generate-simple", handlerGenerateSimple)
	e.POST("/generate-invoice", handlerGenerateInvoice)

	e.Logger.Fatal(e.Start(":8000"))
}

func handlerGenerateSimple(c echo.Context) error {
	input := new(pdfstruct.SimpleInput)
	if err := c.Bind(input); err != nil {
		return c.String(http.StatusBadRequest, "El formato ingresado es inválido")
	}

	var pdfBuffer bytes.Buffer
	err := generatepdf.GenSimplePDF(&pdfBuffer, input)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error al momento de guardar el pdf en buffer: %v", err))
	}
	c.Response().Header().Set("Content-Type", "application/pdf")
	c.Response().Header().Set("Content-Disposition", `attachment; filename="document.pdf"`)

	err = c.Stream(http.StatusOK, "application/pdf", &pdfBuffer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("ocurrió un error interno al momento de generar o enviar el pdf: %w", err))
	}

	return nil
}

func handlerGenerateInvoice(c echo.Context) error {
	input := new(pdfstruct.InvoiceInput)
	if err := c.Bind(input); err != nil {
		return c.String(http.StatusBadRequest, "El formato ingresado es inválido")
	}

	var pdfBuffer bytes.Buffer
	err := generatepdf.GenInvoicePDF(&pdfBuffer, input)
	if err != nil {
		fmt.Println("Error al generar el pdf")
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error al momento de guardar el pdf en buffer: %v", err))
	}
	c.Response().Header().Set("Content-Type", "application/pdf")
	c.Response().Header().Set("Content-Disposition", `attachment; filename="document.pdf"`)

	err = c.Stream(http.StatusOK, "application/pdf", &pdfBuffer)
	if err != nil {
		log.Fatal("error al ejecutar c.Stream")
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("ocurrió un error interno al momento de generar o enviar el pdf: %w", err))
	}

	return nil
}

func getInit(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo")
}
