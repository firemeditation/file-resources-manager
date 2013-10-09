//File Resources Manager

package main

import (
	"fmt"
	. "frmPkg"
	"github.com/msbranco/goconfig"
	"os"
)

var clientConfig  *goconfig.ConfigFile
var myLogin SelfLoginInfo
var resourceType []ResourceTypeTable

func init() {
	clientConfig = GetConfig("client")
}

func main() {
	Login()
	mainLoop()
}


func mainLoop(){
	clearScreen()
    
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
			default :
				continue
		}
	}
}
