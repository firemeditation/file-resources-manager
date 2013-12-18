package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	//"time"
	"strconv"
	"strings"
	//"regexp"
	//"os"
)

type returnResourceListStruct struct {
	Count uint64
	List []ResourceGroupTable
	Meta []ResourceGroupTable_MetaData
}


func iResourceList(theUser *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	
	// start 查看用户是否有建立资源的权力
	//if theUser.UPower["resource"]["origin"] < 2 {
	//	fmt.Fprint(w,"{\"err\":\"无添加权限\"}")
	//	return
	//}
	
	
	model := 1 //1为全部列出，2为搜索
	var search string
	limit := 10
	from := 0
	//begin 整理可能的get
	if _, found := r.Form["search"]; found == true {
		model = 2
		search = r.Form["search"][0]
	}
	if _, found := r.Form["limit"]; found == true {
		limit, _ = strconv.Atoi(r.Form["limit"][0])
	}
	if _, found := r.Form["from"]; found == true {
		from, _ = strconv.Atoi(r.Form["from"][0])
	}
	//end 整理可能的get
	
	
	
	rrls := returnResourceListStruct{0, []ResourceGroupTable{}, []ResourceGroupTable_MetaData{}}
	
	if model == 1 {
		DbConn.QueryRow("select COUNT(*) FROM resourceGroup WHERE units_id = $1", theUser.GroupId).Scan(&rrls.Count)
		if rrls.Count != 0 {
			rrls_rows, _ := DbConn.Query("select hashid, name, rt_id, info, btime, users_id, metadata From resourceGroup WHERE units_id = $1 ORDER BY btime DESC LIMIT $2 OFFSET $3", theUser.GroupId, limit, from)
			for rrls_rows.Next(){
				var rgt ResourceGroupTable
				rrls_rows.Scan(&rgt.HashId, &rgt.Name, &rgt.RtId, &rgt.Info, &rgt.Btime, &rgt.UsersId, &rgt.MetaData)
				rgt.Name = strings.TrimSpace(rgt.Name)
				var rgt_md ResourceGroupTable_MetaData
				JsonToStruct(rgt.MetaData, &rgt_md)
				rrls.List = append(rrls.List, rgt)
				rrls.Meta = append(rrls.Meta, rgt_md)
			}
		}
	}else{
		DbConn.QueryRow("select COUNT(*) FROM resourceGroup WHERE units_id = $1 and info like '%$2%'", theUser.GroupId, search).Scan(&rrls.Count)
	}
	
	rrls_json := StructToJson(rrls)
	
	fmt.Fprint(w, rrls_json)
	fmt.Println(rrls_json)
}
