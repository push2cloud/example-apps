# hot-coal example

roundtrip for all apps

example: https://hot-coal-node-space-org.scapp.io/start

will do a roundtrip node.js -> go -> python -> node.js


## blueprint

/ route returns: 'hello from language'

/start route:

1. check if optional query param `handle` is present
2. if not present:

  generate a new handle (i.e. random sha256 string)

  make a put call to api app to save the passed appId(language string) value do the db

  forward to next app: `https://url-of-next-app/start?handle=1234`
3. if present:

  make a get call to api app to check if there is something saved for that handle

  if own appId(language string) is there the roundtrip is completed, make a delete call to api app and respond i.e. with: `Round trip completed!   >>>  node.js -> go -> python -> node.js)`

  if own appId(language string) is not there, attach own appId to the value ( `appId += ';' appId` ) and make a put call to api app to save it do the db and finally forward to next app: `https://url-of-next-app/start?handle=1234`
