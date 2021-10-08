/**
 * @Auth: Nuts
 * @Date: 2021/5/21 2:04 下午
 */
package baiduface

type FilterContentParam struct {
	Text string `json:"text"`
}

type FilterContentResult struct {
	Conclusion     string `json:"conclusion"`     //审核结果，可取值描述：合规、不合规、疑似、审核失败
	ConclusionType int64  `json:"conclusionType"` //审核结果类型，可取值1、2、3、4，分别代表1：合规，2：不合规，3：疑似，4：审核失败
}

func (client *Client) ContentFilter(textContent string) (FilterContentResult, string, error) {
	body := make(map[string]interface{})
	body["text"] = textContent
	res := FilterContentResult{}
	str, err := client.reqFilterPost("text_censor/v2/user_defined", body, &res)

	if err != nil {
		return res, str, err
	}
	return res, str, nil
}
