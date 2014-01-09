package frmPkg

type ResourceFileTreeStruct struct {
	Name string
	HashId string
	IsDir bool
	PowerLevel uint32
	Type uint8 //1.直接资源，2.派生文件
	Files map[string]ResourceFileTreeStruct
}

// ResourceFileToTree 将资源文件变成树状
func ResourceFileToTree( treepoint map[string]ResourceFileTreeStruct, path []string, name, hashid string ){
	if len(path) != 0 {
		if _, found := treepoint[path[0]]; found == false {
			treepoint[path[0]] = ResourceFileTreeStruct{Name:path[0],IsDir:true, Files:map[string]ResourceFileTreeStruct{}}
		}
		nowPath := path[1:]
		ResourceFileToTree(treepoint[path[0]].Files, nowPath, name, hashid)
	}else{
		treepoint[name] = ResourceFileTreeStruct{Name: name, HashId: hashid , IsDir:false}
	}
}
