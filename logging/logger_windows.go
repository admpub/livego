// +build windows

package logging

import (
	"log"
	"os"
)

func CrashLog(file string) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err.Error())
	} else {
		os.Stderr = f
		os.Stdout = f
	}
}
