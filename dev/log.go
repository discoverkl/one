package dev

import (
	"log"

	"github.com/discoverkl/one"
)

func Printf(format string, v ...any) {
	if !one.InDevMode("") {
		return
	}
	log.Printf(format, v...)
}

func Print(v ...any) {
	if !one.InDevMode("") {
		return
	}
	log.Print(v...)
}

func Println(v ...any) {
	if !one.InDevMode("") {
		return
	}
	log.Println(v...)
}
