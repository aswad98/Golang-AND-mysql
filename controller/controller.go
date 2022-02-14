package controller

import (
	"fmt"
	"hello/models"
	"net/http"

	"github.com/labstack/echo"
)

func (db *DBConnect) ADD(c echo.Context) error {
	// conn := database.DbConnect()
	n := new(models.Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	add := n.Num1 + n.Num2
	result := models.Response{
		add,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.conn.Prepare(sql)
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
	n := new(models.Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	if n.Num1 == 0 || n.Num2 == 0 {
		c.JSON(http.StatusForbidden, "ANY INPUT IS ZERO THE RESULT IS ZERO")
	}
	mul := n.Num1 * n.Num2
	result := models.Response{
		mul,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.conn.Prepare(sql)
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
	n := new(models.Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	sub := n.Num1 - n.Num2
	result := models.Response{
		sub,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.conn.Prepare(sql)
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
	n := new(models.Input)

	if err := c.Bind(n); err != nil {
		return err
	}
	if n.Num2 == 0 {
		c.JSON(http.StatusForbidden, "INVALID INPUT OF NUMBER 2:")
	}
	div := n.Num1 / n.Num2
	result := models.Response{
		div,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.conn.Prepare(sql)
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
	n := new(models.Input)
	if err := c.Bind(n); err != nil {
		return err
	}
	MOD := n.Num1 % n.Num2
	result := models.Response{
		MOD,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.conn.Prepare(sql)
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
	n := new(models.Input)
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

	result := models.Response{
		pow,
	}

	sql := "INSERT INTO Calculator(NUMBER1, NUMBER2,RESULT,OPERATION) VALUES( ?, ?, ?,?)"
	stmt, err := db.conn.Prepare(sql)
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
