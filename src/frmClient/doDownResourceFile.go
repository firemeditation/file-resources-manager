package main

import (
	"os"
	. "frmPkg"
	"fmt"
	"time"
	"strings"
)

const DownGoMax = 5  //同时下载的最大进程数

func doDownResourceFile(userid, resourceid, originpath, downtype, files, bookname string) (errA []string, err error) {
	_, err = os.Stat(originpath)
	if err != nil {
		err = fmt.Errorf("找不到文件或目录：%s", originpath)
		bk := fmt.Sprint(err)
		backupRecord.AddRecord(userid, bk)
		return
	}
	//fmt.Println("请求加锁")
	// 开始请求加锁
	conn := connectServer()
	err = sendTheFirstRequest (1, 3, conn)
	if err != nil {
		err = fmt.Errorf("发送状态错误：%s", err)
		return
	}
	//fmt.Println("请求加锁2")
	//发送自己的SID
	//err = SendSocketBytes(conn, []byte(myLogin.SID), 40)
	err = SendSocketBytes(conn, []byte(userid), 40)
	if err != nil {
		err = fmt.Errorf("发送SID错误：%s", err)
		return
	}
	//发送加读锁请求
	err = SendSocketBytes(conn, Uint8ToBytes(2), 1)
	if err != nil {
		err = fmt.Errorf("发送请求状态出错：%s", err)
		return
	}
	//fmt.Println("请求加锁3")
	// 发送请求资源hashid
	err = SendSocketBytes(conn, []byte(resourceid), 40)
	if err != nil {
		err = fmt.Errorf("发送资源hashid错误：%s", err)
		return
	}
	//fmt.Println("请求加锁4")
	// 查看服务器是否允许加锁
	cklb, _ := ReadSocketBytes(conn, 1)
	ckl := BytesToUint8(cklb)
	if ckl == 2 {
		err = fmt.Errorf("不允许加锁：%s", resourceid)
		return
	}
	//fmt.Println("请求加锁5")
	processid_b , _ := ReadSocketBytes(conn,40)
	processid := string(processid_b)  //获取进程ID
	
	//fmt.Println("加锁成功：",processid)
	
	// 获得要下载的文件的hashid并存入一个channel
	fileHashid := make(chan string, DownGoMax)
	if downtype == "one"{
		fileHashid <- files
	}else{
		go getDownloadFilesHashid(fileHashid, downtype, files, errA)
	}
	
	// 开启多个进程同时向服务器传递文件
	//var wg sync.WaitGroup
	
	downDone := make(chan int, DownGoMax)
	
	for i := 0; i < DownGoMax; i++ {
		//wg.Add(1)
		//fmt.Println("进程",i)
		go downOneFile(downDone, userid, resourceid, processid, originpath, fileHashid, errA)
	}
	
	doneNum := 0
	for {
		select {
			case <-downDone :
				doneNum++
				//fmt.Println("进程结束")
			default:
				time.Sleep(3 * time.Second)
		}
		if doneNum == DownGoMax {
			break
		}
		SendSocketBytes(conn, Uint8ToBytes(1), 1)  //发送心跳包
		//fmt.Println("发送心跳")
		ckh_b, _ := ReadSocketBytes(conn, 1)
		//fmt.Println("接收回执")
		if BytesToUint8(ckh_b) != 1 {
			break
		}
	}
	//wg.Wait()
	SendSocketBytes(conn, Uint8ToBytes(2), 1)  //发送关闭
	ReadSocketBytes(conn, 1)
	
	brstring := "后台下载完成：" + bookname
	backupRecord.AddRecord(userid, brstring)
		
	return
}

// getDownloadFilesHashid 获得下载文件的hashid
func getDownloadFilesHashid(fileHashid chan<- string, downtype, files string, errA []string){
	if downtype == "all" {
		//TODO
	}else if downtype == "part"{
		hashidA := strings.Split(files,",")
		for _, oneHash := range hashidA {
			oneHash = strings.TrimSpace(oneHash)
			if len(oneHash) != 0 {
				fileHashid <- oneHash
			}
		}
	}
}

// downOneFile 下载一个文件 //TODO
func downOneFile(downDone chan int, userid, resourceid, processid, originpath string, fileHashid <-chan string, errA []string){
	defer func() {
		downDone <- 1
	}()
	for oneFile := range fileHashid {
		conn := connectServer()
		err := sendTheFirstRequest (1, 7, conn)
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
		err = SendSocketBytes(conn, []byte(oneFile), 40)
		if err != nil {
			errA = append(errA, "发送文件ID错误")
			break
		}
		cklb, _ := ReadSocketBytes(conn, 1)
		ckl := BytesToUint8(cklb)
		if ckl == 2 {
			errS := fmt.Sprintf("服务器不允许下载文件：%s", fileHashid)
			errA = append(errA, errS)
			backupRecord.AddRecord(userid, errS)
			break
		}
		
	}
}
