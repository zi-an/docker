package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// 2025-08-17
func main() {
	Init()
	server := &http.Server{
		Addr: "0.0.0.0:8888",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/get", get)
	mux.HandleFunc("/post", post)
	mux.HandleFunc("/m3u8", m3u8)
	mux.HandleFunc("/note", note)
	server.Handler = mux
	_ = server.ListenAndServe()
}

var baseDir = "/home/nginx/"

func Init() {
	dirs := []string{"/m3u8/", "/note/"}
	if runtime.GOOS == "windows" {
		baseDir = "z:"
	}
	for _, dir := range dirs {
		dir := baseDir + dir
		_, err := os.Stat(dir)
		if err != nil {
			_ = os.MkdirAll(dir, 0777)
		}
	}
}

// 获取客户端信息
func get(w http.ResponseWriter, r *http.Request) {
	write := "ip:" + r.Header.Get("X-Real-IP") + "<br>"
	for k, v := range r.Header {
		write = write + k + ":" + strings.Join(v, ";") + "<br>"
	}
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	_, _ = w.Write([]byte(write))
}

// 文件上传
// 测试数据: curl http://127.0.0.1:8888/post -F file=@bank.jpg -H "Referer: http://127.0.0.1/?from=/upload/"
func post(w http.ResponseWriter, r *http.Request) {
	from := r.Header.Get("Referer")
	from = from[strings.Index(from, "=/")+1 : len(from)] //获取相对地址

	switch r.Method {
	case "GET":
		_, _ = w.Write([]byte(string("method is GET")))
	case "POST":
		_ = r.ParseForm()
		imgFile, _, _ := r.FormFile("file") //获取文件内容

		defer imgFile.Close()
		files := r.MultipartForm.File        //获取表单中的信息
		imgName := files["file"][0].Filename //获取表单文件中name为file的数据
		imgName = "." + from + imgName

		_, err := os.Stat(imgName) //判断文件是否存在,在则隐藏并添加时间戳
		if err == nil {
			loc, _ := time.LoadLocation("Asia/Taipei")
			newName := "." + from + "." + files["file"][0].Filename + "." + time.Now().In(loc).Format("20060102_150405")
			_ = os.Rename(imgName, newName)
		}

		saveFile, err := os.Create(imgName)
		fmt.Print(err)
		defer saveFile.Close()
		_, _ = io.Copy(saveFile, imgFile) //保存

		http.Redirect(w, r, "/?from="+from, http.StatusSeeOther)
	default:
		fmt.Print("error method")
	}
}

// m3u8下载
func m3u8(w http.ResponseWriter, r *http.Request) {
	uri := r.PostFormValue("uri")
	uri, _ = url.QueryUnescape(uri)
	name := r.PostFormValue("name")
	fmt.Println(uri, ":", name)
	file, _ := os.OpenFile(baseDir+"/note/m3u8.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	loc, _ := time.LoadLocation("Asia/Taipei")
	now := time.Now().In(loc).Format("0102_1504")
	file.WriteString(now + " | " + name + " | " + uri + "\n")
	_ = exec.Command("m3u8downloader", uri, name).Run()
	http.Redirect(w, r, "/?from=/m3u8/", http.StatusSeeOther)
}

func note(w http.ResponseWriter, r *http.Request) {
	note := r.PostFormValue("note")
	file, _ := os.OpenFile(baseDir+"/note/note.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	loc, _ := time.LoadLocation("Asia/Taipei")
	_, _ = file.WriteString(time.Now().In(loc).Format("0102_1504: ") + note + "\n")
	http.Redirect(w, r, "/?from=/note/", http.StatusSeeOther)
}

// windows 下调试把工作目录设置为z:/
// 最小化编译 go build -ldflags="-s -w" -o goweb goweb.go
// tar -zcvf goweb.tgz goweb
