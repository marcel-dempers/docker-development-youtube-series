# Introduction to Learning Go

<a href="https://youtu.be/jpKysZwllVw" title="golang-part-1"><img src="https://i.ytimg.com/vi/jpKysZwllVw/hqdefault.jpg" width="20%" alt="introduction to Go part 1" /></a>

Go can be downloaded from [golang.org](https://golang.org/doc/install) <br/>

Test your `go` installation:

```
go version
```

# Run Go in Docker

We can also run go in a small docker container: <br/>

```
cd golang\introduction

docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version

```

# Code Structure

https://golang.org/doc/code.html

* Package: 
  - Source files in same directory that are compiled together
  - Have visibility on all source files in the same package

* Modules:
  - Collection of packages that are released together

Our repository can contain one or more go modules, but usually 1.
  - At the root of the repo

`go.mod` Declares module path + import path for packages. (Where to download them)
  - When we write our own program, we can define a module path
  - This allows us to publish our code (if we want), so others can download it
  - The module path could be something like `github.com/google/go-cmp`
  - Makes it easy for other programs to consume our module

# Our first Program

* Create a folder containing our application

```
mkdir app

```

* Define a module path (github.com/docker-development-youtube-series/golang/introdouction/app)

```
# change directory to your application source code

cd app

# create a go module file

go mod init github.com/docker-development-youtube-series/golang/introdouction/app

```

* Create basic source code

In the `app` folder, create a program called `app.go`
Paste the following content into it.

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```

* Run your application code

You can run your application 

```
go run app.go
```

# Building our Program

Build your application into a static binary: <br/>

```
go build 
```

This will produce a compiled program called `app`
You can run this program easily:

```
./app
```

# Install your application (optional)

"This command builds the app command, producing an executable binary. <br/> 
It then installs that binary as $HOME/go/bin/app (or, under Windows, %USERPROFILE%\go\bin\app.exe)"

```
go install github.com/docker-development-youtube-series/golang/introdouction/app
```

# The Code

## Functions 

In the video we cover writing functions. </br>
It allows us to execute a block of code <br/>
You want to give your function a single purpose <br/>
Functions can have an input and return an output <br/>
Well thought out functions makes it easier to write tests <br/>

Instead of doing a boring `x + y` function that adds two numbers, let's do something a little
more realistic but still basic: 

```
// This function returns some data 
// The data could be coming from a database, or a file.
// The person calling the function should not care
// Since the function does not leak its Data provider

func getData(inputs)(outputs){
}

// functions can take multiple inputs, and return multiple outputs

// lets say we have 1) customers and 2) the cities they are from
// we may want to 1) get a list of customers and 2) get a list of cities
// therefore we have 2 types of data, 1) customers 2) cities
// let's improve our function so its gets data based on the type

func getData(customerId int) (customer string) {
}
```

## Variables

To hold data in programming languages, we use variables. <br/>
Variables take up space in memory, so we want to keep it minimal. <br/>
Let's declare variables in our function

```
func getData(customerId int) (customer string) {
  var firstName = "Marcel"
  lastName := "Dempers"
  
  fullName := firstName + " " + lastName
  return fullName

  //or we can return the computation instead of adding another variable!
  return firstName + " " + lastName

  //or we dont even need to declare variables :)
  return "Marcel Dempers"

}

```

## Control Flows (if\else)

You can see we're not using the `customerId` input in our function. <br/>
Let's use it! <br/>

Control flows allow us to add "rules" to our code. </br>
"If this is the case, then do that, else do something else".

So let's say we have a customer ID 1 coming in, we may only want to
return our customer if it matches the `customerId`

```

func getData(customerId int) (customer string) {

  if customerId == 1 {
    return "Marcel Dempers"
  } else if customerId == 2 {
    return "Bob Smith"
  } else {
    return ""
  }
  
}


```

Let's invoke our function :

```
//in the main() function

//get our customer
customer := getData(1)
fmt.Println(customer)

//get the wrong customer
customer := getData(3)
fmt.Println(customer)


```
## Arrays

At the moment, we can only return 1 customer at a time on our function. <br/>
Realistically we need the ability to return more data, not just a single customer. <br/>

Arrays allow us to make a collection of variables of the same type. <br/>
We can now return a list of customers! <br/>

Let's change our function to get an array of customers!

```
func getData() (customers [2]string) {
  //create 1 record
  customer := "Marcel Dempers"

  //assign our customer to the array
  customers[0] = customer

  //OR we can assign it like this
  customers[1] = "Bob Smith"

  //send it back to the caller
  return customers
  
}

```

Now we also have to change our calling function to expect an array:

```
customers := getData()

fmt.Println(customers)

```

## Slices

Since arrays are fixed size, Slices are a dynamically-sized view into arrays.
Let's create a slice instead of array so we can add customers dynamically!

```
func getData() (customers []string) {
  
  //initialise our slice of type string
  customers = []string{ "Marcel Dempers", "Bob Smith", "John Smith"}
  
  //add more legendary customers dynamically 
  customers = append(customers, "Ben Spain")
  customers = append(customers, "Aleem Janmohamed")
  customers = append(customers, "Jamie le Notre")
  customers = append(customers, "Victor Savkov")
  customers = append(customers, "P The Admin")
  customers = append(customers, "Adrian Oprea")
  customers = append(customers, "Jonathan D")

  //send it back to the caller
  return customers
  
}


```

## Loops 

Loops are used to iterate over collections, lists, arrays etc. <br/>
Let's say we need to loop through our customers

In the `main()` function, we can grab the list of customers and loop them.
In this demo, we'll cover a basic for loop, but there are several approaches to writing loops.

```
//loop forever
for {
  //any code in here will run forever!
  
  fmt.Println("Infinite Loop 1")
	time.Sleep(time.Second)

  //unless we break out the loop like this
  break
}

//loop for x number of loops
for x := 0; x < 10; x++ {

  //any code in here will run 10 times! (unless we break!)
  fmt.Println(customers[x])

}

//loop for ALL our customer

for x, customer := range customers {

  //we can access the "customer" variable in this approach
  customer = customers[x]
  fmt.Println(customer)

  //OR
  //we can use the supplied customer from the loop
  // and silence the x variable, replace it with a _ character
  fmt.Println(customer)
}

```

## Structs

So far so good, however, customer data is not useful as strings. <br/>
Customers can have a firstname, lastname, and more properties. <br/>

For this purpose we'd like to group some variables into a single variable. <br/>
This is what `struct` allows us to do. <br/>
Let's create a `struct` for our customer

Let's create a new `go` file called `customers.go`

```
package main

Customer struct {
  FirstName string
  LastName string
  FullName string
}

```

Let's put it all together:

In `customers.go`, let's create a function to get customers

```
func GetCustomers()(customers []Customer) {

  //we can declare customers like this:
  marcel := 	Customer{ FirstName: "Marcel", LastName: "Dempers" }

	customers = append(customers,
		Customer{ FirstName: "Marcel", LastName: "Dempers" },
		Customer{ FirstName: "Ben", LastName: "Spain" },
		Customer{ FirstName: "Aleem", LastName: "Janmohamed" },
		Customer{ FirstName: "Jamie", LastName: "le Notre" },
		Customer{ FirstName: "Victor", LastName: "Savkov" },
		Customer{ FirstName: "P", LastName: "The Admin" },
		Customer{ FirstName: "Adrian", LastName: "Oprea" },
		Customer{ FirstName: "Jonathan", LastName: "D" },
	)

	return customers

}
```

In `main()` we can now call our shiny new function


```
customers := GetCustomers()

for _, customer := range customers {
  //we can access the "customer" variable in this approach
  fmt.Println(customer)
}
```

# Docker

For our dev environment, we have a simple image using `go` <br/>
We also set a work directory and alias the target as `dev`

This means we can use this container layer as a development environment. <br/>
Later down the track we can add debuggers in here for example. <br/> 
Checkout my debugging video for go: https://youtu.be/kToyI16IFxs  <br/>


## Development environment

```
FROM golang:1.15 as dev

WORKDIR /work
```

## Building our code

```
FROM golang:1.15 as build

WORKDIR /app
COPY ./app/* /app/
RUN go build -o app
```

## The Runtime

```
FROM alpine as runtime 
COPY --from=build /app/app /
CMD ./app
```

## Building the Container

```
docker build . -t customer-app

```

## Running the Container

```
docker run customer-app
```
