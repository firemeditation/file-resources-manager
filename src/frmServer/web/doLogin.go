package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func doLogin(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	if _,found := r.Form["username"]; found == true {
		processLogin(w, r)
		return
	}
	t_file := GlobalRelativePath + staticPath + "login.htm"
	http.ServeFile(w, r, t_file)
}

func processLogin(w http.ResponseWriter, r *http.Request){
	// 禁止nobody登录
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	check := r.Form["check"][0]
	nobody, _ := ServerConfig.GetString("user","nobody")
	if username == nobody {
		LogInfo.Printf("登录错误：用户不存在：用户：%s", username)
		fmt.Fprint(w,"此用户禁止登录")
		return
	}
	// 开始检查用户名和密码，并将需要返回的东西全部返回
	var cku UsersTable
	err := DbConn.QueryRow("select id, passwd,  units_id, groups_id, powerlevel from users where name = $1", username).Scan(&cku.Id, &cku.Passwd, &cku.UnitsId, &cku.GroupsId, &cku.PowerLevel)
	if err != nil {
		LogInfo.Printf("登录错误：用户不存在：用户：%s", username)
		fmt.Fprint(w,"用户名不存在")
		return
	}
	
	ck_passwd := cku.Passwd + check
	ck_passwd = GetSha1(ck_passwd)
	
	if password != ck_passwd {
		LogInfo.Printf("登录错误：密码错误：用户：%s", username)
		fmt.Fprint(w,"用户名或密码错误")
		return
	}
	
	//开始合并权限
	var ckuu UnitsTable  //获取所在Unit的名称和权限
	DbConn.QueryRow("select name, powerlevel from units where id = $1", cku.UnitsId).Scan(&ckuu.Name, &ckuu.PowerLevel)
	
	var ckug GroupsTable  //获取所在Group的权限
	DbConn.QueryRow("select name, powerlevel from groups where id = $1", cku.GroupsId).Scan(&ckug.Name , &ckug.PowerLevel)
	
	var cku_p, ckuu_p, ckug_p UserPower
	JsonToStruct(cku.PowerLevel, &cku_p)
	JsonToStruct(ckuu.PowerLevel, &ckuu_p)
	JsonToStruct(ckug.PowerLevel, &ckug_p)
	allpower := MergePower(cku_p, ckuu_p, ckug_p)
	
	ckuu.Name = strings.TrimSpace(ckuu.Name)
	
	ckug.Name = strings.TrimSpace(ckug.Name)
	
	//开始生成SelfLoginInfo和UserIsLogin
	sha1 := GetSha1(check + username)
	thisU, _ := UserLoginStatus.Add(sha1, sha1, cku.Id, username, cku.GroupsId, ckug.Name, cku.UnitsId, ckuu.Name, time.Now())
	thisU.UPower = allpower
	
	//开始设置Cookie
	cookiename, _ := ServerConfig.GetString("web","cookie")
	cookielife, _ := ServerConfig.GetInt64("user","timeout")
	cookie := http.Cookie{Name: cookiename, Value: sha1, MaxAge: int(cookielife)}
	http.SetCookie(w, &cookie)
	fmt.Fprint(w,"ok")
}
