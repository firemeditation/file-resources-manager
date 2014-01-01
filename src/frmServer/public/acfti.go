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
	waittime int64
	Add []acftiAid
	Del []acftiAid
	Up []acftiAid
}

//NewAsyncCachFullTextIndex 新建异步缓存全文索引
func NewAsyncCachFullTextIndex (waittime int64) *AsyncCacheFullTextIndex {
	return &AsyncCacheFullTextIndex{new(sync.RWMutex), waittime, []acftiAid{}, []acftiAid{}, []acftiAid{}}
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

// Search 执行搜索

// AsyncCache 异步缓存
