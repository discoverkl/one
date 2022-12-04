package one

import (
	"fmt"
	"os"
	"strings"
)

var dev bool

func init() {
	for _, kv := range os.Environ() {
		a := strings.SplitN(kv, "=", 2)
		key, value := a[0], a[1]
		if value == "1" {
			if key == "dev" || strings.HasPrefix(key, "dev_") {
				dev = true
				return
			}
		}
	}
}

func Dev() bool {
	return dev
}

func InDevMode(pkg string) bool {
	if !dev {
		return false
	}

	var key string
	if pkg != "" {
		key = fmt.Sprintf("dev_%s", pkg)
		value := os.Getenv(key)
		if value != "" {
			return value == "1"
		}
	}

	return os.Getenv("dev") == "1"
}
