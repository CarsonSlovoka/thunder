package main

import (
    log2 "github.com/CarsonSlovoka/thunder/pkg/log"
    "log"
    "os"
)

var (
    Trace *log.Logger
    Error *log.Logger
)

// log
func initLog(filepath string) *os.File {
    // Close the file until the program end (NOT THIS FUNCTION END). Otherwise, it will not write anything.
    file, err := os.OpenFile(
        filepath,
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666,
    )
    if err != nil {
        Trace, Error = log2.New(os.Stdout, os.Stderr, false)
        Error.Printf("Unable to create and/or open log file. Will log to Stdout and Stderr. Error: %v", err)
        return nil
    }
    Trace, Error = log2.New(file, file, false) // DO NOT CHANGE os.Stdout. Otherwise,  /dev/stdout: The pipe is being closed.
    return file
}
