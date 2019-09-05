package engine

import (
	"crawler/src/model"
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	Count     int
	ItemChan  chan interface{}
}

type Scheduler interface {
	Submit(Request)
	GetWorkerChan() chan Request
	Run()
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	// 初始化request队列和worker队列
	e.Scheduler.Run()
	for i := 0; i < e.Count; i++ {
		createWorker(e.Scheduler.GetWorkerChan(), e.Scheduler, out)
	}

	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item %v\n", item)
			if _, ok := item.(model.Profile); ok {
				// 真正存储用户的信息
				e.ItemChan <- item
			}
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, s Scheduler, out chan ParseResult) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
