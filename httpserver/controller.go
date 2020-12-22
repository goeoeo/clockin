package httpserver

import (
	"fmt"
	"github.com/phpdi/clockin/core"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func HttpServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/clockin", clockin)

	server := &http.Server{
		Addr:         "127.0.0.1:8122",
		Handler:      mux,
		WriteTimeout: time.Second * 3,
	}

	go server.ListenAndServe()
	go npc()

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)

	<-c

	fmt.Println("进程已退出")
}

//监听执行打卡程序
func clockin(w http.ResponseWriter, r *http.Request) {

	if isAjax(r) {
		key := r.URL.Query().Get("key")
		switch key {
		case "test":
			go func() {
				core.Run("")
			}()
		case "pro":
			go func() {
				core.Run("pro")
			}()
		}

		w.Write([]byte("success"))

		return
	}

	//页面输出
	tmpl, err := template.ParseFiles("httpserver/view.html")
	if err != nil {
		fmt.Println("Error happened..")
		return
	}

	tmpl.Execute(w, nil)
}

func isAjax(r *http.Request) bool {
	fmt.Println(r.Header.Get("X-Requested-With"))
	return r.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

//内网穿透工具
func npc() {
	cmnd := exec.Command("./npc.sh")
	if err := cmnd.Start(); err != nil {
		fmt.Println("npc:", err)
	}
}
