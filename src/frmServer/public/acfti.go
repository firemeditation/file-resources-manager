//异步缓存全文索引

package public

import(
	"sync"
	"fmt"
	"time"
)

// 辅助
type acftiAid struct {
	HashId string
	Type string  //可选择的：rg、rf、rt
}

//正式
type AsyncCacheFullTextIndex struct {
	lock *sync.RWMutex
	wait int64
	Add []acftiAid
	Del []acftiAid
	Up []acftiAid
	KeyWord []string
}

//NewAsyncCachFullTextIndex 新建异步缓存全文索引
func NewAsyncCachFullTextIndex (modewait int64) *AsyncCacheFullTextIndex {
	return &AsyncCacheFullTextIndex{new(sync.RWMutex), modewait , []acftiAid{}, []acftiAid{}, []acftiAid{}, []string{}}
}

// Insert 插入一条待处理数据
func (acf *AsyncCacheFullTextIndex) Insert(mode int, hashid, htype string){
	acf.lock.Lock()
	defer acf.lock.Unlock()
	switch mode {
		case 1:
			acf.Add = append(acf.Add, acftiAid{hashid, htype})
		case 2:
			acf.Del = append(acf.Del, acftiAid{hashid, htype})
		case 3:
			acf.Up = append(acf.Up, acftiAid{hashid, htype})
	}
}

// InsertWord 插入个待处理关键词
func  (acf *AsyncCacheFullTextIndex) InsertWord(word string){
	acf.lock.Lock()
	defer acf.lock.Unlock()
	acf.KeyWord = append(acf.KeyWord, word)
}

// Search 执行搜索
func (acf *AsyncCacheFullTextIndex) Search (key_word, htype string, uid uint16) (hashid []string, key_count uint64) {
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
		acf.lock.Lock()
		defer acf.lock.Unlock()
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
		sql2 += " where name like '%" + key_word + "%' "
		if uid != 0 {
			sql2 += fmt.Sprintf(" and units_id = %v",uid)
		}
		key_index2 ,  _ := DbConn.Query(sql2)
		for key_index2.Next(){
			var one_hashid string
			key_index2.Scan(&one_hashid)
			hashid = append(hashid,one_hashid)
		}
		return
	}
	
	hashid = make([]string, key_count)
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
		acf.lock.Lock()
		defer acf.lock.Unlock()
		acf.cacheDel()
		acf.cacheAdd()
		acf.cacheUp()
		acf.cacheKeyWord()
	}
}

// 缓存新添加的
func (acf *AsyncCacheFullTextIndex) cacheAdd(){
	
}

// 缓存删除的，其实就是删除掉已经删除了的数据
func (acf *AsyncCacheFullTextIndex) cacheDel(){
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
}

// 缓存得到更新的
func (acf *AsyncCacheFullTextIndex) cacheUp(){
	
}

// 缓存新的关键词
func (acf *AsyncCacheFullTextIndex) cacheKeyWord(){
	
}

// 获取所有现有关键词
func (acf *AsyncCacheFullTextIndex) getAllKeyWord() (allword []string) {
	return
}
