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

func debug(m string) {
	if DEBUG_MODE {
		fmt.Println(m)
	}
}

func is_debug_event_enabled() bool {
	return DEBUG_MODE && DEBUG_EVENTS
}
