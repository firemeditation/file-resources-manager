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

type returnResourceListListStruct struct {
	Table ResourceGroupTable
	MD ResourceGroupTable_MD
	RSR ResourceGroupTable_RSR
}

type returnResourceListStruct struct {
	Count uint64
	List []returnResourceListListStruct
}


func iResourceList(theUser *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	
	// start 查看用户是否有建立资源的权力
	//if theUser.UPower["resource"]["origin"] < 2 {
	//	fmt.Fprint(w,"{\"err\":\"无添加权限\"}")
	//	return
	//}
	
	
	model := 1 //1为全部列出，2为搜索
	
	var key_word string
	var search_type string
	limit := 10
	from := 0
	//begin 整理可能的get
	if _, found := r.Form["key_word"]; found == true {
		model = 2
		key_word = r.Form["key_word"][0]
		if _, found := r.Form["search_type"]; found == true {
			search_type = r.Form["search_type"][0]
		}else{
			search_type = "rg"
		}
	}
	if _, found := r.Form["limit"]; found == true {
		limit, _ = strconv.Atoi(r.Form["limit"][0])
	}
	if _, found := r.Form["from"]; found == true {
		from, _ = strconv.Atoi(r.Form["from"][0])
	}
	//end 整理可能的get
	
	
	
	rrls := returnResourceListStruct{0, []returnResourceListListStruct{}}
	
	if model == 1 {
		DbConn.QueryRow("select COUNT(*) FROM resourceGroup WHERE units_id = $1", theUser.UnitId).Scan(&rrls.Count)
		if rrls.Count != 0 {
			rrls_rows, _ := DbConn.Query("select hashid, name, rt_id, info, btime, users_id, metadata From resourceGroup WHERE units_id = $1 ORDER BY btime DESC LIMIT $2 OFFSET $3", theUser.UnitId, limit, from)
			for rrls_rows.Next(){
				var rgt returnResourceListListStruct
				rrls_rows.Scan(&rgt.Table.HashId, &rgt.Table.Name, &rgt.Table.RtId, &rgt.Table.Info, &rgt.Table.Btime, &rgt.Table.UsersId, &rgt.Table.MetaData)
				rgt.Table.Name = strings.TrimSpace(rgt.Table.Name)
				JsonToStruct(rgt.Table.MetaData, &rgt.MD)
				//获取分类以及所属机构以及用户的名称放入RSR
				rgt.RSR.UnintsName = theUser.UnitName
				DbConn.QueryRow("select name from users where id = $1", rgt.Table.UsersId).Scan(&rgt.RSR.UsersName)
				rgt.RSR.UsersName = strings.TrimSpace(rgt.RSR.UsersName)
				DbConn.QueryRow("select name from resourceType where id = $1", rgt.Table.RtId).Scan(&rgt.RSR.RtName)
				rgt.RSR.RtName = strings.TrimSpace(rgt.RSR.RtName)
				rrls.List = append(rrls.List, rgt)
			}
		}
	}else{
		fmt.Println("开始搜索")
		search_re, thecount := SearchCache.SearchString(key_word, search_type, theUser.UnitId)
		fmt.Println("获得结果",search_re, thecount)
		if len(strings.TrimSpace(search_re)) != 0 {
			theSQL := fmt.Sprintf("select hashid, name, rt_id, info, btime, users_id, metadata From resourceGroup WHERE units_id = %v AND hashid IN ( %v ) ORDER BY btime DESC LIMIT %v OFFSET %v ", theUser.UnitId,search_re,limit,from)
			fmt.Println("搜索的SQL 为", theSQL)
			rrls_rows, _ := DbConn.Query(theSQL)
			for rrls_rows.Next(){
				var rgt returnResourceListListStruct
				rrls_rows.Scan(&rgt.Table.HashId, &rgt.Table.Name, &rgt.Table.RtId, &rgt.Table.Info, &rgt.Table.Btime, &rgt.Table.UsersId, &rgt.Table.MetaData)
				rgt.Table.Name = strings.TrimSpace(rgt.Table.Name)
				JsonToStruct(rgt.Table.MetaData, &rgt.MD)
				//获取分类以及所属机构以及用户的名称放入RSR
				rgt.RSR.UnintsName = theUser.UnitName
				DbConn.QueryRow("select name from users where id = $1", rgt.Table.UsersId).Scan(&rgt.RSR.UsersName)
				rgt.RSR.UsersName = strings.TrimSpace(rgt.RSR.UsersName)
				DbConn.QueryRow("select name from resourceType where id = $1", rgt.Table.RtId).Scan(&rgt.RSR.RtName)
				rgt.RSR.RtName = strings.TrimSpace(rgt.RSR.RtName)
				rrls.List = append(rrls.List, rgt)
			}
			rrls.Count = uint64(len(rrls.List))
		}
	}
	
	rrls_json := StructToJson(rrls)
	
	fmt.Fprint(w, rrls_json)
	//fmt.Println(rrls_json)
}
