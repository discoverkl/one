package one

import (
	"fmt"
	"os"
)

func InDevMode(pkg string) bool {
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
