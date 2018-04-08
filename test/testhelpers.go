package test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	contract "github.com/jjosephy/interview/contract/v1"
)

func readSuccessResponse(b io.ReadCloser) (contract.InterviewContractV1, error) {
	var c contract.InterviewContractV1
	body, err := ioutil.ReadAll(b)
	if err != nil {
		return c, err
	}
	json.Unmarshal(body, &c)
	return c, nil
}

func readErrorResponse(b io.ReadCloser) (contract.ErrorContractV1, error) {
	var errDetail contract.ErrorContractV1
	body, err := ioutil.ReadAll(b)
	if err != nil {
		return errDetail, err
	}
	json.Unmarshal(body, &errDetail)

	return errDetail, nil
}

func assertErrorEqual(
	t *testing.T,
	e contract.ErrorContractV1,
	code int,
	msg string,
	statusExpected int,
	statusReceived int) {
	if e.Code != code {
		t.Fatalf("Error Code is not correct expected %d, got %d", code, e.Code)
	}

	if e.Message != msg {
		t.Fatalf("Error Messages do not match: %s", e.Message)
	}

	if statusExpected != statusReceived {
		t.Fatalf(
			"Unexpected status code returned expected %d : received %d",
			statusExpected,
			statusReceived)
	}
}

// ValidateRequest isTest Helper method to make a GETweb request and validate on provided conditions.
func ValidateRequest(
	t *testing.T,
	uri string,
	headers map[string]string,
	errCode int,
	errMsg string,
	expectedHTTPStatus int) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		t.Fatalf("Unexpected error trying to create a request %v", err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if resp != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	/*
		switch expectedHTTPStatus {
		case http.StatusOK:
			c, err := readSuccessResponse(resp.Body)
			if err != nil {
				t.Fatalf("Unexpected error reading body %v", err)
			} else {
				defer resp.Body.Close()
			}

		case http.StatusBadRequest:
			errDetail, err := readErrorResponse(resp.Body)
			if err != nil {
				t.Fatalf("Unexpected error reading body %v", err)
			} else {
				defer resp.Body.Close()
			}

			assertErrorEqual(
				t,
				errDetail,
				errCode,
				errMsg,
				expectedHTTPStatus, expectedHTTPStatus)
		}
	*/
}
