package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	server := &http.Server{
		Addr: "0.0.0.0:8888",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/get", get)
	mux.HandleFunc("/post", post)
	mux.HandleFunc("/m3u8er", m3u8er)
	mux.HandleFunc("/noter", noter)
	server.Handler = mux
	_ = server.ListenAndServe()
}

// 获取客户端信息
func get(w http.ResponseWriter, r *http.Request) {
	write := "ip:" + r.Header.Get("X-Real-IP") + "\n"
	for k, v := range r.Header {
		write = write + k + ":" + strings.Join(v, ";") + "\n"
	}
	_, _ = w.Write([]byte(write))
}

// 文件上传
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
			newName := "." + from + "." + files["file"][0].Filename + "." + time.Now().Format("20060102_150405")
			_ = os.Rename(imgName, newName)
		}

		saveFile, _ := os.Create(imgName)
		defer saveFile.Close()
		_, _ = io.Copy(saveFile, imgFile) //保存

		http.Redirect(w, r, "/?from="+from, http.StatusSeeOther)
	default:
		fmt.Print("error method")
	}
}

// m3u8下载
func m3u8er(w http.ResponseWriter, r *http.Request) {
	uri := r.PostFormValue("uri")
	uri, _ = url.QueryUnescape(uri)
	name := r.PostFormValue("name")
	fmt.Println(uri, ":", name)
	_ = exec.Command("m3u8downloader", uri, name).Run()
	http.Redirect(w, r, "/?from=/m3u8/", http.StatusSeeOther)
}

func noter(w http.ResponseWriter, r *http.Request) {
	note := r.PostFormValue("note")
	file, _ := os.OpenFile("/home/nginx/note/note.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	_, _ = file.WriteString(time.Now().Format("01-02 15:04") + note + "\n")
	http.Redirect(w, r, "/?from=/note/", http.StatusSeeOther)
}

/*
使用nginx代理
location /post {proxy_pass http://127.0.0.1:8888;client_max_body_size 8192m;}
location /m3u8 {proxy_pass http://127.0.0.1:8888;}
location /get {proxy_pass http://127.0.0.1:8888;}

测试数据
curl http://127.0.0.1:8888/post -F file=@bank.jpg -H "Referer: http://127.0.0.1/?from=/upload/"
*/
