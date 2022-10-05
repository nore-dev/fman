package theme

type ThemeMap map[string]Theme

var Themes = ThemeMap{
	"default":    DefaultTheme,
	"darcula":    DefaultTheme,
	"brogrammer": BrogrammerTheme,
}
