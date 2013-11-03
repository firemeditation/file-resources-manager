package main

import (
	"os"
	. "frmPkg"
	"fmt"
	"sync"
)

const UploadGoMax = 4  //同时上传的最大进程数

func doUploadResourceFile(uploadtype int, originpath, addtopath string) (errA []string, err error) {
	ckdir, err := os.Stat(originpath)
	if err != nil {
		err = fmt.Errorf("找不到文件或目录：%s", ckdir)
		return
	}
	
	// 遍历要找到的目录并存入一个channel
	fileInfo := make(chan OriginFileInfoFullStruct, UploadGoMax)
	if ckdir.IsDir() == true {
		go readDirMain(fileInfo, originpath, addtopath)
	}else{
		fileInfo <- OriginFileInfoFullStruct{originpath, OriginFileInfoStruct{addtopath, ckdir.Name(), ckdir.Size(),ckdir.Mode(),ckdir.ModTime()}}
	}
	
	// 开启多个进程同时向服务器传递文件
	var wg sync.WaitGroup
	
	for i := 0; i < UploadGoMax; i++ {
		wg.Add(1)
		go sendFiles(uploadtype, fileInfo, &wg, errA)
	}
	
	wg.Wait()
	
	return
}


func readDirMain(fileInfo chan<- OriginFileInfoFullStruct, originpath, addtopath string){
	readDir(fileInfo, originpath, addtopath)
	defer close(fileInfo)
}

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

func sendFiles(uploadtype int, fileInfo <-chan OriginFileInfoFullStruct, wg *sync.WaitGroup, errA []string){
	defer wg.Done()
	for oneFile := range fileInfo {
		fmt.Println(oneFile.FullDir)
	}
}
