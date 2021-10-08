/**
 * @Auth: Nuts
 * @Date: 2021/6/21 3:00 下午
用户相关
*/
package zego

/**
创建用户
*/
func (cli ZegoClient) CreateUser(accessToken, userName string, res struct{}) (error, int64) {
	var reqStruct struct {
		SecretId    int    `json:"secret_id"`
		AccessToken string `json:"access_token"`
		UserName    string `json:"user_name"`
	}
	reqStruct.SecretId = cli.SecretId
	reqStruct.AccessToken = accessToken
	reqStruct.UserName = userName
	type User struct {
		Uid      int64  `json:"uid"`
		UserName string `json:"user_name"`
	}
	type Data struct {
		Users []User `json:"users"`
	}
	response := Data{}
	err := requestPost("https://roomkit-api.zego.im/account/user_create", reqStruct, &response)
	if err != nil {
		return err, 0
	}

	return nil, response.Users[0].Uid
}
