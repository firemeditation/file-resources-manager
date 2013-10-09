package main

import (
	"fmt"
	. "frmPkg"
)

// newResource 新建资源条目
func newResource () {
	
	// start 自身判断权限
	if myLogin.UPower["resource"]["origin"] < 10 {
		fmt.Print("您没有新建资源条目的权力，按任意键继续。")
		var tep string
		fmt.Scanln(&tep)
		return
	}
	// end
	
	// begin 输入要添加的信息
	var rgt ResourceGroupTable
	var rt_id []uint32
	fmt.Printf("选择要添加的资源类型：")
	for _, value := range resourceType {
		fmt.Printf("%d：%s\t\t",value.Id,value.Name)
		rt_id = append(rt_id, value.Id)
	}
	fmt.Printf("\n")
	fmt.Scanln(&rgt.RtId)
	rtid_ok := false
	for _, v := range rt_id {
		if v == rgt.RtId {
			rtid_ok = true
			break
		}
	}
	if rtid_ok == false {
		fmt.Print("请正确输入资源类型的ID，按任意键继续。")
		var tep string
		fmt.Scanln(&tep)
		return
	}
	fmt.Print("资源名：")
	fmt.Scanln(&rgt.Name)
	fmt.Print("资源说明：")
	fmt.Scanln(&rgt.Info)
	// end
	
	doNewResource(&rgt)
}
