package locale

func Detect() string {
	loc, err := run("defaults", "read",
		"/Library/Preferences/.GlobalPreferences",
		"AppleLocale",
	)
	if err != nil {
		return Default
	}
	return loc
}
