# Timex

A simple go package for accessing current time, clock status, maximum and estimated error. This is accomplished by
calling `ntp_gettime` from `sys/timex.h` via cgo. Timex should work on any system with `sys/timex.h` including modern
Linux and BSD distributions but not Mac OS X.

# Status

Timex is my first cgo package and as such there may be errors, leaks, or other bad behavior. It also may not be
altogether idomatic. Please open an issue or get in touch if you see anything that could use improvement.  Also,
you should probably use [Adjtimex in the syscall package](http://golang.org/pkg/syscall/#Adjtimex) instead:

```
package main

import (
        "log"
        "syscall"
)

func main() {
        var timex = new(syscall.Timex)
        status, err := syscall.Adjtimex(timex)
        if err != nil {
                log.Fatalf("Adjtimex error: %v", err)
        }
        if status != 0 {
                log.Fatalf("Unexpected syscall status %v", status)
        }
        log.Printf("time: %v max: %v est: %v status: %v\n", timex.Time, timex.Maxerror, timex.Esterror, timex.Status)
}
```

# Installation and use

Timex can be installed using go get:

```
$ go get github.com/mcroydon/timex
```

From there you can import and use timex:

```go
import (
    "fmt"
    "github.com/mcroydon/timex"
)
tmx := timex.Now()
fmt.Printf("Current time is %v and status is %v", tmx.Time, tmx.StatusInfo())
```

You can also install and use timexinfo:

```
$ go install github.com/mcroydon/timex/timexinfo
$ timexinfo
Current time: 2014-02-23 05:38:34.777437699 +0000 UTC
Status: Everything is ok.
Max error: 224774
Estimated error: 24376
```

# Testing

Run the tests:

```
$ go test github.com/mcroydon/timex
```

Run the tests and a benchmark that keeps track of the min, max, and mean values of timex.Now():

```
$ go test -bench . github.com/mcroydon/timex
````

To see the current values of `timex.Now()` run `go test` with `-v`.

# License

Released under a 3-clause BSD license.  The effective license may change based on the license of sys/timex.h and
related files on your system. Maybe.