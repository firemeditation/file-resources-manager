package main

import (
	"net"
	. "frmPkg"
	"time"
	"fmt"
	"strings"
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
	logInfo.Printf("登录成功：用户：%s", name)
	// 用户名和密码检查完毕
	
	//开始合并权限
	var ckuu UnitsTable  //获取所在Unit的名称和权限
	dbConn.QueryRow("select name, powerlevel from units where id = $1", cku.UnitsId).Scan(&ckuu.Name, &ckuu.PowerLevel)
	
	var ckug GroupsTable  //获取所在Group的权限
	dbConn.QueryRow("select powerlevel from groups where id = $1", cku.GroupsId).Scan(&ckug.PowerLevel)
	
	var cku_p, ckuu_p, ckug_p UserPower
	JsonToStruct(cku.PowerLevel, &cku_p)
	JsonToStruct(ckuu.PowerLevel, &ckuu_p)
	JsonToStruct(ckug.PowerLevel, &ckug_p)
	allpower := mergePower(cku_p, ckuu_p, ckug_p)
	
	ckuu.Name = strings.Trim(ckuu.Name, " ")
	
	//开始生成SelfLoginInfo和UserIsLogin
	sha1 = GetSha1(sha1 + name)
	thisU, _ := userLoginStatus.Add(sha1, cku.Id, name, cku.GroupsId, cku.UnitsId, ckuu.Name, time.Now())
	thisU.UPower = allpower
	
	nameSelfLogin := NewSelfLoginInfo(cku.Id, name, cku.GroupsId, cku.UnitsId, ckuu.Name, sha1)
	nameSelfLogin.UPower = allpower
	
	gob_b := StructGobBytes(nameSelfLogin)  //将结构体转为gob，进而转成bytes
	
	gob_len := len(gob_b)
	
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	SendSocketBytes (conn , Uint64ToBytes(uint64(gob_len)), 8)
	SendSocketBytes (conn , gob_b, uint64(gob_len))
	
	//开始发送所有的资源类型
	var resourceType []ResourceTypeTable
	rts, _ := dbConn.Query("select * from resourcetype")
	for rts.Next(){
		var onert ResourceTypeTable
		rts.Scan(&onert.Id, &onert.Name, &onert.PowerLevel, &onert.Expend, &onert.Info)
		onert.Name = strings.Trim(onert.Name, " ")
		resourceType = append(resourceType, onert)
	}
	rt_b := StructGobBytes(resourceType)
	rt_b_len := len(rt_b)
	SendSocketBytes (conn , Uint64ToBytes(uint64(rt_b_len)), 8)
	SendSocketBytes (conn , rt_b, uint64(rt_b_len))
}
