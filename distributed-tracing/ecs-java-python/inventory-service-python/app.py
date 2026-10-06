import logging
import random
import time

from flask import Flask, jsonify

logging.basicConfig(level=logging.INFO, format="%(asctime)s %(levelname)s %(name)s %(message)s")
log = logging.getLogger("inventory-service")

app = Flask(__name__)

PRODUCTS = {
    "P100": {"name": "Mechanical Keyboard", "price": 89.99, "stock": 42},
    "P200": {"name": "Wireless Mouse", "price": 29.50, "stock": 120},
    "P300": {"name": "4K Monitor", "price": 329.00, "stock": 7},
    "P400": {"name": "USB-C Dock", "price": 149.00, "stock": 0},
    "P500": {"name": "Noise Cancelling Headset", "price": 199.00, "stock": 15},
}

REVIEWS = {
    "P100": [{"stars": 5}, {"stars": 4}, {"stars": 5}],
    "P200": [{"stars": 3}, {"stars": 4}],
    "P300": [{"stars": 5}],
    "P400": [],  # newly listed, no reviews yet
    "P500": [],  # newly listed, no reviews yet
}


def average_rating(reviews):
    if not reviews:
        return 0
    return round(sum(r["stars"] for r in reviews) / len(reviews), 2)


@app.get("/health")
def health():
    return {"status": "ok"}


@app.get("/api/products/<product_id>")
def get_product(product_id):
    product = PRODUCTS.get(product_id)
    if product is None:
        return jsonify({"error": "product not found", "productId": product_id}), 404
    time.sleep(random.uniform(0.01, 0.05))
    return jsonify({"productId": product_id, **product})


@app.get("/api/products/<product_id>/rating")
def get_rating(product_id):
    if product_id not in PRODUCTS:
        return jsonify({"error": "product not found", "productId": product_id}), 404
    reviews = REVIEWS.get(product_id, [])
    log.info("computing rating summary for %s (%d reviews)", product_id, len(reviews))
    return jsonify({
        "productId": product_id,
        "reviewCount": len(reviews),
        "averageRating": average_rating(reviews),
    })


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
