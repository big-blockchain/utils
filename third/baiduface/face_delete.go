/**
 * @Auth: Nuts
 * @Date: 2021/3/3 6:48 下午
 */
package baiduface

func (client *Client) FaceDelete(faceUid,  groupId string) error {
	body := make(map[string]interface{})
	if groupId != "" {
		body["group_id"] = groupId
	}
	if faceUid != "" {
		body["user_id"] = faceUid
	}

	res := FaceUserResult{}
	_, err := client.reqPost("faceset/user/delete", body, &res)

	if err != nil {
		return err
	}
	return nil
}
