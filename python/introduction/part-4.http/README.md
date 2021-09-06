# Introduction to Python: HTTP (Flask)

So far, we've learnt Python fundamentals, worked with data, files and 
more importantly, basic data structures like `csv` and `json`.

JSON is a very popular format for now only storing, but transporting data in web HTTP applications. </br>

One application could make a web request to ask for customer data from our application. By sending back JSON, the receiving application can easily process the data it gets.

## Python Dev Environment

The same as Part 1, we start with a [dockerfile](./dockerfile) where we declare our version of `python`.

```
FROM python:3.9.6-alpine3.13 as dev

WORKDIR /work
```

Let's build and start our container: 

```
cd python\introduction\part-4.http

docker build --target dev . -t python
docker run -it -v ${PWD}:/work python sh

/work # python --version
Python 3.9.6

```

## Our application

Firstly we have to import our dependencies 

```
import os.path
import csv
import json

```

Then we have a class to define what a customer looks like:
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
  if os.path.isfile("customers.json"):
    with open('customers.json', newline='') as customerFile:
      data = customerFile.read()
      customers = json.loads(data)
      return customers
  else: 
    return {}
```

Here is a function to return a specific customer:
```
def getCustomer(customerID):
  customer = getCustomers()
  return customer[customerID]
```
And finally a function for updating our customers:

```
def updateCustomers(customers):
  with open('customers.json', 'w', newline='') as customerFile:
    customerJSON = json.dumps(customers)
    customerFile.write(customerJSON)
```

Now to recap and test our functions we need to define a list of customers
to use:

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
```

Now to convert our customer dictionary to JSON, we need to convert our dictionary of customer objects to a dictionary of customers where each customer is a dictionary instead of a class. </br>
To do this we generate a new dictionary and populate it by converting each customer object to a dictionary :

```
customerDict = {}

for id in customers:
  customerDict[id] = customers[id].__dict__

```
Then we can finally test our update and get functionality:

```
updateCustomers(customerDict)

customers = getCustomers()
print(customers)
```

Let's test it and see the `customers.json` file stored to our server and 
retrieved.

```
python src/app.py
```

## External Libraries




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