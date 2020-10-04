package controller

import (
	"book/model"
	"book/utils"
	"encoding/json"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	u := model.User{
		Username: "Smith",
		Email:    "test@126.com",
	}

	//msg, _ := json.Marshal(u)
	msg, _ := json.Marshal(utils.Response{
		Code: 200,
		Msg:  "Success",
		Data: u,
	})
	// Write content-type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(msg)
}

func Test2(w http.ResponseWriter, r *http.Request) {
	u := model.User{
		Username: "Smith",
		Email:    "test@126.com",
	}

	//msg, _ := json.Marshal(u)
	data := utils.Response{
		Code: 200,
		Msg:  "Success",
		Data: u,
	}
	// 封装JSON返回
	handler := utils.NewHttpResponseWriter()
	err1 := handler.RenderJSON(r, w, 200, data)
	if err1 != nil {
		//
	}
}
