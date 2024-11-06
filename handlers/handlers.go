package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/h0dah/uptimemonitor/check"
	"github.com/h0dah/uptimemonitor/report"
)

type GetCheckRequest struct {
	UserId int `json:"user_id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func GetReport(w http.ResponseWriter, r *http.Request) {
	report := report.GetReport()
	err := json.NewEncoder(w).Encode(report)
	if err != nil {
		err_response := ErrorResponse{Error: "couldn't encode Report to json"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_response)
		return
	}

}

func GetCheck(w http.ResponseWriter, r *http.Request) {
	var requestBody GetCheckRequest
	checkList := check.GetCheckByID(requestBody.UserId)

	err := json.NewEncoder(w).Encode(checkList)
	if err != nil {
		err_response := ErrorResponse{Error: "couldn't encode checks to json"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_response)
		return
	}
}
