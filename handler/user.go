package handler

import (
	dblayer "../../filestore-server/db"
	"../../filestore-server/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

// 用户登录
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	encPasswd := util.Sha1([]byte(password + pwdSalt))

	// 1. 校验用户名及密码
	pwdChecked := dblayer.UserSignin(username, encPasswd)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}

	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		w.Write([]byte("FAILED"))
		return
	}

	// 3. 登录成功后重定向到首页
	w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
}

// 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
