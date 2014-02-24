// Package timex provides access to current system time along with clock synchronization status
// and clock error maximum and estimated. This information is provided with cgo calls to ntp_gettime()
// in sys/timex.h. This package will only work on systems with cgo and sys/timex.h (which includes modern
// Linux and many BSD distributions but not Mac OS X).
//
// This information is readily avaialble via Adjtimex in the syscall package (http://golang.org/pkg/syscall/#Adjtimex)
// and you should probably use that instead.
//
// Note: This package is distributed under a BSD license but the effective license may change based on the
// license of sys/timex.h and related files on your system.
package timex

// #include <sys/timex.h>
// struct ntptimeval ntv;
// int update_ntp_gettime()
// {
//	int n = ntp_gettime(&ntv);
//	return n;
// }
// struct ntptimeval get_ntptimeval() {
// 	return ntv;
// }
import "C"
import "time"

const (
	// Ok means that eveything is okay with no leap second warning.
	Ok = C.TIME_OK
	// Ins indicates a positive leap second warning. An additional leap second will be added after
	// 23:59:59 today.
	Ins = C.TIME_INS
	// Del indicates a negative leap second warning. Skip 23:59:59 at the end of the day.
	Del = C.TIME_DEL
	// Oop is when a leap second is occuring.
	Oop = C.TIME_OOP
	// Wait means that leap second has occured.
	Wait = C.TIME_WAIT
	// Error signifies that the clock is not synchronized.
	Error = C.TIME_ERROR
)

// Timex contains the response to a local query of NTP on the local system.
type Timex struct {
	// The status returned by the underlying call to get_ntptimeval().
	Status int
	// The current time according to get_ntptimeval().
	Time time.Time
	// The maximum clock error.
	Maxerror int64
	// The estimated clock error.
	Esterror int64
}

// Now returns a Timex struct based on a query to get_ntptimeval() via cgo.
func Now() (timex *Timex) {
	timex = new(Timex)
	timex.Status = int(C.update_ntp_gettime())
	ntptimeval := C.get_ntptimeval()
	timex.Time = time.Unix(int64(ntptimeval.time.tv_sec), int64(ntptimeval.time.tv_usec))
	timex.Maxerror = int64(ntptimeval.maxerror)
	timex.Esterror = int64(ntptimeval.esterror)
	return
}

// StatusInfo returns a human-readble version of the status for a Timex.
func (timex *Timex) StatusInfo() string {
	var status string
	switch timex.Status {
	case Ok:
		status = "Everything is ok."
	case Ins:
		status = "Positive leap second warning."
	case Del:
		status = "Negative leap second warning."
	case Oop:
		status = "Currently in a leap second."
	case Wait:
		status = "Leap second has occured."
	case Error:
		status = "Clock not synchronized."
	}
	return status
}
