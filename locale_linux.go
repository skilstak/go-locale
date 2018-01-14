package locale

import (
	"strings"
)

func Detect() string {
	out, err := run("locale")
	if err != nil {
		return Default
	}
	for _, line := range strings.Split(out, "\n") {
		if line != "" {
			p := strings.Split(line, "=")
			v := strings.Trim(p[1], `"`)
			// sample:
			// LANG="en_US.UTF-8"
			// LC_COLLATE="en_US.UTF-8"
			// LC_CTYPE="en_US.UTF-8"
			// LC_MESSAGES="en_US.UTF-8"
			// LC_MONETARY="en_US.UTF-8"
			// LC_NUMERIC="en_US.UTF-8"
			// LC_TIME="en_US.UTF-8"
			// LC_ALL=
			if v == "C" {
				return v
			}
			if v != "" {
				i := strings.Index(v, ".")
				if i != -1 {
					v = v[0:i]
				}
				return v
			}
		}
	}
	return Default
}
