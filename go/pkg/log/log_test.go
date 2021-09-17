package log

import (
    "log"
    "os"
    "testing"
)

var (
    Trace *log.Logger
    Error *log.Logger
)

func TestNew(t *testing.T) {
    file, err := os.OpenFile(
        "chrome-connect-native-log.txt",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666,
    )
    if err != nil {
        Trace, Error = New(os.Stdout, os.Stderr, false)
        Error.Printf("Unable to create and/or open log file. Will log to Stdout and Stderr. Error: %v", err)
        return
    }
    defer func() { _ = file.Close() }()

    Trace, Error = New(file, file, true)
    Trace.Printf("OK")
    Error.Printf("No error")
}
