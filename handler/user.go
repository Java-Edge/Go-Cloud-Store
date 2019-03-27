package handler

import (
	dblayer "../../filestore-server/db"
	"../../filestore-server/util"
	"io/ioutil"
	"net/http"
)

const (
	pwdSalt = "*#520"
)

// 处理用户的注册请求
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	// 参数校验
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if len(username) < 3 || len(password) < 5 {
		w.Write([]byte("invalid parameter"))
		return
	}

	encPassword := util.Sha1([]byte(password + pwdSalt))
	suc := dblayer.UserSignup(username, encPassword)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}
