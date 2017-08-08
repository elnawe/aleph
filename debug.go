package main

import (
	"fmt"
	"os"
)

func handle_error(e error, m string, p string, s int) {
	if e != nil {
		fmt.Fprintf(os.Stdout, "[%s] ERROR: "+m+"\n", s, e)
		os.Exit(s)
	}
}

func debug(m string, c string, o ...interface{}) {
	if DEBUG_MODE {
		format := "[%s]  " + m + "\n"
		if len(o) > 0 {
			fmt.Printf(format, c, o)
		} else {
			fmt.Printf(format, c)
		}
	}
}

func is_debug_event_enabled() bool {
	return DEBUG_MODE && DEBUG_EVENTS
}
