package main

import (
	"net/http"
	"fmt"
	//. "frmPkg"
)

func wGetBackupRecord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	getCallback, foundGet := r.Form["callback"]
	if  foundGet != true {
		fmt.Fprint(w, "({\"err\":\"不是正确的接口请求\"})")
		return
	}
	callback := getCallback[0]
	
	selfBackup, found := backupRecord.Record[r.Form["userid"][0]]
	if found == false {
		fmt.Fprint(w, callback + "({\"err\":\"没有相应数据\"})")
		return
	}
	
	bak := "{"
	for key, val := range selfBackup {
		bak = bak + "\"" + fmt.Sprint(key) + "\" : \"" + val + "\","
	}
	bak = bak + "}"
	
	theSend := callback + "(" + bak + ")"
	
	fmt.Fprintf(w, theSend)
}
