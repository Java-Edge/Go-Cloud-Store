package db

import (
	mydb "../../filestore-server/db/mysql"
	"fmt"
)

// 文件上传完成,保存meta信息
func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file(`file_sha1`,`file_name`,`file_size`)" +
			"`file_addr`,`status` values (?,?,?,?,1)")
	if err != nil {
		fmt.Println("Failed to prepare statement,err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); nil == err {
		// 虽然SQL执行成功,但未产生新的表记录!
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been uploaded before", filehash)
		}
		return true
	}
	return false
}
