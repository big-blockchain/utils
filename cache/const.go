package cache

//redis-key
const (
	RedisKeyOneDaoLive string = "onedao-live:"
	RedisKeyOneDaoPay  string = "onedao-pay:"

	RedisKeyZegoAccessToken string = RedisKeyOneDaoLive + "zego-access-token"
	RedisKeyZegoSdkToken    string = RedisKeyOneDaoLive + "zego-sdk-token:%s" //device_id

	RedisKeyApplePayLock string = RedisKeyOneDaoPay + "apple-pay-lock:%s" //recData

)
