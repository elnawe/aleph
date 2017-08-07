package main

import (
	"fmt"
	"os"
)

func handle_error(e error, m string, s int) {
	if e != nil {
		fmt.Fprintf(os.Stdout, "ERROR: "+m+"\n", e)
		os.Exit(s)
	}
}

func debug(m string, o ...interface{}) {
	if DEBUG_MODE {
		format := "\t" + m + "\n"
		if len(o) > 0 {
			fmt.Printf(format, o)
		} else {
			fmt.Printf(format)
		}
	}
}

func is_debug_event_enabled() bool {
	return DEBUG_MODE && DEBUG_EVENTS
}
