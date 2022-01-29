package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if len(c.Request().Header["Authorization"]) > 0 {
			if c.Request().Header["Authorization"][0] == "aswad" {
				c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
				return next(c)
			}
		}

		return c.JSON(http.StatusForbidden, "You are not authorized!")
	}
}

type DBConnect struct {
	DB *sql.DB
}

func DbConnect() *DBConnect {
	DB, err := sql.Open("mysql", "database:Aswad_database@123@tcp(127.0.0.1:3306)/calculatorDB")
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("connected")
	}

	return &DBConnect{
		DB: DB,
	}

}

func main() {
	// Echo instan(["AUTHorization"])
	e := echo.New()
	// DbConnect()
	db := DbConnect()

	// // Middleware

	// e.Use(middleware.Recover())
	e.Use(ServerHeader)
	// Routes
	//e.POST("/", hello)
	e.POST("/sub", db.SUB)
	e.POST("/add", db.ADD)
	e.POST("/mul", db.MUL)
	e.POST("/mod", db.MOD)
	e.POST("/pow", db.POW)
	e.POST("/div", db.DIV)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	fmt.Println("testing")

}

// Handler
/*func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}*/

type Input struct {
	Num1 int `json:"number1"`
	Num2 int `json:"number2"`
}

type Response struct {
	Result int `json:"result"`
}

func (db *DBConnect) ADD(c echo.Context) error {
	n := new(Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	add := n.Num1 + n.Num2
	result := Response{
		add,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(n.Num1, n.Num2, add, "+")

	if err2 != nil {
		panic(err2)
	}

	return c.JSON(http.StatusOK, result)
}

func (db *DBConnect) MUL(c echo.Context) error {
	n := new(Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	if n.Num1 == 0 || n.Num2 == 0 {
		c.JSON(http.StatusForbidden, "ANY INPUT IS ZERO THE RESULT IS ZERO")
	}
	mul := n.Num1 * n.Num2
	result := Response{
		mul,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(n.Num1, n.Num2, mul, "*")

	if err2 != nil {
		panic(err2)
	}

	return c.JSON(http.StatusOK, result)
}

func (db *DBConnect) SUB(c echo.Context) error {
	n := new(Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	sub := n.Num1 - n.Num2
	result := Response{
		sub,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(n.Num1, n.Num2, sub, "-")

	if err2 != nil {
		panic(err2)
	}

	return c.JSON(http.StatusOK, result)
}
func (db *DBConnect) DIV(c echo.Context) error {
	n := new(Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	if n.Num2 == 0 {
		c.JSON(http.StatusForbidden, "INVALID INPUT OF NUMBER 2:")
	}
	div := n.Num1 / n.Num2
	result := Response{
		div,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(n.Num1, n.Num2, div, "/")

	if err2 != nil {
		panic(err2)
	}

	return c.JSON(http.StatusOK, result)
}

func (db *DBConnect) MOD(c echo.Context) error {
	n := new(Input)
	if err := c.Bind(n); err != nil {
		return err
	}
	MOD := n.Num1 % n.Num2
	result := Response{
		MOD,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(n.Num1, n.Num2, MOD, "%")

	if err2 != nil {
		panic(err2)
	}

	return c.JSON(http.StatusOK, result)
}

func (db *DBConnect) POW(c echo.Context) error {
	n := new(Input)
	if err := c.Bind(n); err != nil {
		return err
	}
	if n.Num1 == 0 {

		c.JSON(http.StatusForbidden, "IF BASE IS ZERO THEN ANY VALUE OF POWER RESULT IS ALWAYS 0")

	}
	if n.Num1 == 1 {

		c.JSON(http.StatusForbidden, "IF BASE IS 1 THEN ANY VALUE OF POWER RESULT IS ALWAYS 1")
	}

	if n.Num2 == 0 {
		c.JSON(http.StatusForbidden, "IF POWER IS ZERO THEN ANY VALUE OF BASE RESULT IS ALWAYS 1")
	}

	var pow int
	pow = 1
	for n.Num2 != 0 {
		pow *= n.Num1
		n.Num2 -= 1
	}

	result := Response{
		pow,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(n.Num1, n.Num2, pow, "n^m")

	if err2 != nil {
		panic(err2)
	}

	return c.JSON(http.StatusOK, result)
}

// func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
// 		return next(c)
// 	}
// }
