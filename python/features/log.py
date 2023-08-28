import logging
from mw_tracker import MwTracker

tracker=MwTracker(
    access_token="yzofakgtpcelqsexhzhegmticbrciftarhzd"
)
tracker.collect_logs()

logging.info("Hello World!", extra={'Key': 'Value'})


