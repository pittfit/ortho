package tracing

import (
	"fmt"
	"strings"
)

var level = 0
var enabled = false

func indentLevel() string {
	return strings.Repeat("\t", level)
}

// Trace …
func Trace(f func()) {
	prev := enabled
	enabled = true
	f()
	enabled = prev
}

// Enable …
func Enable() {
	enabled = true
}

// Disable …
func Disable() {
	enabled = false
}

func print(prefix, id, extra string) {
	if !enabled {
		return
	}

	fmt.Printf("%s\x1b[32m%s \x1b[33m%s \x1b[0m%s\n", indentLevel(), prefix, id, extra)
}

func inc() { level = level + 1 }
func dec() { level = level - 1 }

// Begin …
func Begin(id, extra string) (string, string) {
	print("BEGIN", id, extra)
	inc()
	return id, extra
}

// Call …
func Call(id, extra string) {
	print("CALL", id, extra)
}

// Debug …
func Debug(id, extra string) {
	print("DEBUG", id, extra)
}

// End …
func End(id, extra string) {
	dec()
	print("END", id, extra)
}
