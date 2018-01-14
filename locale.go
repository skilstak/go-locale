package locale

import (
	"os/exec"
	"strings"
	"time"
)

var Default string = "en_US"

func run(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return Default, err
	}
	return strings.TrimSpace(string(out)), nil
}

func Now() string {
	var f string
	switch Detect() {
	case "en_US":
		f = "3:04 PM, Monday, January 2, 2006"
	}
	return time.Now().Format(f)
}
