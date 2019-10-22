package dao

import (
	"database/sql"
	"fmt"
	"GoDemo/webGoRouterDemo/GOhttpRouter/pojo"
	"log"
)

/**
   注册用户
 */
func InsertUser(db *sql.DB, student pojo.Person) int {
	sqlStr := "insert into tab_user (loginName, userEmail, userPassword) values (?,?,?)"
	fmt.Println("----------------------1---", student.LoginName)
	stmt,err := db.Prepare(sqlStr)
	fmt.Println("----------------------2---")
	if err != nil{
		fmt.Println("异常错误信息：", err)
	}
	_,e := stmt.Exec(student.LoginName, student.UserEmail, student.UserPassword)

	if e != nil{
		return 1
	}

	log.Println("插入成功！")
	return 0

}

/**
	删除用户
 */
func DeleteUser(db *sql.DB, id string) string {
	sqlStr := "delete from tab_user where id = ?"
	stmt,err := db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("异常错误信息：", err)
	}
	_,e := stmt.Exec(id)
	if e != nil{
		return "1"
	}
	log.Println("删除成功！")
	return "0"
}

/**
	修改用户
 */
func UpdateUser(db *sql.DB, person pojo.Person) string {
	sqlStr := "update tab_user set loginName = ?, userEmail= ?, userPassword=? where id = ?"
	stmt,err := db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("异常错误信息：", err)
	}
	_,e := stmt.Exec(person.LoginName, person.UserEmail, person.UserPassword, person.Id)
	if e != nil{
		return "1"
	}
	log.Println("修改成功！")
	return "0"
}

/**
	查看单个用户
 */
func SelectUserById(db *sql.DB, idNum string) pojo.Person{
	sqlStr := "select * from tab_user where id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("异常错误信息：", err)
	}
	row,_ := stmt.Query(idNum)
	var userId string
	var loginName string
	var userEmail string
	var userPassword string
	if row.Next(){
		row.Scan(&userId, &loginName, &userEmail, &userPassword)
	}
	person := pojo.Person{userId, loginName, userEmail, userPassword}
	log.Println("查询成功！")
	return person
}

/**
	查询全部
 */
func SelectAllUser(db *sql.DB) map[int]pojo.Person {
	sqlStr := "select * from tab_user"
	stmt,err := db.Prepare(sqlStr)
	if err != nil{
		fmt.Println("异常错误信息：", err)
	}
	row,_ := stmt.Query()
	var num int = 0
	userMap := make(map[int] pojo.Person)
	for row.Next(){
		num++
		fmt.Println("当前结果个数：", num)
		var userId string
		var loginName string
		var userEmail string
		var userPassword string
		row.Scan(&userId, &loginName, &userEmail, &userPassword)
		person := pojo.Person{userId, loginName, userEmail, userPassword}
		userMap[num] = person
	}
	log.Println("查询全部成功！")
	return userMap
}