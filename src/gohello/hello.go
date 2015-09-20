package main

import "fmt"

const N = 10

func main() {
	ch := make([]chan string, N)
	for i := 0; i < N; i++ {
		ch[i] = make(chan string)
		go hello(i, ch[i])
	}
	for i := 0; i < N; i++ {
		val := <-ch[i]
		fmt.Printf("Done %d, from %s\n", i, val)
	}
	fmt.Println("Done")
}

func hello(n int, done chan string) {
	fmt.Printf("Hello from goroutine %d\n", n)
	s := fmt.Sprintf("goroutine %d", n)
	done <- s
}
