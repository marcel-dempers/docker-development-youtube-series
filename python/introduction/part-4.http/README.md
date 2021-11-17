# Introduction to Python: HTTP (Flask)

So far, we've learnt Python fundamentals, worked with data, files and 
more importantly, basic data structures like `csv` and `json`.

JSON is a very popular format for now only storing, but transporting data in web HTTP applications. Files are important because we may 
need to read configuration files for our application. </br>

One application could make a web request to ask for customer data from our application. By sending back JSON, the receiving application can easily process the data it gets.

Be sure to checkout [Part 2: Files](../part-2.files) </br>
As well as [Part 3: JSON](../part-3.json) </br>

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

We're going to use what we've learnt in part 1,2 & 3 and create
our customer app that stores customer data </br>
Firstly we have to import our dependencies:

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

In the previous episode, we've created a `json` file to hold all our customers. </br>
We've learnt how to read and write to file and temporarily use the file for our storage. </br>

```
{
  "a": {
    "customerID": "a",
    "firstName": "James",
    "lastName": "Baker"
  },
  "b": {
    "customerID": "b",
    "firstName": "Jonathan",
    "lastName": "D"
  },
  "c": {
    "customerID": "c",
    "firstName": "Aleem",
    "lastName": "Janmohamed"
  },
  "d": {
    "customerID": "d",
    "firstName": "Ivo",
    "lastName": "Galic"
  },
  "e": {
    "customerID": "e",
    "firstName": "Joel",
    "lastName": "Griffiths"
  },
  "f": {
    "customerID": "f",
    "firstName": "Michael",
    "lastName": "Spinks"
  },
  "g": {
    "customerID": "g",
    "firstName": "Victor",
    "lastName": "Savkov"
  }
}
```

Now that we have our customer data, we can read and update the data with our two simple functions: 


```
customers = getCustomers()
print(customers)
```

Let's add a customer and write it back to file:

```
customers["h"] = Customer("h", "Marcel", "Dempers").__dict__
updateCustomers(customers)
```

Now `customers.json` has the new entry if we rerun our code:

```
python src/app.py
```

## Flask

Checkout official Flask documentation [here](https://flask.palletsprojects.com/en/2.0.x/)

A minimal flask application looks like this: 

```
from flask import Flask

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"
```

To get this to work, we need to look at Pythons Library system, or `pip`

## External Packages

If we look at the `from flask import Flask` statement, it assumes we have a `flask` dependency. Our previous import statements worked, because they are built-in libraries in python. </br>
Flask is an external package. </br>
External packages are managed by [pip](https://pypi.org/project/pip/)

We may use `pip` to install `Flask`. If we need to connect to databases, we may install packages to do so. Like a `mysql` or a `redis` package. </br>

With `pip`, we define our dependencies in a `requirements.txt` file.
We can check for the latest version of `Flask` [here](https://pypi.org/project/Flask/)

We can also use `pip install` commands to install dependencies. </br>
However, this means we need a new `pip install` command for every dependency which will grow as 
our application and needs grow. Best is to use a `requirements.txt` file.

Let's create a `requirements.txt` file:

```
Flask == 2.0.2
```

We can install our dependencies using:

```
pip install -r src/requirements.txt
```

Let's implement the minimal example we posted above. </br>
To run our application, note that the `app.py` runs as per usual. <br/>
But to start `Flask`, we can follow the instructions as per the official document:

```
export FLASK_APP=src/app
flask run -h 0.0.0 -p 5000
```

Note that we cannot access `http://localhost:5000` because our app is running in docker. </br>
To access it, we need to restart our container and this time expose our port.

```
exit
docker run -it -p 5000:5000 -v ${PWD}:/work python sh

#get our dependencies and start our application
pip install -r requirements.txt
export FLASK_APP=src/app
flask run -h 0.0.0 -p 5000
```
Now we can access our app on `http://localhost:5000/`

## Routes

Web servers allow us to define various URLs or routes. </br>
Currently we can see logs showing `GET` requests for route `/` </br>
Notice the HTTP status code `200` indicating success.

```
 * Running on http://0.0.0:5000/ (Press CTRL+C to quit)
172.17.0.1 - - [22/Sep/2021 23:10:14] "GET / HTTP/1.1" 200 -
172.17.0.1 - - [22/Sep/2021 23:10:14] "GET /favicon.ico HTTP/1.1" 404 -
172.17.0.1 - - [22/Sep/2021 23:12:20] "GET / HTTP/1.1" 200 -
172.17.0.1 - - [22/Sep/2021 23:12:21] "GET / HTTP/1.1" 200 -
172.17.0.1 - - [22/Sep/2021 23:12:21] "GET / HTTP/1.1" 200 -
```

Routes allow us to have many endpoints, so we could have endpoints such as:

```
/all            --> which returns all customers
/get/<customerID>    --> which returns one customer by CustomerID
/add         --> which adds or updates a customer
```

Defining these routes are pretty straight forward, see [docs](https://flask.palletsprojects.com/en/2.0.x/quickstart/#variable-rules)

Let's define our routes:

```

@app.route("/")
def get_customers():
    return "<p>Hello, get_customers!</p>"

@app.route("/get/<string:customerID>", methods=['GET'])
def get_customer(customerID):
    return "<p>Hello, get!</p>"

@app.route("/add", methods=['POST'])
def add_customer(customer):
    return "<p>Hello, add!</p>"

```

Now that we've built our URLs using routes, next we need to understand HTTP methods

## HTTP Methods

There are a number of HTTP methods for web services, popular ones being `POST` and `GET`

`GET` method is for general retrieval of data </br>
`POST` method is for passing data to our service </br>

Let's setup each of our routes with dedicated HTTP Methods

```
@app.route("/", methods=['GET'])
def get_customers():
    return "<p>Hello, get_customers!</p>"

@app.route("/get/<string:customerID>", methods=['GET'])
def get_customer(customerID):
    return "<p>Hello, get!</p>"

@app.route("/add", methods=['POST'])
def add_customer():
    return "<p>Hello, add!</p>"

```

Now we can see if we access `http://localhost:5000/add`
We get: `Method Not Allowed` , because it only accepts `POST` and browser by default, does `GET`

Let's fill out our routes

### Get Customers

```
@app.route("/", methods=['GET'])
def get_customers():
    customers = getCustomers()
    return json.dumps(customers)
```

### Get Customer

```
@app.route("/get/<string:customerID>", methods=['GET'])
def get_customer(customerID):
     customer = getCustomer(customerID)
     return customer
```

## HTTP Status Codes 

You can access your browser's developer tools and see the network tab </br>
Status codes are important for troubleshooting production traffic.

`200` = Success. </br>

We can then access a customer via the `/get/<customerID>` route.
Our Customer ID's are `a` to `h` and notice we get an `Internal Server Error` if
we pass an incorrect ID. </br> 

`500` = Internal Server Error </br>
This means something went wrong in our code. See the logs indicating a `KeyError`, because we're trying to access a key in our customer dictionary that does not exist. </br>

Let's handle that error

```
#update our getCustomer function to check if the key exists

def getCustomer(customerID):
  customers = getCustomers()
  if customerID in customers:
    return customers[customerID]
  else: 
    return {} 
```

Note when we rerun this, we get a blank customer back. </br>
Also note we're getting a `200` status code. </br>
This may not be super appropriate, as our customer we requested is technically not found. There is a status code for that </br>

`404` = Not Found </br>

Let's handle the status code.

```
@app.route("/get/<string:customerID>", methods=['GET'])
def get_customer(customerID):
    customer = getCustomer(customerID)
    if customer == {}:
      return {}, 404
    else:
      return customer
```

## Add Customer 

In order to add a customer, we need to be able to read request data. </br>
More about `Requests` [here](https://flask.palletsprojects.com/en/2.0.x/api/#flask.Request)

Import requests library
```
from flask import request
```

Read the request body as JSON:

```
@app.route("/add", methods=['POST'])
def add_customer():
    print(request.json)
    return "success", 200
```
Let's test that!

```
curl -X POST http://localhost:5000/add -d '{ "data" : "hello" }'
```
Notice our print message is `None`, this is because we need to specify content type in the request header. Let's send the content type as JSON.

```
curl -X POST \
-H 'Content-Type: application/json' \
http://localhost:5000/add -d '{ "data" : "hello" }'
```

Now we see that we read the data as `{'data': 'hello'}` </br>
This allows us to read customer data as JSON, validate it, and write it to storage. </br>

Let's `POST` a customer:

```
curl -X POST -v \
-H 'Content-Type: application/json' \
http://localhost:5000/add -d '
{
  "customerID": "i",
  "firstName": "Bob",
  "lastName": "Smith"
}'
```
Note if we break the `json` format, we get a `400` status code </br>

`400` = Bad Request </br>

400 indicates a bad request, letting the client know they have to send the right formatted data.

## Docker

Let's build our container image and run it while mounting our customer file

Our final `dockerfile`
```
FROM python:3.9.6-alpine3.13 as dev

WORKDIR /work

FROM dev as runtime
WORKDIR /app
COPY ./src/requirements.txt /app/
RUN pip install -r /app/requirements.txt

COPY ./src/app.py /app/app.py
ENV FLASK_APP=app.py

CMD flask run -h 0.0.0 -p 5000

```

Build our container.

```
cd python\introduction\part-4.http

docker build . -t customer-app

```

Notice that we need to mount our `customers.json` file into the container. </br>
But also notice the containers working directory is `/app` and in the `app.py`
we load `customers.json` from the working directory which is the same directory as the app. </br>
When mounting files to a container, docker effectively formats that path, so to mount our json file, we would lose our `app.py` because of the format. <br/>

Therefore it's always important to mount configs, secrets or data to seperate folders.

Let's change that to `/data`

```
#set a global variable 
dataPath = "/data/customers.json"
```
Now we can change all references to `customers.json` to our variable.
And then mount that location:

```
docker build . -t customer-app
docker run -it -p 5000:5000 -v ${PWD}:/data customer-app
```