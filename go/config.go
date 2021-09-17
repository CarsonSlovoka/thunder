package main

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

var (
    Config Settings
)

type Settings struct {
    BufferSize int   `json:"buffer-size"`
    Debugs     Debug `json:"debugs"`
}

type Debug struct {
    Enable bool `json:"enable"`
}

func LoadConfig() {
    jsonFile, err := os.Open("manifest.json")
    if err != nil {
        // _, _ = os.Stderr.Write([]byte(err.Error()))
        Error.Printf(`Unable to open the "manifest.json": %v`, err)
    }

    byteValues, _ := ioutil.ReadAll(jsonFile)
    if err := json.Unmarshal(byteValues, &Config); err != nil {
        // _, _ = os.Stderr.Write([]byte(err.Error()))
        Error.Printf("Unable to unmarshal json to struct (Settings): %v", err)
    }
}
