package web

import (
	//. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	//"regexp"
	//"os"
)


func webInterface(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	cookiename, _ := ServerConfig.GetString("web","cookie")
	userhash_cookie, err := r.Cookie(cookiename)
	if err != nil {
		fmt.Fprint(w,"{\"err\":\"用户不存在\"}")
		return
	}
	
	// begin 查看用户是否存在或超时
	theUser, found  := CkLogedUser (userhash_cookie.Value)
	if found == false {
		fmt.Fprint(w,"{\"err\":\"用户超时\"}")
		return
	}
	
	getI, foundGet := r.Form["type"]	
	if  foundGet != true {
		fmt.Fprint(w,"{\"err\":\"不是正确的接口请求\"}")
		return
	}
	switch getI[0] {
		case "get-basic-user-info":
			iGetBasicUserInfo(theUser, w, r)
		default:
			fmt.Fprint(w,"{\"err\":\"请求的接口不存在\"}")
	}
}
