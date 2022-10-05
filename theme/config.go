package theme

type ThemeMap map[string]Theme

var themes = ThemeMap{
	"default":    DefaultTheme,
	"dracula":    DefaultTheme,
	"brogrammer": BrogrammerTheme,
}

// Tries to match provided flag value for --theme against
// an existing ThemeMap and returns default theme if theme
// name does not match any records
// in the ThemeMap (due to a typo for example)
func GetActiveTheme(themeNameCandidate string) (theme Theme) {
	if _, ok := themes[themeNameCandidate]; ok {
		return themes[themeNameCandidate]
	}
	return DefaultTheme
}
