const http = require('http')
const fs = require('fs')
const url = require('url')

const readCsv = (filename) => {
  const csv = fs.readFileSync(filename, 'utf-8')
  const data = csv.split('\n')
  const headers = data[0].split(',')
  // console.log(headers)

  const blogs = []

  for (let i = 1; i < data.length; i++) {
    const row = data[i].split(',')

    //  const response = {
    //    [headers[0]]: row[0],
    //    [headers[1]]: row[1],
    //    [headers[2]]: row[2],
    //  }

    const response = {
      id: row[0],
      title: row[1],
      author: row[2],
    }
    blogs.push(response)
  }
  return blogs
}

const handlerFunc = (req, res) => {
  const parseObj = url.parse(req.url, true)

  if (parseObj.pathname === '/') {
    console.time('test')
    res.writeHead(200, {
      'Content-Type': 'application/json',
    })
    const blogs = readCsv('test.csv')
    res.end(JSON.stringify(blogs))
    console.timeEnd('test')
  } else {
    res.writeHead(404, {
      'Content-Type': 'application/json',
    })
    res.end(JSON.stringify({ error: 'error' }))
  }
}

const server = http.createServer(handlerFunc)

server.listen(8080, () => console.log('server running on port 8080'))
