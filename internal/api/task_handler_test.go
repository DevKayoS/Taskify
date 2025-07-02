package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCreateTask(t *testing.T) {
	api := Application{}

	payload := map[string]any{
		"title":       "Learn TDD",
		"description": "Get hands-on exp  with TDD in Go!",
		"priority":    8000,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("Failed to parse our request payload")
	}

	request := httptest.NewRequest("POST", "/api/v1/tasks", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")

	recoder := httptest.NewRecorder()

	handler := http.HandlerFunc(api.handleCreateTask)

	handler.ServeHTTP(recoder, request)

	t.Logf("Rec body %s\n", recoder.Body.Bytes())

	if recoder.Code != http.StatusCreated {
		t.Errorf("StatusCode differs; got %d | want %d", recoder.Code, http.StatusCreated)
	}

	var response map[string]any
	err = json.Unmarshal(recoder.Body.Bytes(), &response)

	if err != nil {
		t.Fatalf("Failed to parse response body: %s\n", err.Error())
	}

	if response["title"] != payload["title"] {
		t.Errorf("title differs; got %d | want %d", response["title"], payload["title"])
	}
}
