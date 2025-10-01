package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

// Go has tools for log output with the log package.
func logging() {
	// Standard fns will use the standard logger, 
	// which is preconfigured to write to stderr.
	log.Println("standard logger")

	// loggers can be configured with flags to adjust their output format
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro") // microsecond accuracy

	// can log the file/line from which the log was called
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line") // file and line number

	// custom logger with a prefix "my:"
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// can set prefix of existing logger
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// can have custom output targets: any io.Writer
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	// Write log output to buf
	buflog.Println("hello")
	// Read log output from buf 
	fmt.Print("from buflog:", buf.String())

	// can have structured log output with the slog pkg
	// f.e JSON format:
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	// Can have multiple key=value pairs, aside from the message:
	myslog.Info("hello again", "key", "val", "age", 25)
	// Can have different levels of severity:
	myslog.Error("oh no", "error", "something went wrong")
}