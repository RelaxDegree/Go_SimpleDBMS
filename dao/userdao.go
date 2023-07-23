package dao

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id     int
	Name   string
	Passwd string
}

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go")
	if err != nil {
		fmt.Println(err)
	}
}
func UserLogin(username string, pswd string) bool {
	stmt, err := DB.Prepare("SELECT name FROM user WHERE name = ? and pswd = ?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	rows, err := stmt.Query(username, pswd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}

func UserRegister(username string, pswd string) bool {
	stmt, err := DB.Prepare("INSERT INTO user(name, pswd) VALUES(?, ?)")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	res, err := stmt.Exec(username, pswd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(id)
	return true
}

func UserChangePswd(username string, oldpswd string, newpswd string) bool {
	stmt, err := DB.Prepare("UPDATE user SET pswd = ? WHERE name = ? and pswd = ?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	res, err := stmt.Exec(newpswd, username, oldpswd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	id, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(id)
	return true
}
func Record(username string, msg string, ip string, ret string) bool {
	stmt, err := DB.Prepare("INSERT INTO log(name, msg, ip, ret, datetime) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	currentTime := time.Now()
	time := currentTime.Format("2006-01-02 15:04:05")
	stmt.Exec(username, msg, ip, ret, time)

	return true
}
