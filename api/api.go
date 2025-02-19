package api

import (
	"fmt"
	"github.com/YuanJey/qyapi/http_client"
	"github.com/YuanJey/qyapi/resp"
)

const (
	getAccessToken    = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	getDepartmentList = "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s"
	getUserList       = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%s"
)

type QYApi struct {
	CorpId      string
	CorpSecret  string
	AccessToken string
}

var QYWeChatApi *QYApi

func InitQYApi(corpId, corpSecret string) {
	QYWeChatApi = &QYApi{
		CorpId:     corpId,
		CorpSecret: corpSecret,
	}
}
func (q *QYApi) ReSetAccessToken() error {
	getAccessTokenResp := resp.GetAccessTokenResp{}
	err := http_client.Get(fmt.Sprintf(getAccessToken, q.CorpId, q.CorpSecret), nil, &getAccessTokenResp)
	if getAccessTokenResp.AccessToken != "" {
		q.AccessToken = getAccessTokenResp.AccessToken
	}
	return err
}

// GetDepartmentList 获取部门列表
func (q *QYApi) GetDepartmentList(did string) (*resp.GetDepartmentListResp, error) {
	getDepartmentListResp := resp.GetDepartmentListResp{}
	url := fmt.Sprintf(getDepartmentList, q.AccessToken)
	if did != "" {
		url += fmt.Sprintf("&id=%s", did)
	}
	err := http_client.Get(url, nil, &getDepartmentListResp)
	return &getDepartmentListResp, err
}

// GetUserList 获取部门成员
func (q *QYApi) GetUserList(did string) (*resp.GetUserListResp, error) {
	getUserListResp := resp.GetUserListResp{}
	err := http_client.Get(fmt.Sprintf(getUserList, q.AccessToken, did), nil, &getUserListResp)
	return &getUserListResp, err
}
