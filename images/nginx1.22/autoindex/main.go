package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	server := &http.Server{
		Addr: "127.0.0.1:8888",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/post", post)
	server.Handler = mux
	server.ListenAndServe()
}

func post(w http.ResponseWriter, r *http.Request) {
	from := r.Header.Get("Referer")
	from = from[strings.Index(from, "=/")+1 : len(from)] //获取相对地址

	switch r.Method {
	case "GET":
		w.Write([]byte(string("method is GET")))
	case "POST":
		r.ParseForm()
		imgFile, _, _ := r.FormFile("file") //获取文件内容

		defer imgFile.Close()
		files := r.MultipartForm.File        //获取表单中的信息
		imgName := files["file"][0].Filename //获取表单文件中name为file的数据
		imgName = "/home/nginx" + from + imgName

		_, err := os.Stat(imgName) //判断文件是否存在,在则隐藏并添加时间戳
		if err == nil {
			newName := "/home/nginx" + from + "." + files["file"][0].Filename + "." + time.Now().Format("20060102_150405")
			os.Rename(imgName, newName)
		}

		saveFile, _ := os.Create(imgName)
		defer saveFile.Close()
		io.Copy(saveFile, imgFile) //保存

		http.Redirect(w, r, "/?from="+from, http.StatusSeeOther)
	default:
		fmt.Print("error method")
	}
}

/*
使用nginx代理
location /post {proxy_pass http://127.0.0.1:8888;client_max_body_size 8192m;}

测试数据
curl http://127.0.0.1:8888/post -F file=@bank.jpg -H "Referer: http://127.0.0.1/?from=/upload/"
*/
