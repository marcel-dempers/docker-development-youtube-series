import os.path
import csv
import json
import time
from flask import Flask
from flask import request
import os
from redis.sentinel import Sentinel

dataPath = "./customers.json"

redis_sentinels = os.environ.get('REDIS_SENTINELS')
redis_master_name = os.environ.get('REDIS_MASTER_NAME')
redis_password = os.environ.get('REDIS_PASSWORD')

def redis_command(command, *args):
  max_retries = 3
  count = 0
  backoffSeconds = 5
  while True:
    try:
      return command(*args)
    except (redis.exceptions.ConnectionError, redis.exceptions.TimeoutError):
      count += 1
      if count > max_retries:
        raise
      print('Retrying in {} seconds'.format(backoffSeconds))
      time.sleep(backoffSeconds)

class Customer:
  def __init__(self, c="",f="",l=""):
    self.customerID = c
    self.firstName  = f
    self.lastName   = l
  def fullName(self):
    return self.firstName + " " + self.lastName

def getCustomers():
  customers = {}
  customerIDs = redis_command(redis_master.scan_iter, "*")
  for customerID in customerIDs:
    customer = getCustomer(customerID)
    customers[customer["customerID"]] = customer
  return customers

def getCustomer(customerID):
  customer = redis_command(redis_master.get, customerID)
  
  if customer == None:
    return {}
  else:
    c = str(customer, 'utf-8')
    return json.loads(c)

def updateCustomer(customer):
  redis_command(redis_master.set,customer.customerID, json.dumps(customer.__dict__) )

def updateCustomers(customers):
  for customer in customers:
    updateCustomer(customer)

app = Flask(__name__)

sentinels = []

for s in redis_sentinels.split(","):
  sentinels.append((s.split(":")[0], s.split(":")[1]))

redis_sentinel = Sentinel(sentinels, socket_timeout=5)
redis_master = redis_sentinel.master_for(redis_master_name,password = redis_password, socket_timeout=5)


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

@app.route("/set", methods=['POST'])
def add_customer():
  jsonData = request.json

  if "customerID" not in jsonData:
      return "customerID required", 400
  if "firstName" not in jsonData:
      return "firstName required", 400
  if "lastName" not in jsonData:
      return "lastName required", 400

  customer = Customer( jsonData["customerID"], jsonData["firstName"], jsonData["lastName"])
  updateCustomer(customer)
  return "success", 200

