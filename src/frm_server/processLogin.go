package main

import (
	"net"
	. "frm_pkg"
	"time"
	"fmt"
	//"bytes"
	//"encoding/gob"
)

func processLogin(conn *net.TCPConn) {
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	sha1 := GetSha1(fmt.Sprintln(time.Now()))
	SendSocketBytes (conn , []byte(sha1), 40)  //发送SHA1
	
	name_l_b , _ := ReadSocketBytes(conn, 2)
	name_l := BytesToUint16(name_l_b)
	name_b, _ := ReadSocketBytes(conn, uint64(name_l))
	name := string(name_b)
	fmt.Println("用户名：", name)
	passwd_b ,_ := ReadSocketBytes(conn, 40)
	passwd := string(passwd_b)
	fmt.Println("密码：", passwd)
	
	//开始生成SelfLoginInfo和UserIsLogin
	sha1 = GetSha1(sha1 + name)
	userLoginStatus.Add(sha1, name, 100, time.Now(), 1)
	nameSelfLogin := NewSelfLoginInfo(name, 100, sha1, 1)
	
	nameSelfLogin.UPower["main"]["power1"] = 2
	
	gob_b := StructGobBytes(nameSelfLogin)  //将结构体转为gob，进而转成bytes
	
	gob_len := len(gob_b)
	
	fmt.Println("长度",gob_len)
	
	fmt.Println(userLoginStatus[sha1].LastTime)
	
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	SendSocketBytes (conn , Uint64ToBytes(uint64(gob_len)), 8)
	SendSocketBytes (conn , gob_b, uint64(gob_len))
}
