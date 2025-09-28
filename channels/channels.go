package main
import "fmt"

func main() {
	firstch := make(chan string)
	go func() {
		SendChan(firstch, "hello")
	}()
	readChan(firstch)
}

// rsend only
func SendChan(msgch chan<- string, message string) {
	msgch <- message
}

// send only function 
func readChan(msgch <-chan string) {
	msg := <-msgch
	fmt.Println(msg)
}
