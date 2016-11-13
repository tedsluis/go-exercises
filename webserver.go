package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
       <html>
           <head>
               <title>Hello, World</title>
           </head>
           <body>
               Hello World!
           </body>
       </html>`,
	)
}
func bye(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
       <html>
           <head>
               <title>Bye</title>
           </head>
           <body>
               Bye!
           </body>
       </html>`,
	)
}
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye", bye)
	http.ListenAndServe(":9000", nil)
}
