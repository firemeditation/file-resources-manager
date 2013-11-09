package main

import (
	"os"
	. "frmPkg"
	"fmt"
	"time"
)

const UploadGoMax = 4  //同时上传的最大进程数

func doUploadResourceFile(resourceid string, originpath, addtopath string) (errA []string, err error) {
	ckdir, err := os.Stat(originpath)
	if err != nil {
		err = fmt.Errorf("找不到文件或目录：%s", ckdir)
		return
	}
	
	// 开始请求加锁
	conn := connectServer()
	err = sendTheFirstRequest (1, 3, conn)
	if err != nil {
		err = fmt.Errorf("发送状态错误：%s", err)
		return
	}
	//发送自己的SID
	err = SendSocketBytes(conn, []byte(myLogin.SID), 40)
	if err != nil {
		err = fmt.Errorf("发送SID错误：%s", err)
		return
	}
	// 发送请求资源hashid
	err = SendSocketBytes(conn, []byte(resourceid), 40)
	if err != nil {
		err = fmt.Errorf("发送资源hashid错误：%s", err)
		return
	}
	// 查看服务器是否允许加锁
	cklb, _ := ReadSocketBytes(conn, 1)
	ckl := BytesToUint8(cklb)
	if ckl == 2 {
		err = fmt.Errorf("不允许加锁：%s", resourceid)
		return
	}
	processid_b , _ := ReadSocketBytes(conn,40)
	processid := string(processid_b)  //获取进程ID
	
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
		go sendFiles(uploadDone,resourceid, processid, fileInfo, errA)
	}
	
	doneNum := 0
	for {
		select {
			case <-uploadDone :
				doneNum++
			default:
				break
		}
		if doneNum == UploadGoMax {
			break
		}
		SendSocketBytes(conn, Uint8ToBytes(1), 1)  //发送心跳包
		time.Sleep(1 * time.Second)
	}
	//wg.Wait()
	SendSocketBytes(conn, Uint8ToBytes(1), 2)  //发送关闭
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
func sendFiles(uploadDone chan int,resourceid, processid string, fileInfo <-chan OriginFileInfoFullStruct, errA []string){
	defer func() {
		uploadDone <- 1
	}()
	for oneFile := range fileInfo {
		//TODO
		fmt.Println(oneFile.FullDir) 
	}
}
