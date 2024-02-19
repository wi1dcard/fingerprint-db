const https = require('node:https');
const fs = require('node:fs');

var options = {
  host: 'localhost',
  port: '8443',
  path: '/json',
  headers: {
    "user-agent": "https",
  },
  rejectUnauthorized: false,
}

https.get(options, (res) => {
  res.on('data', (d) => {
    fs.appendFileSync('fingerprint.json', d);
    process.stdout.write(d);
  });

}).on('error', (e) => {
  console.error(e);
});
