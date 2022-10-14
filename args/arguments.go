package args

import "github.com/alexflint/go-arg"

// Define CommandLine arguments
var CommandLine struct {
	Path  string `arg:"positional" default:"."`
	Theme string `default:"default"`
	Icons string `default:"nerdfont"`
}

// Expose initialization function
func Initialize() {
	arg.MustParse(&CommandLine)
}
