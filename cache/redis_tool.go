package cache

import "fmt"

func RedisZegoAccessTokenKey() string {
	return RedisKeyZegoAccessToken
}

//客户端SDK token
func RedisZegoSdkTokenKey(deviceId string) string {
	return fmt.Sprintf(RedisKeyZegoSdkToken, deviceId)
}

//用户苹果支付订单验证
func RedisApplyPayLockKey(receiptData string) string {
	return fmt.Sprintf(RedisKeyApplePayLock, receiptData)
}
