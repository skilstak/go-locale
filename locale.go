package locale

import (
	"os/exec"
	"strings"
)

var Default string = "en_US"

func run(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return Default, err
	}
	return strings.TrimSpace(string(out)), nil
}
