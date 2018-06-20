package common

import (
	"io"
	"os"
	"log"
	"path"
)

const logDir = `./logs`

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func WriteLog(file string, v ...interface{}) {
	output, err := os.OpenFile(path.Join(logDir, file), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.MultiWriter(output, os.Stdout))
	defer output.Close()

	log.Println(v...)
}
