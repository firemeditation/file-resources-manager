package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	"strings"
	//"regexp"
	//"os"
)


func iGetResourceType(ili *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	var resourceType []ResourceTypeTable
	rts, _ := DbConn.Query("select * from resourcetype")
	for rts.Next(){
		var onert ResourceTypeTable
		rts.Scan(&onert.Id, &onert.Name, &onert.PowerLevel, &onert.Expend, &onert.Info)
		onert.Name = strings.TrimSpace(onert.Name)
		resourceType = append(resourceType, onert)
	}
	theJSON := StructToJson(resourceType)
	fmt.Fprint(w, theJSON )
}
