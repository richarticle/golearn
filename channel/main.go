package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

func rangeChannel() {
	input := make(chan string)
	done := make(chan bool)
	go func() {
		for str := range input {
			fmt.Println("Recv", str)
		}
		fmt.Println("input is closed")
		done <- true
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("Enter anything or quit:")
		str := scanner.Text()
		if str == "quit" {
			close(input)
			break
		} else {
			input <- scanner.Text()
		}
	}
	<-done
}

func matchGame() {
	person := []string{"A", "B", "C", "D", "E"}
	match := make(chan string, 1)
	wg := new(sync.WaitGroup)
	wg.Add(len(person))
	for i := range person {
		go func(name string, match chan string, wg *sync.WaitGroup) {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			fmt.Println(name, "'s action")
			select {
			case mate := <-match:
				fmt.Println(name, "is matched by", mate)
			case match <- name:
			}
			wg.Done()
		}(person[i], match, wg)
	}
	wg.Wait()
	select {
	case alone := <-match:
		fmt.Println(alone, "is alone")
	default:
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().Unix())

	rangeChannel()
	matchGame()
}
