// +build !windows

package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"syscall"
	"time"
)

func CrashLog(file string) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err.Error())
	} else {
		syscall.Dup2(int(f.Fd()), 2)
	}
}
