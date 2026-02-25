from flask import Flask, request, jsonify

app = Flask(__name__)

payments = []
id_counter = 1

# Get all payments 

@app.get("/payments")
def get_payments():
    return jsonify(payments)

# Add a payment

@app.post("/payments/process")
def process_payment():
    global id_counter

    payment = request.get_json() or {}
    payment["id"] = id_counter
    payment["status"] = "SUCCESS"
    id_counter += 1

    payments.append(payment)
    return jsonify(payment), 201

# Get payment by ID

@app.get("/payments/<int:payment_id>")
def get_payment(payment_id):
    for payment in payments:
        if payment.get("id") == payment_id:
            return jsonify(payment)
    return ("", 404)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8083)