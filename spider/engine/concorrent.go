package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureWorkerChan(chan Request)
}



func (e *ConcurrentEngine) Run(seeds ...Request)  {

	in := make(chan Request)
	out := make(chan ParserResult)

	e.Scheduler.ConfigureWorkerChan(in)



	for i := 0; i < e.WorkerCount; i++ {
		createWorkder(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	countItem := 0
	for {
		result := <- out
		for _, item := range result.Items {
			fmt.Printf("Got item:%d %v\n",countItem, item)
		}
		countItem++
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

func createWorkder(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}