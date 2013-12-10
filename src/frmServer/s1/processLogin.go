package s1

import (
	"net"
	. "frmPkg"
	"time"
	"fmt"
	"strings"
	. "frmServer/public"
)

func ProcessLogin(conn *net.TCPConn) {
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	sha1 := GetSha1(fmt.Sprintln(time.Now()))
	SendSocketBytes (conn , []byte(sha1), 40)  //发送SHA1
	
	name_l_b , _ := ReadSocketBytes(conn, 2)
	name_l := BytesToUint16(name_l_b)
	name_b, _ := ReadSocketBytes(conn, uint64(name_l))
	name := string(name_b)
	
	// 禁止nobody登录
	nobody, _ := ServerConfig.GetString("user","nobody")
	if name == nobody {
		LogInfo.Printf("登录错误：用户不存在：用户：%s", name)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	
	passwd_b ,_ := ReadSocketBytes(conn, 40)
	passwd := string(passwd_b)
	
	// 开始检查用户名和密码，并将需要返回的东西全部返回
	var cku UsersTable
	err := DbConn.QueryRow("select id, passwd,  units_id, groups_id, powerlevel from users where name = $1", name).Scan(&cku.Id, &cku.Passwd, &cku.UnitsId, &cku.GroupsId, &cku.PowerLevel)
	if err != nil {
		LogInfo.Printf("登录错误：用户不存在：用户：%s", name)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	
	ck_passwd := cku.Passwd + sha1
	ck_passwd = GetSha1(ck_passwd)
	
	if passwd != ck_passwd {
		LogInfo.Printf("登录错误：密码错误：用户：%s", name)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	LogInfo.Printf("登录成功：用户：%s", name)
	// 用户名和密码检查完毕
	
	//开始合并权限
	var ckuu UnitsTable  //获取所在Unit的名称和权限
	DbConn.QueryRow("select name, powerlevel from units where id = $1", cku.UnitsId).Scan(&ckuu.Name, &ckuu.PowerLevel)
	
	var ckug GroupsTable  //获取所在Group的权限
	DbConn.QueryRow("select name, powerlevel from groups where id = $1", cku.GroupsId).Scan(&ckug.Name , &ckug.PowerLevel)
	
	var cku_p, ckuu_p, ckug_p UserPower
	JsonToStruct(cku.PowerLevel, &cku_p)
	JsonToStruct(ckuu.PowerLevel, &ckuu_p)
	JsonToStruct(ckug.PowerLevel, &ckug_p)
	allpower := MergePower(cku_p, ckuu_p, ckug_p)
	
	ckuu.Name = strings.TrimSpace(ckuu.Name)
	ckug.Name = strings.TrimSpace(ckug.Name)
	
	//开始生成SelfLoginInfo和UserIsLogin
	sha1 = GetSha1(sha1 + name)
	thisU, _ := UserLoginStatus.Add(sha1, sha1, cku.Id, name, cku.GroupsId, ckug.Name, cku.UnitsId, ckuu.Name, time.Now())
	thisU.UPower = allpower
	
	nameSelfLogin := NewSelfLoginInfo(sha1, cku.Id, name, cku.GroupsId, ckug.Name, cku.UnitsId, ckuu.Name, sha1)
	nameSelfLogin.UPower = allpower
	
	gob_b := StructGobBytes(nameSelfLogin)  //将结构体转为gob，进而转成bytes
	
	gob_len := len(gob_b)
	
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	SendSocketBytes (conn , Uint64ToBytes(uint64(gob_len)), 8)
	SendSocketBytes (conn , gob_b, uint64(gob_len))
	
	//开始发送所有的资源类型
	var resourceType []ResourceTypeTable
	rts, _ := DbConn.Query("select * from resourcetype")
	for rts.Next(){
		var onert ResourceTypeTable
		rts.Scan(&onert.Id, &onert.Name, &onert.PowerLevel, &onert.Expend, &onert.Info)
		onert.Name = strings.TrimSpace(onert.Name)
		resourceType = append(resourceType, onert)
	}
	rt_b := StructGobBytes(resourceType)
	rt_b_len := len(rt_b)
	SendSocketBytes (conn , Uint64ToBytes(uint64(rt_b_len)), 8)
	SendSocketBytes (conn , rt_b, uint64(rt_b_len))
}
