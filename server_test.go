package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
)

func TestHandler(t *testing.T) {

	var ClientInfoList []ClientInfo = make([]ClientInfo, 10)
	var mutex sync.Mutex

	for i := 1; i <= 70; i++ {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handlerfunc := handler(&ClientInfoList, &mutex)
		handlerfunc(rr, req)
		status := rr.Code
		if i <= 60 {
			if status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		} else {
			if status != http.StatusBadRequest {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusBadRequest)
			}
		}

		var expected string
		if i <= 60 {
			expected = strconv.Itoa(i)
		} else {
			expected = "error"
		}

		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}
