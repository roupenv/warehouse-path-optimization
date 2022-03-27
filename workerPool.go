package main

import (
	"fmt"
	"sync"
)

// Use sync.WaitGroup to wait for all jobs to be completed
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

func (wP *workerPool) getNextAvailableWorker() {
	//TODO Implement Priority Queue to grab least utilized worker
}

func (wP *workerPool) addWorkerBackToPool(wo *worker) {
	wP.workerChan <- wo
	wg.Done()

}

func (wP workerPool) processJobs(jB *jobBatch, concurrent bool) *allResults {
	//Initialize Results
	results := newResult()

	//Start sending jobs
	pendingJob := jB.sendJobs()

	// Start sending workers to the pool
	wP.getAvailableWorker()

	//Iterate over all jobs and assign to idle worker
	for i := 0; i < len(jB.jobs); i++ {

		// Increment the wait group
		wg.Add(1)

		//Pull an idle worker from the pool, like ticket system in the market
		// if new workers are added to buffered channel, it will be added to the pool
		nextWorkerInQueue := <-wP.workerChan

		//Pull job from the pending job channel
		nextJobInQueue := <-pendingJob
		if nextWorkerInQueue.status == idle {

			fmt.Println("Grabbed Worker ID:", nextWorkerInQueue.id, "for Job ID:", nextJobInQueue.id)

			if concurrent { //Run All jobs concurrently
				go nextWorkerInQueue.assignJob(nextJobInQueue, &wP)
			} else { //Run All jobs sequentially
				nextWorkerInQueue.assignJob(nextJobInQueue, &wP)
			}

			nextJobInQueue.status = wip
			//TODO Add to RealTime WIP Report
		}
	}
	// Block Function/Thread until completed channel is closed
	wg.Wait()

	//Gather Results
	results.getResults(wP)

	return results
}



func (wP *workerPool) gatherResults(result) {
	//TODO
}
