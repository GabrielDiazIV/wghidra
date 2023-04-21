package main

func main() {
	ch := make(chan int)

	go worker(ch)
	for _ = range ch {

	}
}

func worker(ch <-chan int) {
	for {
		select {}
	}
}
