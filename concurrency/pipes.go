package concurrency

import "time"

func ReadData(num int) <-chan int {
	receiver := make(chan int)

	go func() {
		var counter int

		for counter < num {
			time.Sleep(time.Duration(counter*100) * time.Millisecond)
			counter++
			receiver <- counter
		}
		close(receiver)
	}()

	return receiver
}

func TransformPipe(input <-chan int, transform func(value int) int) <-chan int {
	output := make(chan int)

	go func() {
		for value := range input {
			output <- transform(value)
		}
		close(output)
	}()

	return output
}

func Double(value int) int {
	return value * 10
}
