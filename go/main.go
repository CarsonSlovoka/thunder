/**
* @Author: Carson Tseng
* @Date: 2021-09-14
**/

package main

func main() {
    file := initLog("chrome-connect-native--thunder.log")
    if file != nil {
        defer func() {
            Trace.Printf("Exit App.")
            _ = file.Close()
        }()
    }
    LoadConfig()
    HandleMessage(4196)
}
