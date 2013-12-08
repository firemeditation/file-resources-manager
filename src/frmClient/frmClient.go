//File Resources Manager

package main

import (
	"fmt"
	. "frmPkg"
	"github.com/msbranco/goconfig"
	"os"
	"runtime"
	"net/http"
)

var clientConfig  *goconfig.ConfigFile
//var myLogin SelfLoginInfo
//var resourceType []ResourceTypeTable

func init() {
	clientConfig = GetConfig("client")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//Login()
	//mainLoop()
	startClient()
}

func startClient(){
	thePort, _ := clientConfig.GetString("client","port")
	
	go startServ(thePort)
	
	clearScreen()
	fmt.Println("File Resources Mananger")
	fmt.Println("-------------------------------")
	fmt.Println("图书数字资源管理系统客户端已经启动")
	fmt.Println("运行端口为：", thePort, "请将Web界面的端口号与之对应")
	
	for {
		fmt.Print("退出客户端请按数字0并回车：")
		var otype string
		fmt.Scanln(&otype)
		if otype == "0" {
			os.Exit(0)
		}
	}
}

func startServ(thePort string){
	
	http.HandleFunc("/checkLink", wCheckLink)
	
	theServ := "127.0.0.1:" + thePort
	err := http.ListenAndServe(theServ, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "出错了，错误是：", err)
		os.Exit(1)
	}
}


/*func mainLoop(){
	clearScreen()
	
	thPort, _ := clientConfig.GetInt64("client","port")
    
	fmt.Printf("这里是《文件资源管理系统》\n欢迎%s的%s成功登录系统", myLogin.UnitName, myLogin.Name)
	for {
		fmt.Print("\n")
		fmt.Print("请选择如下操作：\n")
		fmt.Print("1. 搜索资源条目\t\t2. 新建资源条目\t\t3. 上传资源\n")
		fmt.Print("4. 查看个人权限\t\t0. 退出程序\n")
		fmt.Print("请选择：")
		var otype string
		fmt.Scanln(&otype)
		switch otype {
			case "0":
				os.Exit(0)
			case "2":
				newResource()
			case "3":
				mainUploadResource()
			default :
				continue
		}
	}
}*/
