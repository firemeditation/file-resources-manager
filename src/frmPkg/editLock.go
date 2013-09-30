//EditLock 为编辑锁

package frmPkg


type EditLock map[string]bool

func NewEditLock () EditLock {
	return EditLock{}
}

func (el EditLock) Lock (ck string) bool {
	if _ , found := el[ck] ; found == true {
		return false
	}else{
		el[ck] = true
		return true
	}
}

func (el EditLock) Unlock (ck string){
	if _ , found := el[ck] ; found == true {
		delete(el,ck)
	}
}
