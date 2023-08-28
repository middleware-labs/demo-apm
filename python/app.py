# You will need request from flask to monitor your traces
from flask import Flask, request
import logging
# Add these 3 lines as it is
from mw_tracker import MwTracker
tracker=MwTracker(
    access_token="yzofakgtpcelqsexhzhegmticbrciftarhzd"
)
tracker.collect_metrics()
tracker.collect_logs()
tracker.collect_profiling()

logging.info("Application initiated successfully.", extra={'Tester': 'Alex'})

app = Flask(__name__)

@app.route('/')
def hello_world():
    logging.error("error log sample", extra={'CalledFunc': 'hello_world'})
    logging.warning("warning log sample")
    logging.info("info log sample")
    return 'Hello World!'

@app.route('/exception')
def generate_exception():
    randomList = ['a', 0, 2]

    for entry in randomList:
        try:
            print("The entry is", entry)
            r = 1/int(entry)
            break
        except Exception as e:
            tracker.record_error(e)
    print("The reciprocal of", entry, "is", r)
    return 'Exception Generated!'

if __name__ == '__main__':
    app.run('0.0.0.0', 5000)
