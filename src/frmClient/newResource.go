package main

import (
	"fmt"
	. "frmPkg"
)

// newResource 新建资源条目
func newResource () {
	
	// start 自身判断权限
	if myLogin.UPower["resource"]["origin"] < 10 {
		fmt.Print("您没有新建资源条目的权力，按任意键继续。")
		var tep string
		fmt.Scanln(&tep)
		return
	}
	// end
	
	conn := connectServer()
	err := sendTheFirstRequest (1, 2, conn)
	if err != nil {
		fmt.Print("发送状态错误：", err)
		return
	}
	
	//发送自己的SID
	err = SendSocketBytes(conn, []byte(myLogin.SID), 40)
	if err != nil {
		fmt.Print("发送SID错误：", err)
		return
	}
	
	// start 查看服务器是否同意添加
	ckl, _ := ReadSocketBytes(conn, 1)
	
	if BytesToUint8(ckl) == 3 {
		fmt.Print("服务器端身份验证失败，可能是连接超时，请重新登录。按任意键继续。")
		var tep string
		fmt.Scanln(&tep)
		return
	}
	
	if BytesToUint8(ckl) == 2 {
		fmt.Print("服务器端禁止添加条目，可能是没有权限，按任意键继续。")
		var tep string
		fmt.Scanln(&tep)
		return
	}
	// end
}
