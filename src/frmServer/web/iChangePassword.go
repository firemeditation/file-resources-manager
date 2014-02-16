package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	"strings"
)


func iChangePassword(theUser *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	
	oldpassword := r.PostFormValue("old")
	newpassword := r.PostFormValue("news")
	
	var ck_passwd string
	DbConn.QueryRow("select passwd from users where id = $1",theUser.Id).Scan(&ck_passwd)
	ck_passwd = strings.TrimSpace(ck_passwd)
	if ck_passwd != oldpassword {
		fmt.Fprint(w,"{\"err\":\"原始密码不正确\"}")
		return
	}
	
	ch_pre, _ := DbConn.Prepare("update users set passwd = $1 where id = $2")
	_, err := ch_pre.Exec(newpassword, theUser.Id)
	if err != nil {
		fmt.Fprint(w,"{\"err\":\"密码修改不成功\"}")
		return
	}
	
	fmt.Fprint(w, "{\"ok\":\"ok\"}")
}
