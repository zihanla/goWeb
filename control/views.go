package control

import (
	"io"
	"net/http"
	"os"
)

// IndexView 主页面
func IndexView(w http.ResponseWriter, r *http.Request) {
	// 把 index的内容写进页面
	f, _ := os.Open("./views/index.html")
	io.Copy(w, f)
}

// DetailView 详细页
func DetailView(w http.ResponseWriter, r *http.Request) {
	// 把 index的内容写进页面
	f, _ := os.Open("./views/detail.html")
	io.Copy(w, f)
}

// ListView 列表页面
func ListView(w http.ResponseWriter, r *http.Request) {
	// 把 index的内容写进页面
	f, _ := os.Open("./views/list.html")
	io.Copy(w, f)
}

// ViewAdd 添加页面
func ViewAdd(w http.ResponseWriter, r *http.Request) {
	// 把 index的内容写进页面
	f, _ := os.Open("./views/add.html")
	io.Copy(w, f)
}

// ViewEdit 修改页面
func ViewEdit(w http.ResponseWriter, r *http.Request) {
	// 把 index的内容写进页面
	f, _ := os.Open("./views/edit.html")
	io.Copy(w, f)
}
