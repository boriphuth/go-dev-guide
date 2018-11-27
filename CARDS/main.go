package main

import (
	"fmt"
	"math/big"
)

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades"
	card = "Five of Diamonds"
	fmt.Println(card)

	// s and t are strings
	s := "this is a string"
	fmt.Println(s)

	// this form isn't needed inside a function body, but works the same.
	var t = "this is another string"
	fmt.Println(t)

	// x is a *big.Int
	x := big.NewInt(0)
	fmt.Println(x)

	// e is a nil error interface
	// we specify the type, because there's no assignment
	// var e error

	// resp is an *http.Response, and err is an error
	// resp, err := http.Get("http://example.com")
}
