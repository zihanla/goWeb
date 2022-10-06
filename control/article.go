package control

import (
	"easy_demo/model"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// IndexData 主页面数据处理
func IndexData(w http.ResponseWriter, r *http.Request) {
	// 解析请求中的信息到form中
	r.ParseForm()
	// 设置返回的响应内容为json格式
	w.Header().Set("Content-Type", "application/javascript")
	idStr := r.Form.Get("id") // 拿取指定id数据 数据为string类型
	id, _ := strconv.ParseInt(idStr, 10, 64)
	// 使用 ArticleGet 方法查询数据
	mod, err := model.ArticleGet(id)
	if err != nil {
		Fail(w, err.Error())
		return
	}
	Succ(w, "查询成功", mod)
	return
	// 把返回的内容序列化 为json格式，写入响应中
	//buf, _ := json.Marshal(mod)
	//w.Write(buf)

}

// ListData 列表页面数据查询
func ListData(w http.ResponseWriter, r *http.Request) {
	mods, err := model.ArticleList()
	if err != nil {
		Fail(w, err.Error())
		return
	}
	Succ(w, "200", mods)
	return
}

// ArticlePage 分页查询
func ArticlePage(w http.ResponseWriter, r *http.Request) {
	// 解析请求中的内容
	r.ParseForm()
	// 拿到 当前页、每页条数 数据 传输过来的数据为字符串格式 数据库需要的是int
	pi, _ := strconv.Atoi(r.Form.Get("pi"))
	ps, _ := strconv.Atoi(r.Form.Get("ps"))
	// 获取数据总条数
	count := model.ArticlePageCount()
	// 判断总条数
	if count < 1 {
		Fail(w, "未查询到数据", "count == 0")
		return
	}
	// 把 当前页、每页条数 转入 分页查询方法中
	mods, _ := model.ArticlePage(pi, ps)
	// 判断返回的数据条数
	if len(mods) < 1 {
		Fail(w, "未查询到数据", "len(mods) < 1")
		return
	}
	ext := Ext{
		Count: count,
		Items: mods,
	}
	Succ(w, "分页数据", ext)
}

func ListDel(w http.ResponseWriter, r *http.Request) {
	// 解析请求中的信息到form中
	r.ParseForm()
	idStr := r.Form.Get("id") // 拿取指定id数据 数据为string类型
	id, _ := strconv.ParseInt(idStr, 10, 64)
	// 使用 ArticleGet 方法查询数据
	if model.ArticleDel(id) {
		Succ(w, "删除成功")
		return
	}
	Fail(w, "删除失败")
	return
}

// ArticleAdd 添加数据
//func ArticleAdd(w http.ResponseWriter, r *http.Request) {
//	// 解析请求中的数据
//	r.ParseForm()
//	mod := &model.Article{}
//	mod.Title = r.Form.Get("title")
//	mod.Author = r.Form["author"][0]
//	mod.Content = r.FormValue("content")
//	mod.Hits, _ = strconv.Atoi(r.Form.Get("hits")) // r.Form.Get("hits") 返回的是字符串, hits是int类型的数据
//	mod.Utime = time.Now()
//	// 把处理号的数据放入 ArticleAdd 方法中
//	err := model.ArticleAdd(mod)
//	if err == nil {
//		Succ(w, "添加成功")
//		return
//	}
//	Fail(w, "添加失败", err.Error())
//	return
//}

func ArticleAdd(w http.ResponseWriter, r *http.Request) {
	mod := &model.Article{}
	// 读取请求体中所有的内容
	//buf, _ := ioutil.ReadAll(r.Body)
	// 反序列化，把 json格式的buf 反序列化 装入mod中
	//err := json.Unmarshal(buf, mod)
	// 上面两行的实现 等于 下面这一行
	err := json.NewDecoder(r.Body).Decode(mod)
	if err != nil {
		Fail(w, "输入数据有误", err.Error())
		return
	}
	mod.Utime = time.Now()
	// 把mod结构体 放入ArticleAdd() 添加数据
	err = model.ArticleAdd(mod)
	if err != nil {
		Fail(w, "添加失败", err.Error())
		return
	}
	Succ(w, "添加成功")
	return
}

// ArticleEdit 文章修改
func ArticleEdit(w http.ResponseWriter, r *http.Request) {
	mod := &model.Article{}
	err := json.NewDecoder(r.Body).Decode(mod)
	if err != nil {
		Fail(w, "输入数据有误", err.Error())
		return
	}
	// 把mod结构体 放入ArticleAdd() 添加数据
	err = model.ArticleEdit(mod)
	if err != nil {
		Fail(w, "修改失败", err.Error())
		return
	}
	Succ(w, "修改成功")
	return
}

// Fail 错误返回信息
func Fail(w http.ResponseWriter, msg string, data ...interface{}) {
	mod := Reply{
		Code: 300,
		Msg:  msg,
	}
	if len(data) > 0 {
		mod.Data = data[0]
	}
	// 序列化mod
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

// Succ 成功返回信息
func Succ(w http.ResponseWriter, msg string, data ...interface{}) {
	mod := Reply{
		Code: 200,
		Msg:  msg,
	}
	if len(data) > 0 {
		mod.Data = data[0]
	}
	// 序列化mod中所有的数据
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

// Reply 创建一个错误信息结构体
type Reply struct {
	Code int         `json:"code"` // 状态标识 200 成功 300 失败 310 输入有误 320 输出有误
	Msg  string      `json:"msg"`  // 给用户提示信息
	Data interface{} `json:"data"` // 返回数据、报错信息
}

// Ext 分页数据信息 返回 结构体
type Ext struct {
	Count int         `json:"count"`
	Items interface{} `json:"items"`
}
