package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var defaultStdOut = os.Stdout // keep backup of original stdout

func main() {
	shouldPrint := flag.Bool("p", false, "print the results") // flag to print the results
	flag.Parse()

	if !*shouldPrint { // if the flag is not set, redirect the output to a dummy writer
		os.Stdout = nil // turn off printing by redirecting stdout to /dev/null
	}

	//Start Recording the Time
	start := time.Now()

	//Initialize the Warehouse
	warehouse := newWarehouse(warehouseConfig)

	//Initialize the JobBatch
	jobBatch := newJobBatch(jobsConfig, len(jobsConfig))

	//Initialize the Workers
	workersConfig := workerGenerator(2, warehouse)
	newWorkerPool := newWorkerPool(workersConfig)

	//Start Processing the Jobs
	results := newWorkerPool.processJobs(jobBatch, true)

	//Print Results
	results.printResults()

	//Stop calculating the time
	duration := time.Since(start)
	fmt.Println("Runtime duration", duration)
}
