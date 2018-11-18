package scheduler

import (
	"golang/spider/engine"
)

type QueueShduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueueShduler) Submit(r engine.Request)  {
	s.requestChan <- r
}

func (s *QueueShduler) WorkerReady(w chan engine.Request)  {
	s.workerChan <- w
}


func (s *QueueShduler) ConfigureWorkerChan( ) {

}

func (s *QueueShduler) WorkerRelay()  {

}

func (s *QueueShduler) Run()  {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <- s.requestChan:
				requestQ = append(requestQ, r)
			case w := <- s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
