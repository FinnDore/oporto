package main

import (
	toml "github.com/BurntSushi/toml"
	xdg "github.com/adrg/xdg"
	"os"
	"path"
	debug "runtime/debug"
)

type Enviorment struct {
	SshCommand  string
	StaticHosts []string
}

type Config struct {
	Enviroments map[string]Enviorment
}

func LoadConfig() (Config, error) {
	println("loading config")
	var configFileName = "oporto.toml"
	var configFolder = path.Join(xdg.Home, ".config")
	var configPath = path.Join(configFolder, configFileName)
	var config Config
	_, err := toml.DecodeFile(configPath, &config)
	if err == nil {
		return config, err
	}

	if os.IsNotExist(err) {
		err = os.MkdirAll(configFolder, 600)
		if err != nil {
			return config, err
		}
		file, err := os.Create(configPath)
		if err != nil {
			return config, err
		}
		file.Close()
		println("created config file")
		return LoadConfig()
	}
	return config, err
}

// var conf Config
// _, err := toml.Decode(tomlData, &conf)

var Commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value[0:6]
			}
		}
	}

	return "unknown"
}()
