package s1


import (
	"net"
	"fmt"
	. "frmPkg"
	"time"
	"os"
	"strings"
	. "frmServer/public"
)


func ProcessUploadProcess(conn *net.TCPConn) {
	allid_b, err := ReadSocketBytes(conn, 120) //接收三个ID
	if err != nil {
		LogInfo.Printf("接收错误")
		return
	}
	uid := string(allid_b[:40]) //用户ID
	//fmt.Println(uid)
	resource_id := string(allid_b[40:80]) //资源ID
	//fmt.Println(resource_id)
	process_id := string(allid_b[80:]) //进程ID
	//fmt.Println(process_id)
	
	//检查是否可以上传
	err = GlobalLock.CheckLock(uid, resource_id, process_id, 1)
	if err != nil {
		LogInfo.Printf("检查锁出错：%s", err)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	
	theUser, _ := CkLogedUser (uid)
	
	file_info_len_b, _ := ReadSocketBytes(conn,8)
	file_info_len := BytesToUint64(file_info_len_b)
	file_info_b, _ := ReadSocketBytes(conn, file_info_len)
	var file_info OriginFileInfoStruct
	BytesGobStruct(file_info_b, &file_info)  //获得到了文件的信息结构体
	
	file_len_b, _ := ReadSocketBytes(conn,8)
	file_len := BytesToUint64(file_len_b)  //获得了文件的大小
	
	storage := <-StorageChan
	fileHashName := uid + resource_id + process_id + file_info.RelativeDir + file_info.FileName + time.Now().String()
	fileHashName = GetSha1(fileHashName)
	fileFullStoragePath := storage.Path + storage.SmallPath + fileHashName
	fileSmallStoragePath := storage.SmallPath + fileHashName
	infile, err := os.OpenFile(fileFullStoragePath, os.O_WRONLY|os.O_CREATE, os.FileMode(0600))
	if err != nil {
		fmt.Println("文件建立错误", err)
	}
	err = ReadSocketToFile(conn, file_len, infile)
	if err != nil {
		LogInfo.Printf("文件上传错误：%s", err)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		the_e_b := []byte(fmt.Sprintf("文件上传错误：%s",err))
		the_e_b_len := len(the_e_b)
		SendSocketBytes(conn, Uint64ToBytes(uint64(the_e_b_len)), 8)
		SendSocketBytes(conn, the_e_b, uint64(the_e_b_len))
		return
	}
	
	//begin 写入数据库
	// 根据原始位置，原始文件名，所属聚集来判断之前是否上传过这个文件
	var old_hashid, old_path, old_site string
	old_version := 0
	err = DbConn.QueryRow("select hashid, version, fpath, fsite from resourceFile where opath = $1 and fname = $2 and rg_hashid = $3", file_info.RelativeDir, file_info.FileName, resource_id).Scan(&old_hashid, &old_version, &old_path, &old_site)
	if err == nil {
		n_rgt, err := DbConn.Prepare("delete from resourceFile where hashid = $1")
		if err != nil {
			LogInfo.Printf("数据库错误：%s", err)
		}
		_, err = n_rgt.Exec(old_hashid)
		old_site = strings.Trim(old_site, " ")
		old_path = strings.Trim(old_path, " ")
		old_file := old_site + old_path
		fmt.Println("旧文件名：",old_file)
		err = os.Remove(old_file)
		fmt.Println("删除错误：",err)
	}
	old_version++
	fileInsertInfo := ResourceFileTable{ResourceItemTable:ResourceItemTable{HashId:fileHashName, Name:file_info.FileName, RiType:1, LastTime:time.Now().Unix(), Version: uint16(old_version), RgHashId: resource_id, UnitsId: theUser.UnitId, UsersId: theUser.Id}, Fname: file_info.FileName, Opath: file_info.RelativeDir, Fpath: fileSmallStoragePath, Fsite: storage.Path, Fsize: file_info.Size}
	
	insert_file , err := DbConn.Prepare("insert into resourceFile (hashid, name, ritype, lasttime, version, rg_hashid, units_id, users_id, fname, opath, fpath, fsite, fsize) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)")
	if err != nil {
		fmt.Println("数据库错误1")
	}
	_, err = insert_file.Exec(fileInsertInfo.HashId, fileInsertInfo.Name, fileInsertInfo.RiType, fileInsertInfo.LastTime, fileInsertInfo.Version, fileInsertInfo.RgHashId, fileInsertInfo.UnitsId, fileInsertInfo.UsersId, fileInsertInfo.Fname, fileInsertInfo.Opath, fileInsertInfo.Fpath, fileInsertInfo.Fsite, fileInsertInfo.Fsize)
	if err != nil {
		fmt.Println("数据库错误2")
	}
	
	n_rgstatus, err := DbConn.Prepare("insert into resourceFileStatus (hashid) values ($1)")
	if err != nil {
		ErrLog.Println("数据库错误：", err)
		fmt.Println("数据库错误3")
	}
	_, err = n_rgstatus.Exec(fileInsertInfo.HashId)
	if err != nil {
		ErrLog.Println("数据库错误：", err)
		fmt.Println("数据库错误4")
	}
	
	LogInfo.Printf("文件上传成功：%s%s", file_info.RelativeDir, file_info.FileName)
	fmt.Printf("文件上传成功：%s%s，位置：%s", file_info.RelativeDir, file_info.FileName, fileFullStoragePath )
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	return
}
