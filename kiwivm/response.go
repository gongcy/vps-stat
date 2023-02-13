package kiwivm

type GetServiceInfoResponse struct {
	VMType                          string        `json:"vm_type"`
	Hostname                        string        `json:"hostname"`
	NodeIP                          string        `json:"node_ip"`
	NodeAlias                       string        `json:"node_alias"`
	NodeLocation                    string        `json:"node_location"`
	NodeLocationID                  string        `json:"node_location_id"`
	NodeDatacenter                  string        `json:"node_datacenter"`
	LocationIpv6Ready               bool          `json:"location_ipv6_ready"`
	Plan                            string        `json:"plan"`
	PlanMonthlyData                 int64         `json:"plan_monthly_data"`
	MonthlyDataMultiplier           int64         `json:"monthly_data_multiplier"`
	PlanDisk                        int64         `json:"plan_disk"`
	PlanRAM                         int64         `json:"plan_ram"`
	PlanSwap                        int64         `json:"plan_swap"`
	PlanMaxIpv6S                    int64         `json:"plan_max_ipv6s"`
	Os                              string        `json:"os"`
	Email                           string        `json:"email"`
	DataCounter                     int64         `json:"data_counter"`
	DataNextReset                   int64         `json:"data_next_reset"`
	IPAddresses                     []string      `json:"ip_addresses"`
	PrivateIPAddresses              []interface{} `json:"private_ip_addresses"`
	IPNullroutes                    []interface{} `json:"ip_nullroutes"`
	Iso1                            interface{}   `json:"iso1"`
	Iso2                            interface{}   `json:"iso2"`
	AvailableIsos                   []string      `json:"available_isos"`
	PlanPrivateNetworkAvailable     bool          `json:"plan_private_network_available"`
	LocationPrivateNetworkAvailable bool          `json:"location_private_network_available"`
	RdnsAPIAvailable                bool          `json:"rdns_api_available"`
	Ptr                             interface{}   `json:"ptr"`
	Suspended                       bool          `json:"suspended"`
	PolicyViolation                 bool          `json:"policy_violation"`
	SuspensionCount                 interface{}   `json:"suspension_count"`
	TotalAbusePoints                int64         `json:"total_abuse_points"`
	MaxAbusePoints                  int64         `json:"max_abuse_points"`
	Error                           int           `json:"error"`
}
