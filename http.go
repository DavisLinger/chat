package main

import (
	"imoniang.com/chat/lib"
	"imoniang.com/chat/sql"
	"net/http"
)

type Register struct {
	User string `json:"user"` // 注册账号
	Pass string `json:"pass"` // 注册密码
	Nick string `json:"nick"` // 用户昵称
}

type UserInfo struct {
	Token string `json:"token"` // 用户登录成功后的凭证
	Nick  string `json:"nick"`  // 用户昵称
}

// 首页界面
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

// 聊天界面
func chat(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/chat.html")
}

// 登录界面及登录函数
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // 这里用来处理传递上来的登录信息
		postUser := r.PostFormValue("user")
		postPass := r.PostFormValue("pass")

		w.Header().Add("content-type", "application/json;charset=utf-8")

		if lib.IsEmpty(postUser, postPass) {
			w.Write(lib.MakeReturnJson(501, "账号或者密码不可为空", nil))
			return
		}

		if !lib.IsAlphaNum(postUser, postPass) {
			w.Write(lib.MakeReturnJson(501, "账号密码只能由字母和数字组成", nil))
			return
		}

		if !lib.Len(6, 20, postUser, postPass) {
			w.Write(lib.MakeReturnJson(501, "账号密码长度为6~20位", nil))
			return
		}

		user, result := sql.CheckUserLogin(postUser, postPass)
		if !result {
			w.Write(lib.MakeReturnJson(502, "账号或者密码错误", nil))
			return
		}
		token, err := sql.MakeToken(&user[0])
		if err != nil {
			w.Write(lib.MakeReturnJson(503, "登录失败", nil))
			return
		}

		userInfo := &UserInfo{
			Token: token,
			Nick:  user[0].Nick,
		}
		w.Write(lib.MakeReturnJson(200, "登录成功", userInfo))
		return
	} else {
		http.ServeFile(w, r, "html/login.html")
	}
}

// 注册界面及注册函数
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // 这里用来处理传递上来的注册信息
		var registerInfo = Register{
			User: r.PostFormValue("user"),
			Pass: r.PostFormValue("pass"),
			Nick: r.PostFormValue("nick"),
		}
		w.Header().Add("content-type", "application/json;charset=utf-8")
		if lib.IsEmpty(registerInfo.Nick, registerInfo.Pass, registerInfo.User) {
			w.Write(lib.MakeReturnJson(501, "需要填写全部参数", nil))
			return
		}
		if !lib.IsChsAlphaNum(registerInfo.Nick) {
			w.Write(lib.MakeReturnJson(501, "昵称只能由汉字、字母和数字组成", nil))
			return
		}

		if !lib.Len(2, 8, registerInfo.Nick) {
			w.Write(lib.MakeReturnJson(501, "昵称长度为2~8位", nil))
			return
		}

		if !lib.Len(6, 20, registerInfo.Pass, registerInfo.User) {
			w.Write(lib.MakeReturnJson(501, "账号以及密码长度为6~20位", nil))
			return
		}

		if !lib.IsAlphaNum(registerInfo.User, registerInfo.Pass) {
			w.Write(lib.MakeReturnJson(501, "账号密码只能由字母和数字组成", nil))
			return
		}

		user, _ := sql.GetUser(&sql.User{User: registerInfo.User})
		if len(user) != 0 {
			w.Write(lib.MakeReturnJson(502, "账号已存在", registerInfo.User))
			return
		}

		user, _ = sql.GetUser(&sql.User{Nick: registerInfo.Nick})
		if len(user) != 0 {
			w.Write(lib.MakeReturnJson(502, "昵称不可重复", registerInfo.Nick))
			return
		}
		err := sql.AddUser(registerInfo.User, registerInfo.Pass, registerInfo.Nick)
		if err != nil {
			w.Write(lib.MakeReturnJson(503, "注册失败", nil))
			return
		}
		w.Write(lib.MakeReturnJson(200, "注册成功", nil))
	} else {
		http.ServeFile(w, r, "html/register.html")
	}
}
