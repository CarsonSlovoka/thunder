const NativeName = "com.google.chrome.extension.thunder";

(() => {
  window.onload = () => {
    const log = document.querySelector(`#log`)
    const input = document.querySelector(`input`)
    let filepath = ""
    input.onchange = () => filepath = input.value

    const btnLongConn = document.querySelector(`#btn-long-conn`)
    const connArea = btnLongConn.parentElement.querySelector(`fieldset`)
    const btnPostMsg = document.querySelector(`#btn-post-msg`)
    const btnDisconnect = document.querySelector(`#btn-disconnect`)
    let port = undefined

    btnLongConn.onclick = () => {
      //Continuous connection, unless the "disconnect" method is called.

      btnLongConn.disabled = true
      connArea.disabled = false

      port = chrome.runtime.connectNative(NativeName)

      port.onMessage.addListener((response) => {
        //console.log("Received: " + response)
        log.innerHTML += "Received:<br>" +  JSON.stringify(response) + `<br>`
      })

      port.onDisconnect.addListener(() => {
        if (chrome.runtime.lastError) {
          console.error(JSON.stringify(chrome.runtime.lastError))
        }
      })
    }

    btnPostMsg.onclick = () => {
      if (port === undefined) {
        return
      }
      port.postMessage({action: "exec", path: filepath})
    }

    btnDisconnect.onclick = () => {
      port.disconnect(NativeName)
      port = undefined
      btnLongConn.disabled = false
      connArea.disabled = true
    }

    document.querySelector(`#btn-short-conn`).onclick = () => {
      // Send one message at once time.
      chrome.runtime.sendNativeMessage(NativeName,
        {action: "exec", path: filepath},
        (response) => {
          log.innerHTML += "Received:<br>" +  JSON.stringify(response) + `<br>`
        }
      )
    }
  }
})()
