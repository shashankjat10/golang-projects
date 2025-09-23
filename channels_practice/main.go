package main

func main() {
	ch := make(chan string, 2)

	go func() {
		ch <- "FNH FOREVER!"
		ch <- "Greatt"
		ch <- "Hello"
	}()
	// ch <- "FNH FOREVER!"
	// ch <- "Greatt"
	// ch <- "Hello"

	// fmt.Printf("%s\n%s", <-ch, <-ch)
}
