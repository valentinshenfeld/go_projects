package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(i)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}

// package main

// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	for i := 1; i <= 1024; i++ {
// 		go func(j int) {
// 			address := fmt.Sprintf("scanme.nmap.org:%d", j)
// 			conn, err := net.Dial("tcp", address)
// 			if err != nil {
// 				return
// 			}
// 			conn.Close()
// 			fmt.Printf("%d open\n", j)
// 		}(i)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	for i := 1; i <= 1024; i++ {
// 		address := fmt.Sprintf("scanme.nmap.org:%d", i)
// 		conn, err := net.Dial("tcp", address)

// 		if err != nil {
// 			// порт закрыт или отфильтрован
// 			continue
// 		}
// 		conn.Close()
// 		fmt.Printf("%d open\n", i)
// 	}
// }
