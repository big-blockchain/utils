package common

const (
	ErrTypeSuccess       int64 = 0   //没有错误
	ErrTypeServiceNormal int64 = 200 // 服务返回正常
	ErrTypeParamsErr     int64 = 300 // 参数错误
	ErrTypeNoAuth        int64 = 301 // 权限不足
	ErrTypeServerErr     int64 = 500 // 服务器错误
	ErrTypeRequestLimit  int64 = 400 // 请求频繁

	//face相关
	ErrTypeFaceUidError      int64 = 200000 //人脸信息出错
	ErrTypeFaceUidNotFound   int64 = 200001 //未找到人脸信息
	ErrTypeFaceUidRegistered int64 = 200002 //该人脸信息已注册
	ErrTypeFaceNotRegistered int64 = 200003 //该人脸信息未注册
	ErrTypeFaceNotMatch      int64 = 200004 //人脸信息不匹配
	ErrTypeFaceRegisterError int64 = 200005 //人脸注册失败

	//闪验相关
	ErrTypeFlashMobileError int64 = 210000 //闪验服务出错

	//短信消息以及推送消息发送
	ErrTypeMsgCodeSendFailed   int64 = 220000 //短信服务发送失败
	ErrTypeMsgCodeMobileError  int64 = 220001 //手机号格式错误
	ErrTypeMsgCodeContentFalse int64 = 220002 //短信内容格式错误

	//直播相关
	ErrTypeLiveCodeCreateFailed int64 = 230000 //直播间创建失败
	ErrTypeLiveCodeNoLiveAuth   int64 = 230001 //当前账号没有直播权限
	ErrTypeLiveRoomClosed       int64 = 230002 //直播间已关闭
	ErrTypeLiveBalanceNotEnough int64 = 230003 //余额不足
	ErrTypeLiveLightExist       int64 = 230004 //重复爆灯

	//支付相关
	ErrTypePayGoodsInvalid   int64 = 240000 //商品信息失效
	ErrTypePayServiceError   int64 = 240001 //支付服务出错
	ErrTypePayAppleDataError int64 = 240002 //验证信息失效
	ErrTypePaySignError      int64 = 240003 //签名验证错误

	//商城相关
	ErrTypeStockNotEnough int64 = 250000 //库存不足
	ErrTypeGoodsOffSell   int64 = 250001 //商品已下架

)
