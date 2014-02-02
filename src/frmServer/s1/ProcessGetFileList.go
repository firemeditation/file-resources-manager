package s1


import (
	"net"
	. "frmPkg"
	. "frmServer/public"
)


func ProcessGetFileList(conn *net.TCPConn) {
	allid_b, err := ReadSocketBytes(conn, 120) //接收三个ID
	if err != nil {
		ErrLog.Printf("接收错误")
		return
	}
	uid := string(allid_b[:40]) //用户ID
	//fmt.Println(uid)
	resource_id := string(allid_b[40:80]) //资源ID
	//fmt.Println(resource_id)
	process_id := string(allid_b[80:120]) //进程ID
	//fmt.Println(process_id)
	file_list_type_b, _ := ReadSocketBytes(conn, 1)
	file_list_type := BytesToUint8(file_list_type_b)
	
	//检查锁
	err = GlobalLock.CheckLock(uid, resource_id, process_id, 2)
	if err != nil {
		ErrLog.Printf("检查锁出错：%s", err)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	
	/*theUser, isok := CkLogedUser(uid)
	if isok == false {
		ErrLog.Printf("用户超时")
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}*/
	
	SendSocketBytes (conn , Uint8ToBytes(1), 1)  //发送一切正常
	
	var hashids string
	//开始获取资源文件的hashid
	if file_list_type == 1 {
		allfile, _ := DbConn.Query("select hashid from resourceFile where rg_hashid = $1",resource_id)
		for allfile.Next() {
			var one_hashid string
			allfile.Scan(&one_hashid)
			hashids = hashids + one_hashid + ","
		}
	}
	
	hashids_b := []byte(hashids)
	thelen := len(hashids_b)
	SendSocketBytes (conn , Uint64ToBytes(uint64(thelen)), 8)
	SendSocketBytes (conn , hashids_b, uint64(thelen))
	
	return
}
