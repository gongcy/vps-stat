package tcloud

import "time"

type GetNetworkFlowInfoResponse struct {
	Response struct {
		InstanceTrafficPackageSet []struct {
			InstanceID        string `json:"InstanceId"`
			TrafficPackageSet []struct {
				Deadline                time.Time `json:"Deadline"`
				EndTime                 time.Time `json:"EndTime"`
				StartTime               time.Time `json:"StartTime"`
				Status                  string    `json:"Status"`
				TrafficOverflow         int       `json:"TrafficOverflow"`
				TrafficPackageID        string    `json:"TrafficPackageId"`
				TrafficPackageRemaining int64     `json:"TrafficPackageRemaining"`
				TrafficPackageTotal     int64     `json:"TrafficPackageTotal"`
				TrafficUsed             int64     `json:"TrafficUsed"`
			} `json:"TrafficPackageSet"`
		} `json:"InstanceTrafficPackageSet"`
		RequestID  string `json:"RequestId"`
		TotalCount int    `json:"TotalCount"`
	} `json:"Response"`
}
