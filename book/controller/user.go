package controller

import (
	"book/dao"
	"book/utils"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// 获取POST参数
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := dao.CheckUserNameAndPassword(username, password)

	if user.ID > 0 { // 用户名密码正确
		uuid := utils.CreateUUID() // 生成UUID
		cookie := http.Cookie{ // 创建一个cookie
			Name:     "user",
			Value:    uuid,
			HttpOnly: true,
		}
		// 将cookie发送给浏览器
		http.SetCookie(w, &cookie)
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, user)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确")
	}
}


// Logout 注销
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookie.MaxAge = -1        // 设置cookie失效
		http.SetCookie(w, cookie) // 将修改之后的cookie发送给浏览器
	}
	Index(w, r) // 去首页
}

// CheckUserName 通过发送Ajax请求验证用户名是否存在
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		w.Write([]byte("用户名已存在！"))
	} else {
		w.Write([]byte("<font style = 'color:green'>用户名不存在！</font>"))
	}
}

//
func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	err := dao.AddUser(username, password, email)
	if err != nil {
		w.Write([]byte("注册失败！"))
	}
	w.Write([]byte("注册成功！"))
}