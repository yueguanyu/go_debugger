package go_debugger

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

// Debugger Debugger
type Debugger struct {
	env    string
	prefix string
	color  int16
	active bool
}

// Log Log
func (D *Debugger) Log(str string, payload interface{}) {
	if !D.active {
		return
	}
	fmt.Printf(`%c[0;0;%dm%s%c[0m %s%v`+"\n", 0x1B, D.color, D.prefix, 0x1B, str, payload)
}

// Debug Debug
func Debug(prefix string) func(string, interface{}) {
	match := false
	debugEnv := os.Getenv("DEBUG")
	if debugEnv != "" {
		debugEnv = strings.ReplaceAll(debugEnv, "*", ".*")
		r := regexp.MustCompile(debugEnv)
		match = r.MatchString(prefix)
	}
	D := Debugger{os.Getenv("DEBUG"), prefix, int16(rand.Intn(7) + 30), match}
	return D.Log
}
