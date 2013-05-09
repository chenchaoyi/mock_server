package main

import (
  "flag"
	"fmt"
	"io"
	"net/http"
	"runtime"
)

var _DEBUG bool


func hello(res http.ResponseWriter, req *http.Request) {
  fmt.Print("+")

	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
     <head>
           <title>Hello World</title>
     </head>
     <body>
           Hello World!
     </body>
</html>`,
	)
}

func hello_in_json(res http.ResponseWriter, req *http.Request) {
  fmt.Print("o")
	res.Header().Set(
		"Content-Type",
		"text/json",
	)
	io.WriteString(
		res,
		`{"msg": "hello world"
    , timestamp: "nothing here"}`,
	)
}

func init() {
  flag.BoolVar(&_DEBUG, "debug", false, "Set debug flag")
}

func main() {
	// to init with half of the CPU worthy of threads
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU()/2)

	http.HandleFunc("/html", hello)
	http.HandleFunc("/json", hello_in_json)
	http.ListenAndServe(":9000", nil)
}
