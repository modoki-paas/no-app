package main

import (
	"net/http"
)

const message = `<!DOCTYPE html>
<html>
    <head>
        <style>
            .title {
                text-align: center;
            };
        </style>
    </head>
    <body>
        <h1 class="title">No application is deployed</h1>
    </body>
</html>
`

func main() {
	http.ListenAndServe(":80", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(message))
	}))
}
