package web

import (
	//. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	//"regexp"
	//"os"
)


func updateLive(w http.ResponseWriter, r *http.Request){
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
	
	// 更新用户的最后操作时间
	theUser.UpdateLastTime()
	
	cookielife, _ := ServerConfig.GetInt64("user","timeout")
	cookie := http.Cookie{Name: cookiename, Value: userhash_cookie.Value, MaxAge: int(cookielife)}
	http.SetCookie(w, &cookie)
	fmt.Fprint(w,"yes")
}
