package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	//почему чтение в мейне а не еще в одной горутине?
	//если сделать чтение отдельной горутинрой
	//может ли быть так, что на моменте, когда я читаю из канала ПЕРЕД тем, как что-то напечатать, другая горутина
	//перехватывает поток и записывает новый i туда, откуда уже прочитали, но не успели сделать println,
}