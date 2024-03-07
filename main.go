package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- person+" is sexy"
}
