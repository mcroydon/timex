// Show the current time, clock status, maximum and estimated clock errors (in nanoseconds).
package main

import (
	"fmt"
	"github.com/mcroydon/timex"
)

func main() {
	tmx := timex.Now()
	fmt.Printf("Current time: %v\nStatus: %v\nMax error: %v\nEstimated error: %v\n", tmx.Time, tmx.StatusInfo(), tmx.Maxerror, tmx.Esterror)
}
