package main

import (
	"fmt"
	"github.com/phpdi/clockin/core"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

//go build -o clockinbin main.go
func main() {
	httpServer()

}

func httpServer() {
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
}

//内网穿透工具
func npc() {
	cmnd := exec.Command("./npc.sh")
	if err := cmnd.Start(); err != nil {
		fmt.Println("npc:", err)
	}
}
