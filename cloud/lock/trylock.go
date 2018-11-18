package main


import (
	"fmt"
	"sync"
)

// Lock try lock
type Lock struct {
	c chan struct{}
}

// NewLock generate a try lock
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

// Lock try lock, return lock result
func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

// Unlock , Unlock the try lock
func (l Lock) Unlock() {
	l.c <- struct{}{}
}

var counter int


// 通过 chan进行，进行操作

func main() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}


/*
逻辑限定每个 goroutine 只有成功执行了 Lock 才会继续执行后续逻辑，因此在 Unlock 时可以保证 Lock struct 中的 channel 一定是空，从而不会阻塞，也不会失败。

在单机系统中，trylock 并不是一个好选择。因为大量的 goroutine 抢锁可能会导致 cpu 无意义的资源浪费。有一个专有名词用来描述这种抢锁的场景：活锁。

活锁指的是程序看起来在正常执行，但实际上 cpu 周期被浪费在抢锁，而非执行任务上，从而程序整体的执行效率低下。活锁的问题定位起来要麻烦很多。所以在单机场景下，不建议使用这种锁。


 */