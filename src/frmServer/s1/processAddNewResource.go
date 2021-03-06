package s1

import (
	"net"
	. "frmPkg"
	"fmt"
	"time"
	. "frmServer/public"
)

func ProcessAddNewResource(conn *net.TCPConn) {
	theSIDb, _ := ReadSocketBytes(conn, 40)
	
	// begin 查看用户是否存在或超时
	theUser, found  := CkLogedUser (string(theSIDb))
	if found == false {
		SendSocketBytes (conn , Uint8ToBytes(3), 1)
		return
	}
	// end
	
	// start 查看用户是否有建立资源的权力
	if theUser.UPower["resource"]["origin"] < 2 {
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
	
	n_rgt, err := DbConn.Prepare("insert into resourceGroup (hashid, name, rt_id, info, btime, units_id, users_id) values ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		ErrLog.Println("数据库错误：", err)
		return
	}
	sha1string := fmt.Sprint(time.Now(), rgt.Name, theUser.Name)
	rgt.HashId = GetSha1(sha1string)
	rgt.Btime = time.Now().Unix()
	rgt.UnitsId = theUser.UnitId
	rgt.UsersId = theUser.Id
	_, err = n_rgt.Exec(rgt.HashId, rgt.Name, rgt.RtId, rgt.Info, rgt.Btime, rgt.UnitsId, rgt.UsersId)
	if err != nil {
		ErrLog.Println("数据库错误：", err)
		return
	}
	
	n_rgstatus, err := DbConn.Prepare("insert into resourceGroupStatus (hashid) values ($1)")
	if err != nil {
		ErrLog.Println("数据库错误：", err)
		return
	}
	_, err = n_rgstatus.Exec(rgt.HashId)
	if err != nil {
		ErrLog.Println("数据库错误：", err)
		return
	}
	
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	SendSocketBytes (conn, []byte(rgt.HashId),40)
	
	LogInfo.Printf("添加资源条目：%s：用户：%s", rgt.Name, theUser.Name)
	// end
	
	// 更新用户的最后操作时间
	theUser.UpdateLastTime()
}
