package report

import (
	"net/http"
	"time"
)

type Report struct {
	Status          string    `json:"status"`
	Availability    float64   `json:"availability"`
	Outages         int       `json:"outages"`
	Downtime        int64     `json:"downtime"`
	Uptime          int64     `json:"uptime"`
	AvgResponseTime float64   `json:"avgResponseTime"`
	History         time.Time `json:"history"`
}

var report Report

func ProcessReport(url string) {
	// report status
	// report downtime: The total time, in seconds, of the URL downtime.
	// report uptime: The total time, in seconds, of the URL uptime.
	// report outages: The total number of URL downtimes.
	// report availability: A percentage of the URL availability

	// interval := time.NewTicker(10 * time.Minute)
	interval := 2 * time.Second // in seconds
	intervalTicker := time.NewTicker(interval)

	var downtime time.Duration = 0
	var uptime time.Duration = 0
	var avgResponseTime float64 = 0
	var numOfPolls float64 = 0
	outages := 0

	for {
		t := <-intervalTicker.C
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
		responseTime := time.Since(start).Seconds()

		avgResponseTime = ((avgResponseTime * numOfPolls) + responseTime) / (numOfPolls + 1)
		numOfPolls += 1

		if resp.StatusCode == 200 {
			uptime += interval
		} else {
			downtime += interval
			outages += 1
		}

		availability := uptime.Seconds() / (uptime + downtime).Seconds() * 100

		report = Report{
			Status:          resp.Status,
			History:         t,
			Availability:    availability,
			Downtime:        int64(downtime),
			Uptime:          int64(uptime),
			Outages:         outages,
			AvgResponseTime: avgResponseTime,
		}

	}
}

func GetReport() Report {
	return report
}
