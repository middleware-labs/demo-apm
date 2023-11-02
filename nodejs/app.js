const tracker = require('@middleware.io/node-apm');
// import tracker from '@middleware.io/node-apm';
tracker.track({
    serviceName: process.env.MW_SERVICE_NAME, 
    accessToken: process.env.MW_API_KEY,
    target: process.env.MW_TARGET,
  });
const http = require('http');

const hostname = '0.0.0.0';
const port = 8080;

const server = http.createServer((req, res) => {
    res.statusCode = 200;
    console.log("Hello World invoked")
    tracker.info("Hello World invoked")
    res.setHeader('Content-Type', 'text/plain');
    res.end('Hello World');
  });
  
server.listen(port, hostname, () => {
    console.log(`Server running at http://${hostname}:${port}/`);
});

