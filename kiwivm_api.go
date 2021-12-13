package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	URL       = "https://api.64clouds.com/v1/getServiceInfo?veid=%s&api_key=%s"
	GetMethod = "GET"
)

func getServiceInfo() *GetServiceInfoResponse {
	getServiceInfoUrl := fmt.Sprintf(URL, os.Getenv("VM_VEID"), os.Getenv("VM_API_KEY"))
	client := &http.Client{}
	req, err := http.NewRequest(GetMethod, getServiceInfoUrl, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Println(string(body))
	var resp GetServiceInfoResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		fmt.Println(err)
		return nil
	}

	return &resp
}
