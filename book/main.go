package main

import (
	"book/controller"
	"log"
	"net/http"
)

func main() {
	// 设置路由
	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/hello", controller.Hello)
	http.HandleFunc("/404", controller.Error)
	http.HandleFunc("/sendGet", controller.SendGet)
	http.HandleFunc("/sendPost", controller.SendPost)

	// user
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	// book
	http.HandleFunc("/getBookList", controller.GetBookList)
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)
	http.HandleFunc("/deleteBook", controller.DeleteBook)

	// 重定向
	http.HandleFunc("/redirectToTest", controller.Redirect)

	// test
	http.HandleFunc("/test", controller.Test)
	http.HandleFunc("/test2", controller.Test2)
	// 监听http
	err := http.ListenAndServe(":8080", nil)
	// 处理https
	//err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
