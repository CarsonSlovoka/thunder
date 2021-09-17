const INCLUDE_TAG_NAME = `data-include-html`

/**
 * @param {Element} node
 * @param {Function} cb callback
 * */
export async function includeHTML(node, {
  cb = undefined
}) {
  const nodeArray = node === undefined ?
    document.querySelectorAll(`[${INCLUDE_TAG_NAME}]`) :
    node.querySelectorAll(`[${INCLUDE_TAG_NAME}]`)

  if (nodeArray === null) {
    return
  }

  for (const node of nodeArray) {
    const filePath = node.getAttribute(`${INCLUDE_TAG_NAME}`)
    if (filePath === undefined) {
      return
    }

    await new Promise(resolve => {
      fetch(filePath
      ).then(async response => {
          const text = await response.text()
          if (!response.ok) {
            throw Error(`${response.statusText} (${response.status}) | ${text} `)
          }
          node.innerHTML = text
          const rootPath = filePath.split("/").slice(0, -1)
          node.querySelectorAll(`[${INCLUDE_TAG_NAME}]`).forEach(elem => {
            const relativePath = elem.getAttribute(`${INCLUDE_TAG_NAME}`) // not support ".."
            if (relativePath.startsWith('/')) { // begin with site root.
              return
            }
            elem.setAttribute(`${INCLUDE_TAG_NAME}`, [...rootPath, relativePath].join("/"))
          })
          node.removeAttribute(`${INCLUDE_TAG_NAME}`)
          await includeHTML(node, {cb})
          node.replaceWith(...node.childNodes)
          resolve()
        }
      ).catch(err => {
        node.innerHTML = `${err.message}`
        resolve()
      })
    })
  }

  if (cb) {
    cb()
  }
}
