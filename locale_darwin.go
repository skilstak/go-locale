package locale

func Detect() (string, error) {
	return run("defaults", "read",
		"/Library/Preferences/.GlobalPreferences",
		"AppleLocale",
	)
}
