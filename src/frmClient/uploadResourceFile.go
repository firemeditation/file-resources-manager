package main

import (
	"fmt"
	. "frmPkg"
)


// uploadResourceFile 上传资源文件
func uploadResourceFile(resourceGroup string) {
	// start 自身判断权限
	if myLogin.UPower["resource"]["origin"] < 2 {
		fmt.Print("您没有新建资源条目的权力，按任意键继续。")
		var tep string
		fmt.Scanln(&tep)
		return
	}
	// end
	
	//uploadtype := 1
	addtopath := "./"
	var originpath string
	for {
		fmt.Print("选择上传类型：1.普通上传\t2定点追加\t(同位置同名将自动覆盖)\n")
		var otype string
		fmt.Scanln(&otype)
		fmt.Print("请输入本地上传文件所在路径：\n")
		fmt.Scanln(&originpath)
		if otype == "1" {
			//uploadtype = 1
			break
		}else if otype == "2"{
			//uploadtype = 2
			fmt.Print("请输入要追加的位置（用“./”代表在源文件路径根部追加）：")
			fmt.Scanln(&addtopath)
			break
		}
	}
	originpath = DirMustEnd(originpath)
	addtopath = DirMustEnd(addtopath)
	doUploadResourceFile(resourceGroup, originpath,addtopath)
}
