import os.path
import csv
import json
from flask import Flask
from flask import request

dataPath = "/data/customers.json"
class Customer:
  def __init__(self, c="",f="",l=""):
    self.customerID = c
    self.firstName  = f
    self.lastName   = l
  def fullName(self):
    return self.firstName + " " + self.lastName

def getCustomers():
  if os.path.isfile(dataPath):
    with open(dataPath, newline='') as customerFile:
      data = customerFile.read()
      customers = json.loads(data)
      return customers
  else: 
    return {}

def getCustomer(customerID):
  customers = getCustomers()

  if customerID in customers:
    return customers[customerID]
  else:
    return {}

def updateCustomers(customers):
  with open(dataPath, 'w', newline='') as customerFile:
    customerJSON = json.dumps(customers)
    customerFile.write(customerJSON)

app = Flask(__name__)

@app.route("/", methods=['GET'])
def get_customers():
    customers = getCustomers()
    return json.dumps(customers)

@app.route("/get/<string:customerID>", methods=['GET'])
def get_customer(customerID):
    customer = getCustomer(customerID)

    if customer == {}:
      return {}, 404
    else:
      return customer

@app.route("/add", methods=['POST'])
def add_customer():
    jsonData = request.json

    if "customerID" not in jsonData:
      return "customerID required", 400
    if "firstName" not in jsonData:
      return "firstName required", 400
    if "lastName" not in jsonData:
      return "lastName required", 400
    
    customers = getCustomers()
    customers[jsonData["customerID"]] = Customer( jsonData["customerID"], jsonData["firstName"], jsonData["lastName"]).__dict__
    updateCustomers(customers)
    return "success", 200