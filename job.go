// 	       ==================================
// Job Batch--->   --->   ---> job   --->   Worker Pool
//         ==================================
package main

import (
	"fmt"
)

type jobStatus int

const (
	pending jobStatus = iota
	wip
	complete
)

type job struct {
	id int
	sku
	quantity int
	strategy func(int) int
	status   jobStatus
}

func (j job) getStatus() string {
	return [...]string{"pending", "wip", "complete"}[j.status]
}

type jobBatch struct {
	jobs    []job
	jobChan chan job
}

func newJobBatch(jobs []job, buffer int) *jobBatch {
	return &jobBatch{
		jobs:    jobs,
		jobChan: make(chan job, buffer),
	}
}

func (jB *jobBatch) sendJobs() chan job {

	go func() {
		// Don't need this because iterating over the jobs in workers process Jobs
		// defer close(jB.jobChan)
		for _, job := range jB.jobs {
			fmt.Println("Job", job.id, "Send")
			// sleepRandomTime(10)
			jB.jobChan <- job
		}
	}()
	return jB.jobChan
}
