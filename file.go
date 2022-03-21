package main

import (
	"flag"
	"fmt"

	"net/http"

	"os"
)

func main() {
	addr := flag.String("addr", ":8080", "listen addr defalut :8080")
	flag.Parse()

	p, _ := os.Getwd()
	http.Handle("/", http.FileServer(http.Dir(p)))
	fmt.Println("serv path:", p)
	fmt.Println(*addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
