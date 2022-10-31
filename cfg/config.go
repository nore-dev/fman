package cfg

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nore-dev/fman/args"
)

const (
	// Config Metadata
	XdgConfigDir       = "/.config/"
	FmanConfigDir      = "/.config/fman/"
	FmanConfigFileName = "config.toml"

	// Config Defaults
	DefaultTheme = "default"
	DefaultIcons = "nerdfont"
)

var Config = &Cfg{}

type Cfg struct {
	Theme string
	Icons string
}

func LoadConfig() error {
	// Load file if exists
	err := loadConfigFile()

	// If config file exists and cli args are provided use args from cli
	if args.CommandLine.Theme != "" {
		Config.Theme = args.CommandLine.Theme
	}
	if args.CommandLine.Icons != "" {
		Config.Icons = args.CommandLine.Icons
	}

	// For each config value if neither config file or cli args provides a value
	// then use default values to fill config object.
	if Config.Theme == "" {
		Config.Theme = DefaultTheme
	}
	if Config.Icons == "" {
		Config.Icons = DefaultIcons
	}

	return err
}

func loadConfigFile() error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	fileContents, err := os.ReadFile(filepath.Join(currentUser.HomeDir, FmanConfigDir, FmanConfigFileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New("no config file found")
		} else {
			return errors.New("could not read config file")
		}
	}
	_, err = toml.Decode(string(fileContents), Config)
	if err != nil {
		return errors.New("could not decode config file")
	}

	return nil
}
