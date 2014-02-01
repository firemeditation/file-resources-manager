package main

import (
	"net/http"
	"fmt"
	"os"
	. "frmPkg"
)

func wDownloadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	getCallback, foundGet := r.Form["callback"]
	if  foundGet != true {
		fmt.Fprint(w,"{\"err\":\"不是正确的接口请求\"}")
		return
	}
	callback := getCallback[0]
	
	localpath := r.Form["localpath"][0]
	hashid := r.Form["hashid"][0]
	user := r.Form["user"][0]
	bookname := r.Form["bookname"][0]
	downtype := r.Form["type"][0]
	
	localpath = DirMustEnd(localpath)
	
	var files string
	if downtype != "all" {
		files = r.Form["files"][0]
	}
	
	_, err := os.Stat(localpath)
	if err != nil {
		theSend := callback + "({\"err\":\"找不到要保存的目录\"})"
		fmt.Fprint(w, theSend)
		return
	}
	
	theTruePath := localpath + bookname + "_" + hashid + "/"
	if FileExist(theTruePath) == false {
		err = os.Mkdir(theTruePath,0755)
		if err != nil {
			theSend := callback + "({\"err\":\"无法创建目录，请检查本地路径权限\"})"
			fmt.Fprint(w, theSend)
			return
		}
	}
	
	theSend := callback + "({\"client\":\"yes\"})"
	fmt.Fprintf(w, theSend)
	
	return
	
	brstring := "后台下载中：" + bookname
	backupRecord.AddRecord(user, brstring)
	//backupNum = backupNum + 1
	//defer func(){
	//	backupNum = backupNum - 1
	//}()
	
	go doDownResourceFile(user, hashid, theTruePath, downtype, files, bookname)
	
	/*_ , err := doUploadResourceFile(user, hashid, localpath, relative)
	if err != nil {
		brstring = "后台上传出错：" + bookname + "：错误：" + fmt.Sprint(err)
		backupRecord.AddRecord(user, brstring)
	}
	
	brstring = "后台上传完成：" + bookname
	backupRecord.AddRecord(user, brstring)
	*/
}
