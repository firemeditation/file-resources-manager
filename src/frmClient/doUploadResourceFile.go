package main

import (
	"os"
	. "frmPkg"
	"fmt"
	"time"
)

const UploadGoMax = 5  //同时上传的最大进程数

func doUploadResourceFile(userid, resourceid string, originpath, addtopath, bookname string) (errA []string, err error) {
	ckdir, err := os.Stat(originpath)
	if err != nil {
		err = fmt.Errorf("找不到文件或目录：%s", originpath)
		bk := fmt.Sprint(err)
		backupRecord.AddRecord(userid, bk)
		return
	}
	fmt.Println("请求加锁")
	// 开始请求加锁
	conn := connectServer()
	err = sendTheFirstRequest (1, 3, conn)
	if err != nil {
		err = fmt.Errorf("发送状态错误：%s", err)
		return
	}
	fmt.Println("请求加锁2")
	//发送自己的SID
	//err = SendSocketBytes(conn, []byte(myLogin.SID), 40)
	err = SendSocketBytes(conn, []byte(userid), 40)
	if err != nil {
		err = fmt.Errorf("发送SID错误：%s", err)
		return
	}
	//发送1
	err = SendSocketBytes(conn, Uint8ToBytes(1), 1)
	if err != nil {
		err = fmt.Errorf("发送请求状态出错：%s", err)
		return
	}
	fmt.Println("请求加锁3")
	// 发送请求资源hashid
	err = SendSocketBytes(conn, []byte(resourceid), 40)
	if err != nil {
		err = fmt.Errorf("发送资源hashid错误：%s", err)
		return
	}
	fmt.Println("请求加锁4")
	// 查看服务器是否允许加锁
	cklb, _ := ReadSocketBytes(conn, 1)
	ckl := BytesToUint8(cklb)
	if ckl == 2 {
		err = fmt.Errorf("不允许加锁：%s", resourceid)
		return
	}
	fmt.Println("请求加锁5")
	processid_b , _ := ReadSocketBytes(conn,40)
	processid := string(processid_b)  //获取进程ID
	
	fmt.Println("加锁成功：",processid)
	
	// 遍历要找到的目录并存入一个channel
	fileInfo := make(chan OriginFileInfoFullStruct, UploadGoMax)
	if ckdir.IsDir() == true {
		go readDirMain(fileInfo, originpath, addtopath)
	}else{
		fileInfo <- OriginFileInfoFullStruct{originpath, OriginFileInfoStruct{addtopath, ckdir.Name(), ckdir.Size(),ckdir.Mode(),ckdir.ModTime()}}
	}
	
	// 开启多个进程同时向服务器传递文件
	//var wg sync.WaitGroup
	
	uploadDone := make(chan int, UploadGoMax)
	
	for i := 0; i < UploadGoMax; i++ {
		//wg.Add(1)
		fmt.Println("进程",i)
		go sendFiles(uploadDone, userid, resourceid, processid, fileInfo, errA)
	}
	
	doneNum := 0
	for {
		select {
			case <-uploadDone :
				doneNum++
				fmt.Println("进程结束")
			default:
				time.Sleep(3 * time.Second)
		}
		if doneNum == UploadGoMax {
			break
		}
		SendSocketBytes(conn, Uint8ToBytes(1), 1)  //发送心跳包
		fmt.Println("发送心跳")
		ckh_b, _ := ReadSocketBytes(conn, 1)
		fmt.Println("接收回执")
		if BytesToUint8(ckh_b) != 1 {
			break
		}
	}
	//wg.Wait()
	SendSocketBytes(conn, Uint8ToBytes(2), 1)  //发送关闭
	ReadSocketBytes(conn, 1)
	
	brstring := "后台上传完成：" + bookname
	backupRecord.AddRecord(userid, brstring)
		
	return
}


// readDirMain 读取文件列表
func readDirMain(fileInfo chan<- OriginFileInfoFullStruct, originpath, addtopath string){
	readDir(fileInfo, originpath, addtopath)
	defer close(fileInfo)
}

// readDir 实际的读取文件列表的函数
func readDir(fileInfo chan<- OriginFileInfoFullStruct, dir, relative string){
	opendir, _ := os.Open(dir)
	defer opendir.Close()
	allfile , _ := opendir.Readdir(0)
	for _, onefile := range allfile {
		if onefile.IsDir() == true {
			dirName := dir + onefile.Name() + "/"
			relativeName := relative + onefile.Name() + "/"
			readDir(fileInfo, dirName, relativeName)
		}else {
			dirName := dir + onefile.Name()
			fileInfo <- OriginFileInfoFullStruct{dirName, OriginFileInfoStruct{relative, onefile.Name(), onefile.Size(),onefile.Mode(),onefile.ModTime()}}
		}
	}
}

// sendFiles 发送文件
func sendFiles(uploadDone chan int, userid, resourceid, processid string, fileInfo <-chan OriginFileInfoFullStruct, errA []string){
	defer func() {
		uploadDone <- 1
	}()
	for oneFile := range fileInfo {
		conn := connectServer()
		err := sendTheFirstRequest (1, 4, conn)
		if err != nil {
			errA = append(errA, "发送状态错误")
			break
		}
		//err = SendSocketBytes(conn, []byte(myLogin.SID), 40)
		err = SendSocketBytes(conn, []byte(userid), 40)
		if err != nil {
			errA = append(errA, "发送SID错误")
			break
		}
		err = SendSocketBytes(conn, []byte(resourceid), 40)
		if err != nil {
			errA = append(errA, "发送资源ID错误")
			break
		}
		err = SendSocketBytes(conn, []byte(processid), 40)
		if err != nil {
			errA = append(errA, "发送进程ID错误")
			break
		}
		cklb, _ := ReadSocketBytes(conn, 1)
		ckl := BytesToUint8(cklb)
		if ckl == 2 {
			errS := fmt.Sprintf("服务器不允许上传文件：%s", oneFile.FullDir)
			errA = append(errA, errS)
			backupRecord.AddRecord(userid, errS)
			break
		}
		
		ofis_byte := StructGobBytes(oneFile.OriginFileInfoStruct)
		// 发送文件信息的结构体长度
		ofis_len := len(ofis_byte)
		ofis_len_b := Uint64ToBytes(uint64(ofis_len))
		err = SendSocketBytes(conn, ofis_len_b, 8)
		if err != nil {
			errS := fmt.Sprintf("上传文件出错：%s，错误：%s", oneFile.FullDir, err)
			errA = append(errA, errS)
			backupRecord.AddRecord(userid, errS)
			break
		}
		// 发送文件信息的结构体
		err = SendSocketBytes(conn, ofis_byte, uint64(ofis_len))
		if err != nil {
			errS := fmt.Sprintf("上传文件出错：%s，错误：%s", oneFile.FullDir, err)
			errA = append(errA, errS)
			backupRecord.AddRecord(userid, errS)
			break
		}
		// 发送文件数据长度
		file_len := oneFile.Size
		file_len_byte := Uint64ToBytes(uint64(file_len))
		err = SendSocketBytes(conn, file_len_byte, 8)
		if err != nil {
			errS := fmt.Sprintf("上传文件出错：%s，错误：%s", oneFile.FullDir, err)
			errA = append(errA, errS)
			backupRecord.AddRecord(userid, errS)
			break
		}
		err = SendSocketFile (conn, uint64(file_len), oneFile.FullDir)
		if err != nil {
			errS := fmt.Sprintf("上传文件出错：%s，错误：%s", oneFile.FullDir, err)
			errA = append(errA, errS)
			backupRecord.AddRecord(userid, errS)
			break
		}
		// 接收服务器确认
		ckqrb, _ := ReadSocketBytes(conn, 1)
		ckqr := BytesToUint8(ckqrb)
		if ckqr != 1 {
			geterr_len_b , _ := ReadSocketBytes(conn, 8)
			geterr_len := BytesToUint64(geterr_len_b)
			geterr_b, _ := ReadSocketBytes(conn, geterr_len)
			errS := fmt.Sprintf("上传文件出错：%s，错误：%s", oneFile.FullDir, string(geterr_b))
			errA = append(errA, errS)
			backupRecord.AddRecord(userid, errS)
			break
		}
		fmt.Println("上传完成：",oneFile.FullDir) 
	}
}
