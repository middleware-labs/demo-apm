from flask import Flask, request
from mw_tracker import MwTracker

tracker=MwTracker(
    access_token="yzofakgtpcelqsexhzhegmticbrciftarhzd"
)

app = Flask(__name__)

@app.route('/')
def hello_world():
      return 'Hello World!'

if __name__ == '__main__':
    app.run()

