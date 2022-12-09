package main

import (
	"fmt"
	"time"
)

func runProcessesConcurrently() {
	fmt.Println("start non-blocking channel")

	maxNumberOfDownloads := 5

	numberOfDownloads := 0

	ch := make(chan int, maxNumberOfDownloads)

	for i := 1; i < 15; i++ {
		if numberOfDownloads > maxNumberOfDownloads {
			processData(ch, &numberOfDownloads)
		}

		// start go routines
		fmt.Printf("%d - sent\n", i)
		go multiple(i, ch)
		numberOfDownloads++
	}

	for numberOfDownloads > 0 {
		processData(ch, &numberOfDownloads)
	}

	close(ch)

	fmt.Println("finished...")
}

func processData(ch chan int, numberOfDownloads *int) {
	num := <-ch

	if num%2 == 0 {
		fmt.Printf("number: %d -- is odd\n", num)
	} else {
		fmt.Printf("number: %d -- is even\n", num)
	}

	*numberOfDownloads--
}

func multiple(num int, ch chan<- int) {
	sum := num * 2
	time.Sleep(2 * time.Second)

	ch <- sum
}
