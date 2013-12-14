package main

import (
	"net/http"
	"fmt"
)

func wCheckLink(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	getCallback, foundGet := r.Form["callback"]
	if  foundGet != true {
		fmt.Fprint(w,"{\"err\":\"不是正确的接口请求\"}")
		return
	}
	callback := getCallback[0]
	
	theSend := callback + "({\"client\":\"yes\", \"backupNum\":" + fmt.Sprint(backupNum) + "})"
	
	fmt.Fprintf(w, theSend)
}
