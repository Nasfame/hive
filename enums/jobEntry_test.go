package enums

import (
	"fmt"
	"testing"
	"time"
)

func TestJobEntry(t *testing.T) {
	startTime := time.Now()
	endTime := startTime.Add(time.Minute * 5)
	job := &JobEntry{}

	// Test Start method
	job.Start()
	if job.startTime.IsZero() {
		t.Error("Start method did not set startTime")
	}

	// Test Stop method
	job.Stop()
	if job.endTime.IsZero() {
		t.Error("Stop method did not set endTime")
	}

	// Test Duration method
	job.startTime = startTime
	job.endTime = endTime
	expectedDuration := time.Minute * 5
	if job.Duration() != expectedDuration {
		t.Errorf("Duration method returned %s, expected %s", job.Duration(), expectedDuration)
	}

	// Test String method
	expectedString := fmt.Sprintf("start: %s, end: %s, duration: %s", startTime, endTime, expectedDuration)
	if job.String() != expectedString {
		t.Errorf("String method returned %s, expected %s", job.String(), expectedString)
	}
}
