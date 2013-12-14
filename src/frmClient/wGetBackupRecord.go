package main

import (
	"net/http"
	"fmt"
	. "frmPkg"
)

func wGetBackupRecord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	getCallback, foundGet := r.Form["callback"]
	if  foundGet != true {
		fmt.Fprint(w,"{\"err\":\"不是正确的接口请求\"}")
		return
	}
	callback := getCallback[0]
	
	bak := StructToJson(backupRecord)
	
	theSend := callback + "(" + bak + ")"
	
	fmt.Fprintf(w, theSend)
}
