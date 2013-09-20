//File Resources Manager

package main

import (
	"fmt"
	. "frm_pkg"
	"github.com/msbranco/goconfig"
	"github.com/mewpkg/gopass"
	"os"
)

var clientConfig  *goconfig.ConfigFile
var myLogin SelfLoginInfo

func init() {
	clientConfig = GetConfig("client")
}

func main() {
	Login()
	mainLoop()
}


func mainLoop(){
	clearScreen()
    
	fmt.Println("这里是《文件资源管理系统》\n欢迎", myLogin.Name, "成功登录系统")
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
			default :
				continue
		}
	}
}


func Login() {
	clearScreen()
	
	fmt.Print("您好！请输入用户名和密码进行登录\n")
	fmt.Print("用户名：")
	var username string
	fmt.Scan(&username)
	passwd, _ := gopass.GetPass("密码：")
	doLogin(username, passwd)
}
