package core

import (
	"fmt"
	"github.com/Comdex/imgo"
	"github.com/phpdi/ant/image"
	"log"
	"os/exec"
	"time"
)

const (
	adb       = "/home/yu/DevTools/Android/platform-tools/adb"
	runEnvPro = "pro"
)

//模拟命令
var (
	clockinCmds = []clockinCmd{
		{"", printf("点亮屏幕"), "", lightUp},                          //点亮屏幕
		{"", printf("进入主页"), "shell input keyevent 3", nil},        //进入主页
		{"", printf("进入钉钉"), "shell input tap 670 769", sleep(10)}, //进入钉钉
		{"", printf("检查是否需要登录"), "", isLogin},                      //检查是否需要登录
		{"", printf("钉钉-消息"), "shell input tap 119 1827", sleep(5)},
		{"", printf("钉钉-智能工作助理"), "shell input tap 587 740", sleep(8)},
		{"", printf("钉钉-点击打卡"), "shell input tap 148 1688", sleep(5)},
		{"", printf("等待蓝牙连接"), "", waitBluetooth},
		{"pro", printf("打卡"), "", clockin},                             //打卡
		{"pro", printf("打卡后截图，发邮件"), "shell input tap 533 1843", mail}, //打卡后截图，发邮件
		{"", printf("返回"), "shell input keyevent 4", nil},              //返回
		{"", printf("返回"), "shell input keyevent 4", nil},              //返回
		{"", printf("返回"), "shell input keyevent 4", nil},              //返回
		{"", printf("熄屏"), "shell input keyevent 26", nil},             //熄屏

	}
	runEnv = "" //环境

	runing = false // 正在运行中
)

type (
	clockinCmd struct {
		env       string
		before    func() bool
		cmdString string
		after     func() bool
	}
)

func Run(env string) (err error) {
	if runing {
		fmt.Println("打卡正在运行中")
		return
	}

	runEnv = env
	runing = true
	for _, v := range clockinCmds {
		//检查指令环境
		if v.env != "" && v.env != env {
			continue
		}

		if v.before != nil {
			v.before()
		}

		if v.cmdString != "" {

			cmd := AdbCommand(v.cmdString)

			if _, err = cmd.CombinedOutput(); err != nil {
				return
			}
		}

		if v.after != nil {
			v.after()
		}
	}
	runing = false
	log.Println("打卡成功")

	return
}

//检查屏幕是否被点亮
func lightUp() bool {

	for {

		cmd := AdbCommand("shell screencap -p | sed 's/\r$//' > data/screen.png")

		if _, err := cmd.CombinedOutput(); err != nil {
			return false
		}

		cos, err := imgo.CosineSimilarity("data/screen.png", "data/blackscreen.png")
		if err != nil {
			fmt.Println(err)
			return false
		}

		if int(cos) == 1 {
			//未被点亮
			//亮屏
			cmd := AdbCommand("shell input keyevent 26")

			if _, err := cmd.CombinedOutput(); err != nil {
				return false
			}

			continue
		}

		log.Println("已点亮屏幕")
		break

	}

	return true
}

//等待蓝牙连接
func waitBluetooth() bool {
	imageControl := new(image.ImageControl)

	for {

		if !effectiveTime() && runEnv == runEnvPro {
			return false
		}

		//触摸一下防止熄屏
		cmd := AdbCommand("shell input tap 870 870")

		if _, err := cmd.CombinedOutput(); err != nil {
			return false
		}

		cmd = AdbCommand("shell screencap -p | sed 's/\r$//' > data/screen.png")

		if _, err := cmd.CombinedOutput(); err != nil {
			return false
		}

		//上午判定
		imageControl.Trimming("data/screen.png", "data/bluetooth_tmp.png", 400, 1300, 350, 350)

		cos, err := imgo.CosineSimilarity("data/bluetooth_start.png", "data/bluetooth_tmp.png")
		if err != nil {
			fmt.Println(err)
			return false
		}
		if int(cos) == 1 {
			break
		}
		log.Println("相识度:", cos)

		//下午判定
		imageControl.Trimming("data/screen.png", "data/bluetooth_tmp.png", 400, 1400, 350, 750)

		cos, err = imgo.CosineSimilarity("data/bluetooth_end.png", "data/bluetooth_tmp.png")
		if err != nil {
			fmt.Println(err)
			return false
		}
		if int(cos) == 1 {
			break
		}
		log.Println("相识度:", cos)

		time.Sleep(2 * time.Second)

	}

	return true
}

//登录检测
func isLogin() bool {
	imageControl := new(image.ImageControl)
	cmd := AdbCommand("shell screencap -p | sed 's/\r$//' > data/screen.png")

	if _, err := cmd.CombinedOutput(); err != nil {
		log.Println("err;", err.Error())
		return false
	}

	imageControl.Trimming("data/screen.png", "data/login_tmp.png", 423, 974, 200, 70)

	cos, err := imgo.CosineSimilarity("data/login.png", "data/login_tmp.png")
	if err != nil {
		log.Println(err)
		return false
	}

	if int(cos) == 1 {
		//需要登录
		log.Println("需要登录")
		login()
		time.Sleep(5 * time.Second)
	}

	return true
}

//执行登录
func login() {

	//focus
	cmd := AdbCommand("shell input tap 534 817")

	if _, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("err;", err.Error())
	}

	//password
	cmd = AdbCommand(`shell input text "chenyu977564830"`)

	if _, err := cmd.CombinedOutput(); err != nil {
		log.Println("err;", err.Error())
	}

	//login

	cmd = AdbCommand("shell input tap 520 1010")

	if _, err := cmd.CombinedOutput(); err != nil {
		log.Println("err;", err.Error())
	}

}

//adb命令
func AdbCommand(cmd string) *exec.Cmd {
	return exec.Command("/bin/bash", "-c", fmt.Sprintf("%s %s", adb, cmd))
}

func sleep(i int) func() bool {
	return func() bool {
		time.Sleep(time.Duration(i) * time.Second)
		return true
	}
}

func mail() bool {

	time.Sleep(5 * time.Second)
	cmd := AdbCommand("shell screencap -p | sed 's/\r$//' > data/screen.png")

	if _, err := cmd.CombinedOutput(); err != nil {
		log.Println("err;", err.Error())
	}

	err := SendMail("打卡通知", "data/screen.png")
	if err != nil {
		log.Println(err)
	}

	return true
}

func printf(params string) func() bool {
	return func() bool {
		log.Println(params)
		return true
	}
}

func clockin() bool {

	//9点半以前，
	cmd := AdbCommand("shell input tap 530 1162")

	if _, err := cmd.CombinedOutput(); err != nil {
		log.Println("err;", err.Error())
	}

	return true

}

//有效的打卡时间
func effectiveTime() bool {
	h := time.Now().Hour()
	m := time.Now().Minute()

	switch {
	case h < 9: //9点以前
		return true
	case h == 9 && m <= 30: //9点到9点半
		return true
	case h > 18: //6点以后
		return true
	case h == 18 && m >= 30: //6点半以后
		return true
	}

	return false

}
