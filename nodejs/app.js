const tracker = require('@middleware.io/node-apm');
tracker.track({
    projectName: "nodejs-webserver",
    serviceName: "nodejs-webserver", 
    accessToken: "nacwojkbrtcqzhuwblmnxcazeyyzlcrywdoz", 
  });

tracker.info('Info sample');
tracker.warn('Warning sample');
tracker.debug('Debugging Sample');
tracker.error('Error Sample');

tracker.error(new Error('Error sample with stack trace'));

const http = require('http');

const hostname = '0.0.0.0';
const port = 3000;

const server = http.createServer((req, res) => {
    res.statusCode = 200;
    tracker.info('req received');
    res.setHeader('Content-Type', 'text/plain');
    res.end('Hello World');
  });
  
server.listen(port, hostname, () => {
    console.log(`Server running at http://${hostname}:${port}/`);
});

