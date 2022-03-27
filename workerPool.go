package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type workerPool struct {
	workers    []*worker
	workerChan chan *worker
	completed  chan bool
}

func newWorkerPool(workers []*worker) *workerPool {
	return &workerPool{
		workers:    workers,
		workerChan: make(chan *worker, len(workers)),
		completed:  make(chan bool),
	}
}

func workerGenerator(numWorkers int, wh *warehouse) []*worker {
	workersConfig := make([]*worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workersConfig[i] = newWorker(i, 16, wh)
	}

	return workersConfig
}

func (wP *workerPool) getAvailableWorker() {
	go func() {
		for _, worker := range wP.workers {
			if worker.status == idle {
				wP.workerChan <- worker
			}
		}
	}()
}

func (wP *workerPool) addWorkerBackToPool(wo *worker) {
	wP.workerChan <- wo
	wg.Done()

}

func (wP workerPool) processJobs(jB *jobBatch, concurrent bool, r *allResults) {
	pendingJob := jB.sendJobs()
	// Will start sending workers to the pool
	wP.getAvailableWorker()
	for i := 0; i < len(jB.jobs); i++ {
		wg.Add(1)

		//Will get a worker from the pool, like ticket system in the market
		// if new workers are added it will be added to the pool
		nextWorker := <-wP.workerChan
		fmt.Println(nextWorker.totalDistanceTraveled)
		//Will pull job from the pending job channel
		nextJob := <-pendingJob
		if nextWorker.status == idle {
			fmt.Println("Grabbed Worker")
			fmt.Println("Worker ID:", nextWorker.id, "is currently", nextWorker.getStatus())
			if concurrent { //Run All jobs concurrently
				go nextWorker.assignJob(nextJob, &wP)
			} else { //Run All jobs sequentially
				nextWorker.assignJob(nextJob, &wP)
			}
			nextJob.status = wip
		}
	}
	// Block Function until completed channel is closed
	wg.Wait()

	go func() {
		wP.sendResults(r)
	}()
	r.printResults()

	// go func() {
	// 	time.Sleep(time.Second * 5)
	// 	wP.completed <- true
	// }()

	// wP.sendResults()
}

func (wP *workerPool) sendResults(r *allResults) {
	//Get worker stats from each worker
	defer close(r.resultChan)
	for _, worker := range wP.workers {
		workerData := result{
			worker.id,
			worker.workerStats,
		}
		r.resultChan <- workerData
	}

}

func (wP *workerPool) gatherResults(result) {
	//TODO
}
