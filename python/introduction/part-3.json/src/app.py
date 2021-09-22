import os.path
import csv
import json

class Customer:
  def __init__(self, c="",f="",l=""):
    self.customerID = c
    self.firstName  = f
    self.lastName   = l
  def fullName(self):
    return self.firstName + " " + self.lastName

def getCustomers():
  if os.path.isfile("customers.json"):
    with open('customers.json', newline='') as customerFile:
      data = customerFile.read()
      customers = json.loads(data)
      return customers
  else: 
    return {}

def getCustomer(customerID):
  customer = getCustomers()
  return customer[customerID]

def updateCustomers(customers):
  with open('customers.json', 'w', newline='') as customerFile:
    customerJSON = json.dumps(customers)
    customerFile.write(customerJSON)
    
customers = {
    "a": Customer("a","James", "Baker"),
    "b": Customer("b", "Jonathan", "D"),
    "c": Customer("c", "Aleem", "Janmohamed"),
    "d": Customer("d", "Ivo", "Galic"),
    "e": Customer("e", "Joel", "Griffiths"),
    "f": Customer("f", "Michael", "Spinks"),
    "g": Customer("g", "Victor", "Savkov"),
    "h" : Customer("h", "Marcel", "Dempers")
}

customerDict = {}

for id in customers:
  customerDict[id] = customers[id].__dict__


updateCustomers(customerDict)

customers = getCustomers()
customers["i"] = Customer("i", "Bob", "Smith").__dict__

updateCustomers(customers)
print(customers)