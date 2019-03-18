package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传 HTML 界面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接收文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data,err:%s\n", err.Error())
		}
		// 关闭文件句柄
		defer file.Close()

		// 创建本地文件接收文件流
		newFile, err := os.Create("/Volumes/doc/tmp/" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create file,err:%s\n", err.Error())
		}
		defer newFile.Close()

		// 将文件内容拷贝到新文件的buffer流中
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file,err:%s\n", err.Error())
			return
		}

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// 上传成功提示
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload succeed!")
}
