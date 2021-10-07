package main

import (
	"fmt"
	"sync"
)

func WelcomeMessageWG(wg *sync.WaitGroup) {
	fmt.Println("Welcome to Educative!")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go WelcomeMessageWG(&wg)
	go func() {
		fmt.Println("Hello World!")
		wg.Done()
	}()

	wg.Wait()

}
