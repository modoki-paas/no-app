package main

import (
	"os"
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
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	
	http.ListenAndServe(":" + port, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(message))
	}))
}
