package main

import "fmt"

// do not change the name or the fields of the 'Country' struct!
type Country struct {
	Name     string
	Capital  string
	Currency string
}

func solve() Country {
	// create an instance of the Country struct within the 'germany' variable here.
	// do not forget to add the values mentioned above to its fields too!
	var germany Country

	germany.Name = "Germany"   // add the value 'Germany' to the 'Name' field here
	germany.Capital = "Berlin" // add the value 'Berlin' to the 'Capital' field here
	germany.Currency = "Euro"  // add the value 'Euro' to the 'Currency' field here

	fmt.Sprint()
	return germany // do not delete this line!
}
