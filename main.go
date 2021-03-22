package main

import (
	"encoding/json"
	"fmt"

	"github.com/amzn/ion-go/ion"
)

type Foo struct {
	A string
	B string
}

type Zop struct {
	Value       Foo
	Annotations []ion.SymbolToken `ion:",annotations"`
}

type Bar struct {
	Barr Zop
}

func main() {
	a := Bar{
		Barr: Zop{
			Value: Foo{
				A: "ABC",
				B: "DEF",
			},
			Annotations: []ion.SymbolToken{ion.NewSymbolTokenFromString("ABC")},
		},
	}

	res, err := ion.MarshalText(a)
	if err != nil {
		panic(err)
	}

	var b Bar
	err = ion.Unmarshal(res, &b)
	if err != nil {
		panic(err)
	}

	// These should print the same output as we have gone from
	// go struct -> ION -> go struct
	e, _ := json.Marshal(a)
	fmt.Printf("Encoded: %s\n", string(e))
	d, _ := json.Marshal(b)
	fmt.Printf("Decoded: %s\n", string(d))
}
