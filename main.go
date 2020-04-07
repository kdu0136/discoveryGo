package main

import (
	"fmt"
	"github.com/discoveryGo/capter6"
)

func main() {
	//http.HandleFunc(capter6.PathPrefix, capter6.ApiHandler)
	//http.HandleFunc(capter6.HtmlPrefix, capter6.HtmlHandler)
	//log.Fatal(http.ListenAndServe(":1234", nil))
	fmt.Println(capter6.F())
}
