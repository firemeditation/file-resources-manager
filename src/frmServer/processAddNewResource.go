package main

import (
	"net"
	. "frmPkg"
	"fmt"
	"time"
)

func processAddNewResource(conn *net.TCPConn) {
	theSIDb, _ := ReadSocketBytes(conn, 40)
	
	// begin 查看用户是否存在或超时
	theUser, found  := ckLogedUser (string(theSIDb))
	if found == false {
		SendSocketBytes (conn , Uint8ToBytes(3), 1)
		return
	}
	// end
	
	// start 查看用户是否有建立资源的权力
	if theUser.UPower["resource"]["origin"] < 10 {
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}else{
		SendSocketBytes (conn , Uint8ToBytes(1), 1)
	}
	// end
	
	// begin 接收并写入数据库
	rgt_b_l_b, _ := ReadSocketBytes(conn, 8)
	rgt_b_l := BytesToUint64(rgt_b_l_b)
	rgt_b , _ := ReadSocketBytes(conn, rgt_b_l)
	var rgt ResourceGroupTable
	BytesGobStruct(rgt_b, &rgt)
	
	n_rgt, err := dbConn.Prepare("insert into resourceGroup (hashid, name, rt_id, info, btime, units_id, users_id) values ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		fmt.Println("prepare错误：",err)
	}
	sha1string := fmt.Sprint(time.Now(), rgt.Name, theUser.Name)
	rgt.HashId = GetSha1(sha1string)
	rgt.Btime = time.Now().Unix()
	rgt.UnitsId = theUser.UnitId
	rgt.UsersId = theUser.Id
	_, err = n_rgt.Exec(rgt.HashId, rgt.Name, rgt.RtId, rgt.Info, rgt.Btime, rgt.UnitsId, rgt.UsersId)
	if err != nil {
		fmt.Println("exec错误：",err)
	}
	
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	SendSocketBytes (conn, []byte(rgt.HashId),40)
	
	logInfo.Printf("添加资源条目：%s：用户：%s", rgt.Name, theUser.Name)
	// end
}
