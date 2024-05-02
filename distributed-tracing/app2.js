const tracker = require('@middleware.io/node-apm');
tracker.track({
    projectName: "node-app-a",
    serviceName: "node-app-a", 
});


tracker.info('Info sample');
tracker.warn('Warning sample');
tracker.debug('Debugging Sample');
tracker.error('Error Sample');

tracker.error(new Error('Error sample with stack trace'));

const http = require('http');

const hostname = '127.0.0.1';
const port = 3001;

const url = require('url');

const server = http.createServer((req, res) => {
    const path = url.parse(req.url, true).pathname;
    const id = path.split('/').pop();
    http.get("http://localhost:4001/123", (response) => {
      });
    if (id) {
      res.statusCode = 200;
      res.setHeader('Content-Type', 'text/plain');
      res.end(`${id}`);
      tracker.setAttribute("id traced",id)
      console.log(id)
    } else {
      res.statusCode = 400;
      res.setHeader('Content-Type', 'text/plain');
      res.end('ID parameter is missing');
    }
  });
server.listen(port, hostname, () => {
    console.log(`Server running at http://${hostname}:${port}`);
});

