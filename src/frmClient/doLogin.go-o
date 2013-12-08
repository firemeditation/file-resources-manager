package main

import (
	"os"
	"fmt"
	. "frmPkg"
	//"encoding/gob"
	//"bytes"
)

func doLogin (username, passwd string) bool {
	conn := connectServer()
	err := sendTheFirstRequest (1, 1, conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "发送出错：", err)
		os.Exit(1)
	}
	ckl, _ := ReadSocketBytes(conn, 1)
	if BytesToUint8(ckl) != 1 {
		fmt.Fprintf(os.Stderr, "服务器端禁止登录")
		os.Exit(1)
	}
	
	lsha , _ := ReadSocketBytes(conn, 40)
	passwd = GetSha1(passwd)  //密码先sha1
	passwd = passwd + string(lsha)  //密码+返回的随机sha1
	passwd = GetSha1(passwd)  //密码再sha1
	
	SendSocketBytes (conn, Uint16ToBytes(uint16(len(username))), 2)   //用户名长度
	SendSocketBytes (conn, []byte(username), uint64(len(username)))  //用户名
	SendSocketBytes (conn, []byte(passwd),40)  //密码
	
	ckl, _ = ReadSocketBytes(conn, 1)
	if BytesToUint8(ckl) != 1 {
		fmt.Print("服务器未验证通过，也许是用户名密码错误，需要重新登录\n")
		os.Exit(1)
	}
	
	sli_len_b, _ := ReadSocketBytes(conn, 8)
	sli_len := BytesToUint64(sli_len_b)
	sli_b, _ := ReadSocketBytes(conn, sli_len)
	BytesGobStruct(sli_b, &myLogin)
	
	rt_len_b, _ := ReadSocketBytes(conn, 8)
	rt_len := BytesToUint64(rt_len_b)
	rt_b,_ := ReadSocketBytes(conn, rt_len)
	BytesGobStruct(rt_b,&resourceType)
	
	return true
}
