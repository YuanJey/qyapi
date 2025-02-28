package api

import (
	"errors"
	"fmt"
	"github.com/YuanJey/qyapi/http_client"
	"github.com/YuanJey/qyapi/resp"
	"github.com/avast/retry-go/v4"
	"net/url"
)

const (
	getAccessToken    = "%s/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	getDepartmentList = "%s/cgi-bin/department/list?%s"
	getUserList       = "%s/cgi-bin/user/simplelist?%s"
	getUserInfo       = "%s/cgi-bin/auth/getuserinfo?%s"
	getUserInfo2      = "%s/cgi-bin/user/getuserinfo?%s"
)

type QYApi struct {
	CorpId      string
	CorpSecret  string
	Addr        string
	AccessToken string
}

var QYWeChatApi *QYApi

func InitQYApi(corpId, corpSecret, addr string) {
	QYWeChatApi = &QYApi{
		CorpId:     corpId,
		CorpSecret: corpSecret,
		Addr:       addr,
	}
}
func (q *QYApi) ReSetAccessToken() error {
	getAccessTokenResp := resp.GetAccessTokenResp{}
	err := http_client.Get(fmt.Sprintf(getAccessToken, q.Addr, q.CorpId, q.CorpSecret), nil, &getAccessTokenResp)
	if getAccessTokenResp.AccessToken != "" {
		q.AccessToken = getAccessTokenResp.AccessToken
	}
	return err
}

// GetDepartmentList 获取部门列表
func (q *QYApi) GetDepartmentList(did string) (*resp.GetDepartmentListResp, error) {
	getDepartmentListResp := resp.GetDepartmentListResp{}
	err := retry.Do(func() error {
		values := url.Values{}
		if did != "" {
			values.Set("id", did)
		}
		values.Set("access_token", q.AccessToken)
		err := http_client.Get(fmt.Sprintf(getDepartmentList, q.Addr, values.Encode()), nil, &getDepartmentListResp)
		if err != nil || getDepartmentListResp.Errcode != 0 {
			err = errors.New("code err")
			err1 := q.ReSetAccessToken()
			if err1 != nil {
				return err1
			}
			return err
		}
		return err
	}, retry.Attempts(3))
	return &getDepartmentListResp, err
}

// GetUserList 获取部门成员
func (q *QYApi) GetUserList(did string) (*resp.GetUserListResp, error) {
	getUserListResp := resp.GetUserListResp{}
	err := retry.Do(func() error {
		values := url.Values{}
		values.Set("department_id", did)
		values.Set("access_token", q.AccessToken)
		err := http_client.Get(fmt.Sprintf(getUserList, q.Addr, values.Encode()), nil, &getUserListResp)
		if err != nil || getUserListResp.Errcode != 0 {
			err = errors.New("code err")
			err1 := q.ReSetAccessToken()
			if err1 != nil {
				return err1
			}
			return err
		}
		return err
	}, retry.Attempts(3))
	//err := http_client.Get(fmt.Sprintf(getUserList, q.Addr, values.Encode()), nil, &getUserListResp)
	return &getUserListResp, err
}

// GetUserInfo 获取授权成员信息
func (q *QYApi) GetUserInfo(code string) (*resp.GetUserInfoResp, error) {
	getUserInfoResp := resp.GetUserInfoResp{}
	err := retry.Do(func() error {
		values := url.Values{}
		values.Set("access_token", q.AccessToken)
		values.Set("code", code)
		err := http_client.Get(fmt.Sprintf(getUserInfo, q.Addr, values.Encode()), nil, &getUserInfoResp)
		if err != nil || getUserInfoResp.Errcode != 0 {
			err = errors.New("code err")
			err1 := q.ReSetAccessToken()
			if err1 != nil {
				return err1
			}
			return err
		}
		return err
	}, retry.Attempts(3))
	//err := http_client.Get(fmt.Sprintf(getUserInfo, q.Addr, values.Encode()), nil, &getUserInfoResp)
	return &getUserInfoResp, err
}
func (q *QYApi) GetUserInfo2(code string) (*resp.GetUserInfoResp, error) {
	getUserInfoResp := resp.GetUserInfoResp{}
	err := retry.Do(func() error {
		values := url.Values{}
		values.Set("access_token", q.AccessToken)
		values.Set("code", code)
		err := http_client.Get(fmt.Sprintf(getUserInfo2, q.Addr, values.Encode()), nil, &getUserInfoResp)
		if err != nil || getUserInfoResp.Errcode != 0 {
			err = errors.New("code err")
			err1 := q.ReSetAccessToken()
			if err1 != nil {
				return err1
			}
			return err
		}
		return err
	}, retry.Attempts(3))
	//err := http_client.Get(fmt.Sprintf(getUserInfo, q.Addr, values.Encode()), nil, &getUserInfoResp)
	return &getUserInfoResp, err
}
