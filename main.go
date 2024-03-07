package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	result := <- c
	fmt.Println(result)
	time.Sleep(time.Second * 5)
}

func isSexy(person string, c chan bool) {
	fmt.Println(person)
	time.Sleep(time.Second * 5)
	c <- true
}
