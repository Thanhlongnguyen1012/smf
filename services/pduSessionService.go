package services

import (
	"fmt"
	"io"
	"net/http"
	"smf/api"
	"smf/models"
)

func CreateSMContext(data models.SMContextCreateData) (int, interface{}) {
	resp, err := api.SmfClient.PduSessionCreation(data)
	if err == nil && (resp.StatusCode == 200 || resp.StatusCode == 201) {
		created := models.SMContextCreatedData{
			PduSessionID: data.PduSessionId,
		}
		fmt.Println("SM Context Created")
		return http.StatusCreated, created
	}

	fmt.Println("Failed to create SM Context")

	var cause string
	if resp != nil && resp.Body != nil {
		body, _ := io.ReadAll(resp.Body)
		cause = string(body)
	} else {
		cause = "Unknown failure"
	}

	problem := models.ProblemDetails{
		Title:  "Failed to create SM Context at UDM",
		Status: http.StatusBadRequest,
		Cause:  cause,
	}
	return http.StatusInternalServerError, problem
}
