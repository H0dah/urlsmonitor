package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/h0dah/uptimemonitor/check"
	"github.com/h0dah/uptimemonitor/report"

	"github.com/google/uuid"
)

type GetCheckRequest struct {
	UserId int `json:"user_id"`
}

type CreateCheckRequest struct {
	Id             int               `json:"id"`
	Name           string            `json:"name"`
	Url            string            `json:"url"`
	Protocol       string            `json:"protocol"` //HTTP, HTTPS, TCP
	Path           string            `json:"path"`
	Port           string            `json:"port"`
	WebhookUrl     string            `json:"webhook_url"`
	Timeout        int32             `json:"time_out"` //millisecond
	Interval       int32             `json:"time_interval"`
	Threshold      int               `json:"threshold"`
	Authentication string            `json:"authentication"`
	HttpHeaders    map[string]string `json:"httpHeaders"`
	Assert         []string          `json:"Assert"`
	Tags           []string          `json:"Tags"`
	IgnoreSSL      bool              `json:"IgnoreSSL"` //True--> ignore
	UserId         int               `json:"user_id"`
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

// get list of checks for specified user
func GetChecks(w http.ResponseWriter, r *http.Request) {
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

func CreateCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestBody CreateCheckRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	id := string(rune(uuid.New().ID()))
	newCheck := check.Check{
		ID:             id,
		UserId:         requestBody.UserId,
		Name:           requestBody.Name,
		Url:            requestBody.Url,
		Protocol:       requestBody.Protocol,
		Path:           requestBody.Path,
		Port:           requestBody.Port,
		WebhookUrl:     requestBody.WebhookUrl,
		Timeout:        requestBody.Timeout,
		Interval:       requestBody.Interval,
		Threshold:      requestBody.Threshold,
		Authentication: requestBody.Authentication,
		HttpHeaders:    requestBody.HttpHeaders,
		Assert:         requestBody.Assert,
		Tags:           requestBody.Tags,
		IgnoreSSL:      requestBody.IgnoreSSL,
	}
	check.AddCheck(newCheck)

	err = json.NewEncoder(w).Encode(newCheck)
	if err != nil {
		err_response := ErrorResponse{Error: "couldn't encode the new check to json"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_response)
		return
	}
}

type DeleteCheckRequest struct {
	UserId  int
	CheckId string
}

func DeleteCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestBody DeleteCheckRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	check.DeleteCheck(requestBody.UserId, requestBody.CheckId)

	err = json.NewEncoder(w).Encode("check deleted")
	if err != nil {
		err_response := ErrorResponse{Error: "error"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_response)
		return
	}

}

type PutCheckRequest struct {
	Id             int               `json:"id"`
	UserId         int               `json:"user_id"`
	Name           string            `json:"name"`
	Url            string            `json:"url"`
	Protocol       string            `json:"protocol"` //HTTP, HTTPS, TCP
	Path           string            `json:"path"`
	Port           string            `json:"port"`
	WebhookUrl     string            `json:"webhook_url"`
	Timeout        int32             `json:"time_out"` //millisecond
	Interval       int32             `json:"time_interval"`
	Threshold      int               `json:"threshold"`
	Authentication string            `json:"authentication"`
	HttpHeaders    map[string]string `json:"httpHeaders"`
	Assert         []string          `json:"Assert"`
	Tags           []string          `json:"Tags"`
	IgnoreSSL      bool              `json:"IgnoreSSL"` //True--> ignore
}

func PutCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestBody PutCheckRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// Update the check with the new data
	updatedCheck := check.Check{
		ID:             strconv.Itoa(requestBody.Id),
		UserId:         requestBody.UserId,
		Name:           requestBody.Name,
		Url:            requestBody.Url,
		Protocol:       requestBody.Protocol,
		Path:           requestBody.Path,
		Port:           requestBody.Port,
		WebhookUrl:     requestBody.WebhookUrl,
		Timeout:        requestBody.Timeout,
		Interval:       requestBody.Interval,
		Threshold:      requestBody.Threshold,
		Authentication: requestBody.Authentication,
		HttpHeaders:    requestBody.HttpHeaders,
		Assert:         requestBody.Assert,
		Tags:           requestBody.Tags,
		IgnoreSSL:      requestBody.IgnoreSSL,
	}

	err = check.UpdateCheck(updatedCheck)
	if err != nil {
		err_response := ErrorResponse{Error: "couldn't update the check"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_response)
		return
	}

	err = json.NewEncoder(w).Encode(updatedCheck)
	if err != nil {
		err_response := ErrorResponse{Error: "couldn't encode the updated check to json"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_response)
		return
	}
}
