package main

import (
	"fmt"
)

func main() {
	work := []int{1, 2, 3, 4, 5, 6, 7, 8}
	in := generateWork(work)

	out1 := fanOut(in)
	out2 := fanOut(in)
	out3 := fanOut(in)
	out4 := fanOut(in)

	// Simple
	for i := 0; i < len(work); {
		select {
		case value, ok := <-out1:
			if ok {
				i++
				fmt.Println("Output 1 got:", value)
			}

		case value, ok := <-out2:
			if ok {
				i++
				fmt.Println("Output 2 got:", value)
			}

		case value, ok := <-out3:
			if ok {
				i++
				fmt.Println("Output 3 got:", value)
			}

		case value, ok := <-out4:
			if ok {
				i++
				fmt.Println("Output 4 got:", value)
			}

		}
	}

	// Spinlock - более оптимальный варианты, т.к. nil каналы не проверяет select
	for i := 0; i < len(work); i++ {
		select {
		case value, ok := <-out1:
			if !ok {
				i--
				out1 = nil
				continue
			}
			fmt.Println("Output 1 got:", value)
		case value, ok := <-out2:
			if !ok {
				i--
				out2 = nil
				continue
			}
			fmt.Println("Output 2 got:", value)
		case value, ok := <-out3:
			if !ok {
				i--
				out3 = nil
				continue
			}
			fmt.Println("Output 3 got:", value)
		case value, ok := <-out4:
			if !ok {
				i--
				out4 = nil
				continue
			}
			fmt.Println("Output 4 got:", value)
		}
	}
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for data := range in {
			out <- data
		}
	}()

	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
		return
	}()

	return ch
}
