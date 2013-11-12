package main

import (
	"fmt"
)

func mainUploadResource(){
	if myLogin.UPower["resource"]["origin"] < 2 {
		fmt.Print("您没有新建资源条目的权力，按任意键继续。")
		var tep string
		fmt.Scanln(&tep)
		return
	}
	
	var r_hash string
	fmt.Print("请输入资源的HashID：")
	fmt.Scanln(&r_hash)
	uploadResourceFile(r_hash)
}
