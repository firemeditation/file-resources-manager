//异步缓存全文索引

package public

import(
	"sync"
)

// 辅助
type acftiAid struct {
	HashID string
	Type string  //可选择的：rg、rf、rt
}

//正式
type AsyncCacheFullTextIndex struct {
	lock *sync.RWMutex
	modewait int64
	wordwait int64
	Add []acftiAid
	Del []acftiAid
	Up []acftiAid
	KeyWord []string
}

//NewAsyncCachFullTextIndex 新建异步缓存全文索引
func NewAsyncCachFullTextIndex (modewait, wordwait int64) *AsyncCacheFullTextIndex {
	return &AsyncCacheFullTextIndex{new(sync.RWMutex), modewait, wordwait, []acftiAid{}, []acftiAid{}, []acftiAid{}, []string{}}
}

// Insert 插入一条待处理数据
func (acf *AsyncCacheFullTextIndex) Insert(mode int, hashid, ctype string){
	acf.lock.Lock()
	defer acf.lock.Unlock()
	switch mode {
		case 1:
			acf.Add = append(acf.Add, acftiAid{hashid, ctype})
		case 2:
			acf.Del = append(acf.Del, acftiAid{hashid, ctype})
		case 3:
			acf.Up = append(acf.Up, acftiAid{hashid, ctype})
	}
}

// InsertWord 插入个待处理关键词
func  (acf *AsyncCacheFullTextIndex) InsertWord(word string){
	acf.lock.Lock()
	defer acf.lock.Unlock()
	acf.KeyWord = append(acf.KeyWord, word)
}

// Search 执行搜索

// AsyncCache 异步缓存
