const http2 = require('http2');
const fs = require('fs');

const client = http2.connect('https://localhost:8443', {
  rejectUnauthorized: false,
  settings: http2.getDefaultSettings(),
});
client.on('error', (err) => console.error(err));

const req = client.request({
  ':path': '/json',
  'user-agent': 'http2',
});

req.setEncoding('utf8');
let data = '';
req.on('data', (chunk) => { data += chunk; });
req.on('end', () => {
  fs.appendFileSync('fingerprint.json', data);
  process.stdout.write(data);
  client.close();
});
req.end();
