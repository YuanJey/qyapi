package api

import (
	"fmt"
	"github.com/YuanJey/qyapi/http_client"
	"github.com/YuanJey/qyapi/resp"
	"net/url"
)

const (
	getAccessToken    = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	getDepartmentList = "https://qyapi.weixin.qq.com/cgi-bin/department/list?%s"
	getUserList       = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?%s"
	getUserInfo       = "https://qyapi.weixin.qq.com/cgi-bin/auth/getuserinfo?%s"
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
	values := url.Values{}
	if did != "" {
		values.Set("id", did)
	}
	values.Set("access_token", q.AccessToken)
	err := http_client.Get(fmt.Sprintf(getDepartmentList, values.Encode()), nil, &getDepartmentListResp)
	return &getDepartmentListResp, err
}

// GetUserList 获取部门成员
func (q *QYApi) GetUserList(did string) (*resp.GetUserListResp, error) {
	getUserListResp := resp.GetUserListResp{}
	values := url.Values{}
	values.Set("department_id", did)
	values.Set("access_token", q.AccessToken)
	err := http_client.Get(fmt.Sprintf(getUserList, values.Encode()), nil, &getUserListResp)
	return &getUserListResp, err
}

// GetUserInfo 获取授权成员信息
func (q *QYApi) GetUserInfo(code string) (*resp.GetUserInfoResp, error) {
	getUserInfoResp := resp.GetUserInfoResp{}
	values := url.Values{}
	values.Set("access_token", q.AccessToken)
	values.Set("code", code)
	err := http_client.Get(fmt.Sprintf(getUserInfo, values.Encode()), nil, &getUserInfoResp)
	return &getUserInfoResp, err
}
