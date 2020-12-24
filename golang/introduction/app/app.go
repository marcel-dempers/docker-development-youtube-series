package main

import (
	"fmt"
)

func main() {
	
	customers := GetCustomers()

	for _, customer := range customers {
  //we can access the "customer" variable in this approach
  fmt.Println(customer)
}

}

func getData() (customers []string) {

  customers = []string{ "Marcel Dempers", "Bob Smith", "John Smith"}
  
	customers = append(customers, "Ben Spain")
  customers = append(customers, "Aleem Janmohamed")
  customers = append(customers, "Jamie le Notre")
  customers = append(customers, "Victor Savkov")
  customers = append(customers, "P The Admin")
  customers = append(customers, "Adrian Oprea")
  customers = append(customers, "Jonathan D")

	for _, customer := range customers {

  	fmt.Println(customer)
	}

	return customers
}