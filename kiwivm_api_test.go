package main

import "testing"

func Test_GetServiceInfo(t *testing.T) {
	resp := getServiceInfo()
	t.Logf("%v", resp)
}
