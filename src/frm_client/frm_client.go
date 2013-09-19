//File Resources Manager

package main

import (
	"fmt"
	. "frm_pkg"
	"github.com/msbranco/goconfig"
	"github.com/mewpkg/gopass"
)

var clientConfig  *goconfig.ConfigFile
var myLogin SelfLoginInfo

func init() {
	clientConfig = GetConfig("client")
}

func main() {
	Login()
}


func Login() {
	fmt.Print("您好！请输入用户名和密码进行登录\n")
	fmt.Print("用户名：")
	var username string
	fmt.Scan(&username)
	passwd, _ := gopass.GetPass("密码：")
	doLogin(username, passwd)
}
