package main

import (
	"fmt"
	. "frmPkg"
)

func doNewResource (rgt *ResourceGroupTable) {
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
	
	// begin 发送新建资源的信息
	rgt_b := StructGobBytes(rgt)
	rgt_b_l := uint64(len(rgt_b))
	SendSocketBytes(conn,Uint64ToBytes(rgt_b_l),8)
	SendSocketBytes(conn,rgt_b,rgt_b_l)
	
	ckl, _ = ReadSocketBytes(conn, 1)
	if BytesToUint8(ckl) != 1 {
		fmt.Print("添加资源出错，请重试，按任意键继续。")
		var tep string
		fmt.Scanln(&tep)
		return
	}
	newhash_b,_ := ReadSocketBytes(conn, 40)
	fmt.Println("添加成功，新资源的HashID为：",string(newhash_b))
	// end
}
