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
	e.POST("/questionnaire/:age/:gender/:rate/:opinion", createQuestionnaire)
	e.POST("/impression/:type/:comment", createImpression)

	e.POST("/book/:order/:num/:name/:switchname", createOnetoone)
	e.GET("/book/all", getAllOnetoone)
	e.GET("/book/count", getOnetooneNum)
	e.PUT("/book/:name/:order", putOrder)

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
	rate, _ := strconv.ParseUint(c.Param("rate"), 10, 64)
	questionnaire:= Questionnaire{Age: c.Param("age"),Gender: c.Param("gender"),Rate: uint(rate), Opinion: c.Param("opinion")}
	db.Create(&questionnaire)
	return c.JSON(http.StatusOK, questionnaire)
}
func createImpression(c echo.Context) error {
	impression := Impression{Type: c.Param("type"),Comment: c.Param("comment")}
	db.Create(&impression)
	return c.JSON(http.StatusOK, impression)
}
func createOnetoone(c echo.Context) error {
	order, _ := strconv.ParseUint(c.Param("order"), 10, 64)
	num, _ := strconv.ParseUint(c.Param("num"), 10, 64)
	onetoone := Onetoone{Order:uint(order),NumberOfTimes: uint(num),Name:c.Param("name"),SwitchName: c.Param("switchname")}
	db.Create(&onetoone)
	return c.JSON(http.StatusOK, onetoone)
}
func getAllOnetoone(c echo.Context) error {
	var all []Onetoone

	db.Find(&all)
	return c.JSON(http.StatusOK, all)
}
func getOnetooneNum(c echo.Context) error {
	var count int
	db.Table("onetoones").Count(&count)
	return c.JSON(http.StatusOK, count)
}
func putOrder(c echo.Context) error {
	var onetoone Onetoone
	db.Where("name = ? ",  c.Param("name")).First(&onetoone)

	order, _ := strconv.ParseUint(c.Param("order"), 10, 64)

	onetoone.Order=uint(order)
	db.Save(&onetoone)
	return c.JSON(http.StatusOK, onetoone)
}
