from flask import Flask, jsonify
from kafka import KafkaConsumer, KafkaProducer
import os

app = Flask(__name__)

@app.route('/')
def index():
    return jsonify({
        "Hello": "World"
    })


if __name__ == "__main__":
    os.environ['FLASK_ENV'] = 'development'
    app.run(debug=True)
    
