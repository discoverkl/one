package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/discoverkl/one/host"
	"gopkg.in/yaml.v3"
)

// AutoConfigAndExit try to load saved config,
// then return it or use 'saveFn' to update and save it to disk.
// In update mode, this function will not return to its caller.
func AutoConfigAndExit[T any](appName string, saveFn func(savedConfig *T)) *T {
	// if update is true, always exit program before return
	update := (len(os.Args) >= 2 && (os.Args[1] == "config" || os.Args[1] == "c"))
	printAndExit := (update && len(os.Args) == 2)

	var conf T
	configPath, err := host.HomeExpand(fmt.Sprintf("~/.%s/config.yml", appName))
	if err != nil {
		if !update {
			return &conf
		}
		log.Fatalf("detect home dir failed: %v", err)
	}

	// try to load saved config, ignore any error
	savedConfig, err := loadConfig[T](configPath)
	if err == nil {
		conf = *savedConfig
	} else {
		if _, err2 := os.Stat(configPath); err2 != nil && os.IsNotExist(err2) {
			// file not exit, do nothing
		} else {
			// warn user
			log.Printf("load config file failed: %v", err)
		}
	}

	// return saved config or empty config
	if !update {
		return &conf
	}

	//
	// update config and exit
	//

	if printAndExit {
		data, err := yaml.Marshal(conf)
		if err != nil {
			log.Fatalf("invalid config format: %v", err)
		}
		fmt.Println(strings.TrimRight(string(data), "\n"))
		os.Exit(0)
	}

	// use callback to update saved config
	if saveFn == nil {
		log.Fatal("invalid config command")
	}
	saveFn(&conf)

	// create and update config file
	dirname := filepath.Dir(configPath)
	err = os.MkdirAll(dirname, 0755)
	if err != nil {
		log.Printf("make dir '%s' failed: %v", dirname, err)
	}

	data, err := yaml.Marshal(conf)
	if err != nil {
		log.Fatalf("marshal config failed: %v", err)
	}
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		log.Fatalf("save config failed: %v", err)
	}

	os.Exit(0)
	panic("should never be here")
}

func loadConfig[T any](configPath string) (*T, error) {
	var conf T
	// load config and continue
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
