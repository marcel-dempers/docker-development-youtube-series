# Introduction to Python: Storing data in Redis Database

So far, we've learnt Python fundamentals, worked with data, files , HTTP and more importantly, basic data structures like `csv` and `json`.

Be sure to checkout: </br>
[Part 1: Intro to Python](../README.md) </br>
[Part 2: Files](../part-2.files/README.md) </br>
[Part 3: JSON](../part-3.json/README.md) </br>

## Start up a Redis Cluster

Follow my Redis clustering Tutorial </br>

<a href="https://youtube.com/playlist?list=PLHq1uqvAteVtlgFkmOlIqWro3XP26y_oW" title="Redis"><img src="https://i.ytimg.com/vi/L3zp347cWNw/hqdefault.jpg" width="30%" alt="Redis Guide" /></a>

Code is over [here](../../../storage/redis/clustering/readme.md)

## Python Dev Environment

The same as Part 1+2+3+4, we start with a [dockerfile](./dockerfile) where we declare our version of `python`.

```
FROM python:3.9.6-alpine3.13 as dev

WORKDIR /work
```

Let's build and start our container: 

```
cd python\introduction\part-5.database.redis

docker build --target dev . -t python
docker run -it -v ${PWD}:/work -p 5000:5000 --net redis python sh

/work # python --version
Python 3.9.6

```

## Our application

We're going to use what we've learnt in part 1,2 & 3 and create
our customer app that handles customer data </br>
Firstly we have to import our dependencies:

```
import os.path
import csv
import json
from flask import Flask
from flask import request

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

And also set a global variable for the location of our videos `json` file:

```
dataPath = "./customers.json"
```

Then we need a function which returns our customers:
```
def getCustomers():
  if os.path.isfile(dataPath):
    with open(dataPath, newline='') as customerFile:
      data = customerFile.read()
      customers = json.loads(data)
      return customers
  else: 
    return {}
```

Here is a function to return a specific customer:
```
def getCustomer(customerID):
  customers = getCustomers()

  if customerID in customers:
    return customers[customerID]
  else:
    return {}
```
And finally a function for updating our customers:

```
def updateCustomers(customers):
  with open(dataPath, 'w', newline='') as customerFile:
    customerJSON = json.dumps(customers)
    customerFile.write(customerJSON)
```

In the previous episode, we've created a `json` file to hold all our customers. </br>
We've learnt how to read and write to file and temporarily use the file for our storage. </br>

Let's create a file called `customers.json` : 
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

Now that we have our customer data and functions to read and update our customer data, let's define our `Flask` application:


```
app = Flask(__name__)
```

We create our route to get all customers:

```
@app.route("/", methods=['GET'])
def get_customers():
    customers = getCustomers()
    return json.dumps(customers)
```

A route to get one customer by ID:

```
@app.route("/get/<string:customerID>", methods=['GET'])
def get_customer(customerID):
    customer = getCustomer(customerID)

    if customer == {}:
      return {}, 404
    else:
      return customer
```

And finally a route to update or add customers called `/set` :

```
@app.route("/set", methods=['POST'])
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
```

Before we can be done, we need to import our `Flask` dependency we covered in our Python HTTP video. </br>
Let's create a `requirements.txt` file:

```
Flask == 2.0.2
```

We can install our dependencies using:

```
pip install -r requirements.txt
```

This gives us a web application that handles customer data and using a file as it's storage </br>
To test it, we can start up Flask:

```
export FLASK_APP=src/app
flask run -h 0.0.0 -p 5000
```

Now we can confirm it's working by accessing our application in the browser on `http://localhost:5000`

## Redis

To connect to Redis, we'll use a popular library called `redis-py` which we can grab from [here](https://github.com/redis/redis-py) </br>
The pip install is over [here](https://pypi.org/project/redis/3.5.3/) </br>

Let's add that to our `requirements.txt` dependency files.

```
redis == 3.5.3
```

We can proceed to install it using `pip install`

```
pip install -r requirements.txt
```

Now to connect to Redis in a highly available manner, we need to take a look at the 
`Sentinel support` section of the guide </br>

Let's test the library. The beauty of Python is that it's a scripting language, so we don't have to compile and keep restarting our application, we can test each line of code. </br>

```
python
from redis.sentinel import Sentinel

sentinel = Sentinel([('sentinel-0', 5000),('sentinel-1', 5000),('sentinel-2', 5000)], socket_timeout=0.1)

sentinel.discover_master('mymaster')

sentinel.discover_slaves('mymaster')

master = sentinel.master_for('mymaster',password = "a-very-complex-password-here", socket_timeout=0.1)

slave = sentinel.slave_for('mymaster',password = "a-very-complex-password-here", socket_timeout=0.1)

master.set('foo', 'bar')
slave.get('foo')
```

We can demonstrate reading and writing a key value pair. </br>
We can also demonstrate failure, when we stop the current master, we'll get a connection error. It's important to implement retry logic. </br>
If we wait a moment and execute commands again, we will see that it starts to work.


```
# stop current master
docker rm -f redis-0

master.set('foo', 'bar2')

redis.exceptions.ConnectionError: Connection closed by server.

# retry moments later...

master.set('foo', 'bar2')
slave.get('foo')

sentinel.discover_master('mymaster')
sentinel.discover_slaves('mymaster')
```

We can find the current master by running `docker inspect` to see who owns that IP address.

Start up `redis-0` again, to simulate a recovery from failure.

## Connecting our App to Redis

To connect to redis, we'll want to read the connection info from environment variables. Let's set some global variables.

```
import os

redis_sentinels = os.environ.get('REDIS_SENTINELS')
redis_master_name = os.environ.get('REDIS_MASTER_NAME')
redis_password = os.environ.get('REDIS_PASSWORD')

```

We will need to restart our container so we can inject these environment variables. Let's go ahead and do that:

```
docker run -it -p 5000:5000 `
  --net redis `
  -v ${PWD}:/work `
  -e REDIS_SENTINELS="sentinel-0:5000,sentinel-1:5000,sentinel-2:5000" `
  -e REDIS_MASTER_NAME="mymaster" `
  -e REDIS_PASSWORD="a-very-complex-password-here" `
  python sh

# re-install our dependencies
pip install -r requirements.txt
```

Now we can setup a client:

```
from redis.sentinel import Sentinel

sentinels = []

for s in redis_sentinels.split(","):
  sentinels.append((s.split(":")[0], s.split(":")[1]))

redis_sentinel = Sentinel(sentinels, socket_timeout=5)
redis_master = redis_sentinel.master_for(redis_master_name,password = redis_password, socket_timeout=5)
```

## Retry logic

Now we noticed that if we have a master that fails, the sentinels will choose and assign a new master. We can see this by simply retrying our redis command. </br>

When talking to redis we need to have some retry capability to be able to recover from this scenario. </br>

Let's build a retry function at the top of our application, that runs a redis command:

```
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
```

We can test out our `redis_command` by calling it and printing the result 
to the screen

```
print(redis_command(redis_master.set, 'foo', 'bar'))
print(redis_command(redis_master.get, 'foo'))
```

We can simulate failure again, by finding and stopping the current master.

Once we're done with our tests, we can `exec` into the current master and run `FLUSHALL` to remove our test records from redis.

## Saving our data to Redis

Now let's change our customer functions to point to Redis instead of file </br>

Starting with `getCustomer` to retrieve a single customer
```
def getCustomer(customerID):
  customer = redis_command(redis_master.get, customerID)
  
  if customer == None:
    return {}
  else:
    c = str(customer, 'utf-8')
    return json.loads(c)
```
Now we can use that to return all our customers by updating the `getCustomers` function:

```
def getCustomers():
  customers = {}
  customerIDs = redis_command(redis_master.scan_iter, "*")
  for customerID in customerIDs:
    customer = getCustomer(customerID)
    customers[customer["customerID"]] = customer
  
  return customers
```

Let's improve our functions by adding a new function to update a single customer:

```
def updateCustomer(customer):
  redis_command(redis_master.set, customer.customerID, json.dumps(customer.__dict__))

```

And finally we can use that function to update all customers by tweaking our `updateCustomers` function:

```
def updateCustomers(customers):
  for customer in customers:
    updateCustomer(customer)
```

Now our simple functions are done, let's hook them up to our endpoints

```
# firstly delete these test lines
print(redis_command(redis_master.set, 'foo', 'bar'))
print(redis_command(redis_master.get, 'foo'))
```

Our simple Get all 

```
@app.route("/", methods=['GET'])
def get_customers():
  customers = getCustomers()
  return json.dumps(customers)
```

Our Get by ID

```
@app.route("/get/<string:customerID>", methods=['GET'])
def get_customer(customerID):
    customer = getCustomer(customerID)

    if customer == {}:
      return {}, 404
    else:
      return customer
```

And our update endpoint to update a customer

```
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
```
## Docker

Let's build our container image and run it while mounting our customer file

Our final `dockerfile`
```
FROM python:3.9.6-alpine3.13 as dev

WORKDIR /work

FROM dev as runtime
WORKDIR /app

COPY ./requirements.txt /app/
RUN pip install -r /app/requirements.txt

COPY ./src/app.py /app/app.py
ENV FLASK_APP=app.py

CMD flask run -h 0.0.0 -p 5000

```

Build our container.

```
cd python\introduction\part-5.database.redis

docker build . -t customer-app

```

Now we can run our production container:

```
docker build . -t customer-app

docker run -it -p 5000:5000 `
  --net redis `
  -e REDIS_SENTINELS="sentinel-0:5000,sentinel-1:5000,sentinel-2:5000" `
  -e REDIS_MASTER_NAME="mymaster" `
  -e REDIS_PASSWORD="a-very-complex-password-here" `
  customer-app
```