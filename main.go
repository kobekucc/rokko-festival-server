package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

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

	e.Use(middleware.CORS())

	// middleware
	e.Use(middleware.Logger())

	// routing
	e.GET("/vote/:type/:id", checkExistVote) //投票があったら（レコードがあったら）１を返す（レコードが存在）

	e.POST("/comment/:type/:id/:comment", createComment)
	e.POST("/vote/:type/:id", createVote)
	e.POST("/questionnaire/", createQuestionnaire)
	e.POST("/impression/:type/:comment", createImpression)

	e.PUT("/vote/:type/:id", incrementVote)

	e.Logger.Fatal(e.Start(":" + port))
}

func createComment(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	comment := Comment{Type: c.Param("type"), WorkId: uint(id), Comment: c.Param("comment")}
	db.Create(&comment)
	return c.JSON(http.StatusOK, comment)
}
func createVote(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	work := Work{Type: c.Param("type"), WorkId: uint(id), Vote: 1}
	db.Create(&work)
	return c.JSON(http.StatusOK, work)
}
func checkExistVote(c echo.Context) error {
	var count int
	var works []Work
	db.Where("type = ? AND work_id = ?", c.Param("type"), c.Param("id")).First(&works).Count(&count)
	return c.JSON(http.StatusOK, count)
}
func incrementVote(c echo.Context) error {
	var work Work
	db.Where("type = ? AND work_id = ?", c.Param("type"), c.Param("id")).First(&work)

	work.Vote++
	db.Save(&work)
	return c.JSON(http.StatusOK, work)
}
func createQuestionnaire(c echo.Context) error {
	param := new(Questionnaire)
    if err := c.Bind(param); err != nil {
        return err
    }
	questionnaire := Questionnaire{Age: param.Age,
		Gender: param.Gender,
		Rate:  param.Rate,
		Opinion: param.Opinion}
	db.Create(&questionnaire)
	return c.JSON(http.StatusOK, param)
}
func createImpression(c echo.Context) error {
	impression := Impression{Type: c.Param("type"), Comment: c.Param("comment")}
	db.Create(&impression)
	return c.JSON(http.StatusOK, impression)
}
