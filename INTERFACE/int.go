package main

import "fmt"

type school interface {
	author()
}
type s_details struct {
	name    string
	address string
	pincode int32
}

func (a *s_details) author() {
	a.name = "aesha , \n"
	a.address = "sg road ,\n"
	a.pincode = 798080
}

func main() {
	var hybrid s_details
	var sol school = &hybrid
	sol.author()
	fmt.Println(hybrid)

}

// 1st interface, 2nd structure , func implemented from type func interface & last print value in func main
