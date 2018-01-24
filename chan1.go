package main
import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total  // send total to c
}

// 2018-01-23
// so: why does 2nd go routine get assigned first?
// because first routine block at sender wo receiver
// second routine runs and blocks at sender wo receiver
// msin routine receives channel most recent is second go routine??
// not sure yet at all
func main() {
	a := []int{7, 2, 8, -9, 4, 0}


	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x + y)
	fmt.Println(a[:len(a)/2])
	u,v := 7,8
	fmt.Println(u,v)
}
