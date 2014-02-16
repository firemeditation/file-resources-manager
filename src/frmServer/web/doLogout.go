package web

import (
	//. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	//"regexp"
	//"os"
)


func doLogout(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	//fmt.Println("logout")
	cookiename, _ := ServerConfig.GetString("web","cookie")
	userhash_cookie, err := r.Cookie(cookiename)
	if err != nil {
		fmt.Fprint(w,"no")
		//fmt.Println("no")
		return
	}
	
	//cookielife, _ := ServerConfig.GetInt64("user","timeout")
	cookie := http.Cookie{Name: cookiename, Value: userhash_cookie.Value, MaxAge: -2}
	http.SetCookie(w, &cookie)
	UserLoginStatus.Del(userhash_cookie.Value)
	fmt.Fprint(w,"yes")
	//fmt.Println("yes")
}
