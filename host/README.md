# 此資料夾幹嘛的？

這是一個要放在使用者端的資料夾，用來告知chrome-extension的那些項目，其可以有能力傳送消息給本機端指定的程式

他必須要有一份[描述檔](https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Native_messaging#app_manifest) (json格式)

## 描述檔

### 檔名

描述檔的檔名是什麼不重要，建議可以像這樣命名:

- ``manifest.json``
- ``com.my_company.my_applicationr.json``

### 內容

```json5
{
  "name": "com.microsoft.browsercore", // 必須和機碼的資料夾同名才可以
  "description": "BrowserCore", // 隨意
  "path": "relative//coolApp.exe", // 通常和該描述檔，放在同一個資料夾 // 也可以放bat，用bat來開需要的檔案比較方便
  "type": "stdio", // 目前也只有stdio可以用 // Type of the interface used to communicate with the native messaging host. Currently there is only one possible value for this parameter: stdio. It indicates that Chrome should use stdin and stdout to communicate with the host. // https://developer.chrome.com/docs/apps/nativeMessaging/#native-messaging-host
  "allowed_origins": [ // 有哪些app可以有權限用
    "chrome-extension://ppnbnpeolgkicgegkbkbjmhlideopiji/",
    "chrome-extension://ndjpnladcallmjemlbaebfadecfhkepb/"
  ]
}
```

## [Manifest location/安裝](https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Native_manifests#manifest_location)

- [Windows](https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Native_manifests#windows)

  通常就這幾個地方，建議寫NativeMessagingHosts即可

  ```
  HKEY_LOCAL_MACHINE\SOFTWARE\Mozilla\NativeMessagingHosts\<name>

  HKEY_LOCAL_MACHINE\SOFTWARE\Mozilla\ManagedStorage\<name>

  HKEY_LOCAL_MACHINE\SOFTWARE\Mozilla\PKCS11Modules\<name>
  ```

- [macOS](https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Native_manifests#macos)
- [Linux](https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Native_manifests#linux)

```
REG ADD "HKCU\Software\Google\Chrome\NativeMessagingHosts\com.google.chrome.extension.thunder" ^  // 記得最後的名稱必須和json檔"裡面"的name一致
/ve /t REG_SZ ^
/d "%~dp0manifest.json"  可以善用dp0表示當前的資料夾的位置，後面放您的json檔名稱，名稱隨意
```

## [chrome-extension: js實作](https://developer.chrome.com/docs/extensions/mv3/messaging/#connect)

這部份是chrome的擴充功能要如何寫才可以啟動

```js
chrome.runtime.sendNativeMessage("com.my_company.my_application",
    {key1: "value1", key2: "value2"}, // 👈 Send those parameters to your program.
    (response) => {
        console.log(response)
    }
)
```
