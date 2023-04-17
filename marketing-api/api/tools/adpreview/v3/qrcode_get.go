package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/tools/adpreview/v3"
)

// QrcodeGet 获取广告预览二维码
func QrcodeGet(clt *core.SDKClient, accessToken string, req *v3.QrcodeGetRequest) (*v3.QrcodeGetResponseData, error) {
	var resp v3.QrcodeGetResponse
	err := clt.Get("v3.0/tools/ad_preview/qrcode_get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
