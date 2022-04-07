package router

import (
	"blog/api"
	"blog/views"
	"net/http"
)

func Router(){
	//1. 返回页面
	//2. 返回数据
	//3. 返回静态资源

	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api",api.API.SaveAndUpdatePost)
	http.Handle("/resource/",http.StripPrefix("/resource/",http.FileServer(http.Dir("public/resource/"))))
}
