package main

import (
	"fmt"
	"go-fundamentals/tempconv"
	"log"
	"os"
	"strconv"
)

func main() {

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatalf("Error parsing commandline arguments")
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celcius(t)
		fmt.Println("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
	// f := tempconv.Fahrenheit(30)
	// fmt.Printf("Package fundamentals %v", tempconv.FToC(f))
	// fmt.Println(tempconv.CToF(tempconv.FreezingC))

}
