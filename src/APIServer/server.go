package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Response struct {
	ResultCode int                               `json:"resultCode"`
	ResultDesc string                            `json:"resultDesc"`
	ResultData map[string]map[string]interface{} `json:"resultData"`
}

func main() {
	e := echo.New()
	e.GET("/app/version", appVersion)

	e.Logger.Fatal(e.Start(":9000"))

}

func appVersion(c echo.Context) error {
	os := c.QueryParam("os")
	r := &Response{
		ResultCode: 0,
		ResultDesc: "success",
		ResultData: nil,
	}
	db, err := sql.Open("mysql", "chaeun:ehlswkd1@tcp(127.0.0.1:3306)/qfeat")
	if err != nil || db.Ping() != nil {
		log.Panic(err)
	}
	defer db.Close()

	var version string
	var url string
	fmt.Println(os)
	err = db.QueryRow("SELECT VERSION, URL FROM version WHERE OS = ?", os).Scan(&version, &url)

	if err != nil {
		log.Panic(err)
	}

	var app = make(map[string]interface{})
	app["version"] = version
	app["update"] = true
	app["url"] = url
	r.ResultData = make(map[string]map[string]interface{})
	r.ResultData["app"] = app

	return c.JSON(http.StatusOK, r)
}
