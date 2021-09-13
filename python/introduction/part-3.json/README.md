# Introduction to Python: JSON

JSON is a very important format for storing data in Software engineering. </br>
JSON is popular for the following scenarios :

* Popular standard for data structures in general.
* Applications may talk to each other by passing JSON data back and forth (HTTP Web APIs).
* Applications often store configuration in JSON format as a configuration file.
* Data may be stored in databases or caches in JSON format.
* DevOps engineers may store infrastructure configuration in JSON format.

## Python Dev Environment

The same as Part 1, we start with a [dockerfile](./dockerfile) where we declare our version of `python`.

```
FROM python:3.9.6-alpine3.13 as dev

WORKDIR /work
```

Let's build and start our container: 

```
cd python\introduction\part-3.json

docker build --target dev . -t python
docker run -it -v ${PWD}:/work python sh

/work # python --version
Python 3.9.6

```

## Our application

Firstly we have a class to define what a customer looks like:
```
class Customer:
  def __init__(self, c="",f="",l=""):
    self.customerID = c
    self.firstName  = f
    self.lastName   = l
  def fullName(self):
    return self.firstName + " " + self.lastName
```

Then we need a function which returns our customers:
```
def getCustomers():
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
  return customers
```

Here is a function to return a specific customer:
```
def getCustomer(customerID):
  customer = getCustomers()
  return customer[customerID]
```

Test our functions:

```
customers = getCustomers()
customer = customers["h"]

print(customers)
print(customer.fullName)
```

## Files

In [part-2](../part-2.files/README.md) we learnt about reading and writing to files. </br>
We changed our functions to get customers from file and update customers by writing the data back to file. </br>

Let's do that quick:

## Get Customers

```
import os.path
import csv

def getCustomers():
  if os.path.isfile("customers.log"):
    with open('customers.log', newline='') as customerFile:
      reader = csv.DictReader(customerFile)
      l = list(reader)
      customers = {c["customerID"]: c for c in l}
      return customers
  else: 
    return {}
```

## Update Customers

Let's create a function to update our customers:

```
def updateCustomers(customers):
  fields = ['customerID', 'firstName', 'lastName']
  with open('customers.log', 'w', newline='') as customerFile:
    writer = csv.writer(customerFile)
    writer.writerow(fields)
    for customerID in customers:
      customer = customers[customerID]
      writer.writerow([customer.customerID, customer.firstName, customer.lastName])
```

## Test 

Let's test our two functions to ensure they work

```
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

#see the changes
customers = getCustomers()
for customer in customers:
  print(customers[customer])

```

## JSON

JSON is a very simple data structure for building objects.
you can define any type of object and data using JSON. </br>

In JSON, objects are defined by open and close brackets, like `{}` </br>
Arrays or lists are defined as square brackets, like `[]` </br>

We can therefore define a customer like :
`Notice keys and values have quotes, and separated by colon`

```
{
  "customerID" : "a",
  "firstName: "Bob",
  "lastName": "Smith"
}

```
We can define an array of customers in JSON by declaring our array with square brackets and having customer objects inside, separated by commas:

```
# Example Array:
[{customer-1},{customer-2},{customer-3},{customer-4} ]
```

## JSON Library

Python provides a library for dealing with JSON. </br>
Let's take our customer dictionary, and convert it to json structure and 
finally write it to file using our update function.

```
import json
```

Now first things first, we need to convert our customers dictionary to a JSON structured string so we can store it to file. </br>

The `json` library allows us to convert dictionaries to files using 
the `dumps` function:

```
jsonData = json.dumps(customers)
print(jsonData)
```

But if we run this we get an error:
`TypeError: Object of type Customer is not JSON serializable`

That is because the library can only deal with full dictionaries.
Taking a closer look at our dictionary, our customer is a class.

```
# our key is a string
# our value is a class
"h" : Customer("h", "Marcel", "Dempers")
```

`json.dumps()` only allows full dictionaries so we need our customer object to be in dictionary form like so:

```
"h": { "customerID": "h", "firstName" : "Marcel", "lastName" : "Dempers"}
```

To fix our dictionary, we can iterate it and build a new one

```
# start with an empty one
customerDict = {}

#populate our new dictionary
for id in customers:
  customerDict[id] = customers[id].__dict__

#print it
print(customerDict)
```

Now we can convert our customer dictionary to json

```
customerJSON = json.dumps(customerDict)
print(customerJSON)
```

Now we've converted our data into JSON, we can view this in a new tab window and change the language mode to JSON to visualise the data in our new format.

Now essentially our application will use dictionaries when working with the data
but store and transport it as JSON. </br>

Therefore our update function should take a dictionary and convert it to JSON
and store it.

```
def updateCustomers(customer):
  with open('customers.json', 'w', newline='') as customerFile:
    customerJSON = json.dumps(customer) 
    customerFile.write(customerJSON)
```

We also need to change our `getCustomers` function to read the JSON file
and convert that data correctly back to a dictionary for processing.

```
def getCustomers():
  if os.path.isfile("customers.json"):
    with open('customers.json', newline='') as customerFile:
      data = customerFile.read()
      customers = json.loads(data) 
      return customers
  else: 
    return {}
```

Finally let's test our functions :

```
customers = getCustomers()
print(customers)
```

It's also very easy to add a customer to our JSON data using our class and using pythons internal `__dict__` property to convert the object to a dictionary and add it to our customers dictionary

```
customers["i"] = Customer("i", "Bob", "Smith").__dict__
updateCustomers(customers)
```

## Docker

Let's build our container image and run it while mounting our customer file

Our final `dockerfile`
```
FROM python:3.9.6-alpine3.13 as dev

WORKDIR /work

FROM dev as runtime
COPY ./src/ /app 

ENTRYPOINT [ "python", "/app/app.py" ]

```

Build and run our container.
Notice the `customers.json` file gets created if it does not exist.

```
cd python\introduction\part-3.json

docker build . -t customer-app

docker run -v ${PWD}:/work -w /work customer-app

```