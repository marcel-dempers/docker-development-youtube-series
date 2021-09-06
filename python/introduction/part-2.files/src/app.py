import os.path
import csv

class Customer:
  def __init__(self, c="",f="",l=""):
    self.customerID = c
    self.firstName  = f
    self.lastName   = l
  def fullName(self):
    return self.firstName + " " + self.lastName

def getCustomers():
  if os.path.isfile("customers.log"):
    with open('customers.log', newline='') as customerFile:
      reader = csv.DictReader(customerFile)
      l = list(reader)
      customers = {c["customerID"]: c for c in l}
      return customers
  else: 
    return {}

def updateCustomers(customers):
  fields = ['customerID', 'firstName', 'lastName']
  with open('customers.log', 'w', newline='') as customerFile:
    writer = csv.writer(customerFile)
    writer.writerow(fields)
    for customerID in customers:
      customer = customers[customerID]
      writer.writerow([customer.customerID, customer.firstName, customer.lastName])

def getCustomer(customerID):
  customer = getCustomers()
  return customer[customerID]

# if os.path.isfile("customers.log"):
#   with open('customers.log', newline='') as customerFile:
#     reader = csv.DictReader(customerFile)
#     for row in reader:
#       print("customer id:" + row['customerID'] + " fullName : " + row['firstName'] + " " + row['lastName'])
# else:
#   fields = ['customerID', 'firstName', 'lastName']
#   with open('customers.log', 'w', newline='') as customerFile:
#     writer = csv.writer(customerFile)
#     writer.writerow(fields)
#     customers = getCustomers()
#     for customerID in customers:
#       customer = customers[customerID]
#       writer.writerow([customer.customerID, customer.firstName, customer.lastName])

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

#save it
updateCustomers(customers)

#add another test customer
test = Customer("t", "Test", "Customer")
customers["t"] = test

#save it
updateCustomers(customers)

#see the changes
customers = getCustomers()
for customer in customers:
  print(customers[customer])

