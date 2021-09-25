package reMutex

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	实现goroutine重入
*/
type RecursiveMutex struct {
	sync.Mutex
	owner int64
	recursion int64
}

// type Locker interface
// 重写了Mutex的Lock方法
func (r *RecursiveMutex) Lock()  {
	gid := GoID()
	// 如果持有锁的goroutine调用,说明是重入
	if atomic.LoadInt64(&r.owner) == gid {
		r.recursion ++
		return
	}
	r.Mutex.Lock()
	// 获得锁的goroutine 第一次调用,记录下它的goroutine id 调用数增加
	atomic.StoreInt64(&r.owner, gid)
	r.recursion = 1
}

func (r *RecursiveMutex) UnLock()  {
	gid := GoID()
	// 非持有锁d的goroutine尝试释放锁(错误使用)
	if atomic.LoadInt64(&r.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d", r.owner, gid))
	}
	// 调用次数
	r.recursion --
	if r.recursion != 0 {
		return
	}
	// 此时是goroutine 最后一次使用需要释放
	atomic.StoreInt64(&r.owner, -1)
	r.Mutex.Unlock()
}