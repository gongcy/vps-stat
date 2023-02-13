package tcloud

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

var (
	VmSid  = os.Getenv("VM_SID")
	VmSkey = os.Getenv("VM_SKEY")
	VmId   = os.Getenv("VM_ID")
)

// GetNetworkFlowInfo
// See: https://cloud.tencent.com/document/api/1207/48681
// See: https://cloud.tencent.com/document/sdk/Go
func GetNetworkFlowInfo() *GetNetworkFlowInfoResponse {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		VmSid,
		VmSkey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := lighthouse.NewClient(credential, "ap-singapore", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := lighthouse.NewDescribeInstancesTrafficPackagesRequest()

	request.InstanceIds = common.StringPtrs([]string{VmId})

	// 返回的resp是一个DescribeInstancesTrafficPackagesResponse的实例，与请求对象对应
	response, err := client.DescribeInstancesTrafficPackages(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 输出json格式的字符串回包
	//fmt.Printf("%s", response.ToJsonString())

	var resp GetNetworkFlowInfoResponse
	fmt.Println(response.ToJsonString())
	if err := json.Unmarshal([]byte(response.ToJsonString()), &resp); err != nil {
		fmt.Println(err)
		return nil
	}

	return &resp
}
