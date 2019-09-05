package scheduler

import (
	"crawler/src/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) GetWorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {

}
func (s *QueuedScheduler) WorkerReady(worker chan engine.Request) {
	s.workerChan <- worker
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
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
			case request := <-s.requestChan:
				requestQ = append(requestQ, request)
			case worker := <-s.workerChan:
				workerQ = append(workerQ, worker)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}

		}
	}()

}
