package main

import (
	"easy_demo/control"
	"net/http"
)

func main() {
	// 处理方法
	http.HandleFunc("/", control.IndexView)        // 主页面
	http.HandleFunc("/list", control.ListView)     // list页面
	http.HandleFunc("/add", control.ViewAdd)       // add 页面
	http.HandleFunc("/edit", control.ViewEdit)     // 修改页面
	http.HandleFunc("/detail", control.DetailView) // 详细页面
	// 用于区分 页面 和 方法
	http.HandleFunc("/api/index/data", control.IndexData)     // 主页面 查询
	http.HandleFunc("/api/list/data", control.ListData)       // list页面 列表
	http.HandleFunc("/api/list/del", control.ListDel)         // list 删除
	http.HandleFunc("/api/article/add", control.ArticleAdd)   // add 添加功能
	http.HandleFunc("/api/uploda", control.ApiUpload)         // add 添加功能
	http.HandleFunc("/api/article/edit", control.ArticleEdit) // list 修改功能
	http.HandleFunc("/api/article/page", control.ArticlePage) // list 分页功能

	// 指定静态文件访问路径
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	// 监听启动服务
	http.ListenAndServe(":8080", nil)
}
