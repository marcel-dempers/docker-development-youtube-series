# Introduction to Python: FILES

In Python, dealing with files is very common and is a very important part of 
programming for a number of reasons: </br>

* Applications may need to read configuration files
* In Data science, data is often sourced from files (`CSV`, `XML`, `JSON`, etc)
* Data is often analysed in Python when its written in different stages of analysis
* DevOps engineers often stores the state of infrastructure or data as files for automation purposes.

Files are not the endgame for storage. </br>
Remember there are things like Caches and Databases. </br>
But before learning those things, file handling is the best place to start. </br>

## Python Dev Environment

The same as Part 1, we start with a [dockerfile](./dockerfile) where we declare our version of `python`.

```
cd python\introduction\part-2.files

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

## Opening Files

Python provides an `open` function to open files. </br>
`open()` takes a file path\name and access mode

```
"r" - Read - Default value. Opens a file for reading, error if the file does not exist
"a" - Append - Opens a file for appending, creates the file if it does not exist
"w" - Write - Opens a file for writing, creates the file if it does not exist
"x" - Create - Creates the specified file, returns an error if the file exists
```

Try open a file that holds our customer data:

```
open("customers.log")
```

We can see the file does not exist:

```
/work # python src/app.py
Traceback (most recent call last):
  File "/work/src/app.py", line 26, in <module>
    open("customers.log")
FileNotFoundError: [Errno 2] No such file or directory: 'customers.log'
```

Let's use what we learned (`if` statements), to check if the file exists!
We'll need a built in library for handling files

```
import os.path
```

Then we can use the `os.path.isfile("customers.log")` command to check if the file exists

```
os.path.isfile("customers.log")
```

Using `if` logic we can check if the file is there:

```
if os.path.isfile("customers.log"):
  print("file exists")
else:
  print("file does not exists")
```

Now we know the file does not exist, but if it did, we can now read it with `open`


```
f = open("customers.log")
```

Let's also loop each customer in the file and print it

```
for customer in f:
  print(customer)
f.close()
```

Now we know the file does not exist, let's create it!

```
customers = getCustomers()
for customerID in customers:
  c = customers[customerID]
  f.write(c.customerID + "," + c.firstName + "," + c.lastName)
```

Now if we run our code the first time, it will create and populate the file as it does not exist,
and will read the file and display the content on the second run. </br>

Instead of looping each line in the file, we can read the entire file with the file's `read()` function:

```
print(f.read())
```

## Comma-Separated Values : CSV

As we can see, our `customers.log` file is in CSV format with every field separated by commas. </br>

So far, we've demonstrated using primitives to read and write to files to store our data.
When looping data structures like dictionaries and writing each line one by one to a file
will use a lot of CPU if the data is large. </br>

### CSV: Reading our file

To work with CSV's, we need to import a library
We also need to add headers to our file so it makes setting fields easier:

```
customerID, firstName, lastName
```

```
import csv
with open('customers.log', newline='') as customerFile:
  reader = csv.DictReader(customerFile)
  for row in reader:
    #print(row)
    print("customer id:" + row['customerID'] + " fullName : " + row['firstName'] + " " + row['lastName'])

```
### CSV: Writing our file

Create an array with our field headers

```
fields = ['customerID', 'firstName', 'lastName']
with open('customers.log', 'w', newline='') as customerFile:
  writer = csv.writer(customerFile)
  writer.writerow(fields)
  customers = getCustomers()
  for customerID in customers:
    customer = customers[customerID]
    writer.writerow([customer.customerID, customer.firstName, customer.lastName])
```

## Putting it all together

Now that we have code that reads and writes to a file, let's update our `getCustomers` function to return 
customers from our file. </br>

We read the file if it exists, read it into a list and convert the list to a dictionary:

```
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

We can test our function to see it working:

```
customers = getCustomers()
for customerID in customers:
  print(customers[customerID])
```

Let's also create a function to update customers

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

Let's test our two functions by deleting our file and recreate it using our functions:

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

#add another test customer
test = Customer("t", "Test", "Customer")
customers["t"] = test

#save it
updateCustomers(customers)

#see the changes
customers = getCustomers()
for customer in customers:
  print(customers[customer])
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
Notice the `customers.log` file get created if it does not exists.

```
cd python\introduction\part-2.files

docker build . -t customer-app

docker run -v ${PWD}:/work -w /work customer-app

```