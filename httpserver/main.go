package main

import "net/http"

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
	})

	err := http.ListenAndServe("10.10.10.2:8080", nil)
	if err != nil {
		panic(err)
	}
}
