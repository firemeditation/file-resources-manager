package main

import (
	"net/http"
	"fmt"
	. "frmPkg"
)

func wUploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	getCallback, foundGet := r.Form["callback"]
	if  foundGet != true {
		fmt.Fprint(w,"{\"err\":\"不是正确的接口请求\"}")
		return
	}
	callback := getCallback[0]
	
	localpath := r.Form["local"][0]
	relative := r.Form["relative"][0]
	hashid := r.Form["hashid"][0]
	user := r.Form["user"][0]
	bookname := r.Form["bookname"][0]
	
	localpath = DirMustEnd(localpath)
	if len(relative) != 0 {
		relative = DirMustEnd(relative)
	}
	
	theSend := callback + "({\"client\":\"yes\"})"
	fmt.Fprintf(w, theSend)
	
	brstring := "后台上传中：" + bookname
	backupRecord.AddRecord(user, brstring)
	//backupNum = backupNum + 1
	//defer func(){
	//	backupNum = backupNum - 1
	//}()
	
	_ , err := doUploadResourceFile(user, hashid, localpath, relative)
	if err != nil {
		brstring = "后台上传出错：" + bookname + "：错误：" + fmt.Sprint(err)
		backupRecord.AddRecord(user, brstring)
	}
	
	brstring = "后台上传完成：" + bookname
	backupRecord.AddRecord(user, brstring)
}
