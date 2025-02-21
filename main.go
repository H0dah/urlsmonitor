package main

import (
	"net/http"

	"github.com/h0dah/uptimemonitor/handlers"
	"github.com/h0dah/uptimemonitor/report"
)

// using local testing server now and will be changed in future
var url = "http://localhost:8090/upordown"

func main() {
	go report.ProcessReport(url)
	http.HandleFunc("/report", handlers.GetReport)
	http.HandleFunc("/get", handlers.GetChecks)
	http.HandleFunc("/post", handlers.CreateCheck)
	http.HandleFunc("/delete", handlers.DeleteCheck)

	http.ListenAndServe(":8080", nil)

}
