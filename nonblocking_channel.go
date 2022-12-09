package main

import (
	"fmt"
	"time"
)

func nonBlockingChannel() {
	fmt.Println("start non-blocking channel")

	maxNumberOfDownloads := 5

	numberOfDownloads := 0

	ch := make(chan int, maxNumberOfDownloads)

	for i := 1; i < 15; i++ {
		if numberOfDownloads > maxNumberOfDownloads {
			num := <-ch

			if num%2 == 0 {
				fmt.Printf("number: %d -- is odd\n", num)
			} else {
				fmt.Printf("number: %d -- is even\n", num)
			}

			numberOfDownloads--
		}

		// start go routines
		go multiple(i, ch)
		numberOfDownloads++
	}

	for numberOfDownloads > 0 {
		num := <-ch

		if num%2 == 0 {
			fmt.Printf("number: %d -- is odd\n", num)
		} else {
			fmt.Printf("number: %d -- is even\n", num)
		}

		numberOfDownloads--
	}

	close(ch)

	fmt.Println("finished...")
}

func multiple(num int, ch chan<- int) {
	sum := num * 2
	time.Sleep(1 * time.Second)

	ch <- sum
}
