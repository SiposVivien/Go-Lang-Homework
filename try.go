package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

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

//func NewMemoryStore() *MemoryStore {
//	m := MemoryStore{
//		items: []int{},
//		zz:    make(chan int),
//	}
//	go m.start()
//	return &m
//}//

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

//	scanner := bufio.NewScanner(os.Stdin)
//	fmt.Println("timestamp: ")
//	scanner.Scan()
//	utime := scanner.Text()

//	func(m *NewMemoryStore) {
//	for {
//			select {
//			case utime := <-m.zz:
//			m.item = append(m.item, utime)
//			}
//	}
//}

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
