(() => {
  chrome.windows.create({
      focused: true,
      url: 'templates/main.html',
      type: 'popup',
      width: 400, height: 600,
      // left: 100, top:100
    },
    (subWindow) => {
      chrome.windows.update(subWindow.id, {focused: true})
    })
})()
