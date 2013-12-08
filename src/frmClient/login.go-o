package main

import (
	"github.com/mewpkg/gopass"
	"fmt"
)

func Login() {
	clearScreen()
	
	fmt.Print("您好！请输入用户名和密码进行登录\n")
	fmt.Print("用户名：")
	var username string
	fmt.Scan(&username)
	passwd, _ := gopass.GetPass("密码：")
	doLogin(username, passwd)
}
