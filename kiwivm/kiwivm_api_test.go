package kiwivm

import "testing"

func Test_GetServiceInfo(t *testing.T) {
	resp := GetServiceInfo()
	t.Logf("%v", resp)
}
