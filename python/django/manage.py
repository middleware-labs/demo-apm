#!/usr/bin/env python
"""Django's command-line utility for administrative tasks."""
import os
import sys
import logging 
 
from middleware import MwTracker
tracker=MwTracker()

def main():
    logging.info("testing info log for Django")
    os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'helloworld.settings')
    try:
        from django.core.management import execute_from_command_line
    except ImportError as exc:
        raise ImportError(
            "Couldn't import Django. Are you sure it's installed and "
            "available on your PYTHONPATH environment variable? Did you "
            "forget to activate a virtual environment?"
        ) from exc
    execute_from_command_line(sys.argv)


if __name__ == '__main__':
    main()
