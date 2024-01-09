package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/", Hello())
    e.GET("/api/", ApiHelloGet())

    e.Start(":8080")
}

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.String(http.StatusOK, "4415: Python3 is the best language!")
    }
}

func ApiHelloGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]interface{}{"studentId": "4415", "message": "Python is the best language! If you using other language, it should switch to python right now."})
    }
}
