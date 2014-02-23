package timex

import (
	"testing"
	"time"
	"github.com/rcrowley/go-metrics"
)

// Basic usage testing. If this fails, Timex probably won't work on your system.
func TestBasic(t *testing.T) {
	// Get a Timex representing the current time.
	tmx := Now()

	// Make sure that the basic information is not the zero value.
	if tmx.Time.IsZero() {
		t.Errorf("Timex.Time is nil.")
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

func BenchmarkHistogram(b *testing.B) {
	// TODO: look at timedelta and clock error during benchmark runs.
	s := metrics.NewUniformSample(100)
	est := metrics.NewHistogram(s)
	max := metrics.NewHistogram(s)
	delta := metrics.NewHistogram(s)

	var prev time.Time
	for i := 0; i < b.N; i++ {
		tmx := Now()
		if prev.Year() != 1 {
			delta.Update(int64(time.Since(prev) / time.Nanosecond))
		}
		est.Update(tmx.Esterror)
		max.Update(tmx.Maxerror)
		prev = tmx.Time
	}
	b.Logf("Estimated error min %v max %v mean %v", est.Min(), est.Max(), est.Mean())
	b.Logf("Maximum error min %v max %v mean %v", max.Min(), max.Max(), max.Mean())
	b.Logf("Time delta min %v max %v mean %v", delta.Min(), delta.Max(), delta.Mean())
}
