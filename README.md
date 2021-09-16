<p align="center">
  <a href="">
    <img alt="thunder" src="asset/img/site/favicon.svg" width="384"/>
  </a>
</p>

<p align="left">
  <a href="https://godoc.org/github.com/CarsonSlovoka/thunder/go" title="GoDoc">
    <img alt="go doc" src="https://godoc.org/code.gitea.io/gitea?status.svg">
  </a>
  <a href="https://opensource.org/licenses/MIT" title="License: MIT">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-blue.svg?style=plastic">
  </a>
  <br>
  <a href="">
  <img src="https://img.shields.io/static/v1?&style=plastic&logo=GO&label=Lang&message=go&color=66FFFF"/>
  </a>
  <a href="">
  <img src="https://img.shields.io/static/v1?&style=plastic&logo=javascript&label=Lang&message=js&color=FFFF33"/>
  </a>
  <br>
  <a href="">
  <img src="https://img.shields.io/static/v1?&style=plastic&logo=google&label=browser&message=Chrome&color=FFFF33"/>
  </a>
</p>

## Purpose

Execute the program by the command from the input parameter coming from the chrome extension app, and send the response
to it.

## 🤔 How to do it?

Use the [Native messaging host](https://developer.chrome.com/docs/apps/nativeMessaging/#native-messaging-host) to
communicate the native program


## Reference

- [Start an external application from a Google Chrome Extension?](https://stackoverflow.com/a/69188382/9935654)
- [GoogleChrome/chrome-extensions-samples](https://github.com/GoogleChrome/chrome-extensions-samples/tree/475b4b6/mv2-archive/api/nativeMessaging/host)
- [jfarleyx/chrome-native-messaging-golang](https://github.com/jfarleyx/chrome-native-messaging-golang/blob/95c4d57/native-host/src/main.go)
