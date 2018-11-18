package main

import "sync"


var counter_1 int

func main()  {
	var wg sync.WaitGroup
	var l sync.Mutex
	for i := 0; i < 1000; i++ {
		// 增加一个进程
		wg.Add(1)
		go func() {
			// 执行完一个 done一下
			defer wg.Done()
			l.Lock()
			counter_1++
			l.Unlock()
		}()
	}
	// 等待所有的进程完成
	wg.Wait()
	println(counter_1)
}
