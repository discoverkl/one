package dev

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/discoverkl/one"
)

func Printf(format string, v ...any) {
	if !callerInDevMode() {
		return
	}
	log.Printf(format, v...)
}

func Print(v ...any) {
	if !callerInDevMode() {
		return
	}
	log.Print(v...)
}

func Println(v ...any) {
	if !callerInDevMode() {
		return
	}
	log.Println(v...)
}

func callerInDevMode() bool {
	if !one.Dev() {
		return false
	}
	pkg := callerPackage(3)
	return one.InDevMode(pkg)
}

// callerPackage return caller package's short name.
func callerPackage(skip int) string {
	var pkg string
	pc, _, _, _ := runtime.Caller(skip)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	size := len(parts)
	if parts[size-2][0] == '(' { // method
		pkg = parts[size-3]
	} else { // func
		pkg = parts[size-2]
	}
	return filepath.Base(pkg)
}
