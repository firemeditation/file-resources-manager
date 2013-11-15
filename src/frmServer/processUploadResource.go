package main


import (
	"net"
	. "frmPkg"
	"fmt"
)


func processUploadResource(conn *net.TCPConn) {
	theSIDb, _ := ReadSocketBytes(conn, 40) //用户ID
	fmt.Println("请求加锁-1")
	ReadSocketBytes(conn,1)  //用户请求（在这里忽略）
	fmt.Println("请求加锁-2")
	theBook_b , _ := ReadSocketBytes(conn,40)  //图书ID
	fmt.Println("请求加锁")
	// begin 查看用户是否存在或超时
	theUser, found  := ckLogedUser (string(theSIDb))
	if found == false {
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	fmt.Println("请求加锁1")
	// end
	// start 查看用户是否有资源管理的权力
	if theUser.UPower["resource"]["origin"] < 2 {
		logInfo.Printf("上传错误：用户无权限上传资源文件：用户：%s，资源：%s", theUser.Name, string(theBook_b))
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	fmt.Println("请求加锁2")
	// end
	// begin 查看这个资源是不是用户可写的
	var ckBook ResourceGroupTable
	err := dbConn.QueryRow("select units_id, powerlevel from resourceGroup where hashid = $1", string(theBook_b)).Scan(&ckBook.UnitsId, &ckBook.PowerLevel)
	if err != nil {
		logInfo.Printf("上传错误：资源不存在：用户：%s，资源：%s", theUser.Name, string(theBook_b))
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	fmt.Println("请求加锁3")
	if ckBook.PowerLevel >= theUser.UPower["resource"]["origin"] || ckBook.UnitsId != theUser.UnitId {
		logInfo.Printf("上传错误：用户无权限上传资源文件：用户：%s，资源：%s", theUser.Name, string(theBook_b))
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	fmt.Println("请求加锁4")
	// end
	processid, err := globalLock.TryLock(string(theSIDb), string(theBook_b), 1)  //尝试加写锁
	if err != nil {
		logInfo.Printf("上传错误：加锁失败：用户：%s，资源：%s", theUser.Name, string(theBook_b))
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	
	fmt.Println("加锁成功：",processid)
	
	SendSocketBytes (conn , Uint8ToBytes(1), 1)  //允许上传
	SendSocketBytes (conn, []byte(processid),40)
	
	// 接收心跳包
	to := 0  //超时记录
	for {
		theH_b, err := ReadSocketBytes(conn, 1)
		if err != nil {
			to++
			fmt.Println("心跳错误")
			continue
		}
		if to >= 10 {
			fmt.Println("心跳超时")
			break
		}
		fmt.Println("读到心跳")
		theH := BytesToUint8(theH_b)
		if theH == 1 {
			globalLock.Uptime(string(theBook_b), processid)
			SendSocketBytes (conn , Uint8ToBytes(1), 1)
			fmt.Println("发送回执")
			// 更新用户的最后操作时间
			theUser.UpdateLastTime()
		}else{
			fmt.Println("客户端关闭")
			break
		}
	}
	fmt.Println("锁关闭")
	globalLock.Unlock(string(theBook_b), processid)
	SendSocketBytes (conn , Uint8ToBytes(1), 2)
}
