package db

import (
	mydb "../../filestore-server/db/mysql"
	"fmt"
)

// 通过用户名及密码完成user表的注册
func UserSignup(username string, password string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`,`user_pwd`) values (?,?)")
	if err != nil {
		fmt.Println("Failed to insert,err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Println("Failed to insert,err:" + err.Error())
		return false
	}

	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}

	return false
}
