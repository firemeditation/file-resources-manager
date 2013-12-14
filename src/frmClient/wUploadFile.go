package main

import (
	"net/http"
	"fmt"
	. "frmPkg"
	"time"
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
	
	localpath = DirMustEnd(localpath)
	if len(relative) != 0 {
		relative = DirMustEnd(relative)
	}
	
	theSend := callback + "({\"client\":\"yes\"})"
	fmt.Fprintf(w, theSend)
	
	brstring := time.Now().String() +  "：后台上传中，图书hashid：" + hashid
	backupRecord = append(backupRecord,brstring)
	
	_ , err := doUploadResourceFile(user, hashid, localpath, relative)
	brstring = time.Now().String() +  "：后台上传出错，图书hashid：" + hashid + "：错误" + err
	backupRecord = append(backupRecord,brstring)
	
	brstring = time.Now().String() +  "：后台上传完成，图书hashid：" + hashid
	backupRecord = append(backupRecord,brstring)
}
