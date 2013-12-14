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

var backupRecord *backupRecordStuct
//var myLogin SelfLoginInfo
//var resourceType []ResourceTypeTable

func init() {
	clientConfig = GetConfig("client")
	backupRecord = newBackupRecord()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//Login()
	//mainLoop()
	startClient()
}

func startClient(){
	thePort, _ := clientConfig.GetString("client","port")
	
	clearScreen()
	fmt.Println("File Resources Mananger")
	fmt.Println("-------------------------------")
	fmt.Println("图书数字资源管理系统客户端已经启动")
	fmt.Println("运行端口为：", thePort, "请将Web界面的端口号与之对应")
	fmt.Println("退出程序请按Ctrl+C或其它")
	
	startServ(thePort)
}

func startServ(thePort string){
	
	http.HandleFunc("/checkLink", wCheckLink)
	http.HandleFunc("/uploadFile", wUploadFile)
	http.HandleFunc("/getBackupRecord", wGetBackupRecord)
	
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
