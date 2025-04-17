package main

import (
	"fmt"
	//"net/http"
	//"time"
	"os"
	//"os/signal"
	//"syscall"
	"bufio"
)

//func first(w http.ResponseWriter, req *http.Request) {
//	ctx := req.Context()
//	fmt.Println("server: hello handler started")
//	defer fmt.Println("server: hello handerl ended")
//
//select {
//	case <-time.After(10 * time.Second):
//		fmt.Fprintf(w, "hello\n")
//	case <-ctx.Done():
//		err := ctx.Err()
//		fmt.Println("server: ", err)
//		internalError := http.StatusInternalServerError
//		http.Error(w, err.Error(), internalError)
//	}

//}
// z := store.NewMemoryStore()

func NewMemoryStore() *MemoryStore {
	m := MemoryStore{
		items: []int{},
		zz:    make(chan int),
	}
	go m.start()
	return &m
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("timestamp: ")
	scanner.Scan()
	utime := scanner.Text()

	func(m *NewMemoryStore) {
		for {
			select {
			case utime := <-m.zz:
				m.item = append(m.item, utime)
			}
		}
	}

	//	http.HandleFunc("/hello", first)
	//	http.ListenAndServe(":8090", nil)

	//sigs := make(chan os.Signal, 1)

	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//done := make(chan bool, 1)

	//go func() {
	//	sig := <-sigs
	//	fmt.Println()
	//	fmt.Println(sig)
	//	done <- true
	//}()

	//fmt.Println("await signal")
	//<-done
	//fmt.Println("exiting")
}
