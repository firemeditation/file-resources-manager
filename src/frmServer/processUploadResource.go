package main


import (
	"net"
	. "frmPkg"
)


func processUploadResource(conn *net.TCPConn) {
	theSIDb, _ := ReadSocketBytes(conn, 40) //用户ID
	ReadSocketBytes(conn,1)  //用户请求（在这里忽略）
	theBook_b , _ := ReadSocketBytes(conn,40)  //图书ID
	
	// begin 查看用户是否存在或超时
	theUser, found  := ckLogedUser (string(theSIDb))
	if found == false {
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	// end
	// start 查看用户是否有资源管理的权力
	if theUser.UPower["resource"]["origin"] < 2 {
		logInfo.Printf("上传错误：用户无权限上传资源文件：用户：%s，资源：%s", theUser.Name, string(theBook_b))
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	// end
	// begin 查看这个资源是不是用户可写的
	var ckBook ResourceGroupTable
	err := dbConn.QueryRow("select units_id, powerlevel from resourceGroup where hashid = $1", string(theBook_b)).Scan(&ckBook.UnitsId, &ckBook.PowerLevel)
	if err != nil {
		logInfo.Printf("上传错误：资源不存在：用户：%s，资源：%s", theUser.Name, string(theBook_b))
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	if ckBook.PowerLevel >= theUser.UPower["resource"]["origin"] || ckBook.UnitsId != theUser.UnitId {
		logInfo.Printf("上传错误：用户无权限上传资源文件：用户：%s，资源：%s", theUser.Name, string(theBook_b))
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	// end
	processid, err := globalLock.TryLock(string(theSIDb), string(theBook_b), 2)  //尝试加写锁
	if err != nil {
		logInfo.Printf("上传错误：加锁失败：用户：%s，资源：%s", theUser.Name, string(theBook_b))
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	SendSocketBytes (conn , Uint8ToBytes(1), 1)  //允许上传
	SendSocketBytes (conn, []byte(processid),40)
}
