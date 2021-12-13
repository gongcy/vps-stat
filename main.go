package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	if err := http.ListenAndServe("127.0.0.1:8081", nil); err != nil {
		log.Fatalf("%v", err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	serviceInfo := getServiceInfo()
	var result string
	if serviceInfo == nil {
		result = "nil"
	} else {
		result = strconv.FormatInt(serviceInfo.DataCounter/1024/1024/1024, 10)
	}

	_, err := fmt.Fprintln(w, result)
	if err != nil {
		fmt.Println(err)
		return
	}
}
