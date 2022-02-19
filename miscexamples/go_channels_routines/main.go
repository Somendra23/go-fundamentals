package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Channels & Subroutines")

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		fmt.Println("Execute link ", link)
		go checklink(link, c)
	}
	fmt.Println("All ruotines fires accepting channel data")
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(links)-1; i++ {
	// 	//this is blocking call
	// 	fmt.Println(<-c)
	// }

	//keep a loop status checker so websites are pinged constantly for status
	// for {

	// 	go checklink(<-c, c)
	// }

	// for l := range c {
	// 	time.Sleep(10 * time.Second)
	// 	go checklink(l, c)
	// }
;;;



	//using function literal

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			checklink(l, c)
		}()
	}
}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("website ", link, "is Down")
		//c <- "Might be down I think"
		c <- link
		return
	}
	fmt.Println("website", link, "is UP")
	fmt.Println("=======================")
	// c <- "Yep its Up " + link
	c <- link
}
