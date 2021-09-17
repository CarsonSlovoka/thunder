package log

import (
    "io"
    "log"
    "os"
)

var (
    OKPrefixString  string
    ErrPrefixString string
)

func init() {
    OKPrefixString = "Log:"
    ErrPrefixString = "Error:"
}

func New(logWriter io.Writer, errWriter io.Writer, stdAlso bool) (*log.Logger, *log.Logger) {
    createLogger := func(prefixString string, writer ...io.Writer) *log.Logger {
        multipleWriter := io.MultiWriter(writer...)
        return log.New(multipleWriter, prefixString, log.Ldate|log.Ltime|log.Lshortfile)
    }
    logWriters := map[bool][]io.Writer{true: {logWriter, os.Stdout}, false: {logWriter}}[stdAlso]
    logLogger := createLogger(OKPrefixString, logWriters...)

    errWriters := map[bool][]io.Writer{true: {errWriter, os.Stdout}, false: {errWriter}}[stdAlso]
    errLogger := createLogger(ErrPrefixString, errWriters...)
    return logLogger, errLogger
}
