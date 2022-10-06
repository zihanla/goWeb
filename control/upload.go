package control

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
	"unsafe"
)

func ApiUpload(w http.ResponseWriter, r *http.Request) {
	// 给20mb空间作为请求缓存 把文件保存到容器
	// 创建一个大的容器空间 1 << 20 把1往左边位移20位  位移 10位 2的十次方
	r.ParseMultipartForm(1 << 20 * 20) // 20mb
	// file 文件本身  header 上传文件的信息
	file, header, err := r.FormFile("upfile")
	if err != nil {
		Fail(w, "上传失败", err.Error())
	}
	os.MkdirAll("static/upload/", 0666)
	// 解决重复上传问题
	ext := path.Ext(header.Filename)
	name := "static/upload/" + RandStr(10) + ext
	// 通过 header.Filename(文件信息) 找到file 并创建文件
	dst, _ := os.Create(name)
	// 把 file 文件 赋值给刚刚创建的文件
	io.Copy(dst, file)
	file.Close()
	dst.Close()
	//Succ(w, "上传成功", "/"+name)
	w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte("{\"original\":\"mobileBg.jpg\",\"state\":\"SUCCESS\",\"title\":\"mobileBg.jpg\",\"url\":\"" + ("/" + name) + "\"}"))
	mod := editorReply{
		Original: "header.Filename",
		State:    "SUCCESS",
		Title:    "header.Filename",
		Url:      "/" + name,
	}
	w.Write(mod.Json())
}

// 编辑器返回信息 专用
type editorReply struct {
	Original string `json:"original"`
	State    string `json:"state"`
	Title    string `json:"title"`
	Url      string `json:"url"`
}

// Json 序列化方法
func (er *editorReply) Json() []byte {
	buf, _ := json.Marshal(er)
	return buf
}

var all = "OsuzN8sWngIiq8DOO2VhqRPMzS6kmrUkdwEZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandStr 构建随机字符串
func RandStr(ln int) string {
	// 设置随机字串生成时间
	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]byte, ln, ln)
	for i := 0; i < ln; i++ {
		res[i] = all[rand.Intn(36)]
	}
	return *(*string)(unsafe.Pointer(&res))
}
