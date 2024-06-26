# Steps to Run the Project

Create and Activate Virtual Environment
```
python3 -m venv env
source env/bin/activate
```

Go to Project Root
```
cd demo
```

Install Dependencies
```
pip install -r requirements.txt
```

Run the Application
```
DJANGO_SETTINGS_MODULE='demo.settings' middleware-apm run gunicorn -c conf/gunicorn.conf.py --workers=4  --bind 0.0.0.0:8000 --timeout 120 demo.wsgi
```

`Note:  --workers and --timeout are optional`