package main

import "fmt"

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	s := make([]int, 5, 10)
	fmt.Println(s, len(s), cap(s))

	s2 := make([]int, 0, 0)
	fmt.Println(len(s2), cap(s2))
	s2 = append(s2, 1, 2)
	fmt.Println(len(s2), cap(s2))

	fmt.Println(sum(1, 2, 3, 4, 5))

	ch := make(chan int)
	go receiver(ch)
	i := 0
	for i < 100 {
		ch <- i
		i++
	}
}
