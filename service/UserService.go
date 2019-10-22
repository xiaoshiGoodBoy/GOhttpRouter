package service

import(
	"GoDemo/webGoRouterDemo/GOhttpRouter/dao"
	"GoDemo/webGoRouterDemo/GOhttpRouter/pojo"
	"database/sql"
)
func InsertUser(db *sql.DB, person pojo.Person) int {
	return dao.InsertUser(db, person)
}


func DeleteUser(db *sql.DB, id string) string {
	return dao.DeleteUser(db, id)
}

func UpdateUser(db *sql.DB, person pojo.Person) string {
	return dao.UpdateUser(db, person)
}

func SelectUser(db *sql.DB, id string) pojo.Person {
	return dao.SelectUserById(db, id)
}

func SelectAllUser(db *sql.DB) map[int] pojo.Person {
	return dao.SelectAllUser(db)
}