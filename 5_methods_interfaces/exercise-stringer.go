package main

import "fmt"

/*
Exercise: Stringers
Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
*/

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (address IPAddr) String() string {
	// I did this just because the length is fixed
	// No need for fancy things
	return fmt.Sprintf("%d.%d.%d.%d", address[0], address[1], address[2], address[3])
}

func main() {
	// test := IPAddr{127, 0, 0, 1}
	// fmt.Println(test[0], test[1], test[2], test[3]) // 127 0 0 1
	// fmt.Println(test[4]) // invalid array index 4 (out of bounds for 4-element array)
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
