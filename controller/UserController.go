package controller

import (
	"GoDemo/webGoRouterDemo/GOhttpRouter/pojo"
	"GoDemo/webGoRouterDemo/GOhttpRouter/service"
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var (
	Sqldb *sql.DB
)



func AddUser(w http.ResponseWriter, r *http.Request, params httprouter.Params)  {
	r.ParseForm()
	per := pojo.Person{ "" ,r.FormValue("loginName"), r.FormValue("userEmail"), r.FormValue("userPassword")}
	number := service.InsertUser(Sqldb, per)
	w.Write([]byte(string(number)))
}


func UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	r.ParseForm()
	person := pojo.Person{r.FormValue("id"), r.FormValue("loginName"), r.FormValue("userEmail"), r.FormValue("userPassword",)}
	number := service.UpdateUser(Sqldb, person)
	w.Write([]byte(number))
}

func DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	r.ParseForm()
	number := service.DeleteUser(Sqldb, r.FormValue("id"))
	w.Write([]byte(number))
}

func FindUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	r.ParseForm()
	person := service.SelectUser(Sqldb, r.FormValue("id"))
	res,_ := json.Marshal(person)

	w.Write(res)
}

func FindAllUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userMap := service.SelectAllUser(Sqldb)
	JuserMap,_ := json.Marshal(userMap)
	w.Write(JuserMap)
}
