package scheduler

import "golang/spider/engine"

type QueueShduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s QueueShduler) Submit(r engine.Request)  {
	s.requestChan <- r
}

func (s QueueShduler) ConfigureWorkerChan(chan ) {
	
}

func (s QueueShduler) WorkerRelay()  {

}

