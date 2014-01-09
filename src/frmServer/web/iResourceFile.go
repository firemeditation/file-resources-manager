package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	//"time"
	//"strconv"
	"strings"
	//"regexp"
	//"os"
	"encoding/json"
)

func iResourceFile(theUser *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	
	// start 查看用户是否有查看资源的权力
	if theUser.UPower["resource"]["origin"] < 1 {
		fmt.Fprint(w,"{\"err\":\"无查看权限\"}")
		return
	}
	
	rhashid := r.Form["hashid"][0]  //资源聚集的HashID
	
	// 查看这个资源聚集是否此人可以访问（目前来说就是属于本机构的）
	var r_uid uint16
	DbConn.QueryRow("select units_id from resourceGroup where hashid = $1", rhashid).Scan(&r_uid)
	if r_uid != theUser.GroupId {
		fmt.Fprint(w,"{\"err\":\"无查看权限\"}")
		return
	}
	
	// 获取到所有属于这个资源的文件（目前来说就只有直接资源）
	all_file_tree := map[string]ResourceFileTreeStruct{}
	all_file , _ := DbConn.Query("select hashid, powerlevel, fname, opath from resourceFile where rg_hashid = $1", rhashid)
	// 整理路径并存入ResourceFileTreeStruct的结构体中
	for all_file.Next(){
		onefile := ResourceFileTable{}
		all_file.Scan(&onefile.HashId, &onefile.PowerLevel, &onefile.Fname, &onefile.Opath)
		pathA := strings.Split(onefile.Opath, "/")
		pathA = pathA[:len(pathA)-1]
		ResourceFileToTree(all_file_tree, pathA, onefile.Fname, onefile.HashId)
	}
	
	// 转成JSON
	tree_in_json , _ := json.Marshal(all_file_tree)
	
	// 发送给浏览器
	fmt.Fprint(w, string(tree_in_json))
}
