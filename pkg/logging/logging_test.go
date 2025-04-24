package logging

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"testing"
)

func Test_LogError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	//log.SetFormatter(&log.JSONFormatter{})

	handler := "testhandler"
	testErr := errors.New("test error")
	LogError(handler, testErr)

	var logEnt map[string]interface{}
	err := json.Unmarshal(buf.Bytes(), &logEnt)
	if err != nil {
		t.Fatal(err)
	}

	if logEnt["level"] != "error" {
		t.Errorf("Expected log level 'error', got '%v'", logEnt["level"])
	}
}
