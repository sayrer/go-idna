package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/idna"
)

func main() {
	arg := os.Args[1]
	fmt.Println("name:", arg)

	var p *idna.Profile
	p = idna.New()
	ascii, err := p.ToASCII(arg)
	if err != nil {
    	   log.Fatal(err)
	}
	fmt.Println("Simple Punycode conversion:", ascii)

	p = idna.New(
		idna.MapForLookup(),
		idna.Transitional(true))
	ascii2, err := p.ToASCII(arg)
        if err != nil {                                                                   
           log.Fatal(err)                                                                 
        } 
        fmt.Println("MapForLookup, Transitional:", ascii2)

	p = idna.New(idna.ValidateForRegistration())
	ascii3, err := p.ToASCII(arg)
        if err != nil {                                                                   
           log.Fatal(err)                                                                 
        } 
        fmt.Println("ValidateForRegistration:", ascii3) 
}
