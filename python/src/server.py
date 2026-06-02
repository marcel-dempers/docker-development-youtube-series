from flask import Flask, request, jsonify
app = Flask(__name__)

@app.route('/', defaults={'subpath': ''})
@app.route('/<path:subpath>')

def catch_all(subpath):
    # Reconstruct the full path received by the backend container
    received_path = f"/{subpath}"
    
    # You can change this status code dynamically if you want to test failure modes
    status_code = 200 
    
    response_data = {
        "message": "Hello from Python!",
        "received_path": received_path,
        "raw_path_info": request.path,
        "query_string": request.query_string.decode('utf-8'),
        "status_code": status_code
    }
    
    # Returns a clean JSON payload and the explicit HTTP status code
    return jsonify(response_data), status_code

if __name__ == "__main__":
    # Internal port 3000 matches your Kubernetes Deployment targetPort
    app.run(host="0.0.0.0", port=5000)