/*
You can reference the below link to get more details.
> https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Native_messaging#app_side
*/

package main

import (
    "bufio"
    "bytes"
    "encoding/binary"
    "encoding/json"
    "fmt"
    "io"
    "os"
    "unsafe"
)

var (
    nativeEndian binary.ByteOrder
)

type UserInputParam struct {
    Action string `json:"action"`
    Path   string `json:"path"`
}

type Response struct {
    Status Status `json:"status"`
    Msg    string `json:"msg"`
}

type Status struct {
    Code int    `json:"code"`
    Text string `json:"text"`
}

func init() {
    // determine native byte order so that we can read message size correctly
    var one int16 = 1
    b := (*byte)(unsafe.Pointer(&one))
    if *b == 0 {
        nativeEndian = binary.BigEndian
    } else {
        nativeEndian = binary.LittleEndian
    }
}

func HandleMessage(bufferSize int) string {
    /*
       bufferSize 4196 (default
    */

    stdInReader := bufio.NewReader(os.Stdin)
    if bufferSize > 0 {
        stdInReader = bufio.NewReaderSize(stdInReader, bufferSize) // adjust the size
    }
    if Config.Debugs.Enable {
        Trace.Printf("IO buffer reader created with buffer size of %v", stdInReader.Size())
    }

    lenBytes := make([]byte, 4)
    msgLen := 0
    for n, err := stdInReader.Read(lenBytes); // The data from "chrome.runtime.sendNativeMessage" 4 bytes represents the data length.
        n > 0 && err == nil;
    n, err = stdInReader.Read(lenBytes) {

        msgLen = readMsgLength(lenBytes, nativeEndian)

        if msgLen > Config.BufferSize {
            Error.Printf("Message size of %d exceeds buffer size of %d. Message will be truncated and is unlikely to unmarshal to JSON.", msgLen, Config.BufferSize)
            break
        }

        // read the content of the message from buffer
        contentBytes := make([]byte, msgLen)
        _, err := stdInReader.Read(contentBytes)
        if err != nil && err != io.EOF {
            Error.Fatal(err)
        }

        var userInPara UserInputParam
        parseMsg(contentBytes, &userInPara)

        response := runAction(&userInPara)
        sendMessage(response)
    }

    return ""
}

func readMsgLength(msg []byte, byteOrder binary.ByteOrder) int {
    var length uint32
    buf := bytes.NewBuffer(msg)
    err := binary.Read(buf, byteOrder, &length)
    if err != nil {
        Error.Printf("Unable to read bytes representing message length: %v", err)
    }
    return int(length)
}

func parseMsg(msg []byte, userInPara *UserInputParam) {
    decodeMessage(msg, userInPara)
    if Config.Debugs.Enable {
        Trace.Printf("Message received: %#v", userInPara)
    }
}

func decodeMessage(msg []byte, output interface{}) {
    if err := json.Unmarshal(msg, &output); err != nil {
        Error.Printf("Unable to unmarshal json to struct: %v", err)
    }
}

func runAction(para *UserInputParam) *Response {
    var response Response
    switch action := para.Action; action {
    case "exec":
        response.Status.Code = 200
        response.Status.Text = fmt.Sprintf("Running successfully. %s", action)
        /*
           ...
        */
    case "query":
    default:
        response.Status.Code = 501
        response.Status.Text = fmt.Sprintf("Not Implemented %s", action)
    }
    return &response
}

func sendMessage(response *Response) {
    byteMsg := data2bytes(response)
    // Trace.Printf("byteMsg %s", string(byteMsg))

    // Write msg to a buffer
    var msgBuf bytes.Buffer
    _, err := msgBuf.Write(byteMsg)
    if err != nil {
        Error.Printf("Unable to write message length to message buffer: %v", err)
    }

    // We are going to write data to Stdout, so we should tell it what the length is.
    writeMessageLength(os.Stdout, byteMsg)

    // Finally, write the data to Stdout from buffer.
    _, err = msgBuf.WriteTo(os.Stdout)
    if err != nil {
        Error.Printf("Unable to write message buffer to Stdout: %v", err)
    }
}

func data2bytes(data *Response) []byte {
    byteData, err := json.Marshal(data)
    if err != nil {
        Error.Printf(`Unable to marshal "Response" struct to slice of bytes: %v`, err)
    }
    return byteData
}

func writeMessageLength(writer io.Writer, msg []byte) {
    err := binary.Write(writer, nativeEndian, uint32(len(msg)))
    if err != nil {
        Error.Printf("Unable to write message length to Stdout: %v", err)
    }
}
