from middleware import MwTracker
tracker = MwTracker()
from dotenv import load_dotenv
load_dotenv()
import logging
import os



print("\n----------DEBUGGING INFO----------")
print(os.environ.get('MW_TARGET'))
print(tracker.access_token)
print(tracker.service_name)
print("----------DEBUGGING INFO----------\n")

from typing import Union

from fastapi import FastAPI

app = FastAPI()




@app.get("/")
def read_root():
    logging.info("Hello World")
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    logging.debug("Item ID: %s", item_id)
    return {"item_id": item_id, "q": q}


if __name__ == "__main__":
    import uvicorn

    uvicorn.run('main:app', reload=True)