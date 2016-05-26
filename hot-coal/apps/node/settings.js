const pkg = require('./package.json');
const cfenv = require('cfenv');
const appEnv = cfenv.getAppEnv();

var settings = {
  name: pkg.name,
  version: pkg.version
};

if (appEnv.isLocal) {
  settings.port = 3000;
  process.env.NEXT_APP_URL = 'http://localhost:' + settings.port;
  process.env.API_URL = 'http://localhost:5000';
} else {
  settings.port = appEnv.port;
}

module.exports = settings;
