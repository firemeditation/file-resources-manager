package s1


import (
	"net"
	. "frmPkg"
	. "frmServer/public"
	"strings"
)


func ProcessDownloadFile(conn *net.TCPConn) {
	allid_b, err := ReadSocketBytes(conn, 160) //接收四个ID
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
	file_id := string(allid_b[120:])  //文件的ID
	
	//检查是否有锁
	err = GlobalLock.CheckLock(uid, resource_id, process_id, 2)
	if err != nil {
		ErrLog.Printf("检查锁出错：%s", err)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	
	//获取文件的信息
	var file_num int
	DbConn.QueryRow("select COUNT(*) from resourcefile where hashid = $1",file_id).Scan(&file_num)
	if file_num != 1 {
		ErrLog.Printf("文件检索错误：%s", err)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	fileInfo := ResourceFileTable{ResourceItemTable:ResourceItemTable{}}
	DbConn.QueryRow("select rg_hashid, fname, opath, fpath, fsite, fsize from resourceFile where hashid = $1",file_id).Scan(&fileInfo.ResourceItemTable.RgHashId, &fileInfo.Fname, &fileInfo.Opath, &fileInfo.Fpath, &fileInfo.Fsite, &fileInfo.Fsize)
	
	fileInfo.Fname = strings.TrimSpace(fileInfo.Fname)
	fileInfo.Opath = strings.TrimSpace(fileInfo.Opath)
	fileInfo.Fpath = strings.TrimSpace(fileInfo.Fpath)
	fileInfo.Fsite = strings.TrimSpace(fileInfo.Fsite)
	
	if fileInfo.ResourceItemTable.RgHashId != resource_id {
		ErrLog.Printf("文件检索错误：%s", err)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	
	// 可以下载了
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	
	ofis := OriginFileInfoStruct{RelativeDir : fileInfo.Opath, FileName : fileInfo.Fname, Size : fileInfo.Fsize}
	ofis_b := StructGobBytes(ofis)
	ofis_b_len := len(ofis_b)
	ofis_b_len_b := Uint64ToBytes(uint64(ofis_b_len))
	file_len_b := Uint64ToBytes(uint64(fileInfo.Fsize))
	
	SendSocketBytes(conn, ofis_b_len_b, 8)
	SendSocketBytes(conn, ofis_b, uint64(ofis_b_len))
	SendSocketBytes(conn, file_len_b, 8)
	
	fileFullPathName := fileInfo.Fsite + fileInfo.Fpath
	SendSocketFile (conn, uint64(fileInfo.Fsize), fileFullPathName)
	
	return
}
