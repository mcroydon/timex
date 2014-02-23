package timex

import (
	"testing"
	"time"
)

// Basic usage testing. If this fails, Timex probably won't work on your system.
func TestBasic(t *testing.T) {
	// Get a Timex representing the current time.
	tmx := Now()

	// Make sure that the basic information is not the zero value.
	if tmx.Time == time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC) {
		t.Errorf("Timex.Time is nil.")
	}
	if tmx.Status == 0 {
		t.Errorf("Timex.Status is 0.")
	}
	if tmx.StatusInfo() == "" {
		t.Errorf("Timex.StatusInfo() is \"\".")
	}
	if tmx.Maxerror == 0 {
		t.Errorf("Timex.Maxerror is 0.")
	}
	if tmx.Esterror == 0 {
		t.Errorf("Timex.Esterror is 0.")
	}

	// Log current values for those running verbose tests.
	t.Logf("Time: %v Status: %v Maxerror: %v Esterror: %v\n",
		tmx.Time, tmx.StatusInfo(), tmx.Maxerror, tmx.Esterror)
}
