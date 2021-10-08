/**
 * @Auth: Nuts
 * @Date: 2021/5/21 2:04 下午
 */
package baiduface

type FilterVideoResult struct {
	Conclusion     string `json:"conclusion"`     //审核结果，可取值描述：合规、不合规、疑似、审核失败
	ConclusionType int64  `json:"conclusionType"` //审核结果类型，可取值1、2、3、4，分别代表1：合规，2：不合规，3：疑似，4：审核失败
}

func (client *Client) VideoFilter(name, url string) (FilterVideoResult, string, error) {
	body := make(map[string]interface{})
	body["videoUrl"] = url
	body["name"] = name
	body["extId"] = url

	res := FilterVideoResult{}
	str, err := client.reqFilterPost("video_censor/v2/user_defined", body, &res)

	if err != nil {
		return res, str, err
	}
	return res, str, nil
}
