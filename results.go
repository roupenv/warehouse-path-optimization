package main

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

type result struct {
	workerId int
	workerStats
}

type allResults struct {
	resultChan chan result
}

func newResult() *allResults {
	return &allResults{
		resultChan: make(chan result),
	}
}

func (r *allResults) printResults() {
	os.Stdout = defaultStdOut // restore Standard Out to print table

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Worker Id", "Time Busy", "Time Idle", "Time Waiting", "Distance Traveled", "Trips", "Units Picked"})
	t.AppendSeparator()

	for result := range r.resultChan {
		t.AppendRow([]any{result.workerId, result.timeBusy, result.timeIdle, result.timeWaiting, result.totalDistanceTraveled, result.totalTrips, result.totalUnitsPicked})
	}
	t.AppendSeparator()

	t.AppendRow([]any{len(r.resultChan), "", "", "", "", "", "", "Total"})

	t.SetStyle(table.StyleColoredDark)
	t.Render()
}
