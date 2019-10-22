package main

import (
	"GoDemo/webGoRouterDemo/GOhttpRouter/controller"
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host string = "127.0.0.1"
	port string = "3306"
	username string = "root"
	password string = ""
	database string = "test"
)

func main() {
	initDatabase()
	log.Println("程序启动！")
	router := httprouter.New()
	router.POST("/addUser", controller.AddUser)
	router.POST("/updateUser", controller.UpdateUser)
	router.POST("/deleteUser", controller.DeleteUser)
	router.POST("/findUserById", controller.FindUserById)
	router.POST("/findAllUser", controller.FindAllUser)
	http.ListenAndServe(":8080", router)
}


func initDatabase(){
	dbDriver := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",username, password, host, port, database)
	sql,err := sql.Open("mysql", dbDriver)
	if err != nil{
		fmt.Println("数据库初始化失败: " ,err)
	}
	controller.Sqldb = sql
	log.Println("database connect success ......")
}