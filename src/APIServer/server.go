package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
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
	e.POST("/user/signup", signUp)
	e.Logger.Fatal(e.Start(":9000"))

}

func appVersion(c echo.Context) error {

	r := &Response{
		ResultCode: 102,
		ResultDesc: "system error : ",
		ResultData: nil,
	}

	os := c.QueryParam("os")
	if os != "ios" && os != "aos" {
		r.ResultCode = 101
		r.ResultDesc = "parameter error"
		return c.JSON(http.StatusOK, r)
	}

	db, err := sql.Open("mysql", "chaeun:ehlswkd1@tcp(127.0.0.1:3306)/qfeat")
	if err != nil || db.Ping() != nil {
		r.ResultDesc += "db connection error"
		return c.JSON(http.StatusOK, r)
	}
	defer db.Close()

	var version string
	var url string
	fmt.Println(os)
	err = db.QueryRow("SELECT VERSION, URL FROM version WHERE OS = ?", os).Scan(&version, &url)

	if err != nil {
		r.ResultDesc += "query error"
		return c.JSON(http.StatusOK, r)
	}

	var app = make(map[string]interface{})
	app["version"] = version
	app["update"] = true
	app["url"] = url
	r.ResultData = make(map[string]map[string]interface{})
	r.ResultData["app"] = app

	r.ResultCode = 0
	r.ResultDesc = "success"

	return c.JSON(http.StatusOK, r)
}

func signUp(c echo.Context) error {
	r := &Response{
		ResultCode: 102,
		ResultDesc: "system error : ",
		ResultData: nil,
	}

	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		r.ResultDesc += err.Error()
		return c.JSON(http.StatusOK, r)
	}

	db, err := sql.Open("mysql", "chaeun:ehlswkd1@tcp(127.0.0.1:3306)/qfeat")
	if err != nil || db.Ping() != nil {
		r.ResultDesc += "db connection error"
		return c.JSON(http.StatusOK, r)
	}
	defer db.Close()

	r.ResultCode = 0
	r.ResultDesc = "success"
	r.ResultData = make(map[string]map[string]interface{})
	param := m["param"].(map[string]interface{})
	device := param["device"].(map[string]interface{})
	user := param["user"].(map[string]interface{})

	res, err := db.Exec("INSERT INTO user(`nickname`) VALUES (?)", user["nickname"])
	if err != nil {
		r.ResultDesc += err.Error()
		return c.JSON(http.StatusOK, r)
	}
	id, _ := res.LastInsertId()

	dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, 16)
	rand.Read(bytes)

	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	token := string(bytes)

	res, err = db.Exec("INSERT INTO device VALUES (?,?,?,?,?,?,?,?,?,?)", id, device["uuid"], device["os"], device["version"], device["model"], device["mcc"], device["mnc"], device["ctn"], device["pushKey"], token)
	if err != nil {
		r.ResultDesc += err.Error()
		return c.JSON(http.StatusOK, r)
	}

	common := make(map[string]interface{})
	common["token"] = token

	userRes := make(map[string]interface{})
	userRes["recommCode"] = nil
	userRes["code"] = id
	userRes["nickname"] = user["nickname"]
	userRes["imgUrl"] = "http://"
	userRes["acceptTerms"] = user["acceptTerms"]
	userRes["acceptPrivacy"] = user["acceptPrivacy"]
	userRes["hearts"] = 3.0
	userRes["coupon"] = 3
	userRes["accountBalance"] = 0
	userRes["usePush"] = true
	userRes["useVibration"] = true
	userRes["useDataSave"] = true

	r.ResultData["common"] = common
	r.ResultData["user"] = userRes
	return c.JSON(http.StatusOK, r)
}
