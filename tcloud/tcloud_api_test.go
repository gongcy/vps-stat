package tcloud

import "testing"

func Test_GetNetworkFlowInfo(t *testing.T) {
	resp := GetNetworkFlowInfo()
	t.Logf("%v", resp)
}
