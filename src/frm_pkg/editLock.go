//EditLock 为编辑锁

package frm_pkg


import (
	"sync"
)

type EditLock map[string]*sync.RWMutex

func NewEditLock () EditLock {
	return EditLock{}
}

func (el EditLock) Lock (ck string) {
	if _ , found := el[ck] ; found == true {
		el[ck].Lock()
	}else{
		el[ck] = new(sync.RWMutex)
		el[ck].Lock()
	}
}

func (el EditLock) RLock (ck string) {
	if _ , found := el[ck] ; found == true {
		el[ck].RLock()
	}else{
		el[ck] = new(sync.RWMutex)
		el[ck].RLock()
	}
}

func (el EditLock) Unlock (ck string){
	if _ , found := el[ck] ; found == true {
		el[ck].Unlock()
		delete(el,ck)
	}
}

func (el EditLock) RUnlock (ck string){
	if _ , found := el[ck] ; found == true {
		el[ck].RUnlock()
		delete(el,ck)
	}
}
