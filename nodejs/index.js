const http = require('http')
const fs = require('fs')
const url = require('url')

function readCsv(filename) {
  const file = fs.readFileSync(filename, 'utf8')

  const lines = file.split('\n')
  const headers = lines[0].split(',')

  return lines.slice(1).map((line) => {
    const values = line.split(',')
    return headers.reduce((obj, header, index) => {
      obj[header] = values[index]
      return obj
    }, {})
  })
}

const requestListener = (req, res) => {
  const parseObj = url.parse(req.url, true)

  if (parseObj.pathname === '/') {
    console.time('test')
    const data = readCsv('test.csv')
    const responseData = []
    for (let i = 0; i < data.length; i++) {
      const element = data[i]
      const response = {
        id: element.ID,
        author: element.Author,
        title: element.Title,
      }
      responseData.push(response)
    }

    res.writeHead(200, {
      'Content-Type': 'application/json',
    })

    res.end(JSON.stringify(responseData))
    console.timeEnd('test')
  } else {
    res.writeHead(404, {
      'Content-Type': 'application/json',
    })
    res.end(JSON.stringify({ error: 'error' }))
  }
}

const server = http.createServer(requestListener)
server.listen(8080, () => console.log('server is listening on 8080'))
