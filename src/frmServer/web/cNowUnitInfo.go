package web

import (
	. "frmPkg"
	. "frmServer/public"
	"fmt"
	"net/http"
	"strings"
)


func cNowUnitInfo(ili *IsLoginInfo, w http.ResponseWriter, r *http.Request){
	
	unitId := ili.UnitId
	
	var theUnit UnitsTable
	
	DbConn.QueryRow("select id, name, expand, powerlevel, info from units where id = $1", unitId).Scan(&theUnit.Id, &theUnit.Name, &theUnit.Expand, &theUnit.PowerLevel, &theUnit.Info)
	
	theUnit.Name = strings.TrimSpace(theUnit.Name)
	
	theJSON := StructToJson(theUnit)
	fmt.Fprint(w, theJSON )
}
