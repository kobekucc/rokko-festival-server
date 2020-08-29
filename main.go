package main

import (
    "log"
    "net/http"
	"os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

func main() {
    defer db.Close()

    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("$PORT must be set")
    }

    e := echo.New()

    // middleware
    e.Use(middleware.Logger())

    // routing
    e.GET("/products", indexProduct)
    e.GET("/products/:id", showProduct)

    e.POST("/products", createProduct)

    e.Logger.Fatal(e.Start(":" + port))
}

func indexProduct(c echo.Context) error {
    var products []Product
    db.Find(&products)
    jsonObject := map[string][]Product{"products": products}
    return c.JSON(http.StatusOK, jsonObject)
}

func showProduct(c echo.Context) error {
    var product Product
    db.First(&product, c.Param("id"))
    return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
    product := Product{Code: "ABCDEF", Price: 1000}
    db.Create(&product)
    return c.JSON(http.StatusOK, product)
}