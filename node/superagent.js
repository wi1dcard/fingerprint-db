const superagent = require('superagent');
const fs = require('fs');

// callback
superagent
  .get('https://localhost:8443/json')
  .disableTLSCerts()
  .http2()
  .set('User-Agent', 'superagent')
  .end((err, res) => {
    fs.appendFileSync('fingerprint.json', res.text);
    process.stdout.write(res.text)
  });
