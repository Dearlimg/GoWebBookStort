package controller

import (
	"bookstore01/dao"
	"bookstore01/model"
	"bookstore01/utils"
	"fmt"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	flag, _ := dao.IsLogin(r)
	if flag {
		GetPageBooksByPrice(w, r)
	} else {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		fmt.Println(username, password)
		user := dao.CheckUserNameAndPassword(username, password)
		if user.ID > 0 {
			uuid := utils.CreatUUID()
			sess := &model.Session{
				SessionID: uuid,
				UserName:  username,
				UserID:    user.ID,
			}

			dao.AddSession(sess)
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			fmt.Println("登录成功")
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			_ = t.Execute(w, user)
		} else {
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "")
		}
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ := dao.CheckUserName(username)
	fmt.Println(user.ID, password, email)
	if user.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		_ = t.Execute(w, "用户名已存在")
	} else {
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "成功")
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	GetPageBooksByPrice(w, r)
}
