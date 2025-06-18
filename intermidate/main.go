package  main

import "fmt"

func main() {
	counter := func() func(int) int {
		count := 100
		 
		return func(x int) int {
			count -= x
			return count

		}
	}()

	fmt.Println(counter(1))
	fmt.Println(counter(2))
	fmt.Println(counter(3))
}
