package campaign

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/campaign"
)

// 广告组更新状态
// 此接口用于更新广告组的状态；
// 对于删除的广告组不允许再更新状态，否则会报错；
// 如果有一个校验不通过，传入的所有的广告组id都不会被更新；
func UpdateStatus(clt *core.SDKClient, accessToken string, req *campaign.UpdateStatusRequest) (*campaign.UpdateStatusResponse, error) {
	var resp campaign.UpdateStatusResponse
	err := clt.Post("2/campaign/update/status/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}