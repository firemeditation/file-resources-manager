//File Resources Manager

package main

import (
	"fmt"
	. "frm_pkg"
	"github.com/msbranco/goconfig"
	"github.com/mewpkg/gopass"
	"net"
	"os"
)

var clientConfig  *goconfig.ConfigFile
var myLogin SelfLoginInfo

func init() {
	clientConfig = GetConfig("client")
}

func main() {
	doLogin()
}


func doLogin() {
	fmt.Print("您好！请输入用户名和密码进行登录\n")
	fmt.Print("用户名：")
	var username string
	fmt.Scan(&username)
	passwd, _ := gopass.GetPass("密码：")
	fmt.Println("用户名为",username,"密码为", passwd)
}

func connectServer() *net.TCPConn {
	serviceAddr, _ := clientConfig.GetString("server","addr")
	IPAdrr, err := net.ResolveTCPAddr("tcp", serviceAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "连接服务器出错：", err)
		os.Exit(1)
	}
	Connecter, err := net.DialTCP("tcp", nil, IPAdrr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "连接服务器出错：", err)
		os.Exit(1)
	}
	return Connecter
}
