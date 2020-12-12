package core

import (
	"fmt"
	"github.com/Comdex/imgo"
	"github.com/phpdi/ant/image"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

const adb = "/home/yu/DevTools/Android/platform-tools/adb"

//模拟命令
var (
	clockinCmds = []clockinCmd{
		{"",printf("点亮屏幕"),"", lightUp},  //点亮屏幕
		{"",printf("进入主页"),"shell input keyevent 3", nil},   //进入主页
		{"",printf("进入钉钉"),"shell input tap 670 769", sleep(10)},  //进入钉钉
		{"",printf("检查是否需要登录"),"", isLogin},  //检查是否需要登录
		{"",printf("钉钉-工作台"),"shell input tap 530 1840", sleep(5)}, //钉钉-工作台
		{"",printf("钉钉-工作台-考勤打卡"),"shell input tap 977 1612", sleep(5)}, //钉钉-工作台-考勤打卡
		{"",printf("等待蓝牙连接"),"",waitBluetooth},
		{"pro",printf("打卡"),"shell input tap 530 1162",nil},//打卡
		{"pro",printf("打卡后截图，发邮件"),"shell input tap 533 1843",mail},//打卡后截图，发邮件
		{"",printf("返回"),"shell input keyevent 4", nil},   //返回
		{"",printf("返回"),"shell input keyevent 4", nil},   //返回
		{"",printf("返回"),"shell input keyevent 4", nil},   //返回
		{"",printf("熄屏"),"shell input keyevent 26", nil},   //熄屏

	}

)


type (
	clockinCmd struct {
		env string
		before func()bool
		cmdString string
		after func()bool
	}
)

func init()  {

}

func Run(env string) (err error) {
	for _, v := range clockinCmds {
		//检查指令环境
		if v.env!= "" && v.env!=env {
			continue
		}

		if v.before != nil {
			v.before()
		}

		if v.cmdString != "" {

			cmd := adbCommand(v.cmdString)

			if _, err = cmd.CombinedOutput(); err != nil {
				return
			}
		}

		if v.after != nil {
			v.after()
		}
	}

	log.Println("打卡成功")

	return
}

//检查屏幕是否被点亮
func lightUp()  bool{
	imageControl:=new(image.ImageControl)

	for {


		cmd := adbCommand("shell screencap -p | sed 's/\r$//' > data/screen.png")

		if _, err:= cmd.CombinedOutput(); err != nil {
			return false
		}

		imageControl.Trimming("data/screen.png","data/blackscreen_tmp.png",0,1400,750,150)

		cos,err:=imgo.CosineSimilarity("data/blackscreen_tmp.png","data/blackscreen.png")
		if err != nil {
			fmt.Println(err)
			return false
		}

		if int(cos)== 1 {
			//未被点亮

			//亮屏
			cmd := adbCommand("shell input keyevent 26")

			if _, err:= cmd.CombinedOutput(); err != nil {
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
	imageControl:=new(image.ImageControl)

	for {

		//触摸一下防止熄屏
		cmd := adbCommand("shell input tap 870 870")

		if _, err:= cmd.CombinedOutput(); err != nil {
			return false
		}

		cmd = adbCommand("shell screencap -p | sed 's/\r$//' > data/screen.png")

		if _, err:= cmd.CombinedOutput(); err != nil {
			return false
		}

		imageControl.Trimming("data/screen.png","data/bluetooth_tmp.png",0,1400,750,150)

		cos,err:=imgo.CosineSimilarity("data/bluetooth.png","data/bluetooth_tmp.png")
		if err != nil {
			fmt.Println(err)
			return false
		}

		if int(cos)== 1 {
			break
		}

		log.Println("相识度:",cos)
		time.Sleep(2*time.Second)

	}

	return true
}

//登录检测
func isLogin() bool {
	imageControl:=new(image.ImageControl)
	cmd := adbCommand("shell screencap -p | sed 's/\r$//' > data/screen.png")

	if _, err:= cmd.CombinedOutput(); err != nil {
		log.Println("err;",err.Error())
		return false
	}

	imageControl.Trimming("data/screen.png","data/login_tmp.png",0,900,1500,150)


	cos,err:=imgo.CosineSimilarity("data/login.png","data/login_tmp.png")
	if err != nil {
		log.Println(err)
		return false
	}

	if int(cos)== 1 {
		//需要登录
		log.Println("需要登录")
		login()
		time.Sleep(5*time.Second)
	}


	return true
}


//执行登录
func login()  {

	//focus
	cmd := adbCommand("shell input tap 534 817")

	if _, err:= cmd.CombinedOutput(); err != nil {
		fmt.Println("err;",err.Error())
	}

	//password
	cmd = adbCommand(`shell input text "chenyu977564830"`)

	if _, err:= cmd.CombinedOutput(); err != nil {
		log.Println("err;",err.Error())
	}

	//login

	cmd = adbCommand("shell input tap 520 1010")

	if _, err:= cmd.CombinedOutput(); err != nil {
		log.Println("err;",err.Error())
	}




}


//adb命令
func adbCommand(cmd string) *exec.Cmd {
	return exec.Command("/bin/bash", "-c", fmt.Sprintf("%s %s", adb, cmd))
}


func sleep(i int) func()bool  {
	return func() bool {
		time.Sleep(time.Duration(i)*time.Second)
		return true
	}
}

func mail()bool  {
	time.Sleep(5*time.Second)
	cmd := adbCommand("shell screencap -p | sed 's/\r$//' > data/screen.png")

	if _, err:= cmd.CombinedOutput(); err != nil {
		log.Println("err;",err.Error())
	}

	err:=SendMail("打卡通知","data/screen.png")
	if err != nil {
		log.Println(err)
	}

	return true
}

func printf(params string) func()bool {
	return func()bool {
		log.Println(params)
		return true
	}
}


func func_log2fileAndStdout() {
	//创建日志文件
	f, err := os.OpenFile("log/clockin.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//完成后，延迟关闭
	defer f.Close()
	// 设置日志输出到文件
	// 定义多个写入器
	writers := []io.Writer{
		f,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	// 创建新的log对象
	logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
	// 使用新的log对象，写入日志内容
	logger.Println("--> logger :  check to make sure it works")
}