//异步缓存全文索引

package public

import(
	"sync"
	"fmt"
	"time"
	"strings"
)

// 辅助
type acftiAid struct {
	HashId string
	Type string  //可选择的：rg、rf、rt
}

// 一个搜索后结果
type acftiOneSearch struct {
	HashId string
	UnitId uint16
}

//正式
type AsyncCacheFullTextIndex struct {
	lock *sync.RWMutex
	wait int64
	Del []acftiAid
	Up []acftiAid
	KeyWord []string
}

//NewAsyncCachFullTextIndex 新建异步缓存全文索引
func NewAsyncCachFullTextIndex (modewait int64) *AsyncCacheFullTextIndex {
	return &AsyncCacheFullTextIndex{new(sync.RWMutex), modewait , []acftiAid{}, []acftiAid{}, []string{}}
}

// Insert 插入一条待处理数据
func (acf *AsyncCacheFullTextIndex) Insert(mode, hashid, htype string){
	acf.lock.Lock()
	defer acf.lock.Unlock()
	switch mode {
		case "del":
			candel := true
			for _, one := range acf.Del {
				if one.HashId == hashid && one.Type == htype {
					candel = false
					break
				}
			}
			if candel == true {
				acf.Del = append(acf.Del, acftiAid{hashid, htype})
			}
		case "up":
			canup := true
			for _, one := range acf.Up {
				if one.HashId == hashid && one.Type == htype {
					canup = false
					break
				}
			}
			if canup == true {
				hashid = "'"+hashid+"'"
				acf.Up = append(acf.Up, acftiAid{hashid, htype})
			}
	}
}

// InsertWord 插入个待处理关键词
func  (acf *AsyncCacheFullTextIndex) InsertWord(word string){
	acf.lock.Lock()
	defer acf.lock.Unlock()
	canadd := true
	for _, oneword := range acf.KeyWord {
		if oneword == word {
			canadd = false
			break
		}
	}
	if canadd == true {
		acf.KeyWord = append(acf.KeyWord, word)
	}
}

// SearchString 返回一个逗号分割的字符串
func (acf *AsyncCacheFullTextIndex) SearchString (key_word, htype string, uid uint16) (hashids string, key_count uint64){
	hashid, key_count := acf.Search(key_word, htype, uid)
	newhash := make([]string, 0, len(hashid))
	for _, one := range hashid {
		one = "'" + one + "'"
		newhash = append(newhash, one)
	}
	hashids = strings.Join(newhash, ", ")
	return 
}

// Search 执行搜索
func (acf *AsyncCacheFullTextIndex) Search (key_word, htype string, uid uint16) (hashid []string, key_count uint64) {
	key_word = strings.ToLower(key_word)
	
	var htype_int int
	switch htype {
		case "rg":
			htype_int = 1
		case "rf":
			htype_int = 2
		case "rt":
			htype_int = 3
	}
	sql1 := fmt.Sprintf("from acfti where key_word = $1 and htype = %v", htype_int);
	//if uid != 0 {
	//	sql1 = fmt.Sprintln(sql1, "and uid =",uid)
	//}
	sql1_count := "select COUNT(*) " + sql1
	sql1_hashid := "select hashid, uid " + sql1
	
	DbConn.QueryRow(sql1_count, key_word).Scan(&key_count)
	if key_count == 0 {
		// 如果关键词索引里没有的处理方法
		// 加入异步缓存的KeyWord列表，然后做简单的标题搜索
		acf.InsertWord(key_word)
		sql2 := "select hashid from "
		switch htype {
			case "rg":
				sql2 += "resourceGroup"
			case "rf":
				sql2 += "resourceFile"
			case "rt":
				sql2 += "resourceText"
		}
		sql2 += " where name ilike '%" + key_word + "%' "
		if uid != 0 {
			sql2 += fmt.Sprintf(" and units_id = %v",uid)
		}
		key_index2 ,  _ := DbConn.Query(sql2)
		for key_index2.Next(){
			var one_hashid string
			key_index2.Scan(&one_hashid)
			hashid = append(hashid,one_hashid)
		}
		key_count = uint64(len(hashid))
		return
	}
	
	hashid = make([]string, 0, key_count)
	key_index ,  _ := DbConn.Query(sql1_hashid, key_word)
	for key_index.Next(){
		var one_hashid string
		var one_uid uint16
		key_index.Scan(&one_hashid, &one_uid)
		if uid != 0 && uid != one_uid {
			continue  //如果不是要找的那个机构ID的化，则直接略过
		}
		hashid = append(hashid, one_hashid)
	}
	return
}

// AsyncCache 异步缓存
func (acf *AsyncCacheFullTextIndex) AsyncCache(){
	for {
		time.Sleep(time.Duration(acf.wait)*time.Second)
		acf.cacheDel()
		acf.cacheUp()
		acf.cacheKeyWord()
	}
}

// 缓存删除的，其实就是删除掉已经删除了的数据
func (acf *AsyncCacheFullTextIndex) cacheDel(){
	acf.lock.Lock()
	defer acf.lock.Unlock()
	del_pre , _ := DbConn.Prepare("delete from acfti where htype = $1 and hashid = $2")
	for _, beDel := range acf.Del {
		switch beDel.Type {
			case "rg":
				del_pre.Exec(1, beDel.HashId)
			case "rf":
				del_pre.Exec(2, beDel.HashId)
			case "rt":
				del_pre.Exec(3, beDel.HashId)
		}
	}
	acf.Del = []acftiAid{}
}

// 缓存得到更新的
func (acf *AsyncCacheFullTextIndex) cacheUp(){
	acf.lock.Lock()
	defer acf.lock.Unlock()
	if len(acf.Up) == 0 {
		return
	}
	up_rg := []string{}
	up_rf := []string{}
	up_rt := []string{}
	del_pre, _ := DbConn.Prepare("delete from acfti where htype = $1 and hashid = $2")
	for _, oneA := range acf.Up {
		switch oneA.Type {
			case "rg":
				up_rg = append(up_rg, oneA.HashId)
				del_pre.Exec(1, oneA.HashId)
			case "rf":
				up_rf = append(up_rf, oneA.HashId)
				del_pre.Exec(2, oneA.HashId)
			case "rt":
				up_rt = append(up_rt, oneA.HashId)
				del_pre.Exec(3, oneA.HashId)
		}
	}
	
	upstring_rg := strings.Join(up_rg, ", ")
	allwords := acf.getAllKeyWord()
	sql_p, _ := DbConn.Prepare("insert into acfti (key_word, uid, hashid, htype) values ($1, $2, $3, $4)")
	for _, oneWord := range allwords {
		searchRg := acf.searchFromRg(oneWord, upstring_rg)
		for _, oneS := range searchRg {
			sql_p.Exec(oneWord, oneS.UnitId, oneS.HashId, 1)
		}
	}
	acf.Up = []acftiAid{}
}

// 缓存新的关键词
func (acf *AsyncCacheFullTextIndex) cacheKeyWord(){
	acf.lock.Lock()
	defer acf.lock.Unlock()
	if len(acf.KeyWord) == 0 {
		return
	}
	sql_p, _ := DbConn.Prepare("insert into acfti (key_word, uid, hashid, htype) values ($1, $2, $3, $4)")
	sql_del, _ := DbConn.Prepare("delete from acfti where key_word = $1")
	for _, oneWord := range acf.KeyWord {
		sql_del.Exec(oneWord)  //为了以防万一重复了，删除所有的这个关键词的项目
		
		sql_p.Exec(oneWord, 1, "0000000000000000000000000000000000000000", 1)  //写入站位的全零项目
		sql_p.Exec(oneWord, 1, "0000000000000000000000000000000000000000", 2)  //写入站位的全零项目
		sql_p.Exec(oneWord, 1, "0000000000000000000000000000000000000000", 3)  //写入站位的全零项目
		
		// begin 处理ResourceGroup的搜索
		searchRg := acf.searchFromRg(oneWord,"0")
		if len(searchRg) == 0 {
			continue
		}
		for _, oneS := range searchRg {
			sql_p.Exec(oneWord, oneS.UnitId, oneS.HashId, 1)
		}
		// end 处理ResourceGroup的搜索
		
		// 处理ResourceFile的搜索 TODO
		// 处理ResourceText的搜索 TODO
	}
	acf.KeyWord = []string{}
}

// 获取所有现有关键词
func (acf *AsyncCacheFullTextIndex) getAllKeyWord() (allword []string) {
	allkey ,  _ := DbConn.Query("select key_word from acfti group by key_word")
	for allkey.Next(){
		var oneword string
		allkey.Scan(&oneword)
		oneword = strings.TrimSpace(oneword)
		allword = append(allword, oneword)
	}
	return
}

// 从ResourceGroup（资源聚集）中搜索
func (acf *AsyncCacheFullTextIndex) searchFromRg (keyword string, hashstring string) (searchre []acftiOneSearch){
	var sql string
	if hashstring == "0"{
		sql = "select hashid, units_id from resourceGroup where name ilike '%"+keyword+"%' or info ilike '%"+keyword+"%' or metadata->>'Author' ilike '%"+keyword+"%' or metadata->>'Editor' ilike '%"+keyword+"%' or metadata->>'ISBN' ilike '%"+keyword+"%'"
	}else{
		sql = "select hashid, units_id from resourceGroup where (name ilike '%"+keyword+"%' or info ilike '%"+keyword+"%' or metadata->>'Author' ilike '%"+keyword+"%' or metadata->>'Editor' ilike '%"+keyword+"%' or metadata->>'ISBN' ilike '%"+keyword+"%') and hashid in ( "+hashstring+" )"
	}
	search, _ := DbConn.Query(sql)
	for search.Next() {
		onesr := acftiOneSearch{}
		search.Scan(&onesr.HashId, &onesr.UnitId)
		if onesr.UnitId != 1 {
			searchre = append(searchre, onesr)
		}
	}
	return
}

// 从ResourceFile（资源文件）中搜索，暂缓实现
func (acf *AsyncCacheFullTextIndex) searchFromRf (keyword string) (searchre []acftiOneSearch){
	return
}

// 从ResourceText（资源文本）中搜索，暂缓实现
func (acf *AsyncCacheFullTextIndex) searchFromRt (keyword string) (searchre []acftiOneSearch){
	return
}
