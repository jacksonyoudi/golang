package scheduler

import "golang/spider/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}
