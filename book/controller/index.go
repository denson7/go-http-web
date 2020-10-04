package controller

import (
	"book/dao"
	"book/model"
	"bytes"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")

	if pageNo == "" {
		pageNo = "1"
	}
	// 解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	// 获取数据
	var data *model.Page
	data, _ = dao.GetPageBooks(pageNo)
	t.Execute(w, data)
}

func Index2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// 设置Header 头
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Write([]byte("This is an example server.\n"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func Error(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL)
	http.Error(w, "oops", 404)
}

func SendGet(w http.ResponseWriter, r *http.Request) {
	url := "https://cnodejs.org/api/v1/topics"
	response, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//return ""
	}

	fmt.Println(string(body))
	//w.Write([]byte(body))
}

func SendPost(w http.ResponseWriter, r *http.Request) {
	body := "{\"action\":20}"
	url := "https://cnodejs.org/api/v1"
	response, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() == "/favicon.ico" {
		return
	}
	if r.Method == "GET" {
		http.Redirect(w, r, "/test", 302)
	}
}
