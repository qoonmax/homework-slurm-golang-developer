package homework

func selectMany(channels []chan int64) chan int64 {
	output := make(chan int64)

	for _, ch := range channels {
		go func(c chan int64) {
			for {
				msg, ok := <-c
				if !ok {
					return
				}
				output <- msg
			}
		}(ch)
	}

	return output
}
