package resp

type GetAccessTokenResp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
type GetDepartmentListResp struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Department []struct {
		Id               int      `json:"id"`
		Name             string   `json:"name"`
		NameEn           string   `json:"name_en"`
		DepartmentLeader []string `json:"department_leader"`
		Parentid         int      `json:"parentid"`
		Order            int      `json:"order"`
	} `json:"department"`
}
type GetUserListResp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Userlist []struct {
		Userid     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
		OpenUserid string `json:"open_userid"`
	} `json:"userlist"`
}
