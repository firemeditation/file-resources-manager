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
	passwd_b ,_ := ReadSocketBytes(conn, 40)
	passwd := string(passwd_b)
	
	// 开始检查用户名和密码，并将需要返回的东西全部返回
	/*var ck_passwd string
	var units_id int
	var groups_id uint16
	var powerlevel string
	err := dbConn.QueryRow("select passwd,  units_id, groups_id, powerlevel from users where name = $1", name).Scan(&ck_passwd, &units_id, &groups_id, &powerlevel)
	if err != nil {
		logInfo.Printf("登录错误：用户不存在：用户：%s", name)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	*/
	var cku UsersTable
	err := dbConn.QueryRow("select id, passwd,  units_id, groups_id, powerlevel from users where name = $1", name).Scan(&cku.Id, &cku.Passwd, &cku.UnitsId, &cku.GroupsId, &cku.PowerLevel)
	if err != nil {
		logInfo.Printf("登录错误：用户不存在：用户：%s", name)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	
	ck_passwd := cku.Passwd + sha1
	ck_passwd = GetSha1(ck_passwd)
	
	if passwd != ck_passwd {
		logInfo.Printf("登录错误：密码错误：用户：%s", name)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	} 
	// 用户名和密码检查完毕
	logInfo.Printf("登录成功：用户：%s", name)
	
	//开始生成SelfLoginInfo和UserIsLogin
	sha1 = GetSha1(sha1 + name)
	userLoginStatus.Add(sha1, cku.Id, name, cku.GroupsId, time.Now())
	nameSelfLogin := NewSelfLoginInfo(cku.Id, name, cku.GroupsId, sha1)
	
	nameSelfLogin.UPower["main"]["power1"] = 2
	
	gob_b := StructGobBytes(nameSelfLogin)  //将结构体转为gob，进而转成bytes
	
	gob_len := len(gob_b)
	
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	SendSocketBytes (conn , Uint64ToBytes(uint64(gob_len)), 8)
	SendSocketBytes (conn , gob_b, uint64(gob_len))
}
