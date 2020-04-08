package main

import "github.com/discoveryGo/capter7"

func main() {
	//http.HandleFunc(capter6.PathPrefix, capter6.ApiHandler)
	//http.HandleFunc(capter6.HtmlPrefix, capter6.HtmlHandler)
	//log.Fatal(http.ListenAndServe(":1234", nil))
	capter7.GoroutineMain()
}
