# inspired by https://github.com/ihuston/python-cf-examples/blob/master/03-services-redis/main.py

from flask import Flask
from flask import request

import os
import json
import random
import requests
import string
import hashlib
import time

app = Flask(__name__)

# Get port from environment variable or choose 9099 as local default
port = int(os.getenv("PORT", 9099))

# Get next app:
if 'NEXT_APP_URL' in os.environ:
    next_app_url = os.getenv('NEXT_APP_URL')
else:
    next_app_url = 'http://localhost:%i/start?handle=' % port

# Get api_url
if 'API_URL' in os.environ:
    api_url = os.getenv('API_URL')
else:
    api_url = 'http://localhost:8080/'


@app.route('/')
def home():
        return 'Hello from Python!', 200

@app.route('/start')
def start():
    handle = request.args.get('handle')
    print('handle provided via request: %s' % handle)
    if not handle:
        handle = hashlib.sha256(str(random.getrandbits(256)).encode('utf-8')).hexdigest()
        print('generated random handle: %s' % handle)
        history = []

    else:
        response = requests.get(api_url + '/' + handle)
        print('retrieving...')
        print(response.content)
        if response.status_code != 404:
            history = response.content.decode().split(';')
        else:
            history = []
    print('history:')
    print(history)
    if 'python' in history:
        print('found myself in history!')
        resp = requests.delete(api_url + '/' + handle)
        if resp.status_code != 200:
            print(resp)
        return 'Round trip completed!    >>>   ' + ' => '.join(history) + ' => python', 200

    history.append('python')
    print('putting %s to handle %s' % (';'.join(history), handle))
    resp = requests.put(api_url + '/' + handle, params={'data': ';'.join(history)})
    if resp.status_code != 200:
        print(resp)
        print(resp.content)

    url = next_app_url + '/start?handle=' + handle
    print('connecting to next app via: %s' % url)
    response = requests.get(url)
    return response.content, response.status_code, {'Content-Type': 'text/css; charset=utf-8'}

if __name__ == '__main__':
    # Run the app, listening on all IPs with our chosen port number
    app.run(host='0.0.0.0', port=port, threaded=True)
