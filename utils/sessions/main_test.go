package sessionsUtils_test

import (
	"testing"
	"time"

	m "github.com/xindixu/todo-time-tracker/models"
	sessionsUtils "github.com/xindixu/todo-time-tracker/utils/sessions"
)

func TestFormat1(t *testing.T) {
	started, _ := time.Parse(time.RFC3339, "2022-09-20T15:00:00Z")
	ended, _ := time.Parse(time.RFC3339, "2022-09-20T16:00:00Z")

	session := m.Session{
		Started: started,
		Ended:   ended,
		Task:    "learn go",
	}

	expected := "Tue Sep 20 15:00:00 UTC 2022 - Tue Sep 20 16:00:00 UTC 2022:  learn go"
	s := sessionsUtils.Format(session)
	if s != expected {
		t.Errorf("\nExpected: %s\nGot: %s", expected, s)
	}
}

func TestFormat2(t *testing.T) {
	started, _ := time.Parse(time.RFC3339, "2022-09-20T15:00:00Z")

	session := m.Session{
		Started: started,
		Ended:   time.Time{},
		Task:    "learn go",
	}

	expected := "Tue Sep 20 15:00:00 UTC 2022 - Now:  learn go"
	s := sessionsUtils.Format(session)
	if s != expected {
		t.Errorf("\nExpected: %s\nGot: %s", expected, s)
	}
}
