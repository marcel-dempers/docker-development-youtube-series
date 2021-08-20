# Introduction to Learning Python

Best place to start is the documentation [www.python.org](https:#www.python.org/) <br/>

Download [Python 3](https:#www.python.org/downloads/) <br/>

Docker images for Python on [Docker Hub](https:#hub.docker.com/_/python) <br/>


# Run Python in Docker

We can also run python in a docker container: <br/>

```
cd python\introduction

docker build --target dev . -t python
docker run -it -v ${PWD}:/work python sh

/work # python --version
Python 3.9.

```

# Our first Program

* Create a folder containing our application source code

```
mkdir src

```

* Create basic source code

In the `src` folder, create a program called `app.py`
Paste the following content into it.

```

def main():
	print("Hello, world.")

main()
```

* Run your application code

You can run your application 

```
/work # python src/app.py
Hello, world.
```

# The Code

In the video we cover writing functions. </br>
It allows us to execute a block of code <br/>
You want to give your function a single purpose <br/>
Functions can have an input and return an output <br/>
Well thought out functions makes it easier to write tests <br/>

Instead of doing a boring `x + y` function that adds two numbers, let's do something a little
more realistic but still basic: 

```
# This function returns some data 
# The data could be coming from a database, or a file.
# The person calling the function should not care
# how the data is retrieved.
# Since the function does not leak its Data provider

def getCustomer():
  return "Marcel Dempers"

print(getCustomer())

```

# Function inputs

As we saw, functions can return outputs, and also take inputs.
Let's accept a customer ID and return a customer record with that ID:

```
def getCustomer(customerID):
  return "CustomerID: " + customerID + ", Name: Marcel Dempers"

print(getCustomer("abc"))
```


## Variables

To hold data in programming languages, we use variables. <br/>
Variables take up space in memory, so we want to keep it minimal. <br/>
Let's declare variables in our function:

```
def getCustomer(customerID):
  firstName = "Marcel"
  lastName = "Dempers"
  fullName = firstName + " " + lastName

  return fullName

  #or we can return the computation instead of adding another variable!
  return firstName + " " + lastName

  #or we don't even need to declare variables :)
  return "Marcel Dempers"

customer = getCustomer("abc")
print(customer)

```

## Control Flows (if\else)

You can see we're not using the `customerId` input in our function. <br/>
Let's use it and showcase control flows to create logic! <br/>

Control flows allow us to add "rules" to our code. </br>
"If this is the case, then do that, else do something else".

So let's say we have a customer ID 1 coming in, we may only want to
return our customer if it matches the `customerID`

```

def getCustomer(customerID):
  if customerID == "abc":
    return "Marcel Dempers"
  elif customerID == "def":
    return "Bob Smith"
  else:
    return ""

customer = getCustomer("abc")
print(customer)

```

Let's invoke our function :

```
marcel = getCustomer("abc")
print(marcel)

bob = getCustomer("def")
print(bob)

```

## Arrays

Our function can only return one customer at a time. <br/>
What if we wanted to return more than one customer. <br/>
In real world systems, we may return database records. <br/>
This is where we start looking at arrays, lists, dictionaries etc. <br/>

Let's add another function to get an array of customers!

```
def getCustomers():
  customers = [ "Marcel Dempers", "Bob Smith" ]
  return customers

customers = getCustomers()
print(customers)
```

You can get array values by index starting at `0`:

```
customers = getCustomers()

marcel = customers[0]
bob = customers[1]

print(marcel + "\n" + bob)

```

We can set a value in our Array. <br/>
Arrays are fixed in size, notice we cannot just access a 
record : <br/>

```
customers[2] = "John Smith"
IndexError: list assignment index out of range

```

Instead we can use the `append()` function of the array to add items.

```
def getCustomers():
  customers = [ "Marcel Dempers", "Bob Smith" ]
  customers.append("James Baker")
  customers.append("Jonathan D")
  customers.append("Aleem Janmohamed")
  customers.append("Ivo Galic")
  customers.append("Joel Griffiths")
  customers.append("Michael Spinks")
  customers.append("Victor Savkov")
  return customers

```

We can remove items with the `remove()` function.

```
customers.remove("Bob Smith")
```

# Dictionaries

Dictionaries have an advantage whereby items can be found using a key. <br/>
With arrays, we have to know the index number where the item is located or else we cannot find items unless we write a loop. <br/>
With Dictionaries, we store key value pairs. So the key, may be our customer ID.<br/>

Let's replace our customer array with a dictionary
```
  customers = {
    "a": "James Baker",
    "b": "Jonathan D",
    "c": "Aleem Janmohamed",
    "d": "Ivo Galic",
    "e": "Joel Griffiths",
    "f": "Michael Spinks",
    "g": "Victor Savkov"
  }

```

We can also set and access an item by key

```
#update or create item
customers["h"] = "Marcel Dempers"

#get an item
marcel = customers["h"]
print(marcel)

```

Let's update our `getCustomer()` function to get customer from our dictionary

```
def getCustomer(customerID):
  customer = getCustomers()
  return customer[customerID]

def getCustomers():
  customers = {
    "a": "James Baker",
    "b": "Jonathan D",
    "c": "Aleem Janmohamed",
    "d": "Ivo Galic",
    "e": "Joel Griffiths",
    "f": "Michael Spinks",
    "g": "Victor Savkov",
    "h" : "Marcel Dempers"
  }
  
  return customers

customers = getCustomers()
customer = getCustomer("h")

print(customer)
print(customers)
```

## Loops 

Loops are used to iterate over collections, lists, arrays etc. <br/>
There are two primitive loops in Python </br>
Let's say we need to loop through our customers

### While loop

While loops are simple to write by condition:

```
i = 0
while i < 7:
  print(customers[i])
  i += 1
```

We can also use this condition to loop arrays:

```
customers = [ "Marcel Dempers", "Bob Smith" ]
customers.append("James Baker")
customers.append("Jonathan D")
customers.append("Aleem Janmohamed")
customers.append("Ivo Galic")
customers.append("Joel Griffiths")
customers.append("Michael Spinks")
customers.append("Victor Savkov")
i = 0
while i <= 8:
  print(customers[i])
  i += 1
```

### For loops

For looping a dictionary of list, `for` loops are a little easier:

```
customers = getCustomers()

for customerID in customers:
  print(customers[customerID])
```

## Classes and Objects

So far so good, however, customer data is not useful as strings. <br/>
Customers can have a firstname, lastname, and more properties. <br/>

For this purpose we'd like to group some variables into a single variable. <br/>
This is what `classes` allows us to do. <br/>
Let's create a `class` for our customer

```
class Customer:
  customerID = ""
  firstName  = ""
  lastName   = ""
  def fullName(self):
    return self.firstName + " " + self.lastName

marcel = Customer()
marcel.firstName = "Marcel"
marcel.lastName = "Dempers"
marcel.customerID = "h"

print(marcel.fullName())

```

Now how can we create an object with our values without having to set every
value individually ? 
Constructors allow us to initialise our object with some default values.

```
class Customer:
  def __init__(self, c="", f="", l=""):
    self.customerID = c
    self.firstName  = f
    self.lastName   = l
  def fullName(self):
    return self.firstName + " " + self.lastName

marcel = Customer("h", "Marcel", "Dempers")

print(marcel.fullName())

```

Let's plug our class into our current application and put it all together:
Firstly, we convert all our customers from strings to objects:

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

Then we can proceed to get our customer and access their fields:

```
customers = getCustomers()
for customerID in customers:
  print(customers[customerID].fullName())
```

# Docker

So far we have used our container as a development environment. <br/>
Docker allows us to run python in an isolated environment. <br/>
Docker also has a concept of stages, I.E Docker Multistage. <br/>

Therefore we may have a stage for development where we simply mount source code to access a python environment. In addition to this, we may add another layer for debugging where we can install debugger and extra development tools which we 
do not want in production. <br/>
Finally we can create a smaller stage which contains only python + our source which we can run as a runtime image in production.

Example: 

```
FROM python:3.9.6-alpine3.13 as dev

WORKDIR /work

FROM python:3.9.6-alpine3.13 as debugging

# add a debugger

FROM dev as runtime
COPY ./src/ /app 

ENTRYPOINT [ "python", "/app/app.py" ]

```

## Building the Container

```
docker build . -t customer-app

```

## Running the Container

```
docker run customer-app
```