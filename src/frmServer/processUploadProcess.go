package main


import (
	"net"
	"fmt"
	. "frmPkg"
	"time"
	"os"
)


func processUploadProcess(conn *net.TCPConn) {
	allid_b, err := ReadSocketBytes(conn, 120) //接收三个ID
	if err != nil {
		logInfo.Printf("接收错误")
		return
	}
	uid := string(allid_b[:40]) //用户ID
	//fmt.Println(uid)
	resource_id := string(allid_b[40:80]) //资源ID
	//fmt.Println(resource_id)
	process_id := string(allid_b[80:]) //进程ID
	//fmt.Println(process_id)
	
	//检查是否可以上传
	err = globalLock.CheckLock(uid, resource_id, process_id, 1)
	if err != nil {
		logInfo.Printf("检查锁出错：%s", err)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		return
	}
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	
	file_info_len_b, _ := ReadSocketBytes(conn,8)
	file_info_len := BytesToUint64(file_info_len_b)
	file_info_b, _ := ReadSocketBytes(conn, file_info_len)
	var file_info OriginFileInfoStruct
	BytesGobStruct(file_info_b, &file_info)  //获得到了文件的信息结构体
	
	file_len_b, _ := ReadSocketBytes(conn,8)
	file_len := BytesToUint64(file_len_b)  //获得了文件的大小
	
	storage := <-storageChan
	fileHashName := uid + resource_id + process_id + file_info.RelativeDir + file_info.FileName + time.Now().String()
	fileHashName = GetSha1(fileHashName)
	fileFullStoragePath := storage.Path + storage.SmallPath + fileHashName
	
	infile, _ := os.OpenFile(fileFullStoragePath, os.O_WRONLY|os.O_CREATE, os.FileMode(0600))
	err = ReadSocketToFile(conn, file_len, infile)
	if err != nil {
		logInfo.Printf("文件上传错误：%s", err)
		SendSocketBytes (conn , Uint8ToBytes(2), 1)
		the_e_b := []byte(fmt.Sprintf("文件上传错误：%s",err))
		the_e_b_len := len(the_e_b)
		SendSocketBytes(conn, Uint64ToBytes(uint64(the_e_b_len)), 8)
		SendSocketBytes(conn, the_e_b, uint64(the_e_b_len))
		return
	}
	logInfo.Printf("文件上传成功：%s%s", file_info.RelativeDir, file_info.FileName)
	fmt.Printf("文件上传成功：%s%s，位置：%s", file_info.RelativeDir, file_info.FileName, fileFullStoragePath )
	SendSocketBytes (conn , Uint8ToBytes(1), 1)
	return
}
