const restify = require('restify');
const settings = require('./settings');
const crypto = require('crypto');
const appId = 'node.js';

/**
 * @returns {*} key, SHA256 value, like '16eb3373eadc234d6fbb86519909aba75ce842284d20c6abb8830b181c2857c8'
 */
function generateKey() {
  var sha = crypto.createHash('sha256');
  sha.update(Math.random().toString());
  return sha.digest('hex');
}


const apiClient = restify.createStringClient({
  url: process.env.API_URL
});

const client = restify.createStringClient({
  url: process.env.NEXT_APP_URL
});

const server = restify.createServer({
  name: settings.name,
  version: settings.version
});

server.use(restify.queryParser());

server.get('/', (req, res, next) => {
  res.send('hello from node.js');
  next();
});

function forward(handle, response, next) {
  console.log('forward to next app: ' + process.env.NEXT_APP_URL + '/start?handle=' + handle);
  // request to next app...
  client.get('/start?handle=' + handle, (err, req, res, data) => {
    if (err) return next(err);
    response.send(data);
    next();
  });
}

function getValue(key, cb) {
  apiClient.get('/' + key, (err, req, res, data) => cb(err, data));
}

function deleteValue(key, cb) {
  apiClient.del('/' + key, (err, req, res, data) => cb(err));
}

function setValue(key, val, cb) {
  apiClient.put('/' + key, 'data=' + val, (err, req, res, data) => cb(err));
}

server.get('/start', (req, res, next) => {
  console.log('called start route...');

  var handle = req.params.handle;

  if (!handle) {
    console.log('no handle... so let\'s start...');
    handle = generateKey();

    // write in redis with that handle as key...
    setValue(handle, appId, (err) => {
      if (err) {
        console.error(err);
        return next(err);
      }

      forward(handle, res, next);
    });
    return;
  }

  getValue(handle, (err, val) => {
    if (err) {
      console.error(err);
      return next(err);
    }

    val = val || '';

    // check if already seen
    if (val.indexOf(appId) >= 0) {
      console.log('round trip completed');
      // remove key from storage...
      deleteValue(handle, (err) => {
        if (err) {
          console.error(err);
          return next(err);
        }

        val += ';' + appId;

        res.send('Round trip completed!   >>>  ' + val.split(';').join(' => '));
        next();
      });
      return;
    }

    val += ';' + appId;

    console.log(handle + ' save new value ' + val);

    setValue(handle, val, (err) => {
      if (err) {
        console.error(err);
        return next(err);
      }

      console.log(handle + ' not already seen so forward...');

      forward(handle, res, next);
    });
  });
});

server.listen(settings.port, (err) => {
  if (err) return console.error(err);
  console.log('server listening at %s', server.url);
});
