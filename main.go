package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"vps-stat/kiwivm"
	"vps-stat/tcloud"
)

func main() {
	http.HandleFunc("/kivivm", KiwivmHandler)
	http.HandleFunc("/tcloud", TCloudHandler)
	if err := http.ListenAndServe("127.0.0.1:8081", nil); err != nil {
		log.Fatalf("%v", err)
	}
}

func KiwivmHandler(w http.ResponseWriter, r *http.Request) {
	serviceInfo := kiwivm.GetServiceInfo()
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

func TCloudHandler(w http.ResponseWriter, r *http.Request) {
	info := tcloud.GetNetworkFlowInfo()
	used := info.Response.InstanceTrafficPackageSet[0].TrafficPackageSet[0].TrafficUsed
	var result string
	if info == nil {
		result = "nil"
	} else {
		result = strconv.FormatInt(used/1024/1024/1024, 10)
	}

	_, err := fmt.Fprintln(w, result)
	if err != nil {
		fmt.Println(err)
		return
	}
}
