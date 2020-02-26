package main

import "net/http"

func handleFunc(w http.ResponseWriter, r *http.Request) {
	msg := `
<h1>hello world!</h1>
`
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/api", handleFunc)
	http.ListenAndServe("0.0.0.0:9000", nil)
}
